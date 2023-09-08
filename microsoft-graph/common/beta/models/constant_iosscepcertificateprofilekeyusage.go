package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileKeyUsage string

const (
	IosScepCertificateProfileKeyUsagedigitalSignature IosScepCertificateProfileKeyUsage = "DigitalSignature"
	IosScepCertificateProfileKeyUsagekeyEncipherment  IosScepCertificateProfileKeyUsage = "KeyEncipherment"
)

func PossibleValuesForIosScepCertificateProfileKeyUsage() []string {
	return []string{
		string(IosScepCertificateProfileKeyUsagedigitalSignature),
		string(IosScepCertificateProfileKeyUsagekeyEncipherment),
	}
}

func parseIosScepCertificateProfileKeyUsage(input string) (*IosScepCertificateProfileKeyUsage, error) {
	vals := map[string]IosScepCertificateProfileKeyUsage{
		"digitalsignature": IosScepCertificateProfileKeyUsagedigitalSignature,
		"keyencipherment":  IosScepCertificateProfileKeyUsagekeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileKeyUsage(input)
	return &out, nil
}
