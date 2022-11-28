package providers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Alias struct {
	DefaultMetadata *AliasPathMetadata `json:"defaultMetadata"`
	DefaultPath     *string            `json:"defaultPath,omitempty"`
	DefaultPattern  *AliasPattern      `json:"defaultPattern"`
	Name            *string            `json:"name,omitempty"`
	Paths           *[]AliasPath       `json:"paths,omitempty"`
	Type            *AliasType         `json:"type,omitempty"`
}
