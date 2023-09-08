package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDeviceImageStatusDetails string

const (
	CloudPCDeviceImageStatusDetailsinternalServerError                  CloudPCDeviceImageStatusDetails = "InternalServerError"
	CloudPCDeviceImageStatusDetailsosVersionNotSupported                CloudPCDeviceImageStatusDetails = "OsVersionNotSupported"
	CloudPCDeviceImageStatusDetailspaidSourceImageNotSupport            CloudPCDeviceImageStatusDetails = "PaidSourceImageNotSupport"
	CloudPCDeviceImageStatusDetailssourceImageInvalid                   CloudPCDeviceImageStatusDetails = "SourceImageInvalid"
	CloudPCDeviceImageStatusDetailssourceImageNotFound                  CloudPCDeviceImageStatusDetails = "SourceImageNotFound"
	CloudPCDeviceImageStatusDetailssourceImageNotGeneralized            CloudPCDeviceImageStatusDetails = "SourceImageNotGeneralized"
	CloudPCDeviceImageStatusDetailssourceImageNotSupportCustomizeVMName CloudPCDeviceImageStatusDetails = "SourceImageNotSupportCustomizeVMName"
	CloudPCDeviceImageStatusDetailssourceImageSizeExceedsLimitation     CloudPCDeviceImageStatusDetails = "SourceImageSizeExceedsLimitation"
	CloudPCDeviceImageStatusDetailsvmAlreadyAzureAdjoined               CloudPCDeviceImageStatusDetails = "VmAlreadyAzureAdjoined"
)

func PossibleValuesForCloudPCDeviceImageStatusDetails() []string {
	return []string{
		string(CloudPCDeviceImageStatusDetailsinternalServerError),
		string(CloudPCDeviceImageStatusDetailsosVersionNotSupported),
		string(CloudPCDeviceImageStatusDetailspaidSourceImageNotSupport),
		string(CloudPCDeviceImageStatusDetailssourceImageInvalid),
		string(CloudPCDeviceImageStatusDetailssourceImageNotFound),
		string(CloudPCDeviceImageStatusDetailssourceImageNotGeneralized),
		string(CloudPCDeviceImageStatusDetailssourceImageNotSupportCustomizeVMName),
		string(CloudPCDeviceImageStatusDetailssourceImageSizeExceedsLimitation),
		string(CloudPCDeviceImageStatusDetailsvmAlreadyAzureAdjoined),
	}
}

func parseCloudPCDeviceImageStatusDetails(input string) (*CloudPCDeviceImageStatusDetails, error) {
	vals := map[string]CloudPCDeviceImageStatusDetails{
		"internalservererror":                  CloudPCDeviceImageStatusDetailsinternalServerError,
		"osversionnotsupported":                CloudPCDeviceImageStatusDetailsosVersionNotSupported,
		"paidsourceimagenotsupport":            CloudPCDeviceImageStatusDetailspaidSourceImageNotSupport,
		"sourceimageinvalid":                   CloudPCDeviceImageStatusDetailssourceImageInvalid,
		"sourceimagenotfound":                  CloudPCDeviceImageStatusDetailssourceImageNotFound,
		"sourceimagenotgeneralized":            CloudPCDeviceImageStatusDetailssourceImageNotGeneralized,
		"sourceimagenotsupportcustomizevmname": CloudPCDeviceImageStatusDetailssourceImageNotSupportCustomizeVMName,
		"sourceimagesizeexceedslimitation":     CloudPCDeviceImageStatusDetailssourceImageSizeExceedsLimitation,
		"vmalreadyazureadjoined":               CloudPCDeviceImageStatusDetailsvmAlreadyAzureAdjoined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDeviceImageStatusDetails(input)
	return &out, nil
}
