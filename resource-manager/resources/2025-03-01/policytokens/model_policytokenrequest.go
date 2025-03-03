package policytokens

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyTokenRequest struct {
	ChangeReference *string               `json:"changeReference,omitempty"`
	Operation       *PolicyTokenOperation `json:"operation,omitempty"`
}
