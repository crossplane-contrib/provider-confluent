# Security Policy

## Supported Versions

provider-confluent supports the latest release only. Security fixes are made on
`main` and released as a new patch version of the latest release. Please upgrade
to the latest release before reporting an issue.

## Reporting a Vulnerability

Please do **not** report security vulnerabilities in public issues. Instead,
send an email to `crossplane-security@lists.cncf.io` with the details of the
issue, as described in the [Crossplane security
policy](https://github.com/crossplane/crossplane/blob/main/SECURITY.md).

Useful information to include: the affected provider-confluent version, a
description of the issue and its impact, and steps to reproduce it.

Reports sent to `crossplane-security@lists.cncf.io` are handled through the
Crossplane security process. The Crossplane security team will coordinate with
the provider maintainers listed in [OWNERS.md](./OWNERS.md) on triage, a fix, and
disclosure. This is a volunteer-run community project, so please allow a
reasonable amount of time for a response.

## Vulnerabilities in dependencies

provider-confluent embeds the Terraform CLI, pinned to 1.5.7, the last
MPL-licensed release. Terraform 1.6.0 and later are licensed under the Business
Source License 1.1 (BUSL-1.1), which is not permitted for Crossplane community
extension projects, so there is no newer permissively licensed Terraform CLI to
upgrade to. Known findings in the embedded Terraform CLI are tracked in
[#50](https://github.com/crossplane-contrib/provider-confluent/issues/50).
Removing them requires an MPL-licensed alternative such as OpenTofu, or the
no-fork architecture tracked in
[#55](https://github.com/crossplane-contrib/provider-confluent/issues/55).

The embedded Confluent Terraform provider is Apache-2.0 and can be upgraded
independently. Vulnerabilities in that provider should be tracked separately
from the Terraform CLI findings.
