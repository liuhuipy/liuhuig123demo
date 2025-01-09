package service

import (
	"context"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"liuhuig123demo/conf"
	"liuhuig123demo/internal/dto"
	"liuhuig123demo/internal/model"
	"liuhuig123demo/internal/model/query"
	"net/http"
	"strconv"
	"time"
)

const consumerTaskCacheKey = "hotel_task_consumer:%s"

const (
	priceDivClass = "div._1jo4hgw"
)

type hotelService struct {
	client *http.Client
}

var HotelService *hotelService

func init() {
	HotelService = &hotelService{}
}

func (srv *hotelService) HandleTask(ctx context.Context, task *dto.Task) error {
	// 根据taskId判断是否已经处理过task
	if srv.checkExistedHandleTask(ctx, task) {
		return nil
	}

	// 获取页面数据并解析
	hotelInfos, err := srv.GetTaskHotelInfos(ctx, task)
	if err != nil {
		return err
	}

	// 入库
	hotels := make([]*model.HotelInfo, 0, len(hotelInfos))
	for _, hotelInfo := range hotelInfos {
		hotels = append(hotels, &model.HotelInfo{
			Name:             hotelInfo.HotelName,
			Star:             hotelInfo.Star,
			Price:            hotelInfo.Price,
			PriceBeforeTaxes: hotelInfo.PriceBeforeTaxes,
			CheckInDate:      hotelInfo.CheckInDate,
			CheckOutDate:     hotelInfo.CheckOutDate,
			Guests:           hotelInfo.Guests,
		})
	}

	q := query.HotelInfo
	err = q.WithContext(ctx).CreateInBatches(hotels, len(hotels))
	if err != nil {
		return err
	}
	return nil
}

func (srv *hotelService) GetTaskHotelInfos(ctx context.Context, task *dto.Task) ([]*dto.Hotel, error) {
	req, err := http.NewRequest(http.MethodGet, task.Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = task.Headers
	resp, err := srv.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("")
	}

	// todo: 解析页面
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	selectors := doc.Find("listing-card-subtitle")

	hotelInfos := make([]*dto.Hotel, 0, len(selectors.Nodes))
	for _, node := range selectors.Nodes {
		hotelInfo := &dto.Hotel{}
		for _, childNode := range node.Attr {
			switch childNode.Key {
			case priceDivClass:
				price, err := strconv.ParseFloat(childNode.Val, 64)
				if err != nil {
					continue
				}
				hotelInfo.Price = price
			}
		}
		hotelInfos = append(hotelInfos, hotelInfo)
	}
	return hotelInfos, nil
}

func (srv *hotelService) checkExistedHandleTask(ctx context.Context, task *dto.Task) bool {
	return conf.RedisCli.SetNX(ctx, consumerTaskCacheKey, task.TaskId, time.Hour*24).Val()
}
