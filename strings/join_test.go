package strings

import "testing"

func TestJoinInt64Slice(t *testing.T) {
	type args struct {
		slice []int64
		sep   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil",
			args: args{slice: nil, sep: ","},
			want: "",
		},
		{
			name: "[1,2,3,4]",
			args: args{slice: []int64{1, 2, 3, 4}, sep: ","},
			want: "1,2,3,4",
		},
		{
			name: "empty",
			args: args{slice: []int64{}, sep: ","},
			want: "",
		},
		{
			name: "one",
			args: args{slice: []int64{1}, sep: ","},
			want: "1",
		},
		{
			name: "comma",
			args: args{slice: []int64{1, 2, 3, 4, 5}, sep: ";"},
			want: "1;2;3;4;5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinInt64Slice(tt.args.slice, tt.args.sep); got != tt.want {
				t.Errorf("JoinInt64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
