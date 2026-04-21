#!/bin/bash
# Быстрый запуск REST API сервера

set -e

cd "$(dirname "$0")"

echo "🔧 Компилирую приложение..."
go build 2>/dev/null || go build

echo ""
echo "════════════════════════════════════════════"
echo "🚀 REST API сервер запускается!"
echo "════════════════════════════════════════════"
echo ""
echo "📍 API URL: http://localhost:8080/api/notes"
echo ""
echo "💡 Нажмите Ctrl + C чтобы остановить сервер"
echo ""
echo "════════════════════════════════════════════"
echo ""

# Передаём "2" для режима API (без фона, в текущем терминале)
./learninggo << 'EOF'
2
EOF
