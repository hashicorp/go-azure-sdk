package siteinformationprotectionthreatassessmentrequestresult

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

type ListSiteInformationProtectionThreatAssessmentRequestResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ThreatAssessmentResult
}

type ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ThreatAssessmentResult
}

type ListSiteInformationProtectionThreatAssessmentRequestResultsOperationOptions struct {
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

func DefaultListSiteInformationProtectionThreatAssessmentRequestResultsOperationOptions() ListSiteInformationProtectionThreatAssessmentRequestResultsOperationOptions {
	return ListSiteInformationProtectionThreatAssessmentRequestResultsOperationOptions{}
}

func (o ListSiteInformationProtectionThreatAssessmentRequestResultsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteInformationProtectionThreatAssessmentRequestResultsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteInformationProtectionThreatAssessmentRequestResultsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteInformationProtectionThreatAssessmentRequestResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteInformationProtectionThreatAssessmentRequestResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteInformationProtectionThreatAssessmentRequestResults - Get results from groups. A collection of threat
// assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless
// you apply $expand on it.
func (c SiteInformationProtectionThreatAssessmentRequestResultClient) ListSiteInformationProtectionThreatAssessmentRequestResults(ctx context.Context, id beta.GroupIdSiteIdInformationProtectionThreatAssessmentRequestId, options ListSiteInformationProtectionThreatAssessmentRequestResultsOperationOptions) (result ListSiteInformationProtectionThreatAssessmentRequestResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteInformationProtectionThreatAssessmentRequestResultsCustomPager{},
		Path:          fmt.Sprintf("%s/results", id.ID()),
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
		Values *[]beta.ThreatAssessmentResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteInformationProtectionThreatAssessmentRequestResultsComplete retrieves all the results into a single object
func (c SiteInformationProtectionThreatAssessmentRequestResultClient) ListSiteInformationProtectionThreatAssessmentRequestResultsComplete(ctx context.Context, id beta.GroupIdSiteIdInformationProtectionThreatAssessmentRequestId, options ListSiteInformationProtectionThreatAssessmentRequestResultsOperationOptions) (ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteResult, error) {
	return c.ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteMatchingPredicate(ctx, id, options, ThreatAssessmentResultOperationPredicate{})
}

// ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteInformationProtectionThreatAssessmentRequestResultClient) ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteIdInformationProtectionThreatAssessmentRequestId, options ListSiteInformationProtectionThreatAssessmentRequestResultsOperationOptions, predicate ThreatAssessmentResultOperationPredicate) (result ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteResult, err error) {
	items := make([]beta.ThreatAssessmentResult, 0)

	resp, err := c.ListSiteInformationProtectionThreatAssessmentRequestResults(ctx, id, options)
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

	result = ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
