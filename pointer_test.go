package pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPointerOf(t *testing.T) {
	t.Parallel()

	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "success: returns pointer of v",
			args: args{
				v: 1,
			},
			want: func(v int) *int {
				return &v
			}(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Of(tt.args.v)
			require.NotNil(t, got)
			assert.Equal(t, *tt.want, *got)
		})
	}
}

func TestValueOrDefault(t *testing.T) {
	t.Parallel()

	type args struct {
		p        *int
		_default int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success: returns value",
			args: args{
				p: func(v int) *int {
					return &v
				}(1),
				_default: 3,
			},
			want: 1,
		},
		{
			name: "success: if p is nil, returns default",
			args: args{
				p:        nil,
				_default: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValueOrDefault(tt.args.p, tt.args._default)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestValueOrEmpty(t *testing.T) {
	t.Parallel()

	type args struct {
		p *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success: returns value",
			args: args{
				p: func(v int) *int {
					return &v
				}(1),
			},
			want: 1,
		},
		{
			name: "success: if p is nil, returns empty",
			args: args{
				p: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValueOrEmpty(tt.args.p)
			assert.Equal(t, tt.want, got)
		})
	}
}
