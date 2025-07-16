# WhatsApp CLI Sender

A simple command-line tool for sending WhatsApp messages using the [whatsmeow](https://github.com/tulir/whatsmeow) library.

## Setup

1. Install dependencies:

   ```bash
   go get go.mau.fi/whatsmeow
   go get github.com/mattn/go-sqlite3
   ```

2. Build the tools:

   ```bash
   go build -o whatsapp-setup ./cmd/setup
   go build -o whatsapp-send ./cmd/send
   ```

3. Run the setup to authenticate:

   ```bash
   ./whatsapp-setup
   ```

   This will generate a QR code that you need to scan with your WhatsApp mobile app:

   - Open WhatsApp on your phone
   - Tap Menu (three dots) > Linked Devices > Link a Device
   - Scan the QR code displayed in the terminal

4. Once authenticated, the setup will complete and you're ready to send messages.

## Sending Messages

Use the `whatsapp-send` tool with `-to` and `-msg` flags:

```bash
./whatsapp-send -to "+5511999999999" -msg "Hello from CLI"
```

- The `-to` flag should contain the recipient's phone number in international format.
- The `-msg` flag contains the message you want to send.

## Notes

- The WhatsApp session is stored in `client.db` in the current directory.
- You only need to run the setup process once (or if your session expires).
- This tool works with WhatsApp's multi-device functionality.
