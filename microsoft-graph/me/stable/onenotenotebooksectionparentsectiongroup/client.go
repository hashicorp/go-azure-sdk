package onenotenotebooksectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionParentSectionGroupClient: %+v", err)
	}

	return &OnenoteNotebookSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
