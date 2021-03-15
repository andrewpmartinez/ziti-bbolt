package main

import (
	"go.etcd.io/bbolt"
	"log"
)

func main() {

	db, err := bbolt.Open("./ctrl.db", 0666, nil)
	if err != nil {
		log.Fatalf("could not open ./ctrl.db: %v", err)
	}

	defer db.Close()

	_ = db.Update(func(tx *bbolt.Tx) error {

		root := tx.Bucket([]byte("ziti"))

		if root == nil {
			log.Fatal("root ziti bucket not found")
			return nil
		}

		if err := root.DeleteBucket([]byte("apiSessions")); err != nil {
			log.Default().Printf("could not delete apiSessions: %v", err)
		} else {
			log.Default().Print("done apiSessions")
		}

		if err := root.DeleteBucket([]byte("sessions")); err != nil {
			log.Default().Printf("could not delete sessions: %v", err)
		} else {
			log.Default().Print("done sessions")
		}

		return nil
	})
}
