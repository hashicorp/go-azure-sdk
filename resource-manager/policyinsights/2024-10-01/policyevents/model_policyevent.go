package policyevents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyEvent struct {
	ComplianceState               *string                  `json:"complianceState,omitempty"`
	Components                    *[]ComponentEventDetails `json:"components,omitempty"`
	EffectiveParameters           *string                  `json:"effectiveParameters,omitempty"`
	IsCompliant                   *bool                    `json:"isCompliant,omitempty"`
	ManagementGroupIds            *string                  `json:"managementGroupIds,omitempty"`
	OdataContext                  *string                  `json:"@odata.context,omitempty"`
	OdataId                       *string                  `json:"@odata.id,omitempty"`
	PolicyAssignmentId            *string                  `json:"policyAssignmentId,omitempty"`
	PolicyAssignmentName          *string                  `json:"policyAssignmentName,omitempty"`
	PolicyAssignmentOwner         *string                  `json:"policyAssignmentOwner,omitempty"`
	PolicyAssignmentParameters    *string                  `json:"policyAssignmentParameters,omitempty"`
	PolicyAssignmentScope         *string                  `json:"policyAssignmentScope,omitempty"`
	PolicyDefinitionAction        *string                  `json:"policyDefinitionAction,omitempty"`
	PolicyDefinitionCategory      *string                  `json:"policyDefinitionCategory,omitempty"`
	PolicyDefinitionId            *string                  `json:"policyDefinitionId,omitempty"`
	PolicyDefinitionName          *string                  `json:"policyDefinitionName,omitempty"`
	PolicyDefinitionReferenceId   *string                  `json:"policyDefinitionReferenceId,omitempty"`
	PolicySetDefinitionCategory   *string                  `json:"policySetDefinitionCategory,omitempty"`
	PolicySetDefinitionId         *string                  `json:"policySetDefinitionId,omitempty"`
	PolicySetDefinitionName       *string                  `json:"policySetDefinitionName,omitempty"`
	PolicySetDefinitionOwner      *string                  `json:"policySetDefinitionOwner,omitempty"`
	PolicySetDefinitionParameters *string                  `json:"policySetDefinitionParameters,omitempty"`
	PrincipalOid                  *string                  `json:"principalOid,omitempty"`
	ResourceGroup                 *string                  `json:"resourceGroup,omitempty"`
	ResourceId                    *string                  `json:"resourceId,omitempty"`
	ResourceLocation              *string                  `json:"resourceLocation,omitempty"`
	ResourceTags                  *string                  `json:"resourceTags,omitempty"`
	ResourceType                  *string                  `json:"resourceType,omitempty"`
	SubscriptionId                *string                  `json:"subscriptionId,omitempty"`
	TenantId                      *string                  `json:"tenantId,omitempty"`
	Timestamp                     *string                  `json:"timestamp,omitempty"`
}

func (o *PolicyEvent) GetTimestampAsTime() (*time.Time, error) {
	if o.Timestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Timestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *PolicyEvent) SetTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Timestamp = &formatted
}
