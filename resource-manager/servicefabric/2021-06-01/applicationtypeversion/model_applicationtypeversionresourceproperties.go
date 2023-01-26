package applicationtypeversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationTypeVersionResourceProperties struct {
	AppPackageUrl        string             `json:"appPackageUrl"`
	DefaultParameterList *map[string]string `json:"defaultParameterList,omitempty"`
	ProvisioningState    *string            `json:"provisioningState,omitempty"`
}
