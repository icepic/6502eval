package main

func Petmap(pet byte) rune {

	var temp rune
	p:=make(map[byte]rune)

	// Author: Rebecca Bettencourt <support@kreativekorp.com>

	p[0x20] = 0x0020	// SPACE
	p[0x21] = 0x0021	// EXCLAMATION MARK
	p[0x22] = 0x0022	// QUOTATION MARK
	p[0x23] = 0x0023	// NUMBER SIGN
	p[0x24] = 0x0024	// DOLLAR SIGN
	p[0x25] = 0x0025	// PERCENT SIGN
	p[0x26] = 0x0026	// AMPERSAND
	p[0x27] = 0x0027	// APOSTROPHE
	p[0x28] = 0x0028	// LEFT PARENTHESIS
	p[0x29] = 0x0029	// RIGHT PARENTHESIS
	p[0x2A] = 0x002A	// ASTERISK
	p[0x2B] = 0x002B	// PLUS SIGN
	p[0x2C] = 0x002C	// COMMA
	p[0x2D] = 0x002D	// HYPHEN-MINUS
	p[0x2E] = 0x002E	// FULL STOP
	p[0x2F] = 0x002F	// SOLIDUS
	p[0x30] = 0x0030	// DIGIT ZERO
	p[0x31] = 0x0031	// DIGIT ONE
	p[0x32] = 0x0032	// DIGIT TWO
	p[0x33] = 0x0033	// DIGIT THREE
	p[0x34] = 0x0034	// DIGIT FOUR
	p[0x35] = 0x0035	// DIGIT FIVE
	p[0x36] = 0x0036	// DIGIT SIX
	p[0x37] = 0x0037	// DIGIT SEVEN
	p[0x38] = 0x0038	// DIGIT EIGHT
	p[0x39] = 0x0039	// DIGIT NINE
	p[0x3A] = 0x003A	// COLON
	p[0x3B] = 0x003B	// SEMICOLON
	p[0x3C] = 0x003C	// LESS-THAN SIGN
	p[0x3D] = 0x003D	// EQUALS SIGN
	p[0x3E] = 0x003E	// GREATER-THAN SIGN
	p[0x3F] = 0x003F	// QUESTION MARK
	p[0x40] = 0x0040	// COMMERCIAL AT
	p[0x41] = 0x0041	// LATIN CAPITAL LETTER A
	p[0x42] = 0x0042	// LATIN CAPITAL LETTER B
	p[0x43] = 0x0043	// LATIN CAPITAL LETTER C
	p[0x44] = 0x0044	// LATIN CAPITAL LETTER D
	p[0x45] = 0x0045	// LATIN CAPITAL LETTER E
	p[0x46] = 0x0046	// LATIN CAPITAL LETTER F
	p[0x47] = 0x0047	// LATIN CAPITAL LETTER G
	p[0x48] = 0x0048	// LATIN CAPITAL LETTER H
	p[0x49] = 0x0049	// LATIN CAPITAL LETTER I
	p[0x4A] = 0x004A	// LATIN CAPITAL LETTER J
	p[0x4B] = 0x004B	// LATIN CAPITAL LETTER K
	p[0x4C] = 0x004C	// LATIN CAPITAL LETTER L
	p[0x4D] = 0x004D	// LATIN CAPITAL LETTER M
	p[0x4E] = 0x004E	// LATIN CAPITAL LETTER N
	p[0x4F] = 0x004F	// LATIN CAPITAL LETTER O
	p[0x50] = 0x0050	// LATIN CAPITAL LETTER P
	p[0x51] = 0x0051	// LATIN CAPITAL LETTER Q
	p[0x52] = 0x0052	// LATIN CAPITAL LETTER R
	p[0x53] = 0x0053	// LATIN CAPITAL LETTER S
	p[0x54] = 0x0054	// LATIN CAPITAL LETTER T
	p[0x55] = 0x0055	// LATIN CAPITAL LETTER U
	p[0x56] = 0x0056	// LATIN CAPITAL LETTER V
	p[0x57] = 0x0057	// LATIN CAPITAL LETTER W
	p[0x58] = 0x0058	// LATIN CAPITAL LETTER X
	p[0x59] = 0x0059	// LATIN CAPITAL LETTER Y
	p[0x5A] = 0x005A	// LATIN CAPITAL LETTER Z
	p[0x5B] = 0x005B	// LEFT SQUARE BRACKET
	p[0x5C] = 0x00A3	// POUND SIGN
	p[0x5D] = 0x005D	// RIGHT SQUARE BRACKET
	p[0x5E] = 0x2191	// UPWARDS ARROW
	p[0x5F] = 0x2190	// LEFTWARDS ARROW
	p[0x60] = 0x2500	// BOX DRAWINGS LIGHT HORIZONTAL
	p[0x61] = 0x2660	// BLACK SPADE SUIT
	//	p[0x62] = 0x1FB72	// VERTICAL ONE EIGHTH BLOCK-4
	//	p[0x63] = 0x1FB78  // HORIZONTAL ONE EIGHTH BLOCK-4
	//	p[0x64] = 0x1FB77  // HORIZONTAL ONE EIGHTH BLOCK-3
	//	p[0x65] = 0x1FB76  // HORIZONTAL ONE EIGHTH BLOCK-2
	//	p[0x66] = 0x1FB7A  // HORIZONTAL ONE EIGHTH BLOCK-6
	//	p[0x67] = 0x1FB71  // VERTICAL ONE EIGHTH BLOCK-3
	//	p[0x68] = 0x1FB74  // VERTICAL ONE EIGHTH BLOCK-6
	p[0x69] = 0x256E	// BOX DRAWINGS LIGHT ARC DOWN AND LEFT
	p[0x6A] = 0x2570	// BOX DRAWINGS LIGHT ARC UP AND RIGHT
	p[0x6B] = 0x256F	// BOX DRAWINGS LIGHT ARC UP AND LEFT
	//	p[0x6C] = 0x1FB7C	// LEFT AND LOWER ONE EIGHTH BLOCK
	p[0x6D] = 0x2572	// BOX DRAWINGS LIGHT DIAGONAL UPPER LEFT TO LOWER RIGHT
	p[0x6E] = 0x2571	// BOX DRAWINGS LIGHT DIAGONAL UPPER RIGHT TO LOWER LEFT
	//	p[0x6F] = 0x1FB7D	// LEFT AND UPPER ONE EIGHTH BLOCK
	//	p[0x70] = 0x1FB7E  // RIGHT AND UPPER ONE EIGHTH BLOCK
	p[0x71] = 0x2022	// BULLET (or 0x25CF BLACK CIRCLE)
	//	p[0x72] = 0x1FB7B	// HORIZONTAL ONE EIGHTH BLOCK-7
	p[0x73] = 0x2665	// BLACK HEART SUIT
	//	p[0x74] = 0x1FB70	// VERTICAL ONE EIGHTH BLOCK-2
	p[0x75] = 0x256D	// BOX DRAWINGS LIGHT ARC DOWN AND RIGHT
	p[0x76] = 0x2573	// BOX DRAWINGS LIGHT DIAGONAL CROSS
	p[0x77] = 0x25CB	// WHITE CIRCLE (or 0x25E6 WHITE BULLET)
	p[0x78] = 0x2663	// BLACK CLUB SUIT
	//	p[0x79] = 0x1FB75	// VERTICAL ONE EIGHTH BLOCK-7
	p[0x7A] = 0x2666	// BLACK DIAMOND SUIT
	p[0x7B] = 0x253C	// BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
	//	p[0x7C] = 0x1FB8C	// LEFT HALF MEDIUM SHADE
	p[0x7D] = 0x2502	// BOX DRAWINGS LIGHT VERTICAL
	p[0x7E] = 0x03C0	// GREEK SMALL LETTER PI
	p[0x7F] = 0x25E5	// BLACK UPPER RIGHT TRIANGLE
	p[0xA0] = 0x00A0	// NO-BREAK SPACE
	p[0xA1] = 0x258C	// LEFT HALF BLOCK
	p[0xA2] = 0x2584	// LOWER HALF BLOCK
	//	p[0xA3] = 0x2594	// UPPER ONE EIGHTH BLOCK
	//	p[0xA4] = 0x2581	// LOWER ONE EIGHTH BLOCK
	//	p[0xA5] = 0x258F	// LEFT ONE EIGHTH BLOCK
	p[0xA6] = 0x2592	// MEDIUM SHADE
	//	p[0xA7] = 0x2595	// RIGHT ONE EIGHTH BLOCK
	//	p[0xA8] = 0x1FB8F	// LOWER HALF MEDIUM SHADE
	p[0xA9] = 0x25E4	// BLACK UPPER LEFT TRIANGLE
	//	p[0xAA] = 0x1FB87	// RIGHT ONE QUARTER BLOCK
	p[0xAB] = 0x251C	// BOX DRAWINGS LIGHT VERTICAL AND RIGHT
	p[0xAC] = 0x2597	// QUADRANT LOWER RIGHT
	p[0xAD] = 0x2514	// BOX DRAWINGS LIGHT UP AND RIGHT
	p[0xAE] = 0x2510	// BOX DRAWINGS LIGHT DOWN AND LEFT
	p[0xAF] = 0x2582	// LOWER ONE QUARTER BLOCK
	p[0xB0] = 0x250C	// BOX DRAWINGS LIGHT DOWN AND RIGHT
	p[0xB1] = 0x2534	// BOX DRAWINGS LIGHT UP AND HORIZONTAL
	p[0xB2] = 0x252C	// BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
	p[0xB3] = 0x2524	// BOX DRAWINGS LIGHT VERTICAL AND LEFT
	p[0xB4] = 0x258E	// LEFT ONE QUARTER BLOCK
	//	p[0xB5] = 0x258D	// LEFT THREE EIGHTHS BLOCK
	//	p[0xB6] = 0x1FB88	// RIGHT THREE EIGHTHS BLOCK
	//	p[0xB7] = 0x1FB82 // UPPER ONE QUARTER BLOCK
	//	p[0xB8] = 0x1FB83 // UPPER THREE EIGHTHS BLOCK
	//	p[0xB9] = 0x2583	// LOWER THREE EIGHTHS BLOCK
	//	p[0xBA] = 0x1FB7F	// RIGHT AND LOWER ONE EIGHTH BLOCK
	p[0xBB] = 0x2596	// QUADRANT LOWER LEFT
	p[0xBC] = 0x259D	// QUADRANT UPPER RIGHT
	p[0xBD] = 0x2518	// BOX DRAWINGS LIGHT UP AND LEFT
	p[0xBE] = 0x2598	// QUADRANT UPPER LEFT
	p[0xBF] = 0x259A	// QUADRANT UPPER LEFT AND LOWER RIGHT
	p[0xC0] = 0x2500	// BOX DRAWINGS LIGHT HORIZONTAL
	p[0xC1] = 0x2660	// BLACK SPADE SUIT
	//	p[0xC2] = 0x1FB72	// VERTICAL ONE EIGHTH BLOCK-4
	//	p[0xC3] = 0x1FB78 // HORIZONTAL ONE EIGHTH BLOCK-4
	//	p[0xC4] = 0x1FB77 // HORIZONTAL ONE EIGHTH BLOCK-3
	//	p[0xC5] = 0x1FB76 // HORIZONTAL ONE EIGHTH BLOCK-2
	//	p[0xC6] = 0x1FB7A // HORIZONTAL ONE EIGHTH BLOCK-6
	//	p[0xC7] = 0x1FB71 // VERTICAL ONE EIGHTH BLOCK-3
	//	p[0xC8] = 0x1FB74 // VERTICAL ONE EIGHTH BLOCK-6
	p[0xC9] = 0x256E	// BOX DRAWINGS LIGHT ARC DOWN AND LEFT
	p[0xCA] = 0x2570	// BOX DRAWINGS LIGHT ARC UP AND RIGHT
	p[0xCB] = 0x256F	// BOX DRAWINGS LIGHT ARC UP AND LEFT
	//	p[0xCC] = 0x1FB7C	// LEFT AND LOWER ONE EIGHTH BLOCK
	p[0xCD] = 0x2572	// BOX DRAWINGS LIGHT DIAGONAL UPPER LEFT TO LOWER RIGHT
	p[0xCE] = 0x2571	// BOX DRAWINGS LIGHT DIAGONAL UPPER RIGHT TO LOWER LEFT
	//	p[0xCF] = 0x1FB7D	// LEFT AND UPPER ONE EIGHTH BLOCK
	//	p[0xD0] = 0x1FB7E // RIGHT AND UPPER ONE EIGHTH BLOCK
	p[0xD1] = 0x2022	// BULLET (or 0x25CF BLACK CIRCLE)
	//	p[0xD2] = 0x1FB7B	// HORIZONTAL ONE EIGHTH BLOCK-7
	p[0xD3] = 0x2665	// BLACK HEART SUIT
	//	p[0xD4] = 0x1FB70	// VERTICAL ONE EIGHTH BLOCK-2
	p[0xD5] = 0x256D	// BOX DRAWINGS LIGHT ARC DOWN AND RIGHT
	p[0xD6] = 0x2573	// BOX DRAWINGS LIGHT DIAGONAL CROSS
	p[0xD7] = 0x25CB	// WHITE CIRCLE (or 0x25E6 WHITE BULLET)
	p[0xD8] = 0x2663	// BLACK CLUB SUIT
	//	p[0xD9] = 0x1FB75	// VERTICAL ONE EIGHTH BLOCK-7
	p[0xDA] = 0x2666	// BLACK DIAMOND SUIT
	p[0xDB] = 0x253C	// BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
	//	p[0xDC] = 0x1FB8C	// LEFT HALF MEDIUM SHADE
	p[0xDD] = 0x2502	// BOX DRAWINGS LIGHT VERTICAL
	p[0xDE] = 0x03C0	// GREEK SMALL LETTER PI
	p[0xDF] = 0x25E5	// BLACK UPPER RIGHT TRIANGLE
	p[0xE0] = 0x00A0	// NO-BREAK SPACE
	p[0xE1] = 0x258C	// LEFT HALF BLOCK
	p[0xE2] = 0x2584	// LOWER HALF BLOCK
	//	p[0xE3] = 0x2594	// UPPER ONE EIGHTH BLOCK
	//	p[0xE4] = 0x2581	// LOWER ONE EIGHTH BLOCK
	//	p[0xE5] = 0x258F	// LEFT ONE EIGHTH BLOCK
	p[0xE6] = 0x2592	// MEDIUM SHADE
	//	p[0xE7] = 0x2595	// RIGHT ONE EIGHTH BLOCK
	//	p[0xE8] = 0x1FB8F	// LOWER HALF MEDIUM SHADE
	p[0xE9] = 0x25E4	// BLACK UPPER LEFT TRIANGLE
	//	p[0xEA] = 0x1FB87	// RIGHT ONE QUARTER BLOCK
	p[0xEB] = 0x251C	// BOX DRAWINGS LIGHT VERTICAL AND RIGHT
	p[0xEC] = 0x2597	// QUADRANT LOWER RIGHT
	p[0xED] = 0x2514	// BOX DRAWINGS LIGHT UP AND RIGHT
	p[0xEE] = 0x2510	// BOX DRAWINGS LIGHT DOWN AND LEFT
	p[0xEF] = 0x2582	// LOWER ONE QUARTER BLOCK
	p[0xF0] = 0x250C	// BOX DRAWINGS LIGHT DOWN AND RIGHT
	p[0xF1] = 0x2534	// BOX DRAWINGS LIGHT UP AND HORIZONTAL
	p[0xF2] = 0x252C	// BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
	p[0xF3] = 0x2524	// BOX DRAWINGS LIGHT VERTICAL AND LEFT
	p[0xF4] = 0x258E	// LEFT ONE QUARTER BLOCK
	//	p[0xF5] = 0x258D	// LEFT THREE EIGHTHS BLOCK
	//	p[0xF6] = 0x1FB88	// RIGHT THREE EIGHTHS BLOCK
	//	p[0xF7] = 0x1FB82 // UPPER ONE QUARTER BLOCK
	//	p[0xF8] = 0x1FB83 // UPPER THREE EIGHTHS BLOCK
	//	p[0xF9] = 0x2583	// LOWER THREE EIGHTHS BLOCK
	//	p[0xFA] = 0x1FB7F	// RIGHT AND LOWER ONE EIGHTH BLOCK
	p[0xFB] = 0x2596	// QUADRANT LOWER LEFT
	p[0xFC] = 0x259D	// QUADRANT UPPER RIGHT
	p[0xFD] = 0x2518	// BOX DRAWINGS LIGHT UP AND LEFT
	p[0xFE] = 0x2598	// QUADRANT UPPER LEFT
	p[0xFF] = 0x259A	// QUADRANT UPPER LEFT AND LOWER RIGHT

	temp, ok := p[pet]
	if ok == true {
		return temp
	} else {
		return 0x0020 // space for unprintables
	}
}
