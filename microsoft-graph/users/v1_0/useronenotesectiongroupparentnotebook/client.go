package useronenotesectiongroupparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionGroupParentNotebookClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionGroupParentNotebookClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionGroupParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectiongroupparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionGroupParentNotebookClient: %+v", err)
	}

	return &UserOnenoteSectionGroupParentNotebookClient{
		Client: client,
	}, nil
}
