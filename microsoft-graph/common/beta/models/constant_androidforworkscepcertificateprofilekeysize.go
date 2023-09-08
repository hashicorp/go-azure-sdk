package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileKeySize string

const (
	AndroidForWorkScepCertificateProfileKeySizesize1024 AndroidForWorkScepCertificateProfileKeySize = "Size1024"
	AndroidForWorkScepCertificateProfileKeySizesize2048 AndroidForWorkScepCertificateProfileKeySize = "Size2048"
	AndroidForWorkScepCertificateProfileKeySizesize4096 AndroidForWorkScepCertificateProfileKeySize = "Size4096"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileKeySize() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileKeySizesize1024),
		string(AndroidForWorkScepCertificateProfileKeySizesize2048),
		string(AndroidForWorkScepCertificateProfileKeySizesize4096),
	}
}

func parseAndroidForWorkScepCertificateProfileKeySize(input string) (*AndroidForWorkScepCertificateProfileKeySize, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileKeySize{
		"size1024": AndroidForWorkScepCertificateProfileKeySizesize1024,
		"size2048": AndroidForWorkScepCertificateProfileKeySizesize2048,
		"size4096": AndroidForWorkScepCertificateProfileKeySizesize4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileKeySize(input)
	return &out, nil
}
