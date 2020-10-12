// Copyright Frontware International
// This package is used by several Frontware projects to handle basic tasks about geo location

package geo

import "testing"

func TestDistance(t *testing.T) {
	type args struct {
		lat1 float64
		lon1 float64
		lat2 float64
		lon2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Office to Office",
			args: args{lat1: 13.7665217, lon1: 100.6068431,
				lat2: 13.7665217, lon2: 100.6068431},
			want: 0,
		},
		{
			name: "Office to BiGC",
			args: args{lat1: 13.7665217, lon1: 100.6068431,
				lat2: 13.7199345, lon2: 100.5197898},
			want: 10747.271299236845,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.lat1, tt.args.lon1, tt.args.lat2, tt.args.lon2); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
