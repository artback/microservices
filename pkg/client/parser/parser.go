package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"micro_services/api/v1/port"
)

type Result struct {
	Port *port.Port
	Err  error
}

func ReadPorts(reader io.Reader, resChan chan Result) {
	dec := json.NewDecoder(reader)
	t, err := dec.Token()
	if err != nil {
		resChan <- Result{nil, err}
		return
	}
	for dec.More() {
		t, err = dec.Token()
		if err == io.EOF {
			resChan <- Result{nil, err}
			break
		}
		if err != nil {
			resChan <- Result{nil, err}
			return
		}
		var p *port.Port
		err = dec.Decode(&p)
		id := fmt.Sprint(t)
		p.ID = id
		resChan <- Result{
			p,
			err,
		}
	}
	close(resChan)
}
