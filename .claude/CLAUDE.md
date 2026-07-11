# go-clock

Injectable time source for the xto-email ecosystem: the `clock.Clock` interface (`Now() time.Time`), the production `System` clock (UTC), and the controllable `Fake` for deterministic tests. Extracted from `xto`'s `internal/clock` when the repo split into `xtod`/`xtoctl` (see `_projects/specs/repo-split/`).

- Library repo (`library.go` marker); flat single-package layout at the root; stdlib only (testify for tests).
- Gate: shared Makefile from `nicerobot/tools.repository` — gofumpt, vet, staticcheck, golangci-lint, govulncheck, gocognit ≤ 7, 100% coverage. Never edit the distributed `Makefile`/`.golangci.yaml`/`.github` in-tree.
- Public docs live in `docs.go-clock`; the README is exactly badges + the docs link.
