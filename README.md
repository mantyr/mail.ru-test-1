# Тестовое задание для Mail.Ru Group

## Текст задания:

Процессу на stdin приходят строки, содержащие интересующие нас URL. Каждый такой URL нужно дернуть и посчитать кол-во вхождений строки "Go". В конце работы приложение выводит на экран общее кол-во найденных строк "Go" во всех переданных URL, например:

```Bash
$ echo -e 'https://golang.org\nhttps://golang.org\nhttps://golang.org' | go run 1.go
Count for https://golang.org: 9
Count for https://golang.org: 9
Count for https://golang.org: 9
Total: 27
```

Введенный URL должен начать обрабатываться сразу после вычитывания и параллельно с вычитыванием следующего. URL должны обрабатываться параллельно, но не более k=5 одновременно.

Обработчики url-ов не должны порождать лишних горутин, т.е. если k=1000 а обрабатываемых URL-ов нет, не должно создаваться 1000 горутин.

Нужно обойтись без глобальных переменных и использовать только стандартные библиотеки.

# Examples:

```Bash
$ echo -e 'https://golang.org\nhttps://golang.org\nhttps://golang.org' | go run ./cmd/v1/*.go -search="Go"
Count for https://golang.org: 9
Count for https://golang.org: 9
Count for https://golang.org: 9
Total:  27

$ echo -e 'https://golang.org\nhttps://golang.org\nhttp://4gophers.ru' | go run ./cmd/v1/*.go -search="Go"
Count for https://golang.org: 9
Count for https://golang.org: 9
Count for http://4gophers.ru: 39
Total:  57


$ echo -e 'https://golang.org\nhttps://golang.org\nhttps://golang.org' | go run ./cmd/v2/*.go -search="Go"
Count for https://golang.org: 9
Count for https://golang.org: 9
Count for https://golang.org: 9
Total:  27

$ echo -e 'https://golang.org\nhttps://golang.org\nhttp://4gophers.ru' | go run ./cmd/v2/*.go -search="Go"
Count for https://golang.org: 9
Count for https://golang.org: 9
Count for http://4gophers.ru: 39
Total:  57

```

# Примечания:

1. v1 - вариант обёртки над WaitGroup с ограничением

- вариант "в лоб"

2. v2 - вариант с ограничением по единичному каналу длиной в максимальное количество горутин

- самый простой

- минус: при ограничении в 1М одновременных горутин создаст канал в 1М элементов


Вариант 1 появился исходя из недостатка варианта 2, однако сам по себе достаточно сложный, без добавления тестов и необходимости добавить какую-то дополнительную логику 
(если такую нельзя реализовать в варианте 2) этот вариант сам по себе избыточен, но как пример сответствует задачи

Вариант 1 реализован в виде обёртки над sync.WaitGroup, а не в виде изменённой реализации по следующей причине: не хотелось усложнять, это требовало бы обязательных более сложных и продуманных тестов

Вариант 2 появился так как он проще, элементарнее и, вероятно, требует меньше операций, однако при ограничении в 1М горутин придётся выделить канал на 1М элементов, 
хотя struct{} по сути не несут в себе значения, всё таки они требуют небольшого выделения памяти в самом канале

Скорее всего это не все варианты которые можно придумать.

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr