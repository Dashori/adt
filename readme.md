
curl -X POST localhost:8089/set_key -H "Content-Type: application/json" -d '{"start": "test"}' -i

HTTP/1.1 200 OK
Server: nginx/1.22.1
Date: Mon, 28 Aug 2023 21:09:01 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 75
Connection: keep-alive

{"message":"Success add key-value into redis","key":"start","value":"test"}


curl -X GET 'localhost:8089/get_key?key=start' -i

HTTP/1.1 200 OK
Server: nginx/1.22.1
Date: Mon, 28 Aug 2023 21:09:10 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 71
Connection: keep-alive

{"message":"Success get value from redis","key":"start","value":"test"}


curl -X DELETE localhost:8089/del_key -H "Content-Type: application/json" -d '{"key": "start"}' -i

HTTP/1.1 200 OK
Server: nginx/1.22.1
Date: Mon, 28 Aug 2023 21:09:25 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 74
Connection: keep-alive

{"message":"Success delete key-value from redis","key":"start","value":""}   


curl -X GET 'localhost:8089/get_key?key=start2' -i
HTTP/1.1 404 Not Found
Server: nginx/1.22.1
Date: Mon, 28 Aug 2023 21:09:52 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 45
Connection: keep-alive

{"message":"No value with this key in redis"}



curl -X POST localhost:8089/update_key -H "Content-Type: application/json" -d '{"start": "test"}' -i

HTTP/1.1 403 Forbidden
Server: nginx/1.22.1
Date: Mon, 28 Aug 2023 21:10:19 GMT
Content-Type: application/octet-stream
Content-Length: 9
Connection: keep-alive

Error uri