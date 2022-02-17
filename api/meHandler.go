package api

import (
	"net/http"
	"strconv"
	"webhooks/lib"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const VALIDATE_TOKEN_URL string = "https://accounts.livechat.com/info"

func meHandler(secret string) func(*gin.Context) {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		if auth == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		response, err := resty.New().R().SetHeader("Authorization", "Bearer "+auth).Post(VALIDATE_TOKEN_URL)

		if err != nil {
			c.AbortWithStatus(http.StatusBadGateway)
			return
		}

		if response.StatusCode() != 200 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		data, _ := gabs.ParseJSON(response.Body())

		licenseId := data.Path("license_id").Data().(float64)

		c.JSON(200, gin.H{
			"username": licenseId,
			"password": lib.GeneratePassword(strconv.FormatFloat(licenseId, 'f', 0, 64), secret)},
		)
	}
}
