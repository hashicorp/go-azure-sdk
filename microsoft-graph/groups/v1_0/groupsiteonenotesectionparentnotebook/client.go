package groupsiteonenotesectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionParentNotebookClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionParentNotebookClient{
		Client: client,
	}, nil
}
