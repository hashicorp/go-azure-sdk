package onenotenotebooksectiongroupsectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionGroupSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionGroupSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectiongroupsectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionGroupSectionParentSectionGroupClient: %+v", err)
	}

	return &OnenoteNotebookSectionGroupSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
