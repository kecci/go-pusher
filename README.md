# go-pusher

Tutorial: https://www.youtube.com/watch?v=Z8ebUfQOYrM&t=512s

## Pubsub

Website: https://pusher.com/

### Publisher with Go

Run Project
```
go run ./pubsub/publisher/main.go
```

CURL:
```
curl --location --request POST 'http://localhost:9090/api/message'
```

### Consumer with HTML & JS

Run Project
```
open ./pubsub/consumer/index.html
```

Waiting in the browser, until the publisher send message

## Push Notification

Website: https://pusher.com/beams
