// Copyright (C) 2014 Max Riveiro <kavu13@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package cbr

import (
	"encoding/xml"
	"github.com/djimenez/iconv-go"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	RU_URL string = "http://www.cbr.ru/scripts/XML_daily.asp"
	EN_URL string = "http://www.cbr.ru/scripts/XML_daily_eng.asp"
)

type Currency struct {
	ID       string `xml:",attr"`
	NumCode  int
	CharCode string
	Nominal  int
	Name     string
	Value    float64
}

type CurrencyReport struct {
	Date       string     `xml:",attr"`
	Currencies []Currency `xml:"Valute"`
}

func (self *CurrencyReport) DateAsTime() (date time.Time, err error) {
	date, err = time.Parse("02.01.2006", self.Date)
	// try to parse again
	// because date attribute format is different, fuck that
	if err != nil {
		date, err = time.Parse("02/01/2006", self.Date)
	}

	return date, err
}

func convertXML(res string) (string, error) {
	s, err := iconv.ConvertString(res, "windows-1251", "utf-8")
	if err != nil {
		return "", err
	}

	return s, nil
}

func decodeXMLBody(res string) (report *CurrencyReport, err error) {
	reader := strings.NewReader(strings.Replace(res, `,`, `.`, -1))

	if err := xml.NewDecoder(reader).Decode(&report); err != nil {
		return nil, err
	}

	return report, err
}

func getXML(URL string) (s string, err error) {
	res, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(b), err
}

func GetRuDaily() (report *CurrencyReport, err error) {
	s, err := getXML(RU_URL)
	if err != nil {
		return nil, err
	}

	converted, err := convertXML(s)
	if err != nil {
		return nil, err
	}

	cleaned := strings.Replace(converted, ` encoding="windows-1251"`, ` encoding="utf-8"`, -1)

	return decodeXMLBody(cleaned)
}

func GetEnDaily() (report *CurrencyReport, err error) {
	s, err := getXML(EN_URL)
	if err != nil {
		return nil, err
	}

	converted := strings.Replace(s, ` encoding="windows-1252"`, ` encoding="utf-8"`, -1)

	return decodeXMLBody(converted)
}

func GetRuDailyForDate(date time.Time) (report *CurrencyReport, err error) {
	s, err := getXML(RU_URL + "?date_req=" + date.Format("02/01/2006"))
	if err != nil {
		return nil, err
	}

	converted, err := convertXML(s)
	if err != nil {
		return nil, err
	}

	cleaned := strings.Replace(converted, ` encoding="windows-1251"`, ` encoding="utf-8"`, -1)

	return decodeXMLBody(cleaned)
}

func GetEnDailyForDate(date time.Time) (report *CurrencyReport, err error) {
	s, err := getXML(EN_URL + "?date_req=" + date.Format("02/01/2006"))
	if err != nil {
		return nil, err
	}

	converted := strings.Replace(s, ` encoding="windows-1252"`, ` encoding="utf-8"`, -1)

	return decodeXMLBody(converted)
}
