package billingaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingAccountProperties struct {
	AccountStatus            *AccountStatus                  `json:"accountStatus,omitempty"`
	AccountStatusReasonCode  *BillingAccountStatusReasonCode `json:"accountStatusReasonCode,omitempty"`
	AccountSubType           *AccountSubType                 `json:"accountSubType,omitempty"`
	AccountType              *AccountType                    `json:"accountType,omitempty"`
	AgreementType            *AgreementType                  `json:"agreementType,omitempty"`
	BillingRelationshipTypes *[]BillingRelationshipType      `json:"billingRelationshipTypes,omitempty"`
	DisplayName              *string                         `json:"displayName,omitempty"`
	EnrollmentDetails        *EnrollmentDetails              `json:"enrollmentDetails,omitempty"`
	HasNoBillingProfiles     *bool                           `json:"hasNoBillingProfiles,omitempty"`
	HasReadAccess            *bool                           `json:"hasReadAccess,omitempty"`
	NotificationEmailAddress *string                         `json:"notificationEmailAddress,omitempty"`
	PrimaryBillingTenantId   *string                         `json:"primaryBillingTenantId,omitempty"`
	ProvisioningState        *ProvisioningState              `json:"provisioningState,omitempty"`
	Qualifications           *[]string                       `json:"qualifications,omitempty"`
	RegistrationNumber       *RegistrationNumber             `json:"registrationNumber,omitempty"`
	SoldTo                   *AddressDetails                 `json:"soldTo,omitempty"`
	TaxIds                   *[]TaxIdentifier                `json:"taxIds,omitempty"`
}
