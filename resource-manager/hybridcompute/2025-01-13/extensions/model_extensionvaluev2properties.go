package extensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionValueV2Properties struct {
	Architecture          *[]string `json:"architecture,omitempty"`
	ExtensionSignatureUri *string   `json:"extensionSignatureUri,omitempty"`
	ExtensionType         *string   `json:"extensionType,omitempty"`
	ExtensionUris         *[]string `json:"extensionUris,omitempty"`
	OperatingSystem       *string   `json:"operatingSystem,omitempty"`
	Publisher             *string   `json:"publisher,omitempty"`
	Version               *string   `json:"version,omitempty"`
}
