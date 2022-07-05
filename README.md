# scrappy-plus

Одним скучным вечером я решил что очень-очень хочу сходить на [Плюс.Дачу от Яндекса](https://plus.yandex.ru/dacha) в
Москве,
но на сайте не было свободных билетов. Тогда, чтобы успеть записаться, я решил, нужен бот, который скажет мне, когда
появится свободный слот.

Так и родился scrappy-plus - бот, который поможет вам записаться на ПлюсДачу Яндекса.

## Настройка

Настройка стандартно выполняется через установу ENV-переменных, приведенных в таблице ниже

| Название   | Описание                                                                                  | Пример                       |
|------------|-------------------------------------------------------------------------------------------|------------------------------|
| `PARSE_URL` | Ссылка на сайт Плюс.Дачи                                                                  | https://plus.yandex.ru/dacha |
| `TELEGRAM_TOKEN`           | Токен вашего бота-напоминальщика в Telegram                                               | 5596938288:thisIsMyToken     |
|        `TELEGRAM_CHANNEL_ID`    | Идентификатор вашего пользователя или группы в Telegram, куда бот будет писать обновления |                 260780920             |

## Деплой

Поддерживается либо запуск из бинарника, либо через Yandex Cloud Functions.
