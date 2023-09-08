package mecloudpc

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountType string

const (
	CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountTypeadministrator CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountType = "Administrator"
	CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountTypestandardUser  CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountType = "StandardUser"
)

func PossibleValuesForCreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountType() []string {
	return []string{
		string(CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountTypeadministrator),
		string(CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountTypestandardUser),
	}
}

func (s *CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountType(input string) (*CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountType, error) {
	vals := map[string]CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountType{
		"administrator": CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountTypeadministrator,
		"standarduser":  CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountTypestandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeCloudPCByIdChangeUserAccountTypeRequestUserAccountType(input)
	return &out, nil
}

type CreateMeCloudPCByIdReprovisionRequestOsVersion string

const (
	CreateMeCloudPCByIdReprovisionRequestOsVersionwindows10 CreateMeCloudPCByIdReprovisionRequestOsVersion = "Windows10"
	CreateMeCloudPCByIdReprovisionRequestOsVersionwindows11 CreateMeCloudPCByIdReprovisionRequestOsVersion = "Windows11"
)

func PossibleValuesForCreateMeCloudPCByIdReprovisionRequestOsVersion() []string {
	return []string{
		string(CreateMeCloudPCByIdReprovisionRequestOsVersionwindows10),
		string(CreateMeCloudPCByIdReprovisionRequestOsVersionwindows11),
	}
}

func (s *CreateMeCloudPCByIdReprovisionRequestOsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeCloudPCByIdReprovisionRequestOsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeCloudPCByIdReprovisionRequestOsVersion(input string) (*CreateMeCloudPCByIdReprovisionRequestOsVersion, error) {
	vals := map[string]CreateMeCloudPCByIdReprovisionRequestOsVersion{
		"windows10": CreateMeCloudPCByIdReprovisionRequestOsVersionwindows10,
		"windows11": CreateMeCloudPCByIdReprovisionRequestOsVersionwindows11,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeCloudPCByIdReprovisionRequestOsVersion(input)
	return &out, nil
}

type CreateMeCloudPCByIdReprovisionRequestUserAccountType string

const (
	CreateMeCloudPCByIdReprovisionRequestUserAccountTypeadministrator CreateMeCloudPCByIdReprovisionRequestUserAccountType = "Administrator"
	CreateMeCloudPCByIdReprovisionRequestUserAccountTypestandardUser  CreateMeCloudPCByIdReprovisionRequestUserAccountType = "StandardUser"
)

func PossibleValuesForCreateMeCloudPCByIdReprovisionRequestUserAccountType() []string {
	return []string{
		string(CreateMeCloudPCByIdReprovisionRequestUserAccountTypeadministrator),
		string(CreateMeCloudPCByIdReprovisionRequestUserAccountTypestandardUser),
	}
}

func (s *CreateMeCloudPCByIdReprovisionRequestUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeCloudPCByIdReprovisionRequestUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeCloudPCByIdReprovisionRequestUserAccountType(input string) (*CreateMeCloudPCByIdReprovisionRequestUserAccountType, error) {
	vals := map[string]CreateMeCloudPCByIdReprovisionRequestUserAccountType{
		"administrator": CreateMeCloudPCByIdReprovisionRequestUserAccountTypeadministrator,
		"standarduser":  CreateMeCloudPCByIdReprovisionRequestUserAccountTypestandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeCloudPCByIdReprovisionRequestUserAccountType(input)
	return &out, nil
}
