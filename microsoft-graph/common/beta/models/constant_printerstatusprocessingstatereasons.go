package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterStatusProcessingStateReasons string

const (
	PrinterStatusProcessingStateReasonsconnectingToDevice                 PrinterStatusProcessingStateReasons = "ConnectingToDevice"
	PrinterStatusProcessingStateReasonscoverOpen                          PrinterStatusProcessingStateReasons = "CoverOpen"
	PrinterStatusProcessingStateReasonsdeveloperEmpty                     PrinterStatusProcessingStateReasons = "DeveloperEmpty"
	PrinterStatusProcessingStateReasonsdeveloperLow                       PrinterStatusProcessingStateReasons = "DeveloperLow"
	PrinterStatusProcessingStateReasonsdoorOpen                           PrinterStatusProcessingStateReasons = "DoorOpen"
	PrinterStatusProcessingStateReasonsfuserOverTemp                      PrinterStatusProcessingStateReasons = "FuserOverTemp"
	PrinterStatusProcessingStateReasonsfuserUnderTemp                     PrinterStatusProcessingStateReasons = "FuserUnderTemp"
	PrinterStatusProcessingStateReasonsinputTrayMissing                   PrinterStatusProcessingStateReasons = "InputTrayMissing"
	PrinterStatusProcessingStateReasonsinterlockOpen                      PrinterStatusProcessingStateReasons = "InterlockOpen"
	PrinterStatusProcessingStateReasonsinterpreterResourceUnavailable     PrinterStatusProcessingStateReasons = "InterpreterResourceUnavailable"
	PrinterStatusProcessingStateReasonsmarkerSupplyEmpty                  PrinterStatusProcessingStateReasons = "MarkerSupplyEmpty"
	PrinterStatusProcessingStateReasonsmarkerSupplyLow                    PrinterStatusProcessingStateReasons = "MarkerSupplyLow"
	PrinterStatusProcessingStateReasonsmarkerWasteAlmostFull              PrinterStatusProcessingStateReasons = "MarkerWasteAlmostFull"
	PrinterStatusProcessingStateReasonsmarkerWasteFull                    PrinterStatusProcessingStateReasons = "MarkerWasteFull"
	PrinterStatusProcessingStateReasonsmediaEmpty                         PrinterStatusProcessingStateReasons = "MediaEmpty"
	PrinterStatusProcessingStateReasonsmediaJam                           PrinterStatusProcessingStateReasons = "MediaJam"
	PrinterStatusProcessingStateReasonsmediaLow                           PrinterStatusProcessingStateReasons = "MediaLow"
	PrinterStatusProcessingStateReasonsmediaNeeded                        PrinterStatusProcessingStateReasons = "MediaNeeded"
	PrinterStatusProcessingStateReasonsmovingToPaused                     PrinterStatusProcessingStateReasons = "MovingToPaused"
	PrinterStatusProcessingStateReasonsnone                               PrinterStatusProcessingStateReasons = "None"
	PrinterStatusProcessingStateReasonsopticalPhotoConductorLifeOver      PrinterStatusProcessingStateReasons = "OpticalPhotoConductorLifeOver"
	PrinterStatusProcessingStateReasonsopticalPhotoConductorNearEndOfLife PrinterStatusProcessingStateReasons = "OpticalPhotoConductorNearEndOfLife"
	PrinterStatusProcessingStateReasonsother                              PrinterStatusProcessingStateReasons = "Other"
	PrinterStatusProcessingStateReasonsoutputAreaAlmostFull               PrinterStatusProcessingStateReasons = "OutputAreaAlmostFull"
	PrinterStatusProcessingStateReasonsoutputAreaFull                     PrinterStatusProcessingStateReasons = "OutputAreaFull"
	PrinterStatusProcessingStateReasonsoutputTrayMissing                  PrinterStatusProcessingStateReasons = "OutputTrayMissing"
	PrinterStatusProcessingStateReasonspaused                             PrinterStatusProcessingStateReasons = "Paused"
	PrinterStatusProcessingStateReasonsshutdown                           PrinterStatusProcessingStateReasons = "Shutdown"
	PrinterStatusProcessingStateReasonsspoolAreaFull                      PrinterStatusProcessingStateReasons = "SpoolAreaFull"
	PrinterStatusProcessingStateReasonsstoppedPartially                   PrinterStatusProcessingStateReasons = "StoppedPartially"
	PrinterStatusProcessingStateReasonsstopping                           PrinterStatusProcessingStateReasons = "Stopping"
	PrinterStatusProcessingStateReasonstimedOut                           PrinterStatusProcessingStateReasons = "TimedOut"
	PrinterStatusProcessingStateReasonstonerEmpty                         PrinterStatusProcessingStateReasons = "TonerEmpty"
	PrinterStatusProcessingStateReasonstonerLow                           PrinterStatusProcessingStateReasons = "TonerLow"
)

