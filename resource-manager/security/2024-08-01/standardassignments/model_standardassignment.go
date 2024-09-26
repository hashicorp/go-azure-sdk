package standardassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardAssignment struct {
	Id         *string                       `json:"id,omitempty"`
	Name       *string                       `json:"name,omitempty"`
	Properties *StandardAssignmentProperties `json:"properties,omitempty"`
	Type       *string                       `json:"type,omitempty"`
}
