package serverkeys

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerKeyProperties struct {
	CreationDate  *string       `json:"creationDate,omitempty"`
	ServerKeyType ServerKeyType `json:"serverKeyType"`
	Uri           *string       `json:"uri,omitempty"`
}
