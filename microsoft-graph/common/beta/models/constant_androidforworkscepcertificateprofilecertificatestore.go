package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileCertificateStore string

const (
	AndroidForWorkScepCertificateProfileCertificateStoremachine AndroidForWorkScepCertificateProfileCertificateStore = "Machine"
	AndroidForWorkScepCertificateProfileCertificateStoreuser    AndroidForWorkScepCertificateProfileCertificateStore = "User"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileCertificateStore() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileCertificateStoremachine),
		string(AndroidForWorkScepCertificateProfileCertificateStoreuser),
	}
}

func parseAndroidForWorkScepCertificateProfileCertificateStore(input string) (*AndroidForWorkScepCertificateProfileCertificateStore, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileCertificateStore{
		"machine": AndroidForWorkScepCertificateProfileCertificateStoremachine,
		"user":    AndroidForWorkScepCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileCertificateStore(input)
	return &out, nil
}
