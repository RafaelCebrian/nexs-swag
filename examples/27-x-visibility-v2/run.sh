#!/bin/bash

# Example 27: X-Visibility with Swagger 2.0
# Tests visibility separation for Swagger 2.0 specification

set -e

echo "=========================================="
echo "Example 27: X-Visibility (Swagger 2.0)"
echo "=========================================="
echo ""

# Run Swagger 2.0 visibility tests
echo "Running Swagger 2.0 visibility tests..."
echo ""

chmod +x ./test-visibility.sh
./test-visibility.sh

echo ""
echo "=========================================="
echo "All tests completed successfully!"
echo "=========================================="
