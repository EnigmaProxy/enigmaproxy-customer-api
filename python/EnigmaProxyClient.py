import requests
from typing import Optional


class EnigmaProxyClient:

    def __init__(
        self,
        api_key: str,
        base_url: str = "https://enigmaproxy.net"
    ):

        self.base_url = base_url.rstrip("/")
        self.session = requests.Session()

        self.session.headers.update({
            "Authorization": f"Bearer {api_key}",
            "Content-Type": "application/json"
        })

    def _request(
        self,
        method: str,
        endpoint: str,
        **kwargs
    ):

        response = self.session.request(
            method=method,
            url=f"{self.base_url}{endpoint}",
            timeout=30,
            **kwargs
        )

        response.raise_for_status()

        return response.json()

    # ----------------------------
    # Packages
    # ----------------------------

    def get_packages(self):
        """
        Get all customer packages.
        """

        return self._request(
            "GET",
            "/api/customer/packages"
        )

    def get_package(
        self,
        package_id: str
    ):
        """
        Get information about a package.
        """

        return self._request(
            "GET",
            f"/api/customer/packages/{package_id}"
        )

    # ----------------------------
    # Proxy Generator
    # ----------------------------

    def generate_proxy(

        self,

        package_id: str,

        protocol: str = "http",

        format: str = "host:port:username:password",

        country: Optional[str] = None,

        state: Optional[str] = None,

        city: Optional[str] = None,

        session: bool = False,

        session_time: Optional[int] = None,

        qty: int = 1,

        lifetime: Optional[int] = None,

        fast_mode: bool = False,

        http3: bool = False

    ):

        payload = {

            "packageId": package_id,

            "protocol": protocol,

            "format": format,

            "qty": qty,

            "session": session,

            "fastMode": fast_mode,

            "http3": http3

        }

        if country:
            payload["country"] = country

        if state:
            payload["state"] = state

        if city:
            payload["city"] = city

        if session_time is not None:
            payload["sessionTime"] = session_time

        if lifetime is not None:
            payload["lifetime"] = lifetime

        return self._request(
            "POST",
            "/api/customer/proxy",
            json=payload
        )