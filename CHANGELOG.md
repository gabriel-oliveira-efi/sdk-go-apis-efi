# 3.0.2

- Updated to Go Modules: Added full compatibility with the Go Modules system.
- The go.mod file was created at the project root.
- The directory structure was reorganized to remove the src/ subdirectory.
- All import paths were updated to reflect the new module structure, replacing github.com/efipay/sdk-go-apis-efi/src/efipay with github.com/efipay/sdk-go-apis-efi/efipay.
- Fixed import paths in example files that were pointing to an outdated repository (github.com/gerencianet/gn-api-sdk-go).

# 3.0.1 

- Added: new endpoints (automatic pix)
- Updated: location (creation new folder locRec)
- Fix: request (add Content-Type configuration: Now, PATCH type requests with body)

# 3.0.0

- Added: new endpoints and apis

# 1.1.1

- Fix: pix creation (createCharge and createImmediateCharge)

# 1.1.0

- Added: new endpoint (one step)

# 1.0.3

- Added: new endpoint (settle charge)
- Added: new endpoint (settle parcel carnet)

# 1.0.2

- Added: new endpoint (create charge balance sheet)

# 1.0.1

- Added: new endpoint (update plan)
- Added: new endpoint (create subscription history)
- Updated: examples

# 1.0.0

- Initial release
