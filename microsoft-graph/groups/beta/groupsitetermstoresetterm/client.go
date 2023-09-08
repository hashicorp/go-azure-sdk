package groupsitetermstoresetterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetTermClient{
		Client: client,
	}, nil
}
