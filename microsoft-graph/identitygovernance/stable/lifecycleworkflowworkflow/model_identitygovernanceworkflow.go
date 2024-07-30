package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflow struct {
	Category                *IdentityGovernanceWorkflowCategory            `json:"category,omitempty"`
	CreatedBy               *User                                          `json:"createdBy,omitempty"`
	CreatedDateTime         *string                                        `json:"createdDateTime,omitempty"`
	DeletedDateTime         *string                                        `json:"deletedDateTime,omitempty"`
	Description             *string                                        `json:"description,omitempty"`
	DisplayName             *string                                        `json:"displayName,omitempty"`
	ExecutionConditions     *IdentityGovernanceWorkflowExecutionConditions `json:"executionConditions,omitempty"`
	ExecutionScope          *[]IdentityGovernanceUserProcessingResult      `json:"executionScope,omitempty"`
	Id                      *string                                        `json:"id,omitempty"`
	IsEnabled               *bool                                          `json:"isEnabled,omitempty"`
	IsSchedulingEnabled     *bool                                          `json:"isSchedulingEnabled,omitempty"`
	LastModifiedBy          *User                                          `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime    *string                                        `json:"lastModifiedDateTime,omitempty"`
	NextScheduleRunDateTime *string                                        `json:"nextScheduleRunDateTime,omitempty"`
	ODataType               *string                                        `json:"@odata.type,omitempty"`
	Runs                    *[]IdentityGovernanceRun                       `json:"runs,omitempty"`
	TaskReports             *[]IdentityGovernanceTaskReport                `json:"taskReports,omitempty"`
	Tasks                   *[]IdentityGovernanceTask                      `json:"tasks,omitempty"`
	UserProcessingResults   *[]IdentityGovernanceUserProcessingResult      `json:"userProcessingResults,omitempty"`
	Version                 *int64                                         `json:"version,omitempty"`
	Versions                *[]IdentityGovernanceWorkflowVersion           `json:"versions,omitempty"`
}
