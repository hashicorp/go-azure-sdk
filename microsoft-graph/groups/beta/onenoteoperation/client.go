package onenoteoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteOperationClient struct {
	Client *msgraph.Client
}

func NewOnenoteOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenoteoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteOperationClient: %+v", err)
	}

	return &OnenoteOperationClient{
		Client: client,
	}, nil
}
