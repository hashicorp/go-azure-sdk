package dscnode

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscNodeOperationPredicate struct {
	AccountId        *string
	Etag             *string
	IP               *string
	Id               *string
	LastSeen         *string
	Name             *string
	NodeId           *string
	RegistrationTime *string
	Status           *string
	Type             *string
}

func (p DscNodeOperationPredicate) Matches(input DscNode) bool {

	if p.AccountId != nil && (input.AccountId == nil || *p.AccountId != *input.AccountId) {
		return false
	}

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

	if p.IP != nil && (input.IP == nil || *p.IP != *input.IP) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastSeen != nil && (input.LastSeen == nil || *p.LastSeen != *input.LastSeen) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.NodeId != nil && (input.NodeId == nil || *p.NodeId != *input.NodeId) {
		return false
	}

	if p.RegistrationTime != nil && (input.RegistrationTime == nil || *p.RegistrationTime != *input.RegistrationTime) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
