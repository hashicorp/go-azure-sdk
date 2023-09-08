package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionCapabilityActionName string

const (
	CloudPCRemoteActionCapabilityActionNamechangeUserAccountType CloudPCRemoteActionCapabilityActionName = "ChangeUserAccountType"
	CloudPCRemoteActionCapabilityActionNameplaceUnderReview      CloudPCRemoteActionCapabilityActionName = "PlaceUnderReview"
	CloudPCRemoteActionCapabilityActionNamerename                CloudPCRemoteActionCapabilityActionName = "Rename"
	CloudPCRemoteActionCapabilityActionNamereprovision           CloudPCRemoteActionCapabilityActionName = "Reprovision"
	CloudPCRemoteActionCapabilityActionNameresize                CloudPCRemoteActionCapabilityActionName = "Resize"
	CloudPCRemoteActionCapabilityActionNamerestart               CloudPCRemoteActionCapabilityActionName = "Restart"
	CloudPCRemoteActionCapabilityActionNamerestore               CloudPCRemoteActionCapabilityActionName = "Restore"
	CloudPCRemoteActionCapabilityActionNametroubleshoot          CloudPCRemoteActionCapabilityActionName = "Troubleshoot"
	CloudPCRemoteActionCapabilityActionNameunknown               CloudPCRemoteActionCapabilityActionName = "Unknown"
)

func PossibleValuesForCloudPCRemoteActionCapabilityActionName() []string {
	return []string{
		string(CloudPCRemoteActionCapabilityActionNamechangeUserAccountType),
		string(CloudPCRemoteActionCapabilityActionNameplaceUnderReview),
		string(CloudPCRemoteActionCapabilityActionNamerename),
		string(CloudPCRemoteActionCapabilityActionNamereprovision),
		string(CloudPCRemoteActionCapabilityActionNameresize),
		string(CloudPCRemoteActionCapabilityActionNamerestart),
		string(CloudPCRemoteActionCapabilityActionNamerestore),
		string(CloudPCRemoteActionCapabilityActionNametroubleshoot),
		string(CloudPCRemoteActionCapabilityActionNameunknown),
	}
}

func parseCloudPCRemoteActionCapabilityActionName(input string) (*CloudPCRemoteActionCapabilityActionName, error) {
	vals := map[string]CloudPCRemoteActionCapabilityActionName{
		"changeuseraccounttype": CloudPCRemoteActionCapabilityActionNamechangeUserAccountType,
		"placeunderreview":      CloudPCRemoteActionCapabilityActionNameplaceUnderReview,
		"rename":                CloudPCRemoteActionCapabilityActionNamerename,
		"reprovision":           CloudPCRemoteActionCapabilityActionNamereprovision,
		"resize":                CloudPCRemoteActionCapabilityActionNameresize,
		"restart":               CloudPCRemoteActionCapabilityActionNamerestart,
		"restore":               CloudPCRemoteActionCapabilityActionNamerestore,
		"troubleshoot":          CloudPCRemoteActionCapabilityActionNametroubleshoot,
		"unknown":               CloudPCRemoteActionCapabilityActionNameunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRemoteActionCapabilityActionName(input)
	return &out, nil
}
