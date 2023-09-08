package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOsVersion string

const (
	CloudPCOsVersionwindows10 CloudPCOsVersion = "Windows10"
	CloudPCOsVersionwindows11 CloudPCOsVersion = "Windows11"
)

func PossibleValuesForCloudPCOsVersion() []string {
	return []string{
		string(CloudPCOsVersionwindows10),
		string(CloudPCOsVersionwindows11),
	}
}

func parseCloudPCOsVersion(input string) (*CloudPCOsVersion, error) {
	vals := map[string]CloudPCOsVersion{
		"windows10": CloudPCOsVersionwindows10,
		"windows11": CloudPCOsVersionwindows11,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOsVersion(input)
	return &out, nil
}
