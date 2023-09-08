package useronenotesectiongroupsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionGroupSectionClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionGroupSectionClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionGroupSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectiongroupsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionGroupSectionClient: %+v", err)
	}

	return &UserOnenoteSectionGroupSectionClient{
		Client: client,
	}, nil
}
