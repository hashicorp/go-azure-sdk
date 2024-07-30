package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExternalPartnerSettingStatus string

const (
	CloudPCExternalPartnerSettingStatus_Available    CloudPCExternalPartnerSettingStatus = "available"
	CloudPCExternalPartnerSettingStatus_Healthy      CloudPCExternalPartnerSettingStatus = "healthy"
	CloudPCExternalPartnerSettingStatus_NotAvailable CloudPCExternalPartnerSettingStatus = "notAvailable"
	CloudPCExternalPartnerSettingStatus_Unhealthy    CloudPCExternalPartnerSettingStatus = "unhealthy"
)

func PossibleValuesForCloudPCExternalPartnerSettingStatus() []string {
	return []string{
		string(CloudPCExternalPartnerSettingStatus_Available),
		string(CloudPCExternalPartnerSettingStatus_Healthy),
		string(CloudPCExternalPartnerSettingStatus_NotAvailable),
		string(CloudPCExternalPartnerSettingStatus_Unhealthy),
	}
}

func (s *CloudPCExternalPartnerSettingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCExternalPartnerSettingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCExternalPartnerSettingStatus(input string) (*CloudPCExternalPartnerSettingStatus, error) {
	vals := map[string]CloudPCExternalPartnerSettingStatus{
		"available":    CloudPCExternalPartnerSettingStatus_Available,
		"healthy":      CloudPCExternalPartnerSettingStatus_Healthy,
		"notavailable": CloudPCExternalPartnerSettingStatus_NotAvailable,
		"unhealthy":    CloudPCExternalPartnerSettingStatus_Unhealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCExternalPartnerSettingStatus(input)
	return &out, nil
}
