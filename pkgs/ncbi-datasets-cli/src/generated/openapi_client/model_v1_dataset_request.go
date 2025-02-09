/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1DatasetRequest struct for V1DatasetRequest
type V1DatasetRequest struct {
	GenomeV1 *V1AssemblyDatasetRequest `json:"genome_v1,omitempty"`
	GeneV1 *V1GeneDatasetRequest `json:"gene_v1,omitempty"`
	VirusV1 *V1VirusDatasetRequest `json:"virus_v1,omitempty"`
}

// NewV1DatasetRequest instantiates a new V1DatasetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DatasetRequest() *V1DatasetRequest {
	this := V1DatasetRequest{}
	return &this
}

// NewV1DatasetRequestWithDefaults instantiates a new V1DatasetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DatasetRequestWithDefaults() *V1DatasetRequest {
	this := V1DatasetRequest{}
	return &this
}

// GetGenomeV1 returns the GenomeV1 field value if set, zero value otherwise.
func (o *V1DatasetRequest) GetGenomeV1() V1AssemblyDatasetRequest {
	if o == nil || o.GenomeV1 == nil {
		var ret V1AssemblyDatasetRequest
		return ret
	}
	return *o.GenomeV1
}

// GetGenomeV1Ok returns a tuple with the GenomeV1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DatasetRequest) GetGenomeV1Ok() (*V1AssemblyDatasetRequest, bool) {
	if o == nil || o.GenomeV1 == nil {
		return nil, false
	}
	return o.GenomeV1, true
}

// HasGenomeV1 returns a boolean if a field has been set.
func (o *V1DatasetRequest) HasGenomeV1() bool {
	if o != nil && o.GenomeV1 != nil {
		return true
	}

	return false
}

// SetGenomeV1 gets a reference to the given V1AssemblyDatasetRequest and assigns it to the GenomeV1 field.
func (o *V1DatasetRequest) SetGenomeV1(v V1AssemblyDatasetRequest) {
	o.GenomeV1 = &v
}

// GetGeneV1 returns the GeneV1 field value if set, zero value otherwise.
func (o *V1DatasetRequest) GetGeneV1() V1GeneDatasetRequest {
	if o == nil || o.GeneV1 == nil {
		var ret V1GeneDatasetRequest
		return ret
	}
	return *o.GeneV1
}

// GetGeneV1Ok returns a tuple with the GeneV1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DatasetRequest) GetGeneV1Ok() (*V1GeneDatasetRequest, bool) {
	if o == nil || o.GeneV1 == nil {
		return nil, false
	}
	return o.GeneV1, true
}

// HasGeneV1 returns a boolean if a field has been set.
func (o *V1DatasetRequest) HasGeneV1() bool {
	if o != nil && o.GeneV1 != nil {
		return true
	}

	return false
}

// SetGeneV1 gets a reference to the given V1GeneDatasetRequest and assigns it to the GeneV1 field.
func (o *V1DatasetRequest) SetGeneV1(v V1GeneDatasetRequest) {
	o.GeneV1 = &v
}

// GetVirusV1 returns the VirusV1 field value if set, zero value otherwise.
func (o *V1DatasetRequest) GetVirusV1() V1VirusDatasetRequest {
	if o == nil || o.VirusV1 == nil {
		var ret V1VirusDatasetRequest
		return ret
	}
	return *o.VirusV1
}

// GetVirusV1Ok returns a tuple with the VirusV1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DatasetRequest) GetVirusV1Ok() (*V1VirusDatasetRequest, bool) {
	if o == nil || o.VirusV1 == nil {
		return nil, false
	}
	return o.VirusV1, true
}

// HasVirusV1 returns a boolean if a field has been set.
func (o *V1DatasetRequest) HasVirusV1() bool {
	if o != nil && o.VirusV1 != nil {
		return true
	}

	return false
}

// SetVirusV1 gets a reference to the given V1VirusDatasetRequest and assigns it to the VirusV1 field.
func (o *V1DatasetRequest) SetVirusV1(v V1VirusDatasetRequest) {
	o.VirusV1 = &v
}

func (o V1DatasetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GenomeV1 != nil  {
		toSerialize["genome_v1"] = o.GenomeV1
	}
	if o.GeneV1 != nil  {
		toSerialize["gene_v1"] = o.GeneV1
	}
	if o.VirusV1 != nil  {
		toSerialize["virus_v1"] = o.VirusV1
	}
	return json.Marshal(toSerialize)
}

type NullableV1DatasetRequest struct {
	value *V1DatasetRequest
	isSet bool
}

func (v NullableV1DatasetRequest) Get() *V1DatasetRequest {
	return v.value
}

func (v *NullableV1DatasetRequest) Set(val *V1DatasetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DatasetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DatasetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DatasetRequest(val *V1DatasetRequest) *NullableV1DatasetRequest {
	return &NullableV1DatasetRequest{value: val, isSet: true}
}

func (v NullableV1DatasetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DatasetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


