package usercloudpc

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountType string

const (
	CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountTypeadministrator CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountType = "Administrator"
	CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountTypestandardUser  CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountType = "StandardUser"
)

func PossibleValuesForCreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountType() []string {
	return []string{
		string(CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountTypeadministrator),
		string(CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountTypestandardUser),
	}
}

func (s *CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountType(input string) (*CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountType, error) {
	vals := map[string]CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountType{
		"administrator": CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountTypeadministrator,
		"standarduser":  CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountTypestandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdCloudPCByIdChangeUserAccountTypeRequestUserAccountType(input)
	return &out, nil
}

type CreateUserByIdCloudPCByIdReprovisionRequestOsVersion string

const (
	CreateUserByIdCloudPCByIdReprovisionRequestOsVersionwindows10 CreateUserByIdCloudPCByIdReprovisionRequestOsVersion = "Windows10"
	CreateUserByIdCloudPCByIdReprovisionRequestOsVersionwindows11 CreateUserByIdCloudPCByIdReprovisionRequestOsVersion = "Windows11"
)

func PossibleValuesForCreateUserByIdCloudPCByIdReprovisionRequestOsVersion() []string {
	return []string{
		string(CreateUserByIdCloudPCByIdReprovisionRequestOsVersionwindows10),
		string(CreateUserByIdCloudPCByIdReprovisionRequestOsVersionwindows11),
	}
}

func (s *CreateUserByIdCloudPCByIdReprovisionRequestOsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdCloudPCByIdReprovisionRequestOsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdCloudPCByIdReprovisionRequestOsVersion(input string) (*CreateUserByIdCloudPCByIdReprovisionRequestOsVersion, error) {
	vals := map[string]CreateUserByIdCloudPCByIdReprovisionRequestOsVersion{
		"windows10": CreateUserByIdCloudPCByIdReprovisionRequestOsVersionwindows10,
		"windows11": CreateUserByIdCloudPCByIdReprovisionRequestOsVersionwindows11,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdCloudPCByIdReprovisionRequestOsVersion(input)
	return &out, nil
}

type CreateUserByIdCloudPCByIdReprovisionRequestUserAccountType string

const (
	CreateUserByIdCloudPCByIdReprovisionRequestUserAccountTypeadministrator CreateUserByIdCloudPCByIdReprovisionRequestUserAccountType = "Administrator"
	CreateUserByIdCloudPCByIdReprovisionRequestUserAccountTypestandardUser  CreateUserByIdCloudPCByIdReprovisionRequestUserAccountType = "StandardUser"
)

func PossibleValuesForCreateUserByIdCloudPCByIdReprovisionRequestUserAccountType() []string {
	return []string{
		string(CreateUserByIdCloudPCByIdReprovisionRequestUserAccountTypeadministrator),
		string(CreateUserByIdCloudPCByIdReprovisionRequestUserAccountTypestandardUser),
	}
}

func (s *CreateUserByIdCloudPCByIdReprovisionRequestUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdCloudPCByIdReprovisionRequestUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdCloudPCByIdReprovisionRequestUserAccountType(input string) (*CreateUserByIdCloudPCByIdReprovisionRequestUserAccountType, error) {
	vals := map[string]CreateUserByIdCloudPCByIdReprovisionRequestUserAccountType{
		"administrator": CreateUserByIdCloudPCByIdReprovisionRequestUserAccountTypeadministrator,
		"standarduser":  CreateUserByIdCloudPCByIdReprovisionRequestUserAccountTypestandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdCloudPCByIdReprovisionRequestUserAccountType(input)
	return &out, nil
}
