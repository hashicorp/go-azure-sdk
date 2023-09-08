package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileKeySize string

const (
	IosScepCertificateProfileKeySizesize1024 IosScepCertificateProfileKeySize = "Size1024"
	IosScepCertificateProfileKeySizesize2048 IosScepCertificateProfileKeySize = "Size2048"
	IosScepCertificateProfileKeySizesize4096 IosScepCertificateProfileKeySize = "Size4096"
)

func PossibleValuesForIosScepCertificateProfileKeySize() []string {
	return []string{
		string(IosScepCertificateProfileKeySizesize1024),
		string(IosScepCertificateProfileKeySizesize2048),
		string(IosScepCertificateProfileKeySizesize4096),
	}
}

func parseIosScepCertificateProfileKeySize(input string) (*IosScepCertificateProfileKeySize, error) {
	vals := map[string]IosScepCertificateProfileKeySize{
		"size1024": IosScepCertificateProfileKeySizesize1024,
		"size2048": IosScepCertificateProfileKeySizesize2048,
		"size4096": IosScepCertificateProfileKeySizesize4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileKeySize(input)
	return &out, nil
}
