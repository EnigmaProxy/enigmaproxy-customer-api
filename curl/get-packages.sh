#!/bin/bash

API_KEY="YOUR_API_KEY"

curl \
  -X GET \
  "https://enigmaproxy.net/api/customer/packages" \
  -H "Authorization: Bearer ${API_KEY}" \
  -H "Content-Type: application/json"