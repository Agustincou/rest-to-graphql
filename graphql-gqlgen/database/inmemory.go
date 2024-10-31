package database

import "rest-to-graphql/graphql-gqlgen/graph/model"

var Items = make(map[int]*model.Item)
var NextID = 1
