package gozvuk

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestNew(t *testing.T) {
	cl := getClient(t)
	cl.Profile()
}

func getClient(t *testing.T) *Client {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}
	cl, err := New(os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		t.Fatal(err)
	}
	return cl
}
