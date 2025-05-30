[![SCM Compliance](https://scm-compliance-api.radix.equinor.com/repos/equinor/radix-api/badge)](https://developer.equinor.com/governance/scm-policy/)

# Radix Prometheus Proxy

Currently this exports uptime statistics using the `probe_success` metric from prometheus.

## How we work

Commits to the main branch must follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) and uses Release Please to create new versions. This is now required by branch protection rules

## Security

Authentication and authorisation are performed through an HTTP bearer token, which is relayed to the Kubernetes API. The Kubernetes AAD integration then performs its authentication and resource authorisation checks, and the result is relayed to the user.

## Contributing

Read our [contributing guidelines](./CONTRIBUTING.md)

------------------

[Security notification](./SECURITY.md)
