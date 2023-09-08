package grouponenotesectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionParentSectionGroupClient: %+v", err)
	}

	return &GroupOnenoteSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
