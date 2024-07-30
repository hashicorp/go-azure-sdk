package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCLaunchInfo struct {
	CloudPCId                           *string `json:"cloudPcId,omitempty"`
	CloudPCLaunchUrl                    *string `json:"cloudPcLaunchUrl,omitempty"`
	ODataType                           *string `json:"@odata.type,omitempty"`
	Windows365SwitchCompatible          *bool   `json:"windows365SwitchCompatible,omitempty"`
	Windows365SwitchNotCompatibleReason *string `json:"windows365SwitchNotCompatibleReason,omitempty"`
}
