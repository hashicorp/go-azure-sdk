package actiongroupsapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReceiverStatus string

const (
	ReceiverStatusDisabled     ReceiverStatus = "Disabled"
	ReceiverStatusEnabled      ReceiverStatus = "Enabled"
	ReceiverStatusNotSpecified ReceiverStatus = "NotSpecified"
)

func PossibleValuesForReceiverStatus() []string {
	return []string{
		string(ReceiverStatusDisabled),
		string(ReceiverStatusEnabled),
		string(ReceiverStatusNotSpecified),
	}
}
