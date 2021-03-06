package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"micro_services/api/v1/port"
)

type PortRepository struct {
	*mongo.Client
}

func (p PortRepository) UpsertPort(ctx context.Context, port port.Port) error {
	collection := p.Database("90poe").Collection("ports")
	filter := bson.M{"id": bson.M{"$eq": port.ID}}
	opt := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, bson.D{
		{"$set", port},
	}, opt)
	return err
}
func (p PortRepository) GetByID(ctx context.Context, id string) (*port.Port, error) {
	collection := p.Database("90poe").Collection("ports")
	fPort := port.Port{}
	err := collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&fPort)
	return &fPort, err
}
