# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.0, 2026-10-16

- Added `contextual`, a REST client generated with oapi-codegen from Contextual AI's OpenAPI document (spec version 1.0, OpenAPI 3.1.0, hosted by Stainless for the official SDKs). It covers all 31 paths: agents and queries, datastores and documents, generate, LMUnit, parse jobs, rerank and users.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` reads the current spec URL from the `.stats.yml` of Contextual AI's Python SDK, downloads the spec and regenerates the client. The spec is not committed.
- Added the `go-contextual-ai` agent skill. Install it with `claude plugin install go-contextual-ai@psyb0t` (Claude Code), `codex plugin add go-contextual-ai@psyb0t` (Codex) or `openclaw skills install @psyb0t/go-contextual-ai` (ClawHub).
