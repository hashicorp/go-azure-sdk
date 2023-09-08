package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationDeadPeerDetectionRate string

const (
	IosikEv2VpnConfigurationDeadPeerDetectionRatehigh   IosikEv2VpnConfigurationDeadPeerDetectionRate = "High"
	IosikEv2VpnConfigurationDeadPeerDetectionRatelow    IosikEv2VpnConfigurationDeadPeerDetectionRate = "Low"
	IosikEv2VpnConfigurationDeadPeerDetectionRatemedium IosikEv2VpnConfigurationDeadPeerDetectionRate = "Medium"
	IosikEv2VpnConfigurationDeadPeerDetectionRatenone   IosikEv2VpnConfigurationDeadPeerDetectionRate = "None"
)

func PossibleValuesForIosikEv2VpnConfigurationDeadPeerDetectionRate() []string {
	return []string{
		string(IosikEv2VpnConfigurationDeadPeerDetectionRatehigh),
		string(IosikEv2VpnConfigurationDeadPeerDetectionRatelow),
		string(IosikEv2VpnConfigurationDeadPeerDetectionRatemedium),
		string(IosikEv2VpnConfigurationDeadPeerDetectionRatenone),
	}
}

func parseIosikEv2VpnConfigurationDeadPeerDetectionRate(input string) (*IosikEv2VpnConfigurationDeadPeerDetectionRate, error) {
	vals := map[string]IosikEv2VpnConfigurationDeadPeerDetectionRate{
		"high":   IosikEv2VpnConfigurationDeadPeerDetectionRatehigh,
		"low":    IosikEv2VpnConfigurationDeadPeerDetectionRatelow,
		"medium": IosikEv2VpnConfigurationDeadPeerDetectionRatemedium,
		"none":   IosikEv2VpnConfigurationDeadPeerDetectionRatenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationDeadPeerDetectionRate(input)
	return &out, nil
}
