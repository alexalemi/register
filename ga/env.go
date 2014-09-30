package main

var p0 Program = Program{
	0: Instruction{inst: DEC, reg: 0, next: 1, jump: 10},
	1: Instruction{inst: DEC, reg: 0, next: 2, jump: 22},
	2: Instruction{inst: DEC, reg: 0, next: 3, jump: 27},
	3: Instruction{inst: DEC, reg: 0, next: 4, jump: 38},
	4: Instruction{inst: DEC, reg: 0, next: 5, jump: 48},
	5: Instruction{inst: DEC, reg: 0, next: 6, jump: 57},
	6: Instruction{inst: DEC, reg: 0, next: 7, jump: 68},
	7: Instruction{inst: DEC, reg: 0, next: 8, jump: 80},
	8: Instruction{inst: DEC, reg: 0, next: 9, jump: 87},
	9: Instruction{inst: DEC, reg: 0, next: 127, jump: 100},
	// draw zero
	10: Instruction{inst: INC, reg: 1, next: 11},
	11: Instruction{inst: INC, reg: 2, next: 12},
	12: Instruction{inst: INC, reg: 3, next: 13},
	13: Instruction{inst: INC, reg: 4, next: 14},
	14: Instruction{inst: INC, reg: 7, next: 15},
	15: Instruction{inst: INC, reg: 10, next: 16},
	16: Instruction{inst: INC, reg: 13, next: 17},
	17: Instruction{inst: INC, reg: 14, next: 18},
	18: Instruction{inst: INC, reg: 15, next: 19},
	19: Instruction{inst: INC, reg: 12, next: 20},
	20: Instruction{inst: INC, reg: 9, next: 21},
	21: Instruction{inst: INC, reg: 6, next: 127},
	// draw one
	22: Instruction{inst: INC, reg: 2, next: 23},
	23: Instruction{inst: INC, reg: 5, next: 24},
	24: Instruction{inst: INC, reg: 8, next: 25},
	25: Instruction{inst: INC, reg: 11, next: 26},
	26: Instruction{inst: INC, reg: 14, next: 127},
	// draw two
	27: Instruction{inst: INC, reg: 1, next: 28},
	28: Instruction{inst: INC, reg: 2, next: 29},
	29: Instruction{inst: INC, reg: 3, next: 30},
	30: Instruction{inst: INC, reg: 6, next: 31},
	31: Instruction{inst: INC, reg: 7, next: 32},
	32: Instruction{inst: INC, reg: 8, next: 33},
	33: Instruction{inst: INC, reg: 9, next: 34},
	34: Instruction{inst: INC, reg: 10, next: 35},
	35: Instruction{inst: INC, reg: 13, next: 36},
	36: Instruction{inst: INC, reg: 14, next: 37},
	37: Instruction{inst: INC, reg: 15, next: 127},
	// draw three
	38: Instruction{inst: INC, reg: 1, next: 39},
	39: Instruction{inst: INC, reg: 2, next: 40},
	40: Instruction{inst: INC, reg: 3, next: 41},
	41: Instruction{inst: INC, reg: 6, next: 42},
	42: Instruction{inst: INC, reg: 7, next: 43},
	43: Instruction{inst: INC, reg: 8, next: 44},
	44: Instruction{inst: INC, reg: 9, next: 45},
	45: Instruction{inst: INC, reg: 12, next: 46},
	46: Instruction{inst: INC, reg: 13, next: 47},
	47: Instruction{inst: INC, reg: 14, next: 112},
	// draw four
	48: Instruction{inst: INC, reg: 1, next: 49},
	49: Instruction{inst: INC, reg: 4, next: 50},
	50: Instruction{inst: INC, reg: 7, next: 51},
	51: Instruction{inst: INC, reg: 8, next: 52},
	52: Instruction{inst: INC, reg: 9, next: 53},
	53: Instruction{inst: INC, reg: 3, next: 54},
	54: Instruction{inst: INC, reg: 6, next: 55},
	55: Instruction{inst: INC, reg: 12, next: 56},
	56: Instruction{inst: INC, reg: 15, next: 127},
	// draw five
	57: Instruction{inst: INC, reg: 3, next: 58},
	58: Instruction{inst: INC, reg: 2, next: 59},
	59: Instruction{inst: INC, reg: 1, next: 60},
	60: Instruction{inst: INC, reg: 4, next: 61},
	61: Instruction{inst: INC, reg: 7, next: 62},
	62: Instruction{inst: INC, reg: 8, next: 63},
	63: Instruction{inst: INC, reg: 9, next: 64},
	64: Instruction{inst: INC, reg: 12, next: 65},
	65: Instruction{inst: INC, reg: 15, next: 66},
	66: Instruction{inst: INC, reg: 14, next: 67},
	67: Instruction{inst: INC, reg: 13, next: 127},
	// draw six
	68: Instruction{inst: INC, reg: 3, next: 69},
	69: Instruction{inst: INC, reg: 2, next: 70},
	70: Instruction{inst: INC, reg: 1, next: 71},
	71: Instruction{inst: INC, reg: 4, next: 72},
	72: Instruction{inst: INC, reg: 7, next: 73},
	73: Instruction{inst: INC, reg: 8, next: 74},
	74: Instruction{inst: INC, reg: 9, next: 75},
	75: Instruction{inst: INC, reg: 10, next: 76},
	76: Instruction{inst: INC, reg: 12, next: 77},
	77: Instruction{inst: INC, reg: 13, next: 78},
	78: Instruction{inst: INC, reg: 14, next: 79},
	79: Instruction{inst: INC, reg: 15, next: 127},
	// draw seven
	80: Instruction{inst: INC, reg: 1, next: 81},
	81: Instruction{inst: INC, reg: 2, next: 82},
	82: Instruction{inst: INC, reg: 3, next: 83},
	83: Instruction{inst: INC, reg: 6, next: 84},
	84: Instruction{inst: INC, reg: 9, next: 85},
	85: Instruction{inst: INC, reg: 11, next: 86},
	86: Instruction{inst: INC, reg: 13, next: 127},
	// draw eight
	87: Instruction{inst: INC, reg: 13, next: 88},
	88: Instruction{inst: INC, reg: 14, next: 89},
	89: Instruction{inst: INC, reg: 15, next: 90},
	90: Instruction{inst: INC, reg: 12, next: 91},
	91: Instruction{inst: INC, reg: 10, next: 92},
	92: Instruction{inst: INC, reg: 9, next: 93},
	93: Instruction{inst: INC, reg: 8, next: 94},
	94: Instruction{inst: INC, reg: 7, next: 95},
	95: Instruction{inst: INC, reg: 4, next: 96},
	96: Instruction{inst: INC, reg: 6, next: 97},
	97: Instruction{inst: INC, reg: 1, next: 98},
	98: Instruction{inst: INC, reg: 2, next: 99},
	99: Instruction{inst: INC, reg: 3, next: 127},
	// draw nine
	100: Instruction{inst: INC, reg: 1, next: 101},
	101: Instruction{inst: INC, reg: 2, next: 102},
	102: Instruction{inst: INC, reg: 3, next: 103},
	103: Instruction{inst: INC, reg: 4, next: 104},
	104: Instruction{inst: INC, reg: 6, next: 105},
	105: Instruction{inst: INC, reg: 7, next: 106},
	106: Instruction{inst: INC, reg: 8, next: 107},
	107: Instruction{inst: INC, reg: 9, next: 108},
	108: Instruction{inst: INC, reg: 12, next: 109},
	109: Instruction{inst: INC, reg: 15, next: 110},
	110: Instruction{inst: INC, reg: 13, next: 111},
	111: Instruction{inst: INC, reg: 14, next: 127},
	// overflow
	112: Instruction{inst: INC, reg: 15, next: 127},
}
var r Recipe = Recipe{
	Expectation{World{0: 0},
		World{1: 1, 2: 1, 3: 1, 4: 1, 7: 1, 10: 1, 13: 1, 14: 1, 15: 1, 12: 1, 9: 1, 6: 1}},
	Expectation{World{0: 1},
		World{2: 1, 5: 1, 8: 1, 11: 1, 14: 1}},
	Expectation{World{0: 2},
		World{1: 1, 2: 1, 3: 1, 6: 1, 9: 1, 8: 1, 7: 1, 10: 1, 13: 1, 14: 1, 15: 1}},
	Expectation{World{0: 3},
		World{1: 1, 2: 1, 3: 1, 6: 1, 9: 1, 8: 1, 7: 1, 12: 1, 15: 1, 14: 1, 13: 1}},
	Expectation{World{0: 4},
		World{1: 1, 4: 1, 7: 1, 8: 1, 3: 1, 6: 1, 9: 1, 12: 1, 15: 1}},
	Expectation{World{0: 5},
		World{3: 1, 2: 1, 1: 1, 4: 1, 7: 1, 8: 1, 9: 1, 12: 1, 15: 1, 14: 1, 13: 1}},
	Expectation{World{0: 6},
		World{3: 1, 2: 1, 1: 1, 4: 1, 7: 1, 8: 1, 9: 1, 12: 1, 15: 1, 14: 1, 13: 1, 10: 1}},
	Expectation{World{0: 7},
		World{1: 1, 2: 1, 3: 1, 6: 1, 9: 1, 11: 1, 13: 1}},
	Expectation{World{0: 8},
		World{1: 1, 2: 1, 3: 1, 6: 1, 4: 1, 7: 1, 8: 1, 9: 1, 10: 1, 12: 1, 13: 1, 14: 1, 15: 1}},
	Expectation{World{0: 9},
		World{1: 1, 2: 1, 3: 1, 4: 1, 6: 1, 7: 1, 8: 1, 9: 1, 12: 1, 15: 1, 13: 1, 14: 1}},
}