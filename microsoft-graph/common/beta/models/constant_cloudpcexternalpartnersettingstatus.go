package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExternalPartnerSettingStatus string

const (
	CloudPCExternalPartnerSettingStatusavailable    CloudPCExternalPartnerSettingStatus = "Available"
	CloudPCExternalPartnerSettingStatushealthy      CloudPCExternalPartnerSettingStatus = "Healthy"
	CloudPCExternalPartnerSettingStatusnotAvailable CloudPCExternalPartnerSettingStatus = "NotAvailable"
	CloudPCExternalPartnerSettingStatusunhealthy    CloudPCExternalPartnerSettingStatus = "Unhealthy"
)

func PossibleValuesForCloudPCExternalPartnerSettingStatus() []string {
	return []string{
		string(CloudPCExternalPartnerSettingStatusavailable),
		string(CloudPCExternalPartnerSettingStatushealthy),
		string(CloudPCExternalPartnerSettingStatusnotAvailable),
		string(CloudPCExternalPartnerSettingStatusunhealthy),
	}
}

func parseCloudPCExternalPartnerSettingStatus(input string) (*CloudPCExternalPartnerSettingStatus, error) {
	vals := map[string]CloudPCExternalPartnerSettingStatus{
		"available":    CloudPCExternalPartnerSettingStatusavailable,
		"healthy":      CloudPCExternalPartnerSettingStatushealthy,
		"notavailable": CloudPCExternalPartnerSettingStatusnotAvailable,
		"unhealthy":    CloudPCExternalPartnerSettingStatusunhealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCExternalPartnerSettingStatus(input)
	return &out, nil
}
