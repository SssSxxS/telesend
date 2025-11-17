package telesend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ParseMode string

const (
	ParseModeHTML       ParseMode = "HTML"
	ParseModeMarkdown   ParseMode = "Markdown"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
)

type TelesendClient struct {
	BotToken              string
	ChatID                int64
	ParseMode             ParseMode
	DisableWebPagePreview bool
	DisableNotification   bool
}

type Message struct {
	ChatID                int64  `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool   `json:"disable_notification,omitempty"`
}

func NewClient(botToken string, chatID int64, parseMode ParseMode, disableWebPagePreview, disableNotification bool) *TelesendClient {
	return &TelesendClient{
		BotToken:              botToken,
		ChatID:                chatID,
		ParseMode:             parseMode,
		DisableWebPagePreview: disableWebPagePreview,
		DisableNotification:   disableNotification,
	}
}

func (c *TelesendClient) SendMessage(text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.BotToken)
	msg := Message{
		ChatID:                c.ChatID,
		Text:                  text,
		ParseMode:             string(c.ParseMode),
		DisableWebPagePreview: c.DisableWebPagePreview,
		DisableNotification:   c.DisableNotification,
	}
	reqBody, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}
	httpClient := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	return nil
}
