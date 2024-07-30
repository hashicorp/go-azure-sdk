package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataFileUploadSession struct {
	ContainerExpirationDateTime *string `json:"containerExpirationDateTime,omitempty"`
	ContainerId                 *string `json:"containerId,omitempty"`
	ODataType                   *string `json:"@odata.type,omitempty"`
	SessionExpirationDateTime   *string `json:"sessionExpirationDateTime,omitempty"`
	SessionUrl                  *string `json:"sessionUrl,omitempty"`
}
