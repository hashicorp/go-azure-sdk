package spacecraft

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizedGroundstation struct {
	ExpirationDate *string `json:"expirationDate,omitempty"`
	GroundStation  *string `json:"groundStation,omitempty"`
}
