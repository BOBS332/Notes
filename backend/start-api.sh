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

# Устанавливаем переменные окружения для подключения к БД
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=newpassword
export DB_NAME=test
export APP_MODE=2

# Передаём "2" для режима API (без фона, в текущем терминале)
./learninggo << 'EOF'
2
EOF
