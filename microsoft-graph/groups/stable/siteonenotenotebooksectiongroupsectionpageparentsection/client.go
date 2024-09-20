package siteonenotenotebooksectiongroupsectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionGroupSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionGroupSectionPageParentSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectiongroupsectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionGroupSectionPageParentSectionClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionGroupSectionPageParentSectionClient{
		Client: client,
	}, nil
}
