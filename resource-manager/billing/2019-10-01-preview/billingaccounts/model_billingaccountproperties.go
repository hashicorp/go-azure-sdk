package billingaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingAccountProperties struct {
	AccountType        *AccountType         `json:"accountType,omitempty"`
	Address            *AddressDetails      `json:"address,omitempty"`
	AgreementType      *AgreementType       `json:"agreementType,omitempty"`
	BillingProfiles    *[]BillingProfile    `json:"billingProfiles,omitempty"`
	CustomerType       *CustomerType        `json:"customerType,omitempty"`
	Departments        *[]Department        `json:"departments,omitempty"`
	DisplayName        *string              `json:"displayName,omitempty"`
	EnrollmentAccounts *[]EnrollmentAccount `json:"enrollmentAccounts,omitempty"`
	EnrollmentDetails  *Enrollment          `json:"enrollmentDetails,omitempty"`
	OrganizationId     *string              `json:"organizationId,omitempty"`
}
