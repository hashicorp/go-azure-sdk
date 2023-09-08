package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountEnterpriseSettingsBindStatus string

const (
	AndroidManagedStoreAccountEnterpriseSettingsBindStatusbound             AndroidManagedStoreAccountEnterpriseSettingsBindStatus = "Bound"
	AndroidManagedStoreAccountEnterpriseSettingsBindStatusboundAndValidated AndroidManagedStoreAccountEnterpriseSettingsBindStatus = "BoundAndValidated"
	AndroidManagedStoreAccountEnterpriseSettingsBindStatusnotBound          AndroidManagedStoreAccountEnterpriseSettingsBindStatus = "NotBound"
	AndroidManagedStoreAccountEnterpriseSettingsBindStatusunbinding         AndroidManagedStoreAccountEnterpriseSettingsBindStatus = "Unbinding"
)

func PossibleValuesForAndroidManagedStoreAccountEnterpriseSettingsBindStatus() []string {
	return []string{
		string(AndroidManagedStoreAccountEnterpriseSettingsBindStatusbound),
		string(AndroidManagedStoreAccountEnterpriseSettingsBindStatusboundAndValidated),
		string(AndroidManagedStoreAccountEnterpriseSettingsBindStatusnotBound),
		string(AndroidManagedStoreAccountEnterpriseSettingsBindStatusunbinding),
	}
}

func parseAndroidManagedStoreAccountEnterpriseSettingsBindStatus(input string) (*AndroidManagedStoreAccountEnterpriseSettingsBindStatus, error) {
	vals := map[string]AndroidManagedStoreAccountEnterpriseSettingsBindStatus{
		"bound":             AndroidManagedStoreAccountEnterpriseSettingsBindStatusbound,
		"boundandvalidated": AndroidManagedStoreAccountEnterpriseSettingsBindStatusboundAndValidated,
		"notbound":          AndroidManagedStoreAccountEnterpriseSettingsBindStatusnotBound,
		"unbinding":         AndroidManagedStoreAccountEnterpriseSettingsBindStatusunbinding,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAccountEnterpriseSettingsBindStatus(input)
	return &out, nil
}
