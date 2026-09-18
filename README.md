# QuickDeal

Доска объявлений, где пользователи размещают и просматривают объявления о продаже товаров

![QuickDeal](docs/screenshots/screenshot.png)

## Стек

**Backend:** Go, Gin, GORM<br>
**СУБД:** PostgreSQL<br>
**Frontend:** TypeScript, Angular<br>
**Инфраструктура:** Docker, Docker Compose, nginx


## Функционал

- Регистрация пользователей
- Авторизация по JWT (access и refresh токены)
- Создание объявлений
- Просмотр объявлений


## Быстрый старт

### Docker

#### 1. Настройка переменных окружения

Для Docker — скопируйте `.env.example` в `.env` в корне проекта:
```bash
cp .env.example .env
```
```env
# порт, на котором backend слушает внутри контейнера
BACKEND_PORT=8080
# пользователь PostgreSQL
POSTGRES_USER=user
# пароль PostgreSQL
POSTGRES_PASSWORD=password
# порт PostgreSQL
POSTGRES_PORT=5432
# имя базы данных
POSTGRES_DB=quickdeal
# секрет для подписи JWT access-токена
JWT_ACCESS_SECRET=jwt-access-secret
# время жизни access-токена
ACCESS_TOKEN_TTL=15m
# время жизни refresh-токена
REFRESH_TOKEN_TTL=120h
# порт, на котором доступен frontend (nginx)
FRONTEND_PORT=80
```

#### 2. Запуск

```bash
make up
```
Приложение доступно на `http://localhost`


### Локальный запуск

#### 1. Настройка переменных окружения

Для локального запуска backend — скопируйте `.env.example` в `.env`
внутри `backend/`:
```bash
cp backend/.env.example backend/.env
```
```env
# порт, на котором слушает backend
PORT=8080
# строка подключения к PostgreSQL
DB_DSN=postgres://user:password@localhost:5432/quickdeal?sslmode=disable
# секрет для подписи JWT access-токена
JWT_ACCESS_SECRET=jwt-access-secret
# время жизни access-токена
ACCESS_TOKEN_TTL=15m
# время жизни refresh-токена
REFRESH_TOKEN_TTL=120h
```

#### 2. Запуск

**Backend:**
```bash
make run
```
> **Примечание:** так же можно использовать команду `make run-air`, обеспечивающую автоматический перезапуск при изменении файлов

**Frontend:**
```bash
cd frontend
npm start
```


## Postman

Коллекция запросов ко всем эндпоинтам API — [postman_collection](postman/postman_collection.json)

Импортируйте файл в Postman и укажите переменную коллекции `base_url`

> **Примечание:** `http://localhost` — для запуска через Docker, при локальном 
> запуске backend замените на `http://localhost:8080`

Токены, полученные при логине (`accessToken`, `refreshToken`) сохраняются автоматически после логина и подставляются в остальные запросы


## Команды

| Команда | Что делает |
|---|---|
| `make up` | Поднять все контейнеры |
| `make down` | Остановить контейнеры |
| `make restart` | Перезапустить контейнеры |
| `make clean` | Остановить контейнеры и удалить volume |
| `make logs` | Логи всех контейнеров |

Полный список в [Makefile](Makefile)


## Зависимости

**Backend:**

| Библиотека | Зачем |
|---|---|
| `gin-gonic/gin` | HTTP-фреймворк |
| `gorm.io/gorm` | ORM |
| `gorm.io/driver/postgres` | драйвер PostgreSQL для GORM |
| `golang-jwt/jwt/v5` | генерация и валидация JWT |
| `golang.org/x/crypto` | хеширование паролей (bcrypt) |
| `google/uuid` | генерация UUID |
| `joho/godotenv` | загрузка переменных окружения из `.env` |

**Frontend:** стандартный стек Angular CLI, без сторонних библиотек
сверх дефолтного набора (`@angular/*`, RxJS)


## Лицензия

Проект распространяется под [лицензией MIT](LICENSE)
