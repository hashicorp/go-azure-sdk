package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppPolicyDeploymentSummary struct {
	ConfigurationDeployedUserCount       *int64                                     `json:"configurationDeployedUserCount,omitempty"`
	ConfigurationDeploymentSummaryPerApp *[]ManagedAppPolicyDeploymentSummaryPerApp `json:"configurationDeploymentSummaryPerApp,omitempty"`
	DisplayName                          *string                                    `json:"displayName,omitempty"`
	Id                                   *string                                    `json:"id,omitempty"`
	LastRefreshTime                      *string                                    `json:"lastRefreshTime,omitempty"`
	ODataType                            *string                                    `json:"@odata.type,omitempty"`
	Version                              *string                                    `json:"version,omitempty"`
}
