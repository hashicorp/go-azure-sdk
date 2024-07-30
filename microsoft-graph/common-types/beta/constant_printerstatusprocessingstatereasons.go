package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterStatusProcessingStateReasons string

const (
	PrinterStatusProcessingStateReasons_ConnectingToDevice                 PrinterStatusProcessingStateReasons = "connectingToDevice"
	PrinterStatusProcessingStateReasons_CoverOpen                          PrinterStatusProcessingStateReasons = "coverOpen"
	PrinterStatusProcessingStateReasons_DeveloperEmpty                     PrinterStatusProcessingStateReasons = "developerEmpty"
	PrinterStatusProcessingStateReasons_DeveloperLow                       PrinterStatusProcessingStateReasons = "developerLow"
	PrinterStatusProcessingStateReasons_DoorOpen                           PrinterStatusProcessingStateReasons = "doorOpen"
	PrinterStatusProcessingStateReasons_FuserOverTemp                      PrinterStatusProcessingStateReasons = "fuserOverTemp"
	PrinterStatusProcessingStateReasons_FuserUnderTemp                     PrinterStatusProcessingStateReasons = "fuserUnderTemp"
	PrinterStatusProcessingStateReasons_InputTrayMissing                   PrinterStatusProcessingStateReasons = "inputTrayMissing"
	PrinterStatusProcessingStateReasons_InterlockOpen                      PrinterStatusProcessingStateReasons = "interlockOpen"
	PrinterStatusProcessingStateReasons_InterpreterResourceUnavailable     PrinterStatusProcessingStateReasons = "interpreterResourceUnavailable"
	PrinterStatusProcessingStateReasons_MarkerSupplyEmpty                  PrinterStatusProcessingStateReasons = "markerSupplyEmpty"
	PrinterStatusProcessingStateReasons_MarkerSupplyLow                    PrinterStatusProcessingStateReasons = "markerSupplyLow"
	PrinterStatusProcessingStateReasons_MarkerWasteAlmostFull              PrinterStatusProcessingStateReasons = "markerWasteAlmostFull"
	PrinterStatusProcessingStateReasons_MarkerWasteFull                    PrinterStatusProcessingStateReasons = "markerWasteFull"
	PrinterStatusProcessingStateReasons_MediaEmpty                         PrinterStatusProcessingStateReasons = "mediaEmpty"
	PrinterStatusProcessingStateReasons_MediaJam                           PrinterStatusProcessingStateReasons = "mediaJam"
	PrinterStatusProcessingStateReasons_MediaLow                           PrinterStatusProcessingStateReasons = "mediaLow"
	PrinterStatusProcessingStateReasons_MediaNeeded                        PrinterStatusProcessingStateReasons = "mediaNeeded"
	PrinterStatusProcessingStateReasons_MovingToPaused                     PrinterStatusProcessingStateReasons = "movingToPaused"
	PrinterStatusProcessingStateReasons_None                               PrinterStatusProcessingStateReasons = "none"
	PrinterStatusProcessingStateReasons_OpticalPhotoConductorLifeOver      PrinterStatusProcessingStateReasons = "opticalPhotoConductorLifeOver"
	PrinterStatusProcessingStateReasons_OpticalPhotoConductorNearEndOfLife PrinterStatusProcessingStateReasons = "opticalPhotoConductorNearEndOfLife"
	PrinterStatusProcessingStateReasons_Other                              PrinterStatusProcessingStateReasons = "other"
	PrinterStatusProcessingStateReasons_OutputAreaAlmostFull               PrinterStatusProcessingStateReasons = "outputAreaAlmostFull"
	PrinterStatusProcessingStateReasons_OutputAreaFull                     PrinterStatusProcessingStateReasons = "outputAreaFull"
	PrinterStatusProcessingStateReasons_OutputTrayMissing                  PrinterStatusProcessingStateReasons = "outputTrayMissing"
	PrinterStatusProcessingStateReasons_Paused                             PrinterStatusProcessingStateReasons = "paused"
	PrinterStatusProcessingStateReasons_Shutdown                           PrinterStatusProcessingStateReasons = "shutdown"
	PrinterStatusProcessingStateReasons_SpoolAreaFull                      PrinterStatusProcessingStateReasons = "spoolAreaFull"
	PrinterStatusProcessingStateReasons_StoppedPartially                   PrinterStatusProcessingStateReasons = "stoppedPartially"
	PrinterStatusProcessingStateReasons_Stopping                           PrinterStatusProcessingStateReasons = "stopping"
	PrinterStatusProcessingStateReasons_TimedOut                           PrinterStatusProcessingStateReasons = "timedOut"
	PrinterStatusProcessingStateReasons_TonerEmpty                         PrinterStatusProcessingStateReasons = "tonerEmpty"
	PrinterStatusProcessingStateReasons_TonerLow                           PrinterStatusProcessingStateReasons = "tonerLow"
)

