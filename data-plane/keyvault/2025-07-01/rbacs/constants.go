package rbacs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataAction string

const (
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashbackupSlashstartSlashaction                 DataAction = "Microsoft.KeyVault/managedHsm/backup/start/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashbackupSlashstatusSlashaction                DataAction = "Microsoft.KeyVault/managedHsm/backup/status/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashbackupSlashaction                  DataAction = "Microsoft.KeyVault/managedHsm/keys/backup/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashcreate                             DataAction = "Microsoft.KeyVault/managedHsm/keys/create"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdecryptSlashaction                 DataAction = "Microsoft.KeyVault/managedHsm/keys/decrypt/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdelete                             DataAction = "Microsoft.KeyVault/managedHsm/keys/delete"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdeletedKeysSlashdelete             DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/delete"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdeletedKeysSlashreadSlashaction    DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/read/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdeletedKeysSlashrecoverSlashaction DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/recover/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashencryptSlashaction                 DataAction = "Microsoft.KeyVault/managedHsm/keys/encrypt/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashexportSlashaction                  DataAction = "Microsoft.KeyVault/managedHsm/keys/export/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashimportSlashaction                  DataAction = "Microsoft.KeyVault/managedHsm/keys/import/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashreadSlashaction                    DataAction = "Microsoft.KeyVault/managedHsm/keys/read/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashreleaseSlashaction                 DataAction = "Microsoft.KeyVault/managedHsm/keys/release/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashrestoreSlashaction                 DataAction = "Microsoft.KeyVault/managedHsm/keys/restore/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashsignSlashaction                    DataAction = "Microsoft.KeyVault/managedHsm/keys/sign/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashunwrapSlashaction                  DataAction = "Microsoft.KeyVault/managedHsm/keys/unwrap/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashverifySlashaction                  DataAction = "Microsoft.KeyVault/managedHsm/keys/verify/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashwrapSlashaction                    DataAction = "Microsoft.KeyVault/managedHsm/keys/wrap/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashwriteSlashaction                   DataAction = "Microsoft.KeyVault/managedHsm/keys/write/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashrestoreSlashstartSlashaction                DataAction = "Microsoft.KeyVault/managedHsm/restore/start/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashrestoreSlashstatusSlashaction               DataAction = "Microsoft.KeyVault/managedHsm/restore/status/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashrngSlashaction                              DataAction = "Microsoft.KeyVault/managedHsm/rng/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleAssignmentsSlashdeleteSlashaction       DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/delete/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleAssignmentsSlashreadSlashaction         DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/read/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleAssignmentsSlashwriteSlashaction        DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/write/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleDefinitionsSlashdeleteSlashaction       DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/delete/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleDefinitionsSlashreadSlashaction         DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/read/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleDefinitionsSlashwriteSlashaction        DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/write/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashdownloadSlashaction      DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/download/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashdownloadSlashread        DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/download/read"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashtransferkeySlashread     DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/transferkey/read"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashuploadSlashaction        DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/upload/action"
	DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashuploadSlashread          DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/upload/read"
)

func PossibleValuesForDataAction() []string {
	return []string{
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashbackupSlashstartSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashbackupSlashstatusSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashbackupSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashcreate),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdecryptSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdelete),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdeletedKeysSlashdelete),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdeletedKeysSlashreadSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdeletedKeysSlashrecoverSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashencryptSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashexportSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashimportSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashreadSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashreleaseSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashrestoreSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashsignSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashunwrapSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashverifySlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashwrapSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashwriteSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashrestoreSlashstartSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashrestoreSlashstatusSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashrngSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleAssignmentsSlashdeleteSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleAssignmentsSlashreadSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleAssignmentsSlashwriteSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleDefinitionsSlashdeleteSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleDefinitionsSlashreadSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleDefinitionsSlashwriteSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashdownloadSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashdownloadSlashread),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashtransferkeySlashread),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashuploadSlashaction),
		string(DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashuploadSlashread),
	}
}

