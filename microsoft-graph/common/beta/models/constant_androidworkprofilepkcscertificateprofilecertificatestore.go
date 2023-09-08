package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfilePkcsCertificateProfileCertificateStore string

const (
	AndroidWorkProfilePkcsCertificateProfileCertificateStoremachine AndroidWorkProfilePkcsCertificateProfileCertificateStore = "Machine"
	AndroidWorkProfilePkcsCertificateProfileCertificateStoreuser    AndroidWorkProfilePkcsCertificateProfileCertificateStore = "User"
)

func PossibleValuesForAndroidWorkProfilePkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(AndroidWorkProfilePkcsCertificateProfileCertificateStoremachine),
		string(AndroidWorkProfilePkcsCertificateProfileCertificateStoreuser),
	}
}

func parseAndroidWorkProfilePkcsCertificateProfileCertificateStore(input string) (*AndroidWorkProfilePkcsCertificateProfileCertificateStore, error) {
	vals := map[string]AndroidWorkProfilePkcsCertificateProfileCertificateStore{
		"machine": AndroidWorkProfilePkcsCertificateProfileCertificateStoremachine,
		"user":    AndroidWorkProfilePkcsCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfilePkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
