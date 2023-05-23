/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package vf_take_home

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ComputePath(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var inputFlights []Flight
	inputFlights, err = unmarshal(body)
	if err != nil {
		http.Error(w, "Error computing flight", http.StatusInternalServerError)
		return
	}

	path := CalculatePath(inputFlights)
	if path == nil {
		http.Error(w, "Error computing flight", http.StatusInternalServerError)
		return
	}
	responseData, err := json.Marshal([]string{path.Start, path.End})
	if err != nil {
		http.Error(w, "Error marshalling response data", http.StatusInternalServerError)
		return
	}

	// Write the status code and response data to the http.ResponseWriter
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
