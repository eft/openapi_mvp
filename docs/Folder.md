# Folder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ContentItemResponse**](ContentItemResponse.md) | Optional files and folders | 

## Methods

### NewFolder

`func NewFolder(children []ContentItemResponse, ) *Folder`

NewFolder instantiates a new Folder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderWithDefaults

`func NewFolderWithDefaults() *Folder`

NewFolderWithDefaults instantiates a new Folder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *Folder) GetChildren() []ContentItemResponse`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Folder) GetChildrenOk() (*[]ContentItemResponse, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Folder) SetChildren(v []ContentItemResponse)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


