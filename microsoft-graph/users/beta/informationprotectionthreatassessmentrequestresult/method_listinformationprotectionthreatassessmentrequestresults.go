package informationprotectionthreatassessmentrequestresult

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

type ListInformationProtectionThreatAssessmentRequestResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ThreatAssessmentResult
}

type ListInformationProtectionThreatAssessmentRequestResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ThreatAssessmentResult
}

type ListInformationProtectionThreatAssessmentRequestResultsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListInformationProtectionThreatAssessmentRequestResultsOperationOptions() ListInformationProtectionThreatAssessmentRequestResultsOperationOptions {
	return ListInformationProtectionThreatAssessmentRequestResultsOperationOptions{}
}

func (o ListInformationProtectionThreatAssessmentRequestResultsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListInformationProtectionThreatAssessmentRequestResultsOperationOptions) ToOData() *odata.Query {
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

func (o ListInformationProtectionThreatAssessmentRequestResultsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListInformationProtectionThreatAssessmentRequestResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInformationProtectionThreatAssessmentRequestResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInformationProtectionThreatAssessmentRequestResults - Get results from users. A collection of threat assessment
// results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply
// $expand on it.
func (c InformationProtectionThreatAssessmentRequestResultClient) ListInformationProtectionThreatAssessmentRequestResults(ctx context.Context, id beta.UserIdInformationProtectionThreatAssessmentRequestId, options ListInformationProtectionThreatAssessmentRequestResultsOperationOptions) (result ListInformationProtectionThreatAssessmentRequestResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListInformationProtectionThreatAssessmentRequestResultsCustomPager{},
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

// ListInformationProtectionThreatAssessmentRequestResultsComplete retrieves all the results into a single object
func (c InformationProtectionThreatAssessmentRequestResultClient) ListInformationProtectionThreatAssessmentRequestResultsComplete(ctx context.Context, id beta.UserIdInformationProtectionThreatAssessmentRequestId, options ListInformationProtectionThreatAssessmentRequestResultsOperationOptions) (ListInformationProtectionThreatAssessmentRequestResultsCompleteResult, error) {
	return c.ListInformationProtectionThreatAssessmentRequestResultsCompleteMatchingPredicate(ctx, id, options, ThreatAssessmentResultOperationPredicate{})
}

// ListInformationProtectionThreatAssessmentRequestResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InformationProtectionThreatAssessmentRequestResultClient) ListInformationProtectionThreatAssessmentRequestResultsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdInformationProtectionThreatAssessmentRequestId, options ListInformationProtectionThreatAssessmentRequestResultsOperationOptions, predicate ThreatAssessmentResultOperationPredicate) (result ListInformationProtectionThreatAssessmentRequestResultsCompleteResult, err error) {
	items := make([]beta.ThreatAssessmentResult, 0)

	resp, err := c.ListInformationProtectionThreatAssessmentRequestResults(ctx, id, options)
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

	result = ListInformationProtectionThreatAssessmentRequestResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
