package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagementMicrosoftStoreForBusinessPortalSelection string

const (
	DeviceAppManagementMicrosoftStoreForBusinessPortalSelectioncompanyPortal DeviceAppManagementMicrosoftStoreForBusinessPortalSelection = "CompanyPortal"
	DeviceAppManagementMicrosoftStoreForBusinessPortalSelectionnone          DeviceAppManagementMicrosoftStoreForBusinessPortalSelection = "None"
	DeviceAppManagementMicrosoftStoreForBusinessPortalSelectionprivateStore  DeviceAppManagementMicrosoftStoreForBusinessPortalSelection = "PrivateStore"
)

func PossibleValuesForDeviceAppManagementMicrosoftStoreForBusinessPortalSelection() []string {
	return []string{
		string(DeviceAppManagementMicrosoftStoreForBusinessPortalSelectioncompanyPortal),
		string(DeviceAppManagementMicrosoftStoreForBusinessPortalSelectionnone),
		string(DeviceAppManagementMicrosoftStoreForBusinessPortalSelectionprivateStore),
	}
}

func parseDeviceAppManagementMicrosoftStoreForBusinessPortalSelection(input string) (*DeviceAppManagementMicrosoftStoreForBusinessPortalSelection, error) {
	vals := map[string]DeviceAppManagementMicrosoftStoreForBusinessPortalSelection{
		"companyportal": DeviceAppManagementMicrosoftStoreForBusinessPortalSelectioncompanyPortal,
		"none":          DeviceAppManagementMicrosoftStoreForBusinessPortalSelectionnone,
		"privatestore":  DeviceAppManagementMicrosoftStoreForBusinessPortalSelectionprivateStore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAppManagementMicrosoftStoreForBusinessPortalSelection(input)
	return &out, nil
}
