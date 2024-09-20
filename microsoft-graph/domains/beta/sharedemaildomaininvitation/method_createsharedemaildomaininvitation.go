package sharedemaildomaininvitation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSharedEmailDomainInvitationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SharedEmailDomainInvitation
}

type CreateSharedEmailDomainInvitationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateSharedEmailDomainInvitationOperationOptions() CreateSharedEmailDomainInvitationOperationOptions {
	return CreateSharedEmailDomainInvitationOperationOptions{}
}

func (o CreateSharedEmailDomainInvitationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSharedEmailDomainInvitationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSharedEmailDomainInvitationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSharedEmailDomainInvitation - Create new navigation property to sharedEmailDomainInvitations for domains
func (c SharedEmailDomainInvitationClient) CreateSharedEmailDomainInvitation(ctx context.Context, id beta.DomainId, input beta.SharedEmailDomainInvitation, options CreateSharedEmailDomainInvitationOperationOptions) (result CreateSharedEmailDomainInvitationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sharedEmailDomainInvitations", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.SharedEmailDomainInvitation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
