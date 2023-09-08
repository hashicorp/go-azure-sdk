package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCGalleryImageStatus string

const (
	CloudPCGalleryImageStatusnotSupported         CloudPCGalleryImageStatus = "NotSupported"
	CloudPCGalleryImageStatussupported            CloudPCGalleryImageStatus = "Supported"
	CloudPCGalleryImageStatussupportedWithWarning CloudPCGalleryImageStatus = "SupportedWithWarning"
)

func PossibleValuesForCloudPCGalleryImageStatus() []string {
	return []string{
		string(CloudPCGalleryImageStatusnotSupported),
		string(CloudPCGalleryImageStatussupported),
		string(CloudPCGalleryImageStatussupportedWithWarning),
	}
}

func parseCloudPCGalleryImageStatus(input string) (*CloudPCGalleryImageStatus, error) {
	vals := map[string]CloudPCGalleryImageStatus{
		"notsupported":         CloudPCGalleryImageStatusnotSupported,
		"supported":            CloudPCGalleryImageStatussupported,
		"supportedwithwarning": CloudPCGalleryImageStatussupportedWithWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCGalleryImageStatus(input)
	return &out, nil
}
