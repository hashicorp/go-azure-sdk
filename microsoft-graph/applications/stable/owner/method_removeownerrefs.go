package owner

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveOwnerRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveOwnerRefsOperationOptions struct {
	Id      *string
	IfMatch *string
}

func DefaultRemoveOwnerRefsOperationOptions() RemoveOwnerRefsOperationOptions {
	return RemoveOwnerRefsOperationOptions{}
}

func (o RemoveOwnerRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveOwnerRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RemoveOwnerRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Id != nil {
		out.Append("@id", fmt.Sprintf("%v", *o.Id))
	}
	return &out
}

// RemoveOwnerRefs - Remove application owner. Remove an owner from an application. As a recommended best practice, apps
// should have at least two owners.
func (c OwnerClient) RemoveOwnerRefs(ctx context.Context, id stable.ApplicationId, options RemoveOwnerRefsOperationOptions) (result RemoveOwnerRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/owners/$ref", id.ID()),
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
