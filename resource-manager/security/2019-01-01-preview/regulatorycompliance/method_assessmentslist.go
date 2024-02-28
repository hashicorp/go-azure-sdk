package regulatorycompliance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RegulatoryComplianceAssessment
}

type AssessmentsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RegulatoryComplianceAssessment
}

type AssessmentsListOperationOptions struct {
	Filter *string
}

func DefaultAssessmentsListOperationOptions() AssessmentsListOperationOptions {
	return AssessmentsListOperationOptions{}
}

func (o AssessmentsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssessmentsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o AssessmentsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// AssessmentsList ...
func (c RegulatoryComplianceClient) AssessmentsList(ctx context.Context, id RegulatoryComplianceControlId, options AssessmentsListOperationOptions) (result AssessmentsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/regulatoryComplianceAssessments", id.ID()),
		OptionsObject: options,
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
		Values *[]RegulatoryComplianceAssessment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssessmentsListComplete retrieves all the results into a single object
func (c RegulatoryComplianceClient) AssessmentsListComplete(ctx context.Context, id RegulatoryComplianceControlId, options AssessmentsListOperationOptions) (AssessmentsListCompleteResult, error) {
	return c.AssessmentsListCompleteMatchingPredicate(ctx, id, options, RegulatoryComplianceAssessmentOperationPredicate{})
}

// AssessmentsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RegulatoryComplianceClient) AssessmentsListCompleteMatchingPredicate(ctx context.Context, id RegulatoryComplianceControlId, options AssessmentsListOperationOptions, predicate RegulatoryComplianceAssessmentOperationPredicate) (result AssessmentsListCompleteResult, err error) {
	items := make([]RegulatoryComplianceAssessment, 0)

	resp, err := c.AssessmentsList(ctx, id, options)
	if err != nil {
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

	result = AssessmentsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
