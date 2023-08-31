## Минималистичное приложение-интерфейс для работы с Redis'ом с проксированием трафика.

## Краткое описание
В приложении есть 3 эндпоинта, формат ответа — json:
1. POST /set_key { "[key]": "[val]" } — создать или перезаписать пару ключ-значение/значение по заданному ключу
2. GET /get_key?key=[key] — вернуть значение по ключу (если ключа нет, то 404)
3. DELETE /det_key { "key": "[key]" } — удалить пару ключ-значение по ключу.

## Структура проекта

```
.
├── .env
├── .gitignore
├── app
│   ├── Dockerfile
│   └── backend
├── docker-compose.yaml
├── nginx
│   ├── Dockerfile
│   └── nginx.conf
├── readme.md
└── redis
    └── Dockerfile
```

## Технические особенности

- В качестве базового образа для контейнеров используется debian:12.1-slim

- Все переменные окружения (порты, хосты, ~~пароли~~) задаются в файле .env. Все порты можно менять. В app они прокидываются через .env файл, в redis подхватываются при запуске, в nginx с помощью envsubst. Не стала прописывать EXPOSE, хотя можно было для наглядной информации при вызове ```docker ps```.

- Трафик к приложению проходит через nginx, порт по умолчанию 8089 (NGINX_EXTERNAL_PORT). Поэтому пользователю локально доступен только он.
 
- Добавлена аутентификация при работе с redis.
    
    Для проверки можно зайти в контейнер, то есть выполнить следующие команды:
    ```
    docker exec -it adt-redis bash
    redis-cli
    set start end
    ```
    Будет ошибка: ```(error) NOAUTH Authentication required.```.
    
    Для решения можно ввести команду ```auth ${REDIS_PASS}```.

- Redis работает только со строками, соответственно команды SET, GET и GETDEL.

## Запуск

```
git clone https://github.com/Dashori/adt.git
cd adt

docker-compose up -d
```

Если все успешно, то:
```
[+] Running 3/3
 ⠿ Container adt-redis    Started                                                                                                                                   1.5s
 ⠿ Container adt-backend  Started                                                                                                                                   1.4s
 ⠿ Container adt-nginx    Started    
```
Более подробную информацию о контейнерах можно получить с помощью команды ```docker ps``` или ```docker logs имя-контейнера```. Логи nginx автоматически появятся в папке logs/nginx.

### Примеры

Запрос на создание пары:
```
curl -X POST localhost:8089/set_key -H "Content-Type: application/json" -d '{"start": "test"}' -i


HTTP/1.1 200 OK
Server: nginx/1.22.1
Date: Mon, 28 Aug 2023 21:09:01 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 75
Connection: keep-alive

{"message":"Success add key-value into redis","key":"start","value":"test"}
```

Запрос на получение значения по ключу:

```
curl -X GET 'localhost:8089/get_key?key=start' -i


HTTP/1.1 200 OK
Server: nginx/1.22.1
Date: Mon, 28 Aug 2023 21:09:10 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 71
Connection: keep-alive

{"message":"Success get value from redis","key":"start","value":"test"}
```

Запрос на удаление пары:
```
curl -X DELETE localhost:8089/del_key -H "Content-Type: application/json" -d '{"key": "start"}' -i


HTTP/1.1 200 OK
Server: nginx/1.22.1
Date: Mon, 28 Aug 2023 21:09:25 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 74
Connection: keep-alive

{"message":"Success delete key-value from redis","key":"start","value":""}   
```

Запрос на удаление пары, которой нет:
```
curl -X DELETE localhost:8089/del_key -H "Content-Type: application/json" -d '{"key": "start2"}' -i


HTTP/1.1 400 Bad Request
Server: nginx/1.22.1
Date: Thu, 31 Aug 2023 13:24:24 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 65
Connection: keep-alive

{"message":"There is no pair with this key to delete","error":""}
```


Запрос на получение несуществующего ключа:
```
curl -X GET 'localhost:8089/get_key?key=start2' -i

HTTP/1.1 404 Not Found
Server: nginx/1.22.1
Date: Mon, 28 Aug 2023 21:09:52 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 45
Connection: keep-alive

{"message":"No value with this key in redis"}
```

Запрос на несуществующий эндпоинт, до приложения не доходит:
```
curl -X POST localhost:8089/update_key -H "Content-Type: application/json" -d '{"start": "test"}' -i


HTTP/1.1 403 Forbidden
Server: nginx/1.22.1
Date: Mon, 28 Aug 2023 21:10:19 GMT
Content-Type: application/octet-stream
Content-Length: 9
Connection: keep-alive

Error request
```
