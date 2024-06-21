package resourceproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlProperties struct {
	ExpirationTime *string `json:"expirationTime,omitempty"`
	RefreshToken   *string `json:"refreshToken,omitempty"`
	Token          *string `json:"token,omitempty"`
	TokenSecret    *string `json:"tokenSecret,omitempty"`
}
