package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

type ChatRequest struct {
    Text string `json:"text"`
}

type Part struct {
    Text string `json:"text"`
}

type Content struct {
    Role  string `json:"role"`
    Parts []Part `json:"parts"`
}

type Candidate struct {
    Content      Content `json:"content"`
    FinishReason string  `json:"finishReason"`
    Index        int     `json:"index"`
}

type APIResponse struct {
    Candidates []Candidate `json:"candidates"`
}

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    r := gin.Default()

    r.POST("/chat", func(c *gin.Context) {
        var req ChatRequest
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

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
            log.Fatalf("Error making API request: %v", err)
        }
        defer resp.Body.Close()

        body, _ := ioutil.ReadAll(resp.Body)

        // Logging the response body for debugging
        log.Printf("Response Body: %s", body)

        var apiResp APIResponse
        if err := json.Unmarshal(body, &apiResp); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response from API"})
            return
        }

        // Check if the API response contains the expected data
        if len(apiResp.Candidates) == 0 {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "No candidates in response", "details": string(body)})
            return
        }

        // Return the first candidate's content
        c.JSON(http.StatusOK, apiResp.Candidates[0].Content)
    })

    r.Run(":8080")
}
