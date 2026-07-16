package main

import (
	"fmt"
)

func main() {

	client := New(

		"YOUR_API_KEY",
	)

	fmt.Println("=== Packages ===")

	packages, err := client.GetPackages()

	if err != nil {

		panic(err)

	}

	fmt.Println(packages)

	packageID := "pkg_123456"

	fmt.Println("\n=== Package ===")

	pkg, err := client.GetPackage(packageID)

	if err != nil {

		panic(err)

	}

	fmt.Println(pkg)

	fmt.Println("\n=== Generate Proxies ===")

	proxies, err := client.GenerateProxy(

		GenerateProxyOptions{

			PackageID: packageID,

			Protocol: "http",

			Country: "us",

			Session: true,

			Qty: 3,
		},
	)

	if err != nil {

		panic(err)

	}

	fmt.Println(proxies)

}
