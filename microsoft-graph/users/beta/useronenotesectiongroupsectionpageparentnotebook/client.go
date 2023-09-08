package useronenotesectiongroupsectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionGroupSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionGroupSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionGroupSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectiongroupsectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionGroupSectionPageParentNotebookClient: %+v", err)
	}

	return &UserOnenoteSectionGroupSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
