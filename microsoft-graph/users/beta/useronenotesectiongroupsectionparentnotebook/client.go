package useronenotesectiongroupsectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionGroupSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionGroupSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectiongroupsectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionGroupSectionParentNotebookClient: %+v", err)
	}

	return &UserOnenoteSectionGroupSectionParentNotebookClient{
		Client: client,
	}, nil
}
