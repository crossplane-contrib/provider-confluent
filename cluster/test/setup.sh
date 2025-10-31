#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"

if [[ -z "${UPTEST_CLOUD_CREDENTIALS:-}" ]]; then
  echo "Error: UPTEST_CLOUD_CREDENTIALS is not set!" >&2
  exit 1
fi

echo "Creating Confluent credentials secret 'confluent-creds'..."
${KUBECTL} -n "${CROSSPLANE_NAMESPACE}" create secret generic confluent-creds \
  --from-literal=confluent_credentials="${UPTEST_CLOUD_CREDENTIALS}" \
  --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Creating 'default' ProviderConfig..."
cat << EOF | ${KUBECTL} apply -f -
apiVersion: confluent.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    secretRef:
      key: confluent_credentials
      name: confluent-creds
      namespace: $CROSSPLANE_NAMESPACE
    source: Secret
EOF

echo "Waiting until all provider packages are healthy..."
"${KUBECTL}" wait provider.pkg --all --for condition=Healthy --timeout 10m

echo "Waiting for all pods to come online..."
"${KUBECTL}" -n "${CROSSPLANE_NAMESPACE}" wait --for=condition=Available deployment --all --timeout=5m