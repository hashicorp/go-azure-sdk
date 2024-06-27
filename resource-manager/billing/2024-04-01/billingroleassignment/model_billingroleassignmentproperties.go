package billingroleassignment

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRoleAssignmentProperties struct {
	BillingAccountDisplayName   *string            `json:"billingAccountDisplayName,omitempty"`
	BillingAccountId            *string            `json:"billingAccountId,omitempty"`
	BillingProfileDisplayName   *string            `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId            *string            `json:"billingProfileId,omitempty"`
	BillingRequestId            *string            `json:"billingRequestId,omitempty"`
	CreatedByPrincipalId        *string            `json:"createdByPrincipalId,omitempty"`
	CreatedByPrincipalPuid      *string            `json:"createdByPrincipalPuid,omitempty"`
	CreatedByPrincipalTenantId  *string            `json:"createdByPrincipalTenantId,omitempty"`
	CreatedByUserEmailAddress   *string            `json:"createdByUserEmailAddress,omitempty"`
	CreatedOn                   *string            `json:"createdOn,omitempty"`
	CustomerDisplayName         *string            `json:"customerDisplayName,omitempty"`
	CustomerId                  *string            `json:"customerId,omitempty"`
	InvoiceSectionDisplayName   *string            `json:"invoiceSectionDisplayName,omitempty"`
	InvoiceSectionId            *string            `json:"invoiceSectionId,omitempty"`
	ModifiedByPrincipalId       *string            `json:"modifiedByPrincipalId,omitempty"`
	ModifiedByPrincipalPuid     *string            `json:"modifiedByPrincipalPuid,omitempty"`
	ModifiedByPrincipalTenantId *string            `json:"modifiedByPrincipalTenantId,omitempty"`
	ModifiedByUserEmailAddress  *string            `json:"modifiedByUserEmailAddress,omitempty"`
	ModifiedOn                  *string            `json:"modifiedOn,omitempty"`
	PrincipalDisplayName        *string            `json:"principalDisplayName,omitempty"`
	PrincipalId                 *string            `json:"principalId,omitempty"`
	PrincipalPuid               *string            `json:"principalPuid,omitempty"`
	PrincipalTenantId           *string            `json:"principalTenantId,omitempty"`
	PrincipalTenantName         *string            `json:"principalTenantName,omitempty"`
	PrincipalType               *PrincipalType     `json:"principalType,omitempty"`
	ProvisioningState           *ProvisioningState `json:"provisioningState,omitempty"`
	RoleDefinitionId            string             `json:"roleDefinitionId"`
	Scope                       *string            `json:"scope,omitempty"`
	UserAuthenticationType      *string            `json:"userAuthenticationType,omitempty"`
	UserEmailAddress            *string            `json:"userEmailAddress,omitempty"`
}

func (o *BillingRoleAssignmentProperties) GetCreatedOnAsTime() (*time.Time, error) {
	if o.CreatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *BillingRoleAssignmentProperties) SetCreatedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedOn = &formatted
}

func (o *BillingRoleAssignmentProperties) GetModifiedOnAsTime() (*time.Time, error) {
	if o.ModifiedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ModifiedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *BillingRoleAssignmentProperties) SetModifiedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ModifiedOn = &formatted
}
