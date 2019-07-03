package golash

import "testing"

func TestIsEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid Email",
			args: args{
				email: "abc@gmail.com",
			},
			want: true,
		},
		{
			name: "Valid Email",
			args: args{
				email: "ab1@gmail.com",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.args.email); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
