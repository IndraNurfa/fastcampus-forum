package main

import (
	"log"

	"github.com/IndraNurfa/fastcampus/internal/configs"
	"github.com/IndraNurfa/fastcampus/internal/handlers/memberships"
	membershipRepo "github.com/IndraNurfa/fastcampus/internal/repository/memberships"
	membershipSvc "github.com/IndraNurfa/fastcampus/internal/service/memberships"
	"github.com/IndraNurfa/fastcampus/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal koneksi ke database", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(membershipRepo)

	membershipsHandler := memberships.NewHandler(r, membershipService)
	membershipsHandler.RegisterRoutes()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
