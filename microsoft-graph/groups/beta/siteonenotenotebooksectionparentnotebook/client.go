package siteonenotenotebooksectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionParentNotebookClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionParentNotebookClient{
		Client: client,
	}, nil
}
