package timex

import (
	"testing"
	"time"

	"github.com/starudream/go-lib/testx"
)

func TestParse(t *testing.T) {
	x := time.Date(2022, 8, 8, 0, 0, 0, 0, time.Local)
	type args struct {
		v any
		l []*time.Location
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "Unix", args: args{x.Unix(), []*time.Location{}}, want: x},
		{name: "DateTime", args: args{x.Format(DateTimeFormat), []*time.Location{}}, want: x},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.v, tt.args.l...)
			testx.RequireEqualf(t, tt.wantErr, err != nil, "error = %v, wantErr %v", err, tt.wantErr)
			testx.RequireEqualf(t, tt.want, got, "got = %v, want %v", got, tt.want)
		})
	}
}
