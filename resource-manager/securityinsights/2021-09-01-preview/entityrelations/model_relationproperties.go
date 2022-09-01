package entityrelations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelationProperties struct {
	RelatedResourceId   string  `json:"relatedResourceId"`
	RelatedResourceKind *string `json:"relatedResourceKind,omitempty"`
	RelatedResourceName *string `json:"relatedResourceName,omitempty"`
	RelatedResourceType *string `json:"relatedResourceType,omitempty"`
}
