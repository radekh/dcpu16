#!/bin/bash
# 

LOGDIR=/tmp
LOG=$LOGDIR/dcpu16b.log

echo "Running $0" >>$LOG
id >>$LOG


cd /home/radek/work/www-data/dcpu16
git pull
