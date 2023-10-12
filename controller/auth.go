package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "time"
    "context"
    "market_place.com/models"
    "market_place.com/config"
		"encoding/json"

)
	type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
	}
	type dataRedis struct {
		Id int `json:"id"`
    Name string `json:"name" `
    Email string `json:"email"`
	}
func GenerateToken(c *gin.Context) {

    var loginData LoginRequest

    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    var user models.User

    if err := config.DB.Where("name = ?", loginData.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if loginData.Password != user.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Name,
        "exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    })
    secretKey := []byte("market_place")
    signedToken, err := token.SignedString(secretKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": signedToken})
		ctx := context.Background()
    client := config.CreateRedisClient()

		dataRedis := dataRedis{Id:user.Id, Name: user.Name, Email: user.Email}
		userDataJSON, err := json.Marshal(dataRedis)
    errRedis := client.Set(ctx, signedToken, userDataJSON, 0).Err()
    if errRedis != nil {
        panic(err)
    }
}

func Authenticate() gin.HandlerFunc {
    ctx := context.Background()
    client := config.CreateRedisClient()
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        value, err := client.Get(ctx, tokenString).Result()
        if err != nil {
            panic(err)
        }
      var userData struct {
            Id int `json:"id"`
            // Add other fields as needed
        }

        if err := json.Unmarshal([]byte(value), &userData); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
            c.Abort()
            return
        }

        // Set the user ID in the Gin context
        c.Set("userID", userData.Id)

        c.Next()
    }
}



