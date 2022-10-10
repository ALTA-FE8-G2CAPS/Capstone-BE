package main

import (
	"capstone-project/config"
	"capstone-project/factory"
	modelSchedule "capstone-project/features/schedule/data"
	"capstone-project/migration"
	"capstone-project/utils/database/mysql"

	"time"

	"fmt"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitMysqlDB(cfg)
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	gmt, _ := time.LoadLocation("Asia/Jakarta")
	s := gocron.NewScheduler(gmt)
	s.Every(1).Day().At("00:00").Do(
		func() {
			// Select
			var dataSchedule []modelSchedule.ScheduleDetail
			tx_select := db.Where("status_schedule = ?", "booked").Find(&dataSchedule)

			if tx_select.Error != nil {
				fmt.Println(tx_select.Error)
			} else {
				fmt.Println("berhasil mencari schedule booked")
				fmt.Println(dataSchedule)
			}

			// Update status
			for _, v := range dataSchedule {
				var data modelSchedule.ScheduleDetail
				tx_selectSchedule := db.Where("id = ?", v.ID).Find(&data)

				if tx_selectSchedule.Error != nil {
					fmt.Println("gagal mencari data")
				} else {
					fmt.Println("berhasil mencari data")
				}

				data.Status_schedule = "Available"
				tx_update := db.Model(&modelSchedule.ScheduleDetail{}).Where("id = ?", data.ID).Updates(&data)

				if tx_update.Error != nil {
					fmt.Println("gagal update status schedule")
				} else {
					fmt.Println("update status schedule berjalan")
				}
			}
		},
	)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	migration.InitMigrate(db)
	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))

}
