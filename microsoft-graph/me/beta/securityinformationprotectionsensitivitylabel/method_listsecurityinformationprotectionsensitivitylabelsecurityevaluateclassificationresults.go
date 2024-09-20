package securityinformationprotectionsensitivitylabel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SecurityInformationProtectionAction
}

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SecurityInformationProtectionAction
}

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions() ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions {
	return ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions{}
}

func (o ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResults - Invoke action
// evaluateClassificationResults. Use the classification results to compute the sensitivity label that should be applied
// and return the set of actions that must be taken to correctly label the information. This API is useful when a label
// should be set automatically based on classification of the file contents, rather than labeled directly by a user or
// service. To evaluate based on classification results, provide the contentInfo, which includes existing content
// metadata key-value pairs, and classification results. The API returns an informationProtectionAction that contains
// one of more of the following
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResults(ctx context.Context, input ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsRequest, options ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions) (result ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsCustomPager{},
		Path:          "/me/security/informationProtection/sensitivityLabels/security.evaluateClassificationResults",
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

	temp := make([]beta.SecurityInformationProtectionAction, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalSecurityInformationProtectionActionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.SecurityInformationProtectionAction (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsComplete retrieves all the results into a single object
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsComplete(ctx context.Context, input ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsRequest, options ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions) (ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsCompleteResult, error) {
	return c.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsCompleteMatchingPredicate(ctx, input, options, SecurityInformationProtectionActionOperationPredicate{})
}

// ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsCompleteMatchingPredicate(ctx context.Context, input ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsRequest, options ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions, predicate SecurityInformationProtectionActionOperationPredicate) (result ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsCompleteResult, err error) {
	items := make([]beta.SecurityInformationProtectionAction, 0)

	resp, err := c.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResults(ctx, input, options)
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

	result = ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
