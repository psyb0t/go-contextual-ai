# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.2.0, 2026-10-28

- The client is generated from the spec Contextual AI publishes at `https://docs.contextual.ai/api-reference/openapi.json`. The Stainless-hosted copy of the spec it used before is no longer publicly readable, so scheduled regeneration had stopped.
- Breaking: `ListDatastores` returns `DatastoreOutputList`, which pages with `NextCursor` and has no `TotalCount`. The 200 response union of `GetAgentMetadata` and the types `GetTwilightAgentResponse`, `AgentUsages`, `PerAgentRoles`, `ParseFigureCaptionMode` and `AgentLevelUserRolesEnum` are removed.
- Added 32 endpoint methods, including extract, billing, workspace consent and ACL query.
- Updated Semgrep in the development image to 1.182.1 and refreshed its locked dependencies. certifi 2026.10.23 drops the emSign root CAs whose intermediate keys were compromised (GHSA-6q48-6c7w-9gw2).
- Dependabot no longer opens pull requests for `mcp` or the OpenTelemetry packages in `requirements-semgrep.txt`. Semgrep pins their versions itself, so bumping one of them on its own cannot install.
- Updated kin-openapi, which the code generator uses, from 0.142.0 to 0.149.0. The newer version fixes several panics on malformed specs reported in GitHub security advisories.
- The security scan merges its two gitleaks runs into one SARIF run. GitHub code scanning rejects a report with two runs from the same tool, so no scan results had been uploaded.

## v0.1.3, 2026-10-27

- The development image is based on Debian (`golang:1.26.10-trixie`) instead of Alpine. Alpine keeps only the latest build of each package, so the exact pins broke the image build whenever a package was updated. System packages now come from the snapshot.debian.org date the base image was built from, which keeps them fixed for a given base image digest. jq 1.8.2, shellcheck 0.11.0 and shfmt 3.13.1 keep their versions, and Semgrep runs in its own virtual environment.
- Raised the minimum Go version in `go.mod` from 1.26.9 to 1.26.10, which fixes the `crypto/x509` vulnerability GO-2026-6809 found by govulncheck.
- The issue-pull workflow runs at hours derived from the repository name instead of the same four hours as every other repository.

## v0.1.2, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.1, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.0, 2026-10-16

- Added `contextual`, a REST client generated with oapi-codegen from Contextual AI's OpenAPI document (spec version 1.0, OpenAPI 3.1.0, hosted by Stainless for the official SDKs). It covers all 31 paths: agents and queries, datastores and documents, generate, LMUnit, parse jobs, rerank and users.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` reads the current spec URL from the `.stats.yml` of Contextual AI's Python SDK, downloads the spec and regenerates the client. The spec is not committed.
- Added the `go-contextual-ai` agent skill. Install it with `claude plugin install go-contextual-ai@psyb0t` (Claude Code), `codex plugin add go-contextual-ai@psyb0t` (Codex) or `openclaw skills install @psyb0t/go-contextual-ai` (ClawHub).
