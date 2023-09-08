package groupsitetermstoregroupsetrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetRelationToTermClient{
		Client: client,
	}, nil
}
