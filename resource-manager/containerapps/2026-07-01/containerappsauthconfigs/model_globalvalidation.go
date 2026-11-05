package containerappsauthconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalValidation struct {
	ExcludedPaths               *[]string                      `json:"excludedPaths,omitempty"`
	RedirectToProvider          *string                        `json:"redirectToProvider,omitempty"`
	UnauthenticatedClientAction *UnauthenticatedClientActionV2 `json:"unauthenticatedClientAction,omitempty"`
}
