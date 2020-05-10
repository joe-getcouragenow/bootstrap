# bootstrap

Bootstrap provides the tooling.

It is concerned with:

1. binary tools that are reused on projects

2. make files that are reused on projects

3. core flutter and golang code that are reused on projects.


It covers the following scenarios:

1. Dev time - when developing.

2. Compile time - when compiling.

3. Deploy time - when deploying locally or to Kubernetes.

4. Ops time - when managing the deployment.


Projects should make their main and modules as they see fit, and use this tooling for them.

## Dev time

Bootstrap your OS with the tools needed.

``` bs init ```

- Installs the required binaries and boilerplate files.


``` bs config ```

- Configures the settings you want for your project, such as project name, domain name, configuration settings.


``` bs scaffold ```

- Scaffold the required workflows and configurations using code generation.


## Compile time

Compile your flutter GUI and MicroServices.

``` bs run ```

- Run it

``` bs build ```

- Build it

``` bs pack ```

- Package it.

``` bs sign ```

- Sign it.

## Run time

Run your project locally or in Kubernetes locally or in the cloud.

``` bs config flag ```

- changes your runtime configurations affected the running system using feature flags that are present in your modules configuration.

``` bs config flag -ps ```

- same as above but globally and persistent in the DB.

## Deploy time

Deploy your project to a linux server or a Kubernetes cloud.

``` bs deploy server ```

- Deploy the server

``` bs deploy app ```

- Deploys the apps to the app stores.


## Examples

You can practice and asserts your system runs correctly using the example folder.

- See the examples folder


## Recommended working environment

[✓] macOS Mojave (10.14.6) or Catalina (10.15 beta)

[✓] Flutter (Channel stable, v1.9.1+hotfix.2)

[✓] Android Studio (version 3.5)

[✓] Xcode - develop for iOS and macOS (Xcode 11.0)

[✓] iOS Simulator 11. (iOS 12.0 or iOS 13.0)
