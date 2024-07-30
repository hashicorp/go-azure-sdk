package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageAction struct {
	Action     *CalendarSharingMessageActionAction     `json:"action,omitempty"`
	ActionType *CalendarSharingMessageActionActionType `json:"actionType,omitempty"`
	Importance *CalendarSharingMessageActionImportance `json:"importance,omitempty"`
	ODataType  *string                                 `json:"@odata.type,omitempty"`
}
