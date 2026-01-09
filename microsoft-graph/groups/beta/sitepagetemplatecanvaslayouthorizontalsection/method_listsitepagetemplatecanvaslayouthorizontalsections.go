package sitepagetemplatecanvaslayouthorizontalsection

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

type ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HorizontalSection
}

type ListSitePageTemplateCanvasLayoutHorizontalSectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HorizontalSection
}

type ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationOptions struct {
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

func DefaultListSitePageTemplateCanvasLayoutHorizontalSectionsOperationOptions() ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationOptions {
	return ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationOptions{}
}

func (o ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationOptions) ToOData() *odata.Query {
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

func (o ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSitePageTemplateCanvasLayoutHorizontalSectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSitePageTemplateCanvasLayoutHorizontalSectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSitePageTemplateCanvasLayoutHorizontalSections - Get horizontalSections from groups. Collection of horizontal
// sections on the SharePoint page.
func (c SitePageTemplateCanvasLayoutHorizontalSectionClient) ListSitePageTemplateCanvasLayoutHorizontalSections(ctx context.Context, id beta.GroupIdSiteIdPageTemplateId, options ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationOptions) (result ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSitePageTemplateCanvasLayoutHorizontalSectionsCustomPager{},
		Path:          fmt.Sprintf("%s/canvasLayout/horizontalSections", id.ID()),
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

// ListSitePageTemplateCanvasLayoutHorizontalSectionsComplete retrieves all the results into a single object
func (c SitePageTemplateCanvasLayoutHorizontalSectionClient) ListSitePageTemplateCanvasLayoutHorizontalSectionsComplete(ctx context.Context, id beta.GroupIdSiteIdPageTemplateId, options ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationOptions) (ListSitePageTemplateCanvasLayoutHorizontalSectionsCompleteResult, error) {
	return c.ListSitePageTemplateCanvasLayoutHorizontalSectionsCompleteMatchingPredicate(ctx, id, options, HorizontalSectionOperationPredicate{})
}

// ListSitePageTemplateCanvasLayoutHorizontalSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePageTemplateCanvasLayoutHorizontalSectionClient) ListSitePageTemplateCanvasLayoutHorizontalSectionsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteIdPageTemplateId, options ListSitePageTemplateCanvasLayoutHorizontalSectionsOperationOptions, predicate HorizontalSectionOperationPredicate) (result ListSitePageTemplateCanvasLayoutHorizontalSectionsCompleteResult, err error) {
	items := make([]beta.HorizontalSection, 0)

	resp, err := c.ListSitePageTemplateCanvasLayoutHorizontalSections(ctx, id, options)
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

	result = ListSitePageTemplateCanvasLayoutHorizontalSectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
