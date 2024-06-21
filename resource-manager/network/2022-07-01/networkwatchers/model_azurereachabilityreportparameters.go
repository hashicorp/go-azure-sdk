package networkwatchers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureReachabilityReportParameters struct {
	AzureLocations   *[]string                       `json:"azureLocations,omitempty"`
	EndTime          string                          `json:"endTime"`
	ProviderLocation AzureReachabilityReportLocation `json:"providerLocation"`
	Providers        *[]string                       `json:"providers,omitempty"`
	StartTime        string                          `json:"startTime"`
}
