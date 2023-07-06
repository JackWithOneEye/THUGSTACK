package main

import (
	"log"
	"net/http"
	"sro/thug-stack/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.NewDB()

	router := gin.Default()
	router.StaticFS("/dist", http.Dir("./dist"))
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "layout.html", nil)
	})

	router.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	router.GET("/frameworks", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "frameworks.html", db.ListFrameworks())
	})

	router.GET("/framework/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		fw, err := db.Framework(uint16(id))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.HTML(http.StatusOK, "framework-detail.html", gin.H{
			"Framework": fw,
			"ReadOnly":  true,
			"Buttons": map[string]bool{
				"Back": true,
				"Edit": true,
			},
		})
	})

	router.GET("/framework/:id/edit", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		fw, err := db.Framework(uint16(id))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.HTML(http.StatusOK, "framework-detail.html", gin.H{
			"Framework": fw,
			"ReadOnly":  false,
			"Buttons": map[string]bool{
				"Back":   true,
				"Cancel": true,
				"Save":   true,
			},
			"Method": http.MethodPatch,
		})
	})

	router.GET("/framework/add", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "framework-detail.html", gin.H{
			"Framework": database.Framework{},
			"ReadOnly":  false,
			"Buttons": map[string]bool{
				"Back": true,
				"Save": true,
			},
			"Method": http.MethodPost,
		})
	})

	router.POST("/framework", func(ctx *gin.Context) {
		var fw database.Framework
		err := ctx.Bind(&fw)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		newFw := db.CreateFramework(fw)
		ctx.HTML(http.StatusOK, "framework-detail.html", gin.H{
			"Framework": newFw,
			"ReadOnly":  true,
			"Buttons": map[string]bool{
				"Back": true,
				"Edit": true,
			},
		})
	})

	router.PATCH("/framework/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		var fw database.Framework
		err = ctx.Bind(&fw)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		updatedFw, err := db.UpdateFramework(uint16(id), fw)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.HTML(http.StatusOK, "framework-detail.html", gin.H{
			"Framework": updatedFw,
			"ReadOnly":  true,
			"Buttons": map[string]bool{
				"Back": true,
				"Edit": true,
			},
		})
	})

	router.DELETE("/framework/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err = db.DeleteFramework(uint16(id))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
	})

	log.Println("THUGSTACK running at http://localhost:8090")

	err := router.Run(":8090")
	if err != nil {
		log.Fatal(err)
	}
}
