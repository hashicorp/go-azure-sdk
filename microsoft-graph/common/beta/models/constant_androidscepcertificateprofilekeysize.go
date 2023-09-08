package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileKeySize string

const (
	AndroidScepCertificateProfileKeySizesize1024 AndroidScepCertificateProfileKeySize = "Size1024"
	AndroidScepCertificateProfileKeySizesize2048 AndroidScepCertificateProfileKeySize = "Size2048"
	AndroidScepCertificateProfileKeySizesize4096 AndroidScepCertificateProfileKeySize = "Size4096"
)

func PossibleValuesForAndroidScepCertificateProfileKeySize() []string {
	return []string{
		string(AndroidScepCertificateProfileKeySizesize1024),
		string(AndroidScepCertificateProfileKeySizesize2048),
		string(AndroidScepCertificateProfileKeySizesize4096),
	}
}

func parseAndroidScepCertificateProfileKeySize(input string) (*AndroidScepCertificateProfileKeySize, error) {
	vals := map[string]AndroidScepCertificateProfileKeySize{
		"size1024": AndroidScepCertificateProfileKeySizesize1024,
		"size2048": AndroidScepCertificateProfileKeySizesize2048,
		"size4096": AndroidScepCertificateProfileKeySizesize4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileKeySize(input)
	return &out, nil
}
