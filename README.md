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
      <a href="#limitations">Limitations</a>
    </li>
    <li>
      <a href="#setup">Setup</a>
    </li>
  </ol>
</details>


---

## About The Project

`multik8s` is a tool to access pod information across multiple Kubernetes contexts at once.


## Usage
- Follow logs from pods in a given namespace :
    - `multik8s get logs [-n namespace] [-p pod] [--follow] [--tail nbr]`
- List pods in a given namespace :
    -  `multik8s get pods [-n namespace] [-p pod]`


## Limitations
Currently, the namespace should be the same on all targeted clusters.


## Setup
Tested on ARM64 using Go 1.18.3.

`make build`
