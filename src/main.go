/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"log"
	"net/http"
	vf_take_home "volume-finance/src/vf-take-home"
)

func main() {
	http.HandleFunc("/calculate", vf_take_home.ComputePath)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
