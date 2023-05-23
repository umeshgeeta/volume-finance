/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package vf_take_home

import (
	"encoding/json"
)

type Flight struct {
	Start string
	End   string
}

func CalculatePath(flights []Flight) *Flight {
	if len(flights) > 0 {
		candidateStarts := make(map[string]bool)
		candidateDestinations := make(map[string]bool)
		for _, f := range flights {
			_, present := candidateStarts[f.Start]
			if !present {
				candidateStarts[f.Start] = true
			}
			candidateDestinations[f.Start] = false
			_, present = candidateDestinations[f.End]
			if !present {
				candidateDestinations[f.End] = true
			}
			candidateStarts[f.End] = false
		}
		// find the start
		start := flights[0].Start
		for airport, val := range candidateStarts {
			if val {
				start = airport
			}
		}
		// find the destination
		end := flights[0].End
		for airport, val := range candidateDestinations {
			if val {
				end = airport
			}
		}
		result := Flight{
			Start: start,
			End:   end,
		}
		return &result
	}
	return nil
}

func unmarshal(body []byte) ([]Flight, error) {
	var tempData [][]string
	err := json.Unmarshal(body, &tempData)
	if err != nil {
		return nil, err
	}

	var flights []Flight
	for _, v := range tempData {
		flights = append(flights, Flight{Start: v[0], End: v[1]})
	}

	return flights, nil
}
