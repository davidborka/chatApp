package handlers

import (
	"github.com/couchbase/gocb"
	"github.com/davidborka/chatApp/api/model"
)

func ListAllMessage(cliens model.Client, bucket *gocb.Bucket) []model.Message {
	var searchCliens model.Client
	bucket.Get(cliens.Inner.LoginName, &searchCliens)
	return searchCliens.Inner.Messages
}
func AddMessageClient(client *model.Client, message model.Message, bucket *gocb.Bucket) {
	client.Inner.Messages = append(client.Inner.Messages, message)
	bucket.Upsert(client.Inner.LoginName, client, 0)

}

func MessageSaveToClient(message *model.Message, client model.Client) {
	var sClient model.Client
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ := cluster.OpenBucket("chatapp", "")
	bucket.Get(client.Inner.LoginName, &sClient)
	sClient.Inner.Messages = append(sClient.Inner.Messages, *message)
	bucket.Upsert(sClient.Inner.LoginName, sClient, 0)
}
