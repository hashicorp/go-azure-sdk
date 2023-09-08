package groupsitetermstoresettermrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetTermRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetTermRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetTermRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresettermrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetTermRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetTermRelationToTermClient{
		Client: client,
	}, nil
}
