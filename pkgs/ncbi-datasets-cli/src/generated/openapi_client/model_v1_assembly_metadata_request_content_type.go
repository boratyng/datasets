/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
	"fmt"
)

// V1AssemblyMetadataRequestContentType the model 'V1AssemblyMetadataRequestContentType'
type V1AssemblyMetadataRequestContentType string

// List of v1AssemblyMetadataRequestContentType
const (
	V1ASSEMBLYMETADATAREQUESTCONTENTTYPE_COMPLETE V1AssemblyMetadataRequestContentType = "COMPLETE"
	V1ASSEMBLYMETADATAREQUESTCONTENTTYPE_ASSM_ACC V1AssemblyMetadataRequestContentType = "ASSM_ACC"
)

// All allowed values of V1AssemblyMetadataRequestContentType enum
var AllowedV1AssemblyMetadataRequestContentTypeEnumValues = []V1AssemblyMetadataRequestContentType{
	"COMPLETE",
	"ASSM_ACC",
}

func (v *V1AssemblyMetadataRequestContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1AssemblyMetadataRequestContentType(value)
	for _, existing := range AllowedV1AssemblyMetadataRequestContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1AssemblyMetadataRequestContentType", value)
}

// NewV1AssemblyMetadataRequestContentTypeFromValue returns a pointer to a valid V1AssemblyMetadataRequestContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1AssemblyMetadataRequestContentTypeFromValue(v string) (*V1AssemblyMetadataRequestContentType, error) {
	ev := V1AssemblyMetadataRequestContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1AssemblyMetadataRequestContentType: valid values are %v", v, AllowedV1AssemblyMetadataRequestContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1AssemblyMetadataRequestContentType) IsValid() bool {
	for _, existing := range AllowedV1AssemblyMetadataRequestContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1AssemblyMetadataRequestContentType value
func (v V1AssemblyMetadataRequestContentType) Ptr() *V1AssemblyMetadataRequestContentType {
	return &v
}

type NullableV1AssemblyMetadataRequestContentType struct {
	value *V1AssemblyMetadataRequestContentType
	isSet bool
}

func (v NullableV1AssemblyMetadataRequestContentType) Get() *V1AssemblyMetadataRequestContentType {
	return v.value
}

func (v *NullableV1AssemblyMetadataRequestContentType) Set(val *V1AssemblyMetadataRequestContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyMetadataRequestContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyMetadataRequestContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyMetadataRequestContentType(val *V1AssemblyMetadataRequestContentType) *NullableV1AssemblyMetadataRequestContentType {
	return &NullableV1AssemblyMetadataRequestContentType{value: val, isSet: true}
}

func (v NullableV1AssemblyMetadataRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyMetadataRequestContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

