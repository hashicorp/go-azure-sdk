package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommsNotification struct {
	ChangeType  *CommsNotificationChangeType `json:"changeType,omitempty"`
	ODataType   *string                      `json:"@odata.type,omitempty"`
	ResourceUrl *string                      `json:"resourceUrl,omitempty"`
}
