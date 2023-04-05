package integrationaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventLevel string

const (
	EventLevelCritical      EventLevel = "Critical"
	EventLevelError         EventLevel = "Error"
	EventLevelInformational EventLevel = "Informational"
	EventLevelLogAlways     EventLevel = "LogAlways"
	EventLevelVerbose       EventLevel = "Verbose"
	EventLevelWarning       EventLevel = "Warning"
)

func PossibleValuesForEventLevel() []string {
	return []string{
		string(EventLevelCritical),
		string(EventLevelError),
		string(EventLevelInformational),
		string(EventLevelLogAlways),
		string(EventLevelVerbose),
		string(EventLevelWarning),
	}
}

type IntegrationAccountSkuName string

const (
	IntegrationAccountSkuNameBasic        IntegrationAccountSkuName = "Basic"
	IntegrationAccountSkuNameFree         IntegrationAccountSkuName = "Free"
	IntegrationAccountSkuNameNotSpecified IntegrationAccountSkuName = "NotSpecified"
	IntegrationAccountSkuNameStandard     IntegrationAccountSkuName = "Standard"
)

func PossibleValuesForIntegrationAccountSkuName() []string {
	return []string{
		string(IntegrationAccountSkuNameBasic),
		string(IntegrationAccountSkuNameFree),
		string(IntegrationAccountSkuNameNotSpecified),
		string(IntegrationAccountSkuNameStandard),
	}
}

type KeyType string

const (
	KeyTypeNotSpecified KeyType = "NotSpecified"
	KeyTypePrimary      KeyType = "Primary"
	KeyTypeSecondary    KeyType = "Secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypeNotSpecified),
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}

type TrackEventsOperationOptions string

const (
	TrackEventsOperationOptionsDisableSourceInfoEnrich TrackEventsOperationOptions = "DisableSourceInfoEnrich"
	TrackEventsOperationOptionsNone                    TrackEventsOperationOptions = "None"
)

func PossibleValuesForTrackEventsOperationOptions() []string {
	return []string{
		string(TrackEventsOperationOptionsDisableSourceInfoEnrich),
		string(TrackEventsOperationOptionsNone),
	}
}

type TrackingRecordType string

const (
	TrackingRecordTypeASTwoMDN                             TrackingRecordType = "AS2MDN"
	TrackingRecordTypeASTwoMessage                         TrackingRecordType = "AS2Message"
	TrackingRecordTypeCustom                               TrackingRecordType = "Custom"
	TrackingRecordTypeEdifactFunctionalGroup               TrackingRecordType = "EdifactFunctionalGroup"
	TrackingRecordTypeEdifactFunctionalGroupAcknowledgment TrackingRecordType = "EdifactFunctionalGroupAcknowledgment"
	TrackingRecordTypeEdifactInterchange                   TrackingRecordType = "EdifactInterchange"
	TrackingRecordTypeEdifactInterchangeAcknowledgment     TrackingRecordType = "EdifactInterchangeAcknowledgment"
	TrackingRecordTypeEdifactTransactionSet                TrackingRecordType = "EdifactTransactionSet"
	TrackingRecordTypeEdifactTransactionSetAcknowledgment  TrackingRecordType = "EdifactTransactionSetAcknowledgment"
	TrackingRecordTypeNotSpecified                         TrackingRecordType = "NotSpecified"
	TrackingRecordTypeXOneTwoFunctionalGroup               TrackingRecordType = "X12FunctionalGroup"
	TrackingRecordTypeXOneTwoFunctionalGroupAcknowledgment TrackingRecordType = "X12FunctionalGroupAcknowledgment"
	TrackingRecordTypeXOneTwoInterchange                   TrackingRecordType = "X12Interchange"
	TrackingRecordTypeXOneTwoInterchangeAcknowledgment     TrackingRecordType = "X12InterchangeAcknowledgment"
	TrackingRecordTypeXOneTwoTransactionSet                TrackingRecordType = "X12TransactionSet"
	TrackingRecordTypeXOneTwoTransactionSetAcknowledgment  TrackingRecordType = "X12TransactionSetAcknowledgment"
)

func PossibleValuesForTrackingRecordType() []string {
	return []string{
		string(TrackingRecordTypeASTwoMDN),
		string(TrackingRecordTypeASTwoMessage),
		string(TrackingRecordTypeCustom),
		string(TrackingRecordTypeEdifactFunctionalGroup),
		string(TrackingRecordTypeEdifactFunctionalGroupAcknowledgment),
		string(TrackingRecordTypeEdifactInterchange),
		string(TrackingRecordTypeEdifactInterchangeAcknowledgment),
		string(TrackingRecordTypeEdifactTransactionSet),
		string(TrackingRecordTypeEdifactTransactionSetAcknowledgment),
		string(TrackingRecordTypeNotSpecified),
		string(TrackingRecordTypeXOneTwoFunctionalGroup),
		string(TrackingRecordTypeXOneTwoFunctionalGroupAcknowledgment),
		string(TrackingRecordTypeXOneTwoInterchange),
		string(TrackingRecordTypeXOneTwoInterchangeAcknowledgment),
		string(TrackingRecordTypeXOneTwoTransactionSet),
		string(TrackingRecordTypeXOneTwoTransactionSetAcknowledgment),
	}
}

type WorkflowState string

const (
	WorkflowStateCompleted    WorkflowState = "Completed"
	WorkflowStateDeleted      WorkflowState = "Deleted"
	WorkflowStateDisabled     WorkflowState = "Disabled"
	WorkflowStateEnabled      WorkflowState = "Enabled"
	WorkflowStateNotSpecified WorkflowState = "NotSpecified"
	WorkflowStateSuspended    WorkflowState = "Suspended"
)

func PossibleValuesForWorkflowState() []string {
	return []string{
		string(WorkflowStateCompleted),
		string(WorkflowStateDeleted),
		string(WorkflowStateDisabled),
		string(WorkflowStateEnabled),
		string(WorkflowStateNotSpecified),
		string(WorkflowStateSuspended),
	}
}
