package transaction

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"strconv"
)

type Stock struct {
	Id            int `json:"id"`
	idStr         string
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Value         float32 `json:"value"`
	stockBytes    []byte
	stockBytesStr string
}

func (s *Stock) AsBytes() []byte {

	if s.stockBytes == nil {
		b, err := json.Marshal(s)

		if err != nil {
			log.Error().Msgf("Error trying to mashall %+v to json bytes array: %s", s, err)
		}

		s.stockBytes = b
	}

	return s.stockBytes
}

func (s *Stock) AsString() string {

	if s.stockBytesStr == "" {
		str := string(s.AsBytes()[:])
		s.stockBytesStr = str
	}

	return s.stockBytesStr
}

func (s *Stock) IdStr() string {

	if s.idStr == "" {
		s.idStr = strconv.FormatInt(int64(s.Id), 10)
	}

	return s.idStr
}