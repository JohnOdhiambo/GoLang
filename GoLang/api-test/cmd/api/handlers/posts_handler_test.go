package handlers

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

// Pass the EndToEndSuite start
func (s *EndToEndSuite) TestPostHandler() {
	//Create a client and test
	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}
func (s *EndToEndSuite) TestPostNoResult() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1223/post/5440")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := ioutil.ReadAll(r.Body)
	s.JSONEq(`{"status:"ok", "data":[]}`, string(b))
}
