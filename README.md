# WhatsApp CLI Tools

A set of command-line tools for WhatsApp messaging using the [whatsmeow](https://github.com/tulir/whatsmeow) library.

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
│   ├── check/     # WhatsApp number verification tool
│   ├── groups/    # WhatsApp groups listing tool
│   ├── send/      # Message sending tool
│   └── setup/     # Authentication setup tool
├── build.sh       # Build script for all tools
├── .gitignore     # Git ignore file
├── go.mod         # Go module file
├── go.sum         # Go module checksum
└── README.md      # This file
```

## Setup

1. **Build the tools**:

   ```bash
   # Using the build script
   ./build.sh

   # Or build manually
   go build -o whatsapp-setup ./cmd/setup
   go build -o whatsapp-send ./cmd/send
   go build -o whatsapp-check ./cmd/check
   go build -o whatsapp-groups ./cmd/groups
   ```

2. **Run the setup to authenticate**:

   ```bash
   ./whatsapp-setup
   ```

   This will generate a QR code image that you need to scan with your WhatsApp mobile app:

   - Open WhatsApp on your phone
   - Tap Menu (three dots) > Linked Devices > Link a Device
   - Scan the QR code displayed by your image viewer

3. Once authenticated, the setup will complete and you're ready to send messages.

## Usage

### Checking if a number is on WhatsApp

```bash
./whatsapp-check -phone "+1234567890"
```

### Listing your WhatsApp groups

```bash
./whatsapp-groups
```

This will show all groups you're a member of, including their group IDs which you need for sending messages to groups.

### Sending Messages

#### To a contact:

```bash
./whatsapp-send -to "+1234567890" -msg "Hello from CLI"
```

#### To a group:

```bash
./whatsapp-send -to "123456789@g.us" -msg "Hello group from CLI"
```

You must use the exact group ID from the `whatsapp-groups` command.

#### Additional options:

- `-debug` - Enable verbose debug output
- `-wait N` - Wait N seconds for message confirmation (default: 5)

Example:

```bash
./whatsapp-send -to "+1234567890" -msg "Hello with debug" -debug -wait 10
```

## Notes

- The WhatsApp session is stored in `client.db` in the current directory.
- You only need to run the setup process once (or if your session expires).
- This tool works with WhatsApp's multi-device functionality.
- If a message is not delivered, try using the `-debug` flag for more information.
