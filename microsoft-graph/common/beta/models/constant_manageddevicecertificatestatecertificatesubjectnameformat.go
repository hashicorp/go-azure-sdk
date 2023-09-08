package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateSubjectNameFormat string

const (
	ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonName                  ManagedDeviceCertificateStateCertificateSubjectNameFormat = "CommonName"
	ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsAadDeviceId     ManagedDeviceCertificateStateCertificateSubjectNameFormat = "CommonNameAsAadDeviceId"
	ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsDurableDeviceId ManagedDeviceCertificateStateCertificateSubjectNameFormat = "CommonNameAsDurableDeviceId"
	ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsEmail           ManagedDeviceCertificateStateCertificateSubjectNameFormat = "CommonNameAsEmail"
	ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsIMEI            ManagedDeviceCertificateStateCertificateSubjectNameFormat = "CommonNameAsIMEI"
	ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsIntuneDeviceId  ManagedDeviceCertificateStateCertificateSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsSerialNumber    ManagedDeviceCertificateStateCertificateSubjectNameFormat = "CommonNameAsSerialNumber"
	ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameIncludingEmail    ManagedDeviceCertificateStateCertificateSubjectNameFormat = "CommonNameIncludingEmail"
	ManagedDeviceCertificateStateCertificateSubjectNameFormatcustom                      ManagedDeviceCertificateStateCertificateSubjectNameFormat = "Custom"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateSubjectNameFormat() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonName),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsAadDeviceId),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsDurableDeviceId),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsEmail),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsIMEI),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsSerialNumber),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameIncludingEmail),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormatcustom),
	}
}

func parseManagedDeviceCertificateStateCertificateSubjectNameFormat(input string) (*ManagedDeviceCertificateStateCertificateSubjectNameFormat, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateSubjectNameFormat{
		"commonname":                  ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    ManagedDeviceCertificateStateCertificateSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      ManagedDeviceCertificateStateCertificateSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateSubjectNameFormat(input)
	return &out, nil
}
