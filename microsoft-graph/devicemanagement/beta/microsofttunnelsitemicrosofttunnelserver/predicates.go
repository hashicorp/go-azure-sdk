package microsofttunnelsitemicrosofttunnelserver

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type KeyLongValuePairOperationPredicate struct {
}

func (p KeyLongValuePairOperationPredicate) Matches(input beta.KeyLongValuePair) bool {

	return true
}

type MetricTimeSeriesDataPointOperationPredicate struct {
}

func (p MetricTimeSeriesDataPointOperationPredicate) Matches(input beta.MetricTimeSeriesDataPoint) bool {

	return true
}

type MicrosoftTunnelServerOperationPredicate struct {
}

func (p MicrosoftTunnelServerOperationPredicate) Matches(input beta.MicrosoftTunnelServer) bool {

	return true
}
