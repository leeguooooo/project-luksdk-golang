# Changelog

## [v1.0.1] - 2025-08-12

### Fixed
- Fixed signature generation bug for `map[string]interface{}` parameters
- SDK now correctly handles interface{} typed values in map parameters
- Resolved issue causing API calls to fail with signature error (code 100005)

### Technical Details
- The bug was in the `Signature` function which incorrectly skipped interface{} types
- Master branch already contains the fix (commit: refactor(sign): 签名算法调整)
- This release makes the fix available through Go modules

## [v1.0.0] - Initial Release

- Initial release of Lucky Game Go SDK