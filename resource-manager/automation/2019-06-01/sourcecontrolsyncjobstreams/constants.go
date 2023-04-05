package sourcecontrolsyncjobstreams

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamType string

const (
	StreamTypeError  StreamType = "Error"
	StreamTypeOutput StreamType = "Output"
)

func PossibleValuesForStreamType() []string {
	return []string{
		string(StreamTypeError),
		string(StreamTypeOutput),
	}
}
