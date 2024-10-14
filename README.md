# Radix Prometheus Proxy

Currently this exports uptime statistics using the `probe_success` metric from prometheus.

## How we work

Commits to the main branch must follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) and uses Release Please to create new versions.

# TODO:

- Manually patch helm chart based on please-release output
- Manually tag docker image based on please-release output

# Notes:

## Release-Please output:
```json
  {
    "releases_created": "true",
    "release_created": "true",
    "id": "179766252",
    "name": "v1.3.1",
    "tag_name": "v1.3.1",
    "sha": "1785100949bdc314681a91e136a42248406d309f",
    "body": "...",
    "html_url": "https://github.com/equinor/radix-prometheus-proxy/releases/tag/v1.3.1",
    "draft": "false",
    "upload_url": "https://uploads.github.com/repos/equinor/radix-prometheus-proxy/releases/179766252/assets{?name,label}",
    "path": ".",
    "version": "1.3.1",
    "major": "1",
    "minor": "3",
    "patch": "1",
    "paths_released": "[\".\"]"
  }
```
