package siteonenotenotebooksection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionClient{
		Client: client,
	}, nil
}
