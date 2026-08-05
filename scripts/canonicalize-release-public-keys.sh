#!/usr/bin/env bash
set -euo pipefail

die() {
  printf 'release public keys: %s\n' "$*" >&2
  exit 1
}

raw=${MINISIGN_PUBLIC_KEYS:-}
[[ -n "$raw" && "$raw" != "UNSET" ]] || die "MINISIGN_PUBLIC_KEYS is unset"
[[ "$raw" != "0000000000000000000000000000000000000000000000000000000000000000" ]] || die "legacy placeholder public key is forbidden"

# A public-key payload is exactly 42 bytes and therefore exactly 56 canonical
# base64 characters without padding. Accept one payload or two distinct
# payloads separated by one comma, with no whitespace or trailing separator.
[[ "$raw" =~ ^[A-Za-z0-9+/]{56}(,[A-Za-z0-9+/]{56})?$ ]] ||
  die "configure one canonical key or a two-key rotation overlap"

IFS=',' read -r -a keys <<<"$raw"

# Membership is tracked in a space-delimited string rather than an associative
# array: `declare -A` needs Bash 4, and macOS still ships Bash 3.2 as /bin/bash,
# so the release check could not run there at all. The regex above already
# restricts every key to [A-Za-z0-9+/], so no key can contain a space and the
# delimiter is unambiguous.
seen=""
for key in "${keys[@]}"; do
  case " $seen " in
    *" $key "*) die "duplicate public key" ;;
  esac

  decoded_hex=$(printf '%s' "$key" | base64 --decode 2>/dev/null | od -An -v -tx1 | tr '\n' ' ') ||
    die "public key is not valid base64"
  read -r -a decoded <<<"$decoded_hex"
  (( ${#decoded[@]} == 42 )) || die "public key payload must decode to 42 bytes"
  [[ "${decoded[0]} ${decoded[1]}" == "45 64" ]] || die "public key algorithm must be Ed"

  canonical=$(printf '%s' "$key" | base64 --decode 2>/dev/null | base64 -w0) ||
    die "public key is not valid base64"
  [[ "$canonical" == "$key" ]] || die "public key must use canonical base64 encoding"
  seen="$seen $key"
done

printf '%s\n' "$raw"
