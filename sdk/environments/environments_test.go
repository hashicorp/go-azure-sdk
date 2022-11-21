package environments_test

import (
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"testing"
)

func TestFromNamed(t *testing.T) {
	type testCase struct {
		envName string
		result  *environments.Environment
		fail    bool
	}
	testCases := []testCase{
		{"global", &environments.Global, false},
		{"public", &environments.Public, false},
		{"usgovernment", &environments.USGovernmentL4, false},
		{"usgovernmentl4", &environments.USGovernmentL4, false},
		{"dod", &environments.USGovernmentL5, false},
		{"usgovernmentl5", &environments.USGovernmentL5, false},
		{"canary", &environments.Canary, false},
		{"china", &environments.China, false},
		{"mars", nil, true},
	}

	for _, c := range testCases {
		env, err := environments.FromNamed(c.envName)
		if c.fail {
			if err != nil {
				continue
			}

			t.Fatalf("test case %q: expected an error but didn't get one", c.envName)
		}
		if err != nil {
			t.Fatalf("test case %q: unexpected error: %+v", c.envName, err)
		}
		if *env != *c.result {
			t.Fatalf("test case %q: unexpected value received.\nexpected: %#v\nreceived: %#v\n", c.envName, c.result, env)
		}
	}
}
