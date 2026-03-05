# Promptdis Python SDK Example

Minimal CLI demonstrating core SDK operations: fetch, render, cache, and error handling.

## Prerequisites

- Python 3.10+
- Promptdis SDK installed locally:
  ```bash
  cd sdk-py
  pip install -e .
  # or
  uv pip install -e .
  ```
- A running Promptdis server with at least one prompt (`demo/myapp/greeting`)

## Run

```bash
PROMPTDIS_URL=http://localhost:8000 \
PROMPTDIS_API_KEY=pm_test_... \
python examples/main.py
```

## What it demonstrates

1. **Fetch by name** — `client.get_by_name("demo", "myapp", "greeting")`
2. **Fetch by ID** — `client.get(prompt.id)`
3. **Local render** — `prompt.render({"name": "Alice"})` (Jinja2 sandboxed)
4. **Server-side render** — `client.render(prompt.id, {"name": "Bob"})`
5. **Cache stats** — `client.cache_stats()`
6. **Error handling** — catch `NotFoundError` for missing prompts

## Full SDK docs

See [`sdk-py/README.md`](../README.md) for the complete API reference.
