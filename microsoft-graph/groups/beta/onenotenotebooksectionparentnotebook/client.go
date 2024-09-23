package onenotenotebooksectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionParentNotebookClient: %+v", err)
	}

	return &OnenoteNotebookSectionParentNotebookClient{
		Client: client,
	}, nil
}