func (s *DataAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataAction(input string) (*DataAction, error) {
	vals := map[string]DataAction{
		"microsoft.keyvault/managedhsm/backup/start/action":             DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashbackupSlashstartSlashaction,
		"microsoft.keyvault/managedhsm/backup/status/action":            DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashbackupSlashstatusSlashaction,
		"microsoft.keyvault/managedhsm/keys/backup/action":              DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashbackupSlashaction,
		"microsoft.keyvault/managedhsm/keys/create":                     DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashcreate,
		"microsoft.keyvault/managedhsm/keys/decrypt/action":             DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdecryptSlashaction,
		"microsoft.keyvault/managedhsm/keys/delete":                     DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdelete,
		"microsoft.keyvault/managedhsm/keys/deletedkeys/delete":         DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdeletedKeysSlashdelete,
		"microsoft.keyvault/managedhsm/keys/deletedkeys/read/action":    DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdeletedKeysSlashreadSlashaction,
		"microsoft.keyvault/managedhsm/keys/deletedkeys/recover/action": DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashdeletedKeysSlashrecoverSlashaction,
		"microsoft.keyvault/managedhsm/keys/encrypt/action":             DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashencryptSlashaction,
		"microsoft.keyvault/managedhsm/keys/export/action":              DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashexportSlashaction,
		"microsoft.keyvault/managedhsm/keys/import/action":              DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashimportSlashaction,
		"microsoft.keyvault/managedhsm/keys/read/action":                DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashreadSlashaction,
		"microsoft.keyvault/managedhsm/keys/release/action":             DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashreleaseSlashaction,
		"microsoft.keyvault/managedhsm/keys/restore/action":             DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashrestoreSlashaction,
		"microsoft.keyvault/managedhsm/keys/sign/action":                DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashsignSlashaction,
		"microsoft.keyvault/managedhsm/keys/unwrap/action":              DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashunwrapSlashaction,
		"microsoft.keyvault/managedhsm/keys/verify/action":              DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashverifySlashaction,
		"microsoft.keyvault/managedhsm/keys/wrap/action":                DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashwrapSlashaction,
		"microsoft.keyvault/managedhsm/keys/write/action":               DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashkeysSlashwriteSlashaction,
		"microsoft.keyvault/managedhsm/restore/start/action":            DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashrestoreSlashstartSlashaction,
		"microsoft.keyvault/managedhsm/restore/status/action":           DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashrestoreSlashstatusSlashaction,
		"microsoft.keyvault/managedhsm/rng/action":                      DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashrngSlashaction,
		"microsoft.keyvault/managedhsm/roleassignments/delete/action":   DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleAssignmentsSlashdeleteSlashaction,
		"microsoft.keyvault/managedhsm/roleassignments/read/action":     DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleAssignmentsSlashreadSlashaction,
		"microsoft.keyvault/managedhsm/roleassignments/write/action":    DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleAssignmentsSlashwriteSlashaction,
		"microsoft.keyvault/managedhsm/roledefinitions/delete/action":   DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleDefinitionsSlashdeleteSlashaction,
		"microsoft.keyvault/managedhsm/roledefinitions/read/action":     DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleDefinitionsSlashreadSlashaction,
		"microsoft.keyvault/managedhsm/roledefinitions/write/action":    DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashroleDefinitionsSlashwriteSlashaction,
		"microsoft.keyvault/managedhsm/securitydomain/download/action":  DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashdownloadSlashaction,
		"microsoft.keyvault/managedhsm/securitydomain/download/read":    DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashdownloadSlashread,
		"microsoft.keyvault/managedhsm/securitydomain/transferkey/read": DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashtransferkeySlashread,
		"microsoft.keyvault/managedhsm/securitydomain/upload/action":    DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashuploadSlashaction,
		"microsoft.keyvault/managedhsm/securitydomain/upload/read":      DataActionMicrosoftPointKeyVaultSlashmanagedHsmSlashsecuritydomainSlashuploadSlashread,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataAction(input)
	return &out, nil
}

type RoleDefinitionType string

const (
	RoleDefinitionTypeMicrosoftPointAuthorizationSlashroleDefinitions RoleDefinitionType = "Microsoft.Authorization/roleDefinitions"
)

func PossibleValuesForRoleDefinitionType() []string {
	return []string{
		string(RoleDefinitionTypeMicrosoftPointAuthorizationSlashroleDefinitions),
	}
}

func (s *RoleDefinitionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleDefinitionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleDefinitionType(input string) (*RoleDefinitionType, error) {
	vals := map[string]RoleDefinitionType{
		"microsoft.authorization/roledefinitions": RoleDefinitionTypeMicrosoftPointAuthorizationSlashroleDefinitions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleDefinitionType(input)
	return &out, nil
}

type RoleScope string

const (
	RoleScopeSlash     RoleScope = "/"
	RoleScopeSlashkeys RoleScope = "/keys"
)

func PossibleValuesForRoleScope() []string {
	return []string{
		string(RoleScopeSlash),
		string(RoleScopeSlashkeys),
	}
}

func (s *RoleScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleScope(input string) (*RoleScope, error) {
	vals := map[string]RoleScope{
		"/":     RoleScopeSlash,
		"/keys": RoleScopeSlashkeys,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleScope(input)
	return &out, nil
}

type RoleType string

const (
	RoleTypeAKVBuiltInRole RoleType = "AKVBuiltInRole"
	RoleTypeCustomRole     RoleType = "CustomRole"
)

func PossibleValuesForRoleType() []string {
	return []string{
		string(RoleTypeAKVBuiltInRole),
		string(RoleTypeCustomRole),
	}
}

func (s *RoleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleType(input string) (*RoleType, error) {
	vals := map[string]RoleType{
		"akvbuiltinrole": RoleTypeAKVBuiltInRole,
		"customrole":     RoleTypeCustomRole,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleType(input)
	return &out, nil
}
