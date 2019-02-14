# Extension API's for the Go Pluggable Transport Library

This package contains some extension API's for the Go Pluggable Transport
Library available at https://gitweb.torproject.org/pluggable-transports/goptlib.git/

## Extension API's

Available since Tor version 0.4.1-alpha:

- Support for Tor's `STATUS` API found in [Tor's pt-spec.txt][ptspec].
- Standard `log.Logger` API for logging via Tor's subprocess logging mechanism
  found in [Tor's pt-spec.txt][ptspec].

[ptspec]: https://gitweb.torproject.org/torspec.git/tree/pt-spec.txt
