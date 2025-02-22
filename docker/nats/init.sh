#! /bin/sh

set -x
NATS_URL="${NATS_URL:-nats://localhost:4222}"
NATS_USER="${NATS_USER:-}"
NATS_PASSWORD="${NATS_PASSWORD:-}"
KV="${KV:-ton}"
set +x

echo "Creating $KV"
nats kv add $KV