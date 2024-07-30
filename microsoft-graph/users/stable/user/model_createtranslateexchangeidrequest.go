package user

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTranslateExchangeIdRequest struct {
	InputIds     *[]string                                     `json:"InputIds,omitempty"`
	SourceIdType *CreateTranslateExchangeIdRequestSourceIdType `json:"SourceIdType,omitempty"`
	TargetIdType *CreateTranslateExchangeIdRequestTargetIdType `json:"TargetIdType,omitempty"`
}
