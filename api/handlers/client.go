package handlers

import (
	"fmt"

	"github.com/couchbase/gocb"
	"github.com/davidborka/chatApp/api/model"
)

func ListAllClientFromDb(bucket *gocb.Bucket) []model.Client {

	myQuery := gocb.NewN1qlQuery("SELECT * FROM chatapp")
	rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)
	if err != nil {
		fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
	} // Interfaces for handling streaming return values
	var row model.Client
	var retValues []model.Client

	for rows.Next(&row) {
		retValues = append(retValues, row)
		row = model.Client{}
	}
	return retValues
}
func AddClient(client model.Client, bucket *gocb.Bucket) bool {
	if _, err := bucket.Upsert(client.Inner.LoginName, client, 0); err != nil {
		fmt.Println("Can't inser new client! The client who caused problem " + client.Inner.LoginName)
		return false
	}
	return true
}
func ListActiveCliens(activeClient model.Connection) model.OnlineClient {
	var active model.OnlineClient

	for _, client := range activeClient.ActiveCLient {

		active.Inner.LoginName = append(active.Inner.LoginName, client.Inner.LoginName)

	}

	return active
}
