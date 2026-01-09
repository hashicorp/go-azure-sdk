package sitepagesitepagecanvaslayoutverticalsectionwebpart

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

type ListSitePageCanvasLayoutVerticalSectionWebpartsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WebPart
}

type ListSitePageCanvasLayoutVerticalSectionWebpartsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WebPart
}

type ListSitePageCanvasLayoutVerticalSectionWebpartsOperationOptions struct {
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

func DefaultListSitePageCanvasLayoutVerticalSectionWebpartsOperationOptions() ListSitePageCanvasLayoutVerticalSectionWebpartsOperationOptions {
	return ListSitePageCanvasLayoutVerticalSectionWebpartsOperationOptions{}
}

func (o ListSitePageCanvasLayoutVerticalSectionWebpartsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSitePageCanvasLayoutVerticalSectionWebpartsOperationOptions) ToOData() *odata.Query {
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

func (o ListSitePageCanvasLayoutVerticalSectionWebpartsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSitePageCanvasLayoutVerticalSectionWebpartsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSitePageCanvasLayoutVerticalSectionWebpartsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSitePageCanvasLayoutVerticalSectionWebparts - Get webparts from groups. The set of web parts in this section.
func (c SitePageSitePageCanvasLayoutVerticalSectionWebpartClient) ListSitePageCanvasLayoutVerticalSectionWebparts(ctx context.Context, id beta.GroupIdSiteIdPageId, options ListSitePageCanvasLayoutVerticalSectionWebpartsOperationOptions) (result ListSitePageCanvasLayoutVerticalSectionWebpartsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSitePageCanvasLayoutVerticalSectionWebpartsCustomPager{},
		Path:          fmt.Sprintf("%s/sitePage/canvasLayout/verticalSection/webparts", id.ID()),
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

	temp := make([]beta.WebPart, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalWebPartImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.WebPart (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListSitePageCanvasLayoutVerticalSectionWebpartsComplete retrieves all the results into a single object
func (c SitePageSitePageCanvasLayoutVerticalSectionWebpartClient) ListSitePageCanvasLayoutVerticalSectionWebpartsComplete(ctx context.Context, id beta.GroupIdSiteIdPageId, options ListSitePageCanvasLayoutVerticalSectionWebpartsOperationOptions) (ListSitePageCanvasLayoutVerticalSectionWebpartsCompleteResult, error) {
	return c.ListSitePageCanvasLayoutVerticalSectionWebpartsCompleteMatchingPredicate(ctx, id, options, WebPartOperationPredicate{})
}

// ListSitePageCanvasLayoutVerticalSectionWebpartsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePageSitePageCanvasLayoutVerticalSectionWebpartClient) ListSitePageCanvasLayoutVerticalSectionWebpartsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteIdPageId, options ListSitePageCanvasLayoutVerticalSectionWebpartsOperationOptions, predicate WebPartOperationPredicate) (result ListSitePageCanvasLayoutVerticalSectionWebpartsCompleteResult, err error) {
	items := make([]beta.WebPart, 0)

	resp, err := c.ListSitePageCanvasLayoutVerticalSectionWebparts(ctx, id, options)
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

	result = ListSitePageCanvasLayoutVerticalSectionWebpartsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
