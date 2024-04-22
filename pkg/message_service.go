package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

type MessageService struct {
	// 在這裡定義你的服務屬性
}

func NewMessageService() *MessageService {
	return &MessageService{
		// 初始化你的服務屬性
	}
}

func (ms *MessageService) SendMessage(message string, mobileNumber int, smsType int) {
	var wg sync.WaitGroup

	smbody := url.QueryEscape(message)
	// 在這裡實作發送訊息的邏輯
	// 構建請求的 payload
	request := map[string]interface{}{
		"username": "your_username",
		"password": "your_password",
		"dstaddr":  mobileNumber,
		"smbody":   smbody,
	}

	wg.Add(1)
	go sendSms(request, &wg)
	wg.Wait()
	fmt.Println("Message sent successfully")
}

func sendSms(request map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var url = "https://api.omgms.com.tw/b2c/mtk/SmSend"

	// 將請求的 payload 轉換為 JSON
	body, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling request:", err)
		return
	}

	// 創建一個新的 HTTP 請求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// 發送 HTTP 請求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	fmt.Println("Message sent successfully")
}
