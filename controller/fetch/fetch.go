package controller_fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	models_fetch "github.com/sureshtamrakar/web-scraper/models/fetch"

	"github.com/sureshtamrakar/web-scraper/util"
)

type FetchRequest struct {
	URL string `json:"url"`
}

func Create(c *gin.Context) {
	if !util.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	var req FetchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	val := util.Scrap(req.URL)
	value, err := json.Marshal(val.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not Encode to JSON")
		return

	}
	urlval, _ := url.Parse(req.URL)
	err = models_fetch.Create(urlval.Host, string(value))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not Save Value")
		return
	}
}

func List(c *gin.Context) {
	if !util.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)

	}
	val, err := models_fetch.LoadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, val)
}

func Get(c *gin.Context) {
	if !util.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)
	}
	url := c.Param("url")
	val, err := models_fetch.Load(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "URL not found")
		return
	}
	c.JSON(http.StatusOK, val)
	return

}

func Date(c *gin.Context) {
	if !util.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	start := c.Query("start")
	last := c.Query("end")
	val, err := models_fetch.Search(start, last)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Date Range Not Found")
		return
	}
	c.JSON(http.StatusOK, val)
	return

}

func Paginate(c *gin.Context) {
	if !util.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)
	}
	url := c.Param("id")
	i1, err := strconv.Atoi(url)
	if err == nil {
		fmt.Println(i1)
	}
	val, err := models_fetch.Paginate(i1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, val)
	return
}
