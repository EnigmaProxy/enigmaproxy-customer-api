# EnigmaProxy Customer API Examples

> Official EnigmaProxy Customer API client examples for cURL, Go,
> Node.js, PHP & Python. Learn how to list your packages, fetch
> package details, and generate residential, ISP, and datacenter
> proxies through the EnigmaProxy Customer API.

![License](https://img.shields.io/badge/license-MIT-green)
![Go](https://img.shields.io/badge/Go-1.24+-00ADD8)
![Node.js](https://img.shields.io/badge/Node.js-18+-success)
![PHP](https://img.shields.io/badge/PHP-8.1+-777BB4)
![Python](https://img.shields.io/badge/Python-3.10+-yellow)

------------------------------------------------------------------------

## Why this repository?

The EnigmaProxy Customer API lets you manage your account
programmatically: list the packages you own, inspect a single
package, and generate proxy credentials on demand for use in
scrapers, browser automation, QA tooling, or any workflow that needs
fresh residential, ISP, or datacenter exit IPs.

This repository provides a minimal, dependency-light client for each
language, plus a runnable example and raw `curl` scripts so you can
see exactly what's sent over the wire. Every client exposes the same
three operations with the same parameters and defaults, so you can
switch languages without relearning the API.

Examples cover:

-   Listing all packages on your account
-   Fetching a single package by ID
-   Generating proxies (protocol, format, geo-targeting, sessions, quantity, lifetime)

------------------------------------------------------------------------

# Table of Contents

1.  Installation
2.  Repository Structure
3.  Authentication
4.  Running the Examples
5.  Listing Packages
6.  Getting a Single Package
7.  Generating Proxies
8.  Geo-Targeting & Sticky Sessions
9.  Error Handling
10. Best Practices
11. FAQ
12. Production Proxies

------------------------------------------------------------------------

# Installation

## cURL

No installation required beyond `curl` itself.

## Go

``` bash
cd go
go build ./...
```

## Node.js

``` bash
cd nodejs
```

Uses the built-in `fetch` API — no dependencies to install.

## PHP

``` bash
cd php
composer install
```

Uses PHP's built-in `curl` extension.

## Python

``` bash
cd python
pip install -r requirements.txt
```

`requirements.txt` includes `requests`.

------------------------------------------------------------------------

# Repository Structure

``` text
curl/
  get-packages.sh     List all packages on your account
  get-package.sh      Get a single package by ID
  generate-proxy.sh   Generate proxies for a package

go/
  EnigmaProxyClient.go   Client: GetPackages, GetPackage, GenerateProxy
  example.go             Runnable example
  go.mod

nodejs/
  EnigmaProxyClient.js   Client: getPackages, getPackage, generateProxy
  example.js             Runnable example

php/
  EnigmaProxyClient.php  Client: getPackages, getPackage, generateProxy
  example.php            Runnable example
  composer.json

python/
  EnigmaProxyClient.py   Client: get_packages, get_package, generate_proxy
  example.py             Runnable example
  requirements.txt
```

Every client implements the same three methods against the same
endpoints:

``` text
GET  /api/customer/packages          List all packages
GET  /api/customer/packages/{id}     Get a single package
POST /api/customer/proxy             Generate proxies
```

------------------------------------------------------------------------

# Authentication

All requests are authenticated with a Bearer token sent as an
`Authorization` header:

``` text
Authorization: Bearer YOUR_API_KEY
```

Replace `YOUR_API_KEY` in each example with your EnigmaProxy API key.
Never hardcode real credentials — prefer an environment variable:

``` bash
export ENIGMAPROXY_API_KEY=...
```

------------------------------------------------------------------------

# Running the Examples

## cURL

``` bash
bash curl/get-packages.sh
bash curl/get-package.sh
bash curl/generate-proxy.sh
```

## Go

``` bash
cd go
go run .
```

## Node.js

``` bash
node nodejs/example.js
```

## PHP

``` bash
php php/example.php
```

## Python

``` bash
python python/example.py
```

------------------------------------------------------------------------

# Listing Packages

Go (`go/example.go`)

``` go
client := New("YOUR_API_KEY")

packages, err := client.GetPackages()
```

Node.js (`nodejs/example.js`)

``` javascript
const client = new EnigmaProxyClient("YOUR_API_KEY");

const packages = await client.getPackages();
```

PHP (`php/example.php`)

``` php
$client = new EnigmaProxyClient("YOUR_API_KEY");

$packages = $client->getPackages();
```

Python (`python/example.py`)

``` python
client = EnigmaProxyClient("YOUR_API_KEY")

packages = client.get_packages()
```

cURL (`curl/get-packages.sh`)

``` bash
curl -X GET "https://enigmaproxy.net/api/customer/packages" \
  -H "Authorization: Bearer ${API_KEY}" \
  -H "Content-Type: application/json"
```

------------------------------------------------------------------------

# Getting a Single Package

Go

``` go
pkg, err := client.GetPackage("pkg_123456")
```

Node.js

``` javascript
const pkg = await client.getPackage("pkg_123456");
```

PHP

``` php
$package = $client->getPackage("pkg_123456");
```

Python

``` python
package = client.get_package("pkg_123456")
```

cURL (`curl/get-package.sh`)

``` bash
curl -X GET "https://enigmaproxy.net/api/customer/packages/${PACKAGE_ID}" \
  -H "Authorization: Bearer ${API_KEY}" \
  -H "Content-Type: application/json"
```

------------------------------------------------------------------------

# Generating Proxies

Every client accepts the same options, all with the same defaults:

``` text
packageId     required
protocol      default "http"
format        default "host:port:username:password"
country       optional
state         optional
city          optional
session       default false
sessionTime   optional, only used with session
qty           default 1
lifetime      optional
fastMode      default false
http3         default false
```

Go

``` go
proxies, err := client.GenerateProxy(GenerateProxyOptions{
    PackageID: "pkg_123456",
    Protocol:  "http",
    Country:   "us",
    Session:   true,
    Qty:       3,
})
```

Node.js

``` javascript
const proxies = await client.generateProxy({
    packageId: "pkg_123456",
    protocol: "http",
    country: "us",
    session: true,
    qty: 3
});
```

PHP

``` php
$proxies = $client->generateProxy(
    packageId: "pkg_123456",
    protocol: "http",
    country: "us",
    session: true,
    qty: 3
);
```

Python

``` python
proxies = client.generate_proxy(
    package_id="pkg_123456",
    protocol="http",
    country="us",
    session=True,
    qty=3
)
```

cURL (`curl/generate-proxy.sh`)

``` bash
curl -X POST "https://enigmaproxy.net/api/customer/proxy" \
  -H "Authorization: Bearer ${API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{
    "packageId": "'"${PACKAGE_ID}"'",
    "protocol": "http",
    "country": "us",
    "session": true,
    "qty": 3
}'
```

------------------------------------------------------------------------

# Geo-Targeting & Sticky Sessions

Pass `country`, `state`, and/or `city` to target a location, and set
`session: true` (with an optional `sessionTime`) to keep the same
exit IP across multiple requests:

``` text
country only              United States
country + state           United States, California
country + state + city     United States, California, Los Angeles
session: true              Sticky session, random exit IP
country + session: true    Sticky session, geo-targeted
```

`qty` controls how many proxy credentials are generated in a single
call, and `lifetime` controls how long generated credentials remain
valid, when supported by your package.

------------------------------------------------------------------------

# Error Handling

Every client raises/returns an error for non-2xx responses, including
the response body so you can inspect the API's error message:

-   Go: `GenerateProxy`, `GetPackage`, and `GetPackages` return `(any, error)` — always check `err`.
-   Node.js: `request()` throws an `Error` containing the status code and response body.
-   PHP: `request()` throws an `Exception` containing the status code and response body.
-   Python: `_request()` calls `response.raise_for_status()`, raising `requests.HTTPError`.

All clients apply a 30-second request timeout by default.

------------------------------------------------------------------------

# Best Practices

-   Store your API key in an environment variable, never in source control
-   Reuse a single client instance instead of creating a new one per request
-   Handle errors from every call — network and API failures are both possible
-   Use `qty` to batch-generate proxies instead of calling the endpoint in a loop
-   Use sticky sessions for any multi-step workflow that must keep the same exit IP

------------------------------------------------------------------------

# FAQ

### Which language should I use?

Whichever matches your existing stack — all five examples call the
same endpoints with the same parameters and defaults.

### Do I need an SDK to use the Customer API?

No. The `curl/` scripts show the raw HTTP requests; the language
clients are thin, dependency-light wrappers around the same calls.

### Can I generate more than one proxy per request?

Yes, set `qty` to the number of proxy credentials you need.

### How do I keep the same exit IP across requests?

Set `session: true` (or `session=True` in Python), and optionally
`sessionTime` to control how long the session is kept alive.

### Can I combine geo-targeting with sticky sessions?

Yes — pass `country` / `state` / `city` together with `session: true`.

------------------------------------------------------------------------

# Production Proxy Infrastructure

The Customer API works alongside EnigmaProxy's proxy network. If
you're looking for production-ready infrastructure, EnigmaProxy
offers:

-   Budget Residential Proxies
-   Residential Proxies
-   Enterprise Residential Proxies
-   Unlimited ISP
-   Static ISP
-   IPv6
-   Datacenter
-   Coverage in 200+ countries
-   HTTP & SOCKS5
-   Rotating & Sticky Sessions
-   Dashboard & API

Products:

-   https://enigmaproxy.net/?plan=budget_residential
-   https://enigmaproxy.net/pricing?plan=residential
-   https://enigmaproxy.net/pricing?plan=enterprise_residential
-   https://enigmaproxy.net/pricing?plan=unlimited_isp
-   https://enigmaproxy.net/pricing?plan=isp
-   https://enigmaproxy.net/pricing?plan=ipv6
-   https://enigmaproxy.net/pricing?plan=datacenter

Website:

https://enigmaproxy.net

------------------------------------------------------------------------

# License

MIT License.

Feel free to fork, improve, and contribute.
