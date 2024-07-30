package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ColumnDefinition struct {
	Boolean               *BooleanColumn               `json:"boolean,omitempty"`
	Calculated            *CalculatedColumn            `json:"calculated,omitempty"`
	Choice                *ChoiceColumn                `json:"choice,omitempty"`
	ColumnGroup           *string                      `json:"columnGroup,omitempty"`
	ContentApprovalStatus *ContentApprovalStatusColumn `json:"contentApprovalStatus,omitempty"`
	Currency              *CurrencyColumn              `json:"currency,omitempty"`
	DateTime              *DateTimeColumn              `json:"dateTime,omitempty"`
	DefaultValue          *DefaultColumnValue          `json:"defaultValue,omitempty"`
	Description           *string                      `json:"description,omitempty"`
	DisplayName           *string                      `json:"displayName,omitempty"`
	EnforceUniqueValues   *bool                        `json:"enforceUniqueValues,omitempty"`
	Geolocation           *GeolocationColumn           `json:"geolocation,omitempty"`
	Hidden                *bool                        `json:"hidden,omitempty"`
	HyperlinkOrPicture    *HyperlinkOrPictureColumn    `json:"hyperlinkOrPicture,omitempty"`
	Id                    *string                      `json:"id,omitempty"`
	Indexed               *bool                        `json:"indexed,omitempty"`
	IsDeletable           *bool                        `json:"isDeletable,omitempty"`
	IsReorderable         *bool                        `json:"isReorderable,omitempty"`
	IsSealed              *bool                        `json:"isSealed,omitempty"`
	Lookup                *LookupColumn                `json:"lookup,omitempty"`
	Name                  *string                      `json:"name,omitempty"`
	Number                *NumberColumn                `json:"number,omitempty"`
	ODataType             *string                      `json:"@odata.type,omitempty"`
	PersonOrGroup         *PersonOrGroupColumn         `json:"personOrGroup,omitempty"`
	PropagateChanges      *bool                        `json:"propagateChanges,omitempty"`
	ReadOnly              *bool                        `json:"readOnly,omitempty"`
	Required              *bool                        `json:"required,omitempty"`
	SourceColumn          *ColumnDefinition            `json:"sourceColumn,omitempty"`
	SourceContentType     *ContentTypeInfo             `json:"sourceContentType,omitempty"`
	Term                  *TermColumn                  `json:"term,omitempty"`
	Text                  *TextColumn                  `json:"text,omitempty"`
	Thumbnail             *ThumbnailColumn             `json:"thumbnail,omitempty"`
	Type                  *ColumnDefinitionType        `json:"type,omitempty"`
	Validation            *ColumnValidation            `json:"validation,omitempty"`
}
