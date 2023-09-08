package meonenoteoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteOperationClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteOperationClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenoteoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteOperationClient: %+v", err)
	}

	return &MeOnenoteOperationClient{
		Client: client,
	}, nil
}
