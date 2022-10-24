# keystoregaen

![Logo](./docs/logo-readme.png)

Generate Java keystores in your browser.

[![hydrun CI](https://github.com/pojntfx/keystoregaen/actions/workflows/hydrun.yaml/badge.svg)](https://github.com/pojntfx/keystoregaen/actions/workflows/hydrun.yaml)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.18-61CFDD.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/pojntfx/keystoregaen.svg)](https://pkg.go.dev/github.com/pojntfx/keystoregaen)
[![Matrix](https://img.shields.io/matrix/keystoregaen:matrix.org)](https://matrix.to/#/#keystoregaen:matrix.org?via=matrix.org)
[![Binary Downloads](https://img.shields.io/github/downloads/pojntfx/keystoregaen/total?label=binary%20downloads)](https://github.com/pojntfx/keystoregaen/releases)

## Overview

keystoregaen is an app to generate Java keystores and compatible certificates without having to install Java or another tool that provides the JDK's `keytool`, such as the Android SDK.

## Installation

The web app is available on [GitHub releases](https://github.com/pojntfx/keystoregaen/releases) in the form of a static `.tar.gz` archive; to deploy it, simply upload it to a CDN or copy it to a web server. For most users, this shouldn't be necessary though; thanks to [@maxence-charriere](https://github.com/maxence-charriere)'s [go-app package](https://go-app.dev/), keystoregaen is a progressive web app. By simply visiting the [public deployment](https://pojntfx.github.io/keystoregaen/) once, it will be available for offline use whenever you need it:

[<img src="https://github.com/alphahorizonio/webnetesctl/raw/main/img/launch.png" width="240">](https://pojntfx.github.io/keystoregaen/)

## Screenshots

Click on an image to see a larger version.

<a display="inline" href="./docs/initial.png?raw=true">
<img src="./docs/initial.png" width="45%" alt="Screenshot of the initial screen" title="Screenshot of the initial screen">
</a>

<a display="inline" href="./docs/filled.png?raw=true">
<img src="./docs/filled.png" width="45%" alt="Screenshot of the form filled in" title="Screenshot of the form filled in">
</a>

<a display="inline" href="./docs/generating.png?raw=true">
<img src="./docs/generating.png" width="45%" alt="Screenshot of the app generating the certificate" title="Screenshot of the app generating the certificate">
</a>

## Acknowledgements

- This project would not have been possible were it not for [@maxence-charriere](https://github.com/maxence-charriere)'s [go-app package](https://go-app.dev/); if you enjoy using keygaen, please donate to him!
- The open source [PatternFly design system](https://www.patternfly.org/v4/) provides the components for the project.
- [pavlo-v-chernykh/keystore-go](https://github.com/pavlo-v-chernykh/keystore-go) provides the implementation of the JKS encoder.

To all the rest of the authors who worked on the dependencies used: **Thanks a lot!**

## Contributing

To contribute, please use the [GitHub flow](https://guides.github.com/introduction/flow/) and follow our [Code of Conduct](./CODE_OF_CONDUCT.md).

To build and start a development version of keystoregaen locally, run the following:

```shell
$ git clone https://github.com/pojntfx/keystoregaen.git
$ cd keystoregaen
$ make depend
$ make run
```

Have any questions or need help? Chat with us [on Matrix](https://matrix.to/#/#keystoregaen:matrix.org?via=matrix.org)!

## License

keystoregaen (c) 2022 Felicitas Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
