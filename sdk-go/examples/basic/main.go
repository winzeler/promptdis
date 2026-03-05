package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	promptdis "github.com/futureself-app/promptdis-go"
)

func main() {
	baseURL := os.Getenv("PROMPTDIS_URL")
	apiKey := os.Getenv("PROMPTDIS_API_KEY")

	if baseURL == "" || apiKey == "" {
		fmt.Fprintln(os.Stderr, "Set PROMPTDIS_URL and PROMPTDIS_API_KEY environment variables.")
		os.Exit(1)
	}

	client, err := promptdis.NewClient(promptdis.ClientOptions{
		BaseURL: baseURL,
		APIKey:  apiKey,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// 1. Fetch a prompt by name
	fmt.Println("--- Fetch by name ---")
	prompt, err := client.GetByName(ctx, "demo", "myapp", "greeting")
	if err != nil {
		log.Fatalf("GetByName failed: %v", err)
	}
	fmt.Printf("  ID:          %s\n", prompt.ID)
	fmt.Printf("  Name:        %s\n", prompt.Name)
	fmt.Printf("  Version:     %s\n", prompt.Version)
	fmt.Printf("  Model:       %s\n", prompt.ModelDefault("unknown"))
	fmt.Printf("  Environment: %s\n", prompt.Environment)
	body := prompt.Body
	if len(body) > 80 {
		body = body[:80]
	}
	fmt.Printf("  Body:        %s...\n", body)

	// 2. Fetch the same prompt by ID
	fmt.Println("\n--- Fetch by ID ---")
	samePrompt, err := client.Get(ctx, prompt.ID)
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	fmt.Printf("  Name:        %s\n", samePrompt.Name)

	// 3. Local render (basic {{var}} substitution)
	fmt.Println("\n--- Local render ---")
	rendered := client.RenderLocal(prompt.Body, map[string]string{"name": "Alice"})
	if len(rendered) > 120 {
		rendered = rendered[:120]
	}
	fmt.Printf("  Rendered:    %s\n", rendered)

	// 4. Server-side render (full Jinja2)
	fmt.Println("\n--- Server-side render ---")
	result, err := client.Render(ctx, prompt.ID, map[string]interface{}{"name": "Bob"})
	if err != nil {
		log.Fatalf("Render failed: %v", err)
	}
	renderedBody := result.RenderedBody
	if len(renderedBody) > 120 {
		renderedBody = renderedBody[:120]
	}
	fmt.Printf("  Rendered:    %s\n", renderedBody)

	// 5. Cache stats
	fmt.Println("\n--- Cache stats ---")
	stats := client.CacheStats()
	fmt.Printf("  Size:        %d\n", stats.Size)
	fmt.Printf("  Max size:    %d\n", stats.MaxSize)
	fmt.Printf("  TTL:         %v\n", stats.TTL)

	// 6. Error handling
	fmt.Println("\n--- Error handling ---")
	_, err = client.Get(ctx, "nonexistent-id")
	if err != nil {
		if errors.Is(err, promptdis.ErrNotFound) {
			fmt.Println("  Caught ErrNotFound (expected)")
		} else {
			var pe *promptdis.PromptdisError
			if errors.As(err, &pe) {
				fmt.Printf("  Caught PromptdisError: %s\n", pe.Message)
			}
		}
	}
}
