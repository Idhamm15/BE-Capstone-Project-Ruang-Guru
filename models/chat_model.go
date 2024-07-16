package models

type ChatRequest struct {
    Text string `json:"text"`
}

type Part struct {
    Text string `json:"text"`
}

type Message struct {
    Role  string `json:"role"`
    Parts []Part `json:"parts"`
}

type ChatHistory struct {
    Messages []Message `json:"messages"`
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
