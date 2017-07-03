package handlers

import "github.com/couchbase/gocb"

func ListAllMessage(cliens Client, bucket *gocb.Bucket) []Message {
	var searchCliens Client
	bucket.Get(cliens.Inner.LoginName, &searchCliens)
	return searchCliens.Inner.Messages
}
func AddMessageClient(client *Client, message Message, bucket *gocb.Bucket) {
	client.Inner.Messages = append(client.Inner.Messages, message)
	bucket.Upsert(client.Inner.LoginName, client, 0)

}

func MessageSaveToClient(message *Message, client Client) {
	var sClient Client
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ := cluster.OpenBucket("chatapp", "")
	bucket.Get(client.Inner.LoginName, &sClient)
	sClient.Inner.Messages = append(sClient.Inner.Messages, *message)
	bucket.Upsert(sClient.Inner.LoginName, sClient, 0)
}
