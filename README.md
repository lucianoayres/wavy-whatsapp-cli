# WhatsApp CLI Tool (wavy)

A command-line interface for WhatsApp messaging using the [whatsmeow](https://github.com/tulir/whatsmeow) library.

## Features

- Set up a WhatsApp session with QR code authentication
- Send messages to WhatsApp contacts via command-line
- Send messages to WhatsApp groups
- Check if phone numbers are registered on WhatsApp
- List all your WhatsApp groups

## Project Structure

```
whatsmeow-go/
├── cmd/
│   └── wavy/       # Consolidated WhatsApp CLI tool
├── bin/            # Build output directory
├── magefile.go     # Mage build system
├── go.mod          # Go module file
├── go.sum          # Go module checksum
└── README.md       # This file
```

## Installation

### Using Mage

This project uses [Mage](https://magefile.org/) for its build system.

1. **Install Mage** (if not already installed):

   ```bash
   go install github.com/magefile/mage@latest
   ```

2. **Build the tool**:

   ```bash
   mage build
   ```

   This will create `bin/wavy` executable.

3. **Install the tool system-wide** (optional):

   ```bash
   mage install
   ```

   This will:

   - Copy the binary to `/usr/local/bin/wavy`
   - Create the required config directories at `~/.config/wavy/` and `~/.local/share/wavy/`

## Usage

### Setting up WhatsApp connection

```bash
wavy setup
```

This will:

- Generate a QR code image that you need to scan with your WhatsApp mobile app
- The QR code image is automatically opened in your default image viewer
- Open WhatsApp on your phone
- Tap Menu (three dots) > Linked Devices > Link a Device
- Scan the QR code displayed by your image viewer

### Checking if a number is on WhatsApp

```bash
wavy check +1234567890
```

Or using flags:

```bash
wavy check --phone +1234567890
```

### Listing your WhatsApp groups

```bash
wavy groups
```

This will show all groups you're a member of, including their group IDs which you need for sending messages to groups.

### Sending Messages

#### To a contact:

```bash
wavy send +1234567890 "Hello from CLI"
```

Or using flags:

```bash
wavy send --to +1234567890 --msg "Hello from CLI"
```

#### To a group:

```bash
wavy send 123456789@g.us "Hello group from CLI"
```

You must use the exact group ID from the `wavy groups` command.

#### Additional options:

- `--debug` - Enable verbose debug output
- `--wait N` - Wait N seconds for message confirmation (default: 5)

Example:

```bash
wavy send --to +1234567890 --msg "Hello with debug" --debug --wait 10
```

## Data Storage

All wavy data is stored according to the XDG Base Directory Specification:

- Configuration: `~/.config/wavy/`
- Data (including WhatsApp session): `~/.local/share/wavy/`

## Maintenance

Additional Mage commands available:

- `mage clean` - Remove build artifacts
- `mage uninstall` - Completely remove wavy from your system
- `mage test` - Run tests
- `mage check` - Run linters and static analysis
