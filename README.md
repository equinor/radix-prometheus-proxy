# Radix Prometheus Proxy

Currently this exports uptime statistics using the `probe_success` metric from prometheus.

## How we work

Commits to the main branch must follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) and uses Release Please to create new versions.

# TODO:

- Manually patch helm chart based on please-release output
- Manually tag docker image based on please-release output
