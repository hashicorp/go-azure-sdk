package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAccidentalDeletionPrevention struct {
	AlertThreshold                *int64                                                               `json:"alertThreshold,omitempty"`
	ODataType                     *string                                                              `json:"@odata.type,omitempty"`
	SynchronizationPreventionType *OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType `json:"synchronizationPreventionType,omitempty"`
}
