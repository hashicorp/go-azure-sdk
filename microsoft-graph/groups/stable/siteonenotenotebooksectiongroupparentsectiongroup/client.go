package siteonenotenotebooksectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
