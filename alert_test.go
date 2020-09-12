package cui

import (
	"errors"
	"testing"

	"github.com/gookit/color"
)

func init() {
	color.Disable()
}

func Test_alertSuccess(t *testing.T) {
	type args struct {
		message string
		details []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty", args{},
			`√ 
`,
		},
		{"Only Success Message", args{message: "Done without any error 1"},
			`√ Done without any error 1
`,
		},
		{"Success Message with Empty Detail",
			args{message: "Done without any error 2", details: []string{}},
			`√ Done without any error 2
`,
		},
		{"Success Message with Single Detail",
			args{message: "Done without any error 3", details: []string{"REF_NO:12"}},
			`√ Done without any error 3
- REF_NO:12
`,
		},
		{"Success Message with Multiple Details",
			args{message: "Done without any error 4", details: []string{"REF_NO:13", "See: ABCd"}},
			`√ Done without any error 4
- REF_NO:13
- See: ABCd
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alertSuccess(tt.args.message, tt.args.details...); got != tt.want {
				t.Error("Got:", got, "Want:", tt.want)
			}
		})
	}
}

func Test_alertError(t *testing.T) {
	type args struct {
		message string
		err     []error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty", args{},
			`X 
`,
		},
		{"Only Error Message", args{message: "Done with an error 1"},
			`X Done with an error 1
`,
		},
		{"Error Message with Empty Error",
			args{message: "Done with an error 2", err: []error{}},
			`X Done with an error 2
`,
		},
		{"Error Message with Single Error",
			args{message: "Done with an error 3", err: []error{errors.New("REF_NO:12")}},
			`X Done with an error 3
  REF_NO:12
`,
		},
		{"Error Message with Multiple Errors",
			args{message: "Done with some errors 4", err: []error{errors.New("REF_NO:13"), errors.New("See: ABCd")}},
			`X Done with some errors 4
  REF_NO:13
  See: ABCd
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alertError(tt.args.message, tt.args.err...); got != tt.want {
				t.Error("Got:", got, "Want:", tt.want)
			}
		})
	}
}

func Test_alertWarning(t *testing.T) {
	type args struct {
		message string
		details []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty", args{},
			`! 
`,
		},
		{"Only Warning Message", args{message: "Done with warning 1"},
			`! Done with warning 1
`,
		},
		{"Warning Message with Empty Detail",
			args{message: "Done with warning 2", details: []string{}},
			`! Done with warning 2
`,
		},
		{"Warning Message with Single Detail",
			args{message: "Done with warning 3", details: []string{"REF_NO:12"}},
			`! Done with warning 3
  REF_NO:12
`,
		},
		{"Warning Message with Multiple Details",
			args{message: "Done with warning 4", details: []string{"REF_NO:13", "See: ABCd"}},
			`! Done with warning 4
  REF_NO:13
  See: ABCd
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alertWarning(tt.args.message, tt.args.details...); got != tt.want {
				t.Error("Got:", got, "Want:", tt.want)
			}
		})
	}
}

func Test_alertInfo(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty", args{},
			`* 
`,
		},
		{"Info Message", args{message: "This is an information"},
			`* This is an information
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alertInfo(tt.args.message); got != tt.want {
				t.Error("Got:", got, "Want:", tt.want)
			}
		})
	}
}
