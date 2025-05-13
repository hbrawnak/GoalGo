/**
 * Given the next set of week of days as strings in the given order: <code>Mon, Tue, Wed, Thu, Fri, Sat, Sun</code> <br/>
 * and given a positive integer <code>K</code> in the range <code>[0..500]</code>. <br/>
 * Find the week of day for a given day plus K days.<br/>
 * For example:<br/>
 * <code>Wed + 5 = Mon</code>
 * <code>Sun + 0 = Sun</code>
 * <code>Tue + 28 = Tue</code>
 * <br/><br/>
 * Max time for resolution: 15 minutes.
 */

package main

import "fmt"

func WeekOfDayNextToK(s string, k int) string {
	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	dayToIndex := map[string]int{
		"Mon": 0, "Tue": 1, "Wed": 2, "Thu": 3,
		"Fri": 4, "Sat": 5, "Sun": 6,
	}

	key := dayToIndex[s]
	idx := (key + k) % 7
	return days[idx]
}

func main() {
	fmt.Println(WeekOfDayNextToK("Tue", 28)) // Tue
	fmt.Println(WeekOfDayNextToK("Wed", 5))  // Mon
	fmt.Println(WeekOfDayNextToK("Sun", 0))  // Sun
}
