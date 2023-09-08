package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileCertificateStore string

const (
	IosScepCertificateProfileCertificateStoremachine IosScepCertificateProfileCertificateStore = "Machine"
	IosScepCertificateProfileCertificateStoreuser    IosScepCertificateProfileCertificateStore = "User"
)

func PossibleValuesForIosScepCertificateProfileCertificateStore() []string {
	return []string{
		string(IosScepCertificateProfileCertificateStoremachine),
		string(IosScepCertificateProfileCertificateStoreuser),
	}
}

func parseIosScepCertificateProfileCertificateStore(input string) (*IosScepCertificateProfileCertificateStore, error) {
	vals := map[string]IosScepCertificateProfileCertificateStore{
		"machine": IosScepCertificateProfileCertificateStoremachine,
		"user":    IosScepCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileCertificateStore(input)
	return &out, nil
}
