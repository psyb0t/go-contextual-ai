# go-contextual-ai

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-contextual-ai.svg)](https://pkg.go.dev/github.com/psyb0t/go-contextual-ai)
[![CI](https://github.com/psyb0t/go-contextual-ai/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-contextual-ai/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-contextual-ai/badges/version.svg)](https://github.com/psyb0t/go-contextual-ai/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-contextual-ai/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-contextual-ai/badges/importers.svg)](https://github.com/psyb0t/go-contextual-ai/blob/badges/importers.md)

Unofficial Go client for the [Contextual AI API](https://docs.contextual.ai), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document Contextual AI publishes. I am not affiliated with Contextual AI.

The package is `contextual`, in the repository root. Everything in `client.gen.go` is generated. Don't edit it by hand, regenerate it.

## Install

```sh
go get github.com/psyb0t/go-contextual-ai
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-contextual-ai"
)

func main() {
	auth := contextual.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+os.Getenv("CONTEXTUAL_API_KEY"))

		return nil
	})

	client, err := contextual.NewClientWithResponses("https://api.contextual.ai/v1", auth)
	if err != nil {
		panic(err)
	}

	res, err := client.ListDatastoresDatastoresGetWithResponse(context.Background(), &contextual.ListDatastoresDatastoresGetParams{})
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
```

The same code is in `example_test.go`, so `make test` checks that it compiles.

Method names come from the spec's operation IDs, which repeat the path and the HTTP method, so they are long: `ListDatastoresDatastoresGet` lists datastores, `QueryAgentsAgentIdQueryPost` queries an agent, `CreateAgentAgentsPost` creates one.

The full API is in the [Go reference](https://pkg.go.dev/github.com/psyb0t/go-contextual-ai); after `go get`, `go doc github.com/psyb0t/go-contextual-ai` prints it locally.

## Auth

Contextual AI authenticates with a bearer token: send the API key as `Authorization: Bearer <key>`. Keys are created in the Contextual AI console. The base URL is `https://api.contextual.ai/v1`.

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` downloads the OpenAPI document Contextual AI publishes with its API reference, `https://docs.contextual.ai/api-reference/openapi.json`, to `api/openapi.json` and checks that it is an OpenAPI 3.x document.
- oapi-codegen runs with `config.yaml`. The config maps `number` to `float64`, because the spec gives scores and other numbers no format and oapi-codegen would otherwise use `float32`.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Agent integrations

The [skill](.agents/skills/go-contextual-ai) teaches a coding agent how to use go-contextual-ai when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-contextual-ai@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-contextual-ai@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-contextual-ai
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from Contextual AI's published API specification. The specification itself is not redistributed here.
