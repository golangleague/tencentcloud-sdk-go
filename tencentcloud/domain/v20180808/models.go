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

package v20180808

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type BatchModifyDomainInfoRequestParams struct {
	// 批量修改的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 模板ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// true： 开启60天内禁止转移注册商锁定
	// false：关闭60天内禁止转移注册商锁定
	// 默认 true
	LockTransfer *bool `json:"LockTransfer,omitnil" name:"LockTransfer"`
}

type BatchModifyDomainInfoRequest struct {
	*tchttp.BaseRequest
	
	// 批量修改的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 模板ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// true： 开启60天内禁止转移注册商锁定
	// false：关闭60天内禁止转移注册商锁定
	// 默认 true
	LockTransfer *bool `json:"LockTransfer,omitnil" name:"LockTransfer"`
}

func (r *BatchModifyDomainInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyDomainInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "TemplateId")
	delete(f, "LockTransfer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyDomainInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyDomainInfoResponseParams struct {
	// 日志ID
	LogId *int64 `json:"LogId,omitnil" name:"LogId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BatchModifyDomainInfoResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyDomainInfoResponseParams `json:"Response"`
}

func (r *BatchModifyDomainInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyDomainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchStatus struct {
	// 批量任务id
	LogId *uint64 `json:"LogId,omitnil" name:"LogId"`

	// 批量任务状态  doing：进行中  success：成功  failed：失败  partial_success：部分成功
	Status *string `json:"Status,omitnil" name:"Status"`

	// 批量任务类型
	BatchAction *string `json:"BatchAction,omitnil" name:"BatchAction"`
}

type CertificateInfo struct {
	// 证件号码。
	CertificateCode *string `json:"CertificateCode,omitnil" name:"CertificateCode"`

	// 证件类型。
	// SFZ: 身份证。
	// HZ: 护照。
	// TXZ: 中国港澳居民来往内地通行证。
	// TWSFZ: 中国台湾居民来往大陆通行证。
	// GWSFZ: 外国人永久居留身份证。
	// ORG: 组织机构代码证
	// YYZZ: 工商营业执照。
	// TYDMZ: 统一社会信用代码证书。
	// BDDH: 部队代号
	// JDXKZ: 军队单位对外有偿服务许可证。
	// SYZS: 事业单位法定代表人证书。
	// GWCZDJZ: 外国企业常驻代表机构登记证。
	// STDJZ: 社会团体法定代表人登记证书。
	// ZJDJZ: 宗教活动场所登记证。
	// MBDJZ: 民办非企业单位登记证书。
	// JJDJZ: 基金会法定代表人登记证书。
	// LSXKZ: 律师事务所执业许可证。
	// GWZHDJZ: 外国在华文化中心登记证。
	// GWLYDJZ: 外国政府旅游部门常驻代表机构批准登记证。
	// SFXKZ: 司法鉴定许可证
	// GWJGZJ: 外国机构证件。
	// SHFWJGZ: 社会服务机构登记证书。
	// MBXXXKZ: 民办学校办学许可证。
	// YLJGXKZ: 医疗机构执业许可证。
	// GAJZZ: 中国港澳居住证。
	// TWJZZ: 中国台湾居住证。
	// QTTYDM: 其他-统一社会信用代码证书。
	// GZJGZY: 公证机构执业证。
	CertificateType *string `json:"CertificateType,omitnil" name:"CertificateType"`

	// 证件照片地址。
	ImgUrl *string `json:"ImgUrl,omitnil" name:"ImgUrl"`

	// 原始照片地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginImgUrl *string `json:"OriginImgUrl,omitnil" name:"OriginImgUrl"`

	// 联系人证件号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrantCertificateCode *string `json:"RegistrantCertificateCode,omitnil" name:"RegistrantCertificateCode"`

	// 联系人证件类型。
	// SFZ: 身份证。
	// HZ: 护照。
	// TXZ: 中国港澳居民来往内地通行证。
	// TWSFZ: 中国台湾居民来往大陆通行证。
	// GWSFZ: 外国人永久居留身份证。
	// ORG: 组织机构代码证
	// YYZZ: 工商营业执照。
	// TYDMZ: 统一社会信用代码证书。
	// BDDH: 部队代号
	// JDXKZ: 军队单位对外有偿服务许可证。
	// SYZS: 事业单位法定代表人证书。
	// GWCZDJZ: 外国企业常驻代表机构登记证。
	// STDJZ: 社会团体法定代表人登记证书。
	// ZJDJZ: 宗教活动场所登记证。
	// MBDJZ: 民办非企业单位登记证书。
	// JJDJZ: 基金会法定代表人登记证书。
	// LSXKZ: 律师事务所执业许可证。
	// GWZHDJZ: 外国在华文化中心登记证。
	// GWLYDJZ: 外国政府旅游部门常驻代表机构批准登记证。
	// SFXKZ: 司法鉴定许可证
	// GWJGZJ: 外国机构证件。
	// SHFWJGZ: 社会服务机构登记证书。
	// MBXXXKZ: 民办学校办学许可证。
	// YLJGXKZ: 医疗机构执业许可证。
	// GAJZZ: 中国港澳居住证。
	// TWJZZ: 中国台湾居住证。
	// QTTYDM: 其他-统一社会信用代码证书。
	// GZJGZY: 公证机构执业证。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrantCertificateType *string `json:"RegistrantCertificateType,omitnil" name:"RegistrantCertificateType"`

	// 联系人证件照片地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrantImgUrl *string `json:"RegistrantImgUrl,omitnil" name:"RegistrantImgUrl"`
}

// Predefined struct for user
type CheckBatchStatusRequestParams struct {
	// 操作日志 ID数组，最多 200 个
	LogIds []*uint64 `json:"LogIds,omitnil" name:"LogIds"`
}

type CheckBatchStatusRequest struct {
	*tchttp.BaseRequest
	
	// 操作日志 ID数组，最多 200 个
	LogIds []*uint64 `json:"LogIds,omitnil" name:"LogIds"`
}

func (r *CheckBatchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBatchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckBatchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBatchStatusResponseParams struct {
	// 批量任务状态集
	StatusSet []*BatchStatus `json:"StatusSet,omitnil" name:"StatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckBatchStatusResponse struct {
	*tchttp.BaseResponse
	Response *CheckBatchStatusResponseParams `json:"Response"`
}

