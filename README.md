# bundle2jwks - Convert an x509 CA bundle to go-jose JSONWebKeySet

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
   bundle2jwks - Convert an x509 CA bundle to go-jose JSONWebKeySet

USAGE:
   bundle2jwks [global options] [CA-BUNDLE-FILE]

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Example
```console
user@host:~$ openssl req -x509 -newkey rsa:4096 -keyout example.key -out example.crt -batch -nodes -sha256 -days 3650 -subj /CN=example
...
-----

user@host:~$ openssl x509 -noout -fingerprint -subject -issuer -dates -in example.crt
SHA1 Fingerprint=9B:FA:45:57:7E:2F:97:6A:66:14:47:41:FC:1C:6D:4A:18:1D:AF:5E
subject=CN = example
issuer=CN = example
notBefore=Dec  3 19:06:24 2024 GMT
notAfter=Dec  1 19:06:24 2034 GMT

user@host:~$ bundle2jwks example.crt
{
  "keys": [
    {
      "kty": "RSA",
      "kid": "4V6G:RPFT:5YP4:YHNF:WDEI:2F6F:JRPI:DXNT:JRDN:BEUB:4ZOO:VT4R",
      "n": "oZpnAF1kemuUTTnWoxzX0bU6NXKTwMANcN6FU-mQSrtsfZXwK7cvM432gb1-JjY2VLAIe0ibqNekE2vEyQ_CJ-AhVscl6TPxxHQutbD5CktUfWABS_V-k-F7gdFOLViA2TVqzBuBlCZ0OrbnLmCsd4vOQP1xkY5z-CZWXlnVfaV0gWWD59NqRIjeRl-O4zAX_8sA9fsDzlwovdYl_PPQ5e4jjWRuJpbY2vB_e7WAfJcWKsLFEEwQ3Lxje0ttNU5y9dEtxjWB_RoAmJ71QZS8hT0juP3_J5EfDPDXY0lGDXGf2SLWM_yYDFGwZ5WnOvzK_dudDhhf4rxRX5ZSBIzD9-9HuoYoWJ8wFvXYCis0P1NwP3f_AAGuAHLPs8ocRMorRN-aWrgAmg2-fP9SDuY05KQTejlCY091JxjRBzX_EG5A1GhBVQ5MFJDIl0us8AreMGHT5xudutnsNcRLJUXSlJQtfwWGeolLYWvifKdMaoYF-rkaPWWtFwmVkNe7C3RDU-eVYPm2-uxpKrpk3U0JES8MTgg_O6L39p9Lf_q5hz9nfX3VlMWObbbHJjChg_Pk6eHSWS76gWkZpYTCqxoS5n4RiFrp2dU385kyf83qDuqvgrkoUqGLrLSE-hYyQn5bIK6T2OdDzgQ4GpLfYO8r_0x_f7OYu2yKrIOpvRS3JTk",
      "e": "AQAB"
    }
  ]
}
```
