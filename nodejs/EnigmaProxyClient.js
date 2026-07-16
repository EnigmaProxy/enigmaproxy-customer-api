class EnigmaProxyClient {

    constructor(apiKey, baseUrl = "https://enigmaproxy.net") {

        this.apiKey = apiKey;
        this.baseUrl = baseUrl.replace(/\/$/, "");

    }

    async request(method, endpoint, body = null) {

        const options = {

            method,

            headers: {

                "Authorization": `Bearer ${this.apiKey}`,

                "Content-Type": "application/json"

            },

            signal: AbortSignal.timeout(30000)

        };

        if (body) {

            options.body = JSON.stringify(body);

        }

        const response = await fetch(

            `${this.baseUrl}${endpoint}`,

            options

        );

        if (!response.ok) {

            throw new Error(

                `Request failed: ${response.status} ${await response.text()}`

            );

        }

        return await response.json();

    }

    // ----------------------------
    // Packages
    // ----------------------------

    async getPackages() {

        return this.request(

            "GET",

            "/api/customer/packages"

        );

    }

    async getPackage(packageId) {

        return this.request(

            "GET",

            `/api/customer/packages/${packageId}`

        );

    }

    // ----------------------------
    // Proxy Generator
    // ----------------------------

    async generateProxy({

        packageId,

        protocol = "http",

        format = "host:port:username:password",

        country,

        state,

        city,

        session = false,

        sessionTime,

        qty = 1,

        lifetime,

        fastMode = false,

        http3 = false

    }) {

        const payload = {

            packageId,

            protocol,

            format,

            qty,

            session,

            fastMode,

            http3

        };

        if (country) payload.country = country;
        if (state) payload.state = state;
        if (city) payload.city = city;
        if (sessionTime !== undefined) payload.sessionTime = sessionTime;
        if (lifetime !== undefined) payload.lifetime = lifetime;

        return this.request(

            "POST",

            "/api/customer/proxy",

            payload

        );

    }

}

module.exports = EnigmaProxyClient;