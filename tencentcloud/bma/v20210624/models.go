// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20210624

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type BrandData struct {
	// 商标名称
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 商标证明
	BrandCertificateName *string `json:"BrandCertificateName,omitnil" name:"BrandCertificateName"`

	// 商标审核状态 1-审核中 2-审核未通过 3-审核通过
	BrandStatus *int64 `json:"BrandStatus,omitnil" name:"BrandStatus"`

	// 审核说明
	BrandNote *string `json:"BrandNote,omitnil" name:"BrandNote"`

	// 商标转让证明
	TransferName *string `json:"TransferName,omitnil" name:"TransferName"`

	// 商标转让证明审核状态
	TransferStatus *int64 `json:"TransferStatus,omitnil" name:"TransferStatus"`

	// 审核说明 1-审核中 2-审核未通过 3-审核通过
	TransferNote *string `json:"TransferNote,omitnil" name:"TransferNote"`
}

// Predefined struct for user
type CreateBPFakeURLRequestParams struct {
	// 保护网址ID
	ProtectURLId *int64 `json:"ProtectURLId,omitnil" name:"ProtectURLId"`

	// 仿冒网址
	FakeURL *string `json:"FakeURL,omitnil" name:"FakeURL"`

	// 截图
	SnapshotNames []*string `json:"SnapshotNames,omitnil" name:"SnapshotNames"`

	// 举报说明
	Note *string `json:"Note,omitnil" name:"Note"`
}

type CreateBPFakeURLRequest struct {
	*tchttp.BaseRequest
	
	// 保护网址ID
	ProtectURLId *int64 `json:"ProtectURLId,omitnil" name:"ProtectURLId"`

	// 仿冒网址
	FakeURL *string `json:"FakeURL,omitnil" name:"FakeURL"`

	// 截图
	SnapshotNames []*string `json:"SnapshotNames,omitnil" name:"SnapshotNames"`

	// 举报说明
	Note *string `json:"Note,omitnil" name:"Note"`
}

func (r *CreateBPFakeURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectURLId")
	delete(f, "FakeURL")
	delete(f, "SnapshotNames")
	delete(f, "Note")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPFakeURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFakeURLResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPFakeURLResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPFakeURLResponseParams `json:"Response"`
}

