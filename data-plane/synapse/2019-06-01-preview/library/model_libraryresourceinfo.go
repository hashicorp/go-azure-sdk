package library

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LibraryResourceInfo struct {
	ArtifactId  *string `json:"artifactId,omitempty"`
	Changed     *string `json:"changed,omitempty"`
	Created     *string `json:"created,omitempty"`
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	OperationId *string `json:"operationId,omitempty"`
	RecordId    *int64  `json:"recordId,omitempty"`
	State       *string `json:"state,omitempty"`
	Type        *string `json:"type,omitempty"`
}
