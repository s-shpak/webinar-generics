# Дженерики

## Что нового добавлено

См. примеры в `internal/intro`

## Частичная мономорфизация

Техническое описание имплементации дженериков можно найти здесь: https://github.com/golang/proposal/blob/master/design/generics-implementation-gcshape.md

Установите `lensm` для более удобного чтения ассемблерного кода: https://github.com/loov/lensm

Скомпилируем пример и откроем код в `lensm`:

```bash
make build-mono
lensm -watch -filter main cmd/monomorphization/monomorphization
```

Проверим, как работает частичная мономорфизация с указателями:

```bash
go test -gcflags '-N -l' --bench=. --benchtime=5s
```

Для того, чтобы понять, почему так, почитаем ассемблер:

```bash
make build-bench
lensm -watch -filter main cmd/bench/bench
```

## Практика

См. упражнения в `internal/practice`.