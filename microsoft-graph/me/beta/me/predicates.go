package me

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type ConvertIdResultOperationPredicate struct {
}

func (p ConvertIdResultOperationPredicate) Matches(input beta.ConvertIdResult) bool {

	return true
}

type MailTipsOperationPredicate struct {
}

func (p MailTipsOperationPredicate) Matches(input beta.MailTips) bool {

	return true
}

type PasswordSingleSignOnCredentialSetOperationPredicate struct {
}

func (p PasswordSingleSignOnCredentialSetOperationPredicate) Matches(input beta.PasswordSingleSignOnCredentialSet) bool {

	return true
}
