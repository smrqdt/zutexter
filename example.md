```sh
curl --request POST \
  --url http://192.168.88.242/queue \
  --header 'content-type: multipart/form-data' \
  --form 'content={
            "type": "scroll",
            "content": "Hallo Rob"
        }'
```

```sh
curl --request POST \
  --url http://192.168.88.242/repertoire \
  --header 'content-type: multipart/form-data' \
  --form 'content={
    "screens": [
        {
            "type": "scroll",
            "content": "Next Stop: Day 3"
        },
        {
            "type": "scroll",
            "content": "Use more bandwidth"
        },
        {
            "type": "static",
            "content": "38c3",
            "timer": 1000
        },
        {
            "type": "static",
            "content": "Illegal Instructions",
            "timer": 2000
        },
        {
            "type": "scroll",
            "content": "+++ Kein Kaese mehr +++"
        }
    ]
}'
```
