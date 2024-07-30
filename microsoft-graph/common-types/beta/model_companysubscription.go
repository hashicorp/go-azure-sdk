package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanySubscription struct {
	CommerceSubscriptionId *string            `json:"commerceSubscriptionId,omitempty"`
	CreatedDateTime        *string            `json:"createdDateTime,omitempty"`
	Id                     *string            `json:"id,omitempty"`
	IsTrial                *bool              `json:"isTrial,omitempty"`
	NextLifecycleDateTime  *string            `json:"nextLifecycleDateTime,omitempty"`
	ODataType              *string            `json:"@odata.type,omitempty"`
	OcpSubscriptionId      *string            `json:"ocpSubscriptionId,omitempty"`
	OwnerId                *string            `json:"ownerId,omitempty"`
	OwnerTenantId          *string            `json:"ownerTenantId,omitempty"`
	OwnerType              *string            `json:"ownerType,omitempty"`
	ServiceStatus          *[]ServicePlanInfo `json:"serviceStatus,omitempty"`
	SkuId                  *string            `json:"skuId,omitempty"`
	SkuPartNumber          *string            `json:"skuPartNumber,omitempty"`
	Status                 *string            `json:"status,omitempty"`
	TotalLicenses          *int64             `json:"totalLicenses,omitempty"`
}
