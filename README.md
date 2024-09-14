<p align="center">
  <img src="https://raw.githubusercontent.com/cert-manager/cert-manager/d53c0b9270f8cd90d908460d69502694e1838f5f/logo/logo-small.png" height="256" width="256" alt="cert-manager project logo" />
  &nbsp;
  <img src="https://raw.githubusercontent.com/linuxserver/docker-templates/master/linuxserver.io/img/duckdns.png" height="256" width="256" alt="DuckDNS project logo" />
</p>

# DuckDNS Webhook for cert-manager

This WebHook solves the DNS01 challenge to prove ownership of DuckDNS domains.

## How to use it

TODO

## Helm Chart

TODO


### Running the test suite

All DNS providers **must** run the DNS01 provider conformance testing suite,
else they will have undetermined behaviour when used with cert-manager.

**It is essential that you configure and run the test suite when creating a
DNS01 webhook.**

An example Go test file has been provided in [main_test.go](https://github.com/cert-manager/webhook-example/blob/master/main_test.go).

You can run the test suite with:

```bash
$ TEST_ZONE_NAME=example.com. make test
```

The example file has a number of areas you must fill in and replace with your
own options in order for tests to pass.

## Acknowledgments

- Inspired by [the wonderful ebrianne's job](https://github.com/ebrianne/cert-manager-webhook-duckdns).