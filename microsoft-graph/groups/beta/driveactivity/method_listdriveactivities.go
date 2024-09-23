package driveactivity

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

type ListDriveActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemActivityOLD
}

type ListDriveActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemActivityOLD
}

type ListDriveActivitiesOperationOptions struct {
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

func DefaultListDriveActivitiesOperationOptions() ListDriveActivitiesOperationOptions {
	return ListDriveActivitiesOperationOptions{}
}

func (o ListDriveActivitiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveActivitiesOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveActivitiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveActivities - Get activities from groups. The list of recent activities that took place under this drive.
func (c DriveActivityClient) ListDriveActivities(ctx context.Context, id beta.GroupIdDriveId, options ListDriveActivitiesOperationOptions) (result ListDriveActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveActivitiesCustomPager{},
		Path:          fmt.Sprintf("%s/activities", id.ID()),
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
		Values *[]beta.ItemActivityOLD `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveActivitiesComplete retrieves all the results into a single object
func (c DriveActivityClient) ListDriveActivitiesComplete(ctx context.Context, id beta.GroupIdDriveId, options ListDriveActivitiesOperationOptions) (ListDriveActivitiesCompleteResult, error) {
	return c.ListDriveActivitiesCompleteMatchingPredicate(ctx, id, options, ItemActivityOLDOperationPredicate{})
}

// ListDriveActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveActivityClient) ListDriveActivitiesCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdDriveId, options ListDriveActivitiesOperationOptions, predicate ItemActivityOLDOperationPredicate) (result ListDriveActivitiesCompleteResult, err error) {
	items := make([]beta.ItemActivityOLD, 0)

	resp, err := c.ListDriveActivities(ctx, id, options)
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

	result = ListDriveActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
