package handlers

import (
	"fmt"
	"log"

	"github.com/couchbase/gocb"
	"github.com/davidborka/chatApp/api/dbconnect"
	"github.com/davidborka/chatApp/api/model"
	"github.com/google/uuid"
)

func NexusIsExixst(login1 string, login2 string) string {

	bucketMessage := dbconnect.DatabaseConnectionMessage()
	var conversations model.Conversations
	var client1 model.Client
	var client2 model.Client
	FinduserUuid(login1, &client1)
	FinduserUuid(login2, &client2)
	bucketMessage.Get("conversations", &conversations)

	for _, v := range conversations.Conv {
		if (v.Partner2Uuid == client1.Uuid && v.Partner1Uuid == client2.Uuid) || (v.Partner2Uuid == client2.Uuid && v.Partner1Uuid == client1.Uuid) {
			return v.NexusUuid
		}
	}
	return "NOT FOUND"

	/*bucket := dbconnect.DatabaseConnection()
	var client model.Client
	var client2 model.Client
	count := 0
	bucket.Get(loginName, &client)
	bucket.Get(loginName2, &client2)

	if client.Nexusid != nil {
		for k, _ := range client.Nexusid {
			if k == loginName2 {
				count++
			}
		}
		for k, _ := range client2.Nexusid {
			if k == loginName {
				count++
			}
		}
		if count == 2 {
			return client.Nexusid[loginName2]
		}
	}
	return "NOT FOUND"*/
}
func AddConversations(lname1 string, lname2 string, nexusuuid string) {

	bucketMessage := dbconnect.DatabaseConnectionMessage()

	var conversation model.Conversation
	var conversations model.Conversations
	var client1 model.Client
	var client2 model.Client
	FinduserUuid(lname1, &client1)
	FinduserUuid(lname2, &client2)
	bucketMessage.Get("conversations", &conversations)
	conversation.Partner1Uuid = client1.Uuid
	conversation.Partner2Uuid = client2.Uuid
	conversation.NexusUuid = nexusuuid
	conversations.Conv = append(conversations.Conv, conversation)
	bucketMessage.Upsert("conversations", conversations, 0)
}
func ListAllClientFromDb(bucket *gocb.Bucket) []model.TypedClient {

	myQuery := gocb.NewN1qlQuery("SELECT * FROM chatappUsers")
	rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)
	if err != nil {
		fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
	} // Interfaces for handling streaming return values
	var row model.TypedClient
	var retValues []model.TypedClient

	for rows.Next(&row) {
		retValues = append(retValues, row)
		row = model.TypedClient{}
	}
	return retValues
}
func AddClient(client model.Client, bucket *gocb.Bucket) bool {
	if _, err := bucket.Upsert(client.Uuid, client, 0); err != nil {
		log.Fatal(err)
		fmt.Println("Can't inser new client! The client who caused problem " + client.LoginName)
		return false
	}
	return true
}
func ListActiveCliens(activeClient model.Connection) model.OnlineClient {
	var active model.OnlineClient

	for _, client := range activeClient.ActiveCLient {

		active.Inner.LoginName = append(active.Inner.LoginName, client.LoginName)

	}

	return active
}

//UuidConGenerator is generate uuid and covert to string
func UuidGenerator() string {
	u1, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	u2 := u1.String()
	return u2
}
func FinduserUuid(loginname string, clientReturn *model.Client) error {

	query := gocb.NewN1qlQuery(`SELECT * FROM chatappUsers WHERE chatappUsers.loginname="` + loginname + `"`)
	bucket := dbconnect.DatabaseConnectionUsers()
	rows, err := bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		log.Fatal(err)

	}
	var row model.TypedClient
	var retvalue []model.TypedClient
	for rows.Next(&row) {
		retvalue = append(retvalue, row)
	}

	clientReturn.Uuid = retvalue[0].Chatapp.Uuid
	clientReturn.LoginName = retvalue[0].Chatapp.LoginName
	clientReturn.Email = retvalue[0].Chatapp.Email
	clientReturn.Password = retvalue[0].Chatapp.Password
	return nil

}
