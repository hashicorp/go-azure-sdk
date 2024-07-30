package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EditionUpgradeConfigurationLicenseType string

const (
	EditionUpgradeConfigurationLicenseType_LicenseFile   EditionUpgradeConfigurationLicenseType = "licenseFile"
	EditionUpgradeConfigurationLicenseType_NotConfigured EditionUpgradeConfigurationLicenseType = "notConfigured"
	EditionUpgradeConfigurationLicenseType_ProductKey    EditionUpgradeConfigurationLicenseType = "productKey"
)

func PossibleValuesForEditionUpgradeConfigurationLicenseType() []string {
	return []string{
		string(EditionUpgradeConfigurationLicenseType_LicenseFile),
		string(EditionUpgradeConfigurationLicenseType_NotConfigured),
		string(EditionUpgradeConfigurationLicenseType_ProductKey),
	}
}

func (s *EditionUpgradeConfigurationLicenseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEditionUpgradeConfigurationLicenseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEditionUpgradeConfigurationLicenseType(input string) (*EditionUpgradeConfigurationLicenseType, error) {
	vals := map[string]EditionUpgradeConfigurationLicenseType{
		"licensefile":   EditionUpgradeConfigurationLicenseType_LicenseFile,
		"notconfigured": EditionUpgradeConfigurationLicenseType_NotConfigured,
		"productkey":    EditionUpgradeConfigurationLicenseType_ProductKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EditionUpgradeConfigurationLicenseType(input)
	return &out, nil
}
