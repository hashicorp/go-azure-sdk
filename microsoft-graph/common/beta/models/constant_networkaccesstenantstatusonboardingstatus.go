package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTenantStatusOnboardingStatus string

const (
	NetworkaccessTenantStatusOnboardingStatusoffboarded               NetworkaccessTenantStatusOnboardingStatus = "Offboarded"
	NetworkaccessTenantStatusOnboardingStatusoffboardingErrorOccurred NetworkaccessTenantStatusOnboardingStatus = "OffboardingErrorOccurred"
	NetworkaccessTenantStatusOnboardingStatusoffboardingInProgress    NetworkaccessTenantStatusOnboardingStatus = "OffboardingInProgress"
	NetworkaccessTenantStatusOnboardingStatusonboarded                NetworkaccessTenantStatusOnboardingStatus = "Onboarded"
	NetworkaccessTenantStatusOnboardingStatusonboardingErrorOccurred  NetworkaccessTenantStatusOnboardingStatus = "OnboardingErrorOccurred"
	NetworkaccessTenantStatusOnboardingStatusonboardingInProgress     NetworkaccessTenantStatusOnboardingStatus = "OnboardingInProgress"
)

func PossibleValuesForNetworkaccessTenantStatusOnboardingStatus() []string {
	return []string{
		string(NetworkaccessTenantStatusOnboardingStatusoffboarded),
		string(NetworkaccessTenantStatusOnboardingStatusoffboardingErrorOccurred),
		string(NetworkaccessTenantStatusOnboardingStatusoffboardingInProgress),
		string(NetworkaccessTenantStatusOnboardingStatusonboarded),
		string(NetworkaccessTenantStatusOnboardingStatusonboardingErrorOccurred),
		string(NetworkaccessTenantStatusOnboardingStatusonboardingInProgress),
	}
}

func parseNetworkaccessTenantStatusOnboardingStatus(input string) (*NetworkaccessTenantStatusOnboardingStatus, error) {
	vals := map[string]NetworkaccessTenantStatusOnboardingStatus{
		"offboarded":               NetworkaccessTenantStatusOnboardingStatusoffboarded,
		"offboardingerroroccurred": NetworkaccessTenantStatusOnboardingStatusoffboardingErrorOccurred,
		"offboardinginprogress":    NetworkaccessTenantStatusOnboardingStatusoffboardingInProgress,
		"onboarded":                NetworkaccessTenantStatusOnboardingStatusonboarded,
		"onboardingerroroccurred":  NetworkaccessTenantStatusOnboardingStatusonboardingErrorOccurred,
		"onboardinginprogress":     NetworkaccessTenantStatusOnboardingStatusonboardingInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTenantStatusOnboardingStatus(input)
	return &out, nil
}
