# Сервер для дисциплины "Разработка мобильных компонент анализа безопасности программного обеспечения"

## Конфиг

Конфиг нужно прописывать в yaml файле<br>
Пример yaml файла для запуска через докер:
```yaml
env: "dev"
secretKey: "hello"
exp: 100
port: 3000
host: 0.0.0.0
database: "mongodb://mongo-database:27017"
```

---

**env** - Четверг<br>
**secretKey** - Ключ для jwt токена<br>
**exp** - Время жизни jwt токена<br>
**port** - Порт сервера<br>
**host** - Адрес сервера (при запуске через докер указать 0.0.0.0, при запуске локально - localhost)<br>
**database** - Строка подключения к бд (при запуске через докер: mongodb://mongo-database:27017, при запуске локально mongodb://localhost:27017)<br>

---

Для запуска через докер в корне прописать:
```bash
sudo docker-compose up -d
```
Для остановки сервера:
```bash
sudo docker-compose down
```