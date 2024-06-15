## Задание: разработать систему управления библиотекой

Разработать REST API на Go для управления коллекцией книг и авторами,которые их пишут. API будет выполнять операции CRUD с книгами и авторами, хранящимися в базе данных PostgreSQL. Система должна быть контейнеризирована с помощью Docker, состоящего из двух отдельных контейнеров: одного для веб-приложения и одного для базы данных.

/Скиллы: Golang, Postgres, SQL, docker, работа с json в Golang /

## 1. Требования

  Написать сервис, который будет слушать входящие запросы по HTTP, преобразовывать их в запрос к соответствующей функции Postgres (по схеме трансляции, приведённой ниже), выполнять запрос и возвращать ответ клиенту.
* приложение должно соответствовать набору архитектурных правил REST, реализуя RESTful API на основе протокола HTTP и его методов GET, POST, PUT и DELETE;
* Формат данных: используйте JSON для передачи данных в приложение;
* Подключение к базе данных: установите соединение с базой данных PostgreSQL, развернутой в Docker-контейнере;
* Операции CRUD: реализуйте функциональность операций, используя стандартный пакет database/sql Go;
* Транзакции: реализуйте транзакцию для операции, которая включает одновременное обновление книги и соответствующего автора, например обновление названия книги и биографии автора за один раз. Это следует выполнять только в том случае, если оба обновления пройдут успешно;
* Обработка ошибок: корректные ответы на ошибки в различных сценариях сбоев, например, книга не найдена, передана некорректная дата рождения, ошибки подключения к базе данных и др. Необходимо корректно использовать коды ответов HTTP для таких сценариев.



## 2. Порядок установки, настройки и запуска

####    2.1. Для работы прогрммы необходима установка следущего ПО:
* Компилятор GO:   <https://go.dev/>
* Программа контейнеризации Docker Desktop: <https://www.docker.com/products/docker-desktop/>


####    2.2. Настройки соединения с сервером Postgres читать из файла config.yml:
 - host - hostname, где установлен Postgres (для локальной работы долно быть указано localhost, для работы через контейнер нужно указать имя сервиса контейнера базы данных из файла docker-compose.yml)
 - port - порт, на котором слушать запросы
 - username - имя пользователя Postgres
 - password - пароль пользователя Postgres
 - dbname - имя базы данных Postgres
 - sslmode - режим использования SSL

####    2.3. Для локального запуска

Необхдимо убедиться, что в config.yml параметр host указан как localhost.
Далее можно использовать команду: 
* ```make run``` 

      или последовательно выполнить команды:
* ```go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest```
* ```go mod tidy```
* ```docker pull postgres```
* ```docker run --name=todo-db -e POSTGRES_PASSWORD='12345' -p 5432:5432 -d --rm postgres```
* ```migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' up```
* ```go run cmd/main.go```

####    2.4. Для запуска в контейнере можно использовать команду:

Необхдимо убедиться, что в config.yml параметр host указан как db.
Далее можно использовать команду: 
* ```make build``` 

      или выполнить команду сборки:
* ```docker pull postgres &&```
* ```docker pull golang &&```
* ```docker-compose up -d --build```


####    2.5. Дополнительные возможности
      Также, внутри приложения организована авторизация пользователя по логину и паролю. По умолчанию, данная особенность отключена. Для включения, необходимо в файле HANDLER.GO раскомментарить следующие строки и пересобрать / перезапустить проект.
* ```	/*auth := router.Group("/auth")                 ```
* ```	{                                               ``` 
* ```		auth.POST("/sign-up", h.signUp)         ```  
* ```		auth.POST("/sign-in", h.signIn)         ```  
* ```	}*/                                             ```
* ```                                                   ```  
* ```	api := router.Group("/" /*, h.userIdentity*/)   ``` 



## 3. Организованные end-point(ы) и порядок обращения к сервису.

###    3.1. Для книг:

**POST/books — Добавить новую книгу;**<br>
Обращение: http://localhost:55000/books<br>
Пример теля обращения: <br>
```json
  {
    "name" : "Сами боги2",
    "authorid" : 4,
    "year": 1985,
    "isbn": "978-5-04-172716-1"
  }
```

Получаемый ответ:
```json
	{
    	  "id": 9
	}
```

**GET /books — Получить все книги;**<br>
Обращение: http://localhost:55000/books<br>
Пример тела обращения: пусто<br>
Получаемый ответ: <br>
```json
   {
    "data": [
        {
            "id": 1,
            "name": "Убийство в доме викария",
            "authorid": 1,
            "year": 1954,
            "isbn": "123-343-5465"
        },
        {
            "id": 2,
            "name": "Смерть на Ниле",
            "authorid": 1,
            "year": 1951,
            "isbn": "123-343-5466"
        },
        ....
	
    ]
   }
```

