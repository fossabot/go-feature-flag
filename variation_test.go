package ffclient

import (
	"github.com/google/go-cmp/cmp"
	"testing"

	"github.com/thomaspoignant/go-feature-flag/ffuser"
	"github.com/thomaspoignant/go-feature-flag/internal/cache"
	"github.com/thomaspoignant/go-feature-flag/internal/flags"
)

func TestBoolVariation(t *testing.T) {
	flagCacheMock := map[string]flags.Flag{
		"test-flag": {
			Rule:       "anonymous eq true",
			Percentage: 50,
			True:       true,
			False:      false,
			Default:    true,
		},
	}

	type args struct {
		flagKey      string
		user         ffuser.User
		defaultValue bool
		flagCache    map[string]flags.Flag
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Get error when not init",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: true,
				flagCache:    nil,
			},
			wantErr: true,
		},
		{
			name: "Get default value with key not exist",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: true,
				flagCache:    flagCacheMock,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Get default value, rule not apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key"),
				defaultValue: true,
				flagCache:    flagCacheMock,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Get true value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key"),
				defaultValue: true,
				flagCache:    flagCacheMock,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Get false value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key-ssss1"),
				defaultValue: true,
				flagCache:    flagCacheMock,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Get default value, when rule apply and not right type",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key-ssss1"),
				defaultValue: true,
				flagCache: map[string]flags.Flag{
					"test-flag": {
						Rule:       "anonymous eq true",
						Percentage: 50,
						True:       "yyy",
						False:      "xxx",
						Default:    "zzz",
					},
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.FlagsCache = tt.args.flagCache
			got, err := BoolVariation(tt.args.flagKey, tt.args.user, tt.args.defaultValue)
			cache.FlagsCache = nil

			if (err != nil) != tt.wantErr {
				t.Errorf("BoolVariation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BoolVariation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Variation(t *testing.T) {
	flagCacheMock := map[string]flags.Flag{
		"test-flag": {
			Rule:       "anonymous eq true",
			Percentage: 50,
			True:       120.0,
			False:      121.0,
			Default:    119.0,
		},
	}

	type args struct {
		flagKey      string
		user         ffuser.User
		defaultValue float64
		flagCache    map[string]flags.Flag
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Get error when not init",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: 118.0,
				flagCache:    nil,
			},
			wantErr: true,
		},
		{
			name: "Get default value with key not exist",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: 118.0,
				flagCache:    flagCacheMock,
			},
			want:    118.0,
			wantErr: false,
		},
		{
			name: "Get default value, rule not apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key"),
				defaultValue: 118.0,
				flagCache:    flagCacheMock,
			},
			want:    119.0,
			wantErr: false,
		},
		{
			name: "Get true value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key"),
				defaultValue: 118.0,
				flagCache:    flagCacheMock,
			},
			want:    120.0,
			wantErr: false,
		},
		{
			name: "Get false value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key-ssss1"),
				defaultValue: 118.0,
				flagCache:    flagCacheMock,
			},
			want:    121.0,
			wantErr: false,
		},
		{
			name: "Get default value, when rule apply and not right type",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key-ssss1"),
				defaultValue: 118.0,
				flagCache: map[string]flags.Flag{
					"test-flag": {
						Rule:       "anonymous eq true",
						Percentage: 50,
						True:       "yyy",
						False:      "xxx",
						Default:    "zzz",
					},
				},
			},
			want:    118.0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.FlagsCache = tt.args.flagCache
			got, err := Float64Variation(tt.args.flagKey, tt.args.user, tt.args.defaultValue)
			cache.FlagsCache = nil

			if (err != nil) != tt.wantErr {
				t.Errorf("Float64Variation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Float64Variation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArrayVariation(t *testing.T) {
	flagCacheMock := map[string]flags.Flag{
		"test-flag": {
			Rule:       "anonymous eq true",
			Percentage: 50,
			True:       []interface{}{"true"},
			False:      []interface{}{"false"},
			Default:    []interface{}{"default"},
		},
	}

	type args struct {
		flagKey      string
		user         ffuser.User
		defaultValue []interface{}
		flagCache    map[string]flags.Flag
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "Get error when not init",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: []interface{}{"toto"},
				flagCache:    nil,
			},
			wantErr: true,
		},
		{
			name: "Get default value with key not exist",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: []interface{}{"toto"},
				flagCache:    flagCacheMock,
			},
			want:    []interface{}{"toto"},
			wantErr: false,
		},
		{
			name: "Get default value, rule not apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key"),
				defaultValue: []interface{}{"toto"},
				flagCache:    flagCacheMock,
			},
			want:    []interface{}{"default"},
			wantErr: false,
		},
		{
			name: "Get true value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key"),
				defaultValue: []interface{}{"toto"},
				flagCache:    flagCacheMock,
			},
			want:    []interface{}{"true"},
			wantErr: false,
		},
		{
			name: "Get false value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key-ssss1"),
				defaultValue: []interface{}{"toto"},
				flagCache:    flagCacheMock,
			},
			want:    []interface{}{"false"},
			wantErr: false,
		},
		{
			name: "Get default value, when rule apply and not right type",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key-ssss1"),
				defaultValue: []interface{}{"toto"},
				flagCache: map[string]flags.Flag{
					"test-flag": {
						Rule:       "anonymous eq true",
						Percentage: 50,
						True:       "yyy",
						False:      "xxx",
						Default:    "zzz",
					},
				},
			},
			want:    []interface{}{"toto"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.FlagsCache = tt.args.flagCache
			got, err := JSONArrayVariation(tt.args.flagKey, tt.args.user, tt.args.defaultValue)
			cache.FlagsCache = nil

			if (err != nil) != tt.wantErr {
				t.Errorf("JSONArrayVariation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("JSONArrayVariation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONVariation(t *testing.T) {
	flagCacheMock := map[string]flags.Flag{
		"test-flag": {
			Rule:       "anonymous eq true",
			Percentage: 50,
			True:       map[string]interface{}{"true": true},
			False:      map[string]interface{}{"false": true},
			Default:    map[string]interface{}{"default": true},
		},
	}

	type args struct {
		flagKey      string
		user         ffuser.User
		defaultValue map[string]interface{}
		flagCache    map[string]flags.Flag
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Get error when not init",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: map[string]interface{}{"default-notkey": true},
				flagCache:    nil,
			},
			wantErr: true,
		},
		{
			name: "Get default value with key not exist",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: map[string]interface{}{"default-notkey": true},
				flagCache:    flagCacheMock,
			},
			want:    map[string]interface{}{"default-notkey": true},
			wantErr: false,
		},
		{
			name: "Get default value, rule not apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key"),
				defaultValue: map[string]interface{}{"default-notkey": true},
				flagCache:    flagCacheMock,
			},
			want:    map[string]interface{}{"default": true},
			wantErr: false,
		},
		{
			name: "Get true value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key"),
				defaultValue: map[string]interface{}{"default-notkey": true},
				flagCache:    flagCacheMock,
			},
			want:    map[string]interface{}{"true": true},
			wantErr: false,
		},
		{
			name: "Get false value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key-ssss1"),
				defaultValue: map[string]interface{}{"default-notkey": true},
				flagCache:    flagCacheMock,
			},
			want:    map[string]interface{}{"false": true},
			wantErr: false,
		},
		{
			name: "Get default value, when rule apply and not right type",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key-ssss1"),
				defaultValue: map[string]interface{}{"default-notkey": true},
				flagCache: map[string]flags.Flag{
					"test-flag": {
						Rule:       "anonymous eq true",
						Percentage: 50,
						True:       1,
						False:      2,
						Default:    3,
					},
				},
			},
			want:    map[string]interface{}{"default-notkey": true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.FlagsCache = tt.args.flagCache
			got, err := JSONVariation(tt.args.flagKey, tt.args.user, tt.args.defaultValue)
			cache.FlagsCache = nil

			if (err != nil) != tt.wantErr {
				t.Errorf("JSONVariation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("JSONVariation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringVariation(t *testing.T) {
	flagCacheMock := map[string]flags.Flag{
		"test-flag": {
			Rule:       "anonymous eq true",
			Percentage: 50,
			True:       "true",
			False:      "false",
			Default:    "default",
		},
	}

	type args struct {
		flagKey      string
		user         ffuser.User
		defaultValue string
		flagCache    map[string]flags.Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get error when not init",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: "default-notkey",
				flagCache:    nil,
			},
			wantErr: true,
		},
		{
			name: "Get default value with key not exist",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: "default-notkey",
				flagCache:    flagCacheMock,
			},
			want:    "default-notkey",
			wantErr: false,
		},
		{
			name: "Get default value, rule not apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key"),
				defaultValue: "default-notkey",
				flagCache:    flagCacheMock,
			},
			want:    "default",
			wantErr: false,
		},
		{
			name: "Get true value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key"),
				defaultValue: "default-notkey",
				flagCache:    flagCacheMock,
			},
			want:    "true",
			wantErr: false,
		},
		{
			name: "Get false value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key-ssss1"),
				defaultValue: "default-notkey",
				flagCache:    flagCacheMock,
			},
			want:    "false",
			wantErr: false,
		},
		{
			name: "Get default value, when rule apply and not right type",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key-ssss1"),
				defaultValue: "default-notkey",
				flagCache: map[string]flags.Flag{
					"test-flag": {
						Rule:       "anonymous eq true",
						Percentage: 50,
						True:       1,
						False:      2,
						Default:    3,
					},
				},
			},
			want:    "default-notkey",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.FlagsCache = tt.args.flagCache
			got, err := StringVariation(tt.args.flagKey, tt.args.user, tt.args.defaultValue)
			cache.FlagsCache = nil

			if (err != nil) != tt.wantErr {
				t.Errorf("StringVariation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringVariation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntVariation(t *testing.T) {
	flagCacheMock := map[string]flags.Flag{
		"test-flag": {
			Rule:       "anonymous eq true",
			Percentage: 50,
			True:       120,
			False:      121,
			Default:    119,
		},
	}

	type args struct {
		flagKey      string
		user         ffuser.User
		defaultValue int
		flagCache    map[string]flags.Flag
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Get error when not init",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: 118,
				flagCache:    nil,
			},
			wantErr: true,
		},
		{
			name: "Get default value with key not exist",
			args: args{
				flagKey:      "key-not-exist",
				user:         ffuser.NewUser("random-key"),
				defaultValue: 118,
				flagCache:    flagCacheMock,
			},
			want:    118,
			wantErr: false,
		},
		{
			name: "Get default value rule not apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key"),
				defaultValue: 118,
				flagCache:    flagCacheMock,
			},
			want:    119,
			wantErr: false,
		},
		{
			name: "Get true value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key"),
				defaultValue: 118,
				flagCache:    flagCacheMock,
			},
			want:    120,
			wantErr: false,
		},
		{
			name: "Get false value, rule apply",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewAnonymousUser("random-key-ssss1"),
				defaultValue: 118,
				flagCache:    flagCacheMock,
			},
			want:    121,
			wantErr: false,
		},
		{
			name: "Get default value, when rule apply and not right type",
			args: args{
				flagKey:      "test-flag",
				user:         ffuser.NewUser("random-key-ssss1"),
				defaultValue: 118,
				flagCache: map[string]flags.Flag{
					"test-flag": {
						Rule:       "anonymous eq true",
						Percentage: 50,
						True:       "yyy",
						False:      "xxx",
						Default:    "zzz",
					},
				},
			},
			want:    118,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.FlagsCache = tt.args.flagCache
			got, err := IntVariation(tt.args.flagKey, tt.args.user, tt.args.defaultValue)
			cache.FlagsCache = nil

			if (err != nil) != tt.wantErr {
				t.Errorf("IntVariation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntVariation() got = %v, want %v", got, tt.want)
			}
		})
	}
}
