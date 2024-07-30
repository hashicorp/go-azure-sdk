package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteUserFromSharedAppleDeviceActionResult struct {
	ActionName          *string                                                 `json:"actionName,omitempty"`
	ActionState         *DeleteUserFromSharedAppleDeviceActionResultActionState `json:"actionState,omitempty"`
	LastUpdatedDateTime *string                                                 `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string                                                 `json:"@odata.type,omitempty"`
	StartDateTime       *string                                                 `json:"startDateTime,omitempty"`
	UserPrincipalName   *string                                                 `json:"userPrincipalName,omitempty"`
}
