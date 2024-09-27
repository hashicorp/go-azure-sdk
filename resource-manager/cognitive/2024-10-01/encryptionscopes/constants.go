package encryptionscopes

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionScopeProvisioningState string

const (
	EncryptionScopeProvisioningStateAccepted  EncryptionScopeProvisioningState = "Accepted"
	EncryptionScopeProvisioningStateCanceled  EncryptionScopeProvisioningState = "Canceled"
	EncryptionScopeProvisioningStateCreating  EncryptionScopeProvisioningState = "Creating"
	EncryptionScopeProvisioningStateDeleting  EncryptionScopeProvisioningState = "Deleting"
	EncryptionScopeProvisioningStateFailed    EncryptionScopeProvisioningState = "Failed"
	EncryptionScopeProvisioningStateMoving    EncryptionScopeProvisioningState = "Moving"
	EncryptionScopeProvisioningStateSucceeded EncryptionScopeProvisioningState = "Succeeded"
)

func PossibleValuesForEncryptionScopeProvisioningState() []string {
	return []string{
		string(EncryptionScopeProvisioningStateAccepted),
		string(EncryptionScopeProvisioningStateCanceled),
		string(EncryptionScopeProvisioningStateCreating),
		string(EncryptionScopeProvisioningStateDeleting),
		string(EncryptionScopeProvisioningStateFailed),
		string(EncryptionScopeProvisioningStateMoving),
		string(EncryptionScopeProvisioningStateSucceeded),
	}
}

func (s *EncryptionScopeProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptionScopeProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptionScopeProvisioningState(input string) (*EncryptionScopeProvisioningState, error) {
	vals := map[string]EncryptionScopeProvisioningState{
		"accepted":  EncryptionScopeProvisioningStateAccepted,
		"canceled":  EncryptionScopeProvisioningStateCanceled,
		"creating":  EncryptionScopeProvisioningStateCreating,
		"deleting":  EncryptionScopeProvisioningStateDeleting,
		"failed":    EncryptionScopeProvisioningStateFailed,
		"moving":    EncryptionScopeProvisioningStateMoving,
		"succeeded": EncryptionScopeProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionScopeProvisioningState(input)
	return &out, nil
}

type EncryptionScopeState string

const (
	EncryptionScopeStateDisabled EncryptionScopeState = "Disabled"
	EncryptionScopeStateEnabled  EncryptionScopeState = "Enabled"
)

func PossibleValuesForEncryptionScopeState() []string {
	return []string{
		string(EncryptionScopeStateDisabled),
		string(EncryptionScopeStateEnabled),
	}
}

func (s *EncryptionScopeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptionScopeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptionScopeState(input string) (*EncryptionScopeState, error) {
	vals := map[string]EncryptionScopeState{
		"disabled": EncryptionScopeStateDisabled,
		"enabled":  EncryptionScopeStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionScopeState(input)
	return &out, nil
}

type KeySource string

const (
	KeySourceMicrosoftPointCognitiveServices KeySource = "Microsoft.CognitiveServices"
	KeySourceMicrosoftPointKeyVault          KeySource = "Microsoft.KeyVault"
)

func PossibleValuesForKeySource() []string {
	return []string{
		string(KeySourceMicrosoftPointCognitiveServices),
		string(KeySourceMicrosoftPointKeyVault),
	}
}

func (s *KeySource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKeySource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKeySource(input string) (*KeySource, error) {
	vals := map[string]KeySource{
		"microsoft.cognitiveservices": KeySourceMicrosoftPointCognitiveServices,
		"microsoft.keyvault":          KeySourceMicrosoftPointKeyVault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeySource(input)
	return &out, nil
}
