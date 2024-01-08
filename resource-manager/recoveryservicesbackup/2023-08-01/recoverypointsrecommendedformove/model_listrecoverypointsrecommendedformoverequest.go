package recoverypointsrecommendedformove

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRecoveryPointsRecommendedForMoveRequest struct {
	ExcludedRPList *[]string `json:"excludedRPList,omitempty"`
	ObjectType     *string   `json:"objectType,omitempty"`
}
