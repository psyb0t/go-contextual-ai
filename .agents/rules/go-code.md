# Go code

Refines `~/.claude/rule-details/working-with-generated-code.md` and `~/.claude/rule-details/coding-style.md`.

- **GO1** `client.gen.go` is generated. Never edit it by hand. Change `config.yaml`, a script in `scripts/`, or `gen.go`, then run `make generate`.
- **GO2** A spec problem that breaks oapi-codegen is fixed by an oapi-codegen config option or a patch script in `scripts/` that `gen.go` runs between the fetch and the generate step. Never by editing the downloaded spec or the output.
- **GO3** The whole spec is generated. No `include-tags` or other trimming.
- **GO4** `gen.go` holds only the package doc comment, the package clause and `//go:generate` lines. No other hand-written code goes in the package.
- **GO5** `make lint` passes with zero issues on the generated and hand-written code. Lines in hand-written files are at most 80 columns, the `.golangci.yml` limit.
- **GO6** Scripts in `scripts/` keep the template shape: `set -euo pipefail`, the `on_error` trap, JSON `log` lines via jq, and they pass shellcheck and shfmt.
