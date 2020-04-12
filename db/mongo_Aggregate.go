package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	aggreate()
}

func aggreate() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // 设置查询的超时时间
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1/dsp"))
	if err != nil {
	}
	slik := []string{"602e5c912ea27c19ca6b08e64f0109d2"}
	ocpc := client.Database("dsp").Collection("dspdata")
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"pid": 3638, "extendmatch": bson.M{"$elemMatch": bson.M{"$in": slik}}}}, // 查询条件
		bson.M{"$unwind": "$extendmatch"}, // 数组拆解
		bson.M{"$match": bson.M{"extendmatch": bson.M{"$in": slik}}},
		bson.M{"$group": bson.M{"_id": "$_id", "count": bson.M{"$sum": 1}}}, // 聚合分组
		bson.M{"$match": bson.M{"count": bson.M{"$gte": 0}}},                // 匹配满足条件的数据
	}
	opts := options.Aggregate()
	cur, err := ocpc.Aggregate(ctx, pipeline, opts)
	defer cur.Close(ctx)

	fmt.Println(cur.Err())

	for cur.Next(ctx) {
		var doc map[string]interface{}
		cur.Decode(&doc)
		fmt.Println(doc)
		//count := interface{}(doc["count"]).(int)
		return
	}
	// dt := []*datatype.Tdtype{}{}
	// cur.All(ctx, dt)
	// cur.
	// dat := &datatype.Tdtype{}
	// // test
	// slik := []string{"6b1f49bb47592950f253ffe38819e81b", "602e5c912ea27c19ca6b08e64f0109d2"}
	// cc, _ := ocpc.Find(ctx, bson.M{"pid": 3638, "extendmatch": bson.M{"$elemMatch": bson.M{"$in": slik}}})
	// for cc.Next(ctx) {
	// 	var doc map[string]interface{}
	// 	cc.Decode(&doc)
	// 	println(doc)
	// }
	// fmt.Println(dat)

}
