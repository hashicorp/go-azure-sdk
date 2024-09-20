package onenotenotebooksectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &OnenoteNotebookSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
