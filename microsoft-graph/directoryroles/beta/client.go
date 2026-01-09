package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/beta/directoryrole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/beta/member"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/beta/scopedmember"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DirectoryRole *directoryrole.DirectoryRoleClient
	Member        *member.MemberClient
	ScopedMember  *scopedmember.ScopedMemberClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	directoryRoleClient, err := directoryrole.NewDirectoryRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRole client: %+v", err)
	}
	configureFunc(directoryRoleClient.Client)

	memberClient, err := member.NewMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Member client: %+v", err)
	}
	configureFunc(memberClient.Client)

	scopedMemberClient, err := scopedmember.NewScopedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScopedMember client: %+v", err)
	}
	configureFunc(scopedMemberClient.Client)

	return &Client{
		DirectoryRole: directoryRoleClient,
		Member:        memberClient,
		ScopedMember:  scopedMemberClient,
	}, nil
}
