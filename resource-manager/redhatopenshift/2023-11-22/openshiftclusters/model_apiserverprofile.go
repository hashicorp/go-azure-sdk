package openshiftclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type APIServerProfile struct {
	IP         *string     `json:"ip,omitempty"`
	Url        *string     `json:"url,omitempty"`
	Visibility *Visibility `json:"visibility,omitempty"`
}
