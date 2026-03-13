# Changelog

## [Unreleased]

### Changed
- Added `Wait()` call to wait for Echo service to be ready before returning client
- Added port exposure in container config
- Added `Port` variable for the service port
- Added `EchoUpWaitTime` variable (default: 10 seconds) to configure wait timeout
- Updated to use Go modules
