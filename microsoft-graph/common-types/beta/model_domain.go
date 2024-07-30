package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Domain struct {
	AuthenticationType               *string                        `json:"authenticationType,omitempty"`
	AvailabilityStatus               *string                        `json:"availabilityStatus,omitempty"`
	DomainNameReferences             *[]DirectoryObject             `json:"domainNameReferences,omitempty"`
	FederationConfiguration          *[]InternalDomainFederation    `json:"federationConfiguration,omitempty"`
	Id                               *string                        `json:"id,omitempty"`
	IsAdminManaged                   *bool                          `json:"isAdminManaged,omitempty"`
	IsDefault                        *bool                          `json:"isDefault,omitempty"`
	IsInitial                        *bool                          `json:"isInitial,omitempty"`
	IsRoot                           *bool                          `json:"isRoot,omitempty"`
	IsVerified                       *bool                          `json:"isVerified,omitempty"`
	ODataType                        *string                        `json:"@odata.type,omitempty"`
	PasswordNotificationWindowInDays *int64                         `json:"passwordNotificationWindowInDays,omitempty"`
	PasswordValidityPeriodInDays     *int64                         `json:"passwordValidityPeriodInDays,omitempty"`
	ServiceConfigurationRecords      *[]DomainDnsRecord             `json:"serviceConfigurationRecords,omitempty"`
	SharedEmailDomainInvitations     *[]SharedEmailDomainInvitation `json:"sharedEmailDomainInvitations,omitempty"`
	State                            *DomainState                   `json:"state,omitempty"`
	SupportedServices                *[]string                      `json:"supportedServices,omitempty"`
	VerificationDnsRecords           *[]DomainDnsRecord             `json:"verificationDnsRecords,omitempty"`
}
