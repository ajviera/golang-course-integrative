package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ajviera/golang-course-integrative/src/maths"
	"github.com/gin-gonic/gin"
)

// Response expose
type Response struct {
	Results []interface{} `json:"results"`
}

func main() {
	r := gin.Default()
	r.GET("/categories/:id/prices", getCategory)
	r.Run()
}

func getCategory(c *gin.Context) {
	url := "https://api.mercadolibre.com/sites/MLA/search?category=" + c.Param("id")
	resp, err := http.Get(url)
	if err != nil {
		fail(c, err)
	} else {
		success(c, resp)
	}
}

func success(c *gin.Context, resp *http.Response) {
	max, suggested, min, err := maths.CalculateSuggestedPrice(parsePrices(parseResponce(c, resp)))
	if err != nil {
		fail(c, err)
	} else {
		c.JSON(200, gin.H{
			"max":       max,
			"suggested": suggested,
			"min":       min,
		})
	}
}

func parseResponce(c *gin.Context, resp *http.Response) Response {
	defer resp.Body.Close()
	var data = Response{}
	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(body, &data)
		if err != nil {
			fail(c, err)
		}
	}

	return data
}

func parsePrices(data Response) []float64 {
	var prices []interface{}
	for i := 0; i < len(data.Results); i++ {
		for k, v := range data.Results[i].(map[string]interface{}) {
			if k == "price" {
				prices = append(prices, v)
			}
		}
	}
	return getPrices(prices)
}

func getPrices(prices []interface{}) []float64 {
	var slice []float64
	for _, value := range prices {
		slice = append(slice, value.(float64))
	}
	return slice
}

func fail(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"error": err,
	})
}
