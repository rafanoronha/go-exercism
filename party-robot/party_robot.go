package partyrobot

import "fmt"

const welcomeMessageFormat = "Welcome to my party, %s!"
const birthdayMessageFormat = "Happy birthday %s! You are now %d years old!"
const assignTableIntroMessageFormat = "%s\nYou have been assigned to table %03d."
const assignTableLocalizationMessageFormat = "Your table is %s, exactly %.1f meters from here."
const assignTableNeighborMessageFormat = "You will be sitting next to %s."
const assignTableFullMessageFormat = "%s %s\n%s"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf(welcomeMessageFormat, name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf(birthdayMessageFormat, name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return assignTableFullMessage(
		name,
		table,
		neighbor,
		direction,
		distance,
	)
}

func assignTableFullMessage(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf(
		assignTableFullMessageFormat,
		assignTableIntroMessage(name, table),
		assignTableLocalizationMessage(direction, distance),
		assignTableNeighborMessage(neighbor),
	)
}

func assignTableIntroMessage(name string, table int) string {
	return fmt.Sprintf(
		assignTableIntroMessageFormat,
		Welcome(name),
		table,
	)
}

func assignTableLocalizationMessage(direction string, distance float64) string {
	return fmt.Sprintf(
		assignTableLocalizationMessageFormat,
		direction,
		distance,
	)
}

func assignTableNeighborMessage(neighbor string) string {
	return fmt.Sprintf(
		assignTableNeighborMessageFormat,
		neighbor,
	)
}
