package groupsitetermstoresetrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreSetRelationClient{
		Client: client,
	}, nil
}
