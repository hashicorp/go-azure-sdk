package policydefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyDefinition struct {
	Name       *string                     `json:"name,omitempty"`
	Properties *PolicyDefinitionProperties `json:"properties,omitempty"`
}
