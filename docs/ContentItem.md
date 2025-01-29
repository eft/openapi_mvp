# ContentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | Folder/File | 
**Name** | **string** | The name of the item. | 
**Readonly** | Pointer to **bool** | Read only if true | [optional] 
**Children** | Pointer to [**[]ContentItem**](ContentItem.md) | Optional files and folders | [optional] 

## Methods

### NewContentItem

`func NewContentItem(contentType string, name string, ) *ContentItem`

NewContentItem instantiates a new ContentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentItemWithDefaults

`func NewContentItemWithDefaults() *ContentItem`

NewContentItemWithDefaults instantiates a new ContentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ContentItem) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ContentItem) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ContentItem) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetName

`func (o *ContentItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentItem) SetName(v string)`

SetName sets Name field to given value.


### GetReadonly

`func (o *ContentItem) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *ContentItem) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *ContentItem) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *ContentItem) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetChildren

`func (o *ContentItem) GetChildren() []ContentItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentItem) GetChildrenOk() (*[]ContentItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentItem) SetChildren(v []ContentItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ContentItem) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


