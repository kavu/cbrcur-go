// Copyright (C) 2014 Max Riveiro <kavu13@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package cbr

import (
	"testing"
	"time"
)

const (
	ruFile = "test_data/ru.xml"
	enFile = "test_data/en.xml"
)

func TestGetRuDaily(t *testing.T) {
	report, err := GetRuDaily()
	if err != nil {
		t.Fatal(err)
	}

	if len(report.Currencies) != 34 {
		t.Errorf("Report Currencies length mismatch — got %d, 1 expected.", len(report.Currencies))
	}

	if report.Currencies[0].CharCode != "AUD" {
		t.Errorf("Currency CharCode mismatch — got %s.", report.Currencies[0].CharCode)
	}

	date, err := report.DateAsTime()
	if err != nil {
		t.Fatal(err)
	}

	tomorrow := time.Now().Add(time.Hour * 24)

	if date.Day() != tomorrow.Day() {
		t.Errorf("Day Mismatch — got %d, %d expected.", date.Day(), tomorrow.Day())
	}

	if date.Month() != tomorrow.Month() {
		t.Errorf("Month Mismatch — got %d, %d expected.", date.Month(), tomorrow.Month())

	}

	if date.Year() != tomorrow.Year() {
		t.Errorf("Year Mismatch — got %d, %d expected.", date.Year(), tomorrow.Year())
	}
}

func TestGetEnDaily(t *testing.T) {
	report, err := GetEnDaily()
	if err != nil {
		t.Fatal(err)
	}

	if len(report.Currencies) != 34 {
		t.Errorf("Report Currencies length mismatch — got %d, 34 expected.", len(report.Currencies))
	}

	if report.Currencies[0].CharCode != "AUD" {
		t.Errorf("Currency CharCode mismatch — got %s.", report.Currencies[0].CharCode)
	}

	date, err := report.DateAsTime()
	if err != nil {
		t.Fatal(err)
	}

	tomorrow := time.Now().Add(time.Hour * 24)

	if date.Day() != tomorrow.Day() {
		t.Errorf("Day Mismatch — got %d, %d expected.", date.Day(), tomorrow.Day())
	}

	if date.Month() != tomorrow.Month() {
		t.Errorf("Month Mismatch — got %d, %d expected.", date.Month(), tomorrow.Month())

	}

	if date.Year() != tomorrow.Year() {
		t.Errorf("Year Mismatch — got %d, %d expected.", date.Year(), tomorrow.Year())
	}
}

func TestGetRuDailyForDate(t *testing.T) {
	forDate := time.Date(2014, 2, 1, 0, 0, 0, 0, time.UTC)

	report, err := GetEnDailyForDate(forDate)
	if err != nil {
		t.Fatal(err)
	}

	if len(report.Currencies) != 34 {
		t.Errorf("Report Currencies length mismatch — got %d, 1 expected.", len(report.Currencies))
	}

	if report.Currencies[0].CharCode != "AUD" {
		t.Errorf("Currency CharCode mismatch — got %s.", report.Currencies[0].CharCode)
	}

	date, err := report.DateAsTime()
	if err != nil {
		t.Fatal(err)
	}

	if date.Day() != forDate.Day() {
		t.Errorf("Day Mismatch — got %d, %d expected.", date.Day(), forDate.Day())
	}

	if date.Month() != forDate.Month() {
		t.Errorf("Month Mismatch — got %d, %d expected.", date.Day(), forDate.Month())

	}

	if date.Year() != forDate.Year() {
		t.Errorf("Year Mismatch — got %d, %d expected.", date.Day(), forDate.Year())
	}
}

func TestGetEnDailyForDate(t *testing.T) {
	forDate := time.Date(2014, 2, 1, 0, 0, 0, 0, time.UTC)

	report, err := GetRuDailyForDate(forDate)
	if err != nil {
		t.Fatal(err)
	}

	if len(report.Currencies) != 34 {
		t.Errorf("Report Currencies length mismatch — got %d, 34 expected.", len(report.Currencies))
	}

	if report.Currencies[0].CharCode != "AUD" {
		t.Errorf("Currency CharCode mismatch — got %s.", report.Currencies[0].CharCode)
	}

	date, err := report.DateAsTime()
	if err != nil {
		t.Fatal(err)
	}

	if date.Day() != forDate.Day() {
		t.Errorf("Day Mismatch — got %d, %d expected.", date.Day(), forDate.Day())
	}

	if date.Month() != forDate.Month() {
		t.Errorf("Month Mismatch — got %d, %d expected.", date.Day(), forDate.Month())

	}

	if date.Year() != forDate.Year() {
		t.Errorf("Year Mismatch — got %d, %d expected.", date.Day(), forDate.Year())
	}
}
