package handlers

import (
	"fmt"

	"github.com/couchbase/gocb"
)

func ListAllClientFromDb(bucket *gocb.Bucket) []Client {

	myQuery := gocb.NewN1qlQuery("SELECT * FROM chatapp")
	rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)
	if err != nil {
		fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
	} // Interfaces for handling streaming return values
	var row Client
	var retValues []Client

	for rows.Next(&row) {
		retValues = append(retValues, row)
		row = Client{}
	}
	return retValues
}
func AddClient(client Client, bucket *gocb.Bucket) bool {
	if _, err := bucket.Upsert(client.Inner.LoginName, client, 0); err != nil {
		fmt.Println("Can't inser new client! The client who caused problem " + client.Inner.LoginName)
		return false
	}
	return true
}
func ListActiveCliens(activeClient Connection) OnlineClient {
	var active OnlineClient

	for _, client := range activeClient.ActiveCLient {

		active.Inner.LoginName = append(active.Inner.LoginName, client.Inner.LoginName)

	}

	return active
}
