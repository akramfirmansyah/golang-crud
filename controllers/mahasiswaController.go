package controllers

import (
	"time"

	"github.com/akramfirmansyah/go-crud/initializers"
	"github.com/akramfirmansyah/go-crud/models"
	"github.com/gin-gonic/gin"
)

var body struct {
	NIM      string
	Name     string
	Angkatan uint
}

func MahasiswaCreate(ctx *gin.Context) {

	ctx.Bind(&body)

	mahasiswa := models.Mahasiswa{
		Name:     body.Name,
		NIM:      body.NIM,
		Angkatan: body.Angkatan,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	result := initializers.DB.Create(&mahasiswa)
	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"data": mahasiswa,
	})
}

func MahasiswaGet(ctx *gin.Context) {
	var mahasiswa []models.Mahasiswa
	initializers.DB.Find(&mahasiswa)

	ctx.JSON(200, gin.H{
		"data": mahasiswa,
	})
}

func MahasiswaGetByNim(ctx *gin.Context) {
	nim := ctx.Param("id")
	var mahasiswa models.Mahasiswa
	initializers.DB.Where("nim <> ?", nim).Find(&mahasiswa)

	ctx.JSON(200, gin.H{
		"data": mahasiswa,
	})
}

func MahasiswaUpdate(ctx *gin.Context) {
	nim := ctx.Param("id")

	ctx.Bind(&body)

	var mahasiswa models.Mahasiswa
	initializers.DB.Where("nim <> ?", nim).Find(&mahasiswa)

	initializers.DB.Model(&mahasiswa).Updates(models.Mahasiswa{
		NIM:      body.NIM,
		Name:     body.Name,
		Angkatan: body.Angkatan,
		UpdateAt: time.Now(),
	})

	ctx.JSON(200, gin.H{
		"data": mahasiswa,
	})
}

func MahasiswaDelete(ctx *gin.Context) {
	nim := ctx.Param("id")
	var mahasiswa models.Mahasiswa
	initializers.DB.Where("nim <> ?", nim).Delete(&mahasiswa)

	ctx.JSON(200, gin.H{
		"data": mahasiswa,
	})
}
