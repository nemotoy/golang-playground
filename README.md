# golang-playground

[![CircleCI](https://circleci.com/gh/nemotoy/golang-playground.svg?style=svg)](https://circleci.com/gh/nemotoy/golang-playground)
[![codecov](https://codecov.io/gh/nemotoy/golang-playground/branch/master/graph/badge.svg)](https://codecov.io/gh/nemotoy/golang-playground)

## CircleCI on local

```sh
# validates a configuration file
circleci config validate .circleci/config.yml

# processes the configuration file of version 2.1 into another file of version 2.0
circleci config process .circleci/config.yml > ci_local.yml

# executes a job
circleci build --job test -c ci_local.yml
```
