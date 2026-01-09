package remotedesktopsecurityconfigurationapprovedclientapp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ApprovedClientApp
}

type ListRemoteDesktopSecurityConfigurationApprovedClientAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ApprovedClientApp
}

type ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationOptions() ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationOptions {
	return ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationOptions{}
}

func (o ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRemoteDesktopSecurityConfigurationApprovedClientAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRemoteDesktopSecurityConfigurationApprovedClientAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRemoteDesktopSecurityConfigurationApprovedClientApps - Get approvedClientApps from servicePrincipals
func (c RemoteDesktopSecurityConfigurationApprovedClientAppClient) ListRemoteDesktopSecurityConfigurationApprovedClientApps(ctx context.Context, id beta.ServicePrincipalId, options ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationOptions) (result ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRemoteDesktopSecurityConfigurationApprovedClientAppsCustomPager{},
		Path:          fmt.Sprintf("%s/remoteDesktopSecurityConfiguration/approvedClientApps", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]beta.ApprovedClientApp `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRemoteDesktopSecurityConfigurationApprovedClientAppsComplete retrieves all the results into a single object
func (c RemoteDesktopSecurityConfigurationApprovedClientAppClient) ListRemoteDesktopSecurityConfigurationApprovedClientAppsComplete(ctx context.Context, id beta.ServicePrincipalId, options ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationOptions) (ListRemoteDesktopSecurityConfigurationApprovedClientAppsCompleteResult, error) {
	return c.ListRemoteDesktopSecurityConfigurationApprovedClientAppsCompleteMatchingPredicate(ctx, id, options, ApprovedClientAppOperationPredicate{})
}

// ListRemoteDesktopSecurityConfigurationApprovedClientAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RemoteDesktopSecurityConfigurationApprovedClientAppClient) ListRemoteDesktopSecurityConfigurationApprovedClientAppsCompleteMatchingPredicate(ctx context.Context, id beta.ServicePrincipalId, options ListRemoteDesktopSecurityConfigurationApprovedClientAppsOperationOptions, predicate ApprovedClientAppOperationPredicate) (result ListRemoteDesktopSecurityConfigurationApprovedClientAppsCompleteResult, err error) {
	items := make([]beta.ApprovedClientApp, 0)

	resp, err := c.ListRemoteDesktopSecurityConfigurationApprovedClientApps(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListRemoteDesktopSecurityConfigurationApprovedClientAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
