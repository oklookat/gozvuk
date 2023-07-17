package gozvuk

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/oklookat/gozvuk/schema"
)

var (
	_artistIds = [4]schema.ID{
		"292150",
		"681777",
		"1221500",
		"76007136",
	}

	_trackIds = [4]schema.ID{
		"132825405",
		"132584581",
		"132507771",
		"132630317",
	}

	_albumIds = [4]schema.ID{
		"24657065",
		"12899060",
		"18973023",
		"10665202",
	}

	_podcastIds = [4]schema.ID{
		"15273974",
		"30431849",
		"29605439",
		"12839134",
	}

	_episodeIds = [4]schema.ID{
		"132916901",
		"132916902",
		"132916903",
		"132916904",
	}

	_playlistIds = [4]schema.ID{
		"4607856",
		"1175767",
		"5669282",
		"6170976",
	}

	_authorIds = [4]schema.ID{
		"279969197",
		"311685330",
		"546938821",
		"1448294917",
	}
)

func TestNew(t *testing.T) {
	getClient(t)
}

func getClient(t *testing.T) *Client {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}
	cl := New(os.Getenv("ACCESS_TOKEN"))
	// cl.Http.SetLogger(loggerDefault{})
	// cl.Http.SetRateLimit(1, time.Duration(1)*time.Second)
	return cl
}

// type loggerDefault struct {
// }

// func (l loggerDefault) Debugf(msg string, args ...any) {
// 	log.Printf(msg, args...)
// }

// func (l loggerDefault) Err(msg string, err error) {
// 	if err == nil {
// 		log.Printf("%s", msg)
// 		return
// 	}
// 	log.Printf("%s. Err: %s", msg, err.Error())
// }
