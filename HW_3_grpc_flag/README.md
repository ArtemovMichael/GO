# Тема проекта
Необходимо написать сервер по работе с аккаунтами и балансами через grpc. Реализовал 5 методов (получение аккаунта, изменение имени аккаунта, изменение баланса аккаунта, создание и удаление аккаунта). Далее реализовал CLI (command-line interface), он же клиент. В нем для реализации интерфейса использовал стандартную библиотеку flag. 


Чтобы запустить сервер необходимо перейти в директорию ```/GO/HW_3_grpc_flag/cmd/server``` и прописать ```go run main.go```.

Чтобы запустить клиента необходимо перейти в директорию ```/GO/HW_3_grpc_flag/cmd/client``` и можно выполнить следующие 5 запросов:
1. Создание аккаунта ```go run main.go -cmd create -name "name" -amount "amount"```.
2. Получение баланса по имени ```go run main.go -cmd get -name "name"```.
3. Удаление аккаунта по имени ```go run main.go -cmd delete -name "name"```.
4. Изменение имени аккаунта ```go run main.go -cmd change_name -name "name" -newname "newname"```.
5. Изменение баланса по имени ```go run main.go -cmd change_amount -name "name" -amount "amount"```.
