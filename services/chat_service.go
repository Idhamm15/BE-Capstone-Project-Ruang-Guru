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

type ChatService struct{}

func NewChatService() *ChatService {
    return &ChatService{}
}

func (s *ChatService) GenerateResponse(req *models.ChatRequest) (*models.Content, error) {
    apiKey := os.Getenv("API_KEY_GEMINI")
    url := "https://generativelanguage.googleapis.com/v1/models/gemini-pro:generateContent?key=" + apiKey

    jsonData := map[string]interface{}{
        "contents": []map[string]interface{}{
            {
                "role": "user",
                "parts": []map[string]string{
                    {"text": req.Text},
                },
            },
        },
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

    return &apiResp.Candidates[0].Content, nil
}
