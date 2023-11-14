package provider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StackMinorVersion struct {
	DisplayVersion           *string `json:"displayVersion,omitempty"`
	IsDefault                *bool   `json:"isDefault,omitempty"`
	IsRemoteDebuggingEnabled *bool   `json:"isRemoteDebuggingEnabled,omitempty"`
	RuntimeVersion           *string `json:"runtimeVersion,omitempty"`
}
