package compute

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -source=compute.go -package=compute -destination=compute_mock.go

func TestComputer_Compute(t *testing.T) {
	type fields struct {
		storageBehaviour  func(storage *MockStorage)
		parserBehaviour   func(parser *MockParser)
		analyzerBehaviour func(analyzer *MockAnalyzer)
	}
	type args struct {
		rawQuery string
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantResult string
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "error: failed to parse query",
			args: args{
				rawQuery: "bad query &^@$@&(",
			},
			fields: fields{
				parserBehaviour: func(parser *MockParser) {
					parser.EXPECT().Parse("bad query &^@$@&(").
						Return(nil, fmt.Errorf("failed to parse query"))
				},
			},
			wantResult: "",
			wantErr:    assert.Error,
		},
		{
			name: "error: failed to analyze query",
			args: args{
				rawQuery: "something x y",
			},
			fields: fields{
				parserBehaviour: func(parser *MockParser) {
					parser.EXPECT().Parse("something x y").
						Return([]string{"something", "x", "y"}, nil)
				},
				analyzerBehaviour: func(analyzer *MockAnalyzer) {
					analyzer.EXPECT().Analyze([]string{"something", "x", "y"}).
						Return(nil, fmt.Errorf("failed to analyze query"))
				},
			},
			wantResult: "",
			wantErr:    assert.Error,
		},
		{
			name: "error: failed to run query",
			args: args{
				rawQuery: "set x y",
			},
			fields: fields{
				parserBehaviour: func(parser *MockParser) {
					parser.EXPECT().Parse("set x y").
						Return([]string{"set", "x", "y"}, nil)
				},
				analyzerBehaviour: func(analyzer *MockAnalyzer) {
					analyzer.EXPECT().Analyze([]string{"set", "x", "y"}).
						Return(&SetQuery{Key: "x", Val: "y"}, nil)
				},
				storageBehaviour: func(storage *MockStorage) {
					storage.EXPECT().Set("x", "y").
						Return(fmt.Errorf("failed to run query"))
				},
			},
			wantResult: "",
			wantErr:    assert.Error,
		},
		{
			name: "success: set query",
			args: args{
				rawQuery: "set x y",
			},
			fields: fields{
				parserBehaviour: func(parser *MockParser) {
					parser.EXPECT().Parse("set x y").
						Return([]string{"set", "x", "y"}, nil)
				},
				analyzerBehaviour: func(analyzer *MockAnalyzer) {
					analyzer.EXPECT().Analyze([]string{"set", "x", "y"}).
						Return(&SetQuery{Key: "x", Val: "y"}, nil)
				},
				storageBehaviour: func(storage *MockStorage) {
					storage.EXPECT().Set("x", "y").
						Return(nil)
				},
			},
			wantResult: "",
			wantErr:    assert.NoError,
		},
		{
			name: "success: get query",
			args: args{
				rawQuery: "get x",
			},
			fields: fields{
				parserBehaviour: func(parser *MockParser) {
					parser.EXPECT().Parse("get x").
						Return([]string{"get", "x"}, nil)
				},
				analyzerBehaviour: func(analyzer *MockAnalyzer) {
					analyzer.EXPECT().Analyze([]string{"get", "x"}).
						Return(&GetQuery{Key: "x"}, nil)
				},
				storageBehaviour: func(storage *MockStorage) {
					storage.EXPECT().Get("x").
						Return("y", nil)
				},
			},
			wantResult: "y",
			wantErr:    assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &Computer{
				storage:  NewMockStorage(ctrl),
				parser:   NewMockParser(ctrl),
				analyzer: NewMockAnalyzer(ctrl),
			}

			if tt.fields.storageBehaviour != nil {
				tt.fields.storageBehaviour(c.storage.(*MockStorage))
			}
			if tt.fields.parserBehaviour != nil {
				tt.fields.parserBehaviour(c.parser.(*MockParser))
			}
			if tt.fields.analyzerBehaviour != nil {
				tt.fields.analyzerBehaviour(c.analyzer.(*MockAnalyzer))
			}

			gotResult, err := c.Compute(tt.args.rawQuery)
			if !tt.wantErr(t, err, fmt.Sprintf("Compute(%v)", tt.args.rawQuery)) {
				return
			}
			assert.Equalf(t, tt.wantResult, gotResult, "Compute(%v)", tt.args.rawQuery)
		})
	}
}
