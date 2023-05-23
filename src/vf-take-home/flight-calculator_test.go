/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package vf_take_home

import (
	"reflect"
	"testing"
)

func TestCalculatePath(t *testing.T) {
	type args struct {
		flights []Flight
	}
	tests := []struct {
		name string
		args args
		want *Flight
	}{
		{
			name: "Single flight test",
			args: args{
				flights: []Flight{
					{
						Start: "SFO",
						End:   "EWR",
					},
				},
			},
			want: &Flight{
				Start: "SFO",
				End:   "EWR",
			},
		},
		{
			name: "Two flights test",
			args: args{
				flights: []Flight{
					{
						Start: "ATL",
						End:   "EWR",
					},
					{
						Start: "SFO",
						End:   "ATL",
					},
				},
			},
			want: &Flight{
				Start: "SFO",
				End:   "EWR",
			},
		},
		{
			name: "Multiple flights test",
			args: args{
				flights: []Flight{
					{
						Start: "IND",
						End:   "EWR",
					},
					{
						Start: "SFO",
						End:   "ATL",
					},
					{
						Start: "GSO",
						End:   "IND",
					},
					{
						Start: "ATL",
						End:   "GSO",
					},
				},
			},
			want: &Flight{
				Start: "SFO",
				End:   "EWR",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatePath(tt.args.flights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculatePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
