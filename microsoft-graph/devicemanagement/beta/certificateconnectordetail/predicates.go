package certificateconnectordetail

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type CertificateConnectorDetailsOperationPredicate struct {
}

func (p CertificateConnectorDetailsOperationPredicate) Matches(input beta.CertificateConnectorDetails) bool {

	return true
}

type CertificateConnectorHealthMetricValueOperationPredicate struct {
}

func (p CertificateConnectorHealthMetricValueOperationPredicate) Matches(input beta.CertificateConnectorHealthMetricValue) bool {

	return true
}

type KeyLongValuePairOperationPredicate struct {
}

func (p KeyLongValuePairOperationPredicate) Matches(input beta.KeyLongValuePair) bool {

	return true
}
