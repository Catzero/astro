#!/bin/bash

go build
killall cmd
nohup ./cmd >> nohup.out 2>&1 &
