package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLogCollectionRequestTemplateType string

const (
	DeviceLogCollectionRequestTemplateTypepredefined DeviceLogCollectionRequestTemplateType = "Predefined"
)

func PossibleValuesForDeviceLogCollectionRequestTemplateType() []string {
	return []string{
		string(DeviceLogCollectionRequestTemplateTypepredefined),
	}
}

func parseDeviceLogCollectionRequestTemplateType(input string) (*DeviceLogCollectionRequestTemplateType, error) {
	vals := map[string]DeviceLogCollectionRequestTemplateType{
		"predefined": DeviceLogCollectionRequestTemplateTypepredefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceLogCollectionRequestTemplateType(input)
	return &out, nil
}
