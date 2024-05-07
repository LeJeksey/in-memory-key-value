package compute

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stateMachine_parse(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "empty query",
			args: args{
				query: "",
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "success",
			args: args{
				query: "set some word",
			},
			want:    []string{"set", "some", "word"},
			wantErr: assert.NoError,
		},
		{
			name: "unknown symbols",
			args: args{
				query: "set some&^%$# word",
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newStateMachine()
			got, err := m.parse(tt.args.query)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
