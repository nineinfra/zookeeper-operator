# Helm Chart for Zookeeper Operator for Apache Zookeeper

This Helm Chart can be used to install Custom Resource Definitions and the Operator for Apache Zookeeper provided by Nineinfra.

## Requirements

- Create a [Kubernetes Cluster](../Readme.md)
- Install [Helm](https://helm.sh/docs/intro/install/)

## Install the Zookeeper Operator for Apache Zookeeper

```bash
# From the root of the operator repository

helm install zookeeper-operator charts/zookeeper-operator
```

## Usage of the CRDs

The usage of this operator and its CRDs is described in the [documentation](https://github.com/nineinfra/zookeeper-operator/blob/main/README.md).

## Links

https://github.com/nineinfra/zookeeper-operator
