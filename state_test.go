package engine

import (
	"testing"
)

func TestStateFrom(t *testing.T) {
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
			name: "game 11",
			args: args{
				m: []Mutator{
					PrepareMove{
						Id: MovePrepare,
						P1: "user1",
						P2: "user2",
						Wonders: WonderList{
							TheHangingGardens,
							TheTempleOfArtemis,
							TheColossus,
							Messe,
							ThePyramids,
							StatueOfLiberty,
							TheMausoleum,
							TheSphinx,
						},
						Tokens: TokenList{
							Economy,
							Agriculture,
							Philosophy,
							Theology,
							Law,
						},
						RandomTokens: TokenList{
							Urbanism,
							Strategy,
							Masonry,
						},
						Cards: map[Age]CardList{
							AgeI: {
								Palisade,
								Theater,
								Tavern,
								Stable,
								Altar,
								Workshop,
								ClayReserve,
								GlassWorks,
								LoggingCamp,
								LumberYard,
								Baths,
								Quarry,
								ClayPit,
								ClayPool,
								Scriptorium,
								Garrison,
								StonePit,
								WoodReserve,
								Pharmacist,
								StoneReserve,
							},
							AgeII: {
								Dispensary,
								CustomHouse,
								CourtHouse,
								Caravansery,
								GlassBlower,
								BrickYard,
								School,
								Laboratory,
								Aqueduct,
								ArcheryRange,
								ParadeGround,
								Brewery,
								Statue,
								HorseBreeders,
								ShelfQuarry,
								Library,
								Walls,
								SawMill,
								Barracks,
								DryingRoom,
							},
							AgeIII: {
								Port,
								Academy,
								Obelisk,
								Observatory,
								Fortifications,
								Palace,
								Senate,
								Armory,
								MagistratesGuild,
								MerchantsGuild,
								SiegeWorkshop,
								ChamberOfCommerce,
								Arsenal,
								Pretorium,
								Arena,
								Lighthouse,
								Gardens,
								Pantheon,
								MoneyLendersGuild,
								TownHall,
							},
						},
					},
					NewMovePickWonder(TheTempleOfArtemis),                   //1
					NewMovePickWonder(TheHangingGardens),                    //2
					NewMovePickWonder(TheColossus),                          //3
					NewMovePickWonder(Messe),                                //4
					NewMovePickWonder(TheSphinx),                            //5
					NewMovePickWonder(StatueOfLiberty),                      //6
					NewMovePickWonder(TheMausoleum),                         //7
					NewMovePickWonder(ThePyramids),                          //8
					NewMoveConstructCard(WoodReserve),                       //9
					NewMoveConstructCard(StoneReserve),                      //10
					NewMoveConstructCard(Scriptorium),                       //11
					NewMoveConstructCard(StonePit),                          //12
					NewMoveConstructCard(Quarry),                            //13
					NewMoveDiscardCard(Garrison),                            //14
					NewMoveConstructCard(Pharmacist),                        //15
					NewMoveConstructCard(ClayPool),                          //16
					NewMoveConstructCard(LumberYard),                        //17
					NewMoveConstructCard(Baths),                             //18
					NewMoveDiscardCard(ClayPit),                             //19
					NewMoveConstructCard(LoggingCamp),                       //20
					NewMoveConstructCard(GlassWorks),                        //21
					NewMoveConstructCard(Altar),                             //22
					NewMoveConstructCard(Workshop),                          //23
					NewMoveDiscardCard(ClayReserve),                         //24
					NewMoveConstructCard(Tavern),                            //25
					NewMoveConstructCard(Stable),                            //26
					NewMoveConstructCard(Theater),                           //27
					NewMoveConstructCard(Palisade),                          //28
					NewMoveSelectWhoBeginsTheNextAge("user1"),               //29
					NewMoveConstructCard(DryingRoom),                        //30
					NewMoveConstructCard(SawMill),                           //31
					NewMoveConstructCard(ShelfQuarry),                       //32
					NewMoveDiscardCard(ParadeGround),                        //33
					NewMoveConstructCard(BrickYard),                         //34
					NewMoveConstructCard(Barracks),                          //35
					NewMoveConstructCard(Library),                           //36
					NewMovePickBoardToken(Theology),                         //37
					NewMoveConstructCard(Walls),                             //38
					NewMoveConstructCard(Brewery),                           //39
					NewMoveDiscardCard(HorseBreeders),                       //40
					NewMoveConstructWonder(Messe, Statue),                   //41
					NewMovePickTopLineCard(Dispensary),                      //42
					NewMovePickBoardToken(Economy),                          //43
					NewMoveConstructCard(Laboratory),                        //44
					NewMovePickBoardToken(Agriculture),                      //45
					NewMoveConstructCard(ArcheryRange),                      //46
					NewMoveConstructCard(Aqueduct),                          //47
					NewMoveConstructCard(GlassBlower),                       //48
					NewMoveConstructCard(School),                            //49
					NewMoveDiscardCard(CourtHouse),                          //50
					NewMoveConstructCard(Caravansery),                       //51
					NewMoveConstructCard(CustomHouse),                       //52
					NewMoveSelectWhoBeginsTheNextAge("user1"),               //53
					NewMoveConstructWonder(TheMausoleum, MoneyLendersGuild), //54
					NewMovePickDiscardedCard(ParadeGround),                  //55
					NewMoveConstructCard(Lighthouse),                        //56
					NewMoveConstructCard(ChamberOfCommerce),                 //57
					NewMoveConstructCard(TownHall),                          //58
					NewMoveConstructWonder(ThePyramids, Gardens),            //59
					NewMoveConstructCard(Arsenal),                           //60
					NewMoveDiscardCard(Pantheon),                            //61
					NewMoveDiscardCard(Pretorium),                           //62
					NewMoveConstructCard(MerchantsGuild),                    //63
					NewMoveConstructWonder(StatueOfLiberty, Senate),         //64
					NewMovePickReturnedCards(Study, Circus),                 //65
					NewMoveConstructWonder(TheTempleOfArtemis, Palace),      //66
					NewMoveConstructCard(Obelisk),                           //67
					NewMoveConstructCard(Arena),                             //68
					NewMoveConstructCard(SiegeWorkshop),                     //69
					NewMoveConstructCard(MagistratesGuild),                  //70
					NewMoveConstructCard(Armory),                            //71
					NewMoveConstructCard(Observatory),                       //72
					NewMoveConstructCard(Fortifications),                    //73
					NewMoveConstructCard(Port),                              //74
					NewMoveConstructCard(Academy),                           //75
					NewMovePickBoardToken(Philosophy),                       //76
				},
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
