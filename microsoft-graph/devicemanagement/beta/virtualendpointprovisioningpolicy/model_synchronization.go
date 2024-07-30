package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Synchronization struct {
	Id        *string                                    `json:"id,omitempty"`
	Jobs      *[]SynchronizationJob                      `json:"jobs,omitempty"`
	ODataType *string                                    `json:"@odata.type,omitempty"`
	Secrets   *[]SynchronizationSecretKeyStringValuePair `json:"secrets,omitempty"`
	Templates *[]SynchronizationTemplate                 `json:"templates,omitempty"`
}
