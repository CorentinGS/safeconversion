package safeconversion_test

import (
	"testing"

	"github.com/corentings/safeconversion"
)

func TestSafeParseInt64(t *testing.T) {
	tests := []struct {
		value   string
		want    int64
		wantErr bool
	}{
		{
			value:   "123",
			want:    123,
			wantErr: false,
		},
		{
			value:   "9223372036854775807",
			want:    9223372036854775807,
			wantErr: false,
		},
		{
			value:   "-9223372036854775808",
			want:    -9223372036854775808,
			wantErr: false,
		},
		{
			value:   "9223372036854775808",
			wantErr: true,
		},
		{
			value:   "-9223372036854775809",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got, err := safeconversion.SafeParse[int64](tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SafeParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SafeParse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeParseInt32(t *testing.T) {
	tests := []struct {
		value   string
		want    int32
		wantErr bool
	}{
		{
			value:   "123",
			want:    123,
			wantErr: false,
		},
		{
			value:   "2147483647",
			want:    2147483647,
			wantErr: false,
		},
		{
			value:   "-2147483648",
			want:    -2147483648,
			wantErr: false,
		},
		{
			value:   "2147483648",
			wantErr: true,
		},
		{
			value:   "-2147483649",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got, err := safeconversion.SafeParse[int32](tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SafeParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SafeParse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