func (r *CheckBatchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBatchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDomainRequestParams struct {
	// 所查询域名名称
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 年限。该参数为空时无法查询溢价词域名
	Period *string `json:"Period,omitnil" name:"Period"`
}

type CheckDomainRequest struct {
	*tchttp.BaseRequest
	
	// 所查询域名名称
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 年限。该参数为空时无法查询溢价词域名
	Period *string `json:"Period,omitnil" name:"Period"`
}

func (r *CheckDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDomainResponseParams struct {
	// 所查询域名名称
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 是否能够注册
	Available *bool `json:"Available,omitnil" name:"Available"`

	// 不能注册原因
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 是否是溢价词
	Premium *bool `json:"Premium,omitnil" name:"Premium"`

	// 域名价格
	Price *uint64 `json:"Price,omitnil" name:"Price"`

	// 是否是敏感词
	BlackWord *bool `json:"BlackWord,omitnil" name:"BlackWord"`

	// 溢价词描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 溢价词的续费价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeRenew *uint64 `json:"FeeRenew,omitnil" name:"FeeRenew"`

	// 域名真实价格, 溢价词时价格跟年限有关，非溢价词时价格为1年的价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealPrice *uint64 `json:"RealPrice,omitnil" name:"RealPrice"`

	// 溢价词的转入价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeTransfer *uint64 `json:"FeeTransfer,omitnil" name:"FeeTransfer"`

	// 溢价词的赎回价格
	FeeRestore *uint64 `json:"FeeRestore,omitnil" name:"FeeRestore"`

	// 检测年限
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// 是否支持北京备案  true 支持  false 不支持
	RecordSupport *bool `json:"RecordSupport,omitnil" name:"RecordSupport"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckDomainResponse struct {
	*tchttp.BaseResponse
	Response *CheckDomainResponseParams `json:"Response"`
}

func (r *CheckDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContactInfo struct {
	// 注册人（中文）
	OrganizationNameCN *string `json:"OrganizationNameCN,omitnil" name:"OrganizationNameCN"`

	// 注册人（英文）
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 联系人（中文）
	RegistrantNameCN *string `json:"RegistrantNameCN,omitnil" name:"RegistrantNameCN"`

	// 联系人（英文）
	RegistrantName *string `json:"RegistrantName,omitnil" name:"RegistrantName"`

	// 省份（中文）
	ProvinceCN *string `json:"ProvinceCN,omitnil" name:"ProvinceCN"`

	// 城市（中文）
	CityCN *string `json:"CityCN,omitnil" name:"CityCN"`

	// 街道（中文）
	StreetCN *string `json:"StreetCN,omitnil" name:"StreetCN"`

	// 街道（英文）
	Street *string `json:"Street,omitnil" name:"Street"`

	// 国家（中文）
	CountryCN *string `json:"CountryCN,omitnil" name:"CountryCN"`

	// 联系人手机号
	Telephone *string `json:"Telephone,omitnil" name:"Telephone"`

	// 联系人邮箱
	Email *string `json:"Email,omitnil" name:"Email"`

	// 邮编
	ZipCode *string `json:"ZipCode,omitnil" name:"ZipCode"`

	// 用户类型 E:组织， I:个人
	RegistrantType *string `json:"RegistrantType,omitnil" name:"RegistrantType"`

	// 省份（英文）。作为入参时可以不填
	Province *string `json:"Province,omitnil" name:"Province"`

	// 城市（英文）。作为入参时可以不填
	City *string `json:"City,omitnil" name:"City"`

	// 国家（英文）。作为入参时可以不填
	Country *string `json:"Country,omitnil" name:"Country"`
}

// Predefined struct for user
type CreateDomainBatchRequestParams struct {
	// 模板ID。详情请查看：[获取模板列表](https://cloud.tencent.com/document/product/242/48940)
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 购买域名的年限，可选值：[1-10]
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 批量购买的域名,最多为4000个
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 付费模式 0手动在线付费，1使用余额付费，2使用特惠包
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 自动续费开关。有两个可选值：
	// 0 表示关闭，不自动续费（默认值）
	// 1 表示开启，将自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 使用的特惠包ID，PayMode为2时必填
	PackageResourceId *string `json:"PackageResourceId,omitnil" name:"PackageResourceId"`

	// 是否开启更新锁：0=默认不开启，1=开启
	UpdateProhibition *int64 `json:"UpdateProhibition,omitnil" name:"UpdateProhibition"`

	// 是否开启转移锁：0=默认不开启，1=开启
	TransferProhibition *int64 `json:"TransferProhibition,omitnil" name:"TransferProhibition"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil" name:"ActivityId"`
}

type CreateDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID。详情请查看：[获取模板列表](https://cloud.tencent.com/document/product/242/48940)
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 购买域名的年限，可选值：[1-10]
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 批量购买的域名,最多为4000个
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 付费模式 0手动在线付费，1使用余额付费，2使用特惠包
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 自动续费开关。有两个可选值：
	// 0 表示关闭，不自动续费（默认值）
	// 1 表示开启，将自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 使用的特惠包ID，PayMode为2时必填
	PackageResourceId *string `json:"PackageResourceId,omitnil" name:"PackageResourceId"`

	// 是否开启更新锁：0=默认不开启，1=开启
	UpdateProhibition *int64 `json:"UpdateProhibition,omitnil" name:"UpdateProhibition"`

	// 是否开启转移锁：0=默认不开启，1=开启
	TransferProhibition *int64 `json:"TransferProhibition,omitnil" name:"TransferProhibition"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil" name:"ActivityId"`
}

func (r *CreateDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Period")
	delete(f, "Domains")
	delete(f, "PayMode")
	delete(f, "AutoRenewFlag")
	delete(f, "PackageResourceId")
	delete(f, "UpdateProhibition")
	delete(f, "TransferProhibition")
	delete(f, "ChannelFrom")
	delete(f, "OrderFrom")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainBatchResponseParams struct {
	// 批量日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogId *int64 `json:"LogId,omitnil" name:"LogId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainBatchResponseParams `json:"Response"`
}

func (r *CreateDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainRedemptionRequestParams struct {
	// 域名 ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`
}

type CreateDomainRedemptionRequest struct {
	*tchttp.BaseRequest
	
	// 域名 ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`
}

func (r *CreateDomainRedemptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainRedemptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainRedemptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainRedemptionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDomainRedemptionResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainRedemptionResponseParams `json:"Response"`
}

func (r *CreateDomainRedemptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainRedemptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePhoneEmailRequestParams struct {
	// 手机号或者邮箱
	Code *string `json:"Code,omitnil" name:"Code"`

	// 1：手机   2：邮箱
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 验证码
	VerifyCode *string `json:"VerifyCode,omitnil" name:"VerifyCode"`
}

type CreatePhoneEmailRequest struct {
	*tchttp.BaseRequest
	
	// 手机号或者邮箱
	Code *string `json:"Code,omitnil" name:"Code"`

	// 1：手机   2：邮箱
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 验证码
	VerifyCode *string `json:"VerifyCode,omitnil" name:"VerifyCode"`
}

func (r *CreatePhoneEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePhoneEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Code")
	delete(f, "Type")
	delete(f, "VerifyCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePhoneEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePhoneEmailResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePhoneEmailResponse struct {
	*tchttp.BaseResponse
	Response *CreatePhoneEmailResponseParams `json:"Response"`
}

func (r *CreatePhoneEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePhoneEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTemplateRequestParams struct {
	// 联系人信息
	ContactInfo *ContactInfo `json:"ContactInfo,omitnil" name:"ContactInfo"`

	// 证件信息
	CertificateInfo *CertificateInfo `json:"CertificateInfo,omitnil" name:"CertificateInfo"`
}

type CreateTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 联系人信息
	ContactInfo *ContactInfo `json:"ContactInfo,omitnil" name:"ContactInfo"`

	// 证件信息
	CertificateInfo *CertificateInfo `json:"CertificateInfo,omitnil" name:"CertificateInfo"`
}

func (r *CreateTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContactInfo")
	delete(f, "CertificateInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTemplateResponseParams struct {
	// 模板信息
	Template *TemplateInfo `json:"Template,omitnil" name:"Template"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateTemplateResponseParams `json:"Response"`
}

func (r *CreateTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePhoneEmailRequestParams struct {
	// 手机或者邮箱
	Code *string `json:"Code,omitnil" name:"Code"`

	// 1：手机  2：邮箱
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

type DeletePhoneEmailRequest struct {
	*tchttp.BaseRequest
	
	// 手机或者邮箱
	Code *string `json:"Code,omitnil" name:"Code"`

	// 1：手机  2：邮箱
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

func (r *DeletePhoneEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePhoneEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Code")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePhoneEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePhoneEmailResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeletePhoneEmailResponse struct {
	*tchttp.BaseResponse
	Response *DeletePhoneEmailResponseParams `json:"Response"`
}

func (r *DeletePhoneEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePhoneEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTemplateRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type DeleteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *DeleteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTemplateResponseParams `json:"Response"`
}

func (r *DeleteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOperationLogDetailsRequestParams struct {
	// 日志ID。
	LogId *int64 `json:"LogId,omitnil" name:"LogId"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为200。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeBatchOperationLogDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 日志ID。
	LogId *int64 `json:"LogId,omitnil" name:"LogId"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为200。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeBatchOperationLogDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOperationLogDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchOperationLogDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOperationLogDetailsResponseParams struct {
	// 总数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 日志详情列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainBatchDetailSet []*DomainBatchDetailSet `json:"DomainBatchDetailSet,omitnil" name:"DomainBatchDetailSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBatchOperationLogDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchOperationLogDetailsResponseParams `json:"Response"`
}

func (r *DescribeBatchOperationLogDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOperationLogDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOperationLogsRequestParams struct {
	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为200。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeBatchOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为200。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeBatchOperationLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOperationLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchOperationLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOperationLogsResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainBatchLogSet []*DomainBatchLogSet `json:"DomainBatchLogSet,omitnil" name:"DomainBatchLogSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBatchOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchOperationLogsResponseParams `json:"Response"`
}

func (r *DescribeBatchOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainBaseInfoRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DescribeDomainBaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DescribeDomainBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainBaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainBaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainBaseInfoResponseParams struct {
	// 域名信息
	DomainInfo *DomainBaseInfo `json:"DomainInfo,omitnil" name:"DomainInfo"`

	// 用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainBaseInfoResponseParams `json:"Response"`
}

func (r *DescribeDomainBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainNameListRequestParams struct {
	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，取值范围[1,100]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeDomainNameListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，取值范围[1,100]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeDomainNameListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainNameListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainNameListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainNameListResponseParams struct {
	// 域名信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainSet []*DomainList `json:"DomainSet,omitnil" name:"DomainSet"`

	// 域名总数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainNameListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainNameListResponseParams `json:"Response"`
}

func (r *DescribeDomainNameListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainPriceListRequestParams struct {
	// 查询价格的后缀列表。默认则为全部后缀
	TldList []*string `json:"TldList,omitnil" name:"TldList"`

	// 查询购买的年份，默认会列出所有年份的价格
	Year []*int64 `json:"Year,omitnil" name:"Year"`

	// 域名的购买类型：new  新购，renew 续费，redem 赎回，tran 转入
	Operation []*string `json:"Operation,omitnil" name:"Operation"`
}

type DescribeDomainPriceListRequest struct {
	*tchttp.BaseRequest
	
	// 查询价格的后缀列表。默认则为全部后缀
	TldList []*string `json:"TldList,omitnil" name:"TldList"`

	// 查询购买的年份，默认会列出所有年份的价格
	Year []*int64 `json:"Year,omitnil" name:"Year"`

	// 域名的购买类型：new  新购，renew 续费，redem 赎回，tran 转入
	Operation []*string `json:"Operation,omitnil" name:"Operation"`
}

func (r *DescribeDomainPriceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainPriceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TldList")
	delete(f, "Year")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainPriceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainPriceListResponseParams struct {
	// 域名价格列表
	PriceList []*PriceInfo `json:"PriceList,omitnil" name:"PriceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainPriceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainPriceListResponseParams `json:"Response"`
}

func (r *DescribeDomainPriceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainPriceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainSimpleInfoRequestParams struct {
	// 域名
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`
}

type DescribeDomainSimpleInfoRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`
}

func (r *DescribeDomainSimpleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainSimpleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainSimpleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainSimpleInfoResponseParams struct {
	// 域名信息
	DomainInfo *DomainSimpleInfo `json:"DomainInfo,omitnil" name:"DomainInfo"`

	// 账号ID
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainSimpleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainSimpleInfoResponseParams `json:"Response"`
}

func (r *DescribeDomainSimpleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainSimpleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePhoneEmailListRequestParams struct {
	// 0：所有类型  1：手机  2：邮箱，默认0
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，取值范围[1,200]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 手机或者邮箱，用于精确搜索
	Code *string `json:"Code,omitnil" name:"Code"`
}

type DescribePhoneEmailListRequest struct {
	*tchttp.BaseRequest
	
	// 0：所有类型  1：手机  2：邮箱，默认0
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，取值范围[1,200]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 手机或者邮箱，用于精确搜索
	Code *string `json:"Code,omitnil" name:"Code"`
}

func (r *DescribePhoneEmailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePhoneEmailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Code")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePhoneEmailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePhoneEmailListResponseParams struct {
	// 手机或者邮箱列表
	PhoneEmailList []*PhoneEmailData `json:"PhoneEmailList,omitnil" name:"PhoneEmailList"`

	// 总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePhoneEmailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePhoneEmailListResponseParams `json:"Response"`
}

func (r *DescribePhoneEmailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePhoneEmailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateListRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 用户注册类型，默认:all , 个人：I ,企业: E
	Type *string `json:"Type,omitnil" name:"Type"`

	// 认证状态：未实名审核:NotUpload, 实名审核中:InAudit，已实名审核:Approved，实名审核失败:Reject，更新手机邮箱:NotVerified。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 关键字，用于域名所有者筛选
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`
}

type DescribeTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 用户注册类型，默认:all , 个人：I ,企业: E
	Type *string `json:"Type,omitnil" name:"Type"`

	// 认证状态：未实名审核:NotUpload, 实名审核中:InAudit，已实名审核:Approved，实名审核失败:Reject，更新手机邮箱:NotVerified。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 关键字，用于域名所有者筛选
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`
}

func (r *DescribeTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	delete(f, "Status")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateListResponseParams struct {
	// 模板数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 模板详细信息列表。
	TemplateSet []*TemplateInfo `json:"TemplateSet,omitnil" name:"TemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateListResponseParams `json:"Response"`
}

func (r *DescribeTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type DescribeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *DescribeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateResponseParams struct {
	// 模板信息
	Template *TemplateInfo `json:"Template,omitnil" name:"Template"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateResponseParams `json:"Response"`
}

func (r *DescribeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainBaseInfo struct {
	// 域名资源ID。
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 域名名称。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 域名实名认证状态。
	// NotUpload：未实名认证
	// InAudit：实名审核中
	// Approved：实名审核通过
	// Reject：实名审核失败
	// NoAudit: 无需实名认证
	RealNameAuditStatus *string `json:"RealNameAuditStatus,omitnil" name:"RealNameAuditStatus"`

	// 域名实名认证不通过原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealNameAuditUnpassReason *string `json:"RealNameAuditUnpassReason,omitnil" name:"RealNameAuditUnpassReason"`

	// 域名命名审核状态。
	// NotAudit：命名审核未上传
	// Pending：命名审核待上传
	// Auditing：域名命名审核中
	// Approved：域名命名审核通过
	// Rejected：域名命名审核拒绝
	DomainNameAuditStatus *string `json:"DomainNameAuditStatus,omitnil" name:"DomainNameAuditStatus"`

	// 域名命名审核不通过原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainNameAuditUnpassReason *string `json:"DomainNameAuditUnpassReason,omitnil" name:"DomainNameAuditUnpassReason"`

	// 注册时间。
	CreationDate *string `json:"CreationDate,omitnil" name:"CreationDate"`

	// 到期时间
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// 域名状态。
	// ok：正常
	// serverHold：注册局暂停解析 
	// clientHold：注册商暂停解析
	// pendingTransfer：转移中
	// renewingPeriod：续费期
	// redemptionPeriod：偿还期
	// pendingDelete：删除期
	// serverTransferProhibited：注册局禁止转移
	// serverUpdateProhibited：注册局禁止更新
	// serverDeleteProhibited：注册局禁止删除
	// clientTransferProhibited：注册商禁止转移
	// clientUpdateProhibited：注册商禁止更新
	// clientDeleteProhibited：注册商禁止删除
	// serverRenewProhibited: 注册局禁止续费
	// clientRenewProhobited: 注册商禁止续费
	DomainStatus []*string `json:"DomainStatus,omitnil" name:"DomainStatus"`

	// 域名购买状态。
	// ok：正常
	// RegisterPending：待注册
	// RegisterDoing：注册中
	// RegisterFailed：注册失败
	// AboutToExpire: 即将过期
	// RenewPending：已进入续费期，需要进行续费
	// RenewDoing：续费中
	// RedemptionPending：已进入赎回期，需要进行续费
	// RedemptionDoing：赎回中
	// TransferPending：待转入中
	// TransferTransing：转入中
	// TransferFailed：转入失败
	BuyStatus *string `json:"BuyStatus,omitnil" name:"BuyStatus"`

	// 注册商类型
	// epp: DNSPod, Inc.（烟台帝思普网络科技有限公司）
	// qcloud: Tencent Cloud Computing (Beijing) Limited Liability Company（腾讯云计算（北京）有限责任公司）
	// yunxun: Guangzhou Yunxun Information Technology Co., Ltd.（广州云讯信息科技有限公司）
	// xinnet: Xin Net Technology Corporation（北京新网数码信息技术有限公司）
	RegistrarType *string `json:"RegistrarType,omitnil" name:"RegistrarType"`

	// 域名绑定的ns
	NameServer []*string `json:"NameServer,omitnil" name:"NameServer"`

	// true：开启锁定
	// false：关闭锁定
	LockTransfer *bool `json:"LockTransfer,omitnil" name:"LockTransfer"`

	// 锁定结束时间
	LockEndTime *string `json:"LockEndTime,omitnil" name:"LockEndTime"`
}

type DomainBatchDetailSet struct {
	// 详情ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 执行状态：
	// doing 执行中。
	// failed 操作失败。
	// success  操作成功。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil" name:"UpdatedOn"`
}

type DomainBatchLogSet struct {
	// 日志ID
	LogId *int64 `json:"LogId,omitnil" name:"LogId"`

	// 数量
	Number *int64 `json:"Number,omitnil" name:"Number"`

	// 执行状态：
	// doing 执行中。
	// done 执行完成。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 提交时间
	CreatedOn *string `json:"CreatedOn,omitnil" name:"CreatedOn"`
}

type DomainList struct {
	// 是否是溢价域名：
	// ture 是    
	// false 不是
	IsPremium *bool `json:"IsPremium,omitnil" name:"IsPremium"`

	// 域名资源ID。
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 域名名称。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 是否已设置自动续费 。
	// 0：未设置 
	// 1：已设置
	// 2：设置后，关闭
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// 注册时间。
	CreationDate *string `json:"CreationDate,omitnil" name:"CreationDate"`

	// 到期时间。
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// 域名后缀
	Tld *string `json:"Tld,omitnil" name:"Tld"`

	// 编码后的后缀（中文会进行编码）
	CodeTld *string `json:"CodeTld,omitnil" name:"CodeTld"`

	// 域名购买状态。
	// ok：正常
	// AboutToExpire: 即将到期
	// RegisterPending：注册中
	// RegisterDoing：注册中
	// RegisterFailed：注册失败
	// RenewPending：续费期
	// RenewDoing：续费中
	// RedemptionPending：赎回期
	// RedemptionDoing：赎回中
	// TransferPending：转入中
	// TransferTransing：转入中
	// TransferFailed：转入失败
	BuyStatus *string `json:"BuyStatus,omitnil" name:"BuyStatus"`
}

type DomainSimpleInfo struct {
	// 域名资源ID。
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 域名名称。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 域名实名认证状态。
	// NotUpload：未实名认证
	// InAudit：实名审核中
	// Approved：实名审核通过
	// Reject：实名审核失败
	// NoAudit: 无需实名认证
	RealNameAuditStatus *string `json:"RealNameAuditStatus,omitnil" name:"RealNameAuditStatus"`

	// 域名实名认证不通过原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealNameAuditUnpassReason *string `json:"RealNameAuditUnpassReason,omitnil" name:"RealNameAuditUnpassReason"`

	// 域名命名审核状态。
	// NotAudit：命名审核未上传
	// Pending：命名审核待上传
	// Auditing：域名命名审核中
	// Approved：域名命名审核通过
	// Rejected：域名命名审核拒绝
	DomainNameAuditStatus *string `json:"DomainNameAuditStatus,omitnil" name:"DomainNameAuditStatus"`

	// 域名命名审核不通过原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainNameAuditUnpassReason *string `json:"DomainNameAuditUnpassReason,omitnil" name:"DomainNameAuditUnpassReason"`

	// 注册时间。
	CreationDate *string `json:"CreationDate,omitnil" name:"CreationDate"`

	// 到期时间
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// 域名状态。
	// ok：正常
	// serverHold：注册局暂停解析 
	// clientHold：注册商暂停解析
	// pendingTransfer：转移中
	// renewingPeriod：续费期
	// redemptionPeriod：偿还期
	// pendingDelete：删除期
	// serverTransferProhibited：注册局禁止转移
	// serverUpdateProhibited：注册局禁止更新
	// serverDeleteProhibited：注册局禁止删除
	// clientTransferProhibited：注册商禁止转移
	// clientUpdateProhibited：注册商禁止更新
	// clientDeleteProhibited：注册商禁止删除
	// serverRenewProhibited: 注册局禁止续费
	// clientRenewProhobited: 注册商禁止续费
	DomainStatus []*string `json:"DomainStatus,omitnil" name:"DomainStatus"`

	// 域名购买状态。
	// ok：正常
	// RegisterPending：待注册
	// RegisterDoing：注册中
	// RegisterFailed：注册失败
	// AboutToExpire: 即将过期
	// RenewPending：已进入续费期，需要进行续费
	// RenewDoing：续费中
	// RedemptionPending：已进入赎回期，需要进行续费
	// RedemptionDoing：赎回中
	// TransferPending：待转入中
	// TransferTransing：转入中
	// TransferFailed：转入失败
	BuyStatus *string `json:"BuyStatus,omitnil" name:"BuyStatus"`

	// 注册商类型
	// epp: DNSPod, Inc.（烟台帝思普网络科技有限公司）
	// qcloud: Tencent Cloud Computing (Beijing) Limited Liability Company（腾讯云计算（北京）有限责任公司）
	// yunxun: Guangzhou Yunxun Information Technology Co., Ltd.（广州云讯信息科技有限公司）
	// xinnet: Xin Net Technology Corporation（北京新网数码信息技术有限公司）
	RegistrarType *string `json:"RegistrarType,omitnil" name:"RegistrarType"`

	// 域名绑定的ns
	NameServer []*string `json:"NameServer,omitnil" name:"NameServer"`

	// true：开启锁定
	// false：关闭锁定
	LockTransfer *bool `json:"LockTransfer,omitnil" name:"LockTransfer"`

	// 锁定结束时间
	LockEndTime *string `json:"LockEndTime,omitnil" name:"LockEndTime"`

	// 认证类型：I=个人，E=企业
	RegistrantType *string `json:"RegistrantType,omitnil" name:"RegistrantType"`

	// 域名所有者，中文
	OrganizationNameCN *string `json:"OrganizationNameCN,omitnil" name:"OrganizationNameCN"`

	// 域名所有者，英文
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 域名联系人，中文
	RegistrantNameCN *string `json:"RegistrantNameCN,omitnil" name:"RegistrantNameCN"`

	// 域名联系人，英文
	RegistrantName *string `json:"RegistrantName,omitnil" name:"RegistrantName"`
}

// Predefined struct for user
type ModifyDomainDNSBatchRequestParams struct {
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 域名DNS 数组。
	Dns []*string `json:"Dns,omitnil" name:"Dns"`
}

type ModifyDomainDNSBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 域名DNS 数组。
	Dns []*string `json:"Dns,omitnil" name:"Dns"`
}

func (r *ModifyDomainDNSBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainDNSBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Dns")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainDNSBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainDNSBatchResponseParams struct {
	// 日志ID。
	LogId *uint64 `json:"LogId,omitnil" name:"LogId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDomainDNSBatchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainDNSBatchResponseParams `json:"Response"`
}

func (r *ModifyDomainDNSBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainDNSBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainOwnerBatchRequestParams struct {
	// 要过户的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 转入账户的uin。
	NewOwnerUin *string `json:"NewOwnerUin,omitnil" name:"NewOwnerUin"`

	// 是否同时转移对应的 DNS 解析域名，默认false
	TransferDns *bool `json:"TransferDns,omitnil" name:"TransferDns"`

	// 转入账户的appid。
	NewOwnerAppId *string `json:"NewOwnerAppId,omitnil" name:"NewOwnerAppId"`
}

type ModifyDomainOwnerBatchRequest struct {
	*tchttp.BaseRequest
	
	// 要过户的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 转入账户的uin。
	NewOwnerUin *string `json:"NewOwnerUin,omitnil" name:"NewOwnerUin"`

	// 是否同时转移对应的 DNS 解析域名，默认false
	TransferDns *bool `json:"TransferDns,omitnil" name:"TransferDns"`

	// 转入账户的appid。
	NewOwnerAppId *string `json:"NewOwnerAppId,omitnil" name:"NewOwnerAppId"`
}

func (r *ModifyDomainOwnerBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainOwnerBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "NewOwnerUin")
	delete(f, "TransferDns")
	delete(f, "NewOwnerAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainOwnerBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainOwnerBatchResponseParams struct {
	// 日志id
	LogId *uint64 `json:"LogId,omitnil" name:"LogId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDomainOwnerBatchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainOwnerBatchResponseParams `json:"Response"`
}

func (r *ModifyDomainOwnerBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainOwnerBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PhoneEmailData struct {
	// 手机号或者邮箱
	Code *string `json:"Code,omitnil" name:"Code"`

	// 1：手机  2：邮箱
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 1=控制台校验，2=第三方校验
	CheckStatus *int64 `json:"CheckStatus,omitnil" name:"CheckStatus"`
}

type PriceInfo struct {
	// 域名后缀，例如.com
	Tld *string `json:"Tld,omitnil" name:"Tld"`

	// 购买年限，范围[1-10]
	Year *uint64 `json:"Year,omitnil" name:"Year"`

	// 域名原价
	Price *uint64 `json:"Price,omitnil" name:"Price"`

	// 域名现价
	RealPrice *uint64 `json:"RealPrice,omitnil" name:"RealPrice"`

	// 商品的购买类型，新购，续费，赎回，转入，续费并转入
	Operation *string `json:"Operation,omitnil" name:"Operation"`
}

// Predefined struct for user
type RenewDomainBatchRequestParams struct {
	// 域名续费的年限。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 批量续费的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 付费模式 0手动在线付费，1使用余额付费，2使用特惠包。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 自动续费开关。有三个可选值：
	// 0 表示关闭，不自动续费
	// 1 表示开启，将自动续费
	// 2 表示不处理，保留域名原有状态（默认值）
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 特惠包ID
	PackageResourceId *string `json:"PackageResourceId,omitnil" name:"PackageResourceId"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil" name:"ActivityId"`
}

type RenewDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// 域名续费的年限。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 批量续费的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 付费模式 0手动在线付费，1使用余额付费，2使用特惠包。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 自动续费开关。有三个可选值：
	// 0 表示关闭，不自动续费
	// 1 表示开启，将自动续费
	// 2 表示不处理，保留域名原有状态（默认值）
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 特惠包ID
	PackageResourceId *string `json:"PackageResourceId,omitnil" name:"PackageResourceId"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil" name:"ActivityId"`
}

func (r *RenewDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "Domains")
	delete(f, "PayMode")
	delete(f, "AutoRenewFlag")
	delete(f, "PackageResourceId")
	delete(f, "ChannelFrom")
	delete(f, "OrderFrom")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDomainBatchResponseParams struct {
	// 操作日志ID。
	LogId *int64 `json:"LogId,omitnil" name:"LogId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RenewDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *RenewDomainBatchResponseParams `json:"Response"`
}

func (r *RenewDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendPhoneEmailCodeRequestParams struct {
	// 手机或者邮箱号。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 1：手机  2：邮箱。
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

type SendPhoneEmailCodeRequest struct {
	*tchttp.BaseRequest
	
	// 手机或者邮箱号。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 1：手机  2：邮箱。
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

func (r *SendPhoneEmailCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendPhoneEmailCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Code")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendPhoneEmailCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendPhoneEmailCodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SendPhoneEmailCodeResponse struct {
	*tchttp.BaseResponse
	Response *SendPhoneEmailCodeResponseParams `json:"Response"`
}

func (r *SendPhoneEmailCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendPhoneEmailCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDomainAutoRenewRequestParams struct {
	// 域名ID 例如：domain-123abc
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// AutoRenew 有三个可选值：
	//  0：不设置自动续费
	// 1：设置自动续费
	// 2：设置到期后不续费
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`
}

type SetDomainAutoRenewRequest struct {
	*tchttp.BaseRequest
	
	// 域名ID 例如：domain-123abc
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// AutoRenew 有三个可选值：
	//  0：不设置自动续费
	// 1：设置自动续费
	// 2：设置到期后不续费
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`
}

func (r *SetDomainAutoRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDomainAutoRenewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetDomainAutoRenewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDomainAutoRenewResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetDomainAutoRenewResponse struct {
	*tchttp.BaseResponse
	Response *SetDomainAutoRenewResponseParams `json:"Response"`
}

func (r *SetDomainAutoRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDomainAutoRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateInfo struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 认证状态：未实名认证:NotUpload, 实名审核中:InAudit，已实名认证:Approved，实名审核失败:Reject
	AuditStatus *string `json:"AuditStatus,omitnil" name:"AuditStatus"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil" name:"UpdatedOn"`

	// 用户UIN
	UserUin *string `json:"UserUin,omitnil" name:"UserUin"`

	// 是否是默认模板: 是:yes，否:no
	IsDefault *string `json:"IsDefault,omitnil" name:"IsDefault"`

	// 认证失败原因
	AuditReason *string `json:"AuditReason,omitnil" name:"AuditReason"`

	// 认证信息
	CertificateInfo *CertificateInfo `json:"CertificateInfo,omitnil" name:"CertificateInfo"`

	// 联系人信息
	ContactInfo *ContactInfo `json:"ContactInfo,omitnil" name:"ContactInfo"`

	// 模板是否符合规范， 1是 0 否
	IsValidTemplate *int64 `json:"IsValidTemplate,omitnil" name:"IsValidTemplate"`

	// 不符合规范原因
	InvalidReason *string `json:"InvalidReason,omitnil" name:"InvalidReason"`

	// 是包含黑名单手机或邮箱
	IsBlack *bool `json:"IsBlack,omitnil" name:"IsBlack"`
}

// Predefined struct for user
type TransferInDomainBatchRequestParams struct {
	// 转入的域名名称数组。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 域名转移码数组。
	PassWords []*string `json:"PassWords,omitnil" name:"PassWords"`

	// 模板ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 付费模式 0手动在线付费，1使用余额付费。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 自动续费开关。有两个可选值：
	// 0 表示关闭，不自动续费（默认值）
	// 1 表示开启，将自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// true： 开启60天内禁止转移注册商锁定
	// false：关闭60天内禁止转移注册商锁定
	// 默认 true
	LockTransfer *bool `json:"LockTransfer,omitnil" name:"LockTransfer"`

	// 是否开启更新锁：0=默认不开启，1=开启
	UpdateProhibition *int64 `json:"UpdateProhibition,omitnil" name:"UpdateProhibition"`

	// 是否开启转移锁：0=默认不开启，1=开启
	TransferProhibition *int64 `json:"TransferProhibition,omitnil" name:"TransferProhibition"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil" name:"ActivityId"`
}

type TransferInDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// 转入的域名名称数组。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 域名转移码数组。
	PassWords []*string `json:"PassWords,omitnil" name:"PassWords"`

	// 模板ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 付费模式 0手动在线付费，1使用余额付费。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 自动续费开关。有两个可选值：
	// 0 表示关闭，不自动续费（默认值）
	// 1 表示开启，将自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// true： 开启60天内禁止转移注册商锁定
	// false：关闭60天内禁止转移注册商锁定
	// 默认 true
	LockTransfer *bool `json:"LockTransfer,omitnil" name:"LockTransfer"`

	// 是否开启更新锁：0=默认不开启，1=开启
	UpdateProhibition *int64 `json:"UpdateProhibition,omitnil" name:"UpdateProhibition"`

	// 是否开启转移锁：0=默认不开启，1=开启
	TransferProhibition *int64 `json:"TransferProhibition,omitnil" name:"TransferProhibition"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil" name:"ActivityId"`
}

func (r *TransferInDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferInDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "PassWords")
	delete(f, "TemplateId")
	delete(f, "PayMode")
	delete(f, "AutoRenewFlag")
	delete(f, "LockTransfer")
	delete(f, "UpdateProhibition")
	delete(f, "TransferProhibition")
	delete(f, "ChannelFrom")
	delete(f, "OrderFrom")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferInDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferInDomainBatchResponseParams struct {
	// 日志ID
	LogId *uint64 `json:"LogId,omitnil" name:"LogId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TransferInDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *TransferInDomainBatchResponseParams `json:"Response"`
}

func (r *TransferInDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferInDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferProhibitionBatchRequestParams struct {
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 是否开启禁止域名转移。
	// True: 开启禁止域名转移状态。
	// False：关闭禁止域名转移状态。
	Status *bool `json:"Status,omitnil" name:"Status"`
}

type TransferProhibitionBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 是否开启禁止域名转移。
	// True: 开启禁止域名转移状态。
	// False：关闭禁止域名转移状态。
	Status *bool `json:"Status,omitnil" name:"Status"`
}

func (r *TransferProhibitionBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferProhibitionBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferProhibitionBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferProhibitionBatchResponseParams struct {
	// 日志ID
	LogId *uint64 `json:"LogId,omitnil" name:"LogId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TransferProhibitionBatchResponse struct {
	*tchttp.BaseResponse
	Response *TransferProhibitionBatchResponseParams `json:"Response"`
}

func (r *TransferProhibitionBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferProhibitionBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProhibitionBatchRequestParams struct {
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 是否开启禁止域名更新。
	// True:开启禁止域名更新状态。
	// False：关闭禁止域名更新状态。
	Status *bool `json:"Status,omitnil" name:"Status"`
}

type UpdateProhibitionBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 是否开启禁止域名更新。
	// True:开启禁止域名更新状态。
	// False：关闭禁止域名更新状态。
	Status *bool `json:"Status,omitnil" name:"Status"`
}

func (r *UpdateProhibitionBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProhibitionBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProhibitionBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProhibitionBatchResponseParams struct {
	// 日志ID
	LogId *uint64 `json:"LogId,omitnil" name:"LogId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateProhibitionBatchResponse struct {
	*tchttp.BaseResponse
	Response *UpdateProhibitionBatchResponseParams `json:"Response"`
}

func (r *UpdateProhibitionBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProhibitionBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadImageRequestParams struct {
	// 资质照片，照片的base64编码。
	ImageFile *string `json:"ImageFile,omitnil" name:"ImageFile"`
}

type UploadImageRequest struct {
	*tchttp.BaseRequest
	
	// 资质照片，照片的base64编码。
	ImageFile *string `json:"ImageFile,omitnil" name:"ImageFile"`
}

func (r *UploadImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageFile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadImageResponseParams struct {
	// 资质照片地址。
	AccessUrl *string `json:"AccessUrl,omitnil" name:"AccessUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UploadImageResponse struct {
	*tchttp.BaseResponse
	Response *UploadImageResponseParams `json:"Response"`
}

func (r *UploadImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}