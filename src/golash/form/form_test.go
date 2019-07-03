package golash

import (
	"testing"
)

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

func TestIsName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid Name",
			args: args{
				name: "Peter son",
			},
			want: true,
		},
		{
			name: "Valid Name",
			args: args{
				name: "Peter",
			},
			want: true,
		},
		{
			name: "InValid Name",
			args: args{
				name: "p23@ter",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsName(tt.args.name); got != tt.want {
				t.Errorf("IsName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPhone(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0718967888",
			args: args{
				number: "0718967888",
			},
			want: true,
		},
		{
			name: "+045366254",
			args: args{
				number: "+045366254",
			},
			want: true,
		},
		{
			name: "+045-366254",
			args: args{
				number: "+045-366254",
			},
			want: true,
		},
		{
			name: "045-366254",
			args: args{
				number: "045-366254",
			},
			want: true,
		},
		{
			name: "045*366254",
			args: args{
				number: "045*366254",
			},
			want: false,
		},
		{
			name: "045anc366254",
			args: args{
				number: "045anc366254",
			},
			want: false,
		},
		{
			name: "abchdyg",
			args: args{
				number: "abchdyg",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPhone(tt.args.number); got != tt.want {
				t.Errorf("IsPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}
