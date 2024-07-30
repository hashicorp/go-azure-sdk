package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppPolicyDeploymentSummaryPerApp struct {
	ConfigurationAppliedUserCount *int64               `json:"configurationAppliedUserCount,omitempty"`
	MobileAppIdentifier           *MobileAppIdentifier `json:"mobileAppIdentifier,omitempty"`
	ODataType                     *string              `json:"@odata.type,omitempty"`
}
