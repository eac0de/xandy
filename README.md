# xandy

Клиент и сервер приложения для сохранения нужной информации

## Запуск

Сначала необходимо запустить контейнера docker - `make up`

Далее выбрать нужный клиент и запустить его в терминале (можно скомпилировать и запустить клиент самому из папки client)

PS: Иногда необходимо изменить размер окна, для корректного вывода в клиенте

## Аутентификация

Коды подтверждения отображаются в логах контейнера xandy_auth

PS: Для того чтобы отправлялись письма на почту необходимо установить переменные окружения в deploy/auth

IS_DEV=false

и данные SFTP

SMTPHost=

SMTPPort=

SMTPUsername=

SMTPPassword=

Скачанные файлы сохраняются в директории - `~/xandyFiles/`
