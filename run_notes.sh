#!/bin/bash

# Скрипт для запуска проекта с заметками

PROJECT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Проект находится в: $PROJECT_DIR"
cd "$PROJECT_DIR" || exit 1

# Проверяем наличие Go
if ! command -v go &> /dev/null; then
    echo "Ошибка: Go не установлен. Пожалуйста, установите Go."
    exit 1
fi

# Проверяем наличие PostgreSQL
if ! command -v psql &> /dev/null; then
    echo "Ошибка: PostgreSQL CLI не установлен. Пожалуйста, убедитесь, что PostgreSQL запущен."
    exit 1
fi

echo "Сборка приложения..."
go build -o notes-app main.go

if [ ! -f notes-app ]; then
    echo "Ошибка: Не удалось собрать приложение."
    exit 1
fi

echo "Запуск приложения..."
./notes-app