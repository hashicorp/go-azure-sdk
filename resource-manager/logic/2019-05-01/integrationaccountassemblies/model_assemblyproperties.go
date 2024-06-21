package integrationaccountassemblies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssemblyProperties struct {
	AssemblyCulture        *string      `json:"assemblyCulture,omitempty"`
	AssemblyName           string       `json:"assemblyName"`
	AssemblyPublicKeyToken *string      `json:"assemblyPublicKeyToken,omitempty"`
	AssemblyVersion        *string      `json:"assemblyVersion,omitempty"`
	ChangedTime            *string      `json:"changedTime,omitempty"`
	Content                *interface{} `json:"content,omitempty"`
	ContentLink            *ContentLink `json:"contentLink,omitempty"`
	ContentType            *string      `json:"contentType,omitempty"`
	CreatedTime            *string      `json:"createdTime,omitempty"`
	Metadata               *interface{} `json:"metadata,omitempty"`
}
