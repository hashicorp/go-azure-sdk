package domains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainPatchResource struct {
	Id         *string                        `json:"id,omitempty"`
	Kind       *string                        `json:"kind,omitempty"`
	Name       *string                        `json:"name,omitempty"`
	Properties *DomainPatchResourceProperties `json:"properties,omitempty"`
	Type       *string                        `json:"type,omitempty"`
}
