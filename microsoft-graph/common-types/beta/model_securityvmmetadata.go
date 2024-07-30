package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityVmMetadata struct {
	CloudProvider  *SecurityVmMetadataCloudProvider `json:"cloudProvider,omitempty"`
	ODataType      *string                          `json:"@odata.type,omitempty"`
	ResourceId     *string                          `json:"resourceId,omitempty"`
	SubscriptionId *string                          `json:"subscriptionId,omitempty"`
	VmId           *string                          `json:"vmId,omitempty"`
}
