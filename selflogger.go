package main

import "time"

type resultLogger struct {
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	OkServices  int       `json:"okServices"`
	BadServices int       `json:"badServices"`
}
