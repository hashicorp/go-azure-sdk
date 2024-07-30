package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryFileProcessingStatus string

const (
	SecurityEdiscoveryFileProcessingStatus_ExtractionException         SecurityEdiscoveryFileProcessingStatus = "extractionException"
	SecurityEdiscoveryFileProcessingStatus_FileBodyIsTooLong           SecurityEdiscoveryFileProcessingStatus = "fileBodyIsTooLong"
	SecurityEdiscoveryFileProcessingStatus_FileDepthLimitExceeded      SecurityEdiscoveryFileProcessingStatus = "fileDepthLimitExceeded"
	SecurityEdiscoveryFileProcessingStatus_FileSizeIsTooLarge          SecurityEdiscoveryFileProcessingStatus = "fileSizeIsTooLarge"
	SecurityEdiscoveryFileProcessingStatus_FileSizeIsZero              SecurityEdiscoveryFileProcessingStatus = "fileSizeIsZero"
	SecurityEdiscoveryFileProcessingStatus_FileTypeIsNotSupported      SecurityEdiscoveryFileProcessingStatus = "fileTypeIsNotSupported"
	SecurityEdiscoveryFileProcessingStatus_FileTypeIsUnknown           SecurityEdiscoveryFileProcessingStatus = "fileTypeIsUnknown"
	SecurityEdiscoveryFileProcessingStatus_InternalError               SecurityEdiscoveryFileProcessingStatus = "internalError"
	SecurityEdiscoveryFileProcessingStatus_InvalidFileId               SecurityEdiscoveryFileProcessingStatus = "invalidFileId"
	SecurityEdiscoveryFileProcessingStatus_MalformedFile               SecurityEdiscoveryFileProcessingStatus = "malformedFile"
	SecurityEdiscoveryFileProcessingStatus_NoReviewSetSummaryGenerated SecurityEdiscoveryFileProcessingStatus = "noReviewSetSummaryGenerated"
	SecurityEdiscoveryFileProcessingStatus_OcrFileSizeExceedsLimit     SecurityEdiscoveryFileProcessingStatus = "ocrFileSizeExceedsLimit"
	SecurityEdiscoveryFileProcessingStatus_OcrProcessingTimeout        SecurityEdiscoveryFileProcessingStatus = "ocrProcessingTimeout"
	SecurityEdiscoveryFileProcessingStatus_PoisonFile                  SecurityEdiscoveryFileProcessingStatus = "poisonFile"
	SecurityEdiscoveryFileProcessingStatus_ProcessingTimeout           SecurityEdiscoveryFileProcessingStatus = "processingTimeout"
	SecurityEdiscoveryFileProcessingStatus_ProtectedFile               SecurityEdiscoveryFileProcessingStatus = "protectedFile"
	SecurityEdiscoveryFileProcessingStatus_Success                     SecurityEdiscoveryFileProcessingStatus = "success"
	SecurityEdiscoveryFileProcessingStatus_UnknownError                SecurityEdiscoveryFileProcessingStatus = "unknownError"
)

func PossibleValuesForSecurityEdiscoveryFileProcessingStatus() []string {
	return []string{
		string(SecurityEdiscoveryFileProcessingStatus_ExtractionException),
		string(SecurityEdiscoveryFileProcessingStatus_FileBodyIsTooLong),
		string(SecurityEdiscoveryFileProcessingStatus_FileDepthLimitExceeded),
		string(SecurityEdiscoveryFileProcessingStatus_FileSizeIsTooLarge),
		string(SecurityEdiscoveryFileProcessingStatus_FileSizeIsZero),
		string(SecurityEdiscoveryFileProcessingStatus_FileTypeIsNotSupported),
		string(SecurityEdiscoveryFileProcessingStatus_FileTypeIsUnknown),
		string(SecurityEdiscoveryFileProcessingStatus_InternalError),
		string(SecurityEdiscoveryFileProcessingStatus_InvalidFileId),
		string(SecurityEdiscoveryFileProcessingStatus_MalformedFile),
		string(SecurityEdiscoveryFileProcessingStatus_NoReviewSetSummaryGenerated),
		string(SecurityEdiscoveryFileProcessingStatus_OcrFileSizeExceedsLimit),
		string(SecurityEdiscoveryFileProcessingStatus_OcrProcessingTimeout),
		string(SecurityEdiscoveryFileProcessingStatus_PoisonFile),
		string(SecurityEdiscoveryFileProcessingStatus_ProcessingTimeout),
		string(SecurityEdiscoveryFileProcessingStatus_ProtectedFile),
		string(SecurityEdiscoveryFileProcessingStatus_Success),
		string(SecurityEdiscoveryFileProcessingStatus_UnknownError),
	}
}

func (s *SecurityEdiscoveryFileProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryFileProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryFileProcessingStatus(input string) (*SecurityEdiscoveryFileProcessingStatus, error) {
	vals := map[string]SecurityEdiscoveryFileProcessingStatus{
		"extractionexception":         SecurityEdiscoveryFileProcessingStatus_ExtractionException,
		"filebodyistoolong":           SecurityEdiscoveryFileProcessingStatus_FileBodyIsTooLong,
		"filedepthlimitexceeded":      SecurityEdiscoveryFileProcessingStatus_FileDepthLimitExceeded,
		"filesizeistoolarge":          SecurityEdiscoveryFileProcessingStatus_FileSizeIsTooLarge,
		"filesizeiszero":              SecurityEdiscoveryFileProcessingStatus_FileSizeIsZero,
		"filetypeisnotsupported":      SecurityEdiscoveryFileProcessingStatus_FileTypeIsNotSupported,
		"filetypeisunknown":           SecurityEdiscoveryFileProcessingStatus_FileTypeIsUnknown,
		"internalerror":               SecurityEdiscoveryFileProcessingStatus_InternalError,
		"invalidfileid":               SecurityEdiscoveryFileProcessingStatus_InvalidFileId,
		"malformedfile":               SecurityEdiscoveryFileProcessingStatus_MalformedFile,
		"noreviewsetsummarygenerated": SecurityEdiscoveryFileProcessingStatus_NoReviewSetSummaryGenerated,
		"ocrfilesizeexceedslimit":     SecurityEdiscoveryFileProcessingStatus_OcrFileSizeExceedsLimit,
		"ocrprocessingtimeout":        SecurityEdiscoveryFileProcessingStatus_OcrProcessingTimeout,
		"poisonfile":                  SecurityEdiscoveryFileProcessingStatus_PoisonFile,
		"processingtimeout":           SecurityEdiscoveryFileProcessingStatus_ProcessingTimeout,
		"protectedfile":               SecurityEdiscoveryFileProcessingStatus_ProtectedFile,
		"success":                     SecurityEdiscoveryFileProcessingStatus_Success,
		"unknownerror":                SecurityEdiscoveryFileProcessingStatus_UnknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryFileProcessingStatus(input)
	return &out, nil
}
