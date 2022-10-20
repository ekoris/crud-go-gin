package news

import (
	"crud/repositories"
	"crud/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchNews(c *gin.Context) {
	var requestObj request.FindParameter
	c.ShouldBindQuery(&requestObj)

	// fmt.Println(requestObj)
	news := repositories.New().Fetch(requestObj)
	c.JSON(http.StatusOK, gin.H{"data": news})
}

func FindNews(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	news := repositories.New().Find(id)
	c.JSON(http.StatusOK, gin.H{"data": news})
}

func CreateNews(c *gin.Context) {
	var input request.NewsRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	news := repositories.New().Create(input)
	c.JSON(http.StatusOK, gin.H{"data": news})
}

func UpdateNews(c *gin.Context) {
	var input request.NewsRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	news := repositories.New().Update(input, id)
	c.JSON(http.StatusOK, gin.H{"data": news})
}

func DeleteNews(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	news := repositories.New().Delete(id)
	c.JSON(http.StatusOK, gin.H{"data": news})
}
