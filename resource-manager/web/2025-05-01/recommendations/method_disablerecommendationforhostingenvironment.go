package recommendations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DisableRecommendationForHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DisableRecommendationForHostingEnvironmentOperationOptions struct {
	EnvironmentName *string
}

func DefaultDisableRecommendationForHostingEnvironmentOperationOptions() DisableRecommendationForHostingEnvironmentOperationOptions {
	return DisableRecommendationForHostingEnvironmentOperationOptions{}
}

func (o DisableRecommendationForHostingEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DisableRecommendationForHostingEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DisableRecommendationForHostingEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EnvironmentName != nil {
		out.Append("environmentName", fmt.Sprintf("%v", *o.EnvironmentName))
	}
	return &out
}

// DisableRecommendationForHostingEnvironment ...
func (c RecommendationsClient) DisableRecommendationForHostingEnvironment(ctx context.Context, id HostingEnvironmentRecommendationId, options DisableRecommendationForHostingEnvironmentOperationOptions) (result DisableRecommendationForHostingEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/disable", id.ID()),
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
