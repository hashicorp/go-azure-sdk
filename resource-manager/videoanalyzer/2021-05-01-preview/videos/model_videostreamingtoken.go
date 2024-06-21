package videos

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VideoStreamingToken struct {
	ExpirationDate *string `json:"expirationDate,omitempty"`
	Token          *string `json:"token,omitempty"`
}
