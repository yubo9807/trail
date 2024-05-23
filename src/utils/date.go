package utils

import (
	"strconv"
	"strings"
	"time"
)

type Date struct {
	Year   int
	Mouth  int
	Day    int
	Hours  int
	Minute int
	Second int
}

func DateToObj(t time.Time) Date {
	return Date{
		Year:   t.Local().Year(),
		Mouth:  int(t.Month()),
		Day:    t.Day(),
		Hours:  t.Hour(),
		Minute: t.Minute(),
		Second: t.Second(),
	}
}

func DateZeroize(num int) string {
	str := "0" + strconv.Itoa(num)
	length := len(str)
	return str[length-2 : length]
}

type DateFormaterParams struct {
	Time     Date
	Formater string
}

func DateFormater(time time.Time, formater string) string {
	if formater == "" {
		formater = "YYYY-MM-DD hh:mm:ss"
	}
	t := DateToObj(time)

	str1 := strings.Replace(formater, "YYYY", strconv.Itoa(t.Year), -1)
	str2 := strings.Replace(str1, "MM", DateZeroize(t.Mouth), -1)
	str3 := strings.Replace(str2, "DD", DateZeroize(t.Day), -1)
	str4 := strings.Replace(str3, "hh", DateZeroize(t.Hours), -1)
	str5 := strings.Replace(str4, "mm", DateZeroize(t.Minute), -1)
	str6 := strings.Replace(str5, "ss", DateZeroize(t.Second), -1)
	return str6
}
