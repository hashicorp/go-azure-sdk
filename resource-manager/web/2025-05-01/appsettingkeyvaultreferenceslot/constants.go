package appsettingkeyvaultreferenceslot

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResolveStatus string

const (
	ResolveStatusAccessToKeyVaultDenied ResolveStatus = "AccessToKeyVaultDenied"
	ResolveStatusFetchTimedOut          ResolveStatus = "FetchTimedOut"
	ResolveStatusInitialized            ResolveStatus = "Initialized"
	ResolveStatusInvalidSyntax          ResolveStatus = "InvalidSyntax"
	ResolveStatusMSINotEnabled          ResolveStatus = "MSINotEnabled"
	ResolveStatusOtherReasons           ResolveStatus = "OtherReasons"
	ResolveStatusResolved               ResolveStatus = "Resolved"
	ResolveStatusSecretNotFound         ResolveStatus = "SecretNotFound"
	ResolveStatusSecretVersionNotFound  ResolveStatus = "SecretVersionNotFound"
	ResolveStatusUnauthorizedClient     ResolveStatus = "UnauthorizedClient"
	ResolveStatusVaultNotFound          ResolveStatus = "VaultNotFound"
)

func PossibleValuesForResolveStatus() []string {
	return []string{
		string(ResolveStatusAccessToKeyVaultDenied),
		string(ResolveStatusFetchTimedOut),
		string(ResolveStatusInitialized),
		string(ResolveStatusInvalidSyntax),
		string(ResolveStatusMSINotEnabled),
		string(ResolveStatusOtherReasons),
		string(ResolveStatusResolved),
		string(ResolveStatusSecretNotFound),
		string(ResolveStatusSecretVersionNotFound),
		string(ResolveStatusUnauthorizedClient),
		string(ResolveStatusVaultNotFound),
	}
}

func (s *ResolveStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResolveStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResolveStatus(input string) (*ResolveStatus, error) {
	vals := map[string]ResolveStatus{
		"accesstokeyvaultdenied": ResolveStatusAccessToKeyVaultDenied,
		"fetchtimedout":          ResolveStatusFetchTimedOut,
		"initialized":            ResolveStatusInitialized,
		"invalidsyntax":          ResolveStatusInvalidSyntax,
		"msinotenabled":          ResolveStatusMSINotEnabled,
		"otherreasons":           ResolveStatusOtherReasons,
		"resolved":               ResolveStatusResolved,
		"secretnotfound":         ResolveStatusSecretNotFound,
		"secretversionnotfound":  ResolveStatusSecretVersionNotFound,
		"unauthorizedclient":     ResolveStatusUnauthorizedClient,
		"vaultnotfound":          ResolveStatusVaultNotFound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResolveStatus(input)
	return &out, nil
}

type Source string

const (
	SourceKeyVault Source = "KeyVault"
)

func PossibleValuesForSource() []string {
	return []string{
		string(SourceKeyVault),
	}
}

func (s *Source) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSource(input string) (*Source, error) {
	vals := map[string]Source{
		"keyvault": SourceKeyVault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Source(input)
	return &out, nil
}
