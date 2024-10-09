package applicationtypeversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationTypeVersionResourceProperties struct {
	AppPackageURL     string  `json:"appPackageUrl"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
}
