package sources

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hlog/internal/conf"
	"log"
)

var MongoSource = &mongoSource{}

type mongoSource struct {
	db              *mongo.Database
	Client          *mongo.Client
	collectionNames []string
}

func (source *mongoSource) Setup(config *conf.MongoDBConfig) {
	source.Connect(config)
	source.getCollectionNames()
}

func (source *mongoSource) Connect(conf *conf.MongoDBConfig) {
	var uri string
	// mongodb://homi-admin:fndsio8rh0fds@111.230.227.84:27016/?connectTimeoutMS=10000&authSource=Homi&authMechanism=SCRAM-SHA-1
	if conf.UserName == "" {
		uri = fmt.Sprintf("mongodb://%s:%s/", conf.Host, conf.Port)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/", conf.UserName, conf.Password, conf.Host, conf.Port)
	}
	var err error
	source.Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("[mongodata.connect] %s\n", err)
		panic(err)
	}
	source.db = source.Client.Database(conf.DataBase)
	log.Printf("[mongodata.connect] connect success")
}

func (source *mongoSource) getCollectionNames() {
	collections, err := source.db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Printf("[mongodata.getCollectionNames] %s", err)
	}
	source.collectionNames = collections
}

func (source *mongoSource) Close() {
	err := source.Client.Disconnect(context.TODO())
	if err != nil {
		// 这里不用panic是因为需要在恢复函数里关闭
		log.Printf("[mongodata.Close] %s", err)
	}
}

func (source *mongoSource) CollectionExist(name string) bool {
	for _, collections := range source.collectionNames {
		if collections == name {
			return true
		}
	}
	return false
}

func (source *mongoSource) GetCollection(name string) *mongo.Collection {
	return source.db.Collection(name)
}

func (source *mongoSource) GetOrCreateCollection(name string) *mongo.Collection {
	if !source.CollectionExist(name) {
		err := source.db.CreateCollection(context.TODO(), name)
		if err != nil {
			log.Panicf("[mongodata.GetOrCreateCollection]:%s", err)
		}
	}
	return source.GetCollection(name)
}
