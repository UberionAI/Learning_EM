## Задание №1 Разработка простого микросервиса
 **Создать микросервис для управления пользователями (регистрация, обновление данных, удаление, получение списка)**

- Go 1.21+ (стандартная библиотека: net/http, encoding/json).

1. `cd task_1`

2. `go run ./cmd/main.go`

Сервер запустится на http://localhost:8080

API:

* POST /users          {"name": "Test"}     — регистрация
* GET /users           — список пользователей
* PUT /users/1         {"id":"1","name":"New"} — обновление
* DELETE /users/1      — удаление

Бэнчмарки:

`go test -bench=BenchmarkListUsers -benchtime=5s -cpu=4`

## Задание №2 Внедрение gRPC
**Задание: Переделать первый микросервис с использованием gRPC вместо REST. Реализовать те же методы для управления пользователями, но с использованием Protocol Buffers для сериализации данных.**

Требования:
Go 1.21+
protoc + protoc-gen-go (установить: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest)

1. Подготовка (генерация pb.go):

`cd task_2/proto`

`protoc --go_out=. --go-grpc_out=. user.proto`
_создаст 2 файла в proto/: user.pb.go и user_grpc.pb.go_

2. Запуск сервера:

`cd task_2`

`go run cmd/main.go`
_в консоли должно отобразиться: "gRPC server:8080"_

3. Запуск клиента
`go run cmd/client/main.go`
_создаст/покажет/обновит/удалит пользователя_

### API (gRPC UserService):
* CreateUser(name) → User
* ListUsers() → []User
* UpdateUser(id, name) → User
* DeleteUser(id)

4. Бэнчмарки

Из корня задания:
`go test -bench=BenchmarkListUsers -benchtime=5s -cpu=4`

## Задание №3 Реализация аутентификации в gRPC
**Задание: Добавить в микросервис аутентификацию с использованием JWT токенов. Токены должны быть переданы в каждом запросе, и на сервере их необходимо валидировать.**

Требования:
* Go 1.21+
* protoc + protoc-gen-go (установить: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest)

Подготовка (генерация pb.go):

`cd task_3/proto`

`protoc --go_out=. --go-grpc_out=. user.proto`
_создаст 2 файла в proto/: user.pb.go и user_grpc.pb.go_

Запуск сервера:

`cd task_3`

`go run cmd/main.go`

_В cmd/main.go создаётся gRPC‑сервер на порту :8080 и вешается JWT‑перехватчик AuthUnaryInterceptor_

Запуск клиента:

`go run cmd/client/main.go`

Методы сервиса те же, что и в задании №2:
* CreateUser(name) → создаёт пользователя и возвращает его.
* ListUsers() → возвращает список всех пользователей.
* UpdateUser(id, name) → обновляет имя пользователя.
* DeleteUser(id) → удаляет пользователя.

_**Все эти вызовы требуют валидного JWT‑токена.**_

## Задание №4 WebSocket сервер
**Задание: Реализовать микросервис на Go, который принимает WebSocket соединения и позволяет клиентам отправлять и получать сообщения в реальном времени. Добавить поддержку нескольких типов сообщений (например, текстовые и бинарные данные).**

Микросервис поднимает WebSocket‑сервер, принимает подключения от клиентов и рассылает сообщения всем подключённым клиентам в реальном времени.

Требования
- Go 1.21+ (подойдёт и любая свежая версия Go).
- Интернет для установки зависимостей (пакет `github.com/gorilla/websocket`)

`go get github.com/gorilla/websocket@latest` _если нет этой библиотеки_

Запуск:

`cd task_4`

`go run main.go`

Сервер теперь доступен по адресу:
[ws://localhost:8080/ws]


## Задание №5 Коммуникация между микросервисами
**Задание: Разработать два микросервиса: первый для управления пользователями (как в задании 1), второй для отправки уведомлений пользователям. Настроить взаимодействие между этими сервисами через Kafka или RabbitMQ. Когда пользователь регистрируется, отправляется сообщение на брокер, а второй сервис обрабатывает это сообщение и отправляет уведомление.**

- `docker-compose.yml` — поднимает брокер сообщений
- `userservice/` — сервис управления пользователями
- `notificationservice/` — сервис отправки уведомлений

Требования:
- Docker и Docker Compose установлены и работают
- Go 1.21+

Запуск брокера сообщения из-под compose (например, в фоне):

`cd task_5`

`docker-compose up -d`

Запуск сервиса userservice (для управления пользователями):

`cd \task_5\userservice\cmd`

`go run main.go`

Запуск сервиса notificationservice (для отправки уведомлений пользователям):

`cd \task_5\notificationservice\cmd`

`go run main.go`

При создании нового юзера теперь в консоли с запущенным notificationservice видим приветствующее сообщение нового пользователя


## Задание №6 Мониторинг и логирование
**Задание: Интегрировать Prometheus и Grafana для мониторинга производительности одного из разработанных микросервисов. Настроить логирование с использованием ELK стека.**

Структура проекта

Корень `task_6`:
- `cmd/main.go` — запуск HTTP‑сервера, регистрация хендлеров и метрик
- `internal/handlers/handlers.go` — логика CRUD по пользователям
- `internal/models/models.go`
- `Dockerfile` — образ приложения (для сервиса `app`)
- `docker-compose.yml` — поднимает:
    - `app` (наш сервис),
    - `prometheus`,
    - `grafana`,
    - `elasticsearch`,
    - `logstash`,
    - `kibana`
- `prometheus.yml` — конфиг Prometheus
- `logstash.conf` — конфиг Logstash (читает логи и шлёт в Elasticsearch)

Требования:
- Docker и Docker Compose установлены и работают
- Go 1.21+

Запуск всех сервисов через композник:

`cd task_6`

`docker-compose up`

_Терминал с docker-compose up лучше не закрывать — в нём видно логи всех компонентов._

Сервисы
* Метрики Prometheus: http://localhost:8081/metrics
* Prometheus: http://localhost:9090
* Grafana http://localhost:3000 (_логин/пароль admin/admin_)
* Kibana http://localhost:5601

Пример настройки логирования с помощью ELK-стека:

- Открыть Кибану http://localhost:5601

- Management → Index Patterns (или аналогичный раздел),

- создайте индекс‑паттерн logs-*

- После этого вы сможете видеть логи, которые попадают из Logstash в Elasticsearch


_Logstash настроен на приём:_
* UDP‑логов на порту 5000
* файлов /var/log/*.log внутри контейнера Logstash
* Logstash складывает всё в Elasticsearch в индекс вида logs-YYYY.MM.dd

Остановка приложения:

`cd task_6`

`docker-compose down`

или Ctrl+C в терминале с запущенным композником

## Задание №7 Конфигурация через environment variables
**Задание: Разработать микросервис с поддержкой конфигурации через переменные окружения. Например, конфигурация базы данных, порта и других параметров должна быть загружена из environment переменных.**

Требования:
- Go 1.21+ 

Создайте .env файл в корне с данным заданием (/task_7/.env). Переменные окружения (приложение читает следующие переменные):
- `SERVER_PORT` — порт HTTP‑сервера
- `DB_HOST` — хост БД
- `DB_PORT` — порт БД
- `DB_NAME` — имя БД
- `DB_USER` — пользователь БД
- `DB_PASSWORD` — пароль БД

Заполняем корректные данные в .env файл

Запуск приложения:

`cd task_7`

`go run main.go`

Проверьте эндпоинт /health в браузере или через curl: http://localhost:8081/health

