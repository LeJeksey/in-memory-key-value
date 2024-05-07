package compute

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_analyzer_Analyze(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name    string
		args    args
		want    Query
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "error: empty query",
			args: args{
				tokens: []string{},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "error: unknown command",
			args: args{
				tokens: []string{"unknown", "x", "y"},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "error: set query with invalid number of arguments",
			args: args{
				tokens: []string{"set", "x"},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "error: get query with invalid number of arguments",
			args: args{
				tokens: []string{"get"},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "error: delete query with invalid number of arguments",
			args: args{
				tokens: []string{"delete"},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "success: set query",
			args: args{
				tokens: []string{"set", "x", "y"},
			},
			want:    &SetQuery{Key: "x", Val: "y"},
			wantErr: assert.NoError,
		},
		{
			name: "success: get query",
			args: args{
				tokens: []string{"get", "x"},
			},
			want:    &GetQuery{Key: "x"},
			wantErr: assert.NoError,
		},
		{
			name: "success: delete query",
			args: args{
				tokens: []string{"delete", "x"},
			},
			want:    &DeleteQuery{Key: "x"},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAnalyzer()
			got, err := a.Analyze(tt.args.tokens)
			if !tt.wantErr(t, err, fmt.Sprintf("Analyze(%v)", tt.args.tokens)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Analyze(%v)", tt.args.tokens)
		})
	}
}
