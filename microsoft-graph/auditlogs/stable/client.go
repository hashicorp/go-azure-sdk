package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/stable/auditlog"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/stable/directoryaudit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/stable/provisioning"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/stable/signin"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AuditLog       *auditlog.AuditLogClient
	DirectoryAudit *directoryaudit.DirectoryAuditClient
	Provisioning   *provisioning.ProvisioningClient
	SignIn         *signin.SignInClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	auditLogClient, err := auditlog.NewAuditLogClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuditLog client: %+v", err)
	}
	configureFunc(auditLogClient.Client)

	directoryAuditClient, err := directoryaudit.NewDirectoryAuditClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryAudit client: %+v", err)
	}
	configureFunc(directoryAuditClient.Client)

	provisioningClient, err := provisioning.NewProvisioningClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Provisioning client: %+v", err)
	}
	configureFunc(provisioningClient.Client)

	signInClient, err := signin.NewSignInClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SignIn client: %+v", err)
	}
	configureFunc(signInClient.Client)

	return &Client{
		AuditLog:       auditLogClient,
		DirectoryAudit: directoryAuditClient,
		Provisioning:   provisioningClient,
		SignIn:         signInClient,
	}, nil
}
