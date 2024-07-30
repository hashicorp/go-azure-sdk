package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionCapabilityActionName string

const (
	CloudPCRemoteActionCapabilityActionName_ChangeUserAccountType CloudPCRemoteActionCapabilityActionName = "changeUserAccountType"
	CloudPCRemoteActionCapabilityActionName_CreateSnapshot        CloudPCRemoteActionCapabilityActionName = "createSnapshot"
	CloudPCRemoteActionCapabilityActionName_MoveRegion            CloudPCRemoteActionCapabilityActionName = "moveRegion"
	CloudPCRemoteActionCapabilityActionName_PlaceUnderReview      CloudPCRemoteActionCapabilityActionName = "placeUnderReview"
	CloudPCRemoteActionCapabilityActionName_PowerOff              CloudPCRemoteActionCapabilityActionName = "powerOff"
	CloudPCRemoteActionCapabilityActionName_PowerOn               CloudPCRemoteActionCapabilityActionName = "powerOn"
	CloudPCRemoteActionCapabilityActionName_Rename                CloudPCRemoteActionCapabilityActionName = "rename"
	CloudPCRemoteActionCapabilityActionName_Reprovision           CloudPCRemoteActionCapabilityActionName = "reprovision"
	CloudPCRemoteActionCapabilityActionName_Resize                CloudPCRemoteActionCapabilityActionName = "resize"
	CloudPCRemoteActionCapabilityActionName_Restart               CloudPCRemoteActionCapabilityActionName = "restart"
	CloudPCRemoteActionCapabilityActionName_Restore               CloudPCRemoteActionCapabilityActionName = "restore"
	CloudPCRemoteActionCapabilityActionName_Troubleshoot          CloudPCRemoteActionCapabilityActionName = "troubleshoot"
	CloudPCRemoteActionCapabilityActionName_Unknown               CloudPCRemoteActionCapabilityActionName = "unknown"
)

func PossibleValuesForCloudPCRemoteActionCapabilityActionName() []string {
	return []string{
		string(CloudPCRemoteActionCapabilityActionName_ChangeUserAccountType),
		string(CloudPCRemoteActionCapabilityActionName_CreateSnapshot),
		string(CloudPCRemoteActionCapabilityActionName_MoveRegion),
		string(CloudPCRemoteActionCapabilityActionName_PlaceUnderReview),
		string(CloudPCRemoteActionCapabilityActionName_PowerOff),
		string(CloudPCRemoteActionCapabilityActionName_PowerOn),
		string(CloudPCRemoteActionCapabilityActionName_Rename),
		string(CloudPCRemoteActionCapabilityActionName_Reprovision),
		string(CloudPCRemoteActionCapabilityActionName_Resize),
		string(CloudPCRemoteActionCapabilityActionName_Restart),
		string(CloudPCRemoteActionCapabilityActionName_Restore),
		string(CloudPCRemoteActionCapabilityActionName_Troubleshoot),
		string(CloudPCRemoteActionCapabilityActionName_Unknown),
	}
}

func (s *CloudPCRemoteActionCapabilityActionName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCRemoteActionCapabilityActionName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCRemoteActionCapabilityActionName(input string) (*CloudPCRemoteActionCapabilityActionName, error) {
	vals := map[string]CloudPCRemoteActionCapabilityActionName{
		"changeuseraccounttype": CloudPCRemoteActionCapabilityActionName_ChangeUserAccountType,
		"createsnapshot":        CloudPCRemoteActionCapabilityActionName_CreateSnapshot,
		"moveregion":            CloudPCRemoteActionCapabilityActionName_MoveRegion,
		"placeunderreview":      CloudPCRemoteActionCapabilityActionName_PlaceUnderReview,
		"poweroff":              CloudPCRemoteActionCapabilityActionName_PowerOff,
		"poweron":               CloudPCRemoteActionCapabilityActionName_PowerOn,
		"rename":                CloudPCRemoteActionCapabilityActionName_Rename,
		"reprovision":           CloudPCRemoteActionCapabilityActionName_Reprovision,
		"resize":                CloudPCRemoteActionCapabilityActionName_Resize,
		"restart":               CloudPCRemoteActionCapabilityActionName_Restart,
		"restore":               CloudPCRemoteActionCapabilityActionName_Restore,
		"troubleshoot":          CloudPCRemoteActionCapabilityActionName_Troubleshoot,
		"unknown":               CloudPCRemoteActionCapabilityActionName_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRemoteActionCapabilityActionName(input)
	return &out, nil
}
