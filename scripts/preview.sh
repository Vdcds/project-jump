#!/usr/bin/env bash

DIR="$1"

if [ -d "$DIR/.git" ]; then

  STATUS=$(git -C "$DIR" status --short)

  if [ -n "$STATUS" ]; then
    echo "󰊢 Git Status"
    echo
    echo "$STATUS"
  else
    echo "✓ Clean working tree"
    echo
    eza -T --level=2 --git-ignore "$DIR"
  fi

else
  eza -T --level=2 "$DIR"
fi
