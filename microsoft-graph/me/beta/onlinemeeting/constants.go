package onlinemeeting

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesType string

const (
	CreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesTypeMins15 CreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesType = "mins15"
)

func PossibleValuesForCreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesType() []string {
	return []string{
		string(CreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesTypeMins15),
	}
}

func (s *CreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesType(input string) (*CreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesType, error) {
	vals := map[string]CreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesType{
		"mins15": CreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesTypeMins15,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesType(input)
	return &out, nil
}

type CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType string

const (
	CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsTypeCancellation CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType = "cancellation"
	CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsTypeConfirmation CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType = "confirmation"
	CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsTypeReschedule   CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType = "reschedule"
)

func PossibleValuesForCreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType() []string {
	return []string{
		string(CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsTypeCancellation),
		string(CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsTypeConfirmation),
		string(CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsTypeReschedule),
	}
}

func (s *CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType(input string) (*CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType, error) {
	vals := map[string]CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType{
		"cancellation": CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsTypeCancellation,
		"confirmation": CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsTypeConfirmation,
		"reschedule":   CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsTypeReschedule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateOnlineMeetingSendVirtualAppointmentSmRequestSmsType(input)
	return &out, nil
}

type MeetingParticipantInfoRole string

const (
	MeetingParticipantInfoRoleAttendee    MeetingParticipantInfoRole = "attendee"
	MeetingParticipantInfoRoleCoorganizer MeetingParticipantInfoRole = "coorganizer"
	MeetingParticipantInfoRolePresenter   MeetingParticipantInfoRole = "presenter"
	MeetingParticipantInfoRoleProducer    MeetingParticipantInfoRole = "producer"
)

func PossibleValuesForMeetingParticipantInfoRole() []string {
	return []string{
		string(MeetingParticipantInfoRoleAttendee),
		string(MeetingParticipantInfoRoleCoorganizer),
		string(MeetingParticipantInfoRolePresenter),
		string(MeetingParticipantInfoRoleProducer),
	}
}

func (s *MeetingParticipantInfoRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingParticipantInfoRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingParticipantInfoRole(input string) (*MeetingParticipantInfoRole, error) {
	vals := map[string]MeetingParticipantInfoRole{
		"attendee":    MeetingParticipantInfoRoleAttendee,
		"coorganizer": MeetingParticipantInfoRoleCoorganizer,
		"presenter":   MeetingParticipantInfoRolePresenter,
		"producer":    MeetingParticipantInfoRoleProducer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingParticipantInfoRole(input)
	return &out, nil
}
