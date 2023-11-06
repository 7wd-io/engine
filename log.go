package engine

import (
	"encoding/json"
	"fmt"
	"log"
)

type Log []LogRecord

type LogRecord struct {
	Move Mutator  `json:"move"`
	Meta MoveMeta `json:"meta"`
}

type MoveMeta struct {
	Actor Nickname `json:"actor"`
}

func (dst *Log) UnmarshalJSON(bytes []byte) error {
	var messages []*json.RawMessage

	if err := json.Unmarshal(bytes, &messages); err != nil {
		panic("moves unmarshal fail")
	}

	var record struct {
		Move map[string]interface{} `json:"move"`
		Meta MoveMeta               `json:"meta"`
	}

	out := make(Log, len(messages))

	for index, message := range messages {
		if err := json.Unmarshal(*message, &record); err != nil {
			log.Fatalln(err)
		}

		//var rawMove b.Buffer
		//
		//if err := gob.NewEncoder(&rawMove).Encode(record.Move); err != nil {
		//	log.Fatalln(err)
		//}

		//fmt.Println(string(rawMove.Bytes()))

		rawMove, err := json.Marshal(record.Move)

		if err != nil {
			log.Fatalln(err)
		}
		//
		//moveBytes := []byte()

		switch MoveId(record.Move["id"].(float64)) {
		case MovePrepare:
			var m1 prepareMove

			if err := json.Unmarshal(rawMove, &m1); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m1,
				Meta: record.Meta,
			}
		case MovePickWonder:
			var m2 pickWonderMove

			if err := json.Unmarshal(rawMove, &m2); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m2,
				Meta: record.Meta,
			}
		case MovePickBoardToken:
			var m3 pickBoardTokenMove

			if err := json.Unmarshal(rawMove, &m3); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m3,
				Meta: record.Meta,
			}
		case MoveConstructCard:
			var m4 constructCardMove

			if err := json.Unmarshal(rawMove, &m4); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m4,
				Meta: record.Meta,
			}
		case MoveConstructWonder:
			var m5 constructWonderMove

			if err := json.Unmarshal(rawMove, &m5); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m5,
				Meta: record.Meta,
			}
		case MoveDiscardCard:
			var m6 discardCardMove

			if err := json.Unmarshal(rawMove, &m6); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m6,
				Meta: record.Meta,
			}
		case MoveSelectWhoBeginsTheNextAge:
			var m7 selectWhoBeginsTheNextAgeMove

			if err := json.Unmarshal(rawMove, &m7); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m7,
				Meta: record.Meta,
			}
		case MoveBurnCard:
			var m8 burnCardMove

			if err := json.Unmarshal(rawMove, &m8); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m8,
				Meta: record.Meta,
			}
		case MovePickRandomToken:
			var m9 pickRandomTokenMove

			if err := json.Unmarshal(rawMove, &m9); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m9,
				Meta: record.Meta,
			}
		case MovePickTopLineCard:
			var m10 pickTopLineCardMove

			if err := json.Unmarshal(rawMove, &m10); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m10,
				Meta: record.Meta,
			}
		case MovePickDiscardedCard:
			var m11 pickDiscardedCardMove

			if err := json.Unmarshal(rawMove, &m11); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m11,
				Meta: record.Meta,
			}
		case MovePickReturnedCards:
			var m12 pickReturnedCardsMove

			if err := json.Unmarshal(rawMove, &m12); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m12,
				Meta: record.Meta,
			}
		case MoveOver:
			var m13 overMove

			if err := json.Unmarshal(rawMove, &m13); err != nil {
				panic(fmt.Errorf("moves unmarshal fail: %w", err))
			}

			out[index] = LogRecord{
				Move: m13,
				Meta: record.Meta,
			}
		default:
			panic("unknown move")
		}
	}

	*dst = out

	return nil
}
