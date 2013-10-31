package dkim

import (
	"fmt"
	"testing"
)

var utilSampleEML = `
From: "Waffles" <waffles@example.com> 
To: "Taco" <taco@example.net>
Subject: Whatup dogg
MIME-Version: 1.0
Content-Type: multipart/alternative;
  boundary="SiMpLeForLife04101984"
 
--SiMpLeForLife04101984
Content-Type: text/plain; charset=UTF-8
 
This is an email
 
--SiMpLeForLife04101984
Content-Type: text/html; charset=UTF-8

<hi>This is an email</h1>
`

func TestReadEML(t *testing.T) {
	msg, err := readEML([]byte(utilSampleEML))
	if err != nil {
		t.Fatal("error not nil", err)
	}
	if len(msg.Header) == 0 {
		t.Fatal("wrong header length", len(header))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(yourReader)
	body := buf.String()

	if len(body) == 0 {
		t.Fatal("wrong body length", len(body))
	}
}
