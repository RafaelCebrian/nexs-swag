#!/bin/bash

# Script para testar a separação de APIs públicas e privadas com Swagger 2.0

set -e

echo "Gerando documentação Swagger 2.0 com separação público/privado..."
echo ""

# Gerar documentação
nexs-swag init --dir ./ --output ./docs --openapi-version 2.0

echo ""
echo "Verificando arquivos gerados..."
echo ""

# Verificar se os arquivos foram criados
if [ -f "docs/swagger_public.json" ]; then
    echo "[OK] docs/swagger_public.json criado"
else
    echo "[ERRO] docs/swagger_public.json NÃO encontrado"
    exit 1
fi

if [ -f "docs/swagger_private.json" ]; then
    echo "[OK] docs/swagger_private.json criado"
else
    echo "[ERRO] docs/swagger_private.json NÃO encontrado"
    exit 1
fi

echo ""
echo "Contando endpoints em cada spec..."
echo ""

# Contar paths no público
PUBLIC_PATHS=$(cat docs/swagger_public.json | jq '.paths | keys | length')
echo "Endpoints no PÚBLICO: $PUBLIC_PATHS"

# Contar paths no privado
PRIVATE_PATHS=$(cat docs/swagger_private.json | jq '.paths | keys | length')
echo "Endpoints no PRIVADO: $PRIVATE_PATHS"

echo ""
echo "Detalhamento dos endpoints PÚBLICOS:"
echo ""
cat docs/swagger_public.json | jq -r '.paths | keys[]' | while read path; do
    methods=$(cat docs/swagger_public.json | jq -r ".paths[\"$path\"] | keys[]" | grep -v "^\$" | tr '\n' ',' | sed 's/,$//')
    echo "  $path [$methods]"
done

echo ""
echo "Detalhamento dos endpoints PRIVADOS:"
echo ""
cat docs/swagger_private.json | jq -r '.paths | keys[]' | while read path; do
    methods=$(cat docs/swagger_private.json | jq -r ".paths[\"$path\"] | keys[]" | grep -v "^\$" | tr '\n' ',' | sed 's/,$//')
    echo "  $path [$methods]"
done

echo ""
echo "Validação:"
echo ""

# Validar que privado tem mais ou igual endpoints que público
if [ "$PRIVATE_PATHS" -ge "$PUBLIC_PATHS" ]; then
    echo "[OK] CORRETO: Privado ($PRIVATE_PATHS) >= Público ($PUBLIC_PATHS)"
else
    echo "[ERRO] ERRO: Privado ($PRIVATE_PATHS) < Público ($PUBLIC_PATHS)"
    exit 1
fi

# Verificar se endpoints públicos estão no privado
echo ""
echo "Verificando se todos os endpoints públicos estão no privado..."
echo ""

PUBLIC_ENDPOINTS=$(cat docs/swagger_public.json | jq -r '.paths | keys[]')
PRIVATE_ENDPOINTS=$(cat docs/swagger_private.json | jq -r '.paths | keys[]')

ALL_PUBLIC_IN_PRIVATE=true
for endpoint in $PUBLIC_ENDPOINTS; do
    if echo "$PRIVATE_ENDPOINTS" | grep -q "^$endpoint$"; then
        echo "  [OK] $endpoint encontrado no privado"
    else
        echo "  [ERRO] $endpoint NÃO encontrado no privado"
        ALL_PUBLIC_IN_PRIVATE=false
    fi
done

echo ""
if [ "$ALL_PUBLIC_IN_PRIVATE" = true ]; then
    echo "[OK] TESTE PASSOU: Todos os endpoints públicos estão no privado!"
else
    echo "[ERRO] TESTE FALHOU: Nem todos os endpoints públicos estão no privado!"
    exit 1
fi

echo ""
echo "Resumo Final (Swagger 2.0):"
echo "  Público:  $PUBLIC_PATHS endpoints"
echo "  Privado:  $PRIVATE_PATHS endpoints"
echo "  Status:   CORRETO"
echo ""
echo "Teste de visibilidade Swagger 2.0 concluído com sucesso!"
