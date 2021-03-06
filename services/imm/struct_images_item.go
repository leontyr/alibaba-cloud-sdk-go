package imm

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ImagesItem is a nested struct in imm response
type ImagesItem struct {
	FacesStatus                  string                   `json:"FacesStatus" xml:"FacesStatus"`
	CroppingSuggestionFailReason string                   `json:"CroppingSuggestionFailReason" xml:"CroppingSuggestionFailReason"`
	CelebrityFailReason          string                   `json:"CelebrityFailReason" xml:"CelebrityFailReason"`
	FacesFailReason              string                   `json:"FacesFailReason" xml:"FacesFailReason"`
	AddressStatus                string                   `json:"AddressStatus" xml:"AddressStatus"`
	Exif                         string                   `json:"Exif" xml:"Exif"`
	CroppingSuggestionModifyTime string                   `json:"CroppingSuggestionModifyTime" xml:"CroppingSuggestionModifyTime"`
	ImageHeight                  int                      `json:"ImageHeight" xml:"ImageHeight"`
	CroppingSuggestionStatus     string                   `json:"CroppingSuggestionStatus" xml:"CroppingSuggestionStatus"`
	RemarksB                     string                   `json:"RemarksB" xml:"RemarksB"`
	FileSize                     int                      `json:"FileSize" xml:"FileSize"`
	RemarksArrayA                string                   `json:"RemarksArrayA" xml:"RemarksArrayA"`
	FacesModifyTime              string                   `json:"FacesModifyTime" xml:"FacesModifyTime"`
	CelebrityModifyTime          string                   `json:"CelebrityModifyTime" xml:"CelebrityModifyTime"`
	TagsFailReason               string                   `json:"TagsFailReason" xml:"TagsFailReason"`
	ImageUri                     string                   `json:"ImageUri" xml:"ImageUri"`
	ExternalId                   string                   `json:"ExternalId" xml:"ExternalId"`
	Location                     string                   `json:"Location" xml:"Location"`
	TagsModifyTime               string                   `json:"TagsModifyTime" xml:"TagsModifyTime"`
	SourcePosition               string                   `json:"SourcePosition" xml:"SourcePosition"`
	ImageQualityFailReason       string                   `json:"ImageQualityFailReason" xml:"ImageQualityFailReason"`
	OCRFailReason                string                   `json:"OCRFailReason" xml:"OCRFailReason"`
	RemarksC                     string                   `json:"RemarksC" xml:"RemarksC"`
	SourceUri                    string                   `json:"SourceUri" xml:"SourceUri"`
	ImageWidth                   int                      `json:"ImageWidth" xml:"ImageWidth"`
	AddressModifyTime            string                   `json:"AddressModifyTime" xml:"AddressModifyTime"`
	ModifyTime                   string                   `json:"ModifyTime" xml:"ModifyTime"`
	AddressFailReason            string                   `json:"AddressFailReason" xml:"AddressFailReason"`
	CreateTime                   string                   `json:"CreateTime" xml:"CreateTime"`
	OCRModifyTime                string                   `json:"OCRModifyTime" xml:"OCRModifyTime"`
	ImageFormat                  string                   `json:"ImageFormat" xml:"ImageFormat"`
	SourceType                   string                   `json:"SourceType" xml:"SourceType"`
	Orientation                  string                   `json:"Orientation" xml:"Orientation"`
	TagsStatus                   string                   `json:"TagsStatus" xml:"TagsStatus"`
	RemarksD                     string                   `json:"RemarksD" xml:"RemarksD"`
	CelebrityStatus              string                   `json:"CelebrityStatus" xml:"CelebrityStatus"`
	OCRStatus                    string                   `json:"OCRStatus" xml:"OCRStatus"`
	ImageTime                    string                   `json:"ImageTime" xml:"ImageTime"`
	ImageQualityStatus           string                   `json:"ImageQualityStatus" xml:"ImageQualityStatus"`
	RemarksArrayB                string                   `json:"RemarksArrayB" xml:"RemarksArrayB"`
	RemarksA                     string                   `json:"RemarksA" xml:"RemarksA"`
	ImageQualityModifyTime       string                   `json:"ImageQualityModifyTime" xml:"ImageQualityModifyTime"`
	ImageQuality                 ImageQuality             `json:"ImageQuality" xml:"ImageQuality"`
	Address                      Address                  `json:"Address" xml:"Address"`
	OCR                          []OCRItem                `json:"OCR" xml:"OCR"`
	Faces                        []FacesItem              `json:"Faces" xml:"Faces"`
	Celebrity                    []CelebrityItem          `json:"Celebrity" xml:"Celebrity"`
	Tags                         []TagsItem               `json:"Tags" xml:"Tags"`
	CroppingSuggestion           []CroppingSuggestionItem `json:"CroppingSuggestion" xml:"CroppingSuggestion"`
}