func (r *CreateBPFakeURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFalseTicketRequestParams struct {
	// 仿冒网址
	FakeURL *string `json:"FakeURL,omitnil" name:"FakeURL"`
}

type CreateBPFalseTicketRequest struct {
	*tchttp.BaseRequest
	
	// 仿冒网址
	FakeURL *string `json:"FakeURL,omitnil" name:"FakeURL"`
}

func (r *CreateBPFalseTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFalseTicketRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FakeURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPFalseTicketRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFalseTicketResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPFalseTicketResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPFalseTicketResponseParams `json:"Response"`
}

func (r *CreateBPFalseTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFalseTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPOfflineAttachmentRequestParams struct {
	// 品牌名字
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 商标证明
	BrandCertificateName *string `json:"BrandCertificateName,omitnil" name:"BrandCertificateName"`

	// 商标转让证明
	TransferName *string `json:"TransferName,omitnil" name:"TransferName"`

	// 授权书
	AuthorizationName *string `json:"AuthorizationName,omitnil" name:"AuthorizationName"`
}

type CreateBPOfflineAttachmentRequest struct {
	*tchttp.BaseRequest
	
	// 品牌名字
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 商标证明
	BrandCertificateName *string `json:"BrandCertificateName,omitnil" name:"BrandCertificateName"`

	// 商标转让证明
	TransferName *string `json:"TransferName,omitnil" name:"TransferName"`

	// 授权书
	AuthorizationName *string `json:"AuthorizationName,omitnil" name:"AuthorizationName"`
}

func (r *CreateBPOfflineAttachmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPOfflineAttachmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandName")
	delete(f, "BrandCertificateName")
	delete(f, "TransferName")
	delete(f, "AuthorizationName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPOfflineAttachmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPOfflineAttachmentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPOfflineAttachmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPOfflineAttachmentResponseParams `json:"Response"`
}

func (r *CreateBPOfflineAttachmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPOfflineAttachmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPOfflineTicketRequestParams struct {
	// 仿冒网址ID
	FakeURLId *int64 `json:"FakeURLId,omitnil" name:"FakeURLId"`
}

type CreateBPOfflineTicketRequest struct {
	*tchttp.BaseRequest
	
	// 仿冒网址ID
	FakeURLId *int64 `json:"FakeURLId,omitnil" name:"FakeURLId"`
}

func (r *CreateBPOfflineTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPOfflineTicketRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FakeURLId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPOfflineTicketRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPOfflineTicketResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPOfflineTicketResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPOfflineTicketResponseParams `json:"Response"`
}

func (r *CreateBPOfflineTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPOfflineTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPProtectURLsRequestParams struct {
	// 企业名称
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// 电话号码
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 营业执照
	LicenseName *string `json:"LicenseName,omitnil" name:"LicenseName"`

	// 保护网站
	ProtectURLs []*string `json:"ProtectURLs,omitnil" name:"ProtectURLs"`

	// 网站名称
	ProtectWebs []*string `json:"ProtectWebs,omitnil" name:"ProtectWebs"`
}

type CreateBPProtectURLsRequest struct {
	*tchttp.BaseRequest
	
	// 企业名称
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// 电话号码
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 营业执照
	LicenseName *string `json:"LicenseName,omitnil" name:"LicenseName"`

	// 保护网站
	ProtectURLs []*string `json:"ProtectURLs,omitnil" name:"ProtectURLs"`

	// 网站名称
	ProtectWebs []*string `json:"ProtectWebs,omitnil" name:"ProtectWebs"`
}

func (r *CreateBPProtectURLsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPProtectURLsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyName")
	delete(f, "Phone")
	delete(f, "LicenseName")
	delete(f, "ProtectURLs")
	delete(f, "ProtectWebs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPProtectURLsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPProtectURLsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPProtectURLsResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPProtectURLsResponseParams `json:"Response"`
}

func (r *CreateBPProtectURLsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPProtectURLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRBlockRequestParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 侵权链接
	TortUrl *string `json:"TortUrl,omitnil" name:"TortUrl"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitnil" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitnil" name:"TortPlat"`

	// 拦截结果回调地址
	BlockUrl *string `json:"BlockUrl,omitnil" name:"BlockUrl"`

	// 授权书下载地址
	FileUrl *string `json:"FileUrl,omitnil" name:"FileUrl"`

	// 授权书生效日期
	ValidStartDate *string `json:"ValidStartDate,omitnil" name:"ValidStartDate"`

	// 授权书截止日期
	ValidEndDate *string `json:"ValidEndDate,omitnil" name:"ValidEndDate"`

	// 侵权截图
	TortPic *string `json:"TortPic,omitnil" name:"TortPic"`

	// 委托书下载地址
	CommFileUrl *string `json:"CommFileUrl,omitnil" name:"CommFileUrl"`

	// 委托书生效日期
	CommValidStartDate *string `json:"CommValidStartDate,omitnil" name:"CommValidStartDate"`

	// 委托书截止日期
	CommValidEndDate *string `json:"CommValidEndDate,omitnil" name:"CommValidEndDate"`

	// 是否著作权人：0-否 1-是
	IsProducer *string `json:"IsProducer,omitnil" name:"IsProducer"`

	// 存证证书下载地址
	EvidenceFileUrl *string `json:"EvidenceFileUrl,omitnil" name:"EvidenceFileUrl"`

	// 存证证书生效日期
	EvidenceValidStartDate *string `json:"EvidenceValidStartDate,omitnil" name:"EvidenceValidStartDate"`

	// 存证证书截止日期
	EvidenceValidEndDate *string `json:"EvidenceValidEndDate,omitnil" name:"EvidenceValidEndDate"`
}

type CreateCRBlockRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 侵权链接
	TortUrl *string `json:"TortUrl,omitnil" name:"TortUrl"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitnil" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitnil" name:"TortPlat"`

	// 拦截结果回调地址
	BlockUrl *string `json:"BlockUrl,omitnil" name:"BlockUrl"`

	// 授权书下载地址
	FileUrl *string `json:"FileUrl,omitnil" name:"FileUrl"`

	// 授权书生效日期
	ValidStartDate *string `json:"ValidStartDate,omitnil" name:"ValidStartDate"`

	// 授权书截止日期
	ValidEndDate *string `json:"ValidEndDate,omitnil" name:"ValidEndDate"`

	// 侵权截图
	TortPic *string `json:"TortPic,omitnil" name:"TortPic"`

	// 委托书下载地址
	CommFileUrl *string `json:"CommFileUrl,omitnil" name:"CommFileUrl"`

	// 委托书生效日期
	CommValidStartDate *string `json:"CommValidStartDate,omitnil" name:"CommValidStartDate"`

	// 委托书截止日期
	CommValidEndDate *string `json:"CommValidEndDate,omitnil" name:"CommValidEndDate"`

	// 是否著作权人：0-否 1-是
	IsProducer *string `json:"IsProducer,omitnil" name:"IsProducer"`

	// 存证证书下载地址
	EvidenceFileUrl *string `json:"EvidenceFileUrl,omitnil" name:"EvidenceFileUrl"`

	// 存证证书生效日期
	EvidenceValidStartDate *string `json:"EvidenceValidStartDate,omitnil" name:"EvidenceValidStartDate"`

	// 存证证书截止日期
	EvidenceValidEndDate *string `json:"EvidenceValidEndDate,omitnil" name:"EvidenceValidEndDate"`
}

func (r *CreateCRBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "TortUrl")
	delete(f, "TortTitle")
	delete(f, "TortPlat")
	delete(f, "BlockUrl")
	delete(f, "FileUrl")
	delete(f, "ValidStartDate")
	delete(f, "ValidEndDate")
	delete(f, "TortPic")
	delete(f, "CommFileUrl")
	delete(f, "CommValidStartDate")
	delete(f, "CommValidEndDate")
	delete(f, "IsProducer")
	delete(f, "EvidenceFileUrl")
	delete(f, "EvidenceValidStartDate")
	delete(f, "EvidenceValidEndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRBlockResponseParams struct {
	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 该字段已废弃
	TortNum *string `json:"TortNum,omitnil" name:"TortNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCRBlockResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRBlockResponseParams `json:"Response"`
}

func (r *CreateCRBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRCompanyVerifyRequestParams struct {
	// 企业名称
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// 企业证件号码
	CompanyID *string `json:"CompanyID,omitnil" name:"CompanyID"`

	// 企业法人姓名
	CompanyLegalName *string `json:"CompanyLegalName,omitnil" name:"CompanyLegalName"`

	// 联系人姓名
	ManagerName *string `json:"ManagerName,omitnil" name:"ManagerName"`

	// 联系人手机号
	ManagerPhone *string `json:"ManagerPhone,omitnil" name:"ManagerPhone"`

	// 手机验证码，接口接入可以置空
	VerificationCode *string `json:"VerificationCode,omitnil" name:"VerificationCode"`

	// 字段已废弃，企业认证号码类型 1：社会信用代码 2：组织机构代码 3：企业工商注册码 4：其他 默认为1
	CompanyIDType *string `json:"CompanyIDType,omitnil" name:"CompanyIDType"`

	// 字段已废弃，认证类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

type CreateCRCompanyVerifyRequest struct {
	*tchttp.BaseRequest
	
	// 企业名称
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// 企业证件号码
	CompanyID *string `json:"CompanyID,omitnil" name:"CompanyID"`

	// 企业法人姓名
	CompanyLegalName *string `json:"CompanyLegalName,omitnil" name:"CompanyLegalName"`

	// 联系人姓名
	ManagerName *string `json:"ManagerName,omitnil" name:"ManagerName"`

	// 联系人手机号
	ManagerPhone *string `json:"ManagerPhone,omitnil" name:"ManagerPhone"`

	// 手机验证码，接口接入可以置空
	VerificationCode *string `json:"VerificationCode,omitnil" name:"VerificationCode"`

	// 字段已废弃，企业认证号码类型 1：社会信用代码 2：组织机构代码 3：企业工商注册码 4：其他 默认为1
	CompanyIDType *string `json:"CompanyIDType,omitnil" name:"CompanyIDType"`

	// 字段已废弃，认证类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *CreateCRCompanyVerifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRCompanyVerifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyName")
	delete(f, "CompanyID")
	delete(f, "CompanyLegalName")
	delete(f, "ManagerName")
	delete(f, "ManagerPhone")
	delete(f, "VerificationCode")
	delete(f, "CompanyIDType")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRCompanyVerifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRCompanyVerifyResponseParams struct {
	// 认证状态：0-认证成功 1-认证失败
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 认证状态说明，包括认证失败的原因
	Note *string `json:"Note,omitnil" name:"Note"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCRCompanyVerifyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRCompanyVerifyResponseParams `json:"Response"`
}

func (r *CreateCRCompanyVerifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRCompanyVerifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRDesktopCodeRequestParams struct {
	// xxx
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// xxx
	DesktopCode *string `json:"DesktopCode,omitnil" name:"DesktopCode"`
}

type CreateCRDesktopCodeRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// xxx
	DesktopCode *string `json:"DesktopCode,omitnil" name:"DesktopCode"`
}

func (r *CreateCRDesktopCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRDesktopCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TortId")
	delete(f, "DesktopCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRDesktopCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRDesktopCodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCRDesktopCodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRDesktopCodeResponseParams `json:"Response"`
}

func (r *CreateCRDesktopCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRDesktopCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRRightFileRequestParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 权属文件列表
	FileList []*File `json:"FileList,omitnil" name:"FileList"`
}

type CreateCRRightFileRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 权属文件列表
	FileList []*File `json:"FileList,omitnil" name:"FileList"`
}

func (r *CreateCRRightFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRRightFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "FileList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRRightFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRRightFileResponseParams struct {
	// 权属文件Id，按提交顺序排序
	FileIds []*int64 `json:"FileIds,omitnil" name:"FileIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCRRightFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRRightFileResponseParams `json:"Response"`
}

func (r *CreateCRRightFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRRightFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRRightRequestParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 侵权链接
	TortUrl *string `json:"TortUrl,omitnil" name:"TortUrl"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitnil" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitnil" name:"TortPlat"`

	// 发函结果回调地址
	RightUrl *string `json:"RightUrl,omitnil" name:"RightUrl"`

	// 授权书下载地址
	FileUrl *string `json:"FileUrl,omitnil" name:"FileUrl"`

	// 授权书生效日期
	ValidStartDate *string `json:"ValidStartDate,omitnil" name:"ValidStartDate"`

	// 授权书截止日期
	ValidEndDate *string `json:"ValidEndDate,omitnil" name:"ValidEndDate"`

	// 委托书下载地址
	CommFileUrl *string `json:"CommFileUrl,omitnil" name:"CommFileUrl"`

	// 委托书生效日期
	CommValidStartDate *string `json:"CommValidStartDate,omitnil" name:"CommValidStartDate"`

	// 委托书截止日期
	CommValidEndDate *string `json:"CommValidEndDate,omitnil" name:"CommValidEndDate"`

	// 主页下载地址
	HomeFileUrl *string `json:"HomeFileUrl,omitnil" name:"HomeFileUrl"`

	// 主页生效日期
	HomeValidStartDate *string `json:"HomeValidStartDate,omitnil" name:"HomeValidStartDate"`

	// 主页截止日期
	HomeValidEndDate *string `json:"HomeValidEndDate,omitnil" name:"HomeValidEndDate"`

	// 是否著作权人：0-否 1-是
	IsProducer *string `json:"IsProducer,omitnil" name:"IsProducer"`

	// 存证证书下载地址
	EvidenceFileUrl *string `json:"EvidenceFileUrl,omitnil" name:"EvidenceFileUrl"`

	// 存证证书生效日期
	EvidenceValidStartDate *string `json:"EvidenceValidStartDate,omitnil" name:"EvidenceValidStartDate"`

	// 存证证书截止日期
	EvidenceValidEndDate *string `json:"EvidenceValidEndDate,omitnil" name:"EvidenceValidEndDate"`
}

type CreateCRRightRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 侵权链接
	TortUrl *string `json:"TortUrl,omitnil" name:"TortUrl"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitnil" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitnil" name:"TortPlat"`

	// 发函结果回调地址
	RightUrl *string `json:"RightUrl,omitnil" name:"RightUrl"`

	// 授权书下载地址
	FileUrl *string `json:"FileUrl,omitnil" name:"FileUrl"`

	// 授权书生效日期
	ValidStartDate *string `json:"ValidStartDate,omitnil" name:"ValidStartDate"`

	// 授权书截止日期
	ValidEndDate *string `json:"ValidEndDate,omitnil" name:"ValidEndDate"`

	// 委托书下载地址
	CommFileUrl *string `json:"CommFileUrl,omitnil" name:"CommFileUrl"`

	// 委托书生效日期
	CommValidStartDate *string `json:"CommValidStartDate,omitnil" name:"CommValidStartDate"`

	// 委托书截止日期
	CommValidEndDate *string `json:"CommValidEndDate,omitnil" name:"CommValidEndDate"`

	// 主页下载地址
	HomeFileUrl *string `json:"HomeFileUrl,omitnil" name:"HomeFileUrl"`

	// 主页生效日期
	HomeValidStartDate *string `json:"HomeValidStartDate,omitnil" name:"HomeValidStartDate"`

	// 主页截止日期
	HomeValidEndDate *string `json:"HomeValidEndDate,omitnil" name:"HomeValidEndDate"`

	// 是否著作权人：0-否 1-是
	IsProducer *string `json:"IsProducer,omitnil" name:"IsProducer"`

	// 存证证书下载地址
	EvidenceFileUrl *string `json:"EvidenceFileUrl,omitnil" name:"EvidenceFileUrl"`

	// 存证证书生效日期
	EvidenceValidStartDate *string `json:"EvidenceValidStartDate,omitnil" name:"EvidenceValidStartDate"`

	// 存证证书截止日期
	EvidenceValidEndDate *string `json:"EvidenceValidEndDate,omitnil" name:"EvidenceValidEndDate"`
}

func (r *CreateCRRightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRRightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "TortUrl")
	delete(f, "TortTitle")
	delete(f, "TortPlat")
	delete(f, "RightUrl")
	delete(f, "FileUrl")
	delete(f, "ValidStartDate")
	delete(f, "ValidEndDate")
	delete(f, "CommFileUrl")
	delete(f, "CommValidStartDate")
	delete(f, "CommValidEndDate")
	delete(f, "HomeFileUrl")
	delete(f, "HomeValidStartDate")
	delete(f, "HomeValidEndDate")
	delete(f, "IsProducer")
	delete(f, "EvidenceFileUrl")
	delete(f, "EvidenceValidStartDate")
	delete(f, "EvidenceValidEndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRRightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRRightResponseParams struct {
	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 该字段已废弃
	TortNum *string `json:"TortNum,omitnil" name:"TortNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCRRightResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRRightResponseParams `json:"Response"`
}

func (r *CreateCRRightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRRightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRTortRequestParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 侵权网址
	TortURL *string `json:"TortURL,omitnil" name:"TortURL"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitnil" name:"TortPlat"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitnil" name:"TortTitle"`
}

type CreateCRTortRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 侵权网址
	TortURL *string `json:"TortURL,omitnil" name:"TortURL"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitnil" name:"TortPlat"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitnil" name:"TortTitle"`
}

func (r *CreateCRTortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRTortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "TortURL")
	delete(f, "TortPlat")
	delete(f, "TortTitle")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRTortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRTortResponseParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitnil" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitnil" name:"TortPlat"`

	// 侵权网址
	TortURL *string `json:"TortURL,omitnil" name:"TortURL"`

	// 侵权域名
	TortDomain *string `json:"TortDomain,omitnil" name:"TortDomain"`

	// 侵权主体
	TortBodyName *string `json:"TortBodyName,omitnil" name:"TortBodyName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCRTortResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRTortResponseParams `json:"Response"`
}

func (r *CreateCRTortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRTortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRUserVerifyRequestParams struct {
	// 用户真实姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户身份证号
	UserID *string `json:"UserID,omitnil" name:"UserID"`

	// 用户手机号码
	UserPhone *string `json:"UserPhone,omitnil" name:"UserPhone"`

	// 短信验证码，接口接入可以置空
	VerificationCode *string `json:"VerificationCode,omitnil" name:"VerificationCode"`

	// 字段已废弃，认证类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

type CreateCRUserVerifyRequest struct {
	*tchttp.BaseRequest
	
	// 用户真实姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户身份证号
	UserID *string `json:"UserID,omitnil" name:"UserID"`

	// 用户手机号码
	UserPhone *string `json:"UserPhone,omitnil" name:"UserPhone"`

	// 短信验证码，接口接入可以置空
	VerificationCode *string `json:"VerificationCode,omitnil" name:"VerificationCode"`

	// 字段已废弃，认证类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *CreateCRUserVerifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRUserVerifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "UserID")
	delete(f, "UserPhone")
	delete(f, "VerificationCode")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRUserVerifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRUserVerifyResponseParams struct {
	// 认证状态：0-认证成功 1-认证失败
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 认证状态说明，包括认证失败原因等
	Note *string `json:"Note,omitnil" name:"Note"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCRUserVerifyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRUserVerifyResponseParams `json:"Response"`
}

func (r *CreateCRUserVerifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRUserVerifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRWorkRequestParams struct {
	// 作品名称
	WorkName *string `json:"WorkName,omitnil" name:"WorkName"`

	// 作品分类
	WorkCategory *string `json:"WorkCategory,omitnil" name:"WorkCategory"`

	// 作品内容类型
	WorkType *string `json:"WorkType,omitnil" name:"WorkType"`

	// 作品标签
	WorkSign *string `json:"WorkSign,omitnil" name:"WorkSign"`

	// 字段已废弃，作品图片
	WorkPic *string `json:"WorkPic,omitnil" name:"WorkPic"`

	// 作品描述
	WorkDesc *string `json:"WorkDesc,omitnil" name:"WorkDesc"`

	// 是否原创：0-否 1-是
	IsOriginal *string `json:"IsOriginal,omitnil" name:"IsOriginal"`

	// 是否发布：0-未发布 1-已发布
	IsRelease *string `json:"IsRelease,omitnil" name:"IsRelease"`

	// 字段已废弃，著作权人ID
	ProducerID *int64 `json:"ProducerID,omitnil" name:"ProducerID"`

	// 创作时间
	ProduceTime *string `json:"ProduceTime,omitnil" name:"ProduceTime"`

	// 字段已废弃
	SampleContentURL *string `json:"SampleContentURL,omitnil" name:"SampleContentURL"`

	// 作品下载地址
	SampleDownloadURL *string `json:"SampleDownloadURL,omitnil" name:"SampleDownloadURL"`

	// 作品在线地址
	SamplePublicURL *string `json:"SamplePublicURL,omitnil" name:"SamplePublicURL"`

	// 字段已废弃，授予类型
	GrantType *string `json:"GrantType,omitnil" name:"GrantType"`

	// 是否监测：0-不监测 1-监测
	IsMonitor *string `json:"IsMonitor,omitnil" name:"IsMonitor"`

	// 是否存证：0-不存证  2-存证 注意是2
	IsCert *string `json:"IsCert,omitnil" name:"IsCert"`

	// 存证回调地址
	CertUrl *string `json:"CertUrl,omitnil" name:"CertUrl"`

	// 监测回调地址
	MonitorUrl *string `json:"MonitorUrl,omitnil" name:"MonitorUrl"`

	// 字段已废弃，创作性质
	ProduceType *string `json:"ProduceType,omitnil" name:"ProduceType"`

	// 白名单列表
	WhiteLists []*string `json:"WhiteLists,omitnil" name:"WhiteLists"`

	// 作品ID，忽略该字段
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 著作权人姓名
	ProducerName *string `json:"ProducerName,omitnil" name:"ProducerName"`

	// 作者，小说类型必填
	Nickname *string `json:"Nickname,omitnil" name:"Nickname"`

	// 授权书下载地址
	Authorization *string `json:"Authorization,omitnil" name:"Authorization"`

	// 授权书开始时间
	AuthorizationStartTime *string `json:"AuthorizationStartTime,omitnil" name:"AuthorizationStartTime"`

	// 授权书结束时间
	AuthorizationEndTime *string `json:"AuthorizationEndTime,omitnil" name:"AuthorizationEndTime"`

	// 内容格式，支持txt、doc等，表示Content的具体格式
	ContentType *string `json:"ContentType,omitnil" name:"ContentType"`

	// 文件内容base64编码，该字段仅在无法提供下载链接时使用
	Content *string `json:"Content,omitnil" name:"Content"`

	// 监测结束时间
	MonitorEndTime *string `json:"MonitorEndTime,omitnil" name:"MonitorEndTime"`

	// 申请人ID，用于存证和取证
	ApplierId *string `json:"ApplierId,omitnil" name:"ApplierId"`

	// 申请人姓名，用于存证和取证
	ApplierName *string `json:"ApplierName,omitnil" name:"ApplierName"`

	// 是否自动续期
	IsAutoRenew *string `json:"IsAutoRenew,omitnil" name:"IsAutoRenew"`
}

type CreateCRWorkRequest struct {
	*tchttp.BaseRequest
	
	// 作品名称
	WorkName *string `json:"WorkName,omitnil" name:"WorkName"`

	// 作品分类
	WorkCategory *string `json:"WorkCategory,omitnil" name:"WorkCategory"`

	// 作品内容类型
	WorkType *string `json:"WorkType,omitnil" name:"WorkType"`

	// 作品标签
	WorkSign *string `json:"WorkSign,omitnil" name:"WorkSign"`

	// 字段已废弃，作品图片
	WorkPic *string `json:"WorkPic,omitnil" name:"WorkPic"`

	// 作品描述
	WorkDesc *string `json:"WorkDesc,omitnil" name:"WorkDesc"`

	// 是否原创：0-否 1-是
	IsOriginal *string `json:"IsOriginal,omitnil" name:"IsOriginal"`

	// 是否发布：0-未发布 1-已发布
	IsRelease *string `json:"IsRelease,omitnil" name:"IsRelease"`

	// 字段已废弃，著作权人ID
	ProducerID *int64 `json:"ProducerID,omitnil" name:"ProducerID"`

	// 创作时间
	ProduceTime *string `json:"ProduceTime,omitnil" name:"ProduceTime"`

	// 字段已废弃
	SampleContentURL *string `json:"SampleContentURL,omitnil" name:"SampleContentURL"`

	// 作品下载地址
	SampleDownloadURL *string `json:"SampleDownloadURL,omitnil" name:"SampleDownloadURL"`

	// 作品在线地址
	SamplePublicURL *string `json:"SamplePublicURL,omitnil" name:"SamplePublicURL"`

	// 字段已废弃，授予类型
	GrantType *string `json:"GrantType,omitnil" name:"GrantType"`

	// 是否监测：0-不监测 1-监测
	IsMonitor *string `json:"IsMonitor,omitnil" name:"IsMonitor"`

	// 是否存证：0-不存证  2-存证 注意是2
	IsCert *string `json:"IsCert,omitnil" name:"IsCert"`

	// 存证回调地址
	CertUrl *string `json:"CertUrl,omitnil" name:"CertUrl"`

	// 监测回调地址
	MonitorUrl *string `json:"MonitorUrl,omitnil" name:"MonitorUrl"`

	// 字段已废弃，创作性质
	ProduceType *string `json:"ProduceType,omitnil" name:"ProduceType"`

	// 白名单列表
	WhiteLists []*string `json:"WhiteLists,omitnil" name:"WhiteLists"`

	// 作品ID，忽略该字段
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 著作权人姓名
	ProducerName *string `json:"ProducerName,omitnil" name:"ProducerName"`

	// 作者，小说类型必填
	Nickname *string `json:"Nickname,omitnil" name:"Nickname"`

	// 授权书下载地址
	Authorization *string `json:"Authorization,omitnil" name:"Authorization"`

	// 授权书开始时间
	AuthorizationStartTime *string `json:"AuthorizationStartTime,omitnil" name:"AuthorizationStartTime"`

	// 授权书结束时间
	AuthorizationEndTime *string `json:"AuthorizationEndTime,omitnil" name:"AuthorizationEndTime"`

	// 内容格式，支持txt、doc等，表示Content的具体格式
	ContentType *string `json:"ContentType,omitnil" name:"ContentType"`

	// 文件内容base64编码，该字段仅在无法提供下载链接时使用
	Content *string `json:"Content,omitnil" name:"Content"`

	// 监测结束时间
	MonitorEndTime *string `json:"MonitorEndTime,omitnil" name:"MonitorEndTime"`

	// 申请人ID，用于存证和取证
	ApplierId *string `json:"ApplierId,omitnil" name:"ApplierId"`

	// 申请人姓名，用于存证和取证
	ApplierName *string `json:"ApplierName,omitnil" name:"ApplierName"`

	// 是否自动续期
	IsAutoRenew *string `json:"IsAutoRenew,omitnil" name:"IsAutoRenew"`
}

func (r *CreateCRWorkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRWorkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkName")
	delete(f, "WorkCategory")
	delete(f, "WorkType")
	delete(f, "WorkSign")
	delete(f, "WorkPic")
	delete(f, "WorkDesc")
	delete(f, "IsOriginal")
	delete(f, "IsRelease")
	delete(f, "ProducerID")
	delete(f, "ProduceTime")
	delete(f, "SampleContentURL")
	delete(f, "SampleDownloadURL")
	delete(f, "SamplePublicURL")
	delete(f, "GrantType")
	delete(f, "IsMonitor")
	delete(f, "IsCert")
	delete(f, "CertUrl")
	delete(f, "MonitorUrl")
	delete(f, "ProduceType")
	delete(f, "WhiteLists")
	delete(f, "WorkId")
	delete(f, "ProducerName")
	delete(f, "Nickname")
	delete(f, "Authorization")
	delete(f, "AuthorizationStartTime")
	delete(f, "AuthorizationEndTime")
	delete(f, "ContentType")
	delete(f, "Content")
	delete(f, "MonitorEndTime")
	delete(f, "ApplierId")
	delete(f, "ApplierName")
	delete(f, "IsAutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRWorkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRWorkResponseParams struct {
	// 作品ID，一个作品对应唯一的workid
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 存证ID，忽略该字段
	EvidenceId *int64 `json:"EvidenceId,omitnil" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCRWorkResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRWorkResponseParams `json:"Response"`
}

func (r *CreateCRWorkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRWorkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPCompanyInfoRequestParams struct {

}

type DescribeBPCompanyInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBPCompanyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPCompanyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPCompanyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPCompanyInfoResponseParams struct {
	// 企业名称
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// 电话号码
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 营业执照
	LicenseName *string `json:"LicenseName,omitnil" name:"LicenseName"`

	// 营业执照审核状态 1-审核中 2-审核未通过，3、审核通过
	LicenseStatus *int64 `json:"LicenseStatus,omitnil" name:"LicenseStatus"`

	// 营业执照备注
	LicenseNote *string `json:"LicenseNote,omitnil" name:"LicenseNote"`

	// 授权书
	AuthorizationName *string `json:"AuthorizationName,omitnil" name:"AuthorizationName"`

	// 授权书审核状态
	AuthorizationStatus *int64 `json:"AuthorizationStatus,omitnil" name:"AuthorizationStatus"`

	// 授权书备注
	AuthorizationNote *string `json:"AuthorizationNote,omitnil" name:"AuthorizationNote"`

	// 品牌信息
	BrandDatas []*BrandData `json:"BrandDatas,omitnil" name:"BrandDatas"`

	// 企业ID
	CompanyId *int64 `json:"CompanyId,omitnil" name:"CompanyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBPCompanyInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPCompanyInfoResponseParams `json:"Response"`
}

func (r *DescribeBPCompanyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPCompanyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPFakeURLsRequestParams struct {
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

type DescribeBPFakeURLsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

func (r *DescribeBPFakeURLsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPFakeURLsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPFakeURLsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPFakeURLsResponseParams struct {
	// 仿冒网址列表
	FakeURLInfos []*FakeURLInfo `json:"FakeURLInfos,omitnil" name:"FakeURLInfos"`

	// 总量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 导出量
	ExportURL *string `json:"ExportURL,omitnil" name:"ExportURL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBPFakeURLsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPFakeURLsResponseParams `json:"Response"`
}

func (r *DescribeBPFakeURLsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPFakeURLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPProtectURLsRequestParams struct {
	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

type DescribeBPProtectURLsRequest struct {
	*tchttp.BaseRequest
	
	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

func (r *DescribeBPProtectURLsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPProtectURLsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPProtectURLsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPProtectURLsResponseParams struct {
	// 保护网址列表
	ProtectURLInfos []*ProtectURLInfo `json:"ProtectURLInfos,omitnil" name:"ProtectURLInfos"`

	// 总量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBPProtectURLsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPProtectURLsResponseParams `json:"Response"`
}

func (r *DescribeBPProtectURLsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPProtectURLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPReportFakeURLsRequestParams struct {
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

type DescribeBPReportFakeURLsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

func (r *DescribeBPReportFakeURLsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPReportFakeURLsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPReportFakeURLsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPReportFakeURLsResponseParams struct {
	// 举报网站列表
	ReportFakeURLInfos []*ReportFakeURLInfo `json:"ReportFakeURLInfos,omitnil" name:"ReportFakeURLInfos"`

	// 总量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBPReportFakeURLsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPReportFakeURLsResponseParams `json:"Response"`
}

func (r *DescribeBPReportFakeURLsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPReportFakeURLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRMonitorDetailRequestParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeCRMonitorDetailRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeCRMonitorDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRMonitorDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCRMonitorDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRMonitorDetailResponseParams struct {
	// 侵权数组
	Torts []*MonitorTort `json:"Torts,omitnil" name:"Torts"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 监测状态
	MonitorStatus *int64 `json:"MonitorStatus,omitnil" name:"MonitorStatus"`

	// 导出地址
	ExportURL *string `json:"ExportURL,omitnil" name:"ExportURL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCRMonitorDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCRMonitorDetailResponseParams `json:"Response"`
}

func (r *DescribeCRMonitorDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRMonitorDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRMonitorsRequestParams struct {
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

type DescribeCRMonitorsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

func (r *DescribeCRMonitorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRMonitorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCRMonitorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRMonitorsResponseParams struct {
	// 监测结果
	Monitors []*Monitor `json:"Monitors,omitnil" name:"Monitors"`

	// 记录总条数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 导出地址
	ExportURL *string `json:"ExportURL,omitnil" name:"ExportURL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCRMonitorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCRMonitorsResponseParams `json:"Response"`
}

func (r *DescribeCRMonitorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRMonitorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRObtainDetailRequestParams struct {
	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`
}

type DescribeCRObtainDetailRequest struct {
	*tchttp.BaseRequest
	
	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`
}

func (r *DescribeCRObtainDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRObtainDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TortId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCRObtainDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRObtainDetailResponseParams struct {
	// 作品名称
	WorkName *string `json:"WorkName,omitnil" name:"WorkName"`

	// 侵权链接
	TortURL *string `json:"TortURL,omitnil" name:"TortURL"`

	// 取证时间
	ObtainTime *string `json:"ObtainTime,omitnil" name:"ObtainTime"`

	// 取证类型
	ObtainType *string `json:"ObtainType,omitnil" name:"ObtainType"`

	// 取证号
	ObtainNum *string `json:"ObtainNum,omitnil" name:"ObtainNum"`

	// 证据地址
	DepositFile *string `json:"DepositFile,omitnil" name:"DepositFile"`

	// 公证信息地址
	DepositCert *string `json:"DepositCert,omitnil" name:"DepositCert"`

	// 内容类型
	WorkType *string `json:"WorkType,omitnil" name:"WorkType"`

	// 作品类型
	WorkCategory *string `json:"WorkCategory,omitnil" name:"WorkCategory"`

	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 侵权编号
	TortNum *string `json:"TortNum,omitnil" name:"TortNum"`

	// 取证状态
	ObtainStatus *int64 `json:"ObtainStatus,omitnil" name:"ObtainStatus"`

	// 取证状态说明
	ObtainNote *string `json:"ObtainNote,omitnil" name:"ObtainNote"`

	// 取证时长
	ObtainDuration *string `json:"ObtainDuration,omitnil" name:"ObtainDuration"`

	// 取证名称
	ObtainName *string `json:"ObtainName,omitnil" name:"ObtainName"`

	// 取证公证信息
	DepositPdfCert *string `json:"DepositPdfCert,omitnil" name:"DepositPdfCert"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCRObtainDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCRObtainDetailResponseParams `json:"Response"`
}

func (r *DescribeCRObtainDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRObtainDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRWorkInfoRequestParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`
}

type DescribeCRWorkInfoRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`
}

func (r *DescribeCRWorkInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRWorkInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCRWorkInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRWorkInfoResponseParams struct {
	// 作品名称
	WorkName *string `json:"WorkName,omitnil" name:"WorkName"`

	// 监测状态
	MonitorStatus *int64 `json:"MonitorStatus,omitnil" name:"MonitorStatus"`

	// 授权文件状态
	AuthStatus *int64 `json:"AuthStatus,omitnil" name:"AuthStatus"`

	// 委托书状态
	CommStatus *int64 `json:"CommStatus,omitnil" name:"CommStatus"`

	// 是否著作权人
	IsProducer *int64 `json:"IsProducer,omitnil" name:"IsProducer"`

	// 存证证书状态
	EvidenceStatus *int64 `json:"EvidenceStatus,omitnil" name:"EvidenceStatus"`

	// 作品类型
	WorkCategory *string `json:"WorkCategory,omitnil" name:"WorkCategory"`

	// 是否原创
	IsOriginal *string `json:"IsOriginal,omitnil" name:"IsOriginal"`

	// 是否已发表
	IsRelease *string `json:"IsRelease,omitnil" name:"IsRelease"`

	// 著作权人姓名
	ProducerName *string `json:"ProducerName,omitnil" name:"ProducerName"`

	// 发表时间
	ProduceTime *string `json:"ProduceTime,omitnil" name:"ProduceTime"`

	// 白名单
	WhiteLists []*string `json:"WhiteLists,omitnil" name:"WhiteLists"`

	// 作品描述
	WorkDesc *string `json:"WorkDesc,omitnil" name:"WorkDesc"`

	// 授权书
	Authorization *string `json:"Authorization,omitnil" name:"Authorization"`

	// 授权书生效日期
	AuthorizationStartTime *string `json:"AuthorizationStartTime,omitnil" name:"AuthorizationStartTime"`

	// 授权书截止日期
	AuthorizationEndTime *string `json:"AuthorizationEndTime,omitnil" name:"AuthorizationEndTime"`

	// 委托书
	Commission *string `json:"Commission,omitnil" name:"Commission"`

	// 委托书生效日期
	CommissionStartTime *string `json:"CommissionStartTime,omitnil" name:"CommissionStartTime"`

	// 委托书截止日期
	CommissionEndTime *string `json:"CommissionEndTime,omitnil" name:"CommissionEndTime"`

	// 存证证书
	EvidenceUrl *string `json:"EvidenceUrl,omitnil" name:"EvidenceUrl"`

	// 存证证书生效日期
	EvidenceStartTime *string `json:"EvidenceStartTime,omitnil" name:"EvidenceStartTime"`

	// 存证证书截止日期
	EvidenceEndTime *string `json:"EvidenceEndTime,omitnil" name:"EvidenceEndTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCRWorkInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCRWorkInfoResponseParams `json:"Response"`
}

func (r *DescribeCRWorkInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRWorkInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FakeURLInfo struct {
	// 仿冒网址ID
	FakeURLId *int64 `json:"FakeURLId,omitnil" name:"FakeURLId"`

	// 保护网站
	ProtectWeb *string `json:"ProtectWeb,omitnil" name:"ProtectWeb"`

	// 检测时间
	DetectTime *string `json:"DetectTime,omitnil" name:"DetectTime"`

	// 仿冒网址
	FakeURL *string `json:"FakeURL,omitnil" name:"FakeURL"`

	// 截图
	Snapshot *string `json:"Snapshot,omitnil" name:"Snapshot"`

	// IP地址
	IP *string `json:"IP,omitnil" name:"IP"`

	// IP地理位置
	IPLoc *string `json:"IPLoc,omitnil" name:"IPLoc"`

	// 热度
	Heat *int64 `json:"Heat,omitnil" name:"Heat"`

	// 网址状态
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 备注
	Note *string `json:"Note,omitnil" name:"Note"`

	// 仿冒网站所属单位
	FakeURLCompany *string `json:"FakeURLCompany,omitnil" name:"FakeURLCompany"`

	// 仿冒网站性质
	FakeURLAttr *string `json:"FakeURLAttr,omitnil" name:"FakeURLAttr"`

	// 仿冒网站名称
	FakeURLName *string `json:"FakeURLName,omitnil" name:"FakeURLName"`

	// 仿冒网站备案号
	FakeURLICP *string `json:"FakeURLICP,omitnil" name:"FakeURLICP"`

	// 仿冒网站创建时间
	FakeURLCreateTime *string `json:"FakeURLCreateTime,omitnil" name:"FakeURLCreateTime"`

	// 仿冒网站过期时间
	FakeURLExpireTime *string `json:"FakeURLExpireTime,omitnil" name:"FakeURLExpireTime"`
}

type File struct {
	// 文件下载地址
	FileUrl *string `json:"FileUrl,omitnil" name:"FileUrl"`

	// 文件类型 1-委托书 2-授权书 5-存证证书 11-营业执照
	FileType *int64 `json:"FileType,omitnil" name:"FileType"`

	// 文件有效开始日期
	ValidStartDate *string `json:"ValidStartDate,omitnil" name:"ValidStartDate"`

	// 文件有效截止日期
	ValidEndDate *string `json:"ValidEndDate,omitnil" name:"ValidEndDate"`
}

type Filter struct {

}

// Predefined struct for user
type ModifyBPOfflineAttachmentRequestParams struct {
	// 营业执照
	LicenseName *string `json:"LicenseName,omitnil" name:"LicenseName"`

	// 授权书
	AuthorizationName *string `json:"AuthorizationName,omitnil" name:"AuthorizationName"`

	// 商标名称
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 商标证明
	BrandCertificateName *string `json:"BrandCertificateName,omitnil" name:"BrandCertificateName"`

	// 商标转让证明
	TransferName *string `json:"TransferName,omitnil" name:"TransferName"`
}

type ModifyBPOfflineAttachmentRequest struct {
	*tchttp.BaseRequest
	
	// 营业执照
	LicenseName *string `json:"LicenseName,omitnil" name:"LicenseName"`

	// 授权书
	AuthorizationName *string `json:"AuthorizationName,omitnil" name:"AuthorizationName"`

	// 商标名称
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 商标证明
	BrandCertificateName *string `json:"BrandCertificateName,omitnil" name:"BrandCertificateName"`

	// 商标转让证明
	TransferName *string `json:"TransferName,omitnil" name:"TransferName"`
}

func (r *ModifyBPOfflineAttachmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBPOfflineAttachmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseName")
	delete(f, "AuthorizationName")
	delete(f, "BrandName")
	delete(f, "BrandCertificateName")
	delete(f, "TransferName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBPOfflineAttachmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBPOfflineAttachmentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBPOfflineAttachmentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBPOfflineAttachmentResponseParams `json:"Response"`
}

func (r *ModifyBPOfflineAttachmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBPOfflineAttachmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRBlockStatusRequestParams struct {
	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 拦截结果回调地址
	BlockUrl *string `json:"BlockUrl,omitnil" name:"BlockUrl"`
}

type ModifyCRBlockStatusRequest struct {
	*tchttp.BaseRequest
	
	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 拦截结果回调地址
	BlockUrl *string `json:"BlockUrl,omitnil" name:"BlockUrl"`
}

func (r *ModifyCRBlockStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRBlockStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TortId")
	delete(f, "BlockUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCRBlockStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRBlockStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCRBlockStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCRBlockStatusResponseParams `json:"Response"`
}

func (r *ModifyCRBlockStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRBlockStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRMonitorRequestParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 监测状态：1-开启监测 2-关闭监测
	MonitorStatus *string `json:"MonitorStatus,omitnil" name:"MonitorStatus"`

	// 监测截止时间
	MonitorEnd *string `json:"MonitorEnd,omitnil" name:"MonitorEnd"`
}

type ModifyCRMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 监测状态：1-开启监测 2-关闭监测
	MonitorStatus *string `json:"MonitorStatus,omitnil" name:"MonitorStatus"`

	// 监测截止时间
	MonitorEnd *string `json:"MonitorEnd,omitnil" name:"MonitorEnd"`
}

func (r *ModifyCRMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "MonitorStatus")
	delete(f, "MonitorEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCRMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRMonitorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCRMonitorResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCRMonitorResponseParams `json:"Response"`
}

func (r *ModifyCRMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRObtainStatusRequestParams struct {
	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 取证类型：1-网页取证 2-过程取证(暂不提供)
	ObtainType *int64 `json:"ObtainType,omitnil" name:"ObtainType"`

	// 过程取证的取证时长，单位分钟，范围0-120
	ObtainDuration *int64 `json:"ObtainDuration,omitnil" name:"ObtainDuration"`

	// 取证结果回调地址
	ObtainUrl *string `json:"ObtainUrl,omitnil" name:"ObtainUrl"`
}

type ModifyCRObtainStatusRequest struct {
	*tchttp.BaseRequest
	
	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 取证类型：1-网页取证 2-过程取证(暂不提供)
	ObtainType *int64 `json:"ObtainType,omitnil" name:"ObtainType"`

	// 过程取证的取证时长，单位分钟，范围0-120
	ObtainDuration *int64 `json:"ObtainDuration,omitnil" name:"ObtainDuration"`

	// 取证结果回调地址
	ObtainUrl *string `json:"ObtainUrl,omitnil" name:"ObtainUrl"`
}

func (r *ModifyCRObtainStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRObtainStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TortId")
	delete(f, "ObtainType")
	delete(f, "ObtainDuration")
	delete(f, "ObtainUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCRObtainStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRObtainStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCRObtainStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCRObtainStatusResponseParams `json:"Response"`
}

func (r *ModifyCRObtainStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRObtainStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRRightStatusRequestParams struct {
	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 发函结果回调地址
	RightUrl *string `json:"RightUrl,omitnil" name:"RightUrl"`
}

type ModifyCRRightStatusRequest struct {
	*tchttp.BaseRequest
	
	// 侵权ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 发函结果回调地址
	RightUrl *string `json:"RightUrl,omitnil" name:"RightUrl"`
}

func (r *ModifyCRRightStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRRightStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TortId")
	delete(f, "RightUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCRRightStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRRightStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCRRightStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCRRightStatusResponseParams `json:"Response"`
}

func (r *ModifyCRRightStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRRightStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRWhiteListRequestParams struct {
	// 该字段已废弃，白名单ID
	WhiteListId *int64 `json:"WhiteListId,omitnil" name:"WhiteListId"`

	// 该字段已废弃，平台名称
	PlatForm *string `json:"PlatForm,omitnil" name:"PlatForm"`

	// 该字段已废弃，平台站点链接
	PlatUrl *string `json:"PlatUrl,omitnil" name:"PlatUrl"`

	// 该字段已废弃，作者ID
	AuthorId *string `json:"AuthorId,omitnil" name:"AuthorId"`

	// 该字段已废弃，作品ID
	WorksId *int64 `json:"WorksId,omitnil" name:"WorksId"`

	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 白名单列表，以\n分割
	WhiteSites *string `json:"WhiteSites,omitnil" name:"WhiteSites"`
}

type ModifyCRWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 该字段已废弃，白名单ID
	WhiteListId *int64 `json:"WhiteListId,omitnil" name:"WhiteListId"`

	// 该字段已废弃，平台名称
	PlatForm *string `json:"PlatForm,omitnil" name:"PlatForm"`

	// 该字段已废弃，平台站点链接
	PlatUrl *string `json:"PlatUrl,omitnil" name:"PlatUrl"`

	// 该字段已废弃，作者ID
	AuthorId *string `json:"AuthorId,omitnil" name:"AuthorId"`

	// 该字段已废弃，作品ID
	WorksId *int64 `json:"WorksId,omitnil" name:"WorksId"`

	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 白名单列表，以\n分割
	WhiteSites *string `json:"WhiteSites,omitnil" name:"WhiteSites"`
}

func (r *ModifyCRWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhiteListId")
	delete(f, "PlatForm")
	delete(f, "PlatUrl")
	delete(f, "AuthorId")
	delete(f, "WorksId")
	delete(f, "WorkId")
	delete(f, "WhiteSites")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCRWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCRWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCRWhiteListResponseParams `json:"Response"`
}

func (r *ModifyCRWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Monitor struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 作品名称
	WorkName *string `json:"WorkName,omitnil" name:"WorkName"`

	// 作品内容类型 01-视频 02-音频 03-文本 04-图片
	WorkType *string `json:"WorkType,omitnil" name:"WorkType"`

	// 侵权平台数量
	TortPlatNum *int64 `json:"TortPlatNum,omitnil" name:"TortPlatNum"`

	// 侵权链接数量
	TortURLNum *int64 `json:"TortURLNum,omitnil" name:"TortURLNum"`

	// 监测时间
	MonitorTime *string `json:"MonitorTime,omitnil" name:"MonitorTime"`

	// 0-待监测 1-监测中 2-不监测 3-暂停监测
	MonitorStatus *int64 `json:"MonitorStatus,omitnil" name:"MonitorStatus"`

	// 作品类型
	WorkCategory *string `json:"WorkCategory,omitnil" name:"WorkCategory"`

	// 新增时间
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`

	// 监测状态说明
	MonitorNote *string `json:"MonitorNote,omitnil" name:"MonitorNote"`

	// 作品类型全部展示
	WorkCategoryAll *string `json:"WorkCategoryAll,omitnil" name:"WorkCategoryAll"`

	// 存证状态
	EvidenceStatus *int64 `json:"EvidenceStatus,omitnil" name:"EvidenceStatus"`

	// 存证状态说明
	EvidenceNote *string `json:"EvidenceNote,omitnil" name:"EvidenceNote"`

	// 侵权站点数量
	TortSiteNum *int64 `json:"TortSiteNum,omitnil" name:"TortSiteNum"`

	// 监测截止时间
	MonitorEndTime *string `json:"MonitorEndTime,omitnil" name:"MonitorEndTime"`

	// 是否自动续费
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`
}

type MonitorTort struct {
	// 侵权信息ID
	TortId *int64 `json:"TortId,omitnil" name:"TortId"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitnil" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitnil" name:"TortPlat"`

	// 侵权链接
	TortURL *string `json:"TortURL,omitnil" name:"TortURL"`

	// 侵权链接发布时间
	PubTime *string `json:"PubTime,omitnil" name:"PubTime"`

	// 作者
	Author *string `json:"Author,omitnil" name:"Author"`

	// 发现时间
	DetectTime *string `json:"DetectTime,omitnil" name:"DetectTime"`

	// 取证状态
	ObtainStatus *int64 `json:"ObtainStatus,omitnil" name:"ObtainStatus"`

	// 维权状态
	RightStatus *int64 `json:"RightStatus,omitnil" name:"RightStatus"`

	// 拦截状态
	BlockStatus *int64 `json:"BlockStatus,omitnil" name:"BlockStatus"`

	// 侵权编号
	TortNum *string `json:"TortNum,omitnil" name:"TortNum"`

	// 取证状态说明
	ObtainNote *string `json:"ObtainNote,omitnil" name:"ObtainNote"`

	// 作品标题
	WorkTitle *string `json:"WorkTitle,omitnil" name:"WorkTitle"`

	// 侵权站点
	TortSite *string `json:"TortSite,omitnil" name:"TortSite"`

	// ICP备案信息
	ICP *string `json:"ICP,omitnil" name:"ICP"`

	// 维权状态说明
	RightNote *string `json:"RightNote,omitnil" name:"RightNote"`

	// 取证类型
	ObtainType *int64 `json:"ObtainType,omitnil" name:"ObtainType"`

	// 拦截状态说明
	BlockNote *string `json:"BlockNote,omitnil" name:"BlockNote"`

	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 作品名称
	WorkName *string `json:"WorkName,omitnil" name:"WorkName"`

	// 授权书状态
	AuthStatus *int64 `json:"AuthStatus,omitnil" name:"AuthStatus"`

	// 委托书状态
	CommStatus *int64 `json:"CommStatus,omitnil" name:"CommStatus"`

	// 存证证书状态
	EvidenceStatus *int64 `json:"EvidenceStatus,omitnil" name:"EvidenceStatus"`

	// 是否著作权人
	IsProducer *int64 `json:"IsProducer,omitnil" name:"IsProducer"`

	// 是否境外网址
	IsOverseas *int64 `json:"IsOverseas,omitnil" name:"IsOverseas"`

	// ip地理位置
	IPLoc *string `json:"IPLoc,omitnil" name:"IPLoc"`
}

type ProtectURLInfo struct {
	// 保护网站ID
	ProtectURLId *int64 `json:"ProtectURLId,omitnil" name:"ProtectURLId"`

	// 保护网站
	ProtectURL *string `json:"ProtectURL,omitnil" name:"ProtectURL"`

	// 保护网站名称
	ProtectWeb *string `json:"ProtectWeb,omitnil" name:"ProtectWeb"`

	// 保护网站审核状态 1-审核中 2-审核不通过 3-审核通过
	ProtectURLStatus *int64 `json:"ProtectURLStatus,omitnil" name:"ProtectURLStatus"`

	// 网站审核不通过原因
	ProtectURLNote *string `json:"ProtectURLNote,omitnil" name:"ProtectURLNote"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type ReportFakeURLInfo struct {
	// 仿冒网址ID
	FakeURLId *int64 `json:"FakeURLId,omitnil" name:"FakeURLId"`

	// 检测时间
	DetectTime *string `json:"DetectTime,omitnil" name:"DetectTime"`

	// 保护网站
	ProtectURL *string `json:"ProtectURL,omitnil" name:"ProtectURL"`

	// 保护网站名称
	ProtectWeb *string `json:"ProtectWeb,omitnil" name:"ProtectWeb"`

	// 仿冒网址
	FakeURL *string `json:"FakeURL,omitnil" name:"FakeURL"`

	// 截图
	Snapshot *string `json:"Snapshot,omitnil" name:"Snapshot"`

	// IP地址
	IP *string `json:"IP,omitnil" name:"IP"`

	// IP地理位置
	IPLoc *string `json:"IPLoc,omitnil" name:"IPLoc"`

	// 热度
	Heat *int64 `json:"Heat,omitnil" name:"Heat"`

	// 网站状态
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 网站不处理原因
	Note *string `json:"Note,omitnil" name:"Note"`

	// 仿冒网站的企业名称
	FakeURLCompany *string `json:"FakeURLCompany,omitnil" name:"FakeURLCompany"`

	// 仿冒网站的网站性质
	FakeURLAttr *string `json:"FakeURLAttr,omitnil" name:"FakeURLAttr"`

	// 仿冒网站的网站名称
	FakeURLName *string `json:"FakeURLName,omitnil" name:"FakeURLName"`

	// 仿冒网站的备案
	FakeURLICP *string `json:"FakeURLICP,omitnil" name:"FakeURLICP"`

	// 仿冒网站创建时间
	FakeURLCreateTime *string `json:"FakeURLCreateTime,omitnil" name:"FakeURLCreateTime"`

	// 仿冒网站过期时间
	FakeURLExpireTime *string `json:"FakeURLExpireTime,omitnil" name:"FakeURLExpireTime"`

	// 协查处置时间
	BlockTime *string `json:"BlockTime,omitnil" name:"BlockTime"`
}

// Predefined struct for user
type UpdateCRWorkRequestParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 文件的扩展名，例如txt，docx
	ContentType *string `json:"ContentType,omitnil" name:"ContentType"`

	// 内容的base64编码
	Content *string `json:"Content,omitnil" name:"Content"`

	// 本次存证类型：0-不存证 1-存当前文件 2-存历史全量文件
	CertType *string `json:"CertType,omitnil" name:"CertType"`
}

type UpdateCRWorkRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 文件的扩展名，例如txt，docx
	ContentType *string `json:"ContentType,omitnil" name:"ContentType"`

	// 内容的base64编码
	Content *string `json:"Content,omitnil" name:"Content"`

	// 本次存证类型：0-不存证 1-存当前文件 2-存历史全量文件
	CertType *string `json:"CertType,omitnil" name:"CertType"`
}

func (r *UpdateCRWorkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCRWorkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "ContentType")
	delete(f, "Content")
	delete(f, "CertType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCRWorkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCRWorkResponseParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitnil" name:"WorkId"`

	// 存证ID
	EvidenceId *int64 `json:"EvidenceId,omitnil" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateCRWorkResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCRWorkResponseParams `json:"Response"`
}

func (r *UpdateCRWorkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCRWorkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}