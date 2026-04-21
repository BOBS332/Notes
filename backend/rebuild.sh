#!/bin/bash

# Скрипт для пересборки Docker образа

set -e  # Выход при ошибке

echo "🔄 Начинаем пересборку Docker образа..."

# Удаляем контейнеры и тома
echo "🗑️  Удаляем старые контейнеры и тома..."
docker compose down -v

# Пересобираем образ без кэша
echo "🔨 Собираем новый образ..."
docker compose build --no-cache

echo "✅ Пересборка завершена!"
echo ""
echo "Для запуска приложения выполни:"
echo "  docker compose run --rm app"
echo ""
echo "Или для фонового запуска:"
echo "  docker compose up -d"
