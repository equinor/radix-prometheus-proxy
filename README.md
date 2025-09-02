[![SCM Compliance](https://scm-compliance-api.radix.equinor.com/repos/equinor/radix-api/badge)](https://developer.equinor.com/governance/scm-policy/)

# Radix Prometheus Proxy

Currently this exports uptime statistics using the `probe_success` metric from prometheus.

## Security

Authentication and authorisation are performed through an HTTP bearer token, which is relayed to the Kubernetes API. The Kubernetes AAD integration then performs its authentication and resource authorisation checks, and the result is relayed to the user.

## Development Process

The `radix-prometheus-proxy` project follows a **trunk-based development** approach.

### 🔁 Workflow

- **External contributors** should:
  - Fork the repository
  - Create a feature branch in their fork

- **Maintainers** may create feature branches directly in the main repository.

### ✅ Merging Changes

All changes must be merged into the `main` branch using **pull requests** with **squash commits**.

The squash commit message must follow the [Conventional Commits](https://www.conventionalcommits.org/en/about/) specification.

## Release Process

Merging a pull request into `main` triggers the **Prepare release pull request** workflow.  
This workflow analyzes the commit messages to determine whether the version number should be bumped — and if so, whether it's a major, minor, or patch change.  

It then creates two pull requests:

- one for the new stable version (e.g. `1.2.3`), and  
- one for a pre-release version where `-rc.[number]` is appended (e.g. `1.2.3-rc.1`).

---

Merging either of these pull requests triggers the **Create releases and tags** workflow.  
This workflow reads the version stored in `version.txt`, creates a GitHub release, and tags it accordingly.

The new tag triggers the **Build and deploy Docker and Helm** workflow, which:

- builds and pushes a new container image and Helm chart to `ghcr.io`, and  
- uploads the Helm chart as an artifact to the corresponding GitHub release.

## Contribution

Want to contribute? Read our [contributing guidelines](./CONTRIBUTING.md)

## Security

This is how we handle [security issues](./SECURITY.md)
