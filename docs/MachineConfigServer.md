# MachineConfigServer

## Goals

1. Provide Ignition config to new machines joining the cluster.

## Non Goals

## Overview

All the machines joining the cluster must receive configuration from component running inside the cluster. MachineConfigServer provides the Ignition endpoint that all the new machines can point to receive their machine configuration. 

The machine can request specific configuration by pointing Ignition to MachineConfigServer and passing appropriate parameters. For example, to fetch configuration for master the machine can point to `/config/master` whereas, to fetch configuration for worker the machine can point to `/config/worker`.

## Detailed Design

### Endpoint

MachineConfigServer serves Ignition at `/config/<machine-pool-name>` endpoint.

* If the server finds the machine pool requested in the URL, it returns the Ignition config stored in the MachineConfig object referenced at `.status.currentMachineConfig` in the MachinePool object.

* If the server cannot find the machine pool requested in the URL, the server returns HTTP Status Code 404 with empty response.

### Special parameter for master machine configuration

The etcd members are co-located with the control plane on `master` machines. Bootstrapping an etcd cluster requires special code path that depends on the index of the machine in the pool. Therefore, the server needs to support `etcd_index` query parameter for `master` machine pools to serve correct Ignition config.

### Query parameters

Currently no parameters are supported.

### Ignition config from MachineConfig

MachineConfigServer serves the Ignition config defined in `spec.config` fields of the appropriate MachineConfig object.

But MachineConfigServer performs following extra actions on the Ignition config defined in the MachineConfig object before serving it,

1. etcd member index templating

    * The etcd unit files generated by MachineConfigController are go text templates with `etcd_index` variable. This variable needs to replaces with the index specified in `etcd_index` query.

2. Ignition file for MachineConfigDaemon

    MachineConfigDaemon requires file on disk to seed the `currentConfig` & `desiredConfig` annotations to its node object. The file is JSON object that contains the reference to `MachineConfig` object used to generate the Ignition config for the machine.

### Running MachineConfigServer

It is recommended that the MachineConfigServer is run as a DaemonSet on all `master` machines with the pods running in host network. So machines can access the Ignition endpoint through load balancer setup for control plane.

### Example requests

1. Worker machine

    Request:

    GET `/config/worker`

    Response:

    **TODO(abhinavdahiya): add example response**

2. Master machine with etcd member `etcd-3`

    Request:

    GET `/config/master?etcd_index=3`

    Response:

    **TODO(abhinavdahiya): add example response**