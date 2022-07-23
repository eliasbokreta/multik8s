<!-- TITLE -->
<br />
<div align="center">
  <img src="assets/logo.png" alt="Logo" width="80" height="80">
  <h1 align="center">multik8s</h3>
  <p align="center">
     A tool to access pod information across multiple Kubernetes contexts at once.
  </p>
</div>

<div align="center">

[![CI workflow](https://github.com/eliasbokreta/multik8s/actions/workflows/main.yml/badge.svg)](https://github.com/eliasbokreta/multik8s/actions/workflows/main.yml/badge.svg)

</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#usage">Usage</a>
    </li>
    <li>
      <a href="#setup">Setup</a>
    </li>
  </ol>
</details>


---

## About The Project

`multik8s` is a tool to access pod information across multiple Kubernetes contexts at once.\
It allows to **tail logs** from several pods in several contexts and also to **list pods**.


## Usage
Tail logs from pods :
```
multik8s get logs [-n namespace] [-p podName] [--follow] [--tail nbr]

-n, --namespace string   Kubernetes namespace (should be the exact name) (default "default")
-p, --podname string     Kubernetes pod name (works as a wildcard)
-f, --follow             Choose whether or not to follow log stream
-t, --tail int           The number of lines from the end of the logs to show (default 5)
```

List pods :
```
multik8s get pods [-n namespace] [-p podName]

-n, --namespace string   Kubernetes namespace (should be the exact name) (default "default")
-p, --podname string     Kubernetes pod name (works as a wildcard)
```

## Setup
- ### Build locally :

```
make build
```

- ### Install the latest version :
  - Download the latest tar released from [Github](https://github.com/eliasbokreta/multik8s/releases)
  - Extract the archive
  - Move the binary to the location of your choice

- ### Update the binary to the latest version :
```
multik8s update
```
