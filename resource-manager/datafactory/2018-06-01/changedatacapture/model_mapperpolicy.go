package changedatacapture

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MapperPolicy struct {
	Mode       *string                 `json:"mode,omitempty"`
	Recurrence *MapperPolicyRecurrence `json:"recurrence,omitempty"`
}
