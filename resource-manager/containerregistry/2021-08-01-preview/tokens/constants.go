package tokens

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}

type TokenCertificateName string

const (
	TokenCertificateNameCertificateOne TokenCertificateName = "certificate1"
	TokenCertificateNameCertificateTwo TokenCertificateName = "certificate2"
)

func PossibleValuesForTokenCertificateName() []string {
	return []string{
		string(TokenCertificateNameCertificateOne),
		string(TokenCertificateNameCertificateTwo),
	}
}

func (s *TokenCertificateName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTokenCertificateName() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TokenCertificateName(decoded)
	return nil
}

type TokenPasswordName string

const (
	TokenPasswordNamePasswordOne TokenPasswordName = "password1"
	TokenPasswordNamePasswordTwo TokenPasswordName = "password2"
)

func PossibleValuesForTokenPasswordName() []string {
	return []string{
		string(TokenPasswordNamePasswordOne),
		string(TokenPasswordNamePasswordTwo),
	}
}

func (s *TokenPasswordName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTokenPasswordName() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TokenPasswordName(decoded)
	return nil
}

type TokenStatus string

const (
	TokenStatusDisabled TokenStatus = "disabled"
	TokenStatusEnabled  TokenStatus = "enabled"
)

func PossibleValuesForTokenStatus() []string {
	return []string{
		string(TokenStatusDisabled),
		string(TokenStatusEnabled),
	}
}

func (s *TokenStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTokenStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TokenStatus(decoded)
	return nil
}
