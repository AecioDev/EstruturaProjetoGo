#!/bin/bash

# Diretório raiz do projeto
PROJECT_DIR="meu-projeto"

# Nome do módulo antigo e o novo
OLD_MODULE="clinica_server"
NEW_MODULE="meu-projeto"

# Encontrar e substituir o nome do módulo em todos os arquivos .go
find "$PROJECT_DIR" -type f -name "*.go" -exec sed -i "s|$OLD_MODULE|$NEW_MODULE|g" {} +

echo "Nome do módulo alterado com sucesso em todos os arquivos .go!"
