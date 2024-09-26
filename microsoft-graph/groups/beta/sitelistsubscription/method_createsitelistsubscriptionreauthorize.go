package sitelistsubscription

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

type CreateSiteListSubscriptionReauthorizeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateSiteListSubscriptionReauthorizeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSiteListSubscriptionReauthorizeOperationOptions() CreateSiteListSubscriptionReauthorizeOperationOptions {
	return CreateSiteListSubscriptionReauthorizeOperationOptions{}
}

func (o CreateSiteListSubscriptionReauthorizeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSiteListSubscriptionReauthorizeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSiteListSubscriptionReauthorizeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSiteListSubscriptionReauthorize - Invoke action reauthorize. Reauthorize a subscription when you receive a
// reauthorizationRequired challenge.
func (c SiteListSubscriptionClient) CreateSiteListSubscriptionReauthorize(ctx context.Context, id beta.GroupIdSiteIdListIdSubscriptionId, options CreateSiteListSubscriptionReauthorizeOperationOptions) (result CreateSiteListSubscriptionReauthorizeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reauthorize", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	return
}
