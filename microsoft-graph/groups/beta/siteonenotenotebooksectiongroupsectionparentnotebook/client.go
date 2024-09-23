package siteonenotenotebooksectiongroupsectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionGroupSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionGroupSectionParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectiongroupsectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionGroupSectionParentNotebookClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionGroupSectionParentNotebookClient{
		Client: client,
	}, nil
}
