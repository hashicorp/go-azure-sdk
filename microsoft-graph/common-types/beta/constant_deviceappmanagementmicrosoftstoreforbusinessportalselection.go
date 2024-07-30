package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagementMicrosoftStoreForBusinessPortalSelection string

const (
	DeviceAppManagementMicrosoftStoreForBusinessPortalSelection_CompanyPortal DeviceAppManagementMicrosoftStoreForBusinessPortalSelection = "companyPortal"
	DeviceAppManagementMicrosoftStoreForBusinessPortalSelection_None          DeviceAppManagementMicrosoftStoreForBusinessPortalSelection = "none"
	DeviceAppManagementMicrosoftStoreForBusinessPortalSelection_PrivateStore  DeviceAppManagementMicrosoftStoreForBusinessPortalSelection = "privateStore"
)

func PossibleValuesForDeviceAppManagementMicrosoftStoreForBusinessPortalSelection() []string {
	return []string{
		string(DeviceAppManagementMicrosoftStoreForBusinessPortalSelection_CompanyPortal),
		string(DeviceAppManagementMicrosoftStoreForBusinessPortalSelection_None),
		string(DeviceAppManagementMicrosoftStoreForBusinessPortalSelection_PrivateStore),
	}
}

func (s *DeviceAppManagementMicrosoftStoreForBusinessPortalSelection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAppManagementMicrosoftStoreForBusinessPortalSelection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAppManagementMicrosoftStoreForBusinessPortalSelection(input string) (*DeviceAppManagementMicrosoftStoreForBusinessPortalSelection, error) {
	vals := map[string]DeviceAppManagementMicrosoftStoreForBusinessPortalSelection{
		"companyportal": DeviceAppManagementMicrosoftStoreForBusinessPortalSelection_CompanyPortal,
		"none":          DeviceAppManagementMicrosoftStoreForBusinessPortalSelection_None,
		"privatestore":  DeviceAppManagementMicrosoftStoreForBusinessPortalSelection_PrivateStore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAppManagementMicrosoftStoreForBusinessPortalSelection(input)
	return &out, nil
}
