package service

import (
	"context"
	"encoding/json"
	"errors"
	"liuhuig123demo/internal/dto"
	"net/http"
)

type hotelService struct {
	client *http.Client
}

var HotelService *hotelService

func init() {
	HotelService = &hotelService{}
}

func (srv *hotelService) SyncTasks(ctx context.Context, taskFile string) error {
	return nil
}

func (srv *hotelService) HandleTask(ctx context.Context, task *dto.Task) error {
	req, err := http.NewRequest(http.MethodGet, task.Url, nil)
	if err != nil {
		return err
	}

	req.Header = task.Headers
	resp, err := srv.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("")
	}

	var result dto.Hotel
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}
