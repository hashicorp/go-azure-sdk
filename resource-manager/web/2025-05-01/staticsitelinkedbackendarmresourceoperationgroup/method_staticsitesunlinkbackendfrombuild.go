package staticsitelinkedbackendarmresourceoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesUnlinkBackendFromBuildOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type StaticSitesUnlinkBackendFromBuildOperationOptions struct {
	IsCleaningAuthConfig *bool
}

func DefaultStaticSitesUnlinkBackendFromBuildOperationOptions() StaticSitesUnlinkBackendFromBuildOperationOptions {
	return StaticSitesUnlinkBackendFromBuildOperationOptions{}
}

func (o StaticSitesUnlinkBackendFromBuildOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StaticSitesUnlinkBackendFromBuildOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o StaticSitesUnlinkBackendFromBuildOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IsCleaningAuthConfig != nil {
		out.Append("isCleaningAuthConfig", fmt.Sprintf("%v", *o.IsCleaningAuthConfig))
	}
	return &out
}

// StaticSitesUnlinkBackendFromBuild ...
func (c StaticSiteLinkedBackendARMResourceOperationGroupClient) StaticSitesUnlinkBackendFromBuild(ctx context.Context, id BuildLinkedBackendId, options StaticSitesUnlinkBackendFromBuildOperationOptions) (result StaticSitesUnlinkBackendFromBuildOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
