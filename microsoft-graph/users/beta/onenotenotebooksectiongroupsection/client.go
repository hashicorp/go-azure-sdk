package onenotenotebooksectiongroupsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionGroupSectionClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionGroupSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionGroupSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectiongroupsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionGroupSectionClient: %+v", err)
	}

	return &OnenoteNotebookSectionGroupSectionClient{
		Client: client,
	}, nil
}
