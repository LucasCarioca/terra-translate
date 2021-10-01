# Terra-translate

![GitHub release (latest by date)](https://img.shields.io/github/v/release/LucasCarioca/terra-translate)
![GitHub Release Date](https://img.shields.io/github/release-date/LucasCarioca/terra-translate)
![GitHub all releases](https://img.shields.io/github/downloads/LucasCarioca/terra-translate/total)

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/LucasCarioca/terra-translate/Release?label=release)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/LucasCarioca/terra-translate/CI?label=CI)
[![Coverage Status](https://coveralls.io/repos/github/LucasCarioca/terra-translate/badge.svg?branch=main)](https://coveralls.io/github/LucasCarioca/terra-translate?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/LucasCarioca/terra-translate)](https://goreportcard.com/report/github.com/LucasCarioca/terra-translate)

a small utility for reading terraform plan output and used to setup basic ci/cd guard rails.

### Usage

**Basic Usage**

This just reads the output and prints out a simplified summary of the changes.
```shell
echo $(terraform plan -json) | terraform read
# If using fish...
echo (terraform plan -json) | terraform read
```

**Destroy Guard Usage**

This option will not only print out the summary but also exit with code 1 in order to be used as a way to abort a CI/CD pipeline.
```shell
echo $(terraform plan -json) | terraform guard -d
# If using fish...
echo (terraform plan -json) | terraform guard -d
```

**Guard Options**

- `terra-translate guard -a` will exit with code 1 if any additional resources would be created
- `terra-translate guard -c` will exit with code 1 if any existing resources would be changed
- `terra-translate guard -d` will exit with code 1 if any existing resources would be destroyed

these flags can also be combined. for example the below snippit would only allow changes to existing resources but no additional resources or destructive changes.: 
```shell
echo $(terraform plan -json) | terraform guard -d -a 
# If using fish...
echo (terraform plan -json) | terraform guard -d -a 
```