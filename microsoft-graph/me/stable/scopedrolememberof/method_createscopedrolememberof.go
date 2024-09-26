package scopedrolememberof

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateScopedRoleMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ScopedRoleMembership
}

type CreateScopedRoleMemberOfOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateScopedRoleMemberOfOperationOptions() CreateScopedRoleMemberOfOperationOptions {
	return CreateScopedRoleMemberOfOperationOptions{}
}

func (o CreateScopedRoleMemberOfOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateScopedRoleMemberOfOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateScopedRoleMemberOfOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateScopedRoleMemberOf - Create new navigation property to scopedRoleMemberOf for me
func (c ScopedRoleMemberOfClient) CreateScopedRoleMemberOf(ctx context.Context, input stable.ScopedRoleMembership, options CreateScopedRoleMemberOfOperationOptions) (result CreateScopedRoleMemberOfOperationResponse, err error) {
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
		Path:          "/me/scopedRoleMemberOf",
		RetryFunc:     options.RetryFunc,
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

	var model stable.ScopedRoleMembership
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
