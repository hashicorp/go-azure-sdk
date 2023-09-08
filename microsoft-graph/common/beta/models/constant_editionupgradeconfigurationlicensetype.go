package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EditionUpgradeConfigurationLicenseType string

const (
	EditionUpgradeConfigurationLicenseTypelicenseFile   EditionUpgradeConfigurationLicenseType = "LicenseFile"
	EditionUpgradeConfigurationLicenseTypenotConfigured EditionUpgradeConfigurationLicenseType = "NotConfigured"
	EditionUpgradeConfigurationLicenseTypeproductKey    EditionUpgradeConfigurationLicenseType = "ProductKey"
)

func PossibleValuesForEditionUpgradeConfigurationLicenseType() []string {
	return []string{
		string(EditionUpgradeConfigurationLicenseTypelicenseFile),
		string(EditionUpgradeConfigurationLicenseTypenotConfigured),
		string(EditionUpgradeConfigurationLicenseTypeproductKey),
	}
}

func parseEditionUpgradeConfigurationLicenseType(input string) (*EditionUpgradeConfigurationLicenseType, error) {
	vals := map[string]EditionUpgradeConfigurationLicenseType{
		"licensefile":   EditionUpgradeConfigurationLicenseTypelicenseFile,
		"notconfigured": EditionUpgradeConfigurationLicenseTypenotConfigured,
		"productkey":    EditionUpgradeConfigurationLicenseTypeproductKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EditionUpgradeConfigurationLicenseType(input)
	return &out, nil
}
