package userexperienceanalyticsremoteconnection

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

type ListUserExperienceAnalyticsRemoteConnectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsRemoteConnection
}

type ListUserExperienceAnalyticsRemoteConnectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsRemoteConnection
}

type ListUserExperienceAnalyticsRemoteConnectionsOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsRemoteConnectionsOperationOptions() ListUserExperienceAnalyticsRemoteConnectionsOperationOptions {
	return ListUserExperienceAnalyticsRemoteConnectionsOperationOptions{}
}

func (o ListUserExperienceAnalyticsRemoteConnectionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsRemoteConnectionsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsRemoteConnectionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsRemoteConnectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsRemoteConnectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsRemoteConnections - Get userExperienceAnalyticsRemoteConnection from deviceManagement.
// User experience analytics remote connection. The report will be retired on December 31, 2024. You can start using the
// Cloud PC connection quality report now via
// https://learn.microsoft.com/windows-365/enterprise/report-cloud-pc-connection-quality.
func (c UserExperienceAnalyticsRemoteConnectionClient) ListUserExperienceAnalyticsRemoteConnections(ctx context.Context, options ListUserExperienceAnalyticsRemoteConnectionsOperationOptions) (result ListUserExperienceAnalyticsRemoteConnectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsRemoteConnectionsCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsRemoteConnection",
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
		Values *[]beta.UserExperienceAnalyticsRemoteConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsRemoteConnectionsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsRemoteConnectionClient) ListUserExperienceAnalyticsRemoteConnectionsComplete(ctx context.Context, options ListUserExperienceAnalyticsRemoteConnectionsOperationOptions) (ListUserExperienceAnalyticsRemoteConnectionsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsRemoteConnectionsCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsRemoteConnectionOperationPredicate{})
}

// ListUserExperienceAnalyticsRemoteConnectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsRemoteConnectionClient) ListUserExperienceAnalyticsRemoteConnectionsCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsRemoteConnectionsOperationOptions, predicate UserExperienceAnalyticsRemoteConnectionOperationPredicate) (result ListUserExperienceAnalyticsRemoteConnectionsCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsRemoteConnection, 0)

	resp, err := c.ListUserExperienceAnalyticsRemoteConnections(ctx, options)
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

	result = ListUserExperienceAnalyticsRemoteConnectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
