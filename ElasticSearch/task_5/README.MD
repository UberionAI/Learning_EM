## Задание 5. Полнотекстовый поиск
1.	Создайте индекс articles, в котором хранятся статьи с полями title и content.

`curl -X PUT "http://localhost:9200/articles" \
  -H 'Content-Type: application/json' \
  -d '{
    "mappings": {
      "properties": {
        "title": {
          "type": "text"
        },
        "content": {
          "type": "text"
        }
      }
    }
  }'
`

`{"acknowledged":true,"shards_acknowledged":true,"index":"articles"}`

2.	Добавьте несколько статей с различным содержимым.
`
curl -X POST "http://localhost:9200/articles/_doc/1" \
  -H 'Content-Type: application/json' \
  -d '{
    "title": "Introduction to Elasticsearch",
    "content": "This article is an Elasticsearch tutorial for beginners."
}'
`

`
{"_index":"articles","_id":"1","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1}r
`

`
curl -X POST "http://localhost:9200/articles/_doc/2" \
  -H 'Content-Type: application/json' \
  -d '{
    "title": "Full-text search with Elasticsearch",
    "content": "Learn how to perform full-text search using Elasticsearch and its query DSL."
}'
`

```
{"_index":"articles","_id":"2","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":2,"_primary_term":1}
```

`
curl -X POST "http://localhost:9200/articles/_doc/3" \
  -H 'Content-Type: application/json' \
  -d '{
    "title": "Working with Docker",
    "content": "This article explains how to run Elasticsearch in Docker containers."
}'
`

```
{"_index":"articles","_id":"3","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":3,"_primary_term":1}
```


3.	Выполните полнотекстовый поиск статей, содержащих определенные слова в заголовке или контенте.

Cтатьи, где искомая фраза встречается в заголовке или в контенте:

`
curl -X GET "http://localhost:9200/articles/_search?pretty" \
  -H 'Content-Type: application/json' \
  -d '{
    "query": {
      "multi_match": {
        "query": "Elasticsearch tutorial",
        "fields": [ "title", "content" ]
      }
    }
  }'
`

```
{
  "took" : 10,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 3,
      "relation" : "eq"
    },
    "max_score" : 1.2277769,
    "hits" : [
      {
        "_index" : "articles",
        "_id" : "1",
        "_score" : 1.2277769,
        "_source" : {
          "title" : "Introduction to Elasticsearch",
          "content" : "This article is an Elasticsearch tutorial for beginners."
        }
      },
      {
        "_index" : "articles",
        "_id" : "2",
        "_score" : 0.4091398,
        "_source" : {
          "title" : "Full-text search with Elasticsearch",
          "content" : "Learn how to perform full-text search using Elasticsearch and its query DSL."
        }
      },
      {
        "_index" : "articles",
        "_id" : "3",
        "_score" : 0.13531709,
        "_source" : {
          "title" : "Working with Docker",
          "content" : "This article explains how to run Elasticsearch in Docker containers."
        }
      }
    ]
  }
}
```

