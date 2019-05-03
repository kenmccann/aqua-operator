<h1 align="center">Aqua Security Operator</h1>
<p align="center">
  <img width="150px" height="150px" src="images/logo.png"/>
</p>

<h2>Contents</h2>

- [About](#about)
- [Deployment Requirements](#deployment-requirements)
- [Documentation](#documentation)
- [Issues and feedback](#issues-and-feedback)

## About

The **aqua-operator** is a group of controllers that runs within a Kubernetes or Openshift cluster that provides a means to deploy and manage Aqua Security cluster and Components as:
* Server (Console)
* Gateway
* Database (Not recommended for production environments)
* Enforcer (Agent)
* Scanner CLI
* CSP (Package of Server, Gateway and Database)

**Use the aqua-operator to:**
 * Deploy Aqua Security components on Kubernetes or Openshift
 * Scale up Aqua Security components with extra replicas
 * Assign metadata tags to Aqua Security components
 * Automatic scale to the Aqua Scanner CLI by the count of the Scanning Queue.

## Deployment Requirements

The Operator deploys on Kubernetes and Openshift clusters.

* **Kubernetes:**  1.11.0 +
* **Openshift:** 3.11 +

## Documentation
The following documentation is provided:
- [Installation](docs/Installation.md)
- [First Steps](docs/FirstSteps.md)
- **Deployments**:
  - [Aqua CSP](docs/AquaCsp.md) - **recommended**!
  - [Aqua Server](docs/AquaServer.md)
  - [Aqua Gateway](docs/AquaGateway.md)
  - [Aqua Database](docs/AquaDatabase.md) - **Not For Production Environment - Please Use External DB with PostgreSQL Operator**
  - [Aqua Enforcer](docs/AquaEnforcer.md)
  - [Aqua Scanner CLI](docs/AquaScanner.md)
- [Extra Aqua Types](docs/Types.md)
- [Official Aqua Security Docs Site](https://read.aquasec.com/)

## Issues and feedback

If you encounter any problems or would like to give us feedback on deployments, we encourage you to raise issues here on GitHub.
