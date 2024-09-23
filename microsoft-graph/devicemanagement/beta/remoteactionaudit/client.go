package remoteactionaudit

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteActionAuditClient struct {
	Client *msgraph.Client
}

func NewRemoteActionAuditClientWithBaseURI(sdkApi sdkEnv.Api) (*RemoteActionAuditClient, error) {
	client, err := msgraph.NewClient(sdkApi, "remoteactionaudit", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RemoteActionAuditClient: %+v", err)
	}

	return &RemoteActionAuditClient{
		Client: client,
	}, nil
}
