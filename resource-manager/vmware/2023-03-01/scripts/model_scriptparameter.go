package scripts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptParameter struct {
	Description *string                  `json:"description,omitempty"`
	Name        *string                  `json:"name,omitempty"`
	Optional    *OptionalParamEnum       `json:"optional,omitempty"`
	Type        *ScriptParameterTypes    `json:"type,omitempty"`
	Visibility  *VisibilityParameterEnum `json:"visibility,omitempty"`
}
