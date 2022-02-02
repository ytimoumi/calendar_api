# giskard_api
> A GraphQl API with Golang

###  DEV Environment DOCKER
``` bash
#build
docker-compose build
#run
docker-compose up -d
#to show logs
docker-compose logs -f

```

###  APIs
``` bash
## Referer: http://test-ytimoumi.com =========> client: test
#creater reservation
curl --location --request POST 'http://localhost:7010/query' \
--header 'Referer: http://test-ytimoumi.com' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation ($idAvailability:Int!,$input:ToReserve!) {\r\n  createReservation(idAvailability:$idAvailability, input:$input) {\r\n      code\r\n      message\r\n  }  \r\n}","variables":{"idAvailability":7,"input":{"start":"2022-02-09T08:04:30Z","end":"2022-02-09T08:30:30Z","title":"Entretien","email":"yassinetimoumi.official@gmail.com"}}}'

#create available slot
curl --location --request POST 'http://localhost:7010/query' \
--header 'Referer: http://test-ytimoumi.com' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation ($input:Availability!) {\r\n  createAvailability(input:$input) {\r\n      code\r\n      message\r\n  }  \r\n}","variables":{"input":{"start":"2022-02-09T12:04:30Z","end":"2022-01-02T09:04:30Z"}}}'

#get available slots
curl --location --request POST 'http://localhost:7010/query' \
--header 'Referer: http://test-ytimoumi.com' \ 
--header 'Content-Type: application/json' \
--data-raw '{"query":"query  {\r\n  getCalendar {\r\n      start\r\n      end\r\n  }  \r\n}","variables":{}}'

#delete reservation
curl --location --request POST 'http://localhost:7010/query' \
--header 'Referer: http://test-ytimoumi.com' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation ($id: Int!, $email: String!) {\r\n  deleteReservation(id:$id, email:$email) {\r\n      code\r\n      message\r\n  }  \r\n}","variables":{"id":5,"email":"yassinetimoumi.official@gmail.com"}}'

#delete availability
curl --location --request POST 'http://localhost:7010/query' \
--header 'Referer: http://test-ytimoumi.com' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation ($id: Int!) {\r\n  deleteAvailability(id:$id) {\r\n      code\r\n      message\r\n  }  \r\n}","variables":{"id":5}}'
```