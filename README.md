# ✈️ Mini-Avia API

Мини-проект на Go и PostgreSQL для работы с авиабилетами: поиск офферов, бронирование и выписка билетов.  
Используется авторизация через **JWT** и драйвер **pgx/pgxpool** для взаимодействия с базой данных.

---

## 🚀 Поднятие проекта

1. **Создай файл окружения**
   ```bash
   cp .env.example .env
2. **Поднятие проекта**
   ```bash
   make up
   
   или
   
   docker-compose up

---
##  🧩 Роуты API

#### 🌍 Публичные (без авторизации)

1. **Получить список актуальных офферов (на будущие даты)**
   ```bash
   curl --location 'http://localhost:8080/all-offers'

2. **Поиск офферов по направлению и дате**
   ```bash
   curl --location 'http://localhost:8080/offers?from=TAS&to=SVO&date=2025-10-30'
   
---

#### 🔐 Авторизация

1. **Регистрация**
   ```bash
   curl --location 'http://localhost:8080/auth/register' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "name": "Fedor",
    "surname": "Lidov",
    "birth_date": "2000-01-01",
    "passport_doc": "AB1234567",
    "email": "ivan@qweqwe.com",
    "password": "12345678"
    }'
   
2. **Логин**
   ```bash
   curl --location 'http://localhost:8080/auth/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "email": "ivan@qweqwe.com",
    "password": "12345678"
    }'

---
#### 🌍 Приватные (с авторизацией)

1. **Бронирование оффера**
   ```bash
   curl --location 'http://localhost:8080/bookings' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer <JWT_TOKEN>' \
    --data '{
    "offer_id": 1,
    "passenger_name": "Aziz",
    "passenger_surname": "Azizov",
    "passport_doc": "AD1311111",
    "passenger_birth": "2004-05-26"
    }'

2. **Выписка билета по брони (только владелец брони)**
   ```bash
   curl --location --request POST 'http://localhost:8080/tickets/1/issue' \
    --header 'Accept: application/json' \
    --header 'Authorization: Bearer <JWT_TOKEN>'

---

##  🧪 Тестирование

1. **Запуск тестов**
   ```bash
    go test ./internal/bookings -v
   
**Есть 2 теста, 1 должен упасть, 2 пройти успешно**
