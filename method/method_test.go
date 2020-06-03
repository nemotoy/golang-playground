package method

import "testing"

func TestUser_updateName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		{"", &User{}, args{}, true},
		{"", &User{}, args{"hoge"}, false},
		{"", nil, args{"hoge"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.updateName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("User.updateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
