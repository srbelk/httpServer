package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error parsing data", err)
	}

	rl := string(b) //convert []byte into one string

	parsedrl, err := parseRequestLine(rl) //seperate string and put them into a RequestLine{}
	if err != nil {
		return nil, err
	}

	newR := &Request{
		RequestLine: parsedrl,
	}

	return newR, nil
}

func parseRequestLine(s string) (RequestLine, error) {
	httpmsg := strings.Split(s, "\r\n")

	rlValues := strings.Split(httpmsg[0], " ")

	if len(rlValues) == 3 {
		newRL := RequestLine{
			Method:        rlValues[0],
			RequestTarget: rlValues[1],
			HttpVersion:   strings.TrimPrefix(rlValues[2], "HTTP/"),
		}

		if newRL.Method != strings.ToUpper(rlValues[0]) {
			return RequestLine{}, fmt.Errorf("Wrong casing")
		}

		if newRL.HttpVersion != "1.1" {
			return RequestLine{}, fmt.Errorf("Wrong version")
		}

		return newRL, nil

	} else {
		return RequestLine{}, fmt.Errorf("Wrong number of Request line values")
	}

}
