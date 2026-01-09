package resourceaccessprofile

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListResourceAccessProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementResourceAccessProfileBase
}

type ListResourceAccessProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementResourceAccessProfileBase
}

type ListResourceAccessProfilesOperationOptions struct {
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

func DefaultListResourceAccessProfilesOperationOptions() ListResourceAccessProfilesOperationOptions {
	return ListResourceAccessProfilesOperationOptions{}
}

func (o ListResourceAccessProfilesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListResourceAccessProfilesOperationOptions) ToOData() *odata.Query {
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

func (o ListResourceAccessProfilesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListResourceAccessProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListResourceAccessProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListResourceAccessProfiles - Get resourceAccessProfiles from deviceManagement. Collection of resource access settings
// associated with account.
func (c ResourceAccessProfileClient) ListResourceAccessProfiles(ctx context.Context, options ListResourceAccessProfilesOperationOptions) (result ListResourceAccessProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListResourceAccessProfilesCustomPager{},
		Path:          "/deviceManagement/resourceAccessProfiles",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.DeviceManagementResourceAccessProfileBase, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDeviceManagementResourceAccessProfileBaseImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DeviceManagementResourceAccessProfileBase (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListResourceAccessProfilesComplete retrieves all the results into a single object
func (c ResourceAccessProfileClient) ListResourceAccessProfilesComplete(ctx context.Context, options ListResourceAccessProfilesOperationOptions) (ListResourceAccessProfilesCompleteResult, error) {
	return c.ListResourceAccessProfilesCompleteMatchingPredicate(ctx, options, DeviceManagementResourceAccessProfileBaseOperationPredicate{})
}

// ListResourceAccessProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceAccessProfileClient) ListResourceAccessProfilesCompleteMatchingPredicate(ctx context.Context, options ListResourceAccessProfilesOperationOptions, predicate DeviceManagementResourceAccessProfileBaseOperationPredicate) (result ListResourceAccessProfilesCompleteResult, err error) {
	items := make([]beta.DeviceManagementResourceAccessProfileBase, 0)

	resp, err := c.ListResourceAccessProfiles(ctx, options)
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

	result = ListResourceAccessProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
