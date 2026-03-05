import { PromptClient, NotFoundError, PromptdisError } from "@promptdis/client";
import type { Prompt } from "@promptdis/client";

const baseUrl = process.env.PROMPTDIS_URL;
const apiKey = process.env.PROMPTDIS_API_KEY;

if (!baseUrl || !apiKey) {
  console.error("Set PROMPTDIS_URL and PROMPTDIS_API_KEY environment variables.");
  process.exit(1);
}

const client = new PromptClient({ baseUrl, apiKey });

// 1. Fetch a prompt by name
console.log("--- Fetch by name ---");
const prompt: Prompt = await client.getByName("demo", "myapp", "greeting");
console.log(`  ID:          ${prompt.id}`);
console.log(`  Name:        ${prompt.name}`);
console.log(`  Version:     ${prompt.version}`);
console.log(`  Model:       ${(prompt.model?.default as string) ?? "unknown"}`);
console.log(`  Environment: ${prompt.environment}`);
console.log(`  Body:        ${prompt.body.slice(0, 80)}...`);

// 2. Fetch the same prompt by ID
console.log("\n--- Fetch by ID ---");
const samePrompt: Prompt = await client.get(prompt.id);
console.log(`  Name:        ${samePrompt.name}`);

// 3. Local render (basic {{var}} substitution)
console.log("\n--- Local render ---");
const rendered: string = client.renderLocal(prompt.body, { name: "Alice" });
console.log(`  Rendered:    ${rendered.slice(0, 120)}`);

// 4. Server-side render (full Jinja2)
console.log("\n--- Server-side render ---");
const result = await client.render(prompt.id, { name: "Bob" });
console.log(`  Rendered:    ${result.rendered_body.slice(0, 120)}`);

// 5. Cache stats
console.log("\n--- Cache stats ---");
const stats = client.cacheStats();
console.log(`  Size:        ${stats.size}`);
console.log(`  Max size:    ${stats.maxSize}`);

// 6. Error handling
console.log("\n--- Error handling ---");
try {
  await client.get("nonexistent-id");
} catch (e) {
  if (e instanceof NotFoundError) {
    console.log("  Caught NotFoundError (expected)");
  } else if (e instanceof PromptdisError) {
    console.log(`  Caught PromptdisError: ${(e as Error).message}`);
  }
}
