# Promptdis JavaScript SDK Example

Minimal CLI demonstrating core SDK operations: fetch, render, cache, and error handling.

## Prerequisites

- Node.js 18+
- A running Promptdis server with at least one prompt (`demo/myapp/greeting`)

## Run

From the `sdk-js/` directory:

```bash
PROMPTDIS_URL=http://localhost:8000 \
PROMPTDIS_API_KEY=pm_test_... \
node examples/main.mjs
```

## What it demonstrates

1. **Fetch by name** — `client.getByName("demo", "myapp", "greeting")`
2. **Fetch by ID** — `client.get(prompt.id)`
3. **Local render** — `client.renderLocal(prompt.body, { name: "Alice" })`
4. **Server-side render** — `client.render(prompt.id, { name: "Bob" })`
5. **Cache stats** — `client.cacheStats()`
6. **Error handling** — catch `NotFoundError` for missing prompts

## Using the published package

To use the published npm package instead of the local source, change the import:

```js
import { PromptClient } from "@promptdis/client-js";
```

## Full SDK docs

See [`sdk-js/README.md`](../README.md) for the complete API reference.
