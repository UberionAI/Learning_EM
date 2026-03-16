Поднимаем контейнер ElasticSearch с следующими параметрами:

`docker run -d  --name elasticsearch -p 9200:9200 -e "discovery.type=single-node" -e "xpack.security.enabled=false" -e "ES_JAVA_OPTS=-Xms1g -Xmx1g" elasticsearch:8.10.2`

Команда позволяет запустить эластик с отключёнными проверками безопасности -> можно постучаться на расшаренный порт эластика:

в бразуере: http://localhost:9200/

Ответ:
```
{
"name" : "056a4788f984",
"cluster_name" : "docker-cluster",
"cluster_uuid" : "eyxbQkhbSo22RpifAo4qtQ",
"version" : {
"number" : "8.10.2",
"build_flavor" : "default",
"build_type" : "docker",
"build_hash" : "6d20dd8ce62365be9b1aca96427de4622e970e9e",
"build_date" : "2023-09-19T08:16:24.564900370Z",
"build_snapshot" : false,
"lucene_version" : "9.7.0",
"minimum_wire_compatibility_version" : "7.17.0",
"minimum_index_compatibility_version" : "7.0.0"
},
"tagline" : "You Know, for Search"
}
```

