#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
FAILED=()

step() { printf "\n\033[1;34m==> %s\033[0m\n" "$1"; }
ok()   { printf "\033[1;32m    OK\033[0m\n"; }
fail() { printf "\033[1;31m    FAILED\033[0m\n"; FAILED+=("$1"); }

# ── Python SDK ──────────────────────────────────────────────
step "Python SDK: install"
if (cd "$ROOT/sdk-py" && pip install -e . --quiet 2>&1); then
    ok
else
    fail "sdk-py install"
fi

step "Python SDK: syntax check example"
if python -m py_compile "$ROOT/sdk-py/examples/main.py"; then
    ok
else
    fail "sdk-py example syntax"
fi

# ── JavaScript SDK ──────────────────────────────────────────
step "JavaScript SDK: install dev deps"
if (cd "$ROOT/sdk-js" && npm install --silent 2>&1); then
    ok
else
    fail "sdk-js install"
fi

step "JavaScript SDK: syntax check example"
if node --check "$ROOT/sdk-js/examples/main.mjs"; then
    ok
else
    fail "sdk-js example syntax"
fi

# ── TypeScript SDK ──────────────────────────────────────────
step "TypeScript SDK: install dev deps"
if (cd "$ROOT/sdk-ts" && npm install --silent 2>&1); then
    ok
else
    fail "sdk-ts install"
fi

step "TypeScript SDK: build (tsc)"
if (cd "$ROOT/sdk-ts" && npm run build 2>&1); then
    ok
else
    fail "sdk-ts build"
fi

step "TypeScript SDK: install example deps"
if (cd "$ROOT/sdk-ts/examples" && npm install --silent 2>&1); then
    ok
else
    fail "sdk-ts example install"
fi

step "TypeScript SDK: type-check example"
if (cd "$ROOT/sdk-ts/examples" && npx tsc --noEmit 2>&1); then
    ok
else
    fail "sdk-ts example type-check"
fi

# ── Go SDK ──────────────────────────────────────────────────
step "Go SDK: build"
if (cd "$ROOT/sdk-go" && go build ./... 2>&1); then
    ok
else
    fail "sdk-go build"
fi

step "Go SDK: build example"
if (cd "$ROOT/sdk-go/examples/basic" && go build . 2>&1); then
    ok
else
    fail "sdk-go example build"
fi

# ── Summary ─────────────────────────────────────────────────
echo ""
if [ ${#FAILED[@]} -eq 0 ]; then
    printf "\033[1;32mAll SDK builds passed.\033[0m\n"
else
    printf "\033[1;31m%d step(s) failed:\033[0m\n" "${#FAILED[@]}"
    for f in "${FAILED[@]}"; do
        printf "  - %s\n" "$f"
    done
    exit 1
fi
