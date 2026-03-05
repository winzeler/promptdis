"""Promptdis Python SDK — example usage."""

import os
import sys

from promptdis import PromptClient, NotFoundError, PromptdisError


def main():
    base_url = os.environ.get("PROMPTDIS_URL")
    api_key = os.environ.get("PROMPTDIS_API_KEY")

    if not base_url or not api_key:
        print("Set PROMPTDIS_URL and PROMPTDIS_API_KEY environment variables.")
        sys.exit(1)

    with PromptClient(base_url=base_url, api_key=api_key) as client:
        # 1. Fetch a prompt by name
        print("--- Fetch by name ---")
        prompt = client.get_by_name("demo", "myapp", "greeting")
        print(f"  ID:          {prompt.id}")
        print(f"  Name:        {prompt.name}")
        print(f"  Version:     {prompt.version}")
        print(f"  Model:       {prompt.model.get('default', 'unknown')}")
        print(f"  Environment: {prompt.environment}")
        print(f"  Body:        {prompt.body[:80]}...")

        # 2. Fetch the same prompt by ID
        print("\n--- Fetch by ID ---")
        same_prompt = client.get(prompt.id)
        print(f"  Name:        {same_prompt.name}")

        # 3. Local render (Jinja2)
        print("\n--- Local render ---")
        rendered = prompt.render({"name": "Alice"})
        print(f"  Rendered:    {rendered[:120]}")

        # 4. Server-side render
        print("\n--- Server-side render ---")
        rendered_server = client.render(prompt.id, {"name": "Bob"})
        print(f"  Rendered:    {rendered_server[:120]}")

        # 5. Cache stats
        print("\n--- Cache stats ---")
        stats = client.cache_stats()
        for key, value in stats.items():
            print(f"  {key}: {value}")

        # 6. Error handling
        print("\n--- Error handling ---")
        try:
            client.get("nonexistent-id")
        except NotFoundError:
            print("  Caught NotFoundError (expected)")
        except PromptdisError as e:
            print(f"  Caught PromptdisError: {e}")


if __name__ == "__main__":
    main()
