package bottlesong

func NumberToStringLowerCase(number int) string {
	switch number {
	case 0:
		return "no"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	default:
		return "invalid number"
	}
}

func NumberToStringCaptialized(number int) string {
	switch number {
	case 0:
		return "No"
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	default:
		return "invalid number"
	}
}
func Recite(startBottles, takeDown int) []string {
	verses := make([]string, 0, takeDown*5)
	const verse1Verse2 = " green bottles hanging on the wall,"
	const verse1Verse2Singular = " green bottle hanging on the wall,"
	const verse3 = "And if one green bottle should accidentally fall,"
	const verse4 = "There'll be "
	const verse4continued = " green bottles hanging on the wall."
	const verse4continuedsingular = " green bottle hanging on the wall."

	for takeDown > 0 {
		if startBottles == 1 {
			verses = append(verses, NumberToStringCaptialized(startBottles)+verse1Verse2Singular)
			verses = append(verses, NumberToStringCaptialized(startBottles)+verse1Verse2Singular)
		} else {
			verses = append(verses, NumberToStringCaptialized(startBottles)+verse1Verse2)
			verses = append(verses, NumberToStringCaptialized(startBottles)+verse1Verse2)
		}
		verses = append(verses, verse3)
		if startBottles-1 == 1 {
			verses = append(verses, verse4+NumberToStringLowerCase(startBottles-1)+verse4continuedsingular)
		} else {
			verses = append(verses, verse4+NumberToStringLowerCase(startBottles-1)+verse4continued)
		}
		if takeDown > 1 {
			verses = append(verses, "")
		}
		startBottles--
		takeDown--
	}
	return verses
}
