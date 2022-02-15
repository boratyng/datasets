/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1reportsVirusPeptide struct for V1reportsVirusPeptide
type V1reportsVirusPeptide struct {
	Accession *string `json:"accession,omitempty"`
	Name *string `json:"name,omitempty"`
	OtherNames *[]string `json:"other_names,omitempty"`
	Nucleotide *V1reportsSeqRangeSetFasta `json:"nucleotide,omitempty"`
	Protein *V1reportsSeqRangeSetFasta `json:"protein,omitempty"`
	PdbIds *[]string `json:"pdb_ids,omitempty"`
	Cdd *[]V1reportsConservedDomain `json:"cdd,omitempty"`
	UniProtKb *V1reportsVirusPeptideUniProtId `json:"uni_prot_kb,omitempty"`
	MaturePeptide *[]V1reportsVirusPeptide `json:"mature_peptide,omitempty"`
	ProteinCompleteness *V1reportsVirusAssemblyCompleteness `json:"protein_completeness,omitempty"`
}

// NewV1reportsVirusPeptide instantiates a new V1reportsVirusPeptide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsVirusPeptide() *V1reportsVirusPeptide {
	this := V1reportsVirusPeptide{}
	var proteinCompleteness V1reportsVirusAssemblyCompleteness = V1REPORTSVIRUSASSEMBLYCOMPLETENESS_UNKNOWN
	this.ProteinCompleteness = &proteinCompleteness
	return &this
}

// NewV1reportsVirusPeptideWithDefaults instantiates a new V1reportsVirusPeptide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsVirusPeptideWithDefaults() *V1reportsVirusPeptide {
	this := V1reportsVirusPeptide{}
	var proteinCompleteness V1reportsVirusAssemblyCompleteness = V1REPORTSVIRUSASSEMBLYCOMPLETENESS_UNKNOWN
	this.ProteinCompleteness = &proteinCompleteness
	return &this
}

// GetAccession returns the Accession field value if set, zero value otherwise.
func (o *V1reportsVirusPeptide) GetAccession() string {
	if o == nil || o.Accession == nil {
		var ret string
		return ret
	}
	return *o.Accession
}

// GetAccessionOk returns a tuple with the Accession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusPeptide) GetAccessionOk() (*string, bool) {
	if o == nil || o.Accession == nil {
		return nil, false
	}
	return o.Accession, true
}

// HasAccession returns a boolean if a field has been set.
func (o *V1reportsVirusPeptide) HasAccession() bool {
	if o != nil && o.Accession != nil {
		return true
	}

	return false
}

// SetAccession gets a reference to the given string and assigns it to the Accession field.
func (o *V1reportsVirusPeptide) SetAccession(v string) {
	o.Accession = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1reportsVirusPeptide) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusPeptide) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1reportsVirusPeptide) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1reportsVirusPeptide) SetName(v string) {
	o.Name = &v
}

// GetOtherNames returns the OtherNames field value if set, zero value otherwise.
func (o *V1reportsVirusPeptide) GetOtherNames() []string {
	if o == nil || o.OtherNames == nil {
		var ret []string
		return ret
	}
	return *o.OtherNames
}

// GetOtherNamesOk returns a tuple with the OtherNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusPeptide) GetOtherNamesOk() (*[]string, bool) {
	if o == nil || o.OtherNames == nil {
		return nil, false
	}
	return o.OtherNames, true
}

// HasOtherNames returns a boolean if a field has been set.
func (o *V1reportsVirusPeptide) HasOtherNames() bool {
	if o != nil && o.OtherNames != nil {
		return true
	}

	return false
}

// SetOtherNames gets a reference to the given []string and assigns it to the OtherNames field.
func (o *V1reportsVirusPeptide) SetOtherNames(v []string) {
	o.OtherNames = &v
}

// GetNucleotide returns the Nucleotide field value if set, zero value otherwise.
func (o *V1reportsVirusPeptide) GetNucleotide() V1reportsSeqRangeSetFasta {
	if o == nil || o.Nucleotide == nil {
		var ret V1reportsSeqRangeSetFasta
		return ret
	}
	return *o.Nucleotide
}

// GetNucleotideOk returns a tuple with the Nucleotide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusPeptide) GetNucleotideOk() (*V1reportsSeqRangeSetFasta, bool) {
	if o == nil || o.Nucleotide == nil {
		return nil, false
	}
	return o.Nucleotide, true
}

