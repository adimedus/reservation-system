package main

import (
	"fmt"
	//	"net/http"
	"os"
	"strings"
	"time"
)

type table struct {
	name      string
	capacity  int
	free      bool
	fulluntil time.Time
}

func filter(data []table, f func(table) bool) []table {

	fltd := make([]table, 0)
	free := true

	for _, table := range data {
		if table.free == free {
			if f(table) {
				fltd = append(fltd, table)
			}
		}
	}

	return fltd
}

func main() {

	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

	now := time.Now()
	czechMonths := map[string]string{
		"January":   "ledna",
		"February":  "února",
		"March":     "března",
		"April":     "dubna",
		"May":       "května",
		"June":      "června",
		"July":      "července",
		"August":    "srpna",
		"September": "září",
		"October":   "října",
		"November":  "listopadu",
		"December":  "prosince",
	}
	longFormat := now.Format("2. January 2006 15:04:05")
	for eng, czech := range czechMonths {
		longFormat = strings.ReplaceAll(longFormat, eng, czech)
	}

	tables := []table{
		{"Velký", 6, true, time.Time{}},
		{"Dveře", 2, true, time.Time{}},
		{"Okno", 2, true, time.Time{}},
		{"Čtyřka okno", 4, true, time.Time{}},
		{"Obraz", 4, true, time.Time{}},
		{"Neon", 2, false, time.Now().Add(2 * time.Hour)},
		{"Záchod", 2, true, time.Time{}},
	}

	freefortwo := 2
	freeforfour := 4
	freeforsix := 6

	freecap2 := filter(tables, func(i table) bool {
		return i.capacity == freefortwo
	})
	freecap4 := filter(tables, func(i table) bool {
		return i.capacity == freeforfour
	})
	freecap6 := filter(tables, func(i table) bool {
		return i.capacity == freeforsix
	})

	println("Otevřeno je od 13:00 do 0:00 pondělí až čtvrtek, pátek a sobota od 13:00 do 1:00, neděle pak 15:00 až 23:00")

	println("Pro kolik lidí chceš rezervovat stůl?")

	var answer int
	var whattable string = "Velký"
	fmt.Scanf("%d\n", &answer)

	if answer <= 2 && len(freecap2) == 0 {
		fmt.Println("Pro dva stůl není, co takhle větší stůl: ")
		for _, t := range freecap4 {
			fmt.Printf("Jméno: %s, Kapacita: %d lidi\n", t.name, t.capacity)
		}
		for _, t := range freecap6 {
			fmt.Printf("Jméno: %s, Kapacita: %d lidí\n", t.name, t.capacity)
		}
	} else if answer <= 2 {
		fmt.Println("Volné stoly jsou:")
		for _, t := range freecap2 {
			fmt.Printf("Jméno: %s, Kapacita: %d lidi\n", t.name, t.capacity)
		}
	} else if answer <= 4 && len(freecap4) == 0 {
		fmt.Println("Nemáme stůl pro čtyři volný, máme ale volný:")
		for _, t := range freecap6 {
			fmt.Printf("Jméno: %s, Kapacita: %d lidí\n", t.name, t.capacity)
		}
	} else if answer <= 4 {
		fmt.Println("Volné stoly jsou:")
		for _, t := range freecap4 {
			fmt.Printf("Jméno: %s, Kapacita: %d lidi\n", t.name, t.capacity)
		}
		for _, t := range freecap6 {
			fmt.Fprintln(os.Stdout, []any{`Také volný je:`}...)
			fmt.Printf("Jméno: %s, Kapacita: %d lidi\n", t.name, t.capacity)
		}
	} else if answer <= 6 && len(freecap6) == 0 {
		fmt.Printf("Nemáme volný stůl pro tolik lidí, chceš stoly spojit?(Ano/Ne)")
	} else if answer <= 6 {
		fmt.Println("Volný je:")
		for _, t := range freecap6 {
			fmt.Printf("Jméno: %s, Kapacita: %d lidí\n", t.name, t.capacity)
		}
	} else {
		fmt.Println("Nemáme tak velký stůl, chceš spojit stoly?(Ano/Ne)")

	}

	fmt.Println("Zadej jméno stolu, co chceš rezervovat: ")
	fmt.Scanf("%s", &whattable)
	for t := range tables {
		if tables[t].name == whattable {
			tables[t].free = false
			tables[t].fulluntil = time.Now().Add(3 * time.Hour)
			fmt.Printf("Zarezervovaný stůl do: Jméno: %s, Kapacita: %d, Volný: %s\n", tables[t].name, tables[t].capacity, formatCzechDate(tables[t].fulluntil, czechMonths))
			break
		}
	}
	for _, t := range tables {
		if !t.free {
			if time.Now().After(t.fulluntil) {
				fmt.Printf("Stůl %s už není rezervovaný.\n", t.name)
				t.free = true
			} else {
				fmt.Printf("Stůl %s je rezervovaný %s\n", t.name, formatCzechDate(t.fulluntil, czechMonths))
			}
		}
	}

}

func formatCzechDate(t time.Time, czechMonths map[string]string) string {
	englishFormat := t.Format("2. January 2006 15:04:05")
	for eng, czech := range czechMonths {
		englishFormat = strings.ReplaceAll(englishFormat, eng, czech)
	}
	return englishFormat
}

//func handler(w http.ResponseWriter, r *http.Request) {

//fmt.Fprintf(w, "Otevřeno je od 13:00 do 0:00 pondělí až čtvrtek, pátek a sobota od 13:00 do 1:00, neděle pak 15:00 až 23:00\n")

//fmt.Fprintf(w, "Pro kolik lidí chceš rezervovat stůl?\n")

//}
