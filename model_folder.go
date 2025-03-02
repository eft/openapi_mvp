/*
Minimum Spec

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mvp

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Folder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Folder{}

// Folder struct for Folder
type Folder struct {
	// The name of the item.
	Name string `json:"name"`
	// Read only if true
	Readonly *bool `json:"readonly,omitempty"`
	// Folder/File
	ContentType string `json:"contentType"`
	// Optional files and folders
	Children []ContentItemResponse `json:"children,omitempty"`
	// contrived property for testing purposes
	Color string `json:"color"`
}

type _Folder Folder

// NewFolder instantiates a new Folder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolder(name string, contentType string, color string) *Folder {
	this := Folder{}
	this.Name = name
	this.ContentType = contentType
	this.Color = color
	return &this
}

// NewFolderWithDefaults instantiates a new Folder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderWithDefaults() *Folder {
	this := Folder{}
	return &this
}

// GetName returns the Name field value
func (o *Folder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Folder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Folder) SetName(v string) {
	o.Name = v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *Folder) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *Folder) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *Folder) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetContentType returns the ContentType field value
func (o *Folder) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *Folder) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *Folder) SetContentType(v string) {
	o.ContentType = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *Folder) GetChildren() []ContentItemResponse {
	if o == nil || IsNil(o.Children) {
		var ret []ContentItemResponse
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetChildrenOk() ([]ContentItemResponse, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Folder) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []ContentItemResponse and assigns it to the Children field.
func (o *Folder) SetChildren(v []ContentItemResponse) {
	o.Children = v
}

// GetColor returns the Color field value
func (o *Folder) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *Folder) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *Folder) SetColor(v string) {
	o.Color = v
}

func (o Folder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Folder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
	}
	toSerialize["contentType"] = o.ContentType
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	toSerialize["color"] = o.Color
	return toSerialize, nil
}

func (o *Folder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"contentType",
		"color",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFolder := _Folder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFolder)

	if err != nil {
		return err
	}

	*o = Folder(varFolder)

	return err
}

type NullableFolder struct {
	value *Folder
	isSet bool
}

func (v NullableFolder) Get() *Folder {
	return v.value
}

func (v *NullableFolder) Set(val *Folder) {
	v.value = val
	v.isSet = true
}

func (v NullableFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolder(val *Folder) *NullableFolder {
	return &NullableFolder{value: val, isSet: true}
}

func (v NullableFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


