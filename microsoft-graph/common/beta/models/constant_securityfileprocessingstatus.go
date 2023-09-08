package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileProcessingStatus string

const (
	SecurityFileProcessingStatusextractionException         SecurityFileProcessingStatus = "ExtractionException"
	SecurityFileProcessingStatusfileBodyIsTooLong           SecurityFileProcessingStatus = "FileBodyIsTooLong"
	SecurityFileProcessingStatusfileDepthLimitExceeded      SecurityFileProcessingStatus = "FileDepthLimitExceeded"
	SecurityFileProcessingStatusfileSizeIsTooLarge          SecurityFileProcessingStatus = "FileSizeIsTooLarge"
	SecurityFileProcessingStatusfileSizeIsZero              SecurityFileProcessingStatus = "FileSizeIsZero"
	SecurityFileProcessingStatusfileTypeIsNotSupported      SecurityFileProcessingStatus = "FileTypeIsNotSupported"
	SecurityFileProcessingStatusfileTypeIsUnknown           SecurityFileProcessingStatus = "FileTypeIsUnknown"
	SecurityFileProcessingStatusinternalError               SecurityFileProcessingStatus = "InternalError"
	SecurityFileProcessingStatusinvalidFileId               SecurityFileProcessingStatus = "InvalidFileId"
	SecurityFileProcessingStatusmalformedFile               SecurityFileProcessingStatus = "MalformedFile"
	SecurityFileProcessingStatusnoReviewSetSummaryGenerated SecurityFileProcessingStatus = "NoReviewSetSummaryGenerated"
	SecurityFileProcessingStatusocrFileSizeExceedsLimit     SecurityFileProcessingStatus = "OcrFileSizeExceedsLimit"
	SecurityFileProcessingStatusocrProcessingTimeout        SecurityFileProcessingStatus = "OcrProcessingTimeout"
	SecurityFileProcessingStatuspoisonFile                  SecurityFileProcessingStatus = "PoisonFile"
	SecurityFileProcessingStatusprocessingTimeout           SecurityFileProcessingStatus = "ProcessingTimeout"
	SecurityFileProcessingStatusprotectedFile               SecurityFileProcessingStatus = "ProtectedFile"
	SecurityFileProcessingStatussuccess                     SecurityFileProcessingStatus = "Success"
	SecurityFileProcessingStatusunknownError                SecurityFileProcessingStatus = "UnknownError"
)

func PossibleValuesForSecurityFileProcessingStatus() []string {
	return []string{
		string(SecurityFileProcessingStatusextractionException),
		string(SecurityFileProcessingStatusfileBodyIsTooLong),
		string(SecurityFileProcessingStatusfileDepthLimitExceeded),
		string(SecurityFileProcessingStatusfileSizeIsTooLarge),
		string(SecurityFileProcessingStatusfileSizeIsZero),
		string(SecurityFileProcessingStatusfileTypeIsNotSupported),
		string(SecurityFileProcessingStatusfileTypeIsUnknown),
		string(SecurityFileProcessingStatusinternalError),
		string(SecurityFileProcessingStatusinvalidFileId),
		string(SecurityFileProcessingStatusmalformedFile),
		string(SecurityFileProcessingStatusnoReviewSetSummaryGenerated),
		string(SecurityFileProcessingStatusocrFileSizeExceedsLimit),
		string(SecurityFileProcessingStatusocrProcessingTimeout),
		string(SecurityFileProcessingStatuspoisonFile),
		string(SecurityFileProcessingStatusprocessingTimeout),
		string(SecurityFileProcessingStatusprotectedFile),
		string(SecurityFileProcessingStatussuccess),
		string(SecurityFileProcessingStatusunknownError),
	}
}

func parseSecurityFileProcessingStatus(input string) (*SecurityFileProcessingStatus, error) {
	vals := map[string]SecurityFileProcessingStatus{
		"extractionexception":         SecurityFileProcessingStatusextractionException,
		"filebodyistoolong":           SecurityFileProcessingStatusfileBodyIsTooLong,
		"filedepthlimitexceeded":      SecurityFileProcessingStatusfileDepthLimitExceeded,
		"filesizeistoolarge":          SecurityFileProcessingStatusfileSizeIsTooLarge,
		"filesizeiszero":              SecurityFileProcessingStatusfileSizeIsZero,
		"filetypeisnotsupported":      SecurityFileProcessingStatusfileTypeIsNotSupported,
		"filetypeisunknown":           SecurityFileProcessingStatusfileTypeIsUnknown,
		"internalerror":               SecurityFileProcessingStatusinternalError,
		"invalidfileid":               SecurityFileProcessingStatusinvalidFileId,
		"malformedfile":               SecurityFileProcessingStatusmalformedFile,
		"noreviewsetsummarygenerated": SecurityFileProcessingStatusnoReviewSetSummaryGenerated,
		"ocrfilesizeexceedslimit":     SecurityFileProcessingStatusocrFileSizeExceedsLimit,
		"ocrprocessingtimeout":        SecurityFileProcessingStatusocrProcessingTimeout,
		"poisonfile":                  SecurityFileProcessingStatuspoisonFile,
		"processingtimeout":           SecurityFileProcessingStatusprocessingTimeout,
		"protectedfile":               SecurityFileProcessingStatusprotectedFile,
		"success":                     SecurityFileProcessingStatussuccess,
		"unknownerror":                SecurityFileProcessingStatusunknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileProcessingStatus(input)
	return &out, nil
}
