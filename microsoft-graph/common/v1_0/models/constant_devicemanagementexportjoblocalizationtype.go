package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExportJobLocalizationType string

const (
	DeviceManagementExportJobLocalizationTypelocalizedValuesAsAdditionalColumn DeviceManagementExportJobLocalizationType = "LocalizedValuesAsAdditionalColumn"
	DeviceManagementExportJobLocalizationTypereplaceLocalizableValues          DeviceManagementExportJobLocalizationType = "ReplaceLocalizableValues"
)

func PossibleValuesForDeviceManagementExportJobLocalizationType() []string {
	return []string{
		string(DeviceManagementExportJobLocalizationTypelocalizedValuesAsAdditionalColumn),
		string(DeviceManagementExportJobLocalizationTypereplaceLocalizableValues),
	}
}

func parseDeviceManagementExportJobLocalizationType(input string) (*DeviceManagementExportJobLocalizationType, error) {
	vals := map[string]DeviceManagementExportJobLocalizationType{
		"localizedvaluesasadditionalcolumn": DeviceManagementExportJobLocalizationTypelocalizedValuesAsAdditionalColumn,
		"replacelocalizablevalues":          DeviceManagementExportJobLocalizationTypereplaceLocalizableValues,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExportJobLocalizationType(input)
	return &out, nil
}