// HasNucleotide returns a boolean if a field has been set.
func (o *V1reportsVirusPeptide) HasNucleotide() bool {
	if o != nil && o.Nucleotide != nil {
		return true
	}

	return false
}

// SetNucleotide gets a reference to the given V1reportsSeqRangeSetFasta and assigns it to the Nucleotide field.
func (o *V1reportsVirusPeptide) SetNucleotide(v V1reportsSeqRangeSetFasta) {
	o.Nucleotide = &v
}

// GetProtein returns the Protein field value if set, zero value otherwise.
func (o *V1reportsVirusPeptide) GetProtein() V1reportsSeqRangeSetFasta {
	if o == nil || o.Protein == nil {
		var ret V1reportsSeqRangeSetFasta
		return ret
	}
	return *o.Protein
}

// GetProteinOk returns a tuple with the Protein field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusPeptide) GetProteinOk() (*V1reportsSeqRangeSetFasta, bool) {
	if o == nil || o.Protein == nil {
		return nil, false
	}
	return o.Protein, true
}

// HasProtein returns a boolean if a field has been set.
func (o *V1reportsVirusPeptide) HasProtein() bool {
	if o != nil && o.Protein != nil {
		return true
	}

	return false
}

// SetProtein gets a reference to the given V1reportsSeqRangeSetFasta and assigns it to the Protein field.
func (o *V1reportsVirusPeptide) SetProtein(v V1reportsSeqRangeSetFasta) {
	o.Protein = &v
}

// GetPdbIds returns the PdbIds field value if set, zero value otherwise.
func (o *V1reportsVirusPeptide) GetPdbIds() []string {
	if o == nil || o.PdbIds == nil {
		var ret []string
		return ret
	}
	return *o.PdbIds
}

// GetPdbIdsOk returns a tuple with the PdbIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusPeptide) GetPdbIdsOk() (*[]string, bool) {
	if o == nil || o.PdbIds == nil {
		return nil, false
	}
	return o.PdbIds, true
}

// HasPdbIds returns a boolean if a field has been set.
func (o *V1reportsVirusPeptide) HasPdbIds() bool {
	if o != nil && o.PdbIds != nil {
		return true
	}

	return false
}

// SetPdbIds gets a reference to the given []string and assigns it to the PdbIds field.
func (o *V1reportsVirusPeptide) SetPdbIds(v []string) {
	o.PdbIds = &v
}

// GetCdd returns the Cdd field value if set, zero value otherwise.
func (o *V1reportsVirusPeptide) GetCdd() []V1reportsConservedDomain {
	if o == nil || o.Cdd == nil {
		var ret []V1reportsConservedDomain
		return ret
	}
	return *o.Cdd
}

// GetCddOk returns a tuple with the Cdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusPeptide) GetCddOk() (*[]V1reportsConservedDomain, bool) {
	if o == nil || o.Cdd == nil {
		return nil, false
	}
	return o.Cdd, true
}

// HasCdd returns a boolean if a field has been set.
func (o *V1reportsVirusPeptide) HasCdd() bool {
	if o != nil && o.Cdd != nil {
		return true
	}

	return false
}

// SetCdd gets a reference to the given []V1reportsConservedDomain and assigns it to the Cdd field.
func (o *V1reportsVirusPeptide) SetCdd(v []V1reportsConservedDomain) {
	o.Cdd = &v
}

// GetUniProtKb returns the UniProtKb field value if set, zero value otherwise.
func (o *V1reportsVirusPeptide) GetUniProtKb() V1reportsVirusPeptideUniProtId {
	if o == nil || o.UniProtKb == nil {
		var ret V1reportsVirusPeptideUniProtId
		return ret
	}
	return *o.UniProtKb
}

// GetUniProtKbOk returns a tuple with the UniProtKb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusPeptide) GetUniProtKbOk() (*V1reportsVirusPeptideUniProtId, bool) {
	if o == nil || o.UniProtKb == nil {
		return nil, false
	}
	return o.UniProtKb, true
}

// HasUniProtKb returns a boolean if a field has been set.
func (o *V1reportsVirusPeptide) HasUniProtKb() bool {
	if o != nil && o.UniProtKb != nil {
		return true
	}

	return false
}

