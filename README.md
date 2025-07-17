# 🌊💬 **Wavy**

## Send WhatsApp messages from your command line, entirely free

Wavy is a lightweight command‑line interface for WhatsApp messaging powered by the [whatsmeow](https://github.com/tulir/whatsmeow) library. It lets you authenticate with a QR code, send messages to individuals or groups, verify phone numbers, and browse your group list, all from your terminal.

[Key features](#key-features) | [Project Structure](#project-structure) | [Installation](#installation) | [Usage](#usage) | [Data Storage](#data-storage) | [Viewing WhatsApp Contact Data](#viewing-whatsapp-contact-data) | [Maintenance](#maintenance) | [License](#license)

## Key features

- 🔒 **QR code authentication**
  Pair your account by scanning a QR code that opens in your image viewer
- 💬 **Send messages to contacts**
  Deliver plain text or formatted messages to any registered WhatsApp number
- 👥 **Send messages to groups**
  Post updates in your WhatsApp groups by using their group IDs
- ✅ **Verify registration status**
  Check if a phone number is registered on WhatsApp before sending a message
- 📋 **List your groups**
  Retrieve all your group chats and their IDs to target them easily

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
   sudo mage install
   ```

   This will:

   - Copy the binary to `/usr/local/bin/wavy`
   - Create the required config directories at `~/.config/wavy/` and `~/.local/share/wavy/`

## Usage

**WhatsApp connection setup**

1. Open your terminal and run:

   ```bash
   wavy setup
   ```

2. A QR code will be generated and displayed in your image viewer.
3. On your phone, open WhatsApp and navigate to
   **Settings > Linked Devices > Link a Device**
4. Scan the QR code on your computer screen.
5. Once pairing is successful, WhatsApp will confirm the new device is connected. You’re now authenticated and ready to send messages.
6. In your terminal, press **Ctrl+C** to exit the setup script.

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

## Viewing WhatsApp Contact Data

The WhatsApp session data is stored in a SQLite database at `~/.local/share/wavy/client.db`. You can inspect this database to view your contacts and other information:

### Installing SQLite CLI

```bash
# For Debian/Ubuntu
sudo apt install sqlite3

# For Fedora
sudo dnf install sqlite

# For macOS
brew install sqlite
```

### Viewing Contact Data

1. Open the database:

   ```bash
   sqlite3 ~/.local/share/wavy/client.db
   ```

2. View available tables:

   ```sql
   .tables
   ```

3. View your contacts:

   ```sql
   SELECT * FROM whatsmeow_contacts;
   ```

4. View specific contact information:

   ```sql
   SELECT jid, name, first_name, push_name FROM whatsmeow_contacts;
   ```

5. Count your contacts:

   ```sql
   SELECT COUNT(*) FROM whatsmeow_contacts;
   ```

6. Find a specific contact by name:

   ```sql
   SELECT * FROM whatsmeow_contacts WHERE name LIKE '%John%';
   ```

7. Exit SQLite:

   ```
   .exit
   ```

### Using SQLite in Script Mode

You can also query the database directly from the command line:

```bash
sqlite3 ~/.local/share/wavy/client.db "SELECT jid, name FROM whatsmeow_contacts"
```

## Maintenance

Additional Mage commands available:

- `mage clean` - Remove build artifacts
- `mage uninstall` - Completely remove wavy from your system
- `mage test` - Run tests
- `mage check` - Run linters and static analysis

## License

Licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
