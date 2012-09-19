#!/bin/bash
# Testing script run in parallel in its own terminal

while :;do
    date +%T
    go test -bench=".*"
    sleep 1m
    echo -e "----------\n"
done
