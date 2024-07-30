package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDomainJoinConfiguration struct {
	ODataType              *string                                    `json:"@odata.type,omitempty"`
	OnPremisesConnectionId *string                                    `json:"onPremisesConnectionId,omitempty"`
	RegionGroup            *CloudPCDomainJoinConfigurationRegionGroup `json:"regionGroup,omitempty"`
	RegionName             *string                                    `json:"regionName,omitempty"`
	Type                   *CloudPCDomainJoinConfigurationType        `json:"type,omitempty"`
}
