package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryFileProcessingStatus string

const (
	SecurityEdiscoveryFileProcessingStatusextractionException         SecurityEdiscoveryFileProcessingStatus = "ExtractionException"
	SecurityEdiscoveryFileProcessingStatusfileBodyIsTooLong           SecurityEdiscoveryFileProcessingStatus = "FileBodyIsTooLong"
	SecurityEdiscoveryFileProcessingStatusfileDepthLimitExceeded      SecurityEdiscoveryFileProcessingStatus = "FileDepthLimitExceeded"
	SecurityEdiscoveryFileProcessingStatusfileSizeIsTooLarge          SecurityEdiscoveryFileProcessingStatus = "FileSizeIsTooLarge"
	SecurityEdiscoveryFileProcessingStatusfileSizeIsZero              SecurityEdiscoveryFileProcessingStatus = "FileSizeIsZero"
	SecurityEdiscoveryFileProcessingStatusfileTypeIsNotSupported      SecurityEdiscoveryFileProcessingStatus = "FileTypeIsNotSupported"
	SecurityEdiscoveryFileProcessingStatusfileTypeIsUnknown           SecurityEdiscoveryFileProcessingStatus = "FileTypeIsUnknown"
	SecurityEdiscoveryFileProcessingStatusinternalError               SecurityEdiscoveryFileProcessingStatus = "InternalError"
	SecurityEdiscoveryFileProcessingStatusinvalidFileId               SecurityEdiscoveryFileProcessingStatus = "InvalidFileId"
	SecurityEdiscoveryFileProcessingStatusmalformedFile               SecurityEdiscoveryFileProcessingStatus = "MalformedFile"
	SecurityEdiscoveryFileProcessingStatusnoReviewSetSummaryGenerated SecurityEdiscoveryFileProcessingStatus = "NoReviewSetSummaryGenerated"
	SecurityEdiscoveryFileProcessingStatusocrFileSizeExceedsLimit     SecurityEdiscoveryFileProcessingStatus = "OcrFileSizeExceedsLimit"
	SecurityEdiscoveryFileProcessingStatusocrProcessingTimeout        SecurityEdiscoveryFileProcessingStatus = "OcrProcessingTimeout"
	SecurityEdiscoveryFileProcessingStatuspoisonFile                  SecurityEdiscoveryFileProcessingStatus = "PoisonFile"
	SecurityEdiscoveryFileProcessingStatusprocessingTimeout           SecurityEdiscoveryFileProcessingStatus = "ProcessingTimeout"
	SecurityEdiscoveryFileProcessingStatusprotectedFile               SecurityEdiscoveryFileProcessingStatus = "ProtectedFile"
	SecurityEdiscoveryFileProcessingStatussuccess                     SecurityEdiscoveryFileProcessingStatus = "Success"
	SecurityEdiscoveryFileProcessingStatusunknownError                SecurityEdiscoveryFileProcessingStatus = "UnknownError"
)

func PossibleValuesForSecurityEdiscoveryFileProcessingStatus() []string {
	return []string{
		string(SecurityEdiscoveryFileProcessingStatusextractionException),
		string(SecurityEdiscoveryFileProcessingStatusfileBodyIsTooLong),
		string(SecurityEdiscoveryFileProcessingStatusfileDepthLimitExceeded),
		string(SecurityEdiscoveryFileProcessingStatusfileSizeIsTooLarge),
		string(SecurityEdiscoveryFileProcessingStatusfileSizeIsZero),
		string(SecurityEdiscoveryFileProcessingStatusfileTypeIsNotSupported),
		string(SecurityEdiscoveryFileProcessingStatusfileTypeIsUnknown),
		string(SecurityEdiscoveryFileProcessingStatusinternalError),
		string(SecurityEdiscoveryFileProcessingStatusinvalidFileId),
		string(SecurityEdiscoveryFileProcessingStatusmalformedFile),
		string(SecurityEdiscoveryFileProcessingStatusnoReviewSetSummaryGenerated),
		string(SecurityEdiscoveryFileProcessingStatusocrFileSizeExceedsLimit),
		string(SecurityEdiscoveryFileProcessingStatusocrProcessingTimeout),
		string(SecurityEdiscoveryFileProcessingStatuspoisonFile),
		string(SecurityEdiscoveryFileProcessingStatusprocessingTimeout),
		string(SecurityEdiscoveryFileProcessingStatusprotectedFile),
		string(SecurityEdiscoveryFileProcessingStatussuccess),
		string(SecurityEdiscoveryFileProcessingStatusunknownError),
	}
}

func parseSecurityEdiscoveryFileProcessingStatus(input string) (*SecurityEdiscoveryFileProcessingStatus, error) {
	vals := map[string]SecurityEdiscoveryFileProcessingStatus{
		"extractionexception":         SecurityEdiscoveryFileProcessingStatusextractionException,
		"filebodyistoolong":           SecurityEdiscoveryFileProcessingStatusfileBodyIsTooLong,
		"filedepthlimitexceeded":      SecurityEdiscoveryFileProcessingStatusfileDepthLimitExceeded,
		"filesizeistoolarge":          SecurityEdiscoveryFileProcessingStatusfileSizeIsTooLarge,
		"filesizeiszero":              SecurityEdiscoveryFileProcessingStatusfileSizeIsZero,
		"filetypeisnotsupported":      SecurityEdiscoveryFileProcessingStatusfileTypeIsNotSupported,
		"filetypeisunknown":           SecurityEdiscoveryFileProcessingStatusfileTypeIsUnknown,
		"internalerror":               SecurityEdiscoveryFileProcessingStatusinternalError,
		"invalidfileid":               SecurityEdiscoveryFileProcessingStatusinvalidFileId,
		"malformedfile":               SecurityEdiscoveryFileProcessingStatusmalformedFile,
		"noreviewsetsummarygenerated": SecurityEdiscoveryFileProcessingStatusnoReviewSetSummaryGenerated,
		"ocrfilesizeexceedslimit":     SecurityEdiscoveryFileProcessingStatusocrFileSizeExceedsLimit,
		"ocrprocessingtimeout":        SecurityEdiscoveryFileProcessingStatusocrProcessingTimeout,
		"poisonfile":                  SecurityEdiscoveryFileProcessingStatuspoisonFile,
		"processingtimeout":           SecurityEdiscoveryFileProcessingStatusprocessingTimeout,
		"protectedfile":               SecurityEdiscoveryFileProcessingStatusprotectedFile,
		"success":                     SecurityEdiscoveryFileProcessingStatussuccess,
		"unknownerror":                SecurityEdiscoveryFileProcessingStatusunknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryFileProcessingStatus(input)
	return &out, nil
}
