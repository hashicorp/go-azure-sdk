package onenotenotebooksectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionPageParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionPageParentNotebookClient: %+v", err)
	}

	return &OnenoteNotebookSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
