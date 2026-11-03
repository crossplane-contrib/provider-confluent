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

Resource configuration lives in `config/`. Generated code is written to the
`zz_*.go` files under `apis/` and `internal/controller/`, plus `package/crds/`,
`examples-generated/`, and the generated config metadata files
(`config/schema.json`, `config/provider-metadata.yaml`, `config/generated.lst`).
Change the configuration and run `make generate` instead of editing those
generated files by hand.

### Toolchain

Use the Go version declared in [`go.mod`](./go.mod). With a newer local Go,
select the matching toolchain and install `goimports` before generating:

```console
GO_VERSION=$(awk '$1 == "go" { print $2; exit }' go.mod)
export GOTOOLCHAIN="go${GO_VERSION}"
export PATH="$PATH:$(go env GOPATH)/bin"
go install golang.org/x/tools/cmd/goimports
```

### Building and testing

Initialize the `build` submodule on a fresh clone, then run the checks:

```console
git submodule update --init
make reviewable test
```

`make reviewable` runs code generation, linting and the unit tests, and the
resulting generated code must be committed. A failed `make generate` can leave
generated files deleted or truncated; restore only the generated outputs. Do
**not** restore all of `config/`, `internal/`, or `package/` as those directories
contain hand-written files:

```console
git restore -- \
  ':(glob)apis/**/zz_*.go' \
  ':(glob)internal/controller/**/zz_*.go' \
  package/crds \
  examples-generated \
  config/schema.json \
  config/provider-metadata.yaml \
  config/generated.lst
```

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
