#!/bin/bash

echo "Building WhatsApp CLI tools..."

echo "1. Building setup tool..."
go build -o whatsapp-setup ./cmd/setup

echo "2. Building send tool..."
go build -o whatsapp-send ./cmd/send

echo "3. Building check tool..."
go build -o whatsapp-check ./cmd/check

echo "Done! All tools built successfully."
echo ""
echo "Usage:"
echo "  ./whatsapp-setup      - Set up WhatsApp connection with QR code"
echo "  ./whatsapp-send -to \"+PHONE\" -msg \"MESSAGE\"  - Send a message"
echo "  ./whatsapp-check -phone \"+PHONE\"   - Check if a number is on WhatsApp" 