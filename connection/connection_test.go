package connection

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-sdk/config"
	"github.com/stretchr/testify/suite"
)

// ConnectionTestSuite : Test suite for connection
type ConnectionTestSuite struct {
	url string
	suite.Suite
	Connection *Conn
}

func testhandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(201)
		w.Write(data)
	default:
		w.Write([]byte(`{"status":"ok"}`))
	}
}

// SetupTest : sets up test suite
func (suite *ConnectionTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/test", testhandler)
	server := httptest.NewServer(mux)

	suite.Connection = New(config.New(server.URL))
}

func (suite *ConnectionTestSuite) TestGet() {
	resp, err := suite.Connection.Get("/api/test")
	body, rerr := ioutil.ReadAll(resp.Body)

	suite.Equal(err, nil)
	suite.Equal(rerr, nil)
	suite.Equal(resp.StatusCode, 200)
	suite.Equal(body, []byte(`{"status":"ok"}`))
}

func (suite *ConnectionTestSuite) TestPost() {
	data := []byte(`{"id":"test"}`)

	resp, err := suite.Connection.Post("/api/test", "application/json", data)
	body, rerr := ioutil.ReadAll(resp.Body)

	suite.Equal(err, nil)
	suite.Equal(rerr, nil)
	suite.Equal(resp.StatusCode, 201)
	suite.Equal(body, data)
}

// TestConnectionTestSuite : Test suite for connection
func TestConnectionTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectionTestSuite))
}
