package volumesreplication

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointType string

const (
	EndpointTypeDst EndpointType = "dst"
	EndpointTypeSrc EndpointType = "src"
)

func PossibleValuesForEndpointType() []string {
	return []string{
		string(EndpointTypeDst),
		string(EndpointTypeSrc),
	}
}

type MirrorState string

const (
	MirrorStateBroken        MirrorState = "Broken"
	MirrorStateMirrored      MirrorState = "Mirrored"
	MirrorStateUninitialized MirrorState = "Uninitialized"
)

func PossibleValuesForMirrorState() []string {
	return []string{
		string(MirrorStateBroken),
		string(MirrorStateMirrored),
		string(MirrorStateUninitialized),
	}
}

type RelationshipStatus string

const (
	RelationshipStatusIdle         RelationshipStatus = "Idle"
	RelationshipStatusTransferring RelationshipStatus = "Transferring"
)

func PossibleValuesForRelationshipStatus() []string {
	return []string{
		string(RelationshipStatusIdle),
		string(RelationshipStatusTransferring),
	}
}

type ReplicationSchedule string

const (
	ReplicationScheduleDaily           ReplicationSchedule = "daily"
	ReplicationScheduleHourly          ReplicationSchedule = "hourly"
	ReplicationScheduleOneZerominutely ReplicationSchedule = "_10minutely"
)

func PossibleValuesForReplicationSchedule() []string {
	return []string{
		string(ReplicationScheduleDaily),
		string(ReplicationScheduleHourly),
		string(ReplicationScheduleOneZerominutely),
	}
}
