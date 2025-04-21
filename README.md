# Telesend

A simple, zero-dependencies Go library for sending messages via the Telegram Bot API.

## Installation

```bash
go get -u github.com/SssSxxS/telesend
```

## Usage

### Standalone Function

For simple use cases or one-off messages, you can use the standalone `SendMessage` function:

```go
package main

import (
	"log"

	"github.com/SssSxxS/telesend"
)

func main() {
	// Send a simple text message
	err := telesend.SendMessage(
		"YOUR_BOT_TOKEN",
		"CHAT_ID",
		"Hello from Telegram!",
		"", // No special formatting
	)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}

	// Send a message with HTML formatting
	err = telesend.SendMessage(
		"YOUR_BOT_TOKEN",
		"CHAT_ID",
		"<b>Bold text</b> and <i>italic text</i>",
		"HTML",
	)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}
```

### Client-Based Approach

For applications that send multiple messages, it's more efficient to create a `Client` instance:

```go
package main

import (
	"log"

	"github.com/SssSxxS/telesend"
)

func main() {
	// Create a new client with your bot token
	client := telesend.NewClient("YOUR_BOT_TOKEN")

	// Send a simple text message
	err := client.SendMessage(
		"CHAT_ID",
		"Hello from Telegram!",
		"", // No special formatting
	)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}

	// Send a message with MarkdownV2 formatting
	err = client.SendMessage(
		"CHAT_ID",
		"*Bold text* and _italic text_",
		"MarkdownV2",
	)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}
```

## Message Formatting

Telesend supports Telegram's message formatting options:

- `HTML` - Use HTML tags like `<b>`, `<i>`, `<a>`, etc.
- `MarkdownV2` - Use Markdown syntax like `*bold*`, `_italic_`, etc.
- Empty string (`""`) - No special formatting

Refer to the [Telegram Bot API documentation](https://core.telegram.org/bots/api#formatting-options) for more details on formatting options.
