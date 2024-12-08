# Сервер для дисциплины "Разработка мобильных компонент анализа безопасности программного обеспечения"
Сервер стартует на 3000 порту<br>
## Конфиг

Создать файл **config.yaml** в корневой папке проекта<br>
В файле прописать переменные:<br>

---

**env** - dev или prod<br>
**secretKey** - Ключ для jwt токена<br>
**exp** - Время жизни jwt токена<br>
**host** - Адрес сервера (при запуске через докер указать 0.0.0.0, при запуске локально - localhost)<br>
**database** - Строка подключения к бд (при запуске через докер: mongodb://mongo-database:27017, при запуске локально mongodb://localhost:27017)<br>

---
Пример yaml файла для запуска через докер:
```yaml
env: "dev"
secretKey: "hello"
exp: 100
host: 0.0.0.0
database: "mongodb://mongo-database:27017"
```

Для запуска через докер в корневой папке проектаы прописать:
```bash
sudo docker-compose up -d
```
Для остановки сервера:
```bash
sudo docker-compose down
```

## Отправка сообщений<br>
Т.к. сваггер не поддерживает документацию для веб сокетов, поэтому пишу тут
для отправки сообщения нужно сначала обновить соединение
<br>
<br>
После этого отправить сообщение в виде жсона
```js
{
    chatroom_id: <айдишник чата>,
    sender_guid: <айдишник отправителя>,
    sender_name: <имя отправителя>,
    content: <текст сообщения>
}
```

Пример сообщения:
```js
{
    "chatroom_id": "7907d1cd-d36c-42fa-8c47-9f9fcba39bd2",
    "sender_guid": "1303c7e3-bf9f-4f78-8f94-1108a8fc4b81",
    "sender_name": "Васёк",
    "content": "Дарова"
}
```