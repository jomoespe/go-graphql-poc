GrahpQL Go sample
=================

Following the medium article [Learn Golang + GraphQL + Relay](https://wehavefaces.net/learn-golang-graphql-relay-1-e59ea174a902#.3o36mv5as), with some changes derivated of [graphql-go](https://github.com/graphql-go/graphql) evolution.


## Install GraphQL packages

    go get github.com/graphql-go/graphql
    go get github.com/graphql-go/graphql-go-handler


## Test

Start the GraphQL server:

    run main.go


Send queries to GraphQL server:

    #curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'query Root{ latestPost }'
    #curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d '{ latestPost }'

    curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d '{Vehicle(id:"2") { name, state }}'


Get the schema meta-data

    curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d '{__schema { queryType { name, fields {name, description }}}}'

