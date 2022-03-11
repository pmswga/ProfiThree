package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	"github.com/swaggo/gin-swagger"        // gin-swagger middleware
	"net/http"
	_ "src/docs"
	"src/internal/prize"
	"src/internal/promo"
	"strconv"
)

type storeServer struct {
	prizeStore *prize.PrizeStore
	promoStore *promo.PromoStore
}

type RequestPromo struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type RequestPrize struct {
	Description string `json:"description" binding:"required"`
}

// @title Products_Store
// @version 1.0

// @BasePath /

// createProductHandler godoc
// @Summary Create new promo
// @Produce json
// @Accept json
// @Param json_body body main.RequestPromo true "json body"
// @Success 200 {object} promo.ResponsePromo
// @Failure 400 {string} string "Field validation failed on the 'required' tag"
// @Failure 400 {string} string "There is no category with id"
// @Router /promo [post]
func (s storeServer) createPromoHandler(context *gin.Context) {
	var rp RequestPromo

	if err := context.ShouldBindJSON(&rp); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	p := s.promoStore.CreatePromo(rp.Name, rp.Description)

	context.JSON(http.StatusOK, p)
}

// getAllPromoHandler godoc
// @Summary Retrieves all promo
// @Produce json
// @Param query query string false "Query string"
// @Success 200 {object} []promo.ResponsePromo
// @Failure 400 {string} string "There is no matches with query"
// @Router /promo [get]
func (s storeServer) getAllPromoHandler(context *gin.Context) {
	query := context.DefaultQuery("query", "")
	if query == "" {
		products := s.promoStore.GetProducts()
		context.JSON(http.StatusOK, products)
	}
}

// getPromoHandler godoc
// @Summary Retrieves promo by id
// @Produce json
// @Param id path integer true "Product ID"
// @Success 200 {object} promo.ResponsePromo
// @Failure 400 {string} string "There is no product with id"
// @Router /promo/{id} [get]
func (s storeServer) getPromoHandler(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	promo, err := s.promoStore.GetPromo(id)

	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, promo)
}

// deletePromoHandler godoc
// @Summary Delete promo
// @Param id path integer true "Product ID"
// @Success 200 {string} string "Product with id was deleted"
// @Failure 400 {string} string "There is no product with id"
// @Router /promo/{id} [delete]
func (s storeServer) deletePromoHandler(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	if err = s.promoStore.DeletePromo(id); err != nil {
		context.String(http.StatusBadRequest, err.Error())
	} else {
		context.String(http.StatusOK, fmt.Sprintf("Promo with id=%d was deleted", id))
	}
}

// createPrizeHandler godoc
// @Summary Create new prize
// @Produce json
// @Accept json
// @Param json_body body main.RequestPrize true "json body"
// @Success 200 {object} prize.ResponsePrize
// @Failure 400 {string} string "Field validation failed on the 'required' tag"
// @Router /prize [post]
func (s storeServer) createPrizeHandler(context *gin.Context) {
	var rp RequestPrize

	if err := context.ShouldBindJSON(&rp); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	prize := s.prizeStore.CreatePrize(rp.Description)
	context.JSON(http.StatusOK, prize)
}

// getAllPrizesHandler godoc
// @Summary Retrieves all prizes
// @Produce json
// @Param query query string false "Query string"
// @Success 200 {object} []prize.ResponsePrize
// @Failure 400 {string} string "There is no matches with query"
// @Router /prize [get]
func (s storeServer) getAllPrizesHandler(context *gin.Context) {
	query := context.DefaultQuery("query", "")
	if query == "" {
		prizes := s.prizeStore.GetPrizes()
		context.JSON(http.StatusOK, prizes)
	}
}

func NewStoreServer() *storeServer {
	prizeStore := prize.New()
	promoStore := promo.New()

	return &storeServer{
		prizeStore: prizeStore,
		promoStore: promoStore,
	}
}

func main() {
	router := gin.Default()
	server := NewStoreServer()

	router.GET("/promo", server.getAllPromoHandler)
	router.POST("/promo", server.createPromoHandler)
	router.GET("/promo/:id", server.getPromoHandler)
	router.DELETE("/promo/:id", server.deletePromoHandler)

	router.GET("/prize", server.getAllPrizesHandler)
	router.POST("/prize", server.createPrizeHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8080")
}
