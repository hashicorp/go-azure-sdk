package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RejectJoinResponse struct {
	ODataType *string                   `json:"@odata.type,omitempty"`
	Reason    *RejectJoinResponseReason `json:"reason,omitempty"`
}