// SetUniProtKb gets a reference to the given V1reportsVirusPeptideUniProtId and assigns it to the UniProtKb field.
func (o *V1reportsVirusPeptide) SetUniProtKb(v V1reportsVirusPeptideUniProtId) {
	o.UniProtKb = &v
}

// GetMaturePeptide returns the MaturePeptide field value if set, zero value otherwise.
func (o *V1reportsVirusPeptide) GetMaturePeptide() []V1reportsVirusPeptide {
	if o == nil || o.MaturePeptide == nil {
		var ret []V1reportsVirusPeptide
		return ret
	}
	return *o.MaturePeptide
}

// GetMaturePeptideOk returns a tuple with the MaturePeptide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusPeptide) GetMaturePeptideOk() (*[]V1reportsVirusPeptide, bool) {
	if o == nil || o.MaturePeptide == nil {
		return nil, false
	}
	return o.MaturePeptide, true
}

// HasMaturePeptide returns a boolean if a field has been set.
func (o *V1reportsVirusPeptide) HasMaturePeptide() bool {
	if o != nil && o.MaturePeptide != nil {
		return true
	}

	return false
}

// SetMaturePeptide gets a reference to the given []V1reportsVirusPeptide and assigns it to the MaturePeptide field.
func (o *V1reportsVirusPeptide) SetMaturePeptide(v []V1reportsVirusPeptide) {
	o.MaturePeptide = &v
}

// GetProteinCompleteness returns the ProteinCompleteness field value if set, zero value otherwise.
func (o *V1reportsVirusPeptide) GetProteinCompleteness() V1reportsVirusAssemblyCompleteness {
	if o == nil || o.ProteinCompleteness == nil {
		var ret V1reportsVirusAssemblyCompleteness
		return ret
	}
	return *o.ProteinCompleteness
}

// GetProteinCompletenessOk returns a tuple with the ProteinCompleteness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusPeptide) GetProteinCompletenessOk() (*V1reportsVirusAssemblyCompleteness, bool) {
	if o == nil || o.ProteinCompleteness == nil {
		return nil, false
	}
	return o.ProteinCompleteness, true
}

// HasProteinCompleteness returns a boolean if a field has been set.
func (o *V1reportsVirusPeptide) HasProteinCompleteness() bool {
	if o != nil && o.ProteinCompleteness != nil {
		return true
	}

	return false
}

// SetProteinCompleteness gets a reference to the given V1reportsVirusAssemblyCompleteness and assigns it to the ProteinCompleteness field.
func (o *V1reportsVirusPeptide) SetProteinCompleteness(v V1reportsVirusAssemblyCompleteness) {
	o.ProteinCompleteness = &v
}

func (o V1reportsVirusPeptide) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accession != nil  {
		toSerialize["accession"] = o.Accession
	}
	if o.Name != nil  {
		toSerialize["name"] = o.Name
	}
	if o.OtherNames != nil && len(o.GetOtherNames()) > 0  {
		toSerialize["other_names"] = o.OtherNames
	}
	if o.Nucleotide != nil  {
		toSerialize["nucleotide"] = o.Nucleotide
	}
	if o.Protein != nil  {
		toSerialize["protein"] = o.Protein
	}
	if o.PdbIds != nil && len(o.GetPdbIds()) > 0  {
		toSerialize["pdb_ids"] = o.PdbIds
	}
	if o.Cdd != nil && len(o.GetCdd()) > 0  {
		toSerialize["cdd"] = o.Cdd
	}
	if o.UniProtKb != nil  {
		toSerialize["uni_prot_kb"] = o.UniProtKb
	}
	if o.MaturePeptide != nil && len(o.GetMaturePeptide()) > 0  {
		toSerialize["mature_peptide"] = o.MaturePeptide
	}
	if o.ProteinCompleteness != nil  {
		toSerialize["protein_completeness"] = o.ProteinCompleteness
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsVirusPeptide struct {
	value *V1reportsVirusPeptide
	isSet bool
}

func (v NullableV1reportsVirusPeptide) Get() *V1reportsVirusPeptide {
	return v.value
}

func (v *NullableV1reportsVirusPeptide) Set(val *V1reportsVirusPeptide) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsVirusPeptide) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsVirusPeptide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsVirusPeptide(val *V1reportsVirusPeptide) *NullableV1reportsVirusPeptide {
	return &NullableV1reportsVirusPeptide{value: val, isSet: true}
}

func (v NullableV1reportsVirusPeptide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsVirusPeptide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


