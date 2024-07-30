package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/beta/auditlog"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/beta/customsecurityattributeaudit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/beta/directoryaudit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/beta/directoryprovisioning"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/beta/provisioning"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/beta/signin"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AuditLog                     *auditlog.AuditLogClient
	CustomSecurityAttributeAudit *customsecurityattributeaudit.CustomSecurityAttributeAuditClient
	DirectoryAudit               *directoryaudit.DirectoryAuditClient
	DirectoryProvisioning        *directoryprovisioning.DirectoryProvisioningClient
	Provisioning                 *provisioning.ProvisioningClient
	SignIn                       *signin.SignInClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	auditLogClient, err := auditlog.NewAuditLogClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuditLog client: %+v", err)
	}
	configureFunc(auditLogClient.Client)

	customSecurityAttributeAuditClient, err := customsecurityattributeaudit.NewCustomSecurityAttributeAuditClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomSecurityAttributeAudit client: %+v", err)
	}
	configureFunc(customSecurityAttributeAuditClient.Client)

	directoryAuditClient, err := directoryaudit.NewDirectoryAuditClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryAudit client: %+v", err)
	}
	configureFunc(directoryAuditClient.Client)

	directoryProvisioningClient, err := directoryprovisioning.NewDirectoryProvisioningClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryProvisioning client: %+v", err)
	}
	configureFunc(directoryProvisioningClient.Client)

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
		AuditLog:                     auditLogClient,
		CustomSecurityAttributeAudit: customSecurityAttributeAuditClient,
		DirectoryAudit:               directoryAuditClient,
		DirectoryProvisioning:        directoryProvisioningClient,
		Provisioning:                 provisioningClient,
		SignIn:                       signInClient,
	}, nil
}
