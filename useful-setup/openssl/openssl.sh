#!/bin/bash

# Variables for the certificate details
COUNTRY="US"                    # Country Code (2 characters)
STATE="California"              # State or Province Name
LOCALITY="San Francisco"         # City or Locality Name
ORGANIZATION="My Company Inc"    # Organization Name
COMMON_NAME="*.example.com"       # Wildcard Common Name (domain name)

# Output files
PRIVATE_KEY="private.key"
CSR_FILE="request.csr"

# Generate the CSR and private key
echo "Generating private key and CSR..."
openssl req -new -newkey rsa:2048 -nodes -keyout $PRIVATE_KEY -out $CSR_FILE \
    -subj "/C=$COUNTRY/ST=$STATE/L=$LOCALITY/O=$ORGANIZATION/CN=$COMMON_NAME"

echo "CSR and private key have been created: $CSR_FILE and $PRIVATE_KEY"


PRIVATE_KEY="private.key"
CSR_FILE="request.csr"
CERTIFICATE="selfsigned.crt"

# Self-sign the CSR to create a self-signed certificate
echo "Self-signing the CSR to create a self-signed certificate..."
openssl x509 -req -in $CSR_FILE -signkey $PRIVATE_KEY -out $CERTIFICATE -days 365

echo "Self-signed certificate has been created: $CERTIFICATE"


# # Generate a private key
# echo "Generating private key..."
# openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -out $PRIVATE_KEY

# # Generate the CSR
# echo "Generating Certificate Signing Request (CSR)..."
# openssl req -new -key $PRIVATE_KEY -out $CSR_FILE -subj "/C=$COUNTRY/ST=$STATE/L=$LOCALITY/O=$ORGANIZATION/CN=$COMMON_NAME"

# echo "CSR has been created: $CSR_FILE"
