#!/bin/sh

./clean.sh

antlr4 -no-listener -no-visitor -package uidl UIDL.g4
