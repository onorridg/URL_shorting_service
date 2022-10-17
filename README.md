# URL shorting service

#### Создайте файл `.env` в корневой директории проекта:
```
PG_HOST=localhost
PG_PORT=5432
PG_USER=postgres
PG_PASSWORD=postgres
PG_DB_NAME=urldb
PG_DB_TABLE_NAME=urls
HOSTNAME=localhost:8080
```

#### Запуск сервиса:
```bash
make run
```
#### Остановить сервис:
```bash
make stop
```

#### Запросы на которые отвечает сервис:
```bash
POST /api/v1 body:{"url": "example.com"}
=>
201:{"short_url": "http://url-shorting-service.ru/CfjShSb9FO", "real_url": "example.com"}
||
# Если real_url уже есть в базе, то вернется уже созданная ранее ссылка 
200:{"short_url": "http://url-shorting-service.ru/CfjShSb9FO", "real_url": "example.com"}
||
500: Internal Server Error


GET /CfjShSb9FO
=>
301: redirect example.com
||
404: данного URL нет в базе, вернёт 404.html
||
500: Internal Server Error
```