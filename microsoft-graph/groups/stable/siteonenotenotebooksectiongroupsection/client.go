package siteonenotenotebooksectiongroupsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionGroupSectionClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionGroupSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionGroupSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectiongroupsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionGroupSectionClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionGroupSectionClient{
		Client: client,
	}, nil
}
