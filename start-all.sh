#!/bin/bash
# Запуск обоих режимов - API сервер + интерактивный CLI

set -e

cd "$(dirname "$0")"

echo "🔧 Компилирую приложение..."
go build 2>/dev/null || go build

echo "🚀 REST API будет запущен на http://localhost:8080"
echo "💻 Интерактивный CLI готов к вводу команд"
echo ""

# Передаём "3" для режима API + CLI
./learninggo << 'EOF'
3
EOF
