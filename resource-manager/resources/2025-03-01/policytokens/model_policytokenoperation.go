package policytokens

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyTokenOperation struct {
	Content    *interface{} `json:"content,omitempty"`
	HTTPMethod *string      `json:"httpMethod,omitempty"`
	Uri        *string      `json:"uri,omitempty"`
}
