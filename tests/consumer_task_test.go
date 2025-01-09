package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"liuhuig123demo/internal/dto"
	"liuhuig123demo/internal/service"
	"testing"
)

func Test_GetTaskHotelInfo(t *testing.T) {
	type args struct {
		ctx  context.Context
		task *dto.Task
	}

	tests := []struct {
		name    string
		args    args
		want    *dto.Hotel
		wantErr error
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hotelInfo, err := service.HotelService.GetTaskHotelInfo(tt.args.ctx, tt.args.task)

			assert.True(t, tt.wantErr == err)
			assert.True(t, tt.want == hotelInfo)
		})
	}
}
