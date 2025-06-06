# VMware Secrets Manager (VSecM) Helm Chart

VMware Secrets Manager keeps your secrets secret. With VSecM, you can rest assured
that your sensitive data is always secure and protected. VSecM is perfect for
securely storing arbitrary configuration information at a central location and
securely dispatching it to workloads.

![Version: 0.28.1](https://img.shields.io/badge/Version-0.28.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.28.1](https://img.shields.io/badge/AppVersion-0.28.1-informational?style=flat-square)

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/vsecm)](https://artifacthub.io/packages/helm/vsecm/vsecm)

## Quickstart

To use VMware Secrets Manager, follow the steps below:

1. Add VMware Secrets Manager Helm repository:

    ```bash
    helm repo add vsecm https://vmware-tanzu.github.io/secrets-manager/
    ```

2. Update the helm repository:

    ```bash
    helm repo update
    ```

3. Install VMware Secrets Manager using Helm:

    ```bash
    helm install vsecm vsecm/vsecm --version 0.28.1
    ```

## Options

The following options can be passed to the `helm install` command to set global
variables:

*`--set global.deploySpire=<true/false>`:
  This flag can be passed to install or skip SPIRE.
*`--set global.baseImage=<distroless|distroless-fips>`:
  This flag can be passed to install VSecM with the given baseImage Docker image.

Default values are `true` and `distroless` for `global.deploySpire`
and `global.baseImage` respectively.

Here's an example command with the above options:

```bash
helm install vsecm vsecm/helm-charts --version 0.28.1 \
  --set global.deploySpire=true --set global.baseImage=distroless
```

Make sure to replace `<true|false>` and
`<distroless|distroless-fips>` with the desired values.

## Environment Configuration

**VMware Secrets Manager** can be tweaked further using environment variables.

[Check out **Configuring VSecM** on the official documentation][configuring-vsecm]
for details.

These environment variable configurations are expose through subcharts.
You can modify them as follows:

```bash
helm install vsecm vsecm/helm-charts --version 0.28.1  \
--set safe.environments.VSECM_LOG_LEVEL="6"
--set sentinel.environments.VSECM_LOGL_LEVEL="5"
# You can update other environment variables too.
# Most of the time VSecM assumes sane defaults if you don't set them.
```

[configuring-vsecm]: https://vsecm.com/docs/configuration/

## Subcharts

For further details about subcharts follow these links:

* [VSecM Safe](charts/safe/README.md)
* [VSecM Sentinel](charts/sentinel/README.md)
* [VsecM Keystone](charts/keystone/README.md)
* [SPIRE](charts/spire/README.md)

Please check out [the official **VSecM** documentation][ducks]
for more information about **VSecM** components and the overall
**VSecM** architecture.

[ducks]: https://vsecm.com/documentation/welcome/overview/

## Detailed Documentation

The sections below are autogenerated from chart source code:

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| file://charts/keystone | keystone | 0.28.1 |
| file://charts/safe | safe | 0.28.1 |
| file://charts/sentinel | sentinel | 0.28.1 |
| file://charts/spire | spire | 0.28.1 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| global.baseImage | string | `"distroless"` | Possible options for baseImage (distroless, distroless-fips). When in doubt, stick with distroless. |
| global.deployKeystone | bool | `true` | Deploy the Keystone VSecM component. VSecM Keystone is a lightweight Pod that is initialized only after VSecM Sentinel completes it `initCommand` initialization sequence. |
| global.deploySentinel | bool | `true` | Deploy VSecM Sentinel. VSecM Sentinel is the only admin interface where you can register secrets. For best security, you might want to disable the initial deployment of it. This way, you can deploy VSecM Sentinel off-cycle later when you need it. |
| global.deploySpire | bool | `true` | Deploy SPIRE components. If set to false, SPIRE components will not be  deployed. This is useful when SPIRE is already deployed in the cluster. |
| global.enableKAppAnnotations | bool | `false` | Set it to true to enable kapp annotations. This is useful when you are using kapp to deploy the VSecM components. (ref: https://carvel.dev/kapp/) |
| global.enableOpenShift | bool | `false` | Set it to true for OpenShift deployments. This will add necessary annotations to the SPIRE components to make them work on OpenShift. |
| global.images | object | `{"initContainer":{"repository":"vsecm-ist-init-container","tag":"0.28.1"},"keystone":{"distrolessFipsRepository":"vsecm-ist-fips-keystone","distrolessRepository":"vsecm-ist-keystone","pullPolicy":"IfNotPresent","tag":"0.28.1"},"nodeDriverRegistrar":{"pullPolicy":"IfNotPresent","repository":"registry.k8s.io/sig-storage/csi-node-driver-registrar","tag":"v2.10.0"},"openShiftHelperUbi9":{"pullPolicy":"IfNotPresent","repository":"registry.access.redhat.com/ubi9","tag":"latest"},"safe":{"distrolessFipsRepository":"vsecm-ist-fips-safe","distrolessRepository":"vsecm-ist-safe","pullPolicy":"IfNotPresent","tag":"0.28.1"},"sentinel":{"distrolessFipsRepository":"vsecm-ist-fips-sentinel","distrolessRepository":"vsecm-ist-sentinel","pullPolicy":"IfNotPresent","tag":"0.28.1"},"spiffeCsiDriver":{"pullPolicy":"IfNotPresent","repository":"ghcr.io/spiffe/spiffe-csi-driver","tag":"0.2.6"},"spireAgent":{"pullPolicy":"IfNotPresent","repository":"ghcr.io/spiffe/spire-agent","tag":"1.9.6"},"spireControllerManager":{"pullPolicy":"IfNotPresent","repository":"ghcr.io/spiffe/spire-controller-manager","tag":"0.5.0"},"spireHelperBash":{"pullPolicy":"IfNotPresent","repository":"cgr.dev/chainguard/bash","tag":"latest@sha256:8c9e5cbb641ced8112c637eb3611dab29bf65448a9d884a03938baf1b352dc4d"},"spireHelperKubectl":{"pullPolicy":"IfNotPresent","repository":"docker.io/rancher/kubectl","tag":"v1.28.0"},"spireServer":{"pullPolicy":"IfNotPresent","repository":"ghcr.io/spiffe/spire-server","tag":"1.9.6"}}` | Where to find the dependent images of VSecM. Normally, you would not need to modify this. |
| global.images.nodeDriverRegistrar | object | `{"pullPolicy":"IfNotPresent","repository":"registry.k8s.io/sig-storage/csi-node-driver-registrar","tag":"v2.10.0"}` | Container registry details of SPIFFE CSI Node Driver Registrar. |
| global.images.spiffeCsiDriver | object | `{"pullPolicy":"IfNotPresent","repository":"ghcr.io/spiffe/spiffe-csi-driver","tag":"0.2.6"}` | Container registry details of SPIFFE CSI Driver. |
| global.images.spireAgent | object | `{"pullPolicy":"IfNotPresent","repository":"ghcr.io/spiffe/spire-agent","tag":"1.9.6"}` | Container registry details of SPIRE Agent. |
| global.images.spireControllerManager | object | `{"pullPolicy":"IfNotPresent","repository":"ghcr.io/spiffe/spire-controller-manager","tag":"0.5.0"}` | Container registry details of SPIRE Controller Manager. |
| global.images.spireServer | object | `{"pullPolicy":"IfNotPresent","repository":"ghcr.io/spiffe/spire-server","tag":"1.9.6"}` | Container registry details of SPIRE Server. |
| global.preInstallSpireNamespaces | bool | `true` | Set it to true to enable the pre-installation of the SPIRE namespaces. If set to false, the SPIRE namespaces will not be pre-installed; you will need to create `spire-system` and `spire-server` namespaces manually. |
| global.preInstallVSecMNamespaces | bool | `true` | Set it to true to enable the pre-installation of the VSecM namespaces. If set to false, the VSecM namespaces will not be pre-installed; you will need to create a `vsecm-system` namespace manually. |
| global.registry | string | `"vsecm"` | Registry url. Defaults to "vsecm", which points to the public vsecm DockerHub registry: <https://hub.docker.com/u/vsecm>. |
| global.spire | object | `{"caCommonName":"vsecm.com","caCountry":"US","caOrganization":"vsecm.com","controllerManagerClassName":"vsecm","federationEnabled":false,"logLevel":"DEBUG","namespace":"spire-system","serverAddress":"spire-server.spire-server.svc.cluster.local","serverNamespace":"spire-server","serverPort":443,"trustDomain":"vsecm.com"}` | SPIRE-related global configuration. |
| global.spire.caCommonName | string | `"vsecm.com"` | The SPIRE CA common name. |
| global.spire.caCountry | string | `"US"` | The SPIRE CA country. |
| global.spire.caOrganization | string | `"vsecm.com"` | The SPIRE CA organization. |
| global.spire.controllerManagerClassName | string | `"vsecm"` | This is the className that ClusterSPIFFEIDs will use to be able to register their SPIFFE IDs with the SPIRE Server. |
| global.spire.federationEnabled | bool | `false` | Enable federation. If set to true, SPIRE Server will be configured to federate with other SPIRE Servers. This is useful when you have multiple clusters, and you want to establish trust between them. |
| global.spire.logLevel | string | `"DEBUG"` | The log level of the SPIRE components. This is useful for debugging. |
| global.spire.namespace | string | `"spire-system"` | This is the namespace where the SPIRE components will be deployed. |
| global.spire.serverAddress | string | `"spire-server.spire-server.svc.cluster.local"` | The SPIRE Server address. This is the address where the SPIRE Server that the agents will connect to. This address is in the form of <service-name>.<namespace>.svc.cluster.local unless you have a custom setup. |
| global.spire.serverNamespace | string | `"spire-server"` | It is best to keep the SPIRE server namespace separate from other SPIRE components for an added layer of security. |
| global.spire.serverPort | int | `443` | The SPIRE Server port. This is the port where the SPIRE Server will listen for incoming connections. This is the port of the SPIRE server k8s Service. |
| global.spire.trustDomain | string | `"vsecm.com"` | The trust domain is the root of the SPIFFE ID hierarchy. It is used to identify the trust domain of a workload. If you use anything other than the default `vsecm.com`, you must also update the relevant environment variables that does SPIFFE ID validation.  To prevent accidental collisions (two trust domains select identical names), operators are advised to select trust domain names which are highly likely to be globally unique. Even though a trust domain name is not a DNS name, using a registered domain name as a suffix of a trust domain name, when available, will reduce chances of an accidental collision; for example, if a trust domain operator owns the domain name `example.com`, then using a trust domain name such as `apps.example.com` would likely not produce a collision. When trust domain names are automatically generated without operator input, randomly generating a unique name (such as a UUID) is strongly advised.  All SPIFFE IDs shall be prefixed with `spiffe://<trustDomain>` unless you have an advanced custom setup. |
| global.vsecm.keystoneSpiffeIdTemplate | string | `"spiffe://vsecm.com/workload/vsecm-keystone/ns/{{ .PodMeta.Namespace }}/sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}"` |  |
| global.vsecm.namespace | string | `"vsecm-system"` |  |
| global.vsecm.safeEndpointUrl | string | `"https://vsecm-safe.vsecm-system.svc.cluster.local:8443/"` |  |
| global.vsecm.safeSpiffeIdPrefix | string | `"^spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/[^/]+$"` |  |
| global.vsecm.safeSpiffeIdTemplate | string | `"spiffe://vsecm.com/workload/vsecm-safe/ns/{{ .PodMeta.Namespace }}/sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}"` |  |
| global.vsecm.sentinelSpiffeIdPrefix | string | `"^spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/[^/]+$"` |  |
| global.vsecm.sentinelSpiffeIdTemplate | string | `"spiffe://vsecm.com/workload/vsecm-sentinel/ns/{{ .PodMeta.Namespace }}/sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}"` |  |
| global.vsecm.workloadNameRegExp | string | `"^spiffe://vsecm.com/workload/([^/]+)/ns/[^/]+/sa/[^/]+/n/[^/]+$"` |  |
| global.vsecm.workloadSpiffeIdPrefix | string | `"^spiffe://vsecm.com/workload/[^/]+/ns/[^/]+/sa/[^/]+/n/[^/]+$"` |  |

## License

This project is licensed under the [BSD 2-Clause License][license].

[license]: https://github.com/vmware/secrets-manager/blob/main/LICENSE
