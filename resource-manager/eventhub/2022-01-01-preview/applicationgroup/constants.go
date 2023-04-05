package applicationgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGroupPolicyType string

const (
	ApplicationGroupPolicyTypeThrottlingPolicy ApplicationGroupPolicyType = "ThrottlingPolicy"
)

func PossibleValuesForApplicationGroupPolicyType() []string {
	return []string{
		string(ApplicationGroupPolicyTypeThrottlingPolicy),
	}
}

type MetricId string

const (
	MetricIdIncomingBytes    MetricId = "IncomingBytes"
	MetricIdIncomingMessages MetricId = "IncomingMessages"
	MetricIdOutgoingBytes    MetricId = "OutgoingBytes"
	MetricIdOutgoingMessages MetricId = "OutgoingMessages"
)

func PossibleValuesForMetricId() []string {
	return []string{
		string(MetricIdIncomingBytes),
		string(MetricIdIncomingMessages),
		string(MetricIdOutgoingBytes),
		string(MetricIdOutgoingMessages),
	}
}
