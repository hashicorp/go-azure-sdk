package staticsiteuserprovidedfunctionapparmresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *StaticSiteUserProvidedFunctionAppARMResource
}

type StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions struct {
	IsForced *bool
}

func DefaultStaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions() StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions {
	return StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions{}
}

func (o StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IsForced != nil {
		out.Append("isForced", fmt.Sprintf("%v", *o.IsForced))
	}
	return &out
}

// StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuild ...
func (c StaticSiteUserProvidedFunctionAppARMResourcesClient) StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuild(ctx context.Context, id BuildUserProvidedFunctionAppId, input StaticSiteUserProvidedFunctionAppARMResource, options StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions) (result StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          id.ID(),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildThenPoll performs StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuild then polls until it's completed
func (c StaticSiteUserProvidedFunctionAppARMResourcesClient) StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildThenPoll(ctx context.Context, id BuildUserProvidedFunctionAppId, input StaticSiteUserProvidedFunctionAppARMResource, options StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions) error {
	result, err := c.StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuild(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuild: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuild: %+v", err)
	}

	return nil
}
