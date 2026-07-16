<?php

class EnigmaProxyClient
{
    private string $apiKey;
    private string $baseUrl;

    public function __construct(
        string $apiKey,
        string $baseUrl = "https://enigmaproxy.net"
    ) {
        $this->apiKey = $apiKey;
        $this->baseUrl = rtrim($baseUrl, "/");
    }

    private function request(
        string $method,
        string $endpoint,
        ?array $body = null
    ): array {

        $ch = curl_init();

        $headers = [
            "Authorization: Bearer {$this->apiKey}",
            "Content-Type: application/json"
        ];

        curl_setopt_array($ch, [

            CURLOPT_URL => $this->baseUrl . $endpoint,

            CURLOPT_RETURNTRANSFER => true,

            CURLOPT_CUSTOMREQUEST => $method,

            CURLOPT_HTTPHEADER => $headers,

            CURLOPT_TIMEOUT => 30

        ]);

        if ($body !== null) {

            curl_setopt(
                $ch,
                CURLOPT_POSTFIELDS,
                json_encode($body)
            );

        }

        $response = curl_exec($ch);

        if ($response === false) {

            throw new Exception(curl_error($ch));

        }

        $status = curl_getinfo($ch, CURLINFO_HTTP_CODE);

        curl_close($ch);

        if ($status >= 400) {

            throw new Exception(
                "Request failed ({$status}): {$response}"
            );

        }

        return json_decode($response, true);

    }

    // ----------------------------
    // Packages
    // ----------------------------

    public function getPackages(): array
    {
        return $this->request(
            "GET",
            "/api/customer/packages"
        );
    }

    public function getPackage(
        string $packageId
    ): array {

        return $this->request(
            "GET",
            "/api/customer/packages/{$packageId}"
        );

    }

    // ----------------------------
    // Proxy Generator
    // ----------------------------

    public function generateProxy(

        string $packageId,

        string $protocol = "http",

        string $format = "host:port:username:password",

        ?string $country = null,

        ?string $state = null,

        ?string $city = null,

        bool $session = false,

        ?int $sessionTime = null,

        int $qty = 1,

        ?int $lifetime = null,

        bool $fastMode = false,

        bool $http3 = false

    ): array {

        $payload = [

            "packageId" => $packageId,

            "protocol" => $protocol,

            "format" => $format,

            "qty" => $qty,

            "session" => $session,

            "fastMode" => $fastMode,

            "http3" => $http3

        ];

        if ($country !== null)
            $payload["country"] = $country;

        if ($state !== null)
            $payload["state"] = $state;

        if ($city !== null)
            $payload["city"] = $city;

        if ($sessionTime !== null)
            $payload["sessionTime"] = $sessionTime;

        if ($lifetime !== null)
            $payload["lifetime"] = $lifetime;

        return $this->request(

            "POST",

            "/api/customer/proxy",

            $payload

        );

    }

}