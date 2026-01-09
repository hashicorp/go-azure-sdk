package delegatedpermissionclassification

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedPermissionClassificationClient struct {
	Client *msgraph.Client
}

func NewDelegatedPermissionClassificationClientWithBaseURI(sdkApi sdkEnv.Api) (*DelegatedPermissionClassificationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "delegatedpermissionclassification", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DelegatedPermissionClassificationClient: %+v", err)
	}

	return &DelegatedPermissionClassificationClient{
		Client: client,
	}, nil
}
