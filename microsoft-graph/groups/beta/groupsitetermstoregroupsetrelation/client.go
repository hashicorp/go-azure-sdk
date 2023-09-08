package groupsitetermstoregroupsetrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetRelationClient{
		Client: client,
	}, nil
}
