package cloudpc

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCloudPCChangeUserAccountTypeRequestUserAccountType string

const (
	CreateCloudPCChangeUserAccountTypeRequestUserAccountTypeAdministrator CreateCloudPCChangeUserAccountTypeRequestUserAccountType = "administrator"
	CreateCloudPCChangeUserAccountTypeRequestUserAccountTypeStandardUser  CreateCloudPCChangeUserAccountTypeRequestUserAccountType = "standardUser"
)

func PossibleValuesForCreateCloudPCChangeUserAccountTypeRequestUserAccountType() []string {
	return []string{
		string(CreateCloudPCChangeUserAccountTypeRequestUserAccountTypeAdministrator),
		string(CreateCloudPCChangeUserAccountTypeRequestUserAccountTypeStandardUser),
	}
}

func (s *CreateCloudPCChangeUserAccountTypeRequestUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateCloudPCChangeUserAccountTypeRequestUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateCloudPCChangeUserAccountTypeRequestUserAccountType(input string) (*CreateCloudPCChangeUserAccountTypeRequestUserAccountType, error) {
	vals := map[string]CreateCloudPCChangeUserAccountTypeRequestUserAccountType{
		"administrator": CreateCloudPCChangeUserAccountTypeRequestUserAccountTypeAdministrator,
		"standarduser":  CreateCloudPCChangeUserAccountTypeRequestUserAccountTypeStandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateCloudPCChangeUserAccountTypeRequestUserAccountType(input)
	return &out, nil
}

type CreateCloudPCReprovisionRequestOsVersion string

const (
	CreateCloudPCReprovisionRequestOsVersionWindows10 CreateCloudPCReprovisionRequestOsVersion = "windows10"
	CreateCloudPCReprovisionRequestOsVersionWindows11 CreateCloudPCReprovisionRequestOsVersion = "windows11"
)

func PossibleValuesForCreateCloudPCReprovisionRequestOsVersion() []string {
	return []string{
		string(CreateCloudPCReprovisionRequestOsVersionWindows10),
		string(CreateCloudPCReprovisionRequestOsVersionWindows11),
	}
}

func (s *CreateCloudPCReprovisionRequestOsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateCloudPCReprovisionRequestOsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateCloudPCReprovisionRequestOsVersion(input string) (*CreateCloudPCReprovisionRequestOsVersion, error) {
	vals := map[string]CreateCloudPCReprovisionRequestOsVersion{
		"windows10": CreateCloudPCReprovisionRequestOsVersionWindows10,
		"windows11": CreateCloudPCReprovisionRequestOsVersionWindows11,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateCloudPCReprovisionRequestOsVersion(input)
	return &out, nil
}

type CreateCloudPCReprovisionRequestUserAccountType string

const (
	CreateCloudPCReprovisionRequestUserAccountTypeAdministrator CreateCloudPCReprovisionRequestUserAccountType = "administrator"
	CreateCloudPCReprovisionRequestUserAccountTypeStandardUser  CreateCloudPCReprovisionRequestUserAccountType = "standardUser"
)

func PossibleValuesForCreateCloudPCReprovisionRequestUserAccountType() []string {
	return []string{
		string(CreateCloudPCReprovisionRequestUserAccountTypeAdministrator),
		string(CreateCloudPCReprovisionRequestUserAccountTypeStandardUser),
	}
}

func (s *CreateCloudPCReprovisionRequestUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateCloudPCReprovisionRequestUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateCloudPCReprovisionRequestUserAccountType(input string) (*CreateCloudPCReprovisionRequestUserAccountType, error) {
	vals := map[string]CreateCloudPCReprovisionRequestUserAccountType{
		"administrator": CreateCloudPCReprovisionRequestUserAccountTypeAdministrator,
		"standarduser":  CreateCloudPCReprovisionRequestUserAccountTypeStandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateCloudPCReprovisionRequestUserAccountType(input)
	return &out, nil
}
