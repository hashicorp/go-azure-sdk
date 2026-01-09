package siteonenotenotebooksectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionParentSectionGroupClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
