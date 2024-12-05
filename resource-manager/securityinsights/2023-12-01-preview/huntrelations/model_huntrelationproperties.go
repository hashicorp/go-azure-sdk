package huntrelations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HuntRelationProperties struct {
	Labels              *[]string `json:"labels,omitempty"`
	RelatedResourceId   string    `json:"relatedResourceId"`
	RelatedResourceKind *string   `json:"relatedResourceKind,omitempty"`
	RelatedResourceName *string   `json:"relatedResourceName,omitempty"`
	RelationType        *string   `json:"relationType,omitempty"`
}
