package onenotenotebooksectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionGroupClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionGroupClient: %+v", err)
	}

	return &OnenoteNotebookSectionGroupClient{
		Client: client,
	}, nil
}
