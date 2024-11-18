#!/bin/bash

dir="$(dirname $0)"
base="$(basename $0 .sh)"
name="$dir/$base"

coproc gui { # $gui[0] coprocs STDOUT; $gui[1] coprocs STDIN
  "$name.uidl"
}
gui_pid="$!"

read -u "${gui[0]}" json
echo "JSON: $json"
if wait -f "$gui_pid"; then
  name=$(echo "$json" | jq -r '.name')
  email=$(echo "$json" | jq -r '.email')
  gender=$(echo "$json" | jq -r '.gender')
  bio=$(echo "$json" | jq -r '.bio')
  echo "name: $name; email: $email; gender: $gender; bio: $bio"
  echo "INSERT INTO person (name, email, gender, bio) VALUES ('$name', '$email', $gender, '$bio')"
  ret=0
else
  ret=$?
fi
echo "RET: $ret"
