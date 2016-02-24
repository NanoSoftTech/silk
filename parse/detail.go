package parse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
)

type Detail struct {
	Key   string
	Value *Value
}

func parseDetail(b []byte, detailregex *regexp.Regexp) (*Detail, error) {
	detail, err := getok(detailregex.FindSubmatch(b), 1)
	if err != nil {
		panic("silk: failed to parse detail: " + err.Error())
	}
	sep := bytes.IndexAny(detail, ":=")
	if sep == -1 || sep > len(detail)-1 {
		return nil, errors.New("malformed detail")
	}
	key := clean(detail[0:sep])
	return &Detail{
		Key:   string(bytes.TrimSpace(key)),
		Value: ParseValue(detail[sep+1:]),
	}, nil
}

func (d *Detail) String() string {
	valbytes, err := json.Marshal(d.Value)
	if err != nil {
		return d.Key + ": " + fmt.Sprint(d.Value)
	}
	return d.Key + ": " + string(valbytes)
}

func (d *Detail) IsCookie() bool {
	log.Println(d.Value.String())
	return false
}

func (d *Detail) Cookie() *Detail {
	return nil
}

func clean(b []byte) []byte {
	return bytes.Trim(bytes.TrimSpace(b), "`")
}
