#!/bin/bash

API_KEY="YOUR_API_KEY"

PACKAGE_ID="pkg_123456"

curl \
  -X POST \
  "https://enigmaproxy.net/api/customer/proxy" \
  -H "Authorization: Bearer ${API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{
    "packageId": "'"${PACKAGE_ID}"'",
    "protocol": "http",
    "country": "us",
    "session": true,
    "qty": 3
}'