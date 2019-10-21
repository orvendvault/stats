package stats

import (
	"reflect"
	"testing"
)

func TestGetTStatistic(t *testing.T) {
	type args struct {
		v     float64
		alpha float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"case 1", args{1.0, 0.1}, 3.078},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTStatistic(tt.args.v, tt.args.alpha); got != tt.want {
				t.Errorf("GetTStatistic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStudentsT(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name    string
		args    args
		want    StudentsT
		wantErr bool
	}{
		{"Normal case", args{1.0}, StudentsT{1.0}, false},
		{"Invalid negative case", args{-1.0}, StudentsT{}, true},
		{"Invalid zero case", args{0.0}, StudentsT{}, true},
		{"Invalid real case", args{2.5}, StudentsT{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStudentsT(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStudentsT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudentsT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewStudentsT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewStudentsT(1.0)
	}
}
