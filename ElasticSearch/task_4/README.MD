## Задание 4. Работа с маппингом

1.	Создайте индекс users с кастомным маппингом, где поле email имеет тип keyword, а name — text.
`curl -X PUT "http://localhost:9200/users" \
  -H 'Content-Type: application/json' \
  -d '{
    "mappings": {
      "properties": {
        "name": {
          "type": "text"
        },
        "email": {
          "type": "keyword"
        }
      }
    }
  }'`

Ответ:
`{"acknowledged":true,"shards_acknowledged":true,"index":"users"}`
2.	Добавьте в индекс несколько пользователей.

```
curl -X POST "http://localhost:9200/users/_doc/1" \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Alex",
    "email": "alex@example.com"
}'
{"_index":"users","_id":"1","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1}
```

```
curl -X POST "http://localhost:9200/users/_doc/2" \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Bob",
    "email": "bob@example.com"
}'
{"_index":"users","_id":"2","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":1,"_primary_term":1}
```

```
curl -X POST "http://localhost:9200/userscurl -X POST "http://localhost:9200/users/_doc/3" \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Vova",
    "email": "vova@example.com"
}'
{"_index":"users","_id":"3","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":2,"_primary_term":1}
```

4. Выполните поиск пользователей по точному значению email.
```
curl -X GET "http://localhost:9200/users/_search?pretty" \
  -H 'Content-Type: application/json' \
  -d '{
    "query": {
      "term": {
        "email": {
          "value": "alex@example.com"
        }
      }
    }
  }'
  
 
{
  "took" : 13,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 1,
      "relation" : "eq"
    },
    "max_score" : 0.9808291,
    "hits" : [
      {
        "_index" : "users",
        "_id" : "1",
        "_score" : 0.9808291,
        "_source" : {
          "name" : "Alex",
          "email" : "alex@example.com"
        }
      }
    ]
  }
}
```
