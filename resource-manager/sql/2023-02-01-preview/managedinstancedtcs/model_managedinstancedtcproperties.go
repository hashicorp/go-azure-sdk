package managedinstancedtcs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceDtcProperties struct {
	DtcEnabled                  *bool                               `json:"dtcEnabled,omitempty"`
	DtcHostNameDnsSuffix        *string                             `json:"dtcHostNameDnsSuffix,omitempty"`
	ExternalDnsSuffixSearchList *[]string                           `json:"externalDnsSuffixSearchList,omitempty"`
	ProvisioningState           *ProvisioningState                  `json:"provisioningState,omitempty"`
	SecuritySettings            *ManagedInstanceDtcSecuritySettings `json:"securitySettings,omitempty"`
}
