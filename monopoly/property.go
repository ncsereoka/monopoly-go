package monopoly

type Property struct {
	owner  int
	name   string
	price  int
	houses uint8
	hotel  bool
}

var PropertyMap = map[int]*Property{
	1: {
		name:  "Old Kent Road",
		price: 60,
	},
	3: {
		name:  "Whitechapel Road",
		price: 60,
	},
	5: {
		name:  "Kings Cross Station",
		price: 200,
	},
	6: {
		name:  "The Angel, Islington",
		price: 100,
	},
	8: {
		name:  "Euston Road",
		price: 100,
	},
	9: {
		name:  "Pentonville Road",
		price: 120,
	},
	11: {
		name:  "Pall Mall",
		price: 140,
	},
	12: {
		name:  "Electric Company",
		price: 150,
	},
	13: {
		name:  "Whitehall",
		price: 140,
	},
	14: {
		name:  "Northumberland Avenue",
		price: 160,
	},
	15: {
		name:  "Marylebone",
		price: 200,
	},
	16: {
		name:  "Bow Street",
		price: 180,
	},
	18: {
		name:  "Marlborough Street",
		price: 180,
	},
	19: {
		name:  "Vine Street",
		price: 200,
	},
	21: {
		name:  "Strand",
		price: 200,
	},
	23: {
		name:  "Fleet Street",
		price: 200,
	},
	24: {
		name:  "Trafalgar Square",
		price: 240,
	},
	25: {
		name:  "Fenchurch Street",
		price: 200,
	},
	26: {
		name:  "Leicester Square",
		price: 260,
	},
	27: {
		name:  "Coventry Street",
		price: 260,
	},
	28: {
		name:  "Water Works",
		price: 150,
	},
	29: {
		name:  "Piccadilly",
		price: 280,
	},
	31: {
		name:  "Regent Street",
		price: 300,
	},
	32: {
		name:  "Oxford Street",
		price: 300,
	},
	34: {
		name:  "Bond Street",
		price: 320,
	},
	35: {
		name:  "Liverpool Street",
		price: 200,
	},
	37: {
		name:  "Park Lane",
		price: 350,
	},
	39: {
		name:  "Mayfair",
		price: 400,
	},
}
