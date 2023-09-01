# Computer club clinets logger
Тестовое задание для Yadro.com  

В logger содержатся функции, необходимые для работы приложения.  

Папки **input**, **output** и **expect** (входные данные, выходные данные и ожидаемые соответственно) нужны для тестирования функции Logger.  

## Запуск приложения: 
```console
git clone https://github.com/SaRu621/CClogger.git
cd CClogger
cd YadroTest
docker build -t go-app .
```
## Запуск тестов:

```console
cd logger
go test -cover -v
```

Для запуска в консоли (linux) достаточно перейти в директорию с файлом main.go с помощью команды cd и исполнить команду go run main.go.  

Для запуска тестов переходим в папку logger и запускаем в консоли команду go test -cover -v. Выводом будет уведомление о прохождении i-го теста и процент покрытия кода тестами.
