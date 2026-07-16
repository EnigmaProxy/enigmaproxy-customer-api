# EnigmaProxy Customer API

> Official EnigmaProxy Customer API libraries and examples for
> **Node.js**, **Python**, **PHP**, **Go**, and **cURL**.

![License](https://img.shields.io/badge/license-MIT-green) ![REST
API](https://img.shields.io/badge/API-REST-blue)
![Python](https://img.shields.io/badge/Python-3.10+-yellow)
![Node.js](https://img.shields.io/badge/Node.js-20+-success)
![PHP](https://img.shields.io/badge/PHP-8.1+-777BB4)
![Go](https://img.shields.io/badge/Go-1.24+-00ADD8)

------------------------------------------------------------------------

## Why this repository?

This repository provides official client libraries and ready-to-run
examples for integrating with the EnigmaProxy Customer API.

Whether you're building dashboards, automation platforms, internal
tools, AI agents, or proxy management software, these examples help you
integrate the API quickly and consistently across multiple languages.

------------------------------------------------------------------------

## Features

-   Official EnigmaProxy Customer API examples
-   Node.js
-   Python
-   PHP
-   Go
-   cURL
-   Bearer authentication
-   Package management
-   Proxy generation
-   Geo-targeting
-   Sticky sessions
-   Lightweight client libraries

------------------------------------------------------------------------

## Clone

``` bash
git clone https://github.com/EnigmaProxy/enigmaproxy-customer-api

cd enigmaproxy-customer-api
```

------------------------------------------------------------------------

## Repository Structure

``` text
enigmaproxy-customer-api/

README.md

python/
├── EnigmaProxyClient.py
├── example.py
└── requirements.txt

nodejs/
├── EnigmaProxyClient.js
└── example.js

php/
├── EnigmaProxyClient.php
├── example.php
└── composer.json

go/
├── EnigmaProxyClient.go
├── example.go
└── go.mod

curl/
├── get-packages.sh
├── get-package.sh
└── generate-proxy.sh
```

------------------------------------------------------------------------

## Requirements

-   EnigmaProxy API Key
-   Active EnigmaProxy package
-   Internet connection

------------------------------------------------------------------------

## API Endpoints

  Method   Endpoint                               Description
  -------- -------------------------------------- --------------------------
  GET      `/api/customer/packages`               List all packages
  GET      `/api/customer/packages/{packageId}`   Get package details
  POST     `/api/customer/proxy`                  Generate proxy endpoints

------------------------------------------------------------------------

## Quick Start

### Python

``` python
from EnigmaProxyClient import EnigmaProxyClient

client = EnigmaProxyClient("YOUR_API_KEY")

print(client.get_packages())
```

### Node.js

``` javascript
const EnigmaProxyClient = require("./EnigmaProxyClient");

const client = new EnigmaProxyClient("YOUR_API_KEY");

console.log(await client.getPackages());
```

### PHP

``` php
$client = new EnigmaProxyClient("YOUR_API_KEY");

print_r($client->getPackages());
```

### Go

``` go
client := New("YOUR_API_KEY")

packages, err := client.GetPackages()
```

### cURL

``` bash
curl -X GET \
"https://enigmaproxy.net/api/customer/packages" \
-H "Authorization: Bearer YOUR_API_KEY" \
-H "Content-Type: application/json"
```

------------------------------------------------------------------------

## Supported Features

-   List packages
-   Get package details
-   Generate proxies
-   HTTP & SOCKS5
-   Country targeting
-   State targeting
-   City targeting
-   Sticky sessions
-   Multiple proxy formats

------------------------------------------------------------------------

## Products

| Product | Use Case |
|---------|----------|
| Budget Residential | Cost-effective scraping & automation |
| Residential | General-purpose proxy workloads |
| Enterprise Residential | Large-scale data collection |
| Unlimited ISP | Long-running sessions |
| Static ISP | Stable persistent IPs |
| IPv6 | IPv6 connectivity |
| Datacenter | High-speed automation |-----------------------------------------------------------

## Learn More

Website:

https://enigmaproxy.net


------------------------------------------------------------------------

## Contributing

Contributions, improvements, and pull requests are welcome.

If this repository helped you, consider giving it a ⭐.

------------------------------------------------------------------------

## License

MIT License.
