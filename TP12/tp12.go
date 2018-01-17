package main

import (
	"fmt"
)

//<editor-fold desc="gender">
type gender string

const (
	m gender = "male"
	f gender = "female"
)
//</editor-fold>

//<editor-fold desc="work">
type work string

const (
	developer work = "developer"
	sysadmin work = "sysadmin"
	manager work = "manager"
)
//</editor-fold>

//<editor-fold desc="currency">
type currency string

const (
	eur currency = "EUR"
	usd currency = "USD"
)
//</editor-fold>

type address struct {
	city    string
	country string
}

//<editor-fold desc="company">
type company struct {
	name         string
	headquarters address
}

var znk = company{"Zenika", address{"Paris", "France"}}
var msf = company{"Microsoft", address{"Redmond", "United States"}}
var apl = company{"Apple", address{"Cupertino", "United States"}}
var fbk = company{"Facebook", address{" Menlo Park", "United States"}}
var amz = company{"Amazon", address{"Seattle", "United States"}}
//</editor-fold>

type salary struct {
	amount   float64
	currency currency
}

type person struct {
	name   string
	sex    gender
	office address
}

//<editor-fold desc="certification">
type certification string

const (
	csm certification = "Certified Scrum Master"
	spring certification = "Spring Professional"
	mongodb certification = "MongoDB Professional"
	java8 certification = "Java SE 8 Programmer OCE"
	pmp certification = "PMI Professional Project Manager"
	cissp certification = "Certified Information Systems Security Professional"
)
//</editor-fold>

//<editor-fold desc="def worker">
type worker struct {
	// To complete
	person
	company        company
	role           work
	monthlyPay     salary
	certifications []certification
}

var worker1 = worker{
	person{
		"nom1",
		m,
		address{"Paris", "France"},
	},
	znk,
	developer,
	salary{
		10000,
		eur,
	},
	[]certification{
		csm,
		spring,
		mongodb,
		java8,
	},
} // To complete
var worker2 = worker{
	person{
		"nom2",
		f,
		address{"city2", "country2"},
	},
	msf,
	developer,
	salary{
		10000,
		usd,
	},
	[]certification{
		java8,
		pmp,
		cissp,
	},
} // To complete

// To complete : create certificationsNum() method
func (w worker) certificationsNum() int {
	return len(w.certifications)
}

// To complete : create isOfficeCompanyHeadquarters() method
func (w worker) isOfficeCompanyHeadquarters() bool {
	return w.person.office == w.company.headquarters
}
//</editor-fold>

func main() {
	workers := [2]worker{worker1, worker2}

	for _, w := range workers {
		var name string = w.person.name    // To complete
		var certsNum int = w.certificationsNum()  // To complete
		var worksAtHq bool = w.isOfficeCompanyHeadquarters() // To complete
		var company string = w.company.name // To complete
		fmt.Printf("%s has %d certifications", name, certsNum)
		if worksAtHq {
			fmt.Printf(" and his/her office is at %s headquarters\n", company)
		} else {
			fmt.Printf(" and his/her office is far from %s headquarters\n", company)
		}
	}
}
