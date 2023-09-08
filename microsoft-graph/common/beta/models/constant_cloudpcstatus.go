package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCStatus string

const (
	CloudPCStatusdeprovisioning          CloudPCStatus = "Deprovisioning"
	CloudPCStatusfailed                  CloudPCStatus = "Failed"
	CloudPCStatusinGracePeriod           CloudPCStatus = "InGracePeriod"
	CloudPCStatusmovingRegion            CloudPCStatus = "MovingRegion"
	CloudPCStatusnotProvisioned          CloudPCStatus = "NotProvisioned"
	CloudPCStatuspendingProvision        CloudPCStatus = "PendingProvision"
	CloudPCStatusprovisioned             CloudPCStatus = "Provisioned"
	CloudPCStatusprovisionedWithWarnings CloudPCStatus = "ProvisionedWithWarnings"
	CloudPCStatusprovisioning            CloudPCStatus = "Provisioning"
	CloudPCStatusresizePendingLicense    CloudPCStatus = "ResizePendingLicense"
	CloudPCStatusresizing                CloudPCStatus = "Resizing"
	CloudPCStatusrestoring               CloudPCStatus = "Restoring"
	CloudPCStatusupdatingSingleSignOn    CloudPCStatus = "UpdatingSingleSignOn"
)

func PossibleValuesForCloudPCStatus() []string {
	return []string{
		string(CloudPCStatusdeprovisioning),
		string(CloudPCStatusfailed),
		string(CloudPCStatusinGracePeriod),
		string(CloudPCStatusmovingRegion),
		string(CloudPCStatusnotProvisioned),
		string(CloudPCStatuspendingProvision),
		string(CloudPCStatusprovisioned),
		string(CloudPCStatusprovisionedWithWarnings),
		string(CloudPCStatusprovisioning),
		string(CloudPCStatusresizePendingLicense),
		string(CloudPCStatusresizing),
		string(CloudPCStatusrestoring),
		string(CloudPCStatusupdatingSingleSignOn),
	}
}

func parseCloudPCStatus(input string) (*CloudPCStatus, error) {
	vals := map[string]CloudPCStatus{
		"deprovisioning":          CloudPCStatusdeprovisioning,
		"failed":                  CloudPCStatusfailed,
		"ingraceperiod":           CloudPCStatusinGracePeriod,
		"movingregion":            CloudPCStatusmovingRegion,
		"notprovisioned":          CloudPCStatusnotProvisioned,
		"pendingprovision":        CloudPCStatuspendingProvision,
		"provisioned":             CloudPCStatusprovisioned,
		"provisionedwithwarnings": CloudPCStatusprovisionedWithWarnings,
		"provisioning":            CloudPCStatusprovisioning,
		"resizependinglicense":    CloudPCStatusresizePendingLicense,
		"resizing":                CloudPCStatusresizing,
		"restoring":               CloudPCStatusrestoring,
		"updatingsinglesignon":    CloudPCStatusupdatingSingleSignOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCStatus(input)
	return &out, nil
}
