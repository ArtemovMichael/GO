# Тема проекта
Необходимо написать http сервер по работе с аккаунтами и балансами. Для этого я использовал фрейворк fiber. Реализовал 5 методов (получение аккаунта, изменение имени аккаунта, изменение баланса аккаунта, создание и удаление аккаунта). Далее реализовал CLI (command-line interface), он же клиент. В нем для реализации интерфейса использовал стандартную библиотеку flag. 

1. В терминале прописать ```docker run -p 5432:5432 --name some-postgres -e POSTGRES_PASSWORD=0000 -d postgres``` для запуска докера и постгреса.
2. В терминале надо прописать ```psql -h 0.0.0.0 -p 5432 -U postgres``` и ввести пароль 0000.
3. Далее создаем базу данных ```CREATE TABLE accounts(name varchar(256) PRIMARY KEY, amount int not null default 0);```.
4. С помощью команды ```SELECT * FROM "accounts";``` можно вывести нашу базу данных.

Чтобы запустить сервер необходимо перейти в директорию ```/GO/Bank_account_Project_fiber_postgres/cmd/server``` и прописать ```go run main.go```.

Чтобы запустить клиента необходимо перейти в директорию ```/GO/Bank_account_Project_fiber_postgres/cmd/client``` и можно выполнить следующие 5 запросов:
1. Создание аккаунта ```go run main.go -cmd create -name "name" -amount "amount"```.
2. Получение баланса по имени ```go run main.go -cmd get -name "name"```.
3. Удаление аккаунта по имени ```go run main.go -cmd delete -name "name"```.
4. Изменение имени аккаунта ```go run main.go -cmd change_name -name "name" -newname "newname"```.
5. Изменение баланса по имени ```go run main.go -cmd change_amount -name "name" -amount "amount"```.


Чтобы закрыть докер необходимо написать:
1. ```docker container stop "CONTAINER ID"```.
2. ```docker container rm "CONTAINER ID"```.