**GET /books/{id} — Получить книгу по ее идентификатору;**<br>
Обращение: http://localhost:55000/books/9<br>
Пример тела обращения: пусто<br>
Получаемый ответ: <br>
```json
   {
    "id": 9,
    "name": "Сами боги2",
    "authorid": 4,
    "year": 1985,
    "isbn": "978-5-04-172716-1"
   }
```

**PUT /books/{id} — обновить книгу по ее идентификатору;**<br>
Обращение: http://localhost:55000/books/9<br>
Пример тела обращения:<br>
```json
   {
    "name" : "adadadad2",
    "authorid" : 4,
    "year": 1977,
    "isbn": "404-233402-3423"
   }    
```

Получаемый ответ: <br>
```json
   {
      "записей обновлено": 1
   }
```

**DELETE /books/{id} — Удалить книгу по ее идентификатору.**<br>
Обращение: http://localhost:55000/books/9<br>
Пример тела обращения: пусто<br>
Получаемый ответ: <br>
```json
   {
      "записей удалено": 1
   }
```


###    3.2. Для авторов:

**POST/authors — Добавить нового автора;**<br>
Обращение: http://localhost:55000/authors<br>
Пример теля обращения: <br>
```json
        {
            "firstname": "aaaa",
            "lastname": "aaaaa",
            "description": "современный писатель-фантаст",
            "birthday": "1919-04-06"
        }
```

Получаемый ответ: <br>
```json
	{
    	  "id": 5
	}
```

**GET /authors — Получить всех авторов;**<br>
Обращение: http://localhost:55000/authors<br>
Пример тела обращения: пусто<br>
Получаемый ответ: <br>
```json
{
  "data": [
        {
            "id": 1,
            "firstname": "Агата",
            "lastname": "Кристи",
            "description": "детективщица!",
            "birthday": "1890-09-15T00:00:00Z"
        },
        {
            "id": 2,
            "firstname": "Артур",
            "lastname": "Конан Дойл",
            "description": "детективщик",
            "birthday": "1859-05-22T00:00:00Z"
        }]
}
```

**GET /authors/{id} — получить автора по его идентификатору;**<br>
Обращение: http://localhost:55000/authors/5<br>
Пример тела обращения: пусто<br>
Получаемый ответ: <br>
```json
{
    "id": 5,
    "firstname": "aaaa",
    "lastname": "aaaaa",
    "description": "современный писатель-фантаст",
    "birthday": "1919-04-06T00:00:00Z"
}
```

**PUT /authors/{id} — обновить автора по его идентификатору;**<br>
Обращение: http://localhost:55000/authors/5<br>
Пример тела обращения:<br>
```json
   {    
    "firstname" : "бббб",
    "lastname" : "бббббб",
    "description": "современный писатель-фантаст",
    "birthday": "1919-04-06"
   }    
```

Получаемый ответ: <br>
```json
   {
      "записей обновлено": 1
   }
```

**DELETE /authors/{id} — удалить автора по его идентификатору.**<br>
Обращение: http://localhost:55000/authors/5<br>
Пример тела обращения: пусто<br>
Получаемый ответ: <br>
```json
   {
      "записей удалено": 1
   }
```

###    3.3. Транзакционное обновление:

**PUT /books/{book_id}/authors/{author_id} - одновременное обновление информации по книге и по автору**<br>
Обращение: hhttp://localhost:55000/books/9/authors/4<br>
Пример тела обращения:<br>
```json
  {
   "author": {
              "firstname": "Айзек",
              "lastname": "Азимов",
              "description": "современный писатель-фантаст!!!!!",
              "birthday": "1919-04-06"
             },
    "book": {
       	      "name" : "Сами боги 22",
              "authorid" : 4,
              "year": 1972,
              "isbn": "978-5-04-172716-1"
            }
   }
```

Получаемый ответ: <br>
```json
  {
    "информация обновлена: ": true
  }
```

###    3.4. Регистрация и логирование (опционально):

**POST /auth/sign-up - регистрация пользователя**<br>
Обращение: http://localhost:55000/auth/sign-up<br>
Пример тела обращения:<br>
```json
{
    "name": "Bug Bunny",
    "username": "rabbit",
    "password": "12345"
}
```

Получаемый ответ: <br>
```json
{
    "id": 1
}
```

**POST /auth/sign-in - логирование пользователя**<br>
Обращение: http://localhost:55000/auth/sign-in<br>
Пример тела обращения:<br>
```json
{    
    "username": "rabbit",
    "password": "12345"
}

Получаемый ответ: <br>
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg1Njg4MDEsImlhdCI6MTcxODQ4MjQwMSwidXNlcl9pZCI6MX0.OcYq0clw4A152yWdKhxuL5F2uw5de8vRzhJgqe_STy4"
}
```


