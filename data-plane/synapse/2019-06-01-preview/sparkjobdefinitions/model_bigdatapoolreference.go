package sparkjobdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BigDataPoolReference struct {
	ReferenceName string                   `json:"referenceName"`
	Type          BigDataPoolReferenceType `json:"type"`
}
