package main

import (
	"fmt"
	"net/http"

	"chensj.com/studygo/go_mod_test/pkg/setting"

	"chensj.com/studygo/go_mod_test/models"
	"chensj.com/studygo/go_mod_test/routers"
)

func main() {
	setting.Init()
	models.Init()
	fmt.Println("setting RunMode is ", setting.RunMode)
	router := routers.InitRouter()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run() // listen and serve on 0.0.0.0:8080
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
