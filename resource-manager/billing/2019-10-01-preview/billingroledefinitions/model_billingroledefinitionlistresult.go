package billingroledefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRoleDefinitionListResult struct {
	NextLink *string                  `json:"nextLink,omitempty"`
	Value    *[]BillingRoleDefinition `json:"value,omitempty"`
}
