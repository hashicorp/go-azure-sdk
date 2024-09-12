package securityinformationprotectionsensitivitylabel

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

type ListSecurityInformationProtectionSensitivityLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SecuritySensitivityLabel
}

type ListSecurityInformationProtectionSensitivityLabelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SecuritySensitivityLabel
}

type ListSecurityInformationProtectionSensitivityLabelsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListSecurityInformationProtectionSensitivityLabelsOperationOptions() ListSecurityInformationProtectionSensitivityLabelsOperationOptions {
	return ListSecurityInformationProtectionSensitivityLabelsOperationOptions{}
}

func (o ListSecurityInformationProtectionSensitivityLabelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSecurityInformationProtectionSensitivityLabelsOperationOptions) ToOData() *odata.Query {
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

func (o ListSecurityInformationProtectionSensitivityLabelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSecurityInformationProtectionSensitivityLabelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSecurityInformationProtectionSensitivityLabelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSecurityInformationProtectionSensitivityLabels - Get sensitivityLabels from me. Read the Microsoft Purview
// Information Protection labels for the user or organization.
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabels(ctx context.Context, options ListSecurityInformationProtectionSensitivityLabelsOperationOptions) (result ListSecurityInformationProtectionSensitivityLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSecurityInformationProtectionSensitivityLabelsCustomPager{},
		Path:          "/me/security/informationProtection/sensitivityLabels",
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
		Values *[]beta.SecuritySensitivityLabel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSecurityInformationProtectionSensitivityLabelsComplete retrieves all the results into a single object
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelsComplete(ctx context.Context, options ListSecurityInformationProtectionSensitivityLabelsOperationOptions) (ListSecurityInformationProtectionSensitivityLabelsCompleteResult, error) {
	return c.ListSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx, options, SecuritySensitivityLabelOperationPredicate{})
}

// ListSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx context.Context, options ListSecurityInformationProtectionSensitivityLabelsOperationOptions, predicate SecuritySensitivityLabelOperationPredicate) (result ListSecurityInformationProtectionSensitivityLabelsCompleteResult, err error) {
	items := make([]beta.SecuritySensitivityLabel, 0)

	resp, err := c.ListSecurityInformationProtectionSensitivityLabels(ctx, options)
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

	result = ListSecurityInformationProtectionSensitivityLabelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
