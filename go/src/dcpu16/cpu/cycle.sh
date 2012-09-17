#!/bin/bash
# Testing script run in parallel in its own terminal

while :;do date +%T;go test; sleep 1m; echo -e "----------\n"; done
