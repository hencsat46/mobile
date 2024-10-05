package dataaccess

import "go.mongodb.org/mongo-driver/mongo"

type DataAccess struct {
	mongoConnection *mongo.Client
}

func NewDataAccess(conn *mongo.Client) *DataAccess {
	return &DataAccess{
		mongoConnection: conn,
	}
}
