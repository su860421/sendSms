package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 建立一個新的 GIN 路由器
	router := gin.Default()

	// 設定路由處理函式
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.POST("/send-message", func(c *gin.Context) {
		var request struct {
			Message string `json:"message" binding:"required"`
			Email   string `json:"email" binding:"required"`
		}

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("無法載入 .env 檔案")
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			fmt.Println("Request body:", c.Request.Body)
			fmt.Println("Error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 對訊息字串進行驗證
		if len(request.Message) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "訊息不能為空"})
			return
		}

		os.Setenv("AWS_ACCESS_KEY_ID", os.Getenv("CLOUD_ACCESS_KEY_ID"))
		os.Setenv("AWS_SECRET_ACCESS_KEY", os.Getenv("CLOUD_SECRET_ACCESS_KEY"))
		// 建立一個新的 AWS 會話
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"), // 替換為你所需的 AWS 區域
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 建立一個新的 AWS SNS 客戶端
		svc := sns.New(sess)
		// 將電子郵件訊息發佈到指定的主題
		result, err := svc.Publish(&sns.PublishInput{
			Message:  aws.String(request.Message),
			Subject:  aws.String("新郵件"),
			TopicArn: aws.String(os.Getenv("TOPIC_ARN")), // 替換為你的 SNS 主題 ARN
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 記錄 SNS 返回的訊息 ID
		fmt.Println("訊息 ID:", *result.MessageId)

		c.JSON(http.StatusOK, gin.H{"message": "訊息接收並成功處理"})
	})

	// 啟動伺服器
	router.Run(":8080")
}
