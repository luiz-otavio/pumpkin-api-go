#!/bin/bash

# Ensure the submodule is initialized to most recent commit
git submodule update --remote --merge

# Generate the Go bindings
rm -rf pkg
mkdir -p pkg

# We generate the Go bindings with stubs for the WIT exports since
# updating the WITs will probably have new exports that need stubs.
# The bridge code will need to be updated accordingly in pkg/exports
wit-bindgen go --pkg-name github.com/Pumpkin-MC/pumpkin-api-go/pkg \
    --generate-stubs \
    --out-dir pkg -w plugin ./wit/repo/crates/pumpkin-plugin-wit/v0.2

echo "Bindings and components generated successfully in pkg/ and wit/repo/crates/pumpkin-plugin-wit/v0.2"
