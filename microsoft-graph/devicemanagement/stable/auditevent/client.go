package auditevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditEventClient struct {
	Client *msgraph.Client
}

func NewAuditEventClientWithBaseURI(sdkApi sdkEnv.Api) (*AuditEventClient, error) {
	client, err := msgraph.NewClient(sdkApi, "auditevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuditEventClient: %+v", err)
	}

	return &AuditEventClient{
		Client: client,
	}, nil
}
