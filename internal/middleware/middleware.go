package middleware

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"net/http"
	"rate-limiter/internal/rules"

	"github.com/gin-gonic/gin"
)

func RateLimit(c *gin.Context) {
	var userType string
	if val, exists := c.Get("user-type"); exists {
		userType = val.(string)
	}
	if userType == "" {
		userType = "default"
	}
	tokenBucket := rules.GetBucket(GetClientIdentifier(c), userType)
	isAllowed, t := tokenBucket.Request(5)
	if !isAllowed {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message":        "Try again after some time",
			"tokenRemaining": t,
		})
		return
	}
	c.Next()
}

func GetClientIdentifier(c *gin.Context) string {
	ip := c.ClientIP()
	url := c.Request.URL.Path
	data := fmt.Sprintf("%s-%s", ip, url)
	h := md5.Sum([]byte(data))
	hash := new(big.Int).SetBytes(h[:]).Text(62)
	return hash
}
