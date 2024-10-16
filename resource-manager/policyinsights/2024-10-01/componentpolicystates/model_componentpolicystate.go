package componentpolicystates

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComponentPolicyState struct {
	ComplianceState               *string                           `json:"complianceState,omitempty"`
	ComponentId                   *string                           `json:"componentId,omitempty"`
	ComponentName                 *string                           `json:"componentName,omitempty"`
	ComponentType                 *string                           `json:"componentType,omitempty"`
	OdataContext                  *string                           `json:"@odata.context,omitempty"`
	OdataId                       *string                           `json:"@odata.id,omitempty"`
	PolicyAssignmentId            *string                           `json:"policyAssignmentId,omitempty"`
	PolicyAssignmentName          *string                           `json:"policyAssignmentName,omitempty"`
	PolicyAssignmentOwner         *string                           `json:"policyAssignmentOwner,omitempty"`
	PolicyAssignmentParameters    *string                           `json:"policyAssignmentParameters,omitempty"`
	PolicyAssignmentScope         *string                           `json:"policyAssignmentScope,omitempty"`
	PolicyAssignmentVersion       *string                           `json:"policyAssignmentVersion,omitempty"`
	PolicyDefinitionAction        *string                           `json:"policyDefinitionAction,omitempty"`
	PolicyDefinitionCategory      *string                           `json:"policyDefinitionCategory,omitempty"`
	PolicyDefinitionGroupNames    *[]string                         `json:"policyDefinitionGroupNames,omitempty"`
	PolicyDefinitionId            *string                           `json:"policyDefinitionId,omitempty"`
	PolicyDefinitionName          *string                           `json:"policyDefinitionName,omitempty"`
	PolicyDefinitionReferenceId   *string                           `json:"policyDefinitionReferenceId,omitempty"`
	PolicyDefinitionVersion       *string                           `json:"policyDefinitionVersion,omitempty"`
	PolicyEvaluationDetails       *ComponentPolicyEvaluationDetails `json:"policyEvaluationDetails,omitempty"`
	PolicySetDefinitionCategory   *string                           `json:"policySetDefinitionCategory,omitempty"`
	PolicySetDefinitionId         *string                           `json:"policySetDefinitionId,omitempty"`
	PolicySetDefinitionName       *string                           `json:"policySetDefinitionName,omitempty"`
	PolicySetDefinitionOwner      *string                           `json:"policySetDefinitionOwner,omitempty"`
	PolicySetDefinitionParameters *string                           `json:"policySetDefinitionParameters,omitempty"`
	PolicySetDefinitionVersion    *string                           `json:"policySetDefinitionVersion,omitempty"`
	ResourceGroup                 *string                           `json:"resourceGroup,omitempty"`
	ResourceId                    *string                           `json:"resourceId,omitempty"`
	ResourceLocation              *string                           `json:"resourceLocation,omitempty"`
	ResourceType                  *string                           `json:"resourceType,omitempty"`
	SubscriptionId                *string                           `json:"subscriptionId,omitempty"`
	Timestamp                     *string                           `json:"timestamp,omitempty"`
}

func (o *ComponentPolicyState) GetTimestampAsTime() (*time.Time, error) {
	if o.Timestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Timestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *ComponentPolicyState) SetTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Timestamp = &formatted
}
