package provider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FunctionAppMinorVersion struct {
	DisplayText   *string              `json:"displayText,omitempty"`
	StackSettings *FunctionAppRuntimes `json:"stackSettings,omitempty"`
	Value         *string              `json:"value,omitempty"`
}
