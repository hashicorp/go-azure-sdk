package drivelistitemversion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDriveListItemVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ListItemVersion
}

type ListDriveListItemVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ListItemVersion
}

type ListDriveListItemVersionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListDriveListItemVersionsOperationOptions() ListDriveListItemVersionsOperationOptions {
	return ListDriveListItemVersionsOperationOptions{}
}

func (o ListDriveListItemVersionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveListItemVersionsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveListItemVersionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveListItemVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveListItemVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveListItemVersions - Get versions from groups. The list of previous versions of the list item.
func (c DriveListItemVersionClient) ListDriveListItemVersions(ctx context.Context, id stable.GroupIdDriveIdListItemId, options ListDriveListItemVersionsOperationOptions) (result ListDriveListItemVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveListItemVersionsCustomPager{},
		Path:          fmt.Sprintf("%s/versions", id.ID()),
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

	temp := make([]stable.ListItemVersion, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalListItemVersionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ListItemVersion (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDriveListItemVersionsComplete retrieves all the results into a single object
func (c DriveListItemVersionClient) ListDriveListItemVersionsComplete(ctx context.Context, id stable.GroupIdDriveIdListItemId, options ListDriveListItemVersionsOperationOptions) (ListDriveListItemVersionsCompleteResult, error) {
	return c.ListDriveListItemVersionsCompleteMatchingPredicate(ctx, id, options, ListItemVersionOperationPredicate{})
}

// ListDriveListItemVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveListItemVersionClient) ListDriveListItemVersionsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdDriveIdListItemId, options ListDriveListItemVersionsOperationOptions, predicate ListItemVersionOperationPredicate) (result ListDriveListItemVersionsCompleteResult, err error) {
	items := make([]stable.ListItemVersion, 0)

	resp, err := c.ListDriveListItemVersions(ctx, id, options)
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

	result = ListDriveListItemVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
