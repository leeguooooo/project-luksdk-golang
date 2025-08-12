# Changelog

## [v1.0.2] - 2025-08-12

### Fixed
- Fixed signature generation bug for `map[string]interface{}` parameters
- SDK now correctly handles interface{} typed values in map parameters
- Resolved issue causing API calls to fail with signature error (code 100005)

### Technical Details
- The bug was in the `Signature` function which incorrectly skipped interface{} types
- The old fix (commit 6a30238) was not migrated when refactoring to luksdk package
- This release properly implements the signature fix in the new code structure

## [v1.0.1] - 2025-08-12

### Fixed (Attempted)
- Initial attempt to fix signature generation bug
- Note: v1.0.1 was cached with old code, use v1.0.2 instead

## [v1.0.0] - Initial Release

- Initial release of Lucky Game Go SDK