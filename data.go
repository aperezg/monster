package monster

var Monsters = map[string]*Monster{
	"01D2HFC5XCDMVHP80ZH44JNHZ4": &Monster{
		ID:      "01D2HFC5XCDMVHP80ZH44JNHZ4",
		Name:    "Chupacabra",
		Attack:  150,
		Defense: 500,
		Type:    EarthType,
	},
	"01D2HV4AYW1YK6TGXE2YZ4DQGR": &Monster{
		ID:      "01D2HV4AYW1YK6TGXE2YZ4DQGR",
		Name:    "Nessie",
		Attack:  10,
		Defense: 999,
		Type:    WaterType,
	},
	"01D2HV56HKYM349958P5ZSM24Q": &Monster{
		ID:      "01D2HV56HKYM349958P5ZSM24Q",
		Name:    "Zombie",
		Attack:  300,
		Defense: 100,
		Type:    EarthType,
	},
	"01D2HV77H0XTA1C5AV317KTNRS": &Monster{
		ID:      "01D2HV77H0XTA1C5AV317KTNRS",
		Name:    "Kraken",
		Attack:  700,
		Defense: 400,
		Type:    WaterType,
	},
	"01D2HVBX3NTV5SB34YJH9K7KW9": &Monster{
		ID:      "01D2HVBX3NTV5SB34YJH9K7KW9",
		Name:    "Godzilla",
		Attack:  999,
		Defense: 800,
		Type:    EarthType,
	},
}
