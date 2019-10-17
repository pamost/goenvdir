# goenvdir (Утилита envdir на Go)

Реализовать утилиту envdir на Go.
Эта утилита позволяет запускать программы получая переменные окружения из определенной директории.

Пример использования:

`go-envdir /path/to/env/dir some_prog`

Если в директории `/path/to/env/dir` содержатся файлы
`A_ENV` с содержимым `123` и `B_VAR` с содержимым `another_val`, то программа `some_prog` должать быть запущена с переменными окружения `A_ENV=123` и `B_VAR=another_val`
