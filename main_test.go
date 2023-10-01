package main

import "testing"

func Test_disguise(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				str: "Here is my spammy page: http://testlink see you",
			},
			want: "Here is my spammy page: http://******** see you",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := disguise(tt.args.str); got != tt.want {
				t.Errorf("disguise() = %v, want %v", got, tt.want)
			}
		})
	}
}
