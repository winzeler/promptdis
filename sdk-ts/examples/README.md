# Promptdis TypeScript SDK Example

Minimal CLI demonstrating core SDK operations with full TypeScript types: fetch, render, cache, and error handling.

## Prerequisites

- Node.js 18+
- A running Promptdis server with at least one prompt (`demo/myapp/greeting`)

## Setup

```bash
cd sdk-ts/examples
npm install
```

## Run

```bash
PROMPTDIS_URL=http://localhost:8000 \
PROMPTDIS_API_KEY=pm_test_... \
npm start
```

## Alternative: compile and run

```bash
npx tsc
PROMPTDIS_URL=http://localhost:8000 \
PROMPTDIS_API_KEY=pm_test_... \
node dist/main.js
```

## What it demonstrates

1. **Fetch by name** — `client.getByName("demo", "myapp", "greeting")` with `Prompt` type
2. **Fetch by ID** — `client.get(prompt.id)`
3. **Local render** — `client.renderLocal(prompt.body, { name: "Alice" })`
4. **Server-side render** — `client.render(prompt.id, { name: "Bob" })`
5. **Cache stats** — `client.cacheStats()`
6. **Error handling** — catch `NotFoundError` for missing prompts

## Full SDK docs

See [`sdk-ts/README.md`](../README.md) for the complete API reference.
