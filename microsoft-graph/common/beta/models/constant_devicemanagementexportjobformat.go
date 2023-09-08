package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExportJobFormat string

const (
	DeviceManagementExportJobFormatcsv  DeviceManagementExportJobFormat = "Csv"
	DeviceManagementExportJobFormatjson DeviceManagementExportJobFormat = "Json"
	DeviceManagementExportJobFormatpdf  DeviceManagementExportJobFormat = "Pdf"
)

func PossibleValuesForDeviceManagementExportJobFormat() []string {
	return []string{
		string(DeviceManagementExportJobFormatcsv),
		string(DeviceManagementExportJobFormatjson),
		string(DeviceManagementExportJobFormatpdf),
	}
}

func parseDeviceManagementExportJobFormat(input string) (*DeviceManagementExportJobFormat, error) {
	vals := map[string]DeviceManagementExportJobFormat{
		"csv":  DeviceManagementExportJobFormatcsv,
		"json": DeviceManagementExportJobFormatjson,
		"pdf":  DeviceManagementExportJobFormatpdf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExportJobFormat(input)
	return &out, nil
}