func PossibleValuesForPrinterStatusProcessingStateReasons() []string {
	return []string{
		string(PrinterStatusProcessingStateReasonsconnectingToDevice),
		string(PrinterStatusProcessingStateReasonscoverOpen),
		string(PrinterStatusProcessingStateReasonsdeveloperEmpty),
		string(PrinterStatusProcessingStateReasonsdeveloperLow),
		string(PrinterStatusProcessingStateReasonsdoorOpen),
		string(PrinterStatusProcessingStateReasonsfuserOverTemp),
		string(PrinterStatusProcessingStateReasonsfuserUnderTemp),
		string(PrinterStatusProcessingStateReasonsinputTrayMissing),
		string(PrinterStatusProcessingStateReasonsinterlockOpen),
		string(PrinterStatusProcessingStateReasonsinterpreterResourceUnavailable),
		string(PrinterStatusProcessingStateReasonsmarkerSupplyEmpty),
		string(PrinterStatusProcessingStateReasonsmarkerSupplyLow),
		string(PrinterStatusProcessingStateReasonsmarkerWasteAlmostFull),
		string(PrinterStatusProcessingStateReasonsmarkerWasteFull),
		string(PrinterStatusProcessingStateReasonsmediaEmpty),
		string(PrinterStatusProcessingStateReasonsmediaJam),
		string(PrinterStatusProcessingStateReasonsmediaLow),
		string(PrinterStatusProcessingStateReasonsmediaNeeded),
		string(PrinterStatusProcessingStateReasonsmovingToPaused),
		string(PrinterStatusProcessingStateReasonsnone),
		string(PrinterStatusProcessingStateReasonsopticalPhotoConductorLifeOver),
		string(PrinterStatusProcessingStateReasonsopticalPhotoConductorNearEndOfLife),
		string(PrinterStatusProcessingStateReasonsother),
		string(PrinterStatusProcessingStateReasonsoutputAreaAlmostFull),
		string(PrinterStatusProcessingStateReasonsoutputAreaFull),
		string(PrinterStatusProcessingStateReasonsoutputTrayMissing),
		string(PrinterStatusProcessingStateReasonspaused),
		string(PrinterStatusProcessingStateReasonsshutdown),
		string(PrinterStatusProcessingStateReasonsspoolAreaFull),
		string(PrinterStatusProcessingStateReasonsstoppedPartially),
		string(PrinterStatusProcessingStateReasonsstopping),
		string(PrinterStatusProcessingStateReasonstimedOut),
		string(PrinterStatusProcessingStateReasonstonerEmpty),
		string(PrinterStatusProcessingStateReasonstonerLow),
	}
}

func parsePrinterStatusProcessingStateReasons(input string) (*PrinterStatusProcessingStateReasons, error) {
	vals := map[string]PrinterStatusProcessingStateReasons{
		"connectingtodevice":                 PrinterStatusProcessingStateReasonsconnectingToDevice,
		"coveropen":                          PrinterStatusProcessingStateReasonscoverOpen,
		"developerempty":                     PrinterStatusProcessingStateReasonsdeveloperEmpty,
		"developerlow":                       PrinterStatusProcessingStateReasonsdeveloperLow,
		"dooropen":                           PrinterStatusProcessingStateReasonsdoorOpen,
		"fuserovertemp":                      PrinterStatusProcessingStateReasonsfuserOverTemp,
		"fuserundertemp":                     PrinterStatusProcessingStateReasonsfuserUnderTemp,
		"inputtraymissing":                   PrinterStatusProcessingStateReasonsinputTrayMissing,
		"interlockopen":                      PrinterStatusProcessingStateReasonsinterlockOpen,
		"interpreterresourceunavailable":     PrinterStatusProcessingStateReasonsinterpreterResourceUnavailable,
		"markersupplyempty":                  PrinterStatusProcessingStateReasonsmarkerSupplyEmpty,
		"markersupplylow":                    PrinterStatusProcessingStateReasonsmarkerSupplyLow,
		"markerwastealmostfull":              PrinterStatusProcessingStateReasonsmarkerWasteAlmostFull,
		"markerwastefull":                    PrinterStatusProcessingStateReasonsmarkerWasteFull,
		"mediaempty":                         PrinterStatusProcessingStateReasonsmediaEmpty,
		"mediajam":                           PrinterStatusProcessingStateReasonsmediaJam,
		"medialow":                           PrinterStatusProcessingStateReasonsmediaLow,
		"medianeeded":                        PrinterStatusProcessingStateReasonsmediaNeeded,
		"movingtopaused":                     PrinterStatusProcessingStateReasonsmovingToPaused,
		"none":                               PrinterStatusProcessingStateReasonsnone,
		"opticalphotoconductorlifeover":      PrinterStatusProcessingStateReasonsopticalPhotoConductorLifeOver,
		"opticalphotoconductornearendoflife": PrinterStatusProcessingStateReasonsopticalPhotoConductorNearEndOfLife,
		"other":                              PrinterStatusProcessingStateReasonsother,
		"outputareaalmostfull":               PrinterStatusProcessingStateReasonsoutputAreaAlmostFull,
		"outputareafull":                     PrinterStatusProcessingStateReasonsoutputAreaFull,
		"outputtraymissing":                  PrinterStatusProcessingStateReasonsoutputTrayMissing,
		"paused":                             PrinterStatusProcessingStateReasonspaused,
		"shutdown":                           PrinterStatusProcessingStateReasonsshutdown,
		"spoolareafull":                      PrinterStatusProcessingStateReasonsspoolAreaFull,
		"stoppedpartially":                   PrinterStatusProcessingStateReasonsstoppedPartially,
		"stopping":                           PrinterStatusProcessingStateReasonsstopping,
		"timedout":                           PrinterStatusProcessingStateReasonstimedOut,
		"tonerempty":                         PrinterStatusProcessingStateReasonstonerEmpty,
		"tonerlow":                           PrinterStatusProcessingStateReasonstonerLow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusProcessingStateReasons(input)
	return &out, nil
}
