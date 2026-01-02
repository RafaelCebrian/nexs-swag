#!/bin/bash

# Example 26: X-Visibility (OpenAPI 3.1)
# Tests visibility separation for OpenAPI 3.1 specification

set -e

echo "=========================================="
echo "Example 26: X-Visibility (OpenAPI 3.1)"
echo "=========================================="
echo ""

# Run OpenAPI 3.1 visibility tests
echo "Running OpenAPI 3.1 visibility tests..."
echo ""

chmod +x ./test-visibility.sh
./test-visibility.sh

echo ""
echo "=========================================="
echo "All tests completed successfully!"
echo "=========================================="

echo ""
echo "Note: For Swagger 2.0 tests, see Example 27"
