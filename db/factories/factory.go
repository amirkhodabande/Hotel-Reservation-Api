package factories

import (
	"bytes"
	"encoding/json"

	"hotel.com/db"
)

type Factory struct {
	*db.Store
}

func New(db *db.Store) *Factory {
	return &Factory{db}
}

func transcode(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}
