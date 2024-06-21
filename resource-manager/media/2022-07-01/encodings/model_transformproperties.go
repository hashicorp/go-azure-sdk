package encodings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransformProperties struct {
	Created      *string           `json:"created,omitempty"`
	Description  *string           `json:"description,omitempty"`
	LastModified *string           `json:"lastModified,omitempty"`
	Outputs      []TransformOutput `json:"outputs"`
}
