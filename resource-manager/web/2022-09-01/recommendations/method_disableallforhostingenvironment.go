package recommendations

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

type DisableAllForHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DisableAllForHostingEnvironmentOperationOptions struct {
	EnvironmentName *string
}

func DefaultDisableAllForHostingEnvironmentOperationOptions() DisableAllForHostingEnvironmentOperationOptions {
	return DisableAllForHostingEnvironmentOperationOptions{}
}

func (o DisableAllForHostingEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DisableAllForHostingEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o DisableAllForHostingEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EnvironmentName != nil {
		out.Append("environmentName", fmt.Sprintf("%v", *o.EnvironmentName))
	}
	return &out
}

// DisableAllForHostingEnvironment ...
func (c RecommendationsClient) DisableAllForHostingEnvironment(ctx context.Context, id commonids.AppServiceEnvironmentId, options DisableAllForHostingEnvironmentOperationOptions) (result DisableAllForHostingEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/recommendations/disable", id.ID()),
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
