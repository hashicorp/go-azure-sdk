package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOsVppAppVppTokenAccountType string

const (
	MacOsVppAppVppTokenAccountTypebusiness  MacOsVppAppVppTokenAccountType = "Business"
	MacOsVppAppVppTokenAccountTypeeducation MacOsVppAppVppTokenAccountType = "Education"
)

func PossibleValuesForMacOsVppAppVppTokenAccountType() []string {
	return []string{
		string(MacOsVppAppVppTokenAccountTypebusiness),
		string(MacOsVppAppVppTokenAccountTypeeducation),
	}
}

func parseMacOsVppAppVppTokenAccountType(input string) (*MacOsVppAppVppTokenAccountType, error) {
	vals := map[string]MacOsVppAppVppTokenAccountType{
		"business":  MacOsVppAppVppTokenAccountTypebusiness,
		"education": MacOsVppAppVppTokenAccountTypeeducation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOsVppAppVppTokenAccountType(input)
	return &out, nil
}
