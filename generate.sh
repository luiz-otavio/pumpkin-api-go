#!/bin/bash

# Ensure the submodule is initialized (based on actual repository state)
git submodule update --init

# Generate the Go bindings
rm -rf pkg
mkdir -p pkg

# Note: We don't use --generate-stubs here because we have hand-written
# bridge code in internal/exports that maps the WIT exports to our
# higher-level Go API in the api/ package.
wit-bindgen go --pkg-name github.com/Pumpkin-MC/pumpkin-api-go/pkg \
    --out-dir pkg -w plugin ./wit/repo/crates/pumpkin-plugin-wit/v0.2

echo "Bindings and components generated successfully in pkg/ and wit/repo/crates/pumpkin-plugin-wit/v0.2"
