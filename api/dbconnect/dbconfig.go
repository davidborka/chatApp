package dbconnect

import (
	"github.com/couchbase/gocb"
)

func DatabaseConnectionUsers() *gocb.Bucket {
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ := cluster.OpenBucket("chatappUsers", "")
	return bucket
}
func DatabaseConnectionMessage() *gocb.Bucket {
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ := cluster.OpenBucket("chatAppMessages", "")
	return bucket
}
