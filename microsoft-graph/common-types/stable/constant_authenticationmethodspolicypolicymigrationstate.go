package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsPolicyPolicyMigrationState string

const (
	AuthenticationMethodsPolicyPolicyMigrationState_MigrationComplete   AuthenticationMethodsPolicyPolicyMigrationState = "migrationComplete"
	AuthenticationMethodsPolicyPolicyMigrationState_MigrationInProgress AuthenticationMethodsPolicyPolicyMigrationState = "migrationInProgress"
	AuthenticationMethodsPolicyPolicyMigrationState_PreMigration        AuthenticationMethodsPolicyPolicyMigrationState = "preMigration"
)

func PossibleValuesForAuthenticationMethodsPolicyPolicyMigrationState() []string {
	return []string{
		string(AuthenticationMethodsPolicyPolicyMigrationState_MigrationComplete),
		string(AuthenticationMethodsPolicyPolicyMigrationState_MigrationInProgress),
		string(AuthenticationMethodsPolicyPolicyMigrationState_PreMigration),
	}
}

func (s *AuthenticationMethodsPolicyPolicyMigrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodsPolicyPolicyMigrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodsPolicyPolicyMigrationState(input string) (*AuthenticationMethodsPolicyPolicyMigrationState, error) {
	vals := map[string]AuthenticationMethodsPolicyPolicyMigrationState{
		"migrationcomplete":   AuthenticationMethodsPolicyPolicyMigrationState_MigrationComplete,
		"migrationinprogress": AuthenticationMethodsPolicyPolicyMigrationState_MigrationInProgress,
		"premigration":        AuthenticationMethodsPolicyPolicyMigrationState_PreMigration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodsPolicyPolicyMigrationState(input)
	return &out, nil
}
