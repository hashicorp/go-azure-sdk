package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppVppTokenAccountType string

const (
	IosVppAppVppTokenAccountTypebusiness  IosVppAppVppTokenAccountType = "Business"
	IosVppAppVppTokenAccountTypeeducation IosVppAppVppTokenAccountType = "Education"
)

func PossibleValuesForIosVppAppVppTokenAccountType() []string {
	return []string{
		string(IosVppAppVppTokenAccountTypebusiness),
		string(IosVppAppVppTokenAccountTypeeducation),
	}
}

func parseIosVppAppVppTokenAccountType(input string) (*IosVppAppVppTokenAccountType, error) {
	vals := map[string]IosVppAppVppTokenAccountType{
		"business":  IosVppAppVppTokenAccountTypebusiness,
		"education": IosVppAppVppTokenAccountTypeeducation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVppAppVppTokenAccountType(input)
	return &out, nil
}
