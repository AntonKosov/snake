#!/bin/bash

golangci-lint run || { echo 'lint failed' ; exit 1; }