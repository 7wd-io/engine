package engine

var R = registry{
	Wonders: map[WonderId]wonder{
		TheAppianWay: {
			Id: TheAppianWay,
			Cost: cost{
				Resources: resourceMap{
					papyrus: 1,
					clay:    2,
					stone:   2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectCoins(3),
					newEffectFine(3),
					newEffectPlayAgain(),
					newEffectPoints(3),
				},
			},
		},
		CircusMaximus: {
			Id: CircusMaximus,
			Cost: cost{
				Resources: resourceMap{
					glass: 1,
					wood:  1,
					stone: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectBurnCard(groupManufacturedGoods),
					newEffectMilitaryWithConfig(militaryEffectConfig{
						power:            1,
						strategyDisabled: true,
					}),
					newEffectPoints(3),
				},
			},
		},
		TheColossus: {
			Id: TheColossus,
			Cost: cost{
				Resources: resourceMap{
					glass: 1,
					clay:  3,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitaryWithConfig(militaryEffectConfig{
						power:            2,
						strategyDisabled: true,
					}),
					newEffectPoints(3),
				},
			},
		},
		TheGreatLibrary: {
			Id: TheGreatLibrary,
			Cost: cost{
				Resources: resourceMap{
					papyrus: 1,
					glass:   1,
					wood:    3,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPickRandomToken(),
					newEffectPoints(4),
				},
			},
		},
		TheGreatLighthouse: {
			Id: TheGreatLighthouse,
			Cost: cost{
				Resources: resourceMap{
					papyrus: 2,
					stone:   1,
					wood:    1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectDiscounter(dcGlobal, 1, clay, wood, stone),
					newEffectPoints(4),
				},
			},
		},
		TheHangingGardens: {
			Id: TheHangingGardens,
			Cost: cost{
				Resources: resourceMap{
					papyrus: 1,
					glass:   1,
					wood:    2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectCoins(6),
					newEffectPlayAgain(),
					newEffectPoints(3),
				},
			},
		},
		TheMausoleum: {
			Id: TheMausoleum,
			Cost: cost{
				Resources: resourceMap{
					papyrus: 1,
					glass:   2,
					clay:    2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPickDiscardedCard(),
					newEffectPoints(2),
				},
			},
		},
		Piraeus: {
			Id: Piraeus,
			Cost: cost{
				Resources: resourceMap{
					clay:  1,
					stone: 1,
					wood:  2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectDiscounter(dcGlobal, 1, glass, papyrus),
					newEffectPlayAgain(),
					newEffectPoints(2),
				},
			},
		},
		ThePyramids: {
			Id: ThePyramids,
			Cost: cost{
				Resources: resourceMap{
					papyrus: 1,
					stone:   3,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(9),
				},
			},
		},
		TheSphinx: {
			Id: TheSphinx,
			Cost: cost{
				Resources: resourceMap{
					glass: 2,
					clay:  1,
					stone: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPlayAgain(),
					newEffectPoints(6),
				},
			},
		},
		TheStatueOfZeus: {
			Id: TheStatueOfZeus,
			Cost: cost{
				Resources: resourceMap{
					papyrus: 2,
					clay:    1,
					wood:    1,
					stone:   1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectBurnCard(groupRawMaterials),
					newEffectMilitaryWithConfig(militaryEffectConfig{
						power:            1,
						strategyDisabled: true,
					}),
					newEffectPoints(3),
				},
			},
		},
		TheTempleOfArtemis: {
			Id: TheTempleOfArtemis,
			Cost: cost{
				Resources: resourceMap{
					papyrus: 1,
					glass:   1,
					stone:   1,
					wood:    1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectCoins(12),
					newEffectPlayAgain(),
				},
			},
		},
		Messe: {
			Id: Messe,
			Cost: cost{
				Resources: resourceMap{
					glass:   1,
					papyrus: 1,
					wood:    1,
					clay:    2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPickTopLineCard(),
					newEffectPoints(2),
				},
			},
		},
		StatueOfLiberty: {
			Id: StatueOfLiberty,
			Cost: cost{
				Resources: resourceMap{
					glass:   1,
					papyrus: 1,
					clay:    1,
					stone:   1,
					wood:    1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPickReturnedCards(),
					newEffectPoints(5),
				},
			},
		},
	},
	Cards: map[CardId]card{
		LumberYard: {
			Id:    LumberYard,
			Age:   ageI,
			Group: groupRawMaterials,
			unit: unit{
				Effects: []interface{}{
					newEffectResource(wood, 1),
				},
			},
		},
		LoggingCamp: {
			Id:    LoggingCamp,
			Age:   ageI,
			Group: groupRawMaterials,
			Cost: cost{
				Coins: 1,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectResource(wood, 1),
				},
			},
		},
		ClayPool: {
			Id:    ClayPool,
			Age:   ageI,
			Group: groupRawMaterials,
			unit: unit{
				Effects: []interface{}{
					newEffectResource(clay, 1),
				},
			},
		},
		ClayPit: {
			Id:    ClayPit,
			Age:   ageI,
			Group: groupRawMaterials,
			Cost: cost{
				Coins: 1,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectResource(clay, 1),
				},
			},
		},
		Quarry: {
			Id:    Quarry,
			Age:   ageI,
			Group: groupRawMaterials,
			unit: unit{
				Effects: []interface{}{
					newEffectResource(stone, 1),
				},
			},
		},
		StonePit: {
			Id:    StonePit,
			Age:   ageI,
			Group: groupRawMaterials,
			Cost: cost{
				Coins: 1,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectResource(stone, 1),
				},
			},
		},
		GlassWorks: {
			Id:    GlassWorks,
			Age:   ageI,
			Group: groupManufacturedGoods,
			Cost: cost{
				Coins: 1,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectResource(glass, 1),
				},
			},
		},
		Press: {
			Id:    Press,
			Age:   ageI,
			Group: groupManufacturedGoods,
			Cost: cost{
				Coins: 1,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectResource(papyrus, 1),
				},
			},
		},
		GuardTower: {
			Id:    GuardTower,
			Age:   ageI,
			Group: groupMilitary,
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(1),
				},
			},
		},
		Workshop: {
			Id:    Workshop,
			Age:   ageI,
			Group: groupScientific,
			Cost: cost{
				Resources: resourceMap{
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(1),
					newEffectScience(symbolCompass),
				},
			},
		},
		Apothecary: {
			Id:    Apothecary,
			Age:   ageI,
			Group: groupScientific,
			Cost: cost{
				Resources: resourceMap{
					glass: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(1),
					newEffectScience(symbolWheel),
				},
			},
		},
		StoneReserve: {
			Id:    StoneReserve,
			Age:   ageI,
			Group: groupCommercial,
			Cost: cost{
				Coins: 3,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectFixedCost(stone),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		ClayReserve: {
			Id:    ClayReserve,
			Age:   ageI,
			Group: groupCommercial,
			Cost: cost{
				Coins: 3,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectFixedCost(clay),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		WoodReserve: {
			Id:    WoodReserve,
			Age:   ageI,
			Group: groupCommercial,
			Cost: cost{
				Coins: 3,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectFixedCost(wood),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		Stable: {
			Id:    Stable,
			Age:   ageI,
			Group: groupMilitary,
			Cost: cost{
				Resources: resourceMap{
					wood: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(1),
					newEffectChain(HorseBreeders),
				},
			},
		},
		Garrison: {
			Id:    Garrison,
			Age:   ageI,
			Group: groupMilitary,
			Cost: cost{
				Resources: resourceMap{
					clay: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(1),
					newEffectChain(Barracks),
				},
			},
		},
		Palisade: {
			Id:    Palisade,
			Age:   ageI,
			Group: groupMilitary,
			Cost: cost{
				Coins: 2,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(1),
					newEffectChain(Fortifications),
				},
			},
		},
		Scriptorium: {
			Id:    Scriptorium,
			Age:   ageI,
			Group: groupScientific,
			Cost: cost{
				Coins: 2,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectScience(symbolWriting),
					newEffectChain(Library),
				},
			},
		},
		Pharmacist: {
			Id:    Pharmacist,
			Age:   ageI,
			Group: groupScientific,
			Cost: cost{
				Coins: 2,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectScience(symbolMortar),
					newEffectChain(Dispensary),
				},
			},
		},
		Theater: {
			Id:    Theater,
			Age:   ageI,
			Group: groupCivilian,
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(3),
					newEffectChain(Statue),
				},
			},
		},
		Altar: {
			Id:    Altar,
			Age:   ageI,
			Group: groupCivilian,
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(3),
					newEffectChain(Temple),
				},
			},
		},
		Baths: {
			Id:    Baths,
			Age:   ageI,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					stone: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(3),
					newEffectChain(Aqueduct),
				},
			},
		},
		Tavern: {
			Id:    Tavern,
			Age:   ageI,
			Group: groupCommercial,
			unit: unit{
				Effects: []interface{}{
					newEffectCoins(4),
					newEffectChain(Lighthouse),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		SawMill: {
			Id:    SawMill,
			Age:   ageII,
			Group: groupRawMaterials,
			Cost: cost{
				Coins: 2,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectResource(wood, 2),
				},
			},
		},
		BrickYard: {
			Id:    BrickYard,
			Age:   ageII,
			Group: groupRawMaterials,
			Cost: cost{
				Coins: 2,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectResource(clay, 2),
				},
			},
		},
		ShelfQuarry: {
			Id:    ShelfQuarry,
			Age:   ageII,
			Group: groupRawMaterials,
			Cost: cost{
				Coins: 2,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectResource(stone, 2),
				},
			},
		},
		GlassBlower: {
			Id:    GlassBlower,
			Age:   ageII,
			Group: groupManufacturedGoods,
			unit: unit{
				Effects: []interface{}{
					newEffectResource(glass, 1),
				},
			},
		},
		DryingRoom: {
			Id:    DryingRoom,
			Age:   ageII,
			Group: groupManufacturedGoods,
			unit: unit{
				Effects: []interface{}{
					newEffectResource(papyrus, 1),
				},
			},
		},
		Walls: {
			Id:    Walls,
			Age:   ageII,
			Group: groupMilitary,
			Cost: cost{
				Resources: resourceMap{
					stone: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(2),
				},
			},
		},
		Forum: {
			Id:    Forum,
			Age:   ageII,
			Group: groupCommercial,
			Cost: cost{
				Coins: 3,
				Resources: resourceMap{
					clay: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectDiscounter(dcGlobal, 1, glass, papyrus),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		Caravansery: {
			Id:    Caravansery,
			Age:   ageII,
			Group: groupCommercial,
			Cost: cost{
				Coins: 2,
				Resources: resourceMap{
					glass:   1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectDiscounter(dcGlobal, 1, wood, clay, stone),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		CustomHouse: {
			Id:    CustomHouse,
			Age:   ageII,
			Group: groupCommercial,
			Cost: cost{
				Coins: 4,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectFixedCost(papyrus, glass),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		CourtHouse: {
			Id:    CourtHouse,
			Age:   ageII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					wood:  2,
					glass: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(5),
				},
			},
		},
		HorseBreeders: {
			Id:    HorseBreeders,
			Age:   ageII,
			Group: groupMilitary,
			Cost: cost{
				Resources: resourceMap{
					clay: 1,
					wood: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(1),
				},
			},
		},
		Barracks: {
			Id:    Barracks,
			Age:   ageII,
			Group: groupMilitary,
			Cost: cost{
				Coins: 3,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(1),
				},
			},
		},
		ArcheryRange: {
			Id:    ArcheryRange,
			Age:   ageII,
			Group: groupMilitary,
			Cost: cost{
				Resources: resourceMap{
					stone:   1,
					wood:    1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(2),
					newEffectChain(SiegeWorkshop),
				},
			},
		},
		ParadeGround: {
			Id:    ParadeGround,
			Age:   ageII,
			Group: groupMilitary,
			Cost: cost{
				Resources: resourceMap{
					clay:  2,
					glass: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(2),
					newEffectChain(Circus),
				},
			},
		},
		Library: {
			Id:    Library,
			Age:   ageII,
			Group: groupScientific,
			Cost: cost{
				Resources: resourceMap{
					stone: 1,
					wood:  1,
					glass: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(2),
					newEffectScience(symbolWriting),
				},
			},
		},
		Dispensary: {
			Id:    Dispensary,
			Age:   ageII,
			Group: groupScientific,
			Cost: cost{
				Resources: resourceMap{
					clay:  2,
					stone: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(2),
					newEffectScience(symbolMortar),
				},
			},
		},
		School: {
			Id:    School,
			Age:   ageII,
			Group: groupScientific,
			Cost: cost{
				Resources: resourceMap{
					wood:    1,
					papyrus: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(1),
					newEffectScience(symbolWheel),
					newEffectChain(University),
				},
			},
		},
		Laboratory: {
			Id:    Laboratory,
			Age:   ageII,
			Group: groupScientific,
			Cost: cost{
				Resources: resourceMap{
					wood:  1,
					glass: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(1),
					newEffectScience(symbolCompass),
					newEffectChain(Observatory),
				},
			},
		},
		Statue: {
			Id:    Statue,
			Age:   ageII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					clay: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(4),
					newEffectChain(Gardens),
				},
			},
		},
		Temple: {
			Id:    Temple,
			Age:   ageII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					wood:    1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(4),
					newEffectChain(Pantheon),
				},
			},
		},
		Aqueduct: {
			Id:    Aqueduct,
			Age:   ageII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					stone: 3,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(5),
				},
			},
		},
		Rostrum: {
			Id:    Rostrum,
			Age:   ageII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					stone: 1,
					wood:  1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(4),
					newEffectChain(Senate),
				},
			},
		},
		Brewery: {
			Id:    Brewery,
			Age:   ageII,
			Group: groupCommercial,
			unit: unit{
				Effects: []interface{}{
					newEffectCoins(6),
					newEffectChain(Arena),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		Arsenal: {
			Id:    Arsenal,
			Age:   ageIII,
			Group: groupMilitary,
			Cost: cost{
				Resources: resourceMap{
					clay: 3,
					wood: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(3),
				},
			},
		},
		Pretorium: {
			Id:    Pretorium,
			Age:   ageIII,
			Group: groupMilitary,
			Cost: cost{
				Coins: 8,
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(3),
				},
			},
		},
		Academy: {
			Id:    Academy,
			Age:   ageIII,
			Group: groupScientific,
			Cost: cost{
				Resources: resourceMap{
					stone: 1,
					wood:  1,
					glass: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(3),
					newEffectScience(symbolSundial),
				},
			},
		},
		Study: {
			Id:    Study,
			Age:   ageIII,
			Group: groupScientific,
			Cost: cost{
				Resources: resourceMap{
					wood:    2,
					glass:   1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(3),
					newEffectScience(symbolSundial),
				},
			},
		},
		ChamberOfCommerce: {
			Id:    ChamberOfCommerce,
			Age:   ageIII,
			Group: groupCommercial,
			Cost: cost{
				Resources: resourceMap{
					papyrus: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(3),
					newEffectCoinsFor(bonusManufacturedGoods, 3),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		Port: {
			Id:    Port,
			Age:   ageIII,
			Group: groupCommercial,
			Cost: cost{
				Resources: resourceMap{
					wood:    1,
					glass:   1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(3),
					newEffectCoinsFor(bonusRawMaterials, 2),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		Armory: {
			Id:    Armory,
			Age:   ageIII,
			Group: groupCommercial,
			Cost: cost{
				Resources: resourceMap{
					stone: 2,
					glass: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(3),
					newEffectCoinsFor(bonusMilitary, 1),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		Palace: {
			Id:    Palace,
			Age:   ageIII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					clay:  1,
					stone: 1,
					wood:  1,
					glass: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(7),
				},
			},
		},
		TownHall: {
			Id:    TownHall,
			Age:   ageIII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					stone: 3,
					wood:  2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(7),
				},
			},
		},
		Obelisk: {
			Id:    Obelisk,
			Age:   ageIII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					stone: 2,
					glass: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(5),
				},
			},
		},
		Fortifications: {
			Id:    Fortifications,
			Age:   ageIII,
			Group: groupMilitary,
			Cost: cost{
				Resources: resourceMap{
					stone:   2,
					clay:    1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(2),
				},
			},
		},
		SiegeWorkshop: {
			Id:    SiegeWorkshop,
			Age:   ageIII,
			Group: groupMilitary,
			Cost: cost{
				Resources: resourceMap{
					wood:  3,
					glass: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(2),
				},
			},
		},
		Circus: {
			Id:    Circus,
			Age:   ageIII,
			Group: groupMilitary,
			Cost: cost{
				Resources: resourceMap{
					clay:  2,
					stone: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectMilitary(2),
				},
			},
		},
		University: {
			Id:    University,
			Age:   ageIII,
			Group: groupScientific,
			Cost: cost{
				Resources: resourceMap{
					clay:    1,
					glass:   1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(2),
					newEffectScience(symbolAstrology),
				},
			},
		},
		Observatory: {
			Id:    Observatory,
			Age:   ageIII,
			Group: groupScientific,
			Cost: cost{
				Resources: resourceMap{
					stone:   1,
					papyrus: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(2),
					newEffectScience(symbolAstrology),
				},
			},
		},
		Gardens: {
			Id:    Gardens,
			Age:   ageIII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					clay: 2,
					wood: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(6),
				},
			},
		},
		Pantheon: {
			Id:    Pantheon,
			Age:   ageIII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					clay:    1,
					wood:    1,
					papyrus: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(6),
				},
			},
		},
		Senate: {
			Id:    Senate,
			Age:   ageIII,
			Group: groupCivilian,
			Cost: cost{
				Resources: resourceMap{
					clay:    2,
					stone:   1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(5),
				},
			},
		},
		Lighthouse: {
			Id:    Lighthouse,
			Age:   ageIII,
			Group: groupCommercial,
			Cost: cost{
				Resources: resourceMap{
					clay:  2,
					glass: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(3),
					newEffectCoinsFor(bonusCommercial, 1),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		Arena: {
			Id:    Arena,
			Age:   ageIII,
			Group: groupCommercial,
			Cost: cost{
				Resources: resourceMap{
					clay:  1,
					stone: 1,
					wood:  1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(3),
					newEffectCoinsFor(bonusWonder, 2),
					newEffectDiscardRewardAdjuster(),
				},
			},
		},
		MerchantsGuild: {
			Id:    MerchantsGuild,
			Age:   ageIII,
			Group: groupGuild,
			Cost: cost{
				Resources: resourceMap{
					clay:    1,
					wood:    1,
					glass:   1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectGuild(bonusCommercial, 1, 1),
				},
			},
		},
		ShipOwnersGuild: {
			Id:    ShipOwnersGuild,
			Age:   ageIII,
			Group: groupGuild,
			Cost: cost{
				Resources: resourceMap{
					clay:    1,
					stone:   1,
					glass:   1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectGuild(bonusResources, 1, 1),
				},
			},
		},
		BuildersGuild: {
			Id:    BuildersGuild,
			Age:   ageIII,
			Group: groupGuild,
			Cost: cost{
				Resources: resourceMap{
					stone: 2,
					clay:  1,
					wood:  1,
					glass: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectGuild(bonusWonder, 2, 0),
				},
			},
		},
		MagistratesGuild: {
			Id:    MagistratesGuild,
			Age:   ageIII,
			Group: groupGuild,
			Cost: cost{
				Resources: resourceMap{
					wood:    2,
					clay:    1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectGuild(bonusCivilian, 1, 1),
				},
			},
		},
		ScientistsGuild: {
			Id:    ScientistsGuild,
			Age:   ageIII,
			Group: groupGuild,
			Cost: cost{
				Resources: resourceMap{
					clay: 2,
					wood: 2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectGuild(bonusScience, 1, 1),
				},
			},
		},
		MoneyLendersGuild: {
			Id:    MoneyLendersGuild,
			Age:   ageIII,
			Group: groupGuild,
			Cost: cost{
				Resources: resourceMap{
					stone: 2,
					wood:  2,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectGuild(bonusCoin, 1, 0),
				},
			},
		},
		TacticiansGuild: {
			Id:    TacticiansGuild,
			Age:   ageIII,
			Group: groupGuild,
			Cost: cost{
				Resources: resourceMap{
					stone:   2,
					clay:    1,
					papyrus: 1,
				},
			},
			unit: unit{
				Effects: []interface{}{
					newEffectGuild(bonusMilitary, 1, 1),
				},
			},
		},
	},
	tokens: map[TokenId]*token{
		Agriculture: {
			unit: unit{
				Effects: []interface{}{
					newEffectCoins(6),
					newEffectPoints(4),
				},
			},
		},
		Architecture: {
			unit: unit{
				Effects: []interface{}{
					newEffectDiscounter(
						dcWonders,
						2,
						clay, wood, stone, glass, papyrus,
					),
				},
			},
		},
		Economy: {
			unit: unit{
				Effects: []interface{}{},
			},
		},
		Law: {
			unit: unit{
				Effects: []interface{}{
					newEffectScience(symbolLaw),
				},
			},
		},
		Masonry: {
			unit: unit{
				Effects: []interface{}{
					newEffectDiscounter(
						dcCivilian,
						2,
						clay, wood, stone, glass, papyrus,
					),
				},
			},
		},
		Mathematics: {
			unit: unit{
				Effects: []interface{}{
					newEffectMathematics(),
				},
			},
		},
		Philosophy: {
			unit: unit{
				Effects: []interface{}{
					newEffectPoints(7),
				},
			},
		},
		Strategy: {
			unit: unit{
				Effects: []interface{}{
					// replace has check to decorator?
					// uses only in has() context
				},
			},
		},
		Theology: {
			unit: unit{
				Effects: []interface{}{
					// uses only in has() context
				},
			},
		},
		Urbanism: {
			unit: unit{
				Effects: []interface{}{
					newEffectCoins(6),
					// + runtime has()
				},
			},
		},
	},
	layouts: map[age]string{
		ageI: `
    [][]
   [][][]
  [][][][]
 [][][][][]
[][][][][][]
`,
		ageII: `
[][][][][][]
 [][][][][]
  [][][][]
   [][][]
    [][]
`,
		ageIII: `
  [][]
 [][][]
[][][][]
 []  []
[][][][]
 [][][]
  [][]
`,
	},
}

type registry struct {
	Wonders   WonderMap
	Cards     CardMap
	wonderIds WonderList
	guilds    CardList
	ageCards  map[age]CardList
	tokens    tokenMap
	tokenIds  TokenList
	layouts   map[age]string
}

func init() {
	var ids WonderList

	for id, _ := range R.Wonders {
		ids = append(ids, id)
	}

	R.wonderIds = ids
}

func init() {
	var ids TokenList

	for id, _ := range R.tokens {
		ids = append(ids, id)
	}

	R.tokenIds = ids
}

func init() {
	var guilds CardList

	ageCards := map[age]CardList{
		ageI:   {},
		ageII:  {},
		ageIII: {},
	}

	for _, c := range R.Cards {
		if c.Group == groupGuild {
			guilds = append(guilds, c.Id)
			continue
		}

		ageCards[c.Age] = append(ageCards[c.Age], c.Id)
	}

	R.guilds = guilds
	R.ageCards = ageCards
}
