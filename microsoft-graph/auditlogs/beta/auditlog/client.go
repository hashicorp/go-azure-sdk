package auditlog

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditLogClient struct {
	Client *msgraph.Client
}

func NewAuditLogClientWithBaseURI(sdkApi sdkEnv.Api) (*AuditLogClient, error) {
	client, err := msgraph.NewClient(sdkApi, "auditlog", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuditLogClient: %+v", err)
	}

	return &AuditLogClient{
		Client: client,
	}, nil
}
