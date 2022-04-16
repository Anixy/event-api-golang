package helpers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicIfError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "SAME ORROR TEST",
			args: args{
				err: errors.New("something wrong"),
			},
		},
		{
			name: "NO ERROR TEST",
			args: args{
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.err != nil {
				assert.Panics(t, func() {
					PanicIfError(tt.args.err)
				})
			}else{
				assert.NotPanics(t, func() {
					PanicIfError(tt.args.err)
				})
			}
		})
	}
}
