package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileKeySize string

const (
	AndroidWorkProfileScepCertificateProfileKeySizesize1024 AndroidWorkProfileScepCertificateProfileKeySize = "Size1024"
	AndroidWorkProfileScepCertificateProfileKeySizesize2048 AndroidWorkProfileScepCertificateProfileKeySize = "Size2048"
	AndroidWorkProfileScepCertificateProfileKeySizesize4096 AndroidWorkProfileScepCertificateProfileKeySize = "Size4096"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileKeySize() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileKeySizesize1024),
		string(AndroidWorkProfileScepCertificateProfileKeySizesize2048),
		string(AndroidWorkProfileScepCertificateProfileKeySizesize4096),
	}
}

func parseAndroidWorkProfileScepCertificateProfileKeySize(input string) (*AndroidWorkProfileScepCertificateProfileKeySize, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileKeySize{
		"size1024": AndroidWorkProfileScepCertificateProfileKeySizesize1024,
		"size2048": AndroidWorkProfileScepCertificateProfileKeySizesize2048,
		"size4096": AndroidWorkProfileScepCertificateProfileKeySizesize4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileKeySize(input)
	return &out, nil
}
