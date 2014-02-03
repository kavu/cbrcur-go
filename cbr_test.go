// Copyright (C) 2014 Max Riveiro <kavu13@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package cbr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	RU_FILE = "test_data/ru.xml"
	EN_FILE = "test_data/en.xml"
)

func TestGetRuDaily(t *testing.T) {
	ru, err := ioutil.ReadFile(RU_FILE)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(ru))
	}))

	RU_URL = ts.URL

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

	if date.Day() != 1 {
		t.Errorf("Day Mismatch — got %d, 1 expected.", date.Day())
	}

	if date.Month() != 2 {
		t.Errorf("Month Mismatch — got %d, 2 expected.", date.Day())

	}

	if date.Year() != 2014 {
		t.Errorf("Year Mismatch — got %d, 2014 expected.", date.Day())
	}
}

func TestGetEnDaily(t *testing.T) {
	data, err := ioutil.ReadFile(EN_FILE)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(data))
	}))

	EN_URL = ts.URL

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

	if date.Day() != 1 {
		t.Errorf("Day Mismatch — got %d, 1 expected.", date.Day())
	}

	if date.Month() != 2 {
		t.Errorf("Month Mismatch — got %d, 2 expected.", date.Day())

	}

	if date.Year() != 2014 {
		t.Errorf("Year Mismatch — got %d, 2014 expected.", date.Day())
	}
}

func TestGetRuDailyForDate(t *testing.T) {
	ru, err := ioutil.ReadFile(RU_FILE)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(ru))
	}))

	RU_URL = ts.URL

	for_date := time.Date(2014, 2, 1, 0, 0, 0, 0, time.UTC)

	report, err := GetEnDailyForDate(for_date)
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

	if date.Day() != 1 {
		t.Errorf("Day Mismatch — got %d, 1 expected.", date.Day())
	}

	if date.Month() != 2 {
		t.Errorf("Month Mismatch — got %d, 2 expected.", date.Day())

	}

	if date.Year() != 2014 {
		t.Errorf("Year Mismatch — got %d, 2014 expected.", date.Day())
	}
}

func TestGetEnDailyForDate(t *testing.T) {
	data, err := ioutil.ReadFile(EN_FILE)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(data))
	}))

	EN_URL = ts.URL

	for_date := time.Date(2014, 2, 1, 0, 0, 0, 0, time.UTC)

	report, err := GetRuDailyForDate(for_date)
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

	if date.Day() != 1 {
		t.Errorf("Day Mismatch — got %d, 1 expected.", date.Day())
	}

	if date.Month() != 2 {
		t.Errorf("Month Mismatch — got %d, 2 expected.", date.Day())

	}

	if date.Year() != 2014 {
		t.Errorf("Year Mismatch — got %d, 2014 expected.", date.Day())
	}
}
