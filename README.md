# Quote API Service

RESTful API сервис для управления цитатами, написанный на Go.

## Архитектура
Проект построен на слоистой архитектуре с использованием handler -> service -> repository слоев

## Требования

- Go 1.24+
- SQLite3

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/xddprog/brandscout_test.git
cd brandscout_test
```

2. Установите зависимости:
```bash
go mod download
```

3. Запустите сервер:
```bash
go run cmd/main.go
```

Сервер запустится на `http://localhost:8080`

## API Endpoints

### 1. Получить случайную цитату
```http
GET /quotes/random
```

Успешный ответ (200 OK):
```json
{
    "id": 1,
    "quote": "To be or not to be",
    "author": "William Shakespeare"
}
```

### 2. Получить цитату по ID
```http
GET /quotes/{id}
```

Успешный ответ (200 OK):
```json
{
    "id": 1,
    "quote": "To be or not to be",
    "author": "William Shakespeare"
}
```


Ошибка (404 Not Found):
```json
{
    "error": "quote not found"
}
```

### 3. Получить цитаты по автору
```http
GET /quotes?author={author}
```

Успешный ответ (200 OK):
```json
[
    {
        "id": 1,
        "quote": "To be or not to be",
        "author": "William Shakespeare"
    },
    {
        "id": 2,
        "quote": "All the world's a stage",
        "author": "William Shakespeare"
    }
]
```


### 4. Создать новую цитату
```http
POST /quotes
Content-Type: application/json

{
    "quote": "New quote text",
    "author": "Author name"
}
```

Успешный ответ (200 OK):
```json
{
    "id": 3,
    "quote": "New quote text",
    "author": "Author name"
}
```

Ошибка (400 Bad Request):
```json
{
    "error": "invalid request body"
}
```

### 5. Удалить цитату
```http
DELETE /quotes/{id}
```

Успешный ответ (200 OK):
```json
{
    "message": "Quote deleted successfully"
}
```

Ошибка (404 Not Found):
```json
{
    "error": "quote not found"
}
```

### 6. Получить все цитаты
```http
GET /quotes
```

Успешный ответ (200 OK):
```json
[
    {
        "id": 1,
        "quote": "To be or not to be",
        "author": "William Shakespeare"
    },
    {
        "id": 2,
        "quote": "All the world's a stage",
        "author": "William Shakespeare"
    }
]
```

## Обработка ошибок

API использует следующие HTTP статус коды:
- 200: Успешный запрос
- 400: Неверный запрос (некорректные параметры)
- 404: Ресурс не найден
- 500: Внутренняя ошибка сервера

Все ошибки возвращаются в формате:
```json
{
    "error": "описание ошибки"
}
```

## Структура кода

### Core (`internal/core/`)
- `repositories/` - интерфейсы репозиториев для работы с данными
- `services/` - бизнес-логика приложения
  - Реализация основных операций с цитатами
  - Валидация данных
  - Обработка бизнес-правил

### Handlers (`internal/handlers/`)
- Обработка HTTP запросов
- Преобразование HTTP запросов в вызовы сервисов

### Infrastructure (`internal/infrastructure/`)
- `database/`
  - `adapters/` - адаптеры для работы с различными БД
    - `sqlite.go` - реализация для SQLite
  - `models/` - модели данных для работы с БД
- `errors/` - API ошибки и методы для работы с ними
- `types/` - интерфейсы, различные типы 

### Utils (`internal/utils/`)
- Вспомогательные функции для работы с HTTP