#!/bin/bash
# Быстрый запуск интерактивного CLI режима

set -e

cd "$(dirname "$0")"

echo "🔧 Компилирую приложение..."
go build 2>/dev/null || go build

echo "🎮 Запускаю интерактивный CLI режим..."
echo ""

# Передаём "1" для режима CLI
./learninggo << 'EOF'
1
EOF
