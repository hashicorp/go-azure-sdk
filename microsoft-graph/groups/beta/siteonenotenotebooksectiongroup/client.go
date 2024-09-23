package siteonenotenotebooksectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionGroupClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionGroupClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionGroupClient{
		Client: client,
	}, nil
}
