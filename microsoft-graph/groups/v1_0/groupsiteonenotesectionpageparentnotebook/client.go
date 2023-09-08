package groupsiteonenotesectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionPageParentNotebookClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
