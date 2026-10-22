# Security Policy

## Supported Versions

provider-confluent follows the standard Crossplane support policy: security
fixes are made on the `main` branch and released as a new patch version of the
latest release. Please upgrade to the latest release before reporting an issue.

## Reporting a Vulnerability

Please do **not** report security vulnerabilities in public issues. To report a
vulnerability, either:

1. Use GitHub's private vulnerability reporting. Open the
   [security tab](https://github.com/crossplane-contrib/provider-confluent/security)
   of this repository, click on 'Advisories', then 'Report a vulnerability', and
   describe the issue. See the [GitHub
   documentation](https://docs.github.com/en/code-security/security-advisories/guidance-on-reporting-and-writing/privately-reporting-a-security-vulnerability)
   for details.

2. Send an email to `crossplane-security@lists.cncf.io` with the details of the
   issue, as described in the [Crossplane security
   policy](https://github.com/crossplane/crossplane/blob/main/SECURITY.md).

Useful information to include: the affected provider-confluent version, a
description of the issue and its impact, and steps to reproduce it.

The maintainers listed in [OWNERS.md](./OWNERS.md) will acknowledge the report
and coordinate a fix and disclosure with the reporter. This is a volunteer-run
community project, so please allow a reasonable amount of time for a response.

## Vulnerabilities in dependencies

provider-confluent embeds the Terraform CLI (pinned to the last MPL-licensed
release, 1.5.7) and the Confluent Terraform provider. Findings in those
binaries are tracked in the issues of this repository and are fixed as soon as
a compatible, permissively licensed upstream version is available.
