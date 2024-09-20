package drivespecial

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDriveSpecialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DriveItem
}

type ListDriveSpecialsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DriveItem
}

type ListDriveSpecialsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListDriveSpecialsOperationOptions() ListDriveSpecialsOperationOptions {
	return ListDriveSpecialsOperationOptions{}
}

func (o ListDriveSpecialsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveSpecialsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveSpecialsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveSpecialsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveSpecialsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveSpecials - Get special from users. Collection of common folders available in OneDrive. Read-only. Nullable.
func (c DriveSpecialClient) ListDriveSpecials(ctx context.Context, id stable.UserIdDriveId, options ListDriveSpecialsOperationOptions) (result ListDriveSpecialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveSpecialsCustomPager{},
		Path:          fmt.Sprintf("%s/special", id.ID()),
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
		Values *[]stable.DriveItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveSpecialsComplete retrieves all the results into a single object
func (c DriveSpecialClient) ListDriveSpecialsComplete(ctx context.Context, id stable.UserIdDriveId, options ListDriveSpecialsOperationOptions) (ListDriveSpecialsCompleteResult, error) {
	return c.ListDriveSpecialsCompleteMatchingPredicate(ctx, id, options, DriveItemOperationPredicate{})
}

// ListDriveSpecialsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveSpecialClient) ListDriveSpecialsCompleteMatchingPredicate(ctx context.Context, id stable.UserIdDriveId, options ListDriveSpecialsOperationOptions, predicate DriveItemOperationPredicate) (result ListDriveSpecialsCompleteResult, err error) {
	items := make([]stable.DriveItem, 0)

	resp, err := c.ListDriveSpecials(ctx, id, options)
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

	result = ListDriveSpecialsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
