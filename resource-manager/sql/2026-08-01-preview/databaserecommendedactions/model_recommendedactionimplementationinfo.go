package databaserecommendedactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionImplementationInfo struct {
	Method *ImplementationMethod `json:"method,omitempty"`
	Script *string               `json:"script,omitempty"`
}
