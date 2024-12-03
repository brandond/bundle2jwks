# bundle2jwks - Convert a x509 CA bundle to go-jose JSONWebKeySet

[distribution/registry](https://github.com/distribution/distribution) - the
reference implimentation of an OCI registry - supports loading auth token
issuer trusted CAs from a file. Unfortunately, as of the v3 release they have
apparently chosen to break compatibility with auth providers that use
libtrust-format JWT key IDs. This is the only key ID format that works with
distribution v2, and has been the de-facto standard for over a decade.

Ref:
* https://github.com/distribution/distribution/issues/4470
* https://github.com/distribution/distribution/discussions/4487
* https://github.com/distribution/distribution/pull/4521

In order to support auth providers that still use this key ID format, a JSON
JWKS file must be provided to the registry server, via the
`REGISTRY_AUTH_TOKEN_JWKS` env var or corresponding YAML key. This is in
addition to still providing the CA bundle path in
`REGISTRY_AUTH_TOKEN_ROOTCERTBUNDLE`.

This is a minimal tool to convert a CA bundle to JWKS JSON, using libtrust-format key IDs.

## Help
```
NAME:
   bundle2jwks - Convert a x509 CA bundle to go-jose JSONWebKeySet

USAGE:
   bundle2jwks [global options] [CA-BUNDLE-FILE]

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
