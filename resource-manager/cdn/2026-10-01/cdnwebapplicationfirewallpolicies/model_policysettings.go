package cdnwebapplicationfirewallpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySettings struct {
	DefaultCustomBlockResponseBody       *string                                             `json:"defaultCustomBlockResponseBody,omitempty"`
	DefaultCustomBlockResponseStatusCode *PolicySettingsDefaultCustomBlockResponseStatusCode `json:"defaultCustomBlockResponseStatusCode,omitempty"`
	DefaultRedirectURL                   *string                                             `json:"defaultRedirectUrl,omitempty"`
	EnabledState                         *PolicyEnabledState                                 `json:"enabledState,omitempty"`
	Mode                                 *PolicyMode                                         `json:"mode,omitempty"`
}
