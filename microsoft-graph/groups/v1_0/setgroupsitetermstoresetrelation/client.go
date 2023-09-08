package setgroupsitetermstoresetrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresetrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetRelationClient{
		Client: client,
	}, nil
}
