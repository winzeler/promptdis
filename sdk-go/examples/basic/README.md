# Promptdis Go SDK Example

Minimal CLI demonstrating core SDK operations: fetch, render, cache, and error handling.

## Prerequisites

- Go 1.21+
- A running Promptdis server with at least one prompt (`demo/myapp/greeting`)

## Run

From the `sdk-go/examples/basic/` directory:

```bash
PROMPTDIS_URL=http://localhost:8000 \
PROMPTDIS_API_KEY=pm_test_... \
go run .
```

## What it demonstrates

1. **Fetch by name** — `client.GetByName(ctx, "demo", "myapp", "greeting")`
2. **Fetch by ID** — `client.Get(ctx, prompt.ID)`
3. **Local render** — `client.RenderLocal(prompt.Body, map[string]string{"name": "Alice"})`
4. **Server-side render** — `client.Render(ctx, prompt.ID, map[string]interface{}{"name": "Bob"})`
5. **Cache stats** — `client.CacheStats()` with `Size`, `MaxSize`, `TTL`
6. **Error handling** — `errors.Is(err, promptdis.ErrNotFound)` for missing prompts

## Full SDK docs

See [`sdk-go/README.md`](../../README.md) for the complete API reference.
