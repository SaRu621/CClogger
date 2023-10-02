# Computer club clients logger
Тестовое задание для Yadro.com  

В logger содержатся функции, необходимые для работы приложения.  

Папки **input**, **output** и **expect** (входные данные, выходные данные и ожидаемые соответственно) нужны для тестирования функции Logger.  

## Запуск приложения: 
```console
git clone https://github.com/SaRu621/CClogger.git
cd CClogger
cd YadroTest
docker build -t go-app .
docker run go-app
```
или

```console
git clone https://github.com/SaRu621/CClogger.git
cd CClogger
cd YadroTest
go run main.go
```

## Запуск тестов:

```console
cd logger
go test -cover -v
```
