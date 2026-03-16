## Задание 2. Создание и индексирование данных

Создаём новый индекс products в Elasticsearch:

PUT http://localhost:9200/products
```
{
"settings": {
"number_of_shards": 1,
"number_of_replicas": 0
    }
}
```

Получаем ответ:
```
{
"acknowledged": true,
"shards_acknowledged": true,
"index": "products"
}
```

Добавим несколько документов в проиндексированный эластик:
POST http://localhost:9200/products/_doc/1

`{
"name": "Laptop",
"category": "Electronics",
"price": 1000
}`

Ответ:

```{
"_index": "products",
"_id": "1",
"_version": 1,
"result": "created",
"_shards": {
"total": 1,
"successful": 1,
"failed": 0
},
"_seq_no": 0,
"_primary_term": 1
}
```
POST http://localhost:9200/products/_doc/2

`{
"name": "Phone",
"category": "Samsung",
"price": 2500
}`

POST http://localhost:9200/products/_doc/3

`{
"name": "PC",
"category": "DEXP",
"price": 25000
}`