<?php

require_once "EnigmaProxyClient.php";

$client = new EnigmaProxyClient("YOUR_API_KEY");

echo "=== Packages ===\n";

$packages = $client->getPackages();

print_r($packages);

$packageId = "pkg_123456";

echo "\n=== Package: {$packageId} ===\n";

$package = $client->getPackage($packageId);

print_r($package);

echo "\n=== Generate Proxies ===\n";

$proxies = $client->generateProxy(

    packageId: $packageId,

    protocol: "http",

    country: "us",

    session: true,

    qty: 3

);

print_r($proxies);