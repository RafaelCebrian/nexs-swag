# Example 27: X-Visibility (Swagger 2.0)

This example demonstrates the `@x-visibility` annotation feature with **Swagger 2.0**, allowing you to generate separate documentation files for public and private APIs.

> **Note**: This example focuses on Swagger 2.0. For OpenAPI 3.1 version, see [Example 26](../26-x-visibility/).

## Overview

The `@x-visibility` annotation enables you to:
- Separate public and private endpoints into different Swagger files
- Automatically filter schemas (definitions) based on usage
- Maintain a single codebase with dual documentation output
- Works seamlessly with Swagger 2.0's `definitions` structure

## How It Works

### Annotation Syntax

Add `@x-visibility` to your operation comments:

```go
// GetUser godoc
// @Summary      Get user (public)
// @Description  Get user details for public consumption
// @Tags         users
// @Success      200  {object}  UserPublic
// @Router       /users/{id} [get]
// @x-visibility public
func GetUser(c *gin.Context) {
    // handler implementation
}
```

### Visibility Logic

- `@x-visibility public` → Endpoint **only** in `swagger_public.json`
- `@x-visibility private` → Endpoint **only** in `swagger_private.json`
- No annotation → Endpoint in **both** files (shared endpoint)

**Important**: Private spec includes ALL operations (public + private). Public spec includes only public operations.

## Generated Files

When using `@x-visibility` with Swagger 2.0, nexs-swag generates:

```
docs/
├── swagger_public.json    # Public API specification
├── swagger_private.json   # Private API specification (includes public)
├── swagger_public.yaml    # Public API (YAML)
├── swagger_private.yaml   # Private API (YAML)
├── docs_public.go         # Public API Go code
└── docs_private.go        # Private API Go code
```

## Running This Example

```bash
# Run the test script
./run.sh

# Or manually:
nexs-swag init --output ./docs --ov 2.0

# Verify public endpoints (should have 3)
jq '.paths | keys | length' docs/swagger_public.json

# Verify private endpoints (should have 5 - includes public)
jq '.paths | keys | length' docs/swagger_private.json
```

## Test Validation

The `test-visibility.sh` script validates:
1. ✅ Swagger files are generated
2. ✅ Public spec has exactly 3 endpoints
3. ✅ Private spec has exactly 5 endpoints
4. ✅ All public endpoints exist in private spec (subset validation)

## Endpoints in This Example

| Endpoint | Method | Visibility | Public Spec | Private Spec |
|----------|--------|------------|-------------|--------------|
| `/users/{id}` | GET | public | ✅ | ✅ |
| `/users` | POST | (none) | ✅ | ✅ |
| `/users/profile` | GET | (none) | ✅ | ✅ |
| `/admin/users/{id}` | DELETE | private | ❌ | ✅ |
| `/admin/users/{id}/role` | PUT | private | ❌ | ✅ |

**Result**: 
- Public spec: 3 endpoints
- Private spec: 5 endpoints (all of them)

## Swagger 2.0 Specifics

This example uses Swagger 2.0 features:
- `definitions` instead of `components.schemas`
- `swagger: "2.0"` version field
- Different file naming: `swagger_*.json` vs `openapi_*.json`

Schema filtering works identically to OpenAPI 3.x:
```bash
# Public definitions (Swagger 2.0)
jq '.definitions | keys' docs/swagger_public.json

# Private definitions (Swagger 2.0)
jq '.definitions | keys' docs/swagger_private.json
```

## Comparison with OpenAPI 3.x

| Feature | Swagger 2.0 (Example 27) | OpenAPI 3.1 (Example 26) |
|---------|--------------------------|--------------------------|
| File Names | `swagger_*.json` | `openapi_*.json` |
| Schemas | `definitions` | `components.schemas` |
| Version Field | `swagger: "2.0"` | `openapi: "3.1.0"` |
| Visibility Logic | ✅ Same | ✅ Same |
| Flag | `--ov 2.0` | `--ov 3.1` (default) |

Both examples use the **same** visibility filtering logic.

## Notes

- Schemas (definitions) are automatically filtered based on endpoint usage
- Private spec includes ALL endpoints (public + private) for complete internal documentation
- Public spec only includes public endpoints for external API consumers
- The `@x-visibility` extension is preserved during OpenAPI version conversion
- All other Swagger 2.0 features work normally within each spec
