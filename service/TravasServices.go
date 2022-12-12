package service

import "github.com/gin-gonic/gin"

type TravasService interface {
	ProcessRegister() gin.HandlerFunc
	ProcessLogin() gin.HandlerFunc
	GetAllTours() gin.HandlerFunc
	UpdateTour() gin.HandlerFunc
	GetTour() gin.HandlerFunc
	DeleteTour() gin.HandlerFunc
	CreateTour() gin.HandlerFunc
	ProcessCheckOut() gin.HandlerFunc
	BookTour() gin.HandlerFunc
	SelectTour() gin.HandlerFunc
}
