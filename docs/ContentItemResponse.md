# ContentItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | Folder/File | 
**Name** | **string** | The name of the item. | 
**Readonly** | Pointer to **bool** | Read only if true | [optional] 
**Children** | Pointer to [**[]ContentItemResponse**](ContentItemResponse.md) | Optional files and folders | [optional] 

## Methods

### NewContentItemResponse

`func NewContentItemResponse(contentType string, name string, ) *ContentItemResponse`

NewContentItemResponse instantiates a new ContentItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentItemResponseWithDefaults

`func NewContentItemResponseWithDefaults() *ContentItemResponse`

NewContentItemResponseWithDefaults instantiates a new ContentItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ContentItemResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ContentItemResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ContentItemResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetName

`func (o *ContentItemResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentItemResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentItemResponse) SetName(v string)`

SetName sets Name field to given value.


### GetReadonly

`func (o *ContentItemResponse) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *ContentItemResponse) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *ContentItemResponse) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *ContentItemResponse) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetChildren

`func (o *ContentItemResponse) GetChildren() []ContentItemResponse`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentItemResponse) GetChildrenOk() (*[]ContentItemResponse, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentItemResponse) SetChildren(v []ContentItemResponse)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ContentItemResponse) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


