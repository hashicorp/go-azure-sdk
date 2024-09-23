package onenotenotebooksectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionPageContentClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionPageContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionPageContentClient: %+v", err)
	}

	return &OnenoteNotebookSectionPageContentClient{
		Client: client,
	}, nil
}
