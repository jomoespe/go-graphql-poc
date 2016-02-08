GrahpQL Go sample
=================

Following the medium article [Learn Golang + GraphQL + Relay](https://wehavefaces.net/learn-golang-graphql-relay-1-e59ea174a902#.3o36mv5as), with some changes derivated of [graphql-go](https://github.com/graphql-go/graphql) evolution.


## Test

	run main.go

    curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'query Root{ latestPost }'
    curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d '{ latestPost }'
