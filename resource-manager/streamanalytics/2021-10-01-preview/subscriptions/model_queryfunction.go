package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryFunction struct {
	BindingType string          `json:"bindingType"`
	Inputs      []FunctionInput `json:"inputs"`
	Name        string          `json:"name"`
	Output      FunctionOutput  `json:"output"`
	Type        string          `json:"type"`
}
