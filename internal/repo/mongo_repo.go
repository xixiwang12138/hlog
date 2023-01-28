package repo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hlog/internal/sources"
	"reflect"
)

type MongoRepo[M any] struct {
	collection     *mongo.Collection
	collectionName string
}

func NewMongoRepo[M any]() *MongoRepo[M] {
	var m M
	rep := &MongoRepo[M]{}
	rep.collectionName = reflect.TypeOf(m).Name()
	rep.collection = sources.MongoSource.GetOrCreateCollection(rep.collectionName)
	return rep
}

func (rep *MongoRepo[M]) SetupRepo() {
	var m M
	rep.collectionName = reflect.TypeOf(m).Name()
	rep.collection = sources.MongoSource.GetOrCreateCollection(rep.collectionName)
}

func (rep *MongoRepo[M]) GetCollection() *mongo.Collection {
	return rep.collection
}

func (rep *MongoRepo[M]) SwitchCollection(collectionName string) {
	rep.collectionName = collectionName
	rep.collection = sources.MongoSource.GetOrCreateCollection(rep.collectionName)
}

// Create 如果mongoid字段存在，则更新，否则插入
func (rep *MongoRepo[M]) Create(instance *M) (err error) {
	_, err = rep.collection.InsertOne(context.TODO(), instance) //可能不需要回写
	if err != nil {
		return errors.Wrap(err, "Insert One ======>")
	}
	return nil
}

func (rep *MongoRepo[M]) FindOne(filter bson.D, opts ...*options.FindOneOptions) (*M, error) {
	res := new(M)
	err := rep.collection.FindOne(context.TODO(), filter).Decode(res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		} else {
			return nil, errors.Wrap(err, "Find One =========> ")
		}
	}
	return res, nil
}

func (rep *MongoRepo[M]) Find(filter bson.D, opts ...*options.FindOptions) ([]*M, error) {
	cursor, err := rep.collection.Find(context.TODO(), filter, opts...)
	defer cursor.Close(context.TODO())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		} else {
			return nil, errors.Wrap(err, "Find Many =========> ")
		}
	}
	var resList []*M
	err = cursor.All(context.TODO(), &resList)
	if err != nil {
		return nil, errors.Wrap(err, "Decode Many Results =========> ")
	}
	return resList, nil
}

// FindByObjectId 根据mongo主键查询
func (rep *MongoRepo[M]) FindByObjectId(_id primitive.ObjectID) (*M, error) {
	return rep.FindOne(bson.D{{"_id", _id}})
}

func (rep *MongoRepo[M]) UpdateOne(filter bson.D, update bson.D) error {
	_, err := rep.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.Wrap(err, "Update One =========> ")
	}
	return nil
}

// endregion
