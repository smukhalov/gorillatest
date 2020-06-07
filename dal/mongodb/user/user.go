package user

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gorillatest/common"
	"gorillatest/model/db"
)

func Insert(user *db.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(common.ConnectionString))
	if err != nil {
		return err
	}

	users := client.Database("smgoz").Collection("users")
	result, err := users.InsertOne(ctx, user)
	if err != nil {
		merr, ok := err.(mongo.WriteException)
		if !ok {
			return err
		}

		errCode := merr.WriteErrors[0].Code
		if errCode == 11000 {
			fmt.Println("Duplicate found")
			return fmt.Errorf("Пользователь %s уже зарегистрирован", user.Email)
		}

		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); !ok {
		return fmt.Errorf("Ошибка преобразования InsertedID")
	} else {
		user.Id = oid.Hex()
	}

	return nil
}

func GetByEmail(email string) (*db.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	var user db.User
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(common.ConnectionString))
	if err != nil {
		return nil, err
	}

	users := client.Database("smgoz").Collection("users")
	result := users.FindOne(ctx, bson.M{"email": email})
	err = result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetById(userid string) (*db.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	var user db.User
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(common.ConnectionString))
	if err != nil {
		return nil, err
	}

	users := client.Database("smgoz").Collection("users")
	if err != nil {
		return nil, err
	}

	//filter := bson.D{primitive.E{Key: "_id", Value: userid}}
	objectId, err := primitive.ObjectIDFromHex(userid)

	result := users.FindOne(ctx, bson.M{"_id": objectId})
	err = result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
