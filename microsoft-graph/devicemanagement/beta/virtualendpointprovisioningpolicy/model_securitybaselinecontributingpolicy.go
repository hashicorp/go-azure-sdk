package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineContributingPolicy struct {
	DisplayName *string                                       `json:"displayName,omitempty"`
	ODataType   *string                                       `json:"@odata.type,omitempty"`
	SourceId    *string                                       `json:"sourceId,omitempty"`
	SourceType  *SecurityBaselineContributingPolicySourceType `json:"sourceType,omitempty"`
}
