#!/bin/bash

echo "Building WhatsApp CLI tools..."

echo "1. Building setup tool..."
go build -o whatsapp-setup ./cmd/setup

echo "2. Building send tool..."
go build -o whatsapp-send ./cmd/send

echo "3. Building check tool..."
go build -o whatsapp-check ./cmd/check

echo "4. Building groups tool..."
go build -o whatsapp-groups ./cmd/groups

echo "Done! All tools built successfully."
echo ""
echo "Usage:"
echo "  ./whatsapp-setup      - Set up WhatsApp connection with QR code"
echo "  ./whatsapp-send -to \"+PHONE\" -msg \"MESSAGE\"  - Send a message to a contact"
echo "  ./whatsapp-send -to \"GROUP_ID@g.us\" -msg \"MESSAGE\"  - Send a message to a group"
echo "  ./whatsapp-check -phone \"+PHONE\"   - Check if a number is on WhatsApp"
echo "  ./whatsapp-groups     - List all your WhatsApp groups" 