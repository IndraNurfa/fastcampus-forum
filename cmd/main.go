package main

import (
	"log"

	"github.com/IndraNurfa/fastcampus/internal/configs"
	"github.com/IndraNurfa/fastcampus/internal/handlers/memberships"
	"github.com/IndraNurfa/fastcampus/internal/handlers/posts"
	membershipRepo "github.com/IndraNurfa/fastcampus/internal/repository/memberships"
	postRepo "github.com/IndraNurfa/fastcampus/internal/repository/posts"
	membershipSvc "github.com/IndraNurfa/fastcampus/internal/service/memberships"
	postSvc "github.com/IndraNurfa/fastcampus/internal/service/posts"
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

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipsHandler := memberships.NewHandler(r, membershipService)
	membershipsHandler.RegisterRoutes()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoutes()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
