package informationprotectionpolicylabel

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

type EvaluateInformationProtectionPolicyLabelsApplicationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InformationProtectionAction
}

type EvaluateInformationProtectionPolicyLabelsApplicationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InformationProtectionAction
}

type EvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultEvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions() EvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions {
	return EvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions{}
}

func (o EvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions) ToOData() *odata.Query {
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

func (o EvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type EvaluateInformationProtectionPolicyLabelsApplicationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *EvaluateInformationProtectionPolicyLabelsApplicationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// EvaluateInformationProtectionPolicyLabelsApplications - Invoke action evaluateApplication. Compute the information
// protection label that should be applied and return the set of actions that must be taken to correctly label the
// information. This API is useful when a label should be set manually or explicitly by a user or service, rather than
// automatically based on file contents. Given contentInfo, which includes existing content metadata key/value pairs,
// and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of
// the following
func (c InformationProtectionPolicyLabelClient) EvaluateInformationProtectionPolicyLabelsApplications(ctx context.Context, input EvaluateInformationProtectionPolicyLabelsApplicationsRequest, options EvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions) (result EvaluateInformationProtectionPolicyLabelsApplicationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &EvaluateInformationProtectionPolicyLabelsApplicationsCustomPager{},
		Path:          "/me/informationProtection/policy/labels/evaluateApplication",
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

// EvaluateInformationProtectionPolicyLabelsApplicationsComplete retrieves all the results into a single object
func (c InformationProtectionPolicyLabelClient) EvaluateInformationProtectionPolicyLabelsApplicationsComplete(ctx context.Context, input EvaluateInformationProtectionPolicyLabelsApplicationsRequest, options EvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions) (EvaluateInformationProtectionPolicyLabelsApplicationsCompleteResult, error) {
	return c.EvaluateInformationProtectionPolicyLabelsApplicationsCompleteMatchingPredicate(ctx, input, options, InformationProtectionActionOperationPredicate{})
}

// EvaluateInformationProtectionPolicyLabelsApplicationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InformationProtectionPolicyLabelClient) EvaluateInformationProtectionPolicyLabelsApplicationsCompleteMatchingPredicate(ctx context.Context, input EvaluateInformationProtectionPolicyLabelsApplicationsRequest, options EvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions, predicate InformationProtectionActionOperationPredicate) (result EvaluateInformationProtectionPolicyLabelsApplicationsCompleteResult, err error) {
	items := make([]beta.InformationProtectionAction, 0)

	resp, err := c.EvaluateInformationProtectionPolicyLabelsApplications(ctx, input, options)
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

	result = EvaluateInformationProtectionPolicyLabelsApplicationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
