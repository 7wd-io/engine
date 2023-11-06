package engine

import (
	"encoding/json"
	"testing"
)

func TestStateFrom(t *testing.T) {
	g1 := `[{"meta": {"actor": ""}, "move": {"id": 1, "p1": "user2", "p2": "user1", "cards": {"1": [112, 101, 110, 120, 114, 104, 121, 103, 100, 111, 118, 108, 122, 102, 113, 115, 116, 117, 119, 106], "2": [213, 218, 206, 200, 211, 217, 201, 205, 221, 203, 209, 212, 222, 214, 220, 202, 208, 210, 215, 207], "3": [318, 313, 307, 302, 309, 403, 303, 400, 301, 306, 304, 300, 315, 311, 308, 314, 316, 401, 312, 310]}, "tokens": [4, 1, 3, 10, 6], "wonders": [10, 5, 2, 14, 12, 8, 6, 7], "randomTokens": [7, 2, 5]}}, {"meta": {"actor": "user2"}, "move": {"id": 2, "wonder": 10}}, {"meta": {"actor": "user1"}, "move": {"id": 2, "wonder": 5}}, {"meta": {"actor": "user1"}, "move": {"id": 2, "wonder": 14}}, {"meta": {"actor": "user2"}, "move": {"id": 2, "wonder": 2}}, {"meta": {"actor": "user1"}, "move": {"id": 2, "wonder": 8}}, {"meta": {"actor": "user2"}, "move": {"id": 2, "wonder": 6}}, {"meta": {"actor": "user2"}, "move": {"id": 2, "wonder": 12}}, {"meta": {"actor": "user1"}, "move": {"id": 2, "wonder": 7}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 113}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 106}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 117}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 119}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 122}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 102}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 100}}, {"meta": {"actor": "user1"}, "move": {"id": 6, "card": 116}}, {"meta": {"actor": "user2"}, "move": {"id": 6, "card": 115}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 111}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 118}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 108}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 104}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 121}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 110}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 103}}, {"meta": {"actor": "user2"}, "move": {"id": 6, "card": 114}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 120}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 112}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 101}}, {"meta": {"actor": "user2"}, "move": {"id": 7, "player": "user2"}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 215}}, {"meta": {"actor": "user2"}, "move": {"id": 3, "token": 3}}, {"meta": {"actor": "user1"}, "move": {"id": 6, "card": 202}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 212}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 207}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 208}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 222}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 205}}, {"meta": {"actor": "user1"}, "move": {"id": 6, "card": 201}}, {"meta": {"actor": "user2"}, "move": {"id": 6, "card": 210}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 220}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 214}}, {"meta": {"actor": "user2"}, "move": {"id": 3, "token": 10}}, {"meta": {"actor": "user1"}, "move": {"id": 5, "card": 203, "wonder": 5}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 213}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 218}}, {"meta": {"actor": "user2"}, "move": {"id": 5, "card": 209, "wonder": 2}}, {"meta": {"actor": "user2"}, "move": {"id": 8, "card": 106}}, {"meta": {"actor": "user1"}, "move": {"id": 6, "card": 211}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 217}}, {"meta": {"actor": "user1"}, "move": {"id": 5, "card": 221, "wonder": 8}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 206}}, {"meta": {"actor": "user2"}, "move": {"id": 6, "card": 200}}, {"meta": {"actor": "user1"}, "move": {"id": 7, "player": "user1"}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 312}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 314}}, {"meta": {"actor": "user1"}, "move": {"id": 5, "card": 300, "wonder": 14}}, {"meta": {"actor": "user1"}, "move": {"id": 12, "give": 319, "pick": 317}}, {"meta": {"actor": "user2"}, "move": {"id": 6, "card": 310}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 401}}, {"meta": {"actor": "user2"}, "move": {"id": 5, "card": 316, "wonder": 10}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 308}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 315}}, {"meta": {"actor": "user2"}, "move": {"id": 5, "card": 306, "wonder": 12}}, {"meta": {"actor": "user2"}, "move": {"id": 5, "card": 311, "wonder": 6}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 403}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 303}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 307}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 304}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 400}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 302}}, {"meta": {"actor": "user1"}, "move": {"id": 3, "token": 1}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 301}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 318}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 309}}, {"meta": {"actor": "user1"}, "move": {"id": 6, "card": 313}}]`
	g2 := `[{"meta": {"actor": ""}, "move": {"id": 1, "p1": "user2", "p2": "user1", "cards": {"1": [102, 116, 101, 106, 120, 109, 121, 122, 119, 115, 104, 110, 117, 111, 112, 103, 100, 108, 114, 113], "2": [200, 221, 206, 222, 203, 215, 220, 219, 212, 216, 211, 208, 205, 217, 204, 201, 210, 214, 218, 207], "3": [301, 305, 319, 306, 315, 308, 312, 313, 307, 310, 317, 400, 406, 318, 304, 314, 303, 401, 316, 302]}, "tokens": [10, 6, 4, 7, 9], "wonders": [12, 10, 7, 9, 11, 14, 13, 5], "randomTokens": [5, 1, 3]}}, {"meta": {"actor": "user2"}, "move": {"id": 2, "wonder": 7}}, {"meta": {"actor": "user1"}, "move": {"id": 2, "wonder": 12}}, {"meta": {"actor": "user1"}, "move": {"id": 2, "wonder": 10}}, {"meta": {"actor": "user2"}, "move": {"id": 2, "wonder": 9}}, {"meta": {"actor": "user1"}, "move": {"id": 2, "wonder": 13}}, {"meta": {"actor": "user2"}, "move": {"id": 2, "wonder": 14}}, {"meta": {"actor": "user2"}, "move": {"id": 2, "wonder": 11}}, {"meta": {"actor": "user1"}, "move": {"id": 2, "wonder": 5}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 113}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 112}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 103}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 100}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 104}}, {"meta": {"actor": "user1"}, "move": {"id": 6, "card": 115}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 109}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 114}}, {"meta": {"actor": "user2"}, "move": {"id": 6, "card": 111}}, {"meta": {"actor": "user1"}, "move": {"id": 6, "card": 108}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 117}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 110}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 122}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 119}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 121}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 106}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 120}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 101}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 102}}, {"meta": {"actor": "user1"}, "move": {"id": 6, "card": 116}}, {"meta": {"actor": "user2"}, "move": {"id": 7, "player": "user1"}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 207}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 214}}, {"meta": {"actor": "user2"}, "move": {"id": 3, "token": 9}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 204}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 218}}, {"meta": {"actor": "user1"}, "move": {"id": 5, "card": 201, "wonder": 13}}, {"meta": {"actor": "user1"}, "move": {"id": 10, "card": 206}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 208}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 210}}, {"meta": {"actor": "user2"}, "move": {"id": 6, "card": 211}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 205}}, {"meta": {"actor": "user2"}, "move": {"id": 5, "card": 219, "wonder": 7}}, {"meta": {"actor": "user2"}, "move": {"id": 11, "card": 111}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 220}}, {"meta": {"actor": "user1"}, "move": {"id": 6, "card": 200}}, {"meta": {"actor": "user2"}, "move": {"id": 6, "card": 221}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 215}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 217}}, {"meta": {"actor": "user2"}, "move": {"id": 3, "token": 10}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 216}}, {"meta": {"actor": "user1"}, "move": {"id": 3, "token": 7}}, {"meta": {"actor": "user2"}, "move": {"id": 5, "card": 212, "wonder": 14}}, {"meta": {"actor": "user2"}, "move": {"id": 12, "give": 202, "pick": 213}}, {"meta": {"actor": "user2"}, "move": {"id": 5, "card": 222, "wonder": 11}}, {"meta": {"actor": "user2"}, "move": {"id": 8, "card": 202}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 203}}, {"meta": {"actor": "user2"}, "move": {"id": 7, "player": "user2"}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 316}}, {"meta": {"actor": "user1"}, "move": {"id": 5, "card": 314, "wonder": 12}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 400}}, {"meta": {"actor": "user2"}, "move": {"id": 6, "card": 302}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 401}}, {"meta": {"actor": "user2"}, "move": {"id": 5, "card": 303, "wonder": 9}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 318}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 304}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 317}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 406}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 307}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 310}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 308}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 312}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 313}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 319}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 315}}, {"meta": {"actor": "user1"}, "move": {"id": 5, "card": 306, "wonder": 10}}, {"meta": {"actor": "user1"}, "move": {"id": 4, "card": 301}}, {"meta": {"actor": "user2"}, "move": {"id": 4, "card": 305}}]`

	var log1 Log
	var log2 Log

	if err := json.Unmarshal([]byte(g1), &log1); err != nil {
		panic("parse log1 panic")
	}

	if err := json.Unmarshal([]byte(g2), &log2); err != nil {
		panic("parse log2 panic")
	}

	var moves1, moves2 []Mutator

	for _, v := range log1 {
		moves1 = append(moves1, v.Move)
	}

	for _, v := range log2 {
		moves2 = append(moves2, v.Move)
	}

	type args struct {
		m []Mutator
	}
	tests := []struct {
		name    string
		args    args
		want    *State
		wantErr bool
	}{
		{
			name: "game 1",
			args: args{
				m: moves1,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "game 2",
			args: args{
				m: moves2,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := StateFrom(tt.args.m...)
			if (err != nil) != tt.wantErr {
				t.Errorf("StateFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("StateFrom() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
