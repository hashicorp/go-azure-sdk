package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpoint struct {
	AuditEvents                             *[]CloudPCAuditEvent                            `json:"auditEvents,omitempty"`
	BulkActions                             *[]CloudPCBulkAction                            `json:"bulkActions,omitempty"`
	CloudPCs                                *[]CloudPC                                      `json:"cloudPCs,omitempty"`
	CrossCloudGovernmentOrganizationMapping *CloudPCCrossCloudGovernmentOrganizationMapping `json:"crossCloudGovernmentOrganizationMapping,omitempty"`
	DeviceImages                            *[]CloudPCDeviceImage                           `json:"deviceImages,omitempty"`
	ExternalPartnerSettings                 *[]CloudPCExternalPartnerSetting                `json:"externalPartnerSettings,omitempty"`
	FrontLineServicePlans                   *[]CloudPCFrontLineServicePlan                  `json:"frontLineServicePlans,omitempty"`
	GalleryImages                           *[]CloudPCGalleryImage                          `json:"galleryImages,omitempty"`
	Id                                      *string                                         `json:"id,omitempty"`
	ODataType                               *string                                         `json:"@odata.type,omitempty"`
	OnPremisesConnections                   *[]CloudPCOnPremisesConnection                  `json:"onPremisesConnections,omitempty"`
	OrganizationSettings                    *CloudPCOrganizationSettings                    `json:"organizationSettings,omitempty"`
	ProvisioningPolicies                    *[]CloudPCProvisioningPolicy                    `json:"provisioningPolicies,omitempty"`
	Reports                                 *CloudPCReports                                 `json:"reports,omitempty"`
	ServicePlans                            *[]CloudPCServicePlan                           `json:"servicePlans,omitempty"`
	SharedUseServicePlans                   *[]CloudPCSharedUseServicePlan                  `json:"sharedUseServicePlans,omitempty"`
	Snapshots                               *[]CloudPCSnapshot                              `json:"snapshots,omitempty"`
	SupportedRegions                        *[]CloudPCSupportedRegion                       `json:"supportedRegions,omitempty"`
	UserSettings                            *[]CloudPCUserSetting                           `json:"userSettings,omitempty"`
}
