package replicationalertsettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertProperties struct {
	CustomEmailAddresses *[]string `json:"customEmailAddresses,omitempty"`
	Locale               *string   `json:"locale,omitempty"`
	SendToOwners         *string   `json:"sendToOwners,omitempty"`
}
