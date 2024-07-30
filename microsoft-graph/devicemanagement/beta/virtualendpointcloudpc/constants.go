package virtualendpointcloudpc

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountType string

const (
	CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountTypeAdministrator CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountType = "administrator"
	CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountTypeStandardUser  CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountType = "standardUser"
)

func PossibleValuesForCreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountType() []string {
	return []string{
		string(CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountTypeAdministrator),
		string(CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountTypeStandardUser),
	}
}

func (s *CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountType(input string) (*CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountType, error) {
	vals := map[string]CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountType{
		"administrator": CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountTypeAdministrator,
		"standarduser":  CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountTypeStandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateVirtualEndpointCloudPCChangeUserAccountTypeRequestUserAccountType(input)
	return &out, nil
}

type CreateVirtualEndpointCloudPCReprovisionRequestOsVersion string

const (
	CreateVirtualEndpointCloudPCReprovisionRequestOsVersionWindows10 CreateVirtualEndpointCloudPCReprovisionRequestOsVersion = "windows10"
	CreateVirtualEndpointCloudPCReprovisionRequestOsVersionWindows11 CreateVirtualEndpointCloudPCReprovisionRequestOsVersion = "windows11"
)

func PossibleValuesForCreateVirtualEndpointCloudPCReprovisionRequestOsVersion() []string {
	return []string{
		string(CreateVirtualEndpointCloudPCReprovisionRequestOsVersionWindows10),
		string(CreateVirtualEndpointCloudPCReprovisionRequestOsVersionWindows11),
	}
}

func (s *CreateVirtualEndpointCloudPCReprovisionRequestOsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateVirtualEndpointCloudPCReprovisionRequestOsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateVirtualEndpointCloudPCReprovisionRequestOsVersion(input string) (*CreateVirtualEndpointCloudPCReprovisionRequestOsVersion, error) {
	vals := map[string]CreateVirtualEndpointCloudPCReprovisionRequestOsVersion{
		"windows10": CreateVirtualEndpointCloudPCReprovisionRequestOsVersionWindows10,
		"windows11": CreateVirtualEndpointCloudPCReprovisionRequestOsVersionWindows11,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateVirtualEndpointCloudPCReprovisionRequestOsVersion(input)
	return &out, nil
}

type CreateVirtualEndpointCloudPCReprovisionRequestUserAccountType string

const (
	CreateVirtualEndpointCloudPCReprovisionRequestUserAccountTypeAdministrator CreateVirtualEndpointCloudPCReprovisionRequestUserAccountType = "administrator"
	CreateVirtualEndpointCloudPCReprovisionRequestUserAccountTypeStandardUser  CreateVirtualEndpointCloudPCReprovisionRequestUserAccountType = "standardUser"
)

func PossibleValuesForCreateVirtualEndpointCloudPCReprovisionRequestUserAccountType() []string {
	return []string{
		string(CreateVirtualEndpointCloudPCReprovisionRequestUserAccountTypeAdministrator),
		string(CreateVirtualEndpointCloudPCReprovisionRequestUserAccountTypeStandardUser),
	}
}

func (s *CreateVirtualEndpointCloudPCReprovisionRequestUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateVirtualEndpointCloudPCReprovisionRequestUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateVirtualEndpointCloudPCReprovisionRequestUserAccountType(input string) (*CreateVirtualEndpointCloudPCReprovisionRequestUserAccountType, error) {
	vals := map[string]CreateVirtualEndpointCloudPCReprovisionRequestUserAccountType{
		"administrator": CreateVirtualEndpointCloudPCReprovisionRequestUserAccountTypeAdministrator,
		"standarduser":  CreateVirtualEndpointCloudPCReprovisionRequestUserAccountTypeStandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateVirtualEndpointCloudPCReprovisionRequestUserAccountType(input)
	return &out, nil
}
