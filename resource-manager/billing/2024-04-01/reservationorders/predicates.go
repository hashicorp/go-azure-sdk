package reservationorders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationOrderOperationPredicate struct {
	Etag *int64
	Id   *string
	Name *string
	Type *string
}

func (p ReservationOrderOperationPredicate) Matches(input ReservationOrder) bool {

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
