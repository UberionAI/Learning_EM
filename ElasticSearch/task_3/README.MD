## Задание 3. Поиск данных

1. Чтобы найти все документы из эластика воспользуемся:

GET http://localhost:9200/products/_search?pretty

Ответ:
```
{
    "took": 84,
    "timed_out": false,
    "_shards": {
        "total": 1,
        "successful": 1,
        "skipped": 0,
        "failed": 0
    },
    "hits": {
        "total": {
            "value": 3,
            "relation": "eq"
        },
        "max_score": 1.0,
        "hits": [
            {
                "_index": "products",
                "_id": "1",
                "_score": 1.0,
                "_source": {
                    "name": "Laptop",
                    "category": "Electronics",
                    "price": 1000
                }
            },
            {
                "_index": "products",
                "_id": "2",
                "_score": 1.0,
                "_source": {
                    "name": "Phone",
                    "category": "Samsung",
                    "price": 2500
                }
            },
            {
                "_index": "products",
                "_id": "3",
                "_score": 1.0,
                "_source": {
                    "name": "PC",
                    "category": "DEXP",
                    "price": 25000
                }
            }
        ]
    }
}
```

2. Находим все товары из категории Electronics:

GET http://localhost:9200/products/_search?pretty
```
{
"query": {
"match": {
"category": "Electronics"
}
}
}
```

Ответ:
```
{
    "took": 12,
    "timed_out": false,
    "_shards": {
        "total": 1,
        "successful": 1,
        "skipped": 0,
        "failed": 0
    },
    "hits": {
        "total": {
            "value": 1,
            "relation": "eq"
        },
        "max_score": 0.9808291,
        "hits": [
            {
                "_index": "products",
                "_id": "1",
                "_score": 0.9808291,
                "_source": {
                    "name": "Laptop",
                    "category": "Electronics",
                    "price": 1000
                }
            }
        ]
    }
}
```

Поиск с фильтрацией по цене, например больше 500:

GET http://localhost:9200/products/_search

```
{
"query": {
        "range": {
            "price": {
                "gte": 500
            }
        }
    }

}
```

Ответ:
```
{
    "took": 10,
    "timed_out": false,
    "_shards": {
        "total": 1,
        "successful": 1,
        "skipped": 0,
        "failed": 0
    },
    "hits": {
        "total": {
            "value": 3,
            "relation": "eq"
        },
        "max_score": 1.0,
        "hits": [
            {
                "_index": "products",
                "_id": "1",
                "_score": 1.0,
                "_source": {
                    "name": "Laptop",
                    "category": "Electronics",
                    "price": 1000
                }
            },
            {
                "_index": "products",
                "_id": "2",
                "_score": 1.0,
                "_source": {
                    "name": "Phone",
                    "category": "Samsung",
                    "price": 2500
                }
            },
            {
                "_index": "products",
                "_id": "3",
                "_score": 1.0,
                "_source": {
                    "name": "PC",
                    "category": "DEXP",
                    "price": 25000
                }
            }
        ]
    }
}
```

