#!/bin/bash

API_KEY="YOUR_API_KEY"

PACKAGE_ID="pkg_123456"

curl \
  -X GET \
  "https://enigmaproxy.net/api/customer/packages/${PACKAGE_ID}" \
  -H "Authorization: Bearer ${API_KEY}" \
  -H "Content-Type: application/json"