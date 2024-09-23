package grouppolicyobjectfile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyObjectFileClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyObjectFileClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyObjectFileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicyobjectfile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyObjectFileClient: %+v", err)
	}

	return &GroupPolicyObjectFileClient{
		Client: client,
	}, nil
}
