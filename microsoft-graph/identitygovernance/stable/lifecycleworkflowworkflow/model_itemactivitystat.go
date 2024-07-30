package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemActivityStat struct {
	Access         *ItemActionStat `json:"access,omitempty"`
	Activities     *[]ItemActivity `json:"activities,omitempty"`
	Create         *ItemActionStat `json:"create,omitempty"`
	Delete         *ItemActionStat `json:"delete,omitempty"`
	Edit           *ItemActionStat `json:"edit,omitempty"`
	EndDateTime    *string         `json:"endDateTime,omitempty"`
	Id             *string         `json:"id,omitempty"`
	IncompleteData *IncompleteData `json:"incompleteData,omitempty"`
	IsTrending     *bool           `json:"isTrending,omitempty"`
	Move           *ItemActionStat `json:"move,omitempty"`
	ODataType      *string         `json:"@odata.type,omitempty"`
	StartDateTime  *string         `json:"startDateTime,omitempty"`
}
