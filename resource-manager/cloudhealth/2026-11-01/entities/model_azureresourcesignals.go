package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureResourceSignals struct {
	AuthenticationSetting string                     `json:"authenticationSetting"`
	AzureResourceId       string                     `json:"azureResourceId"`
	AzureResourceKind     *string                    `json:"azureResourceKind,omitempty"`
	ResourceHealth        *AzureResourceHealthSignal `json:"resourceHealth,omitempty"`
	Signals               *[]AzureResourceSignal     `json:"signals,omitempty"`
}
