#!/bin/bash

dir="$(dirname $0)"
base="$(basename $0 .sh)"
name="$dir/$base"

coproc gui { # $gui[0] coprocs STDOUT; $gui[1] coprocs STDIN
  "$name.uidl"
}
gui_pid="$!"
echo "gui_pid: $gui_pid"
jobs -l

read -u "${gui[0]}" json
echo "JSON: $json"

wait -f "$gui_pid"
ret=$?
echo "RET: $ret"