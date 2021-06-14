package main

import (
	"strings"
)

func phoneMask(phone, mask string) string {
	phoneSlice := strings.Split(phone, "")
	maskSlice := strings.Split(mask, "")

	var s []string

	if len(phone) == strings.Count(mask, "#") {
		numCounter := 0
		for _, charac := range maskSlice {
			if charac == "#" {
				s = append(s, phoneSlice[numCounter])
				numCounter++
			} else {
				s = append(s, charac)
			}

		}

		return strings.Join(s[:], "")
	}

	return phone
}
