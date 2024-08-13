# Тема проекта
Необходимо написать http сервер по работе с аккаунтами и балансами. Для этого можно использовать, фрейворк fiber. Должны быть реализованы 5 методов (получение аккаунта, изменение имени аккаунта, изменение баланса аккаунта, создание и удаление аккаунта). Далее должен быть реализован CLI (command-line interface), он же клиент. В нем для реализации интерфейса использовать можно стандартную библиотеку flag. 

1. В терминале надо прописать ```psql -h 0.0.0.0 -p 5432 -U postgres``` и ввести пароль 0000.
2. Дальше создаем базу данных ```CREATE TABLE accounts(name varchar(256) PRIMARY KEY, amount int not null default 0);```.
3. С помощью команды ```SELECT * FROM "accounts";``` можно вывести нашу базу данных.

Клиент может выполнять такие запросы:
1. Создание аккаунта ```go run main.go -cmd create -name "name" -amount "amount"```.
2. Получение баланса по имени ```go run main.go -cmd get -name "name"```.
3. УДаление аккаунта по имени ```go run main.go -cmd delete -name "name"```.
4. Изменение имени аккаунта ```go run main.go -cmd change_name -name "name" -newname "newname"```.
5. Изменение баланса по имени ```go run main.go -cmd change_amount -name "name" -amount "amount"```.
