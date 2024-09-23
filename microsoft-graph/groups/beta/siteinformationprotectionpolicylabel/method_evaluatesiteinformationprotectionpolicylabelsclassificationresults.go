package siteinformationprotectionpolicylabel

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

type EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InformationProtectionAction
}

type EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InformationProtectionAction
}

type EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultEvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions() EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions {
	return EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions{}
}

func (o EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions) ToOData() *odata.Query {
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

func (o EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// EvaluateSiteInformationProtectionPolicyLabelsClassificationResults - Invoke action evaluateClassificationResults.
// Using classification results, compute the information protection label that should be applied and return the set of
// actions that must be taken to correctly label the information. This API is useful when a label should be set
// automatically based on classification of the file contents, rather than labeled directly by a user or service. To
// evaluate based on classification results, provide contentInfo, which includes existing content metadata key/value
// pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the
// following
func (c SiteInformationProtectionPolicyLabelClient) EvaluateSiteInformationProtectionPolicyLabelsClassificationResults(ctx context.Context, id beta.GroupIdSiteId, input EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsRequest, options EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions) (result EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsCustomPager{},
		Path:          fmt.Sprintf("%s/informationProtection/policy/labels/evaluateClassificationResults", id.ID()),
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

	temp := make([]beta.InformationProtectionAction, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalInformationProtectionActionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.InformationProtectionAction (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsComplete retrieves all the results into a single object
func (c SiteInformationProtectionPolicyLabelClient) EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsComplete(ctx context.Context, id beta.GroupIdSiteId, input EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsRequest, options EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions) (EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsCompleteResult, error) {
	return c.EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsCompleteMatchingPredicate(ctx, id, input, options, InformationProtectionActionOperationPredicate{})
}

// EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteInformationProtectionPolicyLabelClient) EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteId, input EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsRequest, options EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions, predicate InformationProtectionActionOperationPredicate) (result EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsCompleteResult, err error) {
	items := make([]beta.InformationProtectionAction, 0)

	resp, err := c.EvaluateSiteInformationProtectionPolicyLabelsClassificationResults(ctx, id, input, options)
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

	result = EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
