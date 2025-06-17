package sitepagesitepagecanvaslayouthorizontalsection

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

type ListSitePageCanvasLayoutHorizontalSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HorizontalSection
}

type ListSitePageCanvasLayoutHorizontalSectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HorizontalSection
}

type ListSitePageCanvasLayoutHorizontalSectionsOperationOptions struct {
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

func DefaultListSitePageCanvasLayoutHorizontalSectionsOperationOptions() ListSitePageCanvasLayoutHorizontalSectionsOperationOptions {
	return ListSitePageCanvasLayoutHorizontalSectionsOperationOptions{}
}

func (o ListSitePageCanvasLayoutHorizontalSectionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSitePageCanvasLayoutHorizontalSectionsOperationOptions) ToOData() *odata.Query {
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

func (o ListSitePageCanvasLayoutHorizontalSectionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSitePageCanvasLayoutHorizontalSectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSitePageCanvasLayoutHorizontalSectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSitePageCanvasLayoutHorizontalSections - Get horizontalSections from groups. Collection of horizontal sections on
// the SharePoint page.
func (c SitePageSitePageCanvasLayoutHorizontalSectionClient) ListSitePageCanvasLayoutHorizontalSections(ctx context.Context, id beta.GroupIdSiteIdPageId, options ListSitePageCanvasLayoutHorizontalSectionsOperationOptions) (result ListSitePageCanvasLayoutHorizontalSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSitePageCanvasLayoutHorizontalSectionsCustomPager{},
		Path:          fmt.Sprintf("%s/sitePage/canvasLayout/horizontalSections", id.ID()),
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
		Values *[]beta.HorizontalSection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSitePageCanvasLayoutHorizontalSectionsComplete retrieves all the results into a single object
func (c SitePageSitePageCanvasLayoutHorizontalSectionClient) ListSitePageCanvasLayoutHorizontalSectionsComplete(ctx context.Context, id beta.GroupIdSiteIdPageId, options ListSitePageCanvasLayoutHorizontalSectionsOperationOptions) (ListSitePageCanvasLayoutHorizontalSectionsCompleteResult, error) {
	return c.ListSitePageCanvasLayoutHorizontalSectionsCompleteMatchingPredicate(ctx, id, options, HorizontalSectionOperationPredicate{})
}

// ListSitePageCanvasLayoutHorizontalSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePageSitePageCanvasLayoutHorizontalSectionClient) ListSitePageCanvasLayoutHorizontalSectionsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteIdPageId, options ListSitePageCanvasLayoutHorizontalSectionsOperationOptions, predicate HorizontalSectionOperationPredicate) (result ListSitePageCanvasLayoutHorizontalSectionsCompleteResult, err error) {
	items := make([]beta.HorizontalSection, 0)

	resp, err := c.ListSitePageCanvasLayoutHorizontalSections(ctx, id, options)
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

	result = ListSitePageCanvasLayoutHorizontalSectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
