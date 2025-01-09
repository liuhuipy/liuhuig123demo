package main

import (
	"gorm.io/gen"
	"liuhuig123demo/internal/model"
)

type Querier interface {
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.HotelInfo{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {}, model.HotelInfo{})

	// Generate the code
	g.Execute()
}
