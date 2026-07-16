from EnigmaProxyClient import EnigmaProxyClient


def main():

    client = EnigmaProxyClient("YOUR_API_KEY")

    # ----------------------------------------
    # Show all packages
    # ----------------------------------------

    print("=== Packages ===")

    packages = client.get_packages()

    print(packages)

    # ----------------------------------------
    # Get a specific package
    # ----------------------------------------

    package_id = "pkg_123456"

    print(f"\n=== Package: {package_id} ===")

    package = client.get_package(package_id)

    print(package)

    # ----------------------------------------
    # Generate proxies
    # ----------------------------------------

    print("\n=== Generate Proxies ===")

    proxies = client.generate_proxy(

        package_id=package_id,

        protocol="http",

        country="us",

        session=True,

        qty=3

    )

    print(proxies)


if __name__ == "__main__":
    main()