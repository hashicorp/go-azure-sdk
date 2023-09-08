package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsPolicyPolicyMigrationState string

const (
	AuthenticationMethodsPolicyPolicyMigrationStatemigrationComplete   AuthenticationMethodsPolicyPolicyMigrationState = "MigrationComplete"
	AuthenticationMethodsPolicyPolicyMigrationStatemigrationInProgress AuthenticationMethodsPolicyPolicyMigrationState = "MigrationInProgress"
	AuthenticationMethodsPolicyPolicyMigrationStatepreMigration        AuthenticationMethodsPolicyPolicyMigrationState = "PreMigration"
)

func PossibleValuesForAuthenticationMethodsPolicyPolicyMigrationState() []string {
	return []string{
		string(AuthenticationMethodsPolicyPolicyMigrationStatemigrationComplete),
		string(AuthenticationMethodsPolicyPolicyMigrationStatemigrationInProgress),
		string(AuthenticationMethodsPolicyPolicyMigrationStatepreMigration),
	}
}

func parseAuthenticationMethodsPolicyPolicyMigrationState(input string) (*AuthenticationMethodsPolicyPolicyMigrationState, error) {
	vals := map[string]AuthenticationMethodsPolicyPolicyMigrationState{
		"migrationcomplete":   AuthenticationMethodsPolicyPolicyMigrationStatemigrationComplete,
		"migrationinprogress": AuthenticationMethodsPolicyPolicyMigrationStatemigrationInProgress,
		"premigration":        AuthenticationMethodsPolicyPolicyMigrationStatepreMigration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodsPolicyPolicyMigrationState(input)
	return &out, nil
}
