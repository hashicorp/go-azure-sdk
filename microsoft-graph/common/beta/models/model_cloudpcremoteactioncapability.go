package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionCapability struct {
	ActionCapability *CloudPCRemoteActionCapabilityActionCapability `json:"actionCapability,omitempty"`
	ActionName       *CloudPCRemoteActionCapabilityActionName       `json:"actionName,omitempty"`
	ODataType        *string                                        `json:"@odata.type,omitempty"`
}
