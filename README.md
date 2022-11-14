# Launchpad

### A Zero DevOps workflow for building applications on Kubernetes

[![Join Discord](https://img.shields.io/discord/903306922852245526?color=7389D8&label=discord&logo=discord&logoColor=ffffff)](https://discord.gg/agbskCJXk2) ![License: Apache 2.0](https://img.shields.io/github/license/jetpack-io/devbox) [![version](https://img.shields.io/github/v/release/jetpack-io/launchpad?color=green&label=version&sort=semver)](https://github.com/jetpack-io/launchpad/releases) [![tests](https://github.com/jetpack-io/launchpad/actions/workflows/release.yaml/badge.svg)](https://github.com/jetpack-io/launchpad/actions/workflows/release.yaml?branch=main)


## What is it?

[Launchpad](https://www.jetpack.io/launchpad) is a command-line tool that lets you easily create applications on Kubernetes. You start by choosing the cluster you want to deploy to, and launchpad uses that definition to create a deployment just for your application.

In practice, Launchpad works similar to Heroku or Vercel, except everything is on Kubernetes.


## Demo

The example below initializes a web project with `launchpad init`, and deploys to a local Kubernetes cluster with `launchpad up`:

![screen cast](https://user-images.githubusercontent.com/2292093/201768560-b8a4db24-49c4-45cc-a4a4-b27c2815835e.svg)


## Benefits

### Build publish and deploy with a single command

Launchpad can build any image, publish it to your Docker registry, and deploy it to Kubernetes in one step.

### A consistent deployment setup per project

Launchpad defines your project's Kubernetes configuration. Whether you are deploying in development, staging, or production environment, whether you are deploying to a local or remote Kubernetes cluster, Launchpad always creates the same set of Kubernetes resources for you each time.

### Easy onboarding for everyone on the team

Launchpad gives each developer on your team a dedicated namespace to work on and deploy to and generates unique preview URLs for easy testing.

### Secret management built-in

Secrets are tied to your launchpad projects, so they can be shared and updated securely by your team.

## Installing Launchpad

In addition to installing Launchpad itself, you will need to install `docker` since Launchpad depends on it:

1. Install [Docker Desktop](https://www.docker.com/get-started/).

2. Install Launchpad:

   ```sh
   curl -fsSL https://get.jetpack.io/launchpad | bash
   ```

Read more on the [Launchpad docs](http://www.jetpack.io/launchpad/docs/getting-started/any-image-to-k8s-5-min/).


## Quickstart: deploy to your Docker Desktop Kubernetes cluster

In this quickstart, we’ll deploy a cron job to your local Docker Desktop Kubernetes cluster.

1. Open a terminal in a new empty folder called `app/`.

2. Enable [Kubernetes on Docker Desktop](https://docs.docker.com/desktop/kubernetes/)

3. Initialize Launchpad in `app/`:

   ```bash
   > launchpad init
   ```
   
   You will see the following questions:
   ```
   ? What is the name of this project? // Press <Enter> to use the default name
   ? What type of service you would like to add to this project? // Choose `Cron Job`
   ? To which cluster do you want to deploy this project? // Choose `docker-desktop`
   ```
   
   
   This creates a `launchpad.yaml` file in the current directory. You should commit it to source control.

4. Your `launchpad.yaml` file should now look like this:

   ```yaml
      configVersion: 0.1.2
      projectId: ...
      name: app
      cluster: docker-desktop
      services:
        app-cron:
          type: cron
          image: busybox:latest
          command: [/bin/sh, -c, date; echo Hello from Launchpad]
          schedule: '* * * * *'
   ```

5. Start a new deployment to Kubernetes:

   ```bash
   launchpad up
   ```

6. Wait for a minute, and see the cron job in action:

   ```bash
   > kubectl get pods
   > kubectl logs <pod_name>
   ```

   [kubectl](https://www.jetpack.io/launchpad/docs/getting-started/any-image-to-k8s-5-min/#prerequisites), while not required, is a useful tool for inspecting and managing your deployments in Kubernetes.


## Additional commands

`launchpad help` - see all commands

`launchpad env` - use launchpad's secret management toolchain

See the [CLI Reference](https://www.jetpack.io/launchpad/docs/reference/cli/) for the full list of commands.


## Join our Developer Community

- Chat with us by joining the [Jetpack.io Discord Server](https://discord.gg/jetpack-io) – we have a #launchpad channel dedicated to this project.
- File bug reports and feature requests using [Github Issues](https://github.com/jetpack-io/launchpad/issues)
- Follow us on [Jetpack's Twitter](https://twitter.com/jetpack_io) for product updates

## Contributing

Launchpad is an opensource project so contributions are always welcome. Please read [our contributing guide](CONTRIBUTING.md) before submitting pull requests.
