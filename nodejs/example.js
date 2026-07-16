const EnigmaProxyClient = require("./EnigmaProxyClient");

async function main() {

    const client = new EnigmaProxyClient(

        "YOUR_API_KEY"

    );

    // ----------------------------------------
    // Show all packages
    // ----------------------------------------

    console.log("=== Packages ===");

    const packages = await client.getPackages();

    console.log(packages);

    // ----------------------------------------
    // Get package details
    // ----------------------------------------

    const packageId = "pkg_123456";

    console.log(`\n=== Package: ${packageId} ===`);

    const packageInfo = await client.getPackage(

        packageId

    );

    console.log(packageInfo);

    // ----------------------------------------
    // Generate proxies
    // ----------------------------------------

    console.log("\n=== Generate Proxies ===");

    const proxies = await client.generateProxy({

        packageId,

        protocol: "http",

        country: "us",

        session: true,

        qty: 3

    });

    console.log(proxies);

}

main().catch(console.error);