# go-clock

Injectable time source (package `clock`): the `Clock` interface (`Now() time.Time`), the production `System` clock (UTC), and the controllable `Fake` for deterministic tests. Generic — lives in `gomatic`; originated in xto-email's `xto` repo (`internal/clock`) and was promoted here during the xto repo split (see `xto-email/_projects/specs/repo-split/`).

- Library repo (`library.go` marker); flat single-package layout at the root; stdlib only (testify for tests).
- Gate: shared Makefile from `nicerobot/tools.repository` — gofumpt, vet, staticcheck, golangci-lint, govulncheck, gocognit ≤ 7, 100% coverage. Never edit the distributed `Makefile`/`.golangci.yaml`/`.github` in-tree.
- Public docs live in `docs.go-clock`; the README is exactly badges + the docs link.
