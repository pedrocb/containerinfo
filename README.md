# Containerinfo

## Description
Containerinfo is a service that queries the kubernetes clusters for its running containers and what are their resource requests and limits.
This repository contains the source code for the service (with kubernetes manifests for development environment) and a helm chart for packaging it.

## Installation
To install the service on a kubernetes cluster use helm.

```
helm install containerinfo containerinfo/Chart
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | Deployment affinity |
| clusterRole.annotations | object | `{}` | Cluster role annotations |
| clusterRole.create | bool | `true` | Whether a cluster role should be created |
| clusterRole.name | string | `""` | Name of cluster role. If empty and create = true a name is generated from the release name |
| clusterRoleBinding.create | bool | `true` | Whether a cluster role binding should be created  |
| clusterRoleBinding.name | string | `""` | Name of cluster role. If empty and create = true a name is generated from the release name |
| fullnameOverride | string | `""` | Override full name |
| image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| image.repository | string | `"pedrocbelem0/container-info"` | Image repository |
| image.tag | string | `""` | Image tag |
| imagePullSecrets | list | `[]` |  Image pull secrets|
| nameOverride | string | `""` | Override name |
| nodeSelector | object | `{}` | Deployment node seletor |
| podAnnotations | object | `{}` | Annotations for pod |
| podSecurityContext | object | `{}` |  Pod security context |
| replicaCount | int | `2` | Number of deployment replicas |
| resources | object | `{}` | Container resources |
| securityContext | object | `{}` | Security context for pod template |
| service.port | int | `8000` | Service Port |
| service.type | string | `"ClusterIP"` | Type of deployment |
| serviceAccount.annotations | object | `{}` | Annotations for service account |
| serviceAccount.create | bool | `true` | Whether a service  account should be created |
| serviceAccount.name | string | `""` | Service account name. If empty and create = true a name is generated from the release name |
| tolerations | list | `[]` | Deployment tolerations |

## Usage
Having the service installed on a kubernetes cluster, it can be used from inside the cluster using the endpoint `<serviceName>.<namespace>.svc.cluster.local:<port>/container-resources`

#### Example with default values
```
kubectl create -n my-home-assignment
helm install -n my-home-assignment containerinfo containerinfo/Chart

kubectl run --restart Never --image buildpack-deps:curl --rm -ti curl --command -- curl "containerinfo.my-home-assignment.svc.cluster.local:8000/container-resources?pod-label=app.kubernetes.io/instance=containerinfo"
```

## Development
For local development environment, first setup a local kubernetes cluster such as kind or minikube.
With the cluster running, simply start the app with `tilt up`. Tilt will create the minimum necessary resources to deploy the app on the cluster, it will create a port forward to the pod and will automatically build the image and update pod image if the application is modified.

```
kind create cluster

cd containerinfo && tilt up

curl localhost:8000/container-resources
```


