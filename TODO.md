# CLI Interface Changes

1. Consolidate all commands under a single `wavy` binary with subcommands:

- Send message (supports both flag and positional argument styles):

  ```
  wavy send -to "+1234567890" -msg "Hello from CLI"
  wavy send "+1234567890" "Hello from CLI"
  ```

- Check phone number:

  ```
  wavy check "+1234567890"
  ```

- List groups:

  ```
  wavy groups
  ```

- Setup/authentication:
  ```
  wavy setup
  ```

2. Configuration and Data Storage

Move all data files to follow XDG Base Directory Specification:

- Configuration: `~/.config/wavy/`
- Data: `~/.local/share/wavy/`
  - client.db
  - temporary QR code images
  - other runtime data

3. Build System Migration

Replace build.sh with Magefile (magefile.go) that implements:

- `mage build` - Build all components into a single `wavy` binary in the project's ./bin directory (and create the ./bin directory if it does not exist yet)
- `mage install` - Install wavy binary to /usr/local/bin
- `mage uninstall` - Completely remove wavy from your system:
  - Remove the wavy binary from /usr/local/bin
  - Delete all configuration files from ~/.config/wavy/
  - Delete all data files from ~/.local/share/wavy/
  - Clean up any remaining wavy-related files and directories
- `mage clean` - Clean build artifacts
- `mage test` - Run tests
- `mage check` - Run linters and static analysis

The install target should:

- Build the binary
- Copy it to /usr/local/bin/wavy
- Set appropriate permissions
- Create required config directories
