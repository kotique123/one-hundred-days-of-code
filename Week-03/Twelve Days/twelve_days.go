package twelve

func Verse(day int) string {
	if day == 1 {
		return "On the first day of Christmas my true love gave to me: a Partridge in a Pear Tree."
	} else {
		days := []string{"second", "third", "fourth", "fifth", "sixth", "seventh",
			"eighth", "ninth", "tenth", "eleventh", "twelfth"}
		quantities := []string{"two", "three", "four", "five", "six", "seven", "eight",
			"nine", "ten", "eleven", "twelve"}
		gifts := []string{"Turtle Doves", "French Hens", "Calling Birds",
			"Gold Rings", "Geese-a-Laying", "Swans-a-Swimming", "Maids-a-Milking",
			"Ladies Dancing", "Lords-a-Leaping", "Pipers Piping", "Drummers Drumming"}
		result := "On the " + days[day-2] + " day of Christmas my true love gave to me: "
		for day > 1 {
			result += quantities[day-2] + " " + gifts[day-2] + ", "
			day--
		}
		return result + "and a Partridge in a Pear Tree."
	}
}

func Song() string {
	result := ""
	for i := 1; i <= 12; i++ {
		if i == 12 {
			result += Verse(i)
		} else {
			result += Verse(i) + "\n"
		}
	}
	return result
}
