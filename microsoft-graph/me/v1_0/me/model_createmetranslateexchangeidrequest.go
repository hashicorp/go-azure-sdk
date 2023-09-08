package me

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeTranslateExchangeIdRequest struct {
	InputIds     *[]string                                       `json:"InputIds,omitempty"`
	SourceIdType *CreateMeTranslateExchangeIdRequestSourceIdType `json:"SourceIdType,omitempty"`
	TargetIdType *CreateMeTranslateExchangeIdRequestTargetIdType `json:"TargetIdType,omitempty"`
}
