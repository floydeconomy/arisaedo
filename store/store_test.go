package store

import "testing"

func TestStore(t *testing.T) {
	s := New(Options{
    Db: "localhost:4444",
    Chain: "localhost:8080"
  })
}
