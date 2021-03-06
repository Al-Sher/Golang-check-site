# Golang-check-site.
Проверка доступности сайта с помощью Golang.
Внимание! Программа была создана с целью проверки доступности собственных сайтов! Использование её не поназначению может повлечь уголовную ответственность!

# Запуск программы.
Для запуска программы необходимо написать в консоли следующую строку:
```
go run GolangCheckSite.go 
```
Запуск программы с опциями:
```
go run GolangCheckSite.go timer 10 proxy "http://localhost:1080"
```

### Доступные опции.
* `timer` - Таймер проверки сайта;
* `proxy` - Активация прокси сервера в формате scheme://ip:port;
* `timeout` - Таймаут для проверки сайта;
* `file` - Путь до файла с проверяемыми сайтами;
* `useragent` - Кастомный User Agent.

### Результат работы.
В случае доступности сайта, будет выводиться ответ формата
```
{URL}: ok
```
Где {URL} - адрес проверяемого сайта.
В случае ошибок, будут выводиться ошибки следующего формата:
```
error: {TEXTERROR}
```
Если сайт отдал ошибочный http-статус, то ошибка будет иметь следующий формат:
```
Site {URL} returned error code: {RETURNCODE}
```
Где {URL} - адрес проверяемого сайта; {RETURNCODE} - возвращаемый http-статус.

# Файл с сайтами.
Данное приложение берет данные о проверяемых сайтах из файла sites.txt. Его формат должен принимать примерно подобный вид:
```
http://google.com
http://yandex.ru
```
То есть каждый сайт с новой строки.