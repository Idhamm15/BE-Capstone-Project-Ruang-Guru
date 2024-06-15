package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "be_capstone/models"
)

type ChatService struct {
    History *models.ChatHistory
}

func NewChatService() *ChatService {
    // Initialize with predefined history
    initialHistory := &models.ChatHistory{
        Messages: []models.Message{
            {
                Role:  "user",
                Parts: []models.Part{{Text: "Namamu siapa"}},
            },
            {
                Role:  "model",
                Parts: []models.Part{{Text: "saya adalah Asisten Kesehatan Anda"}},
            },
        },
    }
    return &ChatService{
        History: initialHistory,
    }
}

func (s *ChatService) GenerateResponse(req *models.ChatRequest) (*models.Content, error) {
    apiKey := os.Getenv("API_KEY_GEMINI")
    url := "https://generativelanguage.googleapis.com/v1/models/gemini-pro:generateContent?key=" + apiKey

    // Add user message to history
    userMessage := models.Message{
        Role:  "user",
        Parts: []models.Part{{Text: req.Text}},
    }
    s.History.Messages = append(s.History.Messages, userMessage)

    jsonData := map[string]interface{}{
        "contents": s.History.Messages,
    }
    jsonValue, _ := json.Marshal(jsonData)

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    var apiResp models.APIResponse
    if err := json.Unmarshal(body, &apiResp); err != nil {
        return nil, err
    }

    if len(apiResp.Candidates) == 0 {
        return nil, fmt.Errorf("no candidates in response")
    }

    // Customize the response content
    customizedContent := s.CustomizeResponse(apiResp.Candidates[0].Content)

    // Add model response to history
    modelMessage := models.Message{
        Role:  "model",
        Parts: customizedContent.Parts,
    }
    s.History.Messages = append(s.History.Messages, modelMessage)

    return &customizedContent, nil
}

func (s *ChatService) CustomizeResponse(content models.Content) models.Content {
    // Custom logic to modify the response
    for i, part := range content.Parts {
        if part.Text == "Saya tidak punya nama, saya adalah model bahasa yang dikembangkan oleh Google." {
            content.Parts[i].Text = "Saya adalah AI yang dibuat oleh Google, tetapi Anda bisa memanggil saya Asisten Kesehatan Anda."
        }
    }
    return content
}
