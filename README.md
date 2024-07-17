# REST API 
Для роутинга используется framework GIN, для хранения данных Postgrsql c GORM.

При запуске автоматически создается таблица "account"
Структура таблицы: Id, Name, Surname, Balance

**Все методы описаны с помощью swagger, можно попробовать по ссылке** - http://localhost:8080/docs/index.html


**cmd/main.go** - точка входа

**config/.env** - Конфигурация подключения к БД

**handlers/** - хэндлеры

**helpers/** - вспомогательные модели и функции

**models/** - основные модели и методы


