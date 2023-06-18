# Minic

Небольшой бот помощник созданный для личных нужд. В каком то смысле этот репозиторий используется как хранилище с функцией удобной демонстрации. Вдруг вернусь к этому проекту через какое то время ¯\\\_(ツ)_/¯.

Основная логика бота содержится в папке `minic`. Соседние папки по типу `telegram` являются обёртками для взаимодействия с ботом.

Не плохо было бы прикрутить ИИ для распознавания запросов, но сейчас мне кажется это чем то из области фантастики ;)

Запуск: 
```powershell
go run main.go <TelegramBotToken>
```

## Возможности:

`Replace:` находит в тексте определённые символы и заменяет их на требуемые.

&nbsp; &nbsp; &nbsp; `Telegram:` ```/replace text:<text> old:<text> new:<text>```
