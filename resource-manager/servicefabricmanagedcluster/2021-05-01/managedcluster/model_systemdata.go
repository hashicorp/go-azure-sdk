package managedcluster

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SystemData struct {
	CreatedAt          *string `json:"createdAt,omitempty"`
	CreatedBy          *string `json:"createdBy,omitempty"`
	CreatedByType      *string `json:"createdByType,omitempty"`
	LastModifiedAt     *string `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string `json:"lastModifiedByType,omitempty"`
}
