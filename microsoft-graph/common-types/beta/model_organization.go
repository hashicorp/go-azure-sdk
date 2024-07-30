package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Organization struct {
	AssignedPlans                             *[]AssignedPlan                              `json:"assignedPlans,omitempty"`
	Branding                                  *OrganizationalBranding                      `json:"branding,omitempty"`
	BusinessPhones                            *[]string                                    `json:"businessPhones,omitempty"`
	CertificateBasedAuthConfiguration         *[]CertificateBasedAuthConfiguration         `json:"certificateBasedAuthConfiguration,omitempty"`
	CertificateConnectorSetting               *CertificateConnectorSetting                 `json:"certificateConnectorSetting,omitempty"`
	City                                      *string                                      `json:"city,omitempty"`
	Country                                   *string                                      `json:"country,omitempty"`
	CountryLetterCode                         *string                                      `json:"countryLetterCode,omitempty"`
	CreatedDateTime                           *string                                      `json:"createdDateTime,omitempty"`
	DefaultUsageLocation                      *string                                      `json:"defaultUsageLocation,omitempty"`
	DeletedDateTime                           *string                                      `json:"deletedDateTime,omitempty"`
	DirectorySizeQuota                        *DirectorySizeQuota                          `json:"directorySizeQuota,omitempty"`
	DisplayName                               *string                                      `json:"displayName,omitempty"`
	Extensions                                *[]Extension                                 `json:"extensions,omitempty"`
	Id                                        *string                                      `json:"id,omitempty"`
	IsMultipleDataLocationsForServicesEnabled *bool                                        `json:"isMultipleDataLocationsForServicesEnabled,omitempty"`
	MarketingNotificationEmails               *[]string                                    `json:"marketingNotificationEmails,omitempty"`
	MobileDeviceManagementAuthority           *OrganizationMobileDeviceManagementAuthority `json:"mobileDeviceManagementAuthority,omitempty"`
	ODataType                                 *string                                      `json:"@odata.type,omitempty"`
	OnPremisesLastPasswordSyncDateTime        *string                                      `json:"onPremisesLastPasswordSyncDateTime,omitempty"`
	OnPremisesLastSyncDateTime                *string                                      `json:"onPremisesLastSyncDateTime,omitempty"`
	OnPremisesSyncEnabled                     *bool                                        `json:"onPremisesSyncEnabled,omitempty"`
	PartnerInformation                        *PartnerInformation                          `json:"partnerInformation,omitempty"`
	PartnerTenantType                         *OrganizationPartnerTenantType               `json:"partnerTenantType,omitempty"`
	PostalCode                                *string                                      `json:"postalCode,omitempty"`
	PreferredLanguage                         *string                                      `json:"preferredLanguage,omitempty"`
	PrivacyProfile                            *PrivacyProfile                              `json:"privacyProfile,omitempty"`
	ProvisionedPlans                          *[]ProvisionedPlan                           `json:"provisionedPlans,omitempty"`
	SecurityComplianceNotificationMails       *[]string                                    `json:"securityComplianceNotificationMails,omitempty"`
	SecurityComplianceNotificationPhones      *[]string                                    `json:"securityComplianceNotificationPhones,omitempty"`
	Settings                                  *OrganizationSettings                        `json:"settings,omitempty"`
	State                                     *string                                      `json:"state,omitempty"`
	Street                                    *string                                      `json:"street,omitempty"`
	TechnicalNotificationMails                *[]string                                    `json:"technicalNotificationMails,omitempty"`
	VerifiedDomains                           *[]VerifiedDomain                            `json:"verifiedDomains,omitempty"`
}
