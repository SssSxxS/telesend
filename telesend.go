package telesend

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const telegramApiBaseUrl = "https://api.telegram.org/bot%s/sendMessage"

// sendMessageInternal is a helper function that handles the actual message sending logic
// to avoid code duplication between the standalone function and the Client method.
func sendMessageInternal(botToken, chatID, text, parseMode string) error {
	apiURL := fmt.Sprintf(telegramApiBaseUrl, botToken)

	data := url.Values{
		"chat_id": {chatID},
		"text":    {text},
	}
	if parseMode != "" {
		data.Set("parse_mode", parseMode)
	}

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error from Telegram API: status %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

// SendMessage sends a message to a Telegram chat using the provided bot token.
//
// It accepts a bot token, chat ID, message text, and optional parse mode ("HTML", "MarkdownV2", or "").
// This is a standalone function that can be used without creating a Client instance.
//
// Example usage:
//
//	err := telesend.SendMessage("YOUR_BOT_TOKEN", "CHAT_ID", "Hello from Telegram!", "")
//	if err != nil {
//	    log.Fatalf("Failed to send message: %v", err)
//	}
//
//	// With HTML formatting
//	err = telesend.SendMessage("YOUR_BOT_TOKEN", "CHAT_ID", "<b>Bold text</b> in Telegram", "HTML")
//	if err != nil {
//	    log.Fatalf("Failed to send message: %v", err)
//	}
//
// Note: Consider using the Client method for better organization when sending multiple messages.
func SendMessage(botToken, chatID, text, parseMode string) error {
	return sendMessageInternal(botToken, chatID, text, parseMode)
}

// Client represents a client for interacting with the Telegram Bot API.
//
// It stores the bot token to avoid passing it with each request, making it more
// convenient for sending multiple messages to different chats.
type Client struct {
	botToken string
}

// NewClient creates a new Client instance with the provided bot token.
//
// Example usage:
//
//	client := telesend.NewClient("YOUR_BOT_TOKEN")
//	// Now you can use client.SendMessage() without specifying the token each time
func NewClient(botToken string) *Client {
	return &Client{
		botToken: botToken,
	}
}

// SendMessage sends a message to a Telegram chat using the client's bot token.
//
// It accepts a chat ID, message text, and optional parse mode ("HTML", "MarkdownV2", or "").
// This method uses the bot token stored in the Client instance.
//
// Example usage:
//
//	client := telesend.NewClient("YOUR_BOT_TOKEN")
//
//	// Send a simple text message
//	err := client.SendMessage("CHAT_ID", "Hello from Telegram!", "")
//	if err != nil {
//	    log.Fatalf("Failed to send message: %v", err)
//	}
//
//	// Send a message with Markdown formatting
//	err = client.SendMessage("CHAT_ID", "*Bold text* in Telegram", "MarkdownV2")
//	if err != nil {
//	    log.Fatalf("Failed to send message: %v", err)
//	}
func (c *Client) SendMessage(chatID, text, parseMode string) error {
	return sendMessageInternal(c.botToken, chatID, text, parseMode)
}
