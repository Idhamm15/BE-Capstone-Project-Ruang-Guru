package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "be_capstone/models"
    "be_capstone/services"
)

type ChatHandler struct {
    chatService *services.ChatService
}

func NewChatHandler() *ChatHandler {
    return &ChatHandler{
        chatService: services.NewChatService(),
    }
}

func (h *ChatHandler) HandleChat(c *gin.Context) {
    var req models.ChatRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    resp, err := h.chatService.GenerateResponse(&req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}
