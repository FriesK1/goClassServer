package main

import (
	"time"

	DataObject "github.com/FriesK1/goClassData"
)

var (
	Staff DataObject.Employees
)

func init() {
	Staff = make(DataObject.Employees)

	Staff["bharath"] = DataObject.Employee{
		FirstName:       "Bharath",
		LastName:        "Veeramalli",
		StartDate:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		TerminationDate: time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC),
		Active:          true,
	}

	Staff["gabriel"] = DataObject.Employee{
		FirstName:       "Gabriel",
		LastName:        "Matias",
		StartDate:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		TerminationDate: time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC),
		Active:          true,
	}

	Staff["hugo"] = DataObject.Employee{
		FirstName:       "Hugo",
		LastName:        "Bedolla",
		StartDate:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		TerminationDate: time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC),
		Active:          true,
	}

	Staff["jesus"] = DataObject.Employee{
		FirstName:       "Jesus",
		LastName:        "Hernandez",
		StartDate:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		TerminationDate: time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC),
		Active:          true,
	}

	Staff["kevin"] = DataObject.Employee{
		FirstName:       "Kevin",
		LastName:        "Fries",
		StartDate:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		TerminationDate: time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC),
		Active:          true,
	}

	Staff["mario"] = DataObject.Employee{
		FirstName:       "Mario",
		LastName:        "Flores",
		StartDate:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		TerminationDate: time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC),
		Active:          true,
	}

	Staff["sriman"] = DataObject.Employee{
		FirstName:       "Sriman",
		LastName:        "Yarasingu",
		StartDate:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		TerminationDate: time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC),
		Active:          true,
	}

}