func PossibleValuesForPrinterStatusProcessingStateReasons() []string {
	return []string{
		string(PrinterStatusProcessingStateReasons_ConnectingToDevice),
		string(PrinterStatusProcessingStateReasons_CoverOpen),
		string(PrinterStatusProcessingStateReasons_DeveloperEmpty),
		string(PrinterStatusProcessingStateReasons_DeveloperLow),
		string(PrinterStatusProcessingStateReasons_DoorOpen),
		string(PrinterStatusProcessingStateReasons_FuserOverTemp),
		string(PrinterStatusProcessingStateReasons_FuserUnderTemp),
		string(PrinterStatusProcessingStateReasons_InputTrayMissing),
		string(PrinterStatusProcessingStateReasons_InterlockOpen),
		string(PrinterStatusProcessingStateReasons_InterpreterResourceUnavailable),
		string(PrinterStatusProcessingStateReasons_MarkerSupplyEmpty),
		string(PrinterStatusProcessingStateReasons_MarkerSupplyLow),
		string(PrinterStatusProcessingStateReasons_MarkerWasteAlmostFull),
		string(PrinterStatusProcessingStateReasons_MarkerWasteFull),
		string(PrinterStatusProcessingStateReasons_MediaEmpty),
		string(PrinterStatusProcessingStateReasons_MediaJam),
		string(PrinterStatusProcessingStateReasons_MediaLow),
		string(PrinterStatusProcessingStateReasons_MediaNeeded),
		string(PrinterStatusProcessingStateReasons_MovingToPaused),
		string(PrinterStatusProcessingStateReasons_None),
		string(PrinterStatusProcessingStateReasons_OpticalPhotoConductorLifeOver),
		string(PrinterStatusProcessingStateReasons_OpticalPhotoConductorNearEndOfLife),
		string(PrinterStatusProcessingStateReasons_Other),
		string(PrinterStatusProcessingStateReasons_OutputAreaAlmostFull),
		string(PrinterStatusProcessingStateReasons_OutputAreaFull),
		string(PrinterStatusProcessingStateReasons_OutputTrayMissing),
		string(PrinterStatusProcessingStateReasons_Paused),
		string(PrinterStatusProcessingStateReasons_Shutdown),
		string(PrinterStatusProcessingStateReasons_SpoolAreaFull),
		string(PrinterStatusProcessingStateReasons_StoppedPartially),
		string(PrinterStatusProcessingStateReasons_Stopping),
		string(PrinterStatusProcessingStateReasons_TimedOut),
		string(PrinterStatusProcessingStateReasons_TonerEmpty),
		string(PrinterStatusProcessingStateReasons_TonerLow),
	}
}

func (s *PrinterStatusProcessingStateReasons) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterStatusProcessingStateReasons(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterStatusProcessingStateReasons(input string) (*PrinterStatusProcessingStateReasons, error) {
	vals := map[string]PrinterStatusProcessingStateReasons{
		"connectingtodevice":                 PrinterStatusProcessingStateReasons_ConnectingToDevice,
		"coveropen":                          PrinterStatusProcessingStateReasons_CoverOpen,
		"developerempty":                     PrinterStatusProcessingStateReasons_DeveloperEmpty,
		"developerlow":                       PrinterStatusProcessingStateReasons_DeveloperLow,
		"dooropen":                           PrinterStatusProcessingStateReasons_DoorOpen,
		"fuserovertemp":                      PrinterStatusProcessingStateReasons_FuserOverTemp,
		"fuserundertemp":                     PrinterStatusProcessingStateReasons_FuserUnderTemp,
		"inputtraymissing":                   PrinterStatusProcessingStateReasons_InputTrayMissing,
		"interlockopen":                      PrinterStatusProcessingStateReasons_InterlockOpen,
		"interpreterresourceunavailable":     PrinterStatusProcessingStateReasons_InterpreterResourceUnavailable,
		"markersupplyempty":                  PrinterStatusProcessingStateReasons_MarkerSupplyEmpty,
		"markersupplylow":                    PrinterStatusProcessingStateReasons_MarkerSupplyLow,
		"markerwastealmostfull":              PrinterStatusProcessingStateReasons_MarkerWasteAlmostFull,
		"markerwastefull":                    PrinterStatusProcessingStateReasons_MarkerWasteFull,
		"mediaempty":                         PrinterStatusProcessingStateReasons_MediaEmpty,
		"mediajam":                           PrinterStatusProcessingStateReasons_MediaJam,
		"medialow":                           PrinterStatusProcessingStateReasons_MediaLow,
		"medianeeded":                        PrinterStatusProcessingStateReasons_MediaNeeded,
		"movingtopaused":                     PrinterStatusProcessingStateReasons_MovingToPaused,
		"none":                               PrinterStatusProcessingStateReasons_None,
		"opticalphotoconductorlifeover":      PrinterStatusProcessingStateReasons_OpticalPhotoConductorLifeOver,
		"opticalphotoconductornearendoflife": PrinterStatusProcessingStateReasons_OpticalPhotoConductorNearEndOfLife,
		"other":                              PrinterStatusProcessingStateReasons_Other,
		"outputareaalmostfull":               PrinterStatusProcessingStateReasons_OutputAreaAlmostFull,
		"outputareafull":                     PrinterStatusProcessingStateReasons_OutputAreaFull,
		"outputtraymissing":                  PrinterStatusProcessingStateReasons_OutputTrayMissing,
		"paused":                             PrinterStatusProcessingStateReasons_Paused,
		"shutdown":                           PrinterStatusProcessingStateReasons_Shutdown,
		"spoolareafull":                      PrinterStatusProcessingStateReasons_SpoolAreaFull,
		"stoppedpartially":                   PrinterStatusProcessingStateReasons_StoppedPartially,
		"stopping":                           PrinterStatusProcessingStateReasons_Stopping,
		"timedout":                           PrinterStatusProcessingStateReasons_TimedOut,
		"tonerempty":                         PrinterStatusProcessingStateReasons_TonerEmpty,
		"tonerlow":                           PrinterStatusProcessingStateReasons_TonerLow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusProcessingStateReasons(input)
	return &out, nil
}
