package eventsources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventSourceKind string

const (
	EventSourceKindMicrosoftPointEventHub EventSourceKind = "Microsoft.EventHub"
	EventSourceKindMicrosoftPointIoTHub   EventSourceKind = "Microsoft.IoTHub"
)

func PossibleValuesForEventSourceKind() []string {
	return []string{
		string(EventSourceKindMicrosoftPointEventHub),
		string(EventSourceKindMicrosoftPointIoTHub),
	}
}

type IngressStartAtType string

const (
	IngressStartAtTypeCustomEnqueuedTime      IngressStartAtType = "CustomEnqueuedTime"
	IngressStartAtTypeEarliestAvailable       IngressStartAtType = "EarliestAvailable"
	IngressStartAtTypeEventSourceCreationTime IngressStartAtType = "EventSourceCreationTime"
)

func PossibleValuesForIngressStartAtType() []string {
	return []string{
		string(IngressStartAtTypeCustomEnqueuedTime),
		string(IngressStartAtTypeEarliestAvailable),
		string(IngressStartAtTypeEventSourceCreationTime),
	}
}

type Kind string

const (
	KindMicrosoftPointEventHub Kind = "Microsoft.EventHub"
	KindMicrosoftPointIoTHub   Kind = "Microsoft.IoTHub"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindMicrosoftPointEventHub),
		string(KindMicrosoftPointIoTHub),
	}
}

type LocalTimestampFormat string

const (
	LocalTimestampFormatEmbedded LocalTimestampFormat = "Embedded"
)

func PossibleValuesForLocalTimestampFormat() []string {
	return []string{
		string(LocalTimestampFormatEmbedded),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}
