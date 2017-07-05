package handlers

import (
	"github.com/couchbase/gocb"
	"github.com/davidborka/chatApp/api/model"
)

/*
func ListAllMessage(loginname string, bucket *gocb.Bucket) []interface{} {
	var query string
	query = `SELECT * FROM chatapp  WHERE ANY message IN chatapp SATISFIES message.fromloginname="` + loginname + `" OR message.tologinname="` + loginname + `" END;`
	fmt.Println(query)
	myQuery := gocb.NewN1qlQuery(query)
	rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)
	if err != nil {
		fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
	} // Interfaces for handling streaming return values
	var retValues []interface{}
	var row interface{}
	for rows.Next(&row) {
		retValues = append(retValues, row)

	}
	byte, _ := json.Marshal(retValues)

	fmt.Println(string(byte))

	return retValues
}*/
func AddMessageClient(id string, messageToInsert *model.Message, bucket *gocb.Bucket) {
	var nexus model.Nexus

	_, err := bucket.Get(id, &nexus)
	if err != nil {
		bucket.Upsert(id, nexus, 0)
		bucket.Get(id, &nexus)
	}
	nexus.Messages = append(nexus.Messages, *messageToInsert)
	bucket.Upsert(id, nexus, 0)

}
