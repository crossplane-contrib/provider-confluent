# Contributing to provider-confluent

Thank you for helping to improve provider-confluent! This is a community
extension project of [Crossplane](https://crossplane.io). Please read the
Crossplane [contribution
process](https://github.com/crossplane/crossplane/blob/main/contributing/README.md)
if this is your first Crossplane contribution, and find us on the [Crossplane
Slack](https://slack.crossplane.io/).

## Code of Conduct

This project follows the [CNCF Code of Conduct](./CODE_OF_CONDUCT.md).

## Reporting issues

For bugs, improvements or feature requests, please open an
[issue](https://github.com/crossplane-contrib/provider-confluent/issues) and
include the provider and Crossplane versions, the resource manifests involved
and the relevant provider logs. Report security issues as described in
[SECURITY.md](./SECURITY.md), not in public issues.

## Developer Certificate of Origin (DCO) sign-off

All commits must be signed off to certify that you wrote, or have the right to
submit, the code under the project's [Apache 2.0 license](./LICENSE) (see the
[Developer Certificate of Origin](https://developercertificate.org/)). Add the
sign-off with `-s`:

```console
git commit -s -m "feat: add support for X"
```

This appends a `Signed-off-by: Your Name <you@example.com>` line that matches
your Git `user.name` and `user.email`. To fix a commit that is missing the
sign-off, run `git commit --amend -s` (or `git rebase --signoff main` for
several commits). Pull requests with unsigned commits fail the DCO check.

## Development

provider-confluent is built with [Upjet](https://github.com/crossplane/upjet).
Familiarize yourself with its documentation first:

1. [Generating a Crossplane provider using
   Upjet](https://github.com/crossplane/upjet/blob/main/docs/README.md#generating-a-crossplane-provider-using-upjet)
2. [Configuring a
   resource](https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md)

Resource configuration lives in `config/`. The `apis/`, `internal/controller/`,
`package/crds/` and `examples-generated/` directories are generated: change the
configuration and run `make generate` instead of editing them by hand.

### Toolchain

Use the Go version the repository is built with (see `GO_REQUIRED_VERSION` in
the [Makefile](./Makefile)). With a newer local Go, select the right toolchain
and install `goimports` before generating:

```console
export GOTOOLCHAIN=go1.24.2 PATH="$PATH:$(go env GOPATH)/bin"
go install golang.org/x/tools/cmd/goimports@latest
```

### Building and testing

Initialize the `build` submodule on a fresh clone, then run the checks:

```console
git submodule update --init
make reviewable test
```

`make reviewable` runs code generation, linting and the unit tests, and the
resulting generated code must be committed. A failed `make generate` can leave
generated files deleted; restore them with `git checkout -- apis
examples-generated internal package config`.

Run the provider against a Kubernetes cluster, out of cluster, with:

```console
make run
```

### Terraform version

The embedded Terraform CLI must stay on 1.5.x. Terraform 1.6 and later are
licensed under the BSL, which is not permitted for Crossplane community
extension projects.

## Pull requests

- Open a pull request against `main` and describe what it changes and how it
  was tested.
- Make sure `make reviewable test` passes and all commits are signed off.
- Keep pull requests focused. Maintainers listed in [OWNERS.md](./OWNERS.md)
  review them as time allows.
