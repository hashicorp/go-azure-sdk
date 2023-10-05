package endpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialType string

const (
	CredentialTypeAzureKeyVaultSmb CredentialType = "AzureKeyVaultSmb"
)

func PossibleValuesForCredentialType() []string {
	return []string{
		string(CredentialTypeAzureKeyVaultSmb),
	}
}

func (s *CredentialType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCredentialType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCredentialType(input string) (*CredentialType, error) {
	vals := map[string]CredentialType{
		"azurekeyvaultsmb": CredentialTypeAzureKeyVaultSmb,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CredentialType(input)
	return &out, nil
}

type EndpointType string

const (
	EndpointTypeAzureStorageBlobContainer EndpointType = "AzureStorageBlobContainer"
	EndpointTypeAzureStorageSmbFileShare  EndpointType = "AzureStorageSmbFileShare"
	EndpointTypeNfsMount                  EndpointType = "NfsMount"
	EndpointTypeSmbMount                  EndpointType = "SmbMount"
)

func PossibleValuesForEndpointType() []string {
	return []string{
		string(EndpointTypeAzureStorageBlobContainer),
		string(EndpointTypeAzureStorageSmbFileShare),
		string(EndpointTypeNfsMount),
		string(EndpointTypeSmbMount),
	}
}

func (s *EndpointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointType(input string) (*EndpointType, error) {
	vals := map[string]EndpointType{
		"azurestorageblobcontainer": EndpointTypeAzureStorageBlobContainer,
		"azurestoragesmbfileshare":  EndpointTypeAzureStorageSmbFileShare,
		"nfsmount":                  EndpointTypeNfsMount,
		"smbmount":                  EndpointTypeSmbMount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointType(input)
	return &out, nil
}

type NfsVersion string

const (
	NfsVersionNFSauto   NfsVersion = "NFSauto"
	NfsVersionNFSvFour  NfsVersion = "NFSv4"
	NfsVersionNFSvThree NfsVersion = "NFSv3"
)

func PossibleValuesForNfsVersion() []string {
	return []string{
		string(NfsVersionNFSauto),
		string(NfsVersionNFSvFour),
		string(NfsVersionNFSvThree),
	}
}

func (s *NfsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNfsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNfsVersion(input string) (*NfsVersion, error) {
	vals := map[string]NfsVersion{
		"nfsauto": NfsVersionNFSauto,
		"nfsv4":   NfsVersionNFSvFour,
		"nfsv3":   NfsVersionNFSvThree,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NfsVersion(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
