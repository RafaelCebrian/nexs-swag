# Example 26: X-Visibility (OpenAPI 3.1)

This example demonstrates the `@x-visibility` annotation feature with **OpenAPI 3.1**, allowing you to generate separate documentation files for public and private APIs.

> **Note**: This example focuses on OpenAPI 3.1. For Swagger 2.0 version, see [Example 27](../27-x-visibility-v2/).

## 📋 Overview

The `@x-visibility` annotation enables you to:
- Separate public and private endpoints into different OpenAPI files
- Automatically filter schemas based on usage
- Maintain a single codebase with dual documentation output
- Works seamlessly with OpenAPI 3.1's `components.schemas` structure

## 🎯 Endpoints in This Example

### Public Endpoints (appear in both specs):
- `GET /api/v1/products` - List products (no annotation → defaults to public)
- `GET /api/v1/products/{id}` - Get product (`@x-visibility public`)
- `GET /api/v1/health` - Health check (no annotation → defaults to public)

### Private Endpoints (appear ONLY in private spec):
- `POST /api/v1/products` - Create product (`@x-visibility private`)
- `PUT /api/v1/products/{id}` - Update product (`@x-visibility private`)
- `DELETE /api/v1/products/{id}` - Delete product (`@x-visibility private`)
- `GET /api/v1/admin/users` - List users (`@x-visibility private`)
- `DELETE /api/v1/admin/users/{id}` - Delete user (`@x-visibility private`)

### Visibility Logic

**Important**: Private spec includes ALL operations (public + private). Public spec includes only public operations.

## 🚀 Running This Example

### Option 1: Run test script
```bash
./run.sh
```

### Option 2: Manual generation
```bash
nexs-swag init --dir ./ --output ./docs

# Verify public endpoints (should have 3)
jq '.paths | keys | length' docs/openapi_public.json

# Verify private endpoints (should have 5 - includes public)
jq '.paths | keys | length' docs/openapi_private.json
```

## 📊 Expected Results

### Generated Structure:

```
docs/
├── openapi_public.json     # Public API (3 paths)
├── openapi_public.yaml
├── docs_public.go
├── openapi_private.json    # Complete API (5 paths - includes public)
├── openapi_private.yaml
└── docs_private.go
```
```json
{
  "paths": {
    "/api/v1/products": {
      "get": { "summary": "Lista produtos" }
    },
    "/api/v1/products/{id}": {
      "get": { "summary": "Buscar produto" }
    },
    "/api/v1/health": {
      "get": { "summary": "Health check público" }
    }
  }
}
```

### openapi_private.json deve conter:
```json
{
  "paths": {
    "/api/v1/products": {
      "get": { "summary": "Lista produtos" },
      "post": { "summary": "Criar produto" }
    },
    "/api/v1/products/{id}": {
      "get": { "summary": "Buscar produto" },
      "put": { "summary": "Atualizar produto" },
      "delete": { "summary": "Deletar produto" }
    },
    "/api/v1/health": {
      "get": { "summary": "Health check público" }
    },
    "/api/v1/admin/users": {
      "get": { "summary": "Listar usuários" }
    },
    "/api/v1/admin/users/{id}": {
      "delete": { "summary": "Deletar usuário" }
    }
  }
}
```

## ✅ Validações do Teste

O script `test-visibility.sh` valida:

1. ✅ Arquivos `openapi_public.json` e `openapi_private.json` foram criados
2. ✅ Número de endpoints no privado >= número de endpoints no público
3. ✅ Todos os endpoints públicos estão presentes no privado
4. ✅ Endpoints marcados com `@x-visibility private` aparecem APENAS no privado

## 🎯 Casos de Uso

### 1. API SaaS com plano Free e Premium
- **Público**: Endpoints do plano gratuito
- **Privado**: Todos os endpoints (free + premium)

### 2. API com área administrativa
- **Público**: Endpoints para usuários finais
- **Privado**: Endpoints públicos + endpoints administrativos

### 3. Documentação para parceiros vs interna
- **Público**: Endpoints para parceiros externos
- **Privado**: Todos os endpoints para uso interno

## 📝 Como Usar no Seu Projeto

### Marcar endpoints como públicos:
```go
// Opção 1: Sem annotation (público por padrão)
// @Router /products [get]
func ListProducts(w http.ResponseWriter, r *http.Request) { }

// Opção 2: Explicitamente público
// @Router /products/{id} [get]
// @x-visibility public
func GetProduct(w http.ResponseWriter, r *http.Request) { }
```

### Marcar endpoints como privados:
```go
// @Router /products [post]
// @Security Bearer
// @x-visibility private
func CreateProduct(w http.ResponseWriter, r *http.Request) { }
```

### Gerar documentações:
```bash
nexs-swag init --output ./docs
```

### Servir múltiplas documentações:
```go
import (
    _ "seu-projeto/docs"
    "github.com/gofiber/swagger"
)

// Documentação pública
app.Get("/swagger-public/*", swagger.New(swagger.Config{
    URL: "/docs/openapi_public.json",
}))

// Documentação privada (com autenticação)
app.Get("/swagger-private/*", authMiddleware, swagger.New(swagger.Config{
    URL: "/docs/openapi_private.json",
}))
```

## 🐛 Troubleshooting

### Problema: Endpoints privados aparecendo no público
- Verifique se você adicionou `@x-visibility private` no endpoint
- Certifique-se que não há espaços extras na annotation

### Problema: Endpoints públicos não aparecem no privado
- Verifique se a lógica de geração está correta
- Execute `./test-visibility.sh` para diagnosticar

### Problema: JSON não sendo gerado
- Verifique se o comando `nexs-swag init` rodou sem erros
- Confirme que as annotations `@x-visibility` estão corretas

## 📚 Referências

- [Exemplo 07 - Tags Filter](../07-tags-filter/README.md) - Filtrar por tags
- [Exemplo 16 - Instance Name](../16-instance-name/README.md) - Múltiplas instâncias
- [OpenAPI 3.1 Extensions](https://spec.openapis.org/oas/v3.1.0#specification-extensions)
