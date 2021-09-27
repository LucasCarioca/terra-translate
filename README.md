# Terra-translate

![GitHub release (latest by date)](https://img.shields.io/github/v/release/LucasCarioca/terra-translate)
![GitHub Release Date](https://img.shields.io/github/release-date/LucasCarioca/terra-translate)
![GitHub all releases](https://img.shields.io/github/downloads/LucasCarioca/terra-translate/total)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/LucasCarioca/terra-translate/Release?label=release)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/LucasCarioca/terra-translate/CI?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/LucasCarioca/terra-translate)](https://goreportcard.com/report/github.com/LucasCarioca/terra-translate)

a small utility for reading terraform plan output and used to setup basic ci/cd guard rails.

### Usage

**Basic Usage**

This just reads the output and prints out a simplified summary of the changes.
```shell
$ echo (terraform plan -json) | terra-translate
operation: plan
changes: 0
add: 1
destroy: 1
```

**Destroy Guard Usage**

This option will not only print out the summary but also exit with code 1 in order to be used as a way to abort a CI/CD pipeline.
```shell
$ echo (terraform plan -json) | terra-translate -destroy-guard
operation: plan
changes: 0
add: 1
destroy: 1
WARNING: 1 destructive change(s) detected!
```