package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type table struct {
	Name      string    `json:"name"`
	Capacity  int       `json:"capacity"`
	Free      bool      `json:"free"`
	Fulluntil time.Time `json:"fulluntil"`
}

func main() {

	now := time.Now()
	czechMonths := map[string]string{
		"January": "ledna", "February": "února", "March": "března", "April": "dubna",
		"May": "května", "June": "června", "July": "července", "August": "srpna",
		"September": "září", "October": "října", "November": "listopadu", "December": "prosince",
	}
	longFormat := now.Format("2. January 2006 15:04:05")
	for eng, czech := range czechMonths {
		longFormat = strings.ReplaceAll(longFormat, eng, czech)
	}

	// přečtění JSON dat
	file, err := os.ReadFile("tables.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parsování JSON
	var rawTables []map[string]interface{}
	if err := json.Unmarshal(file, &rawTables); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	// prevedeni na struct
	var tables []table
	for _, rawTable := range rawTables {
		fulluntilStr, _ := rawTable["fulluntil"].(string)
		fulluntil, err := time.Parse(time.RFC3339, fulluntilStr)
		if err != nil {
			fulluntil = time.Time{}
		}
		tables = append(tables, table{
			Name:      rawTable["name"].(string),
			Capacity:  int(rawTable["capacity"].(float64)),
			Free:      rawTable["free"].(bool),
			Fulluntil: fulluntil,
		})
	}
	println("Otevřeno je od 13:00 do 0:00 pondělí až čtvrtek, pátek a sobota od 13:00 do 1:00, neděle pak 15:00 až 23:00")

	// Napiš rezervované stoly
	for _, t := range tables {
		if !t.Free {
			if !t.Fulluntil.IsZero() {
				fmt.Printf("Stůl %s je rezervovaný %s\n", t.Name, formatCzechDate(t.Fulluntil, czechMonths))
			} else {
				fmt.Printf("Stůl %s je rezervovaný, ale čas je neznámý.\n", t.Name)
			}
		}
	}

	println("Pro kolik lidí chceš rezervovat stůl?")

	var answer int
	var whattable string
	fmt.Scanf("%d\n", &answer)

	// Napiš volné stoly s vhodnou kapacitou
	fmt.Printf("Tyto stoly jsou volné:\n")
	for _, t := range tables {
		if t.Capacity >= answer && t.Free {
			fmt.Printf("Stůl %s s kapacitou pro %d.\n", t.Name, t.Capacity)
		}
	}

	fmt.Println("Zadej jméno stolu, co chceš rezervovat: ")
	fmt.Scanf("%s", &whattable)
	for t := range tables {
		if tables[t].Name == whattable {
			tables[t].Free = false
			tables[t].Fulluntil = time.Now().Add(3 * time.Hour)
			fmt.Printf("Zarezervovaný stůl do: Jméno: %s, Kapacita: %d, Volný: %s\n", tables[t].Name, tables[t].Capacity, formatCzechDate(tables[t].Fulluntil, czechMonths))
			break
		}
	}
	for _, t := range tables {
		if !t.Free {
			if time.Now().After(t.Fulluntil) {
				fmt.Printf("Stůl %s už není rezervovaný.\n", t.Name)
				t.Free = true
			} else {
				fmt.Printf("Stůl %s je rezervovaný %s\n", t.Name, formatCzechDate(t.Fulluntil, czechMonths))
			}
		}
	}

	updatedJSON, err := json.MarshalIndent(tables, "", "  ")
	if err != nil {
		fmt.Println("Error generování JSONu:", err)
		return
	}

	if err := os.WriteFile("tables.json", updatedJSON, 0644); err != nil {
		fmt.Println("Error zapisování do JSONu:", err)
		return
	}

}

func formatCzechDate(t time.Time, czechMonths map[string]string) string {
	englishFormat := t.Format("2. January 2006 15:04:05")
	for eng, czech := range czechMonths {
		englishFormat = strings.ReplaceAll(englishFormat, eng, czech)
	}
	return englishFormat
}
