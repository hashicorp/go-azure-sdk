package groupsitetermstoresetchildren

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetChildrenClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetChildrenClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetChildrenClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetchildren", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetChildrenClient: %+v", err)
	}

	return &GroupSiteTermStoreSetChildrenClient{
		Client: client,
	}, nil
}
