package denyassignments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DenyAssignmentProperties struct {
	Condition               *string                     `json:"condition,omitempty"`
	ConditionVersion        *string                     `json:"conditionVersion,omitempty"`
	CreatedBy               *string                     `json:"createdBy,omitempty"`
	CreatedOn               *string                     `json:"createdOn,omitempty"`
	DenyAssignmentName      *string                     `json:"denyAssignmentName,omitempty"`
	Description             *string                     `json:"description,omitempty"`
	DoNotApplyToChildScopes *bool                       `json:"doNotApplyToChildScopes,omitempty"`
	ExcludePrincipals       *[]Principal                `json:"excludePrincipals,omitempty"`
	IsSystemProtected       *bool                       `json:"isSystemProtected,omitempty"`
	Permissions             *[]DenyAssignmentPermission `json:"permissions,omitempty"`
	Principals              *[]Principal                `json:"principals,omitempty"`
	Scope                   *string                     `json:"scope,omitempty"`
	UpdatedBy               *string                     `json:"updatedBy,omitempty"`
	UpdatedOn               *string                     `json:"updatedOn,omitempty"`
}

func (o *DenyAssignmentProperties) GetCreatedOnAsTime() (*time.Time, error) {
	if o.CreatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *DenyAssignmentProperties) GetUpdatedOnAsTime() (*time.Time, error) {
	if o.UpdatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.UpdatedOn, "2006-01-02T15:04:05Z07:00")
}
