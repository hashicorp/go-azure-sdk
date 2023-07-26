package customapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomApisListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CustomApiDefinitionCollection
}

type CustomApisListOperationOptions struct {
	Skiptoken *string
	Top       *int64
}

func DefaultCustomApisListOperationOptions() CustomApisListOperationOptions {
	return CustomApisListOperationOptions{}
}

func (o CustomApisListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CustomApisListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o CustomApisListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Skiptoken != nil {
		out.Append("skiptoken", fmt.Sprintf("%v", *o.Skiptoken))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// CustomApisList ...
func (c CustomAPIsClient) CustomApisList(ctx context.Context, id commonids.SubscriptionId, options CustomApisListOperationOptions) (result CustomApisListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Web/customApis", id.ID()),
		OptionsObject: options,
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
