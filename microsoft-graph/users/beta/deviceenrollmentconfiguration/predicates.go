package deviceenrollmentconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceEnrollmentConfigurationOperationPredicate struct {
}

func (p DeviceEnrollmentConfigurationOperationPredicate) Matches(input beta.DeviceEnrollmentConfiguration) bool {

	return true
}

type HasPayloadLinkResultItemOperationPredicate struct {
}

func (p HasPayloadLinkResultItemOperationPredicate) Matches(input beta.HasPayloadLinkResultItem) bool {

	return true
}
