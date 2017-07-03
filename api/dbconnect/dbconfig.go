package dbconnect

import (
	"github.com/couchbase/gocb"
)

func DatabaseConnection() *gocb.Bucket {
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ := cluster.OpenBucket("chatapp", "")
	return bucket
}
