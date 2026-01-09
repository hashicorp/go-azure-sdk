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

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SecurityInformationProtectionAction
}

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SecurityInformationProtectionAction
}

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions() ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions {
	return ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions{}
}

func (o ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions) ToOData() *odata.Query {
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

func (o ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovals - Invoke action evaluateRemoval. Indicate
// to the consuming application what actions it should take to remove the label information. Given contentInfo as an
// input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that
// contains some combination of one or more of the following
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovals(ctx context.Context, id beta.UserId, input ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsRequest, options ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions) (result ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsCustomPager{},
		Path:          fmt.Sprintf("%s/security/informationProtection/sensitivityLabels/security.evaluateRemoval", id.ID()),
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

// ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsComplete retrieves all the results into a single object
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsComplete(ctx context.Context, id beta.UserId, input ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsRequest, options ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions) (ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsCompleteResult, error) {
	return c.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsCompleteMatchingPredicate(ctx, id, input, options, SecurityInformationProtectionActionOperationPredicate{})
}

// ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, input ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsRequest, options ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions, predicate SecurityInformationProtectionActionOperationPredicate) (result ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsCompleteResult, err error) {
	items := make([]beta.SecurityInformationProtectionAction, 0)

	resp, err := c.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovals(ctx, id, input, options)
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

	result = ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
