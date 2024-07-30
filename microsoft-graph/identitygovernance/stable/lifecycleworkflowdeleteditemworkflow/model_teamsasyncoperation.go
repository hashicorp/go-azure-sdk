package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAsyncOperation struct {
	AttemptsCount          *int64                            `json:"attemptsCount,omitempty"`
	CreatedDateTime        *string                           `json:"createdDateTime,omitempty"`
	Error                  *OperationError                   `json:"error,omitempty"`
	Id                     *string                           `json:"id,omitempty"`
	LastActionDateTime     *string                           `json:"lastActionDateTime,omitempty"`
	ODataType              *string                           `json:"@odata.type,omitempty"`
	OperationType          *TeamsAsyncOperationOperationType `json:"operationType,omitempty"`
	Status                 *TeamsAsyncOperationStatus        `json:"status,omitempty"`
	TargetResourceId       *string                           `json:"targetResourceId,omitempty"`
	TargetResourceLocation *string                           `json:"targetResourceLocation,omitempty"`
}
