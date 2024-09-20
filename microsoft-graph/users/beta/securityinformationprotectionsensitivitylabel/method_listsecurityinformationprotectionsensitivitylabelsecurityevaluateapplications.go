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

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SecurityInformationProtectionAction
}

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SecurityInformationProtectionAction
}

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions() ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions {
	return ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions{}
}

func (o ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions) ToOData() *odata.Query {
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

func (o ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplications - Invoke action evaluateApplication.
// Compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly
// label the information. This API is useful when a label should be set manually or explicitly by a user or service,
// rather than automatically based on file contents. Given contentInfo, which includes existing content metadata
// key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains
// one of more of the following
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplications(ctx context.Context, id beta.UserId, input ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsRequest, options ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions) (result ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsCustomPager{},
		Path:          fmt.Sprintf("%s/security/informationProtection/sensitivityLabels/security.evaluateApplication", id.ID()),
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

// ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsComplete retrieves all the results into a single object
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsComplete(ctx context.Context, id beta.UserId, input ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsRequest, options ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions) (ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsCompleteResult, error) {
	return c.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsCompleteMatchingPredicate(ctx, id, input, options, SecurityInformationProtectionActionOperationPredicate{})
}

// ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, input ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsRequest, options ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions, predicate SecurityInformationProtectionActionOperationPredicate) (result ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsCompleteResult, err error) {
	items := make([]beta.SecurityInformationProtectionAction, 0)

	resp, err := c.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplications(ctx, id, input, options)
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

	result = ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
