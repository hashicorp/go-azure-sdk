package user

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdTranslateExchangeIdRequest struct {
	InputIds     *[]string                                             `json:"InputIds,omitempty"`
	SourceIdType *CreateUserByIdTranslateExchangeIdRequestSourceIdType `json:"SourceIdType,omitempty"`
	TargetIdType *CreateUserByIdTranslateExchangeIdRequestTargetIdType `json:"TargetIdType,omitempty"`
}
