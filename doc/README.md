# Documentation

## Deployment Architecture
Kubernetes resources and link to Azure Active Directory (AAD) are outlined below.
[OAuth2 Proxy](https://oauth2-proxy.github.io/oauth2-proxy/) forwards authenticated
users up to a [Gin](https://github.com/gin-gonic/gin) app running in the same
namespaces as [OpenCost](https://github.com/opencost/opencost) and [Prometheus](https://prometheus.io/). OCost queries the OpenCost API and returns a templated HTML response
based on the AAD security groups the user is a member of.

<p align="center">
  <img src="assets/architecture.png" alt="deployment-architecture" width="800"/>
</p>

## Configuration
There are two configuration files required at the root level of the repository `.env` and `config.yaml`.
The sample files should provide a reasonable starting place once the `__CHANGE_ME__` values are
replaced.

### `.env`
Contains secrets and other variables that will be injected as environment variables and
are required for deployment. In CI these variables could be added from [Github secrets](https://docs.github.com/en/actions/security-guides/using-secrets-in-github-actions)
or [variables](https://docs.github.com/en/actions/learn-github-actions/variables) depending on their
sensitivity.

### `config.yaml`
Contains app-level configuration such as pricing and the group ⇔ namespace mapping. These
should be checked into the repository and push to a **private** remote.
