package environmentcontainer

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssetProvisioningState string

const (
	AssetProvisioningStateCanceled  AssetProvisioningState = "Canceled"
	AssetProvisioningStateCreating  AssetProvisioningState = "Creating"
	AssetProvisioningStateDeleting  AssetProvisioningState = "Deleting"
	AssetProvisioningStateFailed    AssetProvisioningState = "Failed"
	AssetProvisioningStateSucceeded AssetProvisioningState = "Succeeded"
	AssetProvisioningStateUpdating  AssetProvisioningState = "Updating"
)

func PossibleValuesForAssetProvisioningState() []string {
	return []string{
		string(AssetProvisioningStateCanceled),
		string(AssetProvisioningStateCreating),
		string(AssetProvisioningStateDeleting),
		string(AssetProvisioningStateFailed),
		string(AssetProvisioningStateSucceeded),
		string(AssetProvisioningStateUpdating),
	}
}

func parseAssetProvisioningState(input string) (*AssetProvisioningState, error) {
	vals := map[string]AssetProvisioningState{
		"canceled":  AssetProvisioningStateCanceled,
		"creating":  AssetProvisioningStateCreating,
		"deleting":  AssetProvisioningStateDeleting,
		"failed":    AssetProvisioningStateFailed,
		"succeeded": AssetProvisioningStateSucceeded,
		"updating":  AssetProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssetProvisioningState(input)
	return &out, nil
}

type ListViewType string

const (
	ListViewTypeActiveOnly   ListViewType = "ActiveOnly"
	ListViewTypeAll          ListViewType = "All"
	ListViewTypeArchivedOnly ListViewType = "ArchivedOnly"
)

func PossibleValuesForListViewType() []string {
	return []string{
		string(ListViewTypeActiveOnly),
		string(ListViewTypeAll),
		string(ListViewTypeArchivedOnly),
	}
}

func parseListViewType(input string) (*ListViewType, error) {
	vals := map[string]ListViewType{
		"activeonly":   ListViewTypeActiveOnly,
		"all":          ListViewTypeAll,
		"archivedonly": ListViewTypeArchivedOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ListViewType(input)
	return &out, nil
}
