package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppContentFileUploadState string

const (
	MobileAppContentFileUploadStateazureStorageUriRenewalFailed   MobileAppContentFileUploadState = "AzureStorageUriRenewalFailed"
	MobileAppContentFileUploadStateazureStorageUriRenewalPending  MobileAppContentFileUploadState = "AzureStorageUriRenewalPending"
	MobileAppContentFileUploadStateazureStorageUriRenewalSuccess  MobileAppContentFileUploadState = "AzureStorageUriRenewalSuccess"
	MobileAppContentFileUploadStateazureStorageUriRenewalTimedOut MobileAppContentFileUploadState = "AzureStorageUriRenewalTimedOut"
	MobileAppContentFileUploadStateazureStorageUriRequestFailed   MobileAppContentFileUploadState = "AzureStorageUriRequestFailed"
	MobileAppContentFileUploadStateazureStorageUriRequestPending  MobileAppContentFileUploadState = "AzureStorageUriRequestPending"
	MobileAppContentFileUploadStateazureStorageUriRequestSuccess  MobileAppContentFileUploadState = "AzureStorageUriRequestSuccess"
	MobileAppContentFileUploadStateazureStorageUriRequestTimedOut MobileAppContentFileUploadState = "AzureStorageUriRequestTimedOut"
	MobileAppContentFileUploadStatecommitFileFailed               MobileAppContentFileUploadState = "CommitFileFailed"
	MobileAppContentFileUploadStatecommitFilePending              MobileAppContentFileUploadState = "CommitFilePending"
	MobileAppContentFileUploadStatecommitFileSuccess              MobileAppContentFileUploadState = "CommitFileSuccess"
	MobileAppContentFileUploadStatecommitFileTimedOut             MobileAppContentFileUploadState = "CommitFileTimedOut"
	MobileAppContentFileUploadStateerror                          MobileAppContentFileUploadState = "Error"
	MobileAppContentFileUploadStatesuccess                        MobileAppContentFileUploadState = "Success"
	MobileAppContentFileUploadStatetransientError                 MobileAppContentFileUploadState = "TransientError"
	MobileAppContentFileUploadStateunknown                        MobileAppContentFileUploadState = "Unknown"
)

func PossibleValuesForMobileAppContentFileUploadState() []string {
	return []string{
		string(MobileAppContentFileUploadStateazureStorageUriRenewalFailed),
		string(MobileAppContentFileUploadStateazureStorageUriRenewalPending),
		string(MobileAppContentFileUploadStateazureStorageUriRenewalSuccess),
		string(MobileAppContentFileUploadStateazureStorageUriRenewalTimedOut),
		string(MobileAppContentFileUploadStateazureStorageUriRequestFailed),
		string(MobileAppContentFileUploadStateazureStorageUriRequestPending),
		string(MobileAppContentFileUploadStateazureStorageUriRequestSuccess),
		string(MobileAppContentFileUploadStateazureStorageUriRequestTimedOut),
		string(MobileAppContentFileUploadStatecommitFileFailed),
		string(MobileAppContentFileUploadStatecommitFilePending),
		string(MobileAppContentFileUploadStatecommitFileSuccess),
		string(MobileAppContentFileUploadStatecommitFileTimedOut),
		string(MobileAppContentFileUploadStateerror),
		string(MobileAppContentFileUploadStatesuccess),
		string(MobileAppContentFileUploadStatetransientError),
		string(MobileAppContentFileUploadStateunknown),
	}
}

func parseMobileAppContentFileUploadState(input string) (*MobileAppContentFileUploadState, error) {
	vals := map[string]MobileAppContentFileUploadState{
		"azurestorageurirenewalfailed":   MobileAppContentFileUploadStateazureStorageUriRenewalFailed,
		"azurestorageurirenewalpending":  MobileAppContentFileUploadStateazureStorageUriRenewalPending,
		"azurestorageurirenewalsuccess":  MobileAppContentFileUploadStateazureStorageUriRenewalSuccess,
		"azurestorageurirenewaltimedout": MobileAppContentFileUploadStateazureStorageUriRenewalTimedOut,
		"azurestorageurirequestfailed":   MobileAppContentFileUploadStateazureStorageUriRequestFailed,
		"azurestorageurirequestpending":  MobileAppContentFileUploadStateazureStorageUriRequestPending,
		"azurestorageurirequestsuccess":  MobileAppContentFileUploadStateazureStorageUriRequestSuccess,
		"azurestorageurirequesttimedout": MobileAppContentFileUploadStateazureStorageUriRequestTimedOut,
		"commitfilefailed":               MobileAppContentFileUploadStatecommitFileFailed,
		"commitfilepending":              MobileAppContentFileUploadStatecommitFilePending,
		"commitfilesuccess":              MobileAppContentFileUploadStatecommitFileSuccess,
		"commitfiletimedout":             MobileAppContentFileUploadStatecommitFileTimedOut,
		"error":                          MobileAppContentFileUploadStateerror,
		"success":                        MobileAppContentFileUploadStatesuccess,
		"transienterror":                 MobileAppContentFileUploadStatetransientError,
		"unknown":                        MobileAppContentFileUploadStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppContentFileUploadState(input)
	return &out, nil
}
