package rolemanagementpolicyassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAssignmentPropertiesPolicy struct {
	Id                   *string    `json:"id,omitempty"`
	LastModifiedBy       *Principal `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string    `json:"lastModifiedDateTime,omitempty"`
}
