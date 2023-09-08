package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileCertificateStore string

const (
	AndroidWorkProfileScepCertificateProfileCertificateStoremachine AndroidWorkProfileScepCertificateProfileCertificateStore = "Machine"
	AndroidWorkProfileScepCertificateProfileCertificateStoreuser    AndroidWorkProfileScepCertificateProfileCertificateStore = "User"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileCertificateStore() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileCertificateStoremachine),
		string(AndroidWorkProfileScepCertificateProfileCertificateStoreuser),
	}
}

func parseAndroidWorkProfileScepCertificateProfileCertificateStore(input string) (*AndroidWorkProfileScepCertificateProfileCertificateStore, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileCertificateStore{
		"machine": AndroidWorkProfileScepCertificateProfileCertificateStoremachine,
		"user":    AndroidWorkProfileScepCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileCertificateStore(input)
	return &out, nil
}
