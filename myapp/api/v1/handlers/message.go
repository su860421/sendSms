package handlers

import (
	"net/http"
)

// MessageHandler 是處理訊息的 handler
func MessageHandler(w http.ResponseWriter, r *http.Request) {
	// 在這裡處理訊息的邏輯

	// 回傳回應
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("訊息處理成功"))
}
