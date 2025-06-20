#!/bin/bash

echo "📦 Baixando o binário ImageGO..."
curl -LJO https://github.com/rohitaryal/imageGO/releases/latest/download/imagego-linux

echo "🔄 Renomeando e dando permissão de execução..."
mv imagego-linux imagego
chmod +x imagego

echo "✅ Binário pronto para uso."
