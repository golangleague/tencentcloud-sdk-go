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

package v20201111

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Admin struct {
	// 超管名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 超管手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`
}

type Agent struct {
	// 代理机构的应用编号,32位字符串，一般不用传
	//
	// Deprecated: AppId is deprecated.
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 被代理机构的应用号，一般不用传
	//
	// Deprecated: ProxyAppId is deprecated.
	ProxyAppId *string `json:"ProxyAppId,omitnil" name:"ProxyAppId"`

	// 被代理机构在电子签平台的机构编号，集团代理下场景必传
	ProxyOrganizationId *string `json:"ProxyOrganizationId,omitnil" name:"ProxyOrganizationId"`

	// 被代理机构的经办人，一般不用传
	//
	// Deprecated: ProxyOperator is deprecated.
	ProxyOperator *string `json:"ProxyOperator,omitnil" name:"ProxyOperator"`
}

type ApproverInfo struct {
	// 在指定签署方时，可选择企业B端或个人C端等不同的参与者类型，可选类型如下:
	// **0**：企业
	// **1**：个人
	// **3**：企业静默签署
	// 注：`类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。静默签署仅进行盖章操作，不能自动签名。`
	// **7**: 个人自动签署，适用于个人自动签场景。
	// 注: `个人自动签场景为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverType *int64 `json:"ApproverType,omitnil" name:"ApproverType"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	ApproverName *string `json:"ApproverName,omitnil" name:"ApproverName"`

	// 签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此合同签署方。
	ApproverMobile *string `json:"ApproverMobile,omitnil" name:"ApproverMobile"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 如果签署方是企业签署方(approverType = 0 或者 approverType = 3)， 则企业名称必填。
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 合同中的签署控件列表，列表中可支持下列多种签署控件,控件的详细定义参考开发者中心的Component结构体
	// <ul><li> 个人签名/印章</li>
	// <li> 企业印章</li>
	// <li> 骑缝章等签署控件</li></ul>
	SignComponents []*Component `json:"SignComponents,omitnil" name:"SignComponents"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD 居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)</li>
	// <li>OTHER_CARD_TYPE 其他证件</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverIdCardType *string `json:"ApproverIdCardType,omitnil" name:"ApproverIdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母（但“I”、“O”除外），后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitnil" name:"ApproverIdCardNumber"`

	// 通知签署方经办人的方式,  有以下途径:
	// <ul><li>  **sms**  :  (默认)短信</li>
	// <li>   **none**   : 不通知</li></ul>
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 收据场景设置签署人角色类型, 可以设置如下****类型****:
	// <ul><li> **1**  :收款人</li>
	// <li>   **2**   :开具人</li>
	// <li>   **3** :见证人</li></ul>
	// 注: `收据场景为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverRole *int64 `json:"ApproverRole,omitnil" name:"ApproverRole"`

	// 签署意愿确认渠道，默认为WEIXINAPP:人脸识别
	// 
	// 注: 将要废弃, 用ApproverSignTypes签署人签署合同时的认证方式代替, 新客户可请用ApproverSignTypes来设置
	VerifyChannel []*string `json:"VerifyChannel,omitnil" name:"VerifyChannel"`

	// 签署方在签署合同之前，需要强制阅读合同的时长，可指定为3秒至300秒之间的任意值。
	// 
	// 若未指定阅读时间，则会按照合同页数大小计算阅读时间，计算规则如下：
	// <ul><li>合同页数少于等于2页，阅读时间为3秒；</li>
	// <li>合同页数为3到5页，阅读时间为5秒；</li>
	// <li>合同页数大于等于6页，阅读时间为10秒。</li></ul>
	PreReadTime *int64 `json:"PreReadTime,omitnil" name:"PreReadTime"`

	// 签署人userId，仅支持本企业的员工userid， 可在控制台组织管理处获得
	// 
	// 注: `若传此字段 则以userid的信息为主，会覆盖传递过来的签署人基本信息， 包括姓名，手机号，证件类型等信息`
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 在企微场景下使用，需设置参数为**WEWORKAPP**，以表明合同来源于企微。
	ApproverSource *string `json:"ApproverSource,omitnil" name:"ApproverSource"`

	// 在企业微信场景下，表明该合同流程为或签，其最大长度为64位字符串。
	// 所有参与或签的人员均需具备该标识。
	// 注意，在合同中，不同的或签参与人必须保证其CustomApproverTag唯一。
	// 如果或签签署人为本方企业微信参与人，则需要指定ApproverSource参数为WEWORKAPP。
	CustomApproverTag *string `json:"CustomApproverTag,omitnil" name:"CustomApproverTag"`

	// 可以控制签署方在签署合同时能否进行某些操作，例如拒签、转交他人等。
	// 详细操作可以参考开发者中心的ApproverOption结构体。
	ApproverOption *ApproverOption `json:"ApproverOption,omitnil" name:"ApproverOption"`

	// 指定个人签署方查看合同的校验方式,可以传值如下:
	// <ul><li>  **1**   : （默认）人脸识别,人脸识别后才能合同内容</li>
	// <li>  **2**  : 手机号验证, 用户手机号和参与方手机号(ApproverMobile)相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// </li></ul>
	// 注: 
	// <ul><li>如果合同流程设置ApproverVerifyType查看合同的校验方式,    则忽略此签署人的查看合同的校验方式</li>
	// <li>此字段不可传多个校验方式</li></ul>
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil" name:"ApproverVerifyTypes"`

	// 您可以指定签署方签署合同的认证校验方式，可传递以下值：
	// <ul><li>**1**：人脸认证，需进行人脸识别成功后才能签署合同；</li>
	// <li>**2**：签署密码，需输入与用户在腾讯电子签设置的密码一致才能校验成功进行合同签署；</li>
	// <li>**3**：运营商三要素，需到运营商处比对手机号实名信息（名字、手机号、证件号）校验一致才能成功进行合同签署。</li></ul>
	// 注：
	// <ul><li>默认情况下，认证校验方式为人脸认证和签署密码两种形式；</li>
	// <li>您可以传递多种值，表示可用多种认证校验方式。</li></ul>
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil" name:"ApproverSignTypes"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li>**false**：（默认）不需要审批，直接签署。</li>
	// <li>**true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li>如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li>如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	ApproverNeedSignReview *bool `json:"ApproverNeedSignReview,omitnil" name:"ApproverNeedSignReview"`
}

type ApproverOption struct {
	// 签署方是否可以拒签
	// 
	// <ul><li> **false** : ( 默认)可以拒签</li>
	// <li> **true** :不可以拒签</li></ul>
	NoRefuse *bool `json:"NoRefuse,omitnil" name:"NoRefuse"`

	// 签署方是否可以转他人处理
	// 
	// <ul><li> **false** : ( 默认)可以转他人处理</li>
	// <li> **true** :不可以转他人处理</li></ul>
	NoTransfer *bool `json:"NoTransfer,omitnil" name:"NoTransfer"`
}

type ApproverRestriction struct {
	// 指定签署人名字
	Name *string `json:"Name,omitnil" name:"Name"`

	// 指定签署人手机号，11位数字
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 指定签署人证件类型，ID_CARD-身份证
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 指定签署人证件号码，字母大写
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`
}

type AuthorizedUser struct {
	// 电子签系统中的用户id
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type AutoSignConfig struct {
	// 自动签开通个人用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 是否回调证书信息:
	// <ul><li>**false**: 不需要(默认)</li>
	// <li>**true**:需要</li></ul>
	CertInfoCallback *bool `json:"CertInfoCallback,omitnil" name:"CertInfoCallback"`

	// 是否支持用户自定义签名印章:
	// <ul><li>**false**: 不能自己定义(默认)</li>
	// <li>**true**: 可以自己定义</li></ul>
	UserDefineSeal *bool `json:"UserDefineSeal,omitnil" name:"UserDefineSeal"`

	// 回调中是否需要自动签将要使用的印章(签名) 图片的 base64:
	// <ul><li>**false**: 不需要(默认)</li>
	// <li>**true**: 需要</li></ul>
	SealImgCallback *bool `json:"SealImgCallback,omitnil" name:"SealImgCallback"`

	// 执行结果的回调URL，该URL仅支持HTTP或HTTPS协议，建议采用HTTPS协议以保证数据传输的安全性。
	// 腾讯电子签服务器将通过POST方式，application/json格式通知执行结果，请确保外网可以正常访问该URL。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 开通时候的身份验证方式, 取值为：
	// <ul><li>**WEIXINAPP** : 微信人脸识别</li>
	// <li>**INSIGHT** : 慧眼人脸认别</li>
	// <li>**TELECOM** : 运营商三要素验证</li></ul>
	// 注：
	// <ul><li>如果是小程序开通链接，支持传 WEIXINAPP / TELECOM。为空默认 WEIXINAPP</li>
	// <li>如果是 H5 开通链接，支持传 INSIGHT / TELECOM。为空默认 INSIGHT </li></ul>
	VerifyChannels []*string `json:"VerifyChannels,omitnil" name:"VerifyChannels"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。
	// 
	// <ul><li>**0**: (默认) 使用个人自动签账号许可进行开通，个人自动签账号许可有效期1年，注: `不可解绑释放更换他人`</li>
	// <li>**1**: 不使用个人自动签账号许可进行开通</li></ul>
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`
}

// Predefined struct for user
type BindEmployeeUserIdWithClientOpenIdRequestParams struct {
	// 用户信息，OpenId与UserId二选一必填一个，OpenId是第三方客户ID，userId是用户实名后的电子签生成的ID,当传入客户系统openId，传入的openId需与电子签员工userId绑定，且参数Channel必填，Channel值为INTEGRATE；当传入参数UserId，Channel无需指定。（参数参考示例）
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子签系统员工UserId
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 客户系统OpenId
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type BindEmployeeUserIdWithClientOpenIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息，OpenId与UserId二选一必填一个，OpenId是第三方客户ID，userId是用户实名后的电子签生成的ID,当传入客户系统openId，传入的openId需与电子签员工userId绑定，且参数Channel必填，Channel值为INTEGRATE；当传入参数UserId，Channel无需指定。（参数参考示例）
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子签系统员工UserId
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 客户系统OpenId
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *BindEmployeeUserIdWithClientOpenIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEmployeeUserIdWithClientOpenIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "UserId")
	delete(f, "OpenId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindEmployeeUserIdWithClientOpenIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindEmployeeUserIdWithClientOpenIdResponseParams struct {
	// 绑定是否成功，1表示成功，0表示失败
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BindEmployeeUserIdWithClientOpenIdResponse struct {
	*tchttp.BaseResponse
	Response *BindEmployeeUserIdWithClientOpenIdResponseParams `json:"Response"`
}

func (r *BindEmployeeUserIdWithClientOpenIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEmployeeUserIdWithClientOpenIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallbackInfo struct {
	// 回调url,。请确保回调地址能够接收并处理 HTTP POST 请求，并返回状态码 200 以表示处理正常。
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 回调加密key，已废弃
	//
	// Deprecated: Token is deprecated.
	Token *string `json:"Token,omitnil" name:"Token"`

	// 回调加密key，用于回调消息加解密。
	CallbackKey *string `json:"CallbackKey,omitnil" name:"CallbackKey"`

	// 回调验签token，用于回调通知校验。
	CallbackToken *string `json:"CallbackToken,omitnil" name:"CallbackToken"`
}

type Caller struct {
	// 应用号
	//
	// Deprecated: ApplicationId is deprecated.
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 主机构ID
	//
	// Deprecated: OrganizationId is deprecated.
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 经办人的用户ID，同UserId
	OperatorId *string `json:"OperatorId,omitnil" name:"OperatorId"`

	// 下属机构ID
	//
	// Deprecated: SubOrganizationId is deprecated.
	SubOrganizationId *string `json:"SubOrganizationId,omitnil" name:"SubOrganizationId"`
}

// Predefined struct for user
type CancelFlowRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程ID, 为32位字符串。
	// 建议开发者保存此流程ID方便后续其他操作。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 撤销此合同(流程)的原因，最长200个字。
	CancelMessage *string `json:"CancelMessage,omitnil" name:"CancelMessage"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CancelFlowRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程ID, 为32位字符串。
	// 建议开发者保存此流程ID方便后续其他操作。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 撤销此合同(流程)的原因，最长200个字。
	CancelMessage *string `json:"CancelMessage,omitnil" name:"CancelMessage"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CancelFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "CancelMessage")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelFlowResponse struct {
	*tchttp.BaseResponse
	Response *CancelFlowResponseParams `json:"Response"`
}

func (r *CancelFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelMultiFlowSignQRCodeRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 二维码ID，为32位字符串。
	QrCodeId *string `json:"QrCodeId,omitnil" name:"QrCodeId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CancelMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 二维码ID，为32位字符串。
	QrCodeId *string `json:"QrCodeId,omitnil" name:"QrCodeId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CancelMultiFlowSignQRCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelMultiFlowSignQRCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "QrCodeId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelMultiFlowSignQRCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelMultiFlowSignQRCodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelMultiFlowSignQRCodeResponse struct {
	*tchttp.BaseResponse
	Response *CancelMultiFlowSignQRCodeResponseParams `json:"Response"`
}

func (r *CancelMultiFlowSignQRCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelMultiFlowSignQRCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelUserAutoSignEnableUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li></ul>
	// 
	// 注: `现在仅支持电子处方场景`
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 预撤销链接的用户信息，包含姓名、证件类型、证件号码等信息。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CancelUserAutoSignEnableUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li></ul>
	// 
	// 注: `现在仅支持电子处方场景`
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 预撤销链接的用户信息，包含姓名、证件类型、证件号码等信息。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CancelUserAutoSignEnableUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelUserAutoSignEnableUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelUserAutoSignEnableUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelUserAutoSignEnableUrlResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelUserAutoSignEnableUrlResponse struct {
	*tchttp.BaseResponse
	Response *CancelUserAutoSignEnableUrlResponseParams `json:"Response"`
}

func (r *CancelUserAutoSignEnableUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelUserAutoSignEnableUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcInfo struct {
	// 被抄送方手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此业务通知方。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 被抄送方姓名。
	// 抄送方的姓名将用于身份认证，请确保填写的姓名为抄送方的真实姓名，而非昵称等代名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 被抄送方类型, 可设置以下类型:
	// <ul><li> **0** :个人抄送方</li>
	// <li> **1** :企业员工抄送方</li></ul>
	CcType *int64 `json:"CcType,omitnil" name:"CcType"`

	// 被抄送方权限, 可设置如下权限:
	// <ul><li> **0** :可查看合同内容</li>
	// <li> **1** :可查看合同内容也可下载原文</li></ul>
	CcPermission *int64 `json:"CcPermission,omitnil" name:"CcPermission"`

	// 通知签署方经办人的方式,  有以下途径:
	// <ul><li> **sms** :  (默认)短信</li>
	// <li> **none** : 不通知</li></ul>
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`
}

type Component struct {
	// 如果是Component填写控件类型，则可选的字段为：
	// TEXT - 普通文本控件，输入文本字符串；
	// MULTI_LINE_TEXT - 多行文本控件，输入文本字符串；
	// CHECK_BOX - 勾选框控件，若选中填写ComponentValue 填写 true或者 false 字符串；
	// FILL_IMAGE - 图片控件，ComponentValue 填写图片的资源 ID；
	// DYNAMIC_TABLE - 动态表格控件；
	// ATTACHMENT - 附件控件,ComponentValue 填写附件图片的资源 ID列表，以逗号分隔；
	// SELECTOR - 选择器控件，ComponentValue填写选择的字符串内容；
	// DATE - 日期控件；默认是格式化为xxxx年xx月xx日字符串；
	// DISTRICT - 省市区行政区控件，ComponentValue填写省市区行政区字符串内容；
	// 
	// 如果是SignComponent签署控件类型，则可选的字段为
	// SIGN_SEAL - 签署印章控件；
	// SIGN_DATE - 签署日期控件；
	// SIGN_SIGNATURE - 用户签名控件；
	// SIGN_PERSONAL_SEAL - 个人签署印章控件（使用文件发起暂不支持此类型）；
	// SIGN_PAGING_SEAL - 骑缝章；若文件发起，需要对应填充ComponentPosY、ComponentWidth、ComponentHeight
	// SIGN_OPINION - 签署意见控件，用户需要根据配置的签署意见内容，完成对意见内容的确认；
	// SIGN_LEGAL_PERSON_SEAL - 企业法定代表人控件。
	// 
	// 表单域的控件不能作为印章和签名控件
	ComponentType *string `json:"ComponentType,omitnil" name:"ComponentType"`

	// 控件所属文件的序号（取值为：0-N）。
	// 目前单文件的情况下，值是0
	FileIndex *int64 `json:"FileIndex,omitnil" name:"FileIndex"`

	// 参数控件高度，单位pt
	ComponentHeight *float64 `json:"ComponentHeight,omitnil" name:"ComponentHeight"`

	// 参数控件宽度，单位pt
	ComponentWidth *float64 `json:"ComponentWidth,omitnil" name:"ComponentWidth"`

	// 参数控件所在页码，取值为：1-N
	ComponentPage *int64 `json:"ComponentPage,omitnil" name:"ComponentPage"`

	// 参数控件X位置，单位pt
	ComponentPosX *float64 `json:"ComponentPosX,omitnil" name:"ComponentPosX"`

	// 参数控件Y位置，单位pt
	ComponentPosY *float64 `json:"ComponentPosY,omitnil" name:"ComponentPosY"`

	// 控件唯一ID。
	// 或使用文件发起合同时用于GenerateMode==KEYWORD 指定关键字
	ComponentId *string `json:"ComponentId,omitnil" name:"ComponentId"`

	// 控件名。
	// 或使用文件发起合同时用于GenerateMode==FIELD 指定表单域名称
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// 是否必选，默认为false-非必选
	ComponentRequired *bool `json:"ComponentRequired,omitnil" name:"ComponentRequired"`

	// 控件关联的参与方ID，对应Recipient结构体中的RecipientId
	ComponentRecipientId *string `json:"ComponentRecipientId,omitnil" name:"ComponentRecipientId"`

	// 扩展参数：
	// 为JSON格式。
	// 
	// ComponentType为FILL_IMAGE时，支持以下参数：
	// NotMakeImageCenter：bool。是否设置图片居中。false：居中（默认）。 true: 不居中
	// FillMethod: int. 填充方式。0-铺满（默认）；1-等比例缩放
	// 
	// ComponentType为SIGN_SIGNATURE类型可以控制签署方式
	// {“ComponentTypeLimit”: [“xxx”]}
	// xxx可以为：
	// HANDWRITE – 手写签名
	// OCR_ESIGN -- AI智能识别手写签名
	// ESIGN -- 个人印章类型
	// SYSTEM_ESIGN -- 系统签名（该类型可以在用户签署时根据用户姓名一键生成一个签名来进行签署）
	// 如：{“ComponentTypeLimit”: [“SYSTEM_ESIGN”]}
	// 
	// ComponentType为SIGN_DATE时，支持以下参数：
	// 1 Font：字符串类型目前只支持"黑体"、"宋体"，如果不填默认为"黑体"
	// 2 FontSize： 数字类型，范围6-72，默认值为12
	// 3 FontAlign： 字符串类型，可取Left/Right/Center，对应左对齐/居中/右对齐
	// 4 Format： 字符串类型，日期格式，必须是以下五种之一 “yyyy m d”，”yyyy年m月d日”，”yyyy/m/d”，”yyyy-m-d”，”yyyy.m.d”。
	// 5 Gaps:： 字符串类型，仅在Format为“yyyy m d”时起作用，格式为用逗号分开的两个整数，例如”2,2”，两个数字分别是日期格式的前后两个空隙中的空格个数
	// 如果extra参数为空，默认为”yyyy年m月d日”格式的居中日期
	// 特别地，如果extra中Format字段为空或无法被识别，则extra参数会被当作默认值处理（Font，FontSize，Gaps和FontAlign都不会起效）
	// 参数样例：    "ComponentExtra": "{\"Format\":“yyyy m d”,\"FontSize\":12,\"Gaps\":\"2,2\", \"FontAlign\":\"Right\"}"
	// 
	// ComponentType为SIGN_SEAL类型时，支持以下参数：
	// 1.PageRanges：PageRange的数组，通过PageRanges属性设置该印章在PDF所有页面上盖章（适用于标书在所有页面盖章的情况）
	// 参数样例："ComponentExtra":"{\"PageRanges\":[{\"BeginPage\":1,\"EndPage\":-1}]}"
	ComponentExtra *string `json:"ComponentExtra,omitnil" name:"ComponentExtra"`

	// 是否是表单域类型，默认false-不是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFormType *bool `json:"IsFormType,omitnil" name:"IsFormType"`

	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// TEXT - 文本内容
	// MULTI_LINE_TEXT - 文本内容
	// CHECK_BOX - true/false
	// FILL_IMAGE、ATTACHMENT - 附件的FileId，需要通过UploadFiles接口上传获取
	// SELECTOR - 选项值
	// DYNAMIC_TABLE - 传入json格式的表格内容，具体见数据结构FlowInfo：https://cloud.tencent.com/document/api/1420/61525#FlowInfo
	// DATE - 默认是格式化为xxxx年xx月xx日
	// SIGN_SEAL - 印章ID，于控制台查询获取
	// SIGN_PAGING_SEAL - 可以指定印章ID，于控制台查询获取
	// 
	// 控件值约束说明：
	// 企业全称控件：
	//   约束：企业名称中文字符中文括号
	//   检查正则表达式：/^[\u3400-\u4dbf\u4e00-\u9fa5（）]+$/
	// 
	// 统一社会信用代码控件：
	//   检查正则表达式：/^[A-Z0-9]{1,18}$/
	// 
	// 法人名称控件：
	//   约束：最大50个字符，2到25个汉字或者1到50个字母
	//   检查正则表达式：/^([\u3400-\u4dbf\u4e00-\u9fa5.·]{2,25}|[a-zA-Z·,\s-]{1,50})$/
	// 
	// 签署意见控件：
	//   约束：签署意见最大长度为50字符
	// 
	// 签署人手机号控件：
	//   约束：国内手机号 13,14,15,16,17,18,19号段长度11位
	// 
	// 签署人身份证控件：
	//   约束：合法的身份证号码检查
	// 
	// 控件名称：
	//   约束：控件名称最大长度为20字符
	// 
	// 单行文本控件：
	//   约束：只允许输入中文，英文，数字，中英文标点符号
	// 
	// 多行文本控件：
	//   约束：只允许输入中文，英文，数字，中英文标点符号
	// 
	// 勾选框控件：
	//   约束：选择填字符串true，不选填字符串false
	// 
	// 选择器控件：
	//   约束：同单行文本控件约束，填写选择值中的字符串
	// 
	// 数字控件：
	//   约束：请输入有效的数字(可带小数点) 
	//   检查正则表达式：/^(-|\+)?\d+(\.\d+)?$/
	// 
	// 日期控件：
	//   约束：格式：yyyy年mm月dd日
	// 
	// 附件控件：
	//   约束：JPG或PNG图片，上传数量限制，1到6个，最大6个附件
	// 
	// 图片控件：
	//   约束：JPG或PNG图片，填写上传的图片资源ID
	// 
	// 邮箱控件：
	//   约束：请输入有效的邮箱地址, w3c标准
	//   检查正则表达式：/^([A-Za-z0-9_\-.!#$%&])+@([A-Za-z0-9_\-.])+\.([A-Za-z]{2,4})$/
	//   参考：https://emailregex.com/
	// 
	// 地址控件：
	//   同单行文本控件约束
	// 
	// 省市区控件：
	//   同单行文本控件约束
	// 
	// 性别控件：
	//   同单行文本控件约束，填写选择值中的字符串
	// 
	// 学历控件：
	//   同单行文本控件约束，填写选择值中的字符串
	ComponentValue *string `json:"ComponentValue,omitnil" name:"ComponentValue"`

	// NORMAL 正常模式，使用坐标制定签署控件位置
	// FIELD 表单域，需使用ComponentName指定表单域名称
	// KEYWORD 关键字，使用ComponentId指定关键字
	GenerateMode *string `json:"GenerateMode,omitnil" name:"GenerateMode"`

	// 日期签署控件的字号，默认为 12
	ComponentDateFontSize *int64 `json:"ComponentDateFontSize,omitnil" name:"ComponentDateFontSize"`

	// 第三方应用集成平台模板控件 ID 标识
	ChannelComponentId *string `json:"ChannelComponentId,omitnil" name:"ChannelComponentId"`

	// 指定关键字时横坐标偏移量，单位pt
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetX *float64 `json:"OffsetX,omitnil" name:"OffsetX"`

	// 指定关键字时纵坐标偏移量，单位pt
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetY *float64 `json:"OffsetY,omitnil" name:"OffsetY"`

	// 第三方应用集成中子客企业控件来源。
	// 0-平台指定；
	// 1-用户自定义
	ChannelComponentSource *uint64 `json:"ChannelComponentSource,omitnil" name:"ChannelComponentSource"`

	// 指定关键字排序规则，Positive-正序，Reverse-倒序。
	// 传入Positive时会根据关键字在PDF文件内的顺序进行排列。在指定KeywordIndexes时，0代表在PDF内查找内容时，查找到的第一个关键字。
	// 传入Reverse时会根据关键字在PDF文件内的反序进行排列。在指定KeywordIndexes时，0代表在PDF内查找内容时，查找到的最后一个关键字。
	KeywordOrder *string `json:"KeywordOrder,omitnil" name:"KeywordOrder"`

	// 指定关键字页码。
	// 指定页码后，将只在指定的页码内查找关键字，非该页码的关键字将不会查询出来
	KeywordPage *int64 `json:"KeywordPage,omitnil" name:"KeywordPage"`

	// 关键字位置模式，
	// Middle-居中，
	// Below-正下方，
	// Right-正右方，
	// LowerRight-右上角，
	// UpperRight-右下角。
	// 示例：如果设置Middle的关键字盖章，则印章的中心会和关键字的中心重合，如果设置Below，则印章在关键字的正下方
	RelativeLocation *string `json:"RelativeLocation,omitnil" name:"RelativeLocation"`

	// 关键字索引。
	// 如果一个关键字在PDF文件中存在多个，可以通过关键字索引指定使用第几个关键字作为最后的结果，可指定多个索引。
	// 示例：[0,2]，说明使用PDF文件内第1个和第3个关键字位置。
	KeywordIndexes []*int64 `json:"KeywordIndexes,omitnil" name:"KeywordIndexes"`

	// 是否锁定控件值不允许编辑（嵌入式发起使用）
	// <br/>默认false：不锁定控件值，允许在页面编辑控件值
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockComponentValue *bool `json:"LockComponentValue,omitnil" name:"LockComponentValue"`

	// 是否禁止移动和删除控件
	// <br/>默认false，不禁止移动和删除控件
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForbidMoveAndDelete *bool `json:"ForbidMoveAndDelete,omitnil" name:"ForbidMoveAndDelete"`
}

// Predefined struct for user
type CreateBatchCancelFlowUrlRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 需要执行撤回的流程(合同)的编号列表，最多100个.
	// 列表中的流程(合同)编号不要重复.
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreateBatchCancelFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 需要执行撤回的流程(合同)的编号列表，最多100个.
	// 列表中的流程(合同)编号不要重复.
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreateBatchCancelFlowUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchCancelFlowUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchCancelFlowUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchCancelFlowUrlResponseParams struct {
	// 批量撤回签署流程链接
	BatchCancelFlowUrl *string `json:"BatchCancelFlowUrl,omitnil" name:"BatchCancelFlowUrl"`

	// 签署流程撤回失败信息
	// 数组里边的错误原因与传进来的FlowIds一一对应,如果是空字符串则标识没有出错
	FailMessages []*string `json:"FailMessages,omitnil" name:"FailMessages"`

	// 签署连接过期时间字符串：年月日-时分秒
	// 
	// 例如:2023-07-28 17:25:59
	UrlExpireOn *string `json:"UrlExpireOn,omitnil" name:"UrlExpireOn"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBatchCancelFlowUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchCancelFlowUrlResponseParams `json:"Response"`
}

func (r *CreateBatchCancelFlowUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchCancelFlowUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelSubOrganizationModifyQrCodeRequestParams struct {
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 应用编号
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

type CreateChannelSubOrganizationModifyQrCodeRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 应用编号
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

func (r *CreateChannelSubOrganizationModifyQrCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelSubOrganizationModifyQrCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChannelSubOrganizationModifyQrCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelSubOrganizationModifyQrCodeResponseParams struct {
	// 二维码下载链接
	QrCodeUrl *string `json:"QrCodeUrl,omitnil" name:"QrCodeUrl"`

	// 二维码失效时间 UNIX 时间戳 精确到秒
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateChannelSubOrganizationModifyQrCodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateChannelSubOrganizationModifyQrCodeResponseParams `json:"Response"`
}

func (r *CreateChannelSubOrganizationModifyQrCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelSubOrganizationModifyQrCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConvertTaskApiRequestParams struct {
	// 需要进行转换的资源文件类型
	// 支持的文件类型如下：
	// <ul><li>doc</li>
	// <li>docx</li>
	// <li>xls</li>
	// <li>xlsx</li>
	// <li>jpg</li>
	// <li>jpeg</li>
	// <li>png</li>
	// <li>bmp</li>
	// <li>txt</li></ul>
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 需要进行转换操作的文件资源名称，带资源后缀名。
	// 
	// 注:  `资源名称长度限制为256个字符`
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 需要进行转换操作的文件资源Id，通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles" target="_blank">UploadFiles</a>接口获取文件资源Id。
	// 
	// 注:  `目前，此接口仅支持单个文件进行转换。`
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`
}

type CreateConvertTaskApiRequest struct {
	*tchttp.BaseRequest
	
	// 需要进行转换的资源文件类型
	// 支持的文件类型如下：
	// <ul><li>doc</li>
	// <li>docx</li>
	// <li>xls</li>
	// <li>xlsx</li>
	// <li>jpg</li>
	// <li>jpeg</li>
	// <li>png</li>
	// <li>bmp</li>
	// <li>txt</li></ul>
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 需要进行转换操作的文件资源名称，带资源后缀名。
	// 
	// 注:  `资源名称长度限制为256个字符`
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 需要进行转换操作的文件资源Id，通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles" target="_blank">UploadFiles</a>接口获取文件资源Id。
	// 
	// 注:  `目前，此接口仅支持单个文件进行转换。`
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`
}

func (r *CreateConvertTaskApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConvertTaskApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "ResourceName")
	delete(f, "ResourceId")
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "Organization")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConvertTaskApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConvertTaskApiResponseParams struct {
	// 接口返回的文件转换任务Id，可以调用接口<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/GetTaskResultApi" target="_blank">查询转换任务状态</a>获取转换任务的状态和转换后的文件资源Id。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateConvertTaskApiResponse struct {
	*tchttp.BaseResponse
	Response *CreateConvertTaskApiResponseParams `json:"Response"`
}

func (r *CreateConvertTaskApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConvertTaskApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentRequestParams struct {
	// 调用方用户信息，userId 必填。支持填入集团子公司经办人 userId代发合同。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 此接口的合同流程ID需要由<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow" target="_blank">创建签署流程</a>接口创建得到。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 用户配置的合同模板ID，会基于此模板创建合同文档，为32位字符串。
	// 可登录腾讯电子签控制台，在 "模板"->"模板中心"->"列表展示设置"选中模板 ID 中查看某个模板的TemplateId(在页面中展示为模板ID)。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 文件名列表，单个文件名最大长度200个字符，暂时仅支持单文件发起。设置后流程对应的文件名称当前设置的值。
	FileNames []*string `json:"FileNames,omitnil" name:"FileNames"`

	// 电子文档的填写控件的填充内容。具体方式可以参考<a href="https://qian.tencent.com/developers/companyApis/dataTypes/#formfield" target="_blank">FormField</a>结构体的定义。
	FormFields []*FormField `json:"FormFields,omitnil" name:"FormFields"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li>
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。</li></ul>
	NeedPreview *bool `json:"NeedPreview,omitnil" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil" name:"PreviewType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 已废弃字段，客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

type CreateDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填。支持填入集团子公司经办人 userId代发合同。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 此接口的合同流程ID需要由<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow" target="_blank">创建签署流程</a>接口创建得到。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 用户配置的合同模板ID，会基于此模板创建合同文档，为32位字符串。
	// 可登录腾讯电子签控制台，在 "模板"->"模板中心"->"列表展示设置"选中模板 ID 中查看某个模板的TemplateId(在页面中展示为模板ID)。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 文件名列表，单个文件名最大长度200个字符，暂时仅支持单文件发起。设置后流程对应的文件名称当前设置的值。
	FileNames []*string `json:"FileNames,omitnil" name:"FileNames"`

	// 电子文档的填写控件的填充内容。具体方式可以参考<a href="https://qian.tencent.com/developers/companyApis/dataTypes/#formfield" target="_blank">FormField</a>结构体的定义。
	FormFields []*FormField `json:"FormFields,omitnil" name:"FormFields"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li>
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。</li></ul>
	NeedPreview *bool `json:"NeedPreview,omitnil" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil" name:"PreviewType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 已废弃字段，客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

func (r *CreateDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "TemplateId")
	delete(f, "FileNames")
	delete(f, "FormFields")
	delete(f, "NeedPreview")
	delete(f, "PreviewType")
	delete(f, "Agent")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentResponseParams struct {
	// 合同流程的底层电子文档ID，为32位字符串。
	// 
	// 注:
	// 后续需用同样的FlowId再次调用<a href="https://qian.tencent.com/developers/companyApis/startFlows/StartFlow" target="_blank">发起签署流程</a>，合同才能进入签署环节
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`

	// 合同预览链接URL。
	// 
	// 注：如果是预览模式(即NeedPreview设置为true)时, 才会有此预览链接URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewFileUrl *string `json:"PreviewFileUrl,omitnil" name:"PreviewFileUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDocumentResponse struct {
	*tchttp.BaseResponse
	Response *CreateDocumentResponseParams `json:"Response"`
}

func (r *CreateDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmbedWebUrlRequestParams struct {
	// 操作者信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// WEB嵌入资源类型。
	// <br/>CREATE_SEAL: 生成创建印章的嵌入页面
	// <br/>CREATE_TEMPLATE：生成创建模板的嵌入页面
	// <br/>MODIFY_TEMPLATE：生成编辑模板的嵌入页面
	// <br/>PREVIEW_TEMPLATE：生成预览模板的嵌入页面
	// <br/>PREVIEW_SEAL_LIST：生成预览印章列表的嵌入页面
	// <br/>PREVIEW_SEAL_DETAIL：生成预览印章详情的嵌入页面
	// <br/>EXTEND_SERVICE：生成拓展服务的嵌入页面
	// <br/>PREVIEW_FLOW：生成预览合同的嵌入页面
	// <br/>PREVIEW_FLOW_DETAIL：生成查看合同详情的嵌入页面
	EmbedType *string `json:"EmbedType,omitnil" name:"EmbedType"`

	// WEB嵌入的业务资源ID
	// <br/>PREVIEW_SEAL_DETAIL，必填，取值为印章id
	// <br/>MODIFY_TEMPLATE，PREVIEW_TEMPLATE，必填，取值为模板id
	// <br/>PREVIEW_FLOW，PREVIEW_FLOW_DETAIL，必填，取值为合同id
	BusinessId *string `json:"BusinessId,omitnil" name:"BusinessId"`

	// 代理相关应用信息，如集团主企业代子企业操作
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 抄送方信息
	Reviewer *ReviewerInfo `json:"Reviewer,omitnil" name:"Reviewer"`

	// 个性化参数
	Option *EmbedUrlOption `json:"Option,omitnil" name:"Option"`
}

type CreateEmbedWebUrlRequest struct {
	*tchttp.BaseRequest
	
	// 操作者信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// WEB嵌入资源类型。
	// <br/>CREATE_SEAL: 生成创建印章的嵌入页面
	// <br/>CREATE_TEMPLATE：生成创建模板的嵌入页面
	// <br/>MODIFY_TEMPLATE：生成编辑模板的嵌入页面
	// <br/>PREVIEW_TEMPLATE：生成预览模板的嵌入页面
	// <br/>PREVIEW_SEAL_LIST：生成预览印章列表的嵌入页面
	// <br/>PREVIEW_SEAL_DETAIL：生成预览印章详情的嵌入页面
	// <br/>EXTEND_SERVICE：生成拓展服务的嵌入页面
	// <br/>PREVIEW_FLOW：生成预览合同的嵌入页面
	// <br/>PREVIEW_FLOW_DETAIL：生成查看合同详情的嵌入页面
	EmbedType *string `json:"EmbedType,omitnil" name:"EmbedType"`

	// WEB嵌入的业务资源ID
	// <br/>PREVIEW_SEAL_DETAIL，必填，取值为印章id
	// <br/>MODIFY_TEMPLATE，PREVIEW_TEMPLATE，必填，取值为模板id
	// <br/>PREVIEW_FLOW，PREVIEW_FLOW_DETAIL，必填，取值为合同id
	BusinessId *string `json:"BusinessId,omitnil" name:"BusinessId"`

	// 代理相关应用信息，如集团主企业代子企业操作
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 抄送方信息
	Reviewer *ReviewerInfo `json:"Reviewer,omitnil" name:"Reviewer"`

	// 个性化参数
	Option *EmbedUrlOption `json:"Option,omitnil" name:"Option"`
}

func (r *CreateEmbedWebUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmbedWebUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "EmbedType")
	delete(f, "BusinessId")
	delete(f, "Agent")
	delete(f, "Reviewer")
	delete(f, "Option")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmbedWebUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmbedWebUrlResponseParams struct {
	// 嵌入的web链接，有效期：5分钟
	// EmbedType=PREVIEW_CC_FLOW，该url为h5链接
	WebUrl *string `json:"WebUrl,omitnil" name:"WebUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateEmbedWebUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmbedWebUrlResponseParams `json:"Response"`
}

func (r *CreateEmbedWebUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmbedWebUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowApproversRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 补充签署人信息
	Approvers []*FillApproverInfo `json:"Approvers,omitnil" name:"Approvers"`

	// 企微消息中的发起人
	Initiator *string `json:"Initiator,omitnil" name:"Initiator"`

	// 代理相关应用信息，如集团主企业代子企业操作
	// 
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreateFlowApproversRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 补充签署人信息
	Approvers []*FillApproverInfo `json:"Approvers,omitnil" name:"Approvers"`

	// 企微消息中的发起人
	Initiator *string `json:"Initiator,omitnil" name:"Initiator"`

	// 代理相关应用信息，如集团主企业代子企业操作
	// 
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreateFlowApproversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowApproversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "Approvers")
	delete(f, "Initiator")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowApproversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowApproversResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFlowApproversResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowApproversResponseParams `json:"Response"`
}

func (r *CreateFlowApproversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowApproversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowByFilesRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 
	// 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。
	// 
	// 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序，请确保列表中参与人的顺序符合实际签署顺序。
	Approvers []*ApproverInfo `json:"Approvers,omitnil" name:"Approvers"`

	// 本合同流程需包含的PDF文件资源编号列表，通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles" target="_blank">UploadFiles</a>接口获取PDF文件资源编号。
	// 
	// 注:  `目前，此接口仅支持单个文件发起。`
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li> 单行文本控件      </li>
	// <li> 多行文本控件      </li>
	// <li> 勾选框控件        </li>
	// <li> 数字控件          </li>
	// <li> 图片控件          </li>
	// <li> 动态表格等填写控件</li></ul>
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// 注:`此功能为白名单功能，使用前请联系对接的客户经理沟通。`
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 可以设置以下时间节点来给抄送人发送短信通知来查看合同内容：
	// <ul><li> **0**：合同发起时通知（默认值）</li>
	// <li> **1**：签署完成后通知</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil" name:"CcNotifyType"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li>
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。</li></ul>
	NeedPreview *bool `json:"NeedPreview,omitnil" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil" name:"PreviewType"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 合同到期提醒时间，为Unix标准时间戳（秒）格式，支持的范围是从发起时间开始到后10年内。
	// 
	// 到达提醒时间后，腾讯电子签会短信通知发起方企业合同提醒，可用于处理合同到期事务，如合同续签等事宜。
	RemindedOn *int64 `json:"RemindedOn,omitnil" name:"RemindedOn"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 您可以自定义腾讯电子签小程序合同列表页展示的合同内容模板，模板中支持以下变量：
	// <ul><li>{合同名称}   </li>
	// <li>{发起方企业} </li>
	// <li>{发起方姓名} </li>
	// <li>{签署方N企业}</li>
	// <li>{签署方N姓名}</li></ul>
	// 其中，N表示签署方的编号，从1开始，不能超过签署人的数量。
	// 
	// 例如，如果是腾讯公司张三发给李四名称为“租房合同”的合同，您可以将此字段设置为：`合同名称:{合同名称};发起方: {发起方企业}({发起方姓名});签署方:{签署方1姓名}`，则小程序中列表页展示此合同为以下样子
	// 
	// 合同名称：租房合同 
	// 发起方：腾讯公司(张三) 
	// 签署方：李四
	// 
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 指定个人签署方查看合同的校验方式
	// <ul><li>   **VerifyCheck**  :（默认）人脸识别,人脸识别后才能合同内容 </li>
	// <li>   **MobileCheck**  :  手机号验证, 用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）</li></ul>
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 签署方签署控件（印章/签名等）的生成方式：
	// <ul><li> **0**：在合同流程发起时，由发起人指定签署方的签署控件的位置和数量。</li>
	// <li> **1**：签署方在签署时自行添加签署控件，可以拖动位置和控制数量。</li></ul>
	SignBeanTag *int64 `json:"SignBeanTag,omitnil" name:"SignBeanTag"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：处方单（医疗自动签）  </li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil" name:"AutoSignScene"`
}

type CreateFlowByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 
	// 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。
	// 
	// 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序，请确保列表中参与人的顺序符合实际签署顺序。
	Approvers []*ApproverInfo `json:"Approvers,omitnil" name:"Approvers"`

	// 本合同流程需包含的PDF文件资源编号列表，通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles" target="_blank">UploadFiles</a>接口获取PDF文件资源编号。
	// 
	// 注:  `目前，此接口仅支持单个文件发起。`
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li> 单行文本控件      </li>
	// <li> 多行文本控件      </li>
	// <li> 勾选框控件        </li>
	// <li> 数字控件          </li>
	// <li> 图片控件          </li>
	// <li> 动态表格等填写控件</li></ul>
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// 注:`此功能为白名单功能，使用前请联系对接的客户经理沟通。`
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 可以设置以下时间节点来给抄送人发送短信通知来查看合同内容：
	// <ul><li> **0**：合同发起时通知（默认值）</li>
	// <li> **1**：签署完成后通知</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil" name:"CcNotifyType"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li>
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。</li></ul>
	NeedPreview *bool `json:"NeedPreview,omitnil" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil" name:"PreviewType"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 合同到期提醒时间，为Unix标准时间戳（秒）格式，支持的范围是从发起时间开始到后10年内。
	// 
	// 到达提醒时间后，腾讯电子签会短信通知发起方企业合同提醒，可用于处理合同到期事务，如合同续签等事宜。
	RemindedOn *int64 `json:"RemindedOn,omitnil" name:"RemindedOn"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 您可以自定义腾讯电子签小程序合同列表页展示的合同内容模板，模板中支持以下变量：
	// <ul><li>{合同名称}   </li>
	// <li>{发起方企业} </li>
	// <li>{发起方姓名} </li>
	// <li>{签署方N企业}</li>
	// <li>{签署方N姓名}</li></ul>
	// 其中，N表示签署方的编号，从1开始，不能超过签署人的数量。
	// 
	// 例如，如果是腾讯公司张三发给李四名称为“租房合同”的合同，您可以将此字段设置为：`合同名称:{合同名称};发起方: {发起方企业}({发起方姓名});签署方:{签署方1姓名}`，则小程序中列表页展示此合同为以下样子
	// 
	// 合同名称：租房合同 
	// 发起方：腾讯公司(张三) 
	// 签署方：李四
	// 
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 指定个人签署方查看合同的校验方式
	// <ul><li>   **VerifyCheck**  :（默认）人脸识别,人脸识别后才能合同内容 </li>
	// <li>   **MobileCheck**  :  手机号验证, 用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）</li></ul>
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 签署方签署控件（印章/签名等）的生成方式：
	// <ul><li> **0**：在合同流程发起时，由发起人指定签署方的签署控件的位置和数量。</li>
	// <li> **1**：签署方在签署时自行添加签署控件，可以拖动位置和控制数量。</li></ul>
	SignBeanTag *int64 `json:"SignBeanTag,omitnil" name:"SignBeanTag"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：处方单（医疗自动签）  </li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil" name:"AutoSignScene"`
}

func (r *CreateFlowByFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowByFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowName")
	delete(f, "Approvers")
	delete(f, "FileIds")
	delete(f, "FlowDescription")
	delete(f, "FlowType")
	delete(f, "Components")
	delete(f, "CcInfos")
	delete(f, "CcNotifyType")
	delete(f, "NeedPreview")
	delete(f, "PreviewType")
	delete(f, "Deadline")
	delete(f, "RemindedOn")
	delete(f, "Unordered")
	delete(f, "CustomShowMap")
	delete(f, "NeedSignReview")
	delete(f, "UserData")
	delete(f, "ApproverVerifyType")
	delete(f, "SignBeanTag")
	delete(f, "Agent")
	delete(f, "AutoSignScene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowByFilesResponseParams struct {
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	// 
	// 注: 如果是预览模式(即NeedPreview设置为true)时, 此处不会有值返回。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 合同预览链接URL。
	// 
	// 注：如果是预览模式(即NeedPreview设置为true)时, 才会有此预览链接URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitnil" name:"PreviewUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFlowByFilesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowByFilesResponseParams `json:"Response"`
}

func (r *CreateFlowByFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowByFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowEvidenceReportRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreateFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreateFlowEvidenceReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowEvidenceReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowEvidenceReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowEvidenceReportResponseParams struct {
	// 出证报告 ID，可用于DescribeFlowEvidenceReport接口查询出证PDF的下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`

	// 出证任务执行的状态, 可能会有以下状态：
	// 
	// <ul><li>EvidenceStatusExecuting：  出证任务在执行中</li>
	// <li>EvidenceStatusSuccess：  出证任务执行成功</li>
	// <li>EvidenceStatusFailed ： 出征任务执行失败</li></ul>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 此字段已经废除,不再使用.
	// 出证的PDF下载地址请调用DescribeChannelFlowEvidenceReport接口获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ReportUrl is deprecated.
	ReportUrl *string `json:"ReportUrl,omitnil" name:"ReportUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFlowEvidenceReportResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowEvidenceReportResponseParams `json:"Response"`
}

func (r *CreateFlowEvidenceReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowEvidenceReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowGroupByFilesRequestParams struct {
	// 调用方用户信息，userId 必填。支持填入集团子公司经办人 userId 代发合同
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同（流程）组名称,最大长度200个字符
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`

	// 合同（流程）组的子合同信息，支持2-50个子合同
	FlowGroupInfos []*FlowGroupInfo `json:"FlowGroupInfos,omitnil" name:"FlowGroupInfos"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同（流程）组的配置项信息。包括是否通知本企业签署方，是否通知其他签署方
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil" name:"FlowGroupOptions"`
}

type CreateFlowGroupByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填。支持填入集团子公司经办人 userId 代发合同
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同（流程）组名称,最大长度200个字符
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`

	// 合同（流程）组的子合同信息，支持2-50个子合同
	FlowGroupInfos []*FlowGroupInfo `json:"FlowGroupInfos,omitnil" name:"FlowGroupInfos"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同（流程）组的配置项信息。包括是否通知本企业签署方，是否通知其他签署方
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil" name:"FlowGroupOptions"`
}

func (r *CreateFlowGroupByFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowGroupByFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowGroupName")
	delete(f, "FlowGroupInfos")
	delete(f, "Agent")
	delete(f, "FlowGroupOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowGroupByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowGroupByFilesResponseParams struct {
	// 合同(流程)组的合同组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 合同(流程)组中子合同列表.
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFlowGroupByFilesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowGroupByFilesResponseParams `json:"Response"`
}

func (r *CreateFlowGroupByFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowGroupByFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowGroupByTemplatesRequestParams struct {
	// 调用方用户信息，userId 必填。支持填入集团子公司经办人 userId 代发合同
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同组名称,最大长度200个字符
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`

	// 合同组的子合同信息，支持2-50个子合同
	FlowGroupInfos []*FlowGroupInfo `json:"FlowGroupInfos,omitnil" name:"FlowGroupInfos"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同组的配置信息。包括是否通知本企业签署方，是否通知其他签署方
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil" name:"FlowGroupOptions"`
}

type CreateFlowGroupByTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填。支持填入集团子公司经办人 userId 代发合同
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同组名称,最大长度200个字符
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`

	// 合同组的子合同信息，支持2-50个子合同
	FlowGroupInfos []*FlowGroupInfo `json:"FlowGroupInfos,omitnil" name:"FlowGroupInfos"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同组的配置信息。包括是否通知本企业签署方，是否通知其他签署方
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil" name:"FlowGroupOptions"`
}

func (r *CreateFlowGroupByTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowGroupByTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowGroupName")
	delete(f, "FlowGroupInfos")
	delete(f, "Agent")
	delete(f, "FlowGroupOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowGroupByTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowGroupByTemplatesResponseParams struct {
	// 合同(流程)组的合同组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 合同(流程)组中子合同列表.
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFlowGroupByTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowGroupByTemplatesResponseParams `json:"Response"`
}

func (r *CreateFlowGroupByTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowGroupByTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlowOption struct {
	// 是否允许修改发起合同时确认弹窗的合同信息（合同名称、合同类型、签署截止时间），若不允许编辑，则表单字段将被禁止输入。
	// <br/>true：允许编辑（默认），<br/>false：不允许编辑<br/>默认：false：不允许编辑
	CanEditFlow *bool `json:"CanEditFlow,omitnil" name:"CanEditFlow"`

	// 是否允许编辑模板控件
	// <br/>true:允许编辑模板控件信息
	// <br/>false:不允许编辑模板控件信息
	// <br/>默认false:不允许编辑模板控件信息
	CanEditFormField *bool `json:"CanEditFormField,omitnil" name:"CanEditFormField"`

	// 发起页面隐藏合同名称展示
	// <br/>true:发起页面隐藏合同名称展示
	// <br/>false:发起页面不隐藏合同名称展示
	// <br/>默认false:发起页面不隐藏合同名称展示
	HideShowFlowName *bool `json:"HideShowFlowName,omitnil" name:"HideShowFlowName"`

	// 发起页面隐藏合同类型展示
	// <br/>true:发起页面隐藏合同类型展示
	// <br/>false:发起页面不隐藏合同类型展示
	// <br/>默认false:发起页面不隐藏合同类型展示
	HideShowFlowType *bool `json:"HideShowFlowType,omitnil" name:"HideShowFlowType"`

	// 发起页面隐藏合同截止日期展示
	// <br/>true:发起页面隐藏合同截止日期展示
	// <br/>false:发起页面不隐藏合同截止日期展示
	// <br/>默认false:发起页面不隐藏合同截止日期展示
	HideShowDeadline *bool `json:"HideShowDeadline,omitnil" name:"HideShowDeadline"`

	// 发起页面允许跳过添加签署人环节
	// <br/>true:发起页面允许跳过添加签署人环节
	// <br/>false:发起页面不允许跳过添加签署人环节
	// <br/>默认false:发起页面不允许跳过添加签署人环节
	CanSkipAddApprover *bool `json:"CanSkipAddApprover,omitnil" name:"CanSkipAddApprover"`

	// 文件发起页面跳过文件上传步骤
	// <br/>true:文件发起页面跳过文件上传步骤
	// <br/>false:文件发起页面不跳过文件上传步骤
	// <br/>默认false:文件发起页面不跳过文件上传步骤
	SkipUploadFile *bool `json:"SkipUploadFile,omitnil" name:"SkipUploadFile"`

	// 禁止编辑填写控件
	// <br/>true:禁止编辑填写控件
	// <br/>false:允许编辑填写控件
	// <br/>默认false:允许编辑填写控件
	ForbidEditFillComponent *bool `json:"ForbidEditFillComponent,omitnil" name:"ForbidEditFillComponent"`

	// 定制化发起合同弹窗的描述信息，描述信息最长500
	CustomCreateFlowDescription *string `json:"CustomCreateFlowDescription,omitnil" name:"CustomCreateFlowDescription"`
}

// Predefined struct for user
type CreateFlowRemindsRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 需执行催办的签署流程ID数组，最多包含100个。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreateFlowRemindsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 需执行催办的签署流程ID数组，最多包含100个。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreateFlowRemindsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowRemindsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowRemindsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowRemindsResponseParams struct {
	// 合同催办结果的详细信息列表。
	RemindFlowRecords []*RemindFlowRecords `json:"RemindFlowRecords,omitnil" name:"RemindFlowRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFlowRemindsResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowRemindsResponseParams `json:"Response"`
}

func (r *CreateFlowRemindsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowRemindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 
	// 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。
	// 
	// 注:  `approver中的顺序需要和模板中的顺序保持一致， 否则会导致模板中配置的信息无效`
	Approvers []*FlowCreateApprover `json:"Approvers,omitnil" name:"Approvers"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 已经废弃字段，客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	DeadLine *int64 `json:"DeadLine,omitnil" name:"DeadLine"`

	// 合同到期提醒时间，为Unix标准时间戳（秒）格式，支持的范围是从发起时间开始到后10年内。
	// 
	// 到达提醒时间后，腾讯电子签会短信通知发起方企业合同提醒，可用于处理合同到期事务，如合同续签等事宜。
	RemindedOn *int64 `json:"RemindedOn,omitnil" name:"RemindedOn"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	// 注：`请和模板中的配置保持一致`
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 您可以自定义腾讯电子签小程序合同列表页展示的合同内容模板，模板中支持以下变量：
	// <ul><li>{合同名称}   </li>
	// <li>{发起方企业} </li>
	// <li>{发起方姓名} </li>
	// <li>{签署方N企业}</li>
	// <li>{签署方N姓名}</li></ul>
	// 其中，N表示签署方的编号，从1开始，不能超过签署人的数量。
	// 
	// 例如，如果是腾讯公司张三发给李四名称为“租房合同”的合同，您可以将此字段设置为：`合同名称:{合同名称};发起方: {发起方企业}({发起方姓名});签署方:{签署方1姓名}`，则小程序中列表页展示此合同为以下样子
	// 
	// 合同名称：租房合同 
	// 发起方：腾讯公司(张三) 
	// 签署方：李四
	// 
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// 注:`此功能为白名单功能，使用前请联系对接的客户经理沟通。`
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：处方单（医疗自动签）  </li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil" name:"AutoSignScene"`

	// 暂未开放
	//
	// Deprecated: RelatedFlowId is deprecated.
	RelatedFlowId *string `json:"RelatedFlowId,omitnil" name:"RelatedFlowId"`

	// 暂未开放
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`
}

type CreateFlowRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 
	// 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。
	// 
	// 注:  `approver中的顺序需要和模板中的顺序保持一致， 否则会导致模板中配置的信息无效`
	Approvers []*FlowCreateApprover `json:"Approvers,omitnil" name:"Approvers"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 已经废弃字段，客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	DeadLine *int64 `json:"DeadLine,omitnil" name:"DeadLine"`

	// 合同到期提醒时间，为Unix标准时间戳（秒）格式，支持的范围是从发起时间开始到后10年内。
	// 
	// 到达提醒时间后，腾讯电子签会短信通知发起方企业合同提醒，可用于处理合同到期事务，如合同续签等事宜。
	RemindedOn *int64 `json:"RemindedOn,omitnil" name:"RemindedOn"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	// 注：`请和模板中的配置保持一致`
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 您可以自定义腾讯电子签小程序合同列表页展示的合同内容模板，模板中支持以下变量：
	// <ul><li>{合同名称}   </li>
	// <li>{发起方企业} </li>
	// <li>{发起方姓名} </li>
	// <li>{签署方N企业}</li>
	// <li>{签署方N姓名}</li></ul>
	// 其中，N表示签署方的编号，从1开始，不能超过签署人的数量。
	// 
	// 例如，如果是腾讯公司张三发给李四名称为“租房合同”的合同，您可以将此字段设置为：`合同名称:{合同名称};发起方: {发起方企业}({发起方姓名});签署方:{签署方1姓名}`，则小程序中列表页展示此合同为以下样子
	// 
	// 合同名称：租房合同 
	// 发起方：腾讯公司(张三) 
	// 签署方：李四
	// 
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// 注:`此功能为白名单功能，使用前请联系对接的客户经理沟通。`
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：处方单（医疗自动签）  </li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil" name:"AutoSignScene"`

	// 暂未开放
	RelatedFlowId *string `json:"RelatedFlowId,omitnil" name:"RelatedFlowId"`

	// 暂未开放
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`
}

func (r *CreateFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowName")
	delete(f, "Approvers")
	delete(f, "FlowDescription")
	delete(f, "FlowType")
	delete(f, "ClientToken")
	delete(f, "DeadLine")
	delete(f, "RemindedOn")
	delete(f, "UserData")
	delete(f, "Unordered")
	delete(f, "CustomShowMap")
	delete(f, "NeedSignReview")
	delete(f, "Agent")
	delete(f, "CcInfos")
	delete(f, "AutoSignScene")
	delete(f, "RelatedFlowId")
	delete(f, "CallbackUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowResponseParams struct {
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 
	// 注:
	// 此返回的合同流程ID，需再次调用<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateDocument" target="_blank">创建电子文档</a>和<a href="https://qian.tencent.com/developers/companyApis/startFlows/StartFlow" target="_blank">发起签署流程</a>接口将合同开始后，合同才能进入签署环节
	// 
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowResponseParams `json:"Response"`
}

func (r *CreateFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowSignReviewRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 企业内部审核结果
	// PASS: 通过 
	// REJECT: 拒绝
	ReviewType *string `json:"ReviewType,omitnil" name:"ReviewType"`

	// 审核原因 
	// 当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	ReviewMessage *string `json:"ReviewMessage,omitnil" name:"ReviewMessage"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 审核签署节点使用 非必填 如果填写则审核该签署节点。给个人审核时必填。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 操作类型：（接口通过该字段区分操作类型）
	// 
	// SignReview:签署审核
	// CreateReview:发起审核
	// 
	// 默认：SignReview；SignReview:签署审核
	// 
	// 该字段不传或者为空，则默认为SignReview签署审核，走签署审核流程
	// 若发起个人审核，则指定该字段为：SignReview
	OperateType *string `json:"OperateType,omitnil" name:"OperateType"`
}

type CreateFlowSignReviewRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 企业内部审核结果
	// PASS: 通过 
	// REJECT: 拒绝
	ReviewType *string `json:"ReviewType,omitnil" name:"ReviewType"`

	// 审核原因 
	// 当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	ReviewMessage *string `json:"ReviewMessage,omitnil" name:"ReviewMessage"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 审核签署节点使用 非必填 如果填写则审核该签署节点。给个人审核时必填。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 操作类型：（接口通过该字段区分操作类型）
	// 
	// SignReview:签署审核
	// CreateReview:发起审核
	// 
	// 默认：SignReview；SignReview:签署审核
	// 
	// 该字段不传或者为空，则默认为SignReview签署审核，走签署审核流程
	// 若发起个人审核，则指定该字段为：SignReview
	OperateType *string `json:"OperateType,omitnil" name:"OperateType"`
}

func (r *CreateFlowSignReviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowSignReviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "ReviewType")
	delete(f, "ReviewMessage")
	delete(f, "Agent")
	delete(f, "RecipientId")
	delete(f, "OperateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowSignReviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowSignReviewResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFlowSignReviewResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowSignReviewResponseParams `json:"Response"`
}

func (r *CreateFlowSignReviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowSignReviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowSignUrlRequestParams struct {
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 流程签署人列表，其中结构体的ApproverName，ApproverMobile和ApproverType必传，其他可不传，
	// 
	// 注:
	// `1. ApproverType目前只支持个人类型的签署人。`
	// `2. 签署人只能有手写签名和时间类型的签署控件，其他类型的填写控件和签署控件暂时都未支持。`
	FlowApproverInfos []*FlowCreateApprover `json:"FlowApproverInfos,omitnil" name:"FlowApproverInfos"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 机构信息，暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 签署完之后的H5页面的跳转链接，此链接及支持http://和https://，最大长度1000个字符。(建议https协议)
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`
}

type CreateFlowSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 流程签署人列表，其中结构体的ApproverName，ApproverMobile和ApproverType必传，其他可不传，
	// 
	// 注:
	// `1. ApproverType目前只支持个人类型的签署人。`
	// `2. 签署人只能有手写签名和时间类型的签署控件，其他类型的填写控件和签署控件暂时都未支持。`
	FlowApproverInfos []*FlowCreateApprover `json:"FlowApproverInfos,omitnil" name:"FlowApproverInfos"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 机构信息，暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 签署完之后的H5页面的跳转链接，此链接及支持http://和https://，最大长度1000个字符。(建议https协议)
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`
}

func (r *CreateFlowSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "FlowApproverInfos")
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "Organization")
	delete(f, "JumpUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowSignUrlResponseParams struct {
	// 签署人签署链接信息
	FlowApproverUrlInfos []*FlowApproverUrlInfo `json:"FlowApproverUrlInfos,omitnil" name:"FlowApproverUrlInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFlowSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowSignUrlResponseParams `json:"Response"`
}

func (r *CreateFlowSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationDepartmentRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 部门名称，不超过50个字符
	DeptName *string `json:"DeptName,omitnil" name:"DeptName"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 电子签父部门ID，与ParentDeptOpenId二选一,优先ParentDeptId,都为空时自动填充至根节点下
	ParentDeptId *string `json:"ParentDeptId,omitnil" name:"ParentDeptId"`

	// 第三方平台中父部门ID,与ParentDeptId二选一,优先ParentDeptId,都为空时自动填充至根节点下
	ParentDeptOpenId *string `json:"ParentDeptOpenId,omitnil" name:"ParentDeptOpenId"`

	// 客户系统部门ID，不超过64个字符
	DeptOpenId *string `json:"DeptOpenId,omitnil" name:"DeptOpenId"`

	// 排序号,1~30000范围内
	OrderNo *uint64 `json:"OrderNo,omitnil" name:"OrderNo"`
}

type CreateIntegrationDepartmentRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 部门名称，不超过50个字符
	DeptName *string `json:"DeptName,omitnil" name:"DeptName"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 电子签父部门ID，与ParentDeptOpenId二选一,优先ParentDeptId,都为空时自动填充至根节点下
	ParentDeptId *string `json:"ParentDeptId,omitnil" name:"ParentDeptId"`

	// 第三方平台中父部门ID,与ParentDeptId二选一,优先ParentDeptId,都为空时自动填充至根节点下
	ParentDeptOpenId *string `json:"ParentDeptOpenId,omitnil" name:"ParentDeptOpenId"`

	// 客户系统部门ID，不超过64个字符
	DeptOpenId *string `json:"DeptOpenId,omitnil" name:"DeptOpenId"`

	// 排序号,1~30000范围内
	OrderNo *uint64 `json:"OrderNo,omitnil" name:"OrderNo"`
}

func (r *CreateIntegrationDepartmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationDepartmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "DeptName")
	delete(f, "Agent")
	delete(f, "ParentDeptId")
	delete(f, "ParentDeptOpenId")
	delete(f, "DeptOpenId")
	delete(f, "OrderNo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationDepartmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationDepartmentResponseParams struct {
	// 电子签部门ID
	DeptId *string `json:"DeptId,omitnil" name:"DeptId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateIntegrationDepartmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationDepartmentResponseParams `json:"Response"`
}

func (r *CreateIntegrationDepartmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationEmployeesRequestParams struct {
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 待创建员工的信息，不超过20个。
	// 所有类型的企业支持的入参：Mobile和DisplayName必填,OpenId、Email和Department.DepartmentId选填，其他字段暂不支持。
	// 企微类型的企业特有支持的入参：WeworkOpenId，传入此字段无需在传入其他信息
	Employees []*Staff `json:"Employees,omitnil" name:"Employees"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreateIntegrationEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 待创建员工的信息，不超过20个。
	// 所有类型的企业支持的入参：Mobile和DisplayName必填,OpenId、Email和Department.DepartmentId选填，其他字段暂不支持。
	// 企微类型的企业特有支持的入参：WeworkOpenId，传入此字段无需在传入其他信息
	Employees []*Staff `json:"Employees,omitnil" name:"Employees"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreateIntegrationEmployeesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationEmployeesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Employees")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationEmployeesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationEmployeesResponseParams struct {
	// 创建员工的结果
	CreateEmployeeResult *CreateStaffResult `json:"CreateEmployeeResult,omitnil" name:"CreateEmployeeResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateIntegrationEmployeesResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationEmployeesResponseParams `json:"Response"`
}

func (r *CreateIntegrationEmployeesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationEmployeesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationRoleRequestParams struct {
	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 角色类型，0:saas角色，1:集团角色
	// 默认0，saas角色
	IsGroupRole *int64 `json:"IsGroupRole,omitnil" name:"IsGroupRole"`

	// 权限树
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`

	// 集团角色的话，需要传递集团子企业列表，如果是全选，则传1
	SubOrganizationIds *string `json:"SubOrganizationIds,omitnil" name:"SubOrganizationIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreateIntegrationRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 角色类型，0:saas角色，1:集团角色
	// 默认0，saas角色
	IsGroupRole *int64 `json:"IsGroupRole,omitnil" name:"IsGroupRole"`

	// 权限树
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`

	// 集团角色的话，需要传递集团子企业列表，如果是全选，则传1
	SubOrganizationIds *string `json:"SubOrganizationIds,omitnil" name:"SubOrganizationIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreateIntegrationRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Operator")
	delete(f, "Description")
	delete(f, "IsGroupRole")
	delete(f, "PermissionGroups")
	delete(f, "SubOrganizationIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationRoleResponseParams struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateIntegrationRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationRoleResponseParams `json:"Response"`
}

func (r *CreateIntegrationRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationUserRolesRequestParams struct {
	// 操作人信息，UserId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 绑定角色的用户id列表，不能重复，不能大于 100 个
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 绑定角色的角色id列表，不能重复，不能大于 100，可以通过DescribeIntegrationRoles接口获取
	RoleIds []*string `json:"RoleIds,omitnil" name:"RoleIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreateIntegrationUserRolesRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息，UserId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 绑定角色的用户id列表，不能重复，不能大于 100 个
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 绑定角色的角色id列表，不能重复，不能大于 100，可以通过DescribeIntegrationRoles接口获取
	RoleIds []*string `json:"RoleIds,omitnil" name:"RoleIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreateIntegrationUserRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationUserRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "UserIds")
	delete(f, "RoleIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationUserRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationUserRolesResponseParams struct {
	// 绑定角色失败列表信息
	FailedCreateRoleData []*FailedCreateRoleData `json:"FailedCreateRoleData,omitnil" name:"FailedCreateRoleData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateIntegrationUserRolesResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationUserRolesResponseParams `json:"Response"`
}

func (r *CreateIntegrationUserRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationUserRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiFlowSignQRCodeRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同模板ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "模板"->"模板中心"->"列表展示设置"选中模板 ID 中查看某个模板的TemplateId(在页面中展示为模板ID)。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 通过此二维码可发起的流程最大限额，如未明确指定，默认为5份。
	// 一旦发起流程数超越该限制，该二维码将自动失效。
	MaxFlowNum *int64 `json:"MaxFlowNum,omitnil" name:"MaxFlowNum"`

	// 二维码的有效期限，默认为7天，最高设定不得超过90天。
	// 一旦超过二维码的有效期限，该二维码将自动失效。
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitnil" name:"QrEffectiveDay"`

	// 合同流程的签署有效期限，若未设定签署截止日期，则默认为自合同流程创建起的7天内截止。
	// 若在签署截止日期前未完成签署，合同状态将变更为已过期，从而导致合同无效。
	// 最长设定期限不得超过30天。
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitnil" name:"FlowEffectiveDay"`

	// 指定签署人信息。
	// 在指定签署人后，仅允许特定签署人通过扫描二维码进行签署。
	Restrictions []*ApproverRestriction `json:"Restrictions,omitnil" name:"Restrictions"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 已废弃，回调配置统一使用企业应用管理-应用集成-企业版应用中的配置 
	// <br/> 通过一码多扫二维码发起的合同，回调消息可参考文档 https://qian.tencent.com/developers/company/callback_types_contracts_sign
	// <br/> 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档 https://qian.tencent.com/developers/company/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 限制二维码用户条件（已弃用）
	//
	// Deprecated: ApproverRestrictions is deprecated.
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitnil" name:"ApproverRestrictions"`
}

type CreateMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同模板ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "模板"->"模板中心"->"列表展示设置"选中模板 ID 中查看某个模板的TemplateId(在页面中展示为模板ID)。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 通过此二维码可发起的流程最大限额，如未明确指定，默认为5份。
	// 一旦发起流程数超越该限制，该二维码将自动失效。
	MaxFlowNum *int64 `json:"MaxFlowNum,omitnil" name:"MaxFlowNum"`

	// 二维码的有效期限，默认为7天，最高设定不得超过90天。
	// 一旦超过二维码的有效期限，该二维码将自动失效。
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitnil" name:"QrEffectiveDay"`

	// 合同流程的签署有效期限，若未设定签署截止日期，则默认为自合同流程创建起的7天内截止。
	// 若在签署截止日期前未完成签署，合同状态将变更为已过期，从而导致合同无效。
	// 最长设定期限不得超过30天。
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitnil" name:"FlowEffectiveDay"`

	// 指定签署人信息。
	// 在指定签署人后，仅允许特定签署人通过扫描二维码进行签署。
	Restrictions []*ApproverRestriction `json:"Restrictions,omitnil" name:"Restrictions"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 已废弃，回调配置统一使用企业应用管理-应用集成-企业版应用中的配置 
	// <br/> 通过一码多扫二维码发起的合同，回调消息可参考文档 https://qian.tencent.com/developers/company/callback_types_contracts_sign
	// <br/> 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档 https://qian.tencent.com/developers/company/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 限制二维码用户条件（已弃用）
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitnil" name:"ApproverRestrictions"`
}

func (r *CreateMultiFlowSignQRCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiFlowSignQRCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "TemplateId")
	delete(f, "FlowName")
	delete(f, "MaxFlowNum")
	delete(f, "QrEffectiveDay")
	delete(f, "FlowEffectiveDay")
	delete(f, "Restrictions")
	delete(f, "UserData")
	delete(f, "CallbackUrl")
	delete(f, "Agent")
	delete(f, "ApproverRestrictions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultiFlowSignQRCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiFlowSignQRCodeResponseParams struct {
	// 签署二维码的基本信息，用于创建二维码，用户可扫描该二维码进行签署操作。
	QrCode *SignQrCode `json:"QrCode,omitnil" name:"QrCode"`

	// 流程签署二维码的签署信息，适用于客户系统整合二维码功能。通过链接，用户可直接访问电子签名小程序并签署合同。
	SignUrls *SignUrl `json:"SignUrls,omitnil" name:"SignUrls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateMultiFlowSignQRCodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultiFlowSignQRCodeResponseParams `json:"Response"`
}

func (r *CreateMultiFlowSignQRCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiFlowSignQRCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationBatchSignUrlRequestParams struct {
	// 调用方用户信息，UserId 必填，支持填入集团子公司经办人UserId。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 指定需要进行批量签署的流程id，数量1-100，填写后用户将通过链接对这些合同进行批量签署。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 员工的UserId，该UserId对应的员工必须已经加入企业并实名，Name和Mobile为空时该字段不能为空。（优先使用UserId对应的员工）
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 员工姓名，该字段需要与Mobile组合使用，UserId为空时该字段不能为空。（UserId为空时，使用Name和Mbile对应的员工）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 员工手机号码，该字段需要与Name组合使用，UserId为空时该字段不能为空。（UserId为空时，使用Name和Mbile对应的员工）
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`
}

type CreateOrganizationBatchSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，UserId 必填，支持填入集团子公司经办人UserId。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 指定需要进行批量签署的流程id，数量1-100，填写后用户将通过链接对这些合同进行批量签署。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 员工的UserId，该UserId对应的员工必须已经加入企业并实名，Name和Mobile为空时该字段不能为空。（优先使用UserId对应的员工）
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 员工姓名，该字段需要与Mobile组合使用，UserId为空时该字段不能为空。（UserId为空时，使用Name和Mbile对应的员工）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 员工手机号码，该字段需要与Name组合使用，UserId为空时该字段不能为空。（UserId为空时，使用Name和Mbile对应的员工）
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`
}

func (r *CreateOrganizationBatchSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationBatchSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "Agent")
	delete(f, "UserId")
	delete(f, "Name")
	delete(f, "Mobile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationBatchSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationBatchSignUrlResponseParams struct {
	// 批量签署入口链接
	SignUrl *string `json:"SignUrl,omitnil" name:"SignUrl"`

	// 链接过期时间戳
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateOrganizationBatchSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationBatchSignUrlResponseParams `json:"Response"`
}

func (r *CreateOrganizationBatchSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationBatchSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonAuthCertificateImageRequestParams struct {
	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 个人用户名称
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 身份证件类型取值：
	// ID_CARD 身居民身份证
	// PASSPORT 护照
	// HONGKONG_AND_MACAO 港澳居民来往内地通行证
	// FOREIGN_ID_CARD 外国人永久居留身份证
	// HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreatePersonAuthCertificateImageRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 个人用户名称
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 身份证件类型取值：
	// ID_CARD 身居民身份证
	// PASSPORT 护照
	// HONGKONG_AND_MACAO 港澳居民来往内地通行证
	// FOREIGN_ID_CARD 外国人永久居留身份证
	// HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreatePersonAuthCertificateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonAuthCertificateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "UserName")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonAuthCertificateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonAuthCertificateImageResponseParams struct {
	// 个人用户证明证书的下载链接
	AuthCertUrl *string `json:"AuthCertUrl,omitnil" name:"AuthCertUrl"`

	// 证书图片上的证书编号，20位数字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageCertId *string `json:"ImageCertId,omitnil" name:"ImageCertId"`

	// 图片证明对应的CA证书序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialNumber *string `json:"SerialNumber,omitnil" name:"SerialNumber"`

	// CA证书颁发时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidFrom *uint64 `json:"ValidFrom,omitnil" name:"ValidFrom"`

	// CA证书有效截止时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidTo *uint64 `json:"ValidTo,omitnil" name:"ValidTo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePersonAuthCertificateImageResponse struct {
	*tchttp.BaseResponse
	Response *CreatePersonAuthCertificateImageResponseParams `json:"Response"`
}

func (r *CreatePersonAuthCertificateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonAuthCertificateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrepareFlowRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 资源id，与ResourceType对应
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 合同名称
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 是否顺序签署
	// true:无序签
	// false:顺序签
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 签署流程的签署截止时间。
	// 值为unix时间戳,精确到秒
	// 不传默认为当前时间一年后
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 用户自定义合同类型Id
	// 
	// 该id为电子签企业内的合同类型id， 可以在自定义合同类型处获取
	UserFlowTypeId *string `json:"UserFlowTypeId,omitnil" name:"UserFlowTypeId"`

	// 合同类型名称
	// 该字段用于客户自定义合同类型
	// 建议使用时指定合同类型，便于之后合同分类以及查看
	// 如果合同类型与自定义的合同类型描述一致，会自动归类到自定义的合同类型处，如果不一致，则会创建一个新的自定义合同类型
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 签署流程参与者信息，最大限制50方
	Approvers []*FlowCreateApprover `json:"Approvers,omitnil" name:"Approvers"`

	// 打开智能添加填写区
	// 默认开启，打开:"OPEN"
	//  关闭："CLOSE"
	IntelligentStatus *string `json:"IntelligentStatus,omitnil" name:"IntelligentStatus"`

	// 资源类型，
	// 1：模板
	// 2：文件，
	// 不传默认为2：文件
	ResourceType *int64 `json:"ResourceType,omitnil" name:"ResourceType"`

	// 发起方填写控件
	// 该类型控件由发起方完成填写
	Components *Component `json:"Components,omitnil" name:"Components"`

	// 发起合同个性化参数
	// 用于满足创建及页面操作过程中的个性化要求
	// 具体定制化内容详见数据接口说明
	FlowOption *CreateFlowOption `json:"FlowOption,omitnil" name:"FlowOption"`

	// 是否开启发起方签署审核
	// true:开启发起方签署审核
	// false:不开启发起方签署审核
	// 默认false:不开启发起方签署审核
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 开启发起方发起合同审核
	// true:开启发起方发起合同审核
	// false:不开启发起方发起合同审核
	// 默认false:不开启发起方发起合同审核
	NeedCreateReview *bool `json:"NeedCreateReview,omitnil" name:"NeedCreateReview"`

	// 用户自定义参数
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 合同id,用于通过已web页面发起的合同id快速生成一个web发起合同链接
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填	
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreatePrepareFlowRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 资源id，与ResourceType对应
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 合同名称
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 是否顺序签署
	// true:无序签
	// false:顺序签
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 签署流程的签署截止时间。
	// 值为unix时间戳,精确到秒
	// 不传默认为当前时间一年后
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 用户自定义合同类型Id
	// 
	// 该id为电子签企业内的合同类型id， 可以在自定义合同类型处获取
	UserFlowTypeId *string `json:"UserFlowTypeId,omitnil" name:"UserFlowTypeId"`

	// 合同类型名称
	// 该字段用于客户自定义合同类型
	// 建议使用时指定合同类型，便于之后合同分类以及查看
	// 如果合同类型与自定义的合同类型描述一致，会自动归类到自定义的合同类型处，如果不一致，则会创建一个新的自定义合同类型
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 签署流程参与者信息，最大限制50方
	Approvers []*FlowCreateApprover `json:"Approvers,omitnil" name:"Approvers"`

	// 打开智能添加填写区
	// 默认开启，打开:"OPEN"
	//  关闭："CLOSE"
	IntelligentStatus *string `json:"IntelligentStatus,omitnil" name:"IntelligentStatus"`

	// 资源类型，
	// 1：模板
	// 2：文件，
	// 不传默认为2：文件
	ResourceType *int64 `json:"ResourceType,omitnil" name:"ResourceType"`

	// 发起方填写控件
	// 该类型控件由发起方完成填写
	Components *Component `json:"Components,omitnil" name:"Components"`

	// 发起合同个性化参数
	// 用于满足创建及页面操作过程中的个性化要求
	// 具体定制化内容详见数据接口说明
	FlowOption *CreateFlowOption `json:"FlowOption,omitnil" name:"FlowOption"`

	// 是否开启发起方签署审核
	// true:开启发起方签署审核
	// false:不开启发起方签署审核
	// 默认false:不开启发起方签署审核
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 开启发起方发起合同审核
	// true:开启发起方发起合同审核
	// false:不开启发起方发起合同审核
	// 默认false:不开启发起方发起合同审核
	NeedCreateReview *bool `json:"NeedCreateReview,omitnil" name:"NeedCreateReview"`

	// 用户自定义参数
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 合同id,用于通过已web页面发起的合同id快速生成一个web发起合同链接
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填	
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreatePrepareFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrepareFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ResourceId")
	delete(f, "FlowName")
	delete(f, "Unordered")
	delete(f, "Deadline")
	delete(f, "UserFlowTypeId")
	delete(f, "FlowType")
	delete(f, "Approvers")
	delete(f, "IntelligentStatus")
	delete(f, "ResourceType")
	delete(f, "Components")
	delete(f, "FlowOption")
	delete(f, "NeedSignReview")
	delete(f, "NeedCreateReview")
	delete(f, "UserData")
	delete(f, "FlowId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrepareFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrepareFlowResponseParams struct {
	// 快速发起预览链接，有效期5分钟
	Url *string `json:"Url,omitnil" name:"Url"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePrepareFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrepareFlowResponseParams `json:"Response"`
}

func (r *CreatePrepareFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrepareFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePreparedPersonalEsignRequestParams struct {
	// 个人用户姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 印章名称
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 调用方用户信息，userId 必填。支持填入集团子公司经办人 userId代发合同。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 身份证件类型:
	// ID_CARD 身份证
	// PASSPORT 护照
	// HONGKONG_AND_MACAO 中国香港
	// FOREIGN_ID_CARD 境外身份
	// HONGKONG_MACAO_AND_TAIWAN 中国台湾
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 印章图片的base64
	// 注：已废弃
	// 请先通过UploadFiles接口上传文件，获取 FileId
	//
	// Deprecated: SealImage is deprecated.
	SealImage *string `json:"SealImage,omitnil" name:"SealImage"`

	// 是否开启印章图片压缩处理，默认不开启，如需开启请设置为 true。当印章超过 2M 时建议开启，开启后图片的 hash 将发生变化。
	SealImageCompress *bool `json:"SealImageCompress,omitnil" name:"SealImageCompress"`

	// 手机号码；当需要开通自动签时，该参数必传
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 是否开通自动签，该功能需联系运营工作人员开通后使用
	EnableAutoSign *bool `json:"EnableAutoSign,omitnil" name:"EnableAutoSign"`

	// 印章颜色（参数ProcessSeal=true时生效）
	// 默认值：BLACK黑色
	// 取值: 
	// BLACK 黑色,
	// RED 红色,
	// BLUE 蓝色。
	SealColor *string `json:"SealColor,omitnil" name:"SealColor"`

	// 是否处理印章
	// 默认不做印章处理。
	// 取值：false：不做任何处理；
	// true：做透明化处理和颜色增强。
	ProcessSeal *bool `json:"ProcessSeal,omitnil" name:"ProcessSeal"`

	// 印章图片文件 id
	// 取值：
	// 填写的FileId通过UploadFiles接口上传文件获取。
	FileId *string `json:"FileId,omitnil" name:"FileId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次 1-不绑定，发起合同时将按标准合同套餐进行扣减	
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`
}

type CreatePreparedPersonalEsignRequest struct {
	*tchttp.BaseRequest
	
	// 个人用户姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 印章名称
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 调用方用户信息，userId 必填。支持填入集团子公司经办人 userId代发合同。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 身份证件类型:
	// ID_CARD 身份证
	// PASSPORT 护照
	// HONGKONG_AND_MACAO 中国香港
	// FOREIGN_ID_CARD 境外身份
	// HONGKONG_MACAO_AND_TAIWAN 中国台湾
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 印章图片的base64
	// 注：已废弃
	// 请先通过UploadFiles接口上传文件，获取 FileId
	SealImage *string `json:"SealImage,omitnil" name:"SealImage"`

	// 是否开启印章图片压缩处理，默认不开启，如需开启请设置为 true。当印章超过 2M 时建议开启，开启后图片的 hash 将发生变化。
	SealImageCompress *bool `json:"SealImageCompress,omitnil" name:"SealImageCompress"`

	// 手机号码；当需要开通自动签时，该参数必传
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 是否开通自动签，该功能需联系运营工作人员开通后使用
	EnableAutoSign *bool `json:"EnableAutoSign,omitnil" name:"EnableAutoSign"`

	// 印章颜色（参数ProcessSeal=true时生效）
	// 默认值：BLACK黑色
	// 取值: 
	// BLACK 黑色,
	// RED 红色,
	// BLUE 蓝色。
	SealColor *string `json:"SealColor,omitnil" name:"SealColor"`

	// 是否处理印章
	// 默认不做印章处理。
	// 取值：false：不做任何处理；
	// true：做透明化处理和颜色增强。
	ProcessSeal *bool `json:"ProcessSeal,omitnil" name:"ProcessSeal"`

	// 印章图片文件 id
	// 取值：
	// 填写的FileId通过UploadFiles接口上传文件获取。
	FileId *string `json:"FileId,omitnil" name:"FileId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次 1-不绑定，发起合同时将按标准合同套餐进行扣减	
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`
}

func (r *CreatePreparedPersonalEsignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePreparedPersonalEsignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "IdCardNumber")
	delete(f, "SealName")
	delete(f, "Operator")
	delete(f, "IdCardType")
	delete(f, "SealImage")
	delete(f, "SealImageCompress")
	delete(f, "Mobile")
	delete(f, "EnableAutoSign")
	delete(f, "SealColor")
	delete(f, "ProcessSeal")
	delete(f, "FileId")
	delete(f, "Agent")
	delete(f, "LicenseType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePreparedPersonalEsignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePreparedPersonalEsignResponseParams struct {
	// 导入生成的印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePreparedPersonalEsignResponse struct {
	*tchttp.BaseResponse
	Response *CreatePreparedPersonalEsignResponseParams `json:"Response"`
}

func (r *CreatePreparedPersonalEsignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePreparedPersonalEsignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseFlowRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 待解除的签署流程编号（即原签署流程的编号）。
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitnil" name:"NeedRelievedFlowId"`

	// 解除协议内容。
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitnil" name:"ReliveInfo"`

	// 关于渠道应用的相关信息，包括子客企业及应用编、号等详细内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 解除协议的签署人列表(如不指定该参数，默认使用原流程的签署人列表)。 <br/>
	// 如需更换原合同中的签署人，可通过指定该签署人的RecipientId编号更换此签署人。(可通过接口<a href="https://qian.tencent.com/developers/companyApis/queryFlows/DescribeFlowInfo/">DescribeFlowInfo</a>查询签署人的RecipientId编号)<br/>
	// 解除协议的签署人数量不能多于原流程的签署人数量。<br/>
	// 
	// `注意：只能更换同企业的签署人。`<br/>
	// `注意：不支持更换个人类型的签署人。`<br/>
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitnil" name:"ReleasedApprovers"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的7天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 调用方自定义的个性化字段，该字段的值可以是字符串JSON或其他字符串形式，客户可以根据自身需求自定义数据格式并在需要时进行解析。该字段的信息将以Base64编码的形式传输，支持的最大数据大小为20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`
}

type CreateReleaseFlowRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 待解除的签署流程编号（即原签署流程的编号）。
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitnil" name:"NeedRelievedFlowId"`

	// 解除协议内容。
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitnil" name:"ReliveInfo"`

	// 关于渠道应用的相关信息，包括子客企业及应用编、号等详细内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 解除协议的签署人列表(如不指定该参数，默认使用原流程的签署人列表)。 <br/>
	// 如需更换原合同中的签署人，可通过指定该签署人的RecipientId编号更换此签署人。(可通过接口<a href="https://qian.tencent.com/developers/companyApis/queryFlows/DescribeFlowInfo/">DescribeFlowInfo</a>查询签署人的RecipientId编号)<br/>
	// 解除协议的签署人数量不能多于原流程的签署人数量。<br/>
	// 
	// `注意：只能更换同企业的签署人。`<br/>
	// `注意：不支持更换个人类型的签署人。`<br/>
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitnil" name:"ReleasedApprovers"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的7天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 调用方自定义的个性化字段，该字段的值可以是字符串JSON或其他字符串形式，客户可以根据自身需求自定义数据格式并在需要时进行解析。该字段的信息将以Base64编码的形式传输，支持的最大数据大小为20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`
}

func (r *CreateReleaseFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "NeedRelievedFlowId")
	delete(f, "ReliveInfo")
	delete(f, "Agent")
	delete(f, "ReleasedApprovers")
	delete(f, "Deadline")
	delete(f, "UserData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReleaseFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseFlowResponseParams struct {
	// 解除协议流程编号
	// `注意：这里的流程编号对应的合同是本次发起的解除协议。`
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateReleaseFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateReleaseFlowResponseParams `json:"Response"`
}

func (r *CreateReleaseFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchemeUrlRequestParams struct {
	// 执行本接口操作的员工信息, userId 必填。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程签署方的组织机构名称。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 合同流程里边签署方经办人的姓名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 合同流程里边签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 要跳转的链接类型
	// 
	// <ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li>
	// <li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li>
	// <li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li></ul>
	EndPoint *string `json:"EndPoint,omitnil" name:"EndPoint"`

	// 合同流程ID 
	// 注: `如果准备跳转到合同流程签署的详情页面(即PathType=1时)必传,   跳转其他页面可不传`
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 合同流程组的组ID, 在合同流程组场景下，生成合同流程组的签署链接时需要赋值
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 要跳转到的页面类型 
	// 
	// <ul><li> **0** : 腾讯电子签小程序个人首页 (默认)</li>
	// <li> **1** : 腾讯电子签小程序流程合同的详情页 (即合同签署页面)</li>
	// <li> **2** : 腾讯电子签小程序合同列表页</li></ul>
	PathType *uint64 `json:"PathType,omitnil" name:"PathType"`

	// 签署完成后是否自动回跳
	// <ul><li>**false**：否, 签署完成不会自动跳转回来(默认)</li><li>**true**：是, 签署完成会自动跳转回来</li></ul>
	// 注:  ` 该参数只针对"APP" 类型的签署链接有效`
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil" name:"AutoJumpBack"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 生成的签署链接在签署页面隐藏的按钮列表，可设置如下：
	// 
	// <ul><li> **0** :合同签署页面更多操作按钮</li>
	// <li> **1** :合同签署页面更多操作的拒绝签署按钮</li>
	// <li> **2** :合同签署页面更多操作的转他人处理按钮</li>
	// <li> **3** :签署成功页的查看详情按钮</li></ul>
	// 
	// 注:  `字段为数组, 可以传值隐藏多个按钮`
	Hides []*int64 `json:"Hides,omitnil" name:"Hides"`
}

type CreateSchemeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息, userId 必填。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程签署方的组织机构名称。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 合同流程里边签署方经办人的姓名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 合同流程里边签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 要跳转的链接类型
	// 
	// <ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li>
	// <li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li>
	// <li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li></ul>
	EndPoint *string `json:"EndPoint,omitnil" name:"EndPoint"`

	// 合同流程ID 
	// 注: `如果准备跳转到合同流程签署的详情页面(即PathType=1时)必传,   跳转其他页面可不传`
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 合同流程组的组ID, 在合同流程组场景下，生成合同流程组的签署链接时需要赋值
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 要跳转到的页面类型 
	// 
	// <ul><li> **0** : 腾讯电子签小程序个人首页 (默认)</li>
	// <li> **1** : 腾讯电子签小程序流程合同的详情页 (即合同签署页面)</li>
	// <li> **2** : 腾讯电子签小程序合同列表页</li></ul>
	PathType *uint64 `json:"PathType,omitnil" name:"PathType"`

	// 签署完成后是否自动回跳
	// <ul><li>**false**：否, 签署完成不会自动跳转回来(默认)</li><li>**true**：是, 签署完成会自动跳转回来</li></ul>
	// 注:  ` 该参数只针对"APP" 类型的签署链接有效`
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil" name:"AutoJumpBack"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 生成的签署链接在签署页面隐藏的按钮列表，可设置如下：
	// 
	// <ul><li> **0** :合同签署页面更多操作按钮</li>
	// <li> **1** :合同签署页面更多操作的拒绝签署按钮</li>
	// <li> **2** :合同签署页面更多操作的转他人处理按钮</li>
	// <li> **3** :签署成功页的查看详情按钮</li></ul>
	// 
	// 注:  `字段为数组, 可以传值隐藏多个按钮`
	Hides []*int64 `json:"Hides,omitnil" name:"Hides"`
}

func (r *CreateSchemeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchemeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "OrganizationName")
	delete(f, "Name")
	delete(f, "Mobile")
	delete(f, "EndPoint")
	delete(f, "FlowId")
	delete(f, "FlowGroupId")
	delete(f, "PathType")
	delete(f, "AutoJumpBack")
	delete(f, "Agent")
	delete(f, "Hides")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchemeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchemeUrlResponseParams struct {
	// 腾讯电子签小程序的签署链接。
	// 
	// <ul><li>如果EndPoint是**APP**，得到的链接类似于`pages/guide?from=default&where=mini&id=yDwJSUUirqauh***7jNSxwdirTSGuH&to=CONTRACT_DETAIL&name=&phone=&shortKey=yDw***k1xFc5`, 用法可以参加接口描述中的"跳转到小程序的实现"</li>
	// <li>如果EndPoint是**HTTP**，得到的链接类似于 `https://res.ess.tencent.cn/cdn/h5-activity/jump-mp.html?where=mini&from=SFY&id=yDwfEUUw**4rV6Avz&to=MVP_CONTRACT_COVER&name=%E9%83%**5%86%9B`，点击后会跳转到腾讯电子签小程序进行签署</li>
	// <li>如果EndPoint是**HTTP_SHORT_URL**，得到的链接类似于 `https://essurl.cn/2n**42Nd`，点击后会跳转到腾讯电子签小程序进行签署</li></ul>
	SchemeUrl *string `json:"SchemeUrl,omitnil" name:"SchemeUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSchemeUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateSchemeUrlResponseParams `json:"Response"`
}

func (r *CreateSchemeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchemeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealPolicyRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 用户在电子文件签署平台标识信息，具体参考UserInfo结构体。可跟下面的UserIds可叠加起作用
	Users []*UserInfo `json:"Users,omitnil" name:"Users"`

	// 印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 授权有效期。时间戳秒级
	Expired *int64 `json:"Expired,omitnil" name:"Expired"`

	// 需要授权的用户UserId集合。跟上面的SealId参数配合使用。选填，跟上面的Users同时起作用
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 印章授权内容
	Policy *string `json:"Policy,omitnil" name:"Policy"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreateSealPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 用户在电子文件签署平台标识信息，具体参考UserInfo结构体。可跟下面的UserIds可叠加起作用
	Users []*UserInfo `json:"Users,omitnil" name:"Users"`

	// 印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 授权有效期。时间戳秒级
	Expired *int64 `json:"Expired,omitnil" name:"Expired"`

	// 需要授权的用户UserId集合。跟上面的SealId参数配合使用。选填，跟上面的Users同时起作用
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 印章授权内容
	Policy *string `json:"Policy,omitnil" name:"Policy"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreateSealPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Users")
	delete(f, "SealId")
	delete(f, "Expired")
	delete(f, "UserIds")
	delete(f, "Policy")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSealPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealPolicyResponseParams struct {
	// 最终授权成功的。其他的跳过的是已经授权了的
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSealPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateSealPolicyResponseParams `json:"Response"`
}

func (r *CreateSealPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子印章名字，1-50个中文字符。
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 本接口支持上传图片印章及系统直接生成印章；
	// 如果要使用系统生成印章，此值传：SealGenerateSourceSystem；
	// 如果要使用图片上传请传字段 Image
	GenerateSource *string `json:"GenerateSource,omitnil" name:"GenerateSource"`

	// 电子印章类型：
	// OFFICIAL-公章；
	// CONTRACT-合同专用章;
	// FINANCE-合财务专用章;
	// PERSONNEL-人事专用章.
	SealType *string `json:"SealType,omitnil" name:"SealType"`

	// 电子印章图片文件名称，1-50个中文字符。
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 电子印章图片base64编码
	// 参数Image,FileToken或GenerateSource=SealGenerateSourceSystem三选一。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 电子印章宽度,单位px
	// 参数不再启用，系统会设置印章大小为标准尺寸。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 电子印章高度,单位px
	// 参数不再启用，系统会设置印章大小为标准尺寸。
	Height *int64 `json:"Height,omitnil" name:"Height"`

	// 电子印章印章颜色(默认红色RED),RED-红色
	// 
	// 系统目前只支持红色印章创建。
	Color *string `json:"Color,omitnil" name:"Color"`

	// 企业印章横向文字，最多可填15个汉字（若超过印章最大宽度，优先压缩字间距，其次缩小字号）
	SealHorizontalText *string `json:"SealHorizontalText,omitnil" name:"SealHorizontalText"`

	// 暂时不支持下弦文字设置
	SealChordText *string `json:"SealChordText,omitnil" name:"SealChordText"`

	// 系统生成的印章只支持STAR
	SealCentralType *string `json:"SealCentralType,omitnil" name:"SealCentralType"`

	// 通过文件上传时，服务端生成的电子印章上传图片的token
	FileToken *string `json:"FileToken,omitnil" name:"FileToken"`

	// 印章样式:
	// 
	// cycle:圆形印章;
	// ellipse:椭圆印章;
	// 注：默认圆形印章
	SealStyle *string `json:"SealStyle,omitnil" name:"SealStyle"`

	// 印章尺寸取值描述：
	// 42_42 圆形企业公章直径42mm；
	// 40_40 圆形企业印章直径40mm；
	// 45_30 椭圆形印章45mm x 30mm;
	SealSize *string `json:"SealSize,omitnil" name:"SealSize"`
}

type CreateSealRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子印章名字，1-50个中文字符。
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 本接口支持上传图片印章及系统直接生成印章；
	// 如果要使用系统生成印章，此值传：SealGenerateSourceSystem；
	// 如果要使用图片上传请传字段 Image
	GenerateSource *string `json:"GenerateSource,omitnil" name:"GenerateSource"`

	// 电子印章类型：
	// OFFICIAL-公章；
	// CONTRACT-合同专用章;
	// FINANCE-合财务专用章;
	// PERSONNEL-人事专用章.
	SealType *string `json:"SealType,omitnil" name:"SealType"`

	// 电子印章图片文件名称，1-50个中文字符。
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 电子印章图片base64编码
	// 参数Image,FileToken或GenerateSource=SealGenerateSourceSystem三选一。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 电子印章宽度,单位px
	// 参数不再启用，系统会设置印章大小为标准尺寸。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 电子印章高度,单位px
	// 参数不再启用，系统会设置印章大小为标准尺寸。
	Height *int64 `json:"Height,omitnil" name:"Height"`

	// 电子印章印章颜色(默认红色RED),RED-红色
	// 
	// 系统目前只支持红色印章创建。
	Color *string `json:"Color,omitnil" name:"Color"`

	// 企业印章横向文字，最多可填15个汉字（若超过印章最大宽度，优先压缩字间距，其次缩小字号）
	SealHorizontalText *string `json:"SealHorizontalText,omitnil" name:"SealHorizontalText"`

	// 暂时不支持下弦文字设置
	SealChordText *string `json:"SealChordText,omitnil" name:"SealChordText"`

	// 系统生成的印章只支持STAR
	SealCentralType *string `json:"SealCentralType,omitnil" name:"SealCentralType"`

	// 通过文件上传时，服务端生成的电子印章上传图片的token
	FileToken *string `json:"FileToken,omitnil" name:"FileToken"`

	// 印章样式:
	// 
	// cycle:圆形印章;
	// ellipse:椭圆印章;
	// 注：默认圆形印章
	SealStyle *string `json:"SealStyle,omitnil" name:"SealStyle"`

	// 印章尺寸取值描述：
	// 42_42 圆形企业公章直径42mm；
	// 40_40 圆形企业印章直径40mm；
	// 45_30 椭圆形印章45mm x 30mm;
	SealSize *string `json:"SealSize,omitnil" name:"SealSize"`
}

func (r *CreateSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SealName")
	delete(f, "Agent")
	delete(f, "GenerateSource")
	delete(f, "SealType")
	delete(f, "FileName")
	delete(f, "Image")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "Color")
	delete(f, "SealHorizontalText")
	delete(f, "SealChordText")
	delete(f, "SealCentralType")
	delete(f, "FileToken")
	delete(f, "SealStyle")
	delete(f, "SealSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealResponseParams struct {
	// 电子印章编号
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSealResponse struct {
	*tchttp.BaseResponse
	Response *CreateSealResponseParams `json:"Response"`
}

func (r *CreateSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStaffResult struct {
	// 创建员工的成功列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessEmployeeData []*SuccessCreateStaffData `json:"SuccessEmployeeData,omitnil" name:"SuccessEmployeeData"`

	// 创建员工的失败列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedEmployeeData []*FailedCreateStaffData `json:"FailedEmployeeData,omitnil" name:"FailedEmployeeData"`
}

// Predefined struct for user
type CreateUserAutoSignEnableUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li></ul>
	// 
	// 注: `现在仅支持电子处方场景`
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 自动签开通配置信息, 包括开通的人员的信息等
	AutoSignConfig *AutoSignConfig `json:"AutoSignConfig,omitnil" name:"AutoSignConfig"`

	// 生成的链接类型：
	// <ul><li> 不传(即为空值) 则会生成小程序端开通链接(默认)</li>
	// <li> **H5SIGN** : 生成H5端开通链接</li></ul>
	UrlType *string `json:"UrlType,omitnil" name:"UrlType"`

	// 是否通知开通方，通知类型:
	// <ul><li>默认不设置为不通知开通方</li>
	// <li>**SMS** :  短信通知 ,如果需要短信通知则NotifyAddress填写对方的手机号</li><ul>
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 如果通知类型NotifyType选择为SMS，则此处为手机号, 其他通知类型不需要设置此项
	NotifyAddress *string `json:"NotifyAddress,omitnil" name:"NotifyAddress"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreateUserAutoSignEnableUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li></ul>
	// 
	// 注: `现在仅支持电子处方场景`
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 自动签开通配置信息, 包括开通的人员的信息等
	AutoSignConfig *AutoSignConfig `json:"AutoSignConfig,omitnil" name:"AutoSignConfig"`

	// 生成的链接类型：
	// <ul><li> 不传(即为空值) 则会生成小程序端开通链接(默认)</li>
	// <li> **H5SIGN** : 生成H5端开通链接</li></ul>
	UrlType *string `json:"UrlType,omitnil" name:"UrlType"`

	// 是否通知开通方，通知类型:
	// <ul><li>默认不设置为不通知开通方</li>
	// <li>**SMS** :  短信通知 ,如果需要短信通知则NotifyAddress填写对方的手机号</li><ul>
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 如果通知类型NotifyType选择为SMS，则此处为手机号, 其他通知类型不需要设置此项
	NotifyAddress *string `json:"NotifyAddress,omitnil" name:"NotifyAddress"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreateUserAutoSignEnableUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserAutoSignEnableUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "AutoSignConfig")
	delete(f, "UrlType")
	delete(f, "NotifyType")
	delete(f, "NotifyAddress")
	delete(f, "ExpiredTime")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserAutoSignEnableUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserAutoSignEnableUrlResponseParams struct {
	// 个人用户自动签的开通链接, 短链形式
	Url *string `json:"Url,omitnil" name:"Url"`

	// 腾讯电子签小程序的 AppID，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用
	// 
	// 注: `如果获取的是H5链接, 则不会返回此值`
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 腾讯电子签小程序的原始 Id,  ，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用
	// 
	// 注: `如果获取的是H5链接, 则不会返回此值`
	AppOriginalId *string `json:"AppOriginalId,omitnil" name:"AppOriginalId"`

	// 腾讯电子签小程序的跳转路径，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用
	// 
	// 注: `如果获取的是H5链接, 则不会返回此值`
	Path *string `json:"Path,omitnil" name:"Path"`

	// base64 格式的跳转二维码图片，可通过微信扫描后跳转到腾讯电子签小程序的开通界面。
	// 
	// 注: `如果获取的是H5链接, 则不会返回此二维码图片`
	QrCode *string `json:"QrCode,omitnil" name:"QrCode"`

	// 返回的链接类型
	// <ul><li> 空: 默认小程序端链接</li>
	// <li> **H5SIGN** : h5端链接</li></ul>
	UrlType *string `json:"UrlType,omitnil" name:"UrlType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateUserAutoSignEnableUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserAutoSignEnableUrlResponseParams `json:"Response"`
}

func (r *CreateUserAutoSignEnableUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserAutoSignEnableUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebThemeConfigRequestParams struct {
	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 主题类型
	// <br/>EMBED_WEB_THEME：嵌入式主题
	// <br/>目前只支持EMBED_WEB_THEME，web页面嵌入的主题风格配置
	ThemeType *string `json:"ThemeType,omitnil" name:"ThemeType"`

	// 主题配置
	WebThemeConfig *WebThemeConfig `json:"WebThemeConfig,omitnil" name:"WebThemeConfig"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。	
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type CreateWebThemeConfigRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 主题类型
	// <br/>EMBED_WEB_THEME：嵌入式主题
	// <br/>目前只支持EMBED_WEB_THEME，web页面嵌入的主题风格配置
	ThemeType *string `json:"ThemeType,omitnil" name:"ThemeType"`

	// 主题配置
	WebThemeConfig *WebThemeConfig `json:"WebThemeConfig,omitnil" name:"WebThemeConfig"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。	
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *CreateWebThemeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebThemeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ThemeType")
	delete(f, "WebThemeConfig")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWebThemeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebThemeConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWebThemeConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateWebThemeConfigResponseParams `json:"Response"`
}

func (r *CreateWebThemeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebThemeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationDepartmentRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子签中的部门id,通过DescribeIntegrationDepartments接口可获得
	DeptId *string `json:"DeptId,omitnil" name:"DeptId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 交接部门ID。待删除部门中的合同、印章和模板数据，交接至该部门ID下，未填写交接至公司根部门。
	ReceiveDeptId *string `json:"ReceiveDeptId,omitnil" name:"ReceiveDeptId"`
}

type DeleteIntegrationDepartmentRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子签中的部门id,通过DescribeIntegrationDepartments接口可获得
	DeptId *string `json:"DeptId,omitnil" name:"DeptId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 交接部门ID。待删除部门中的合同、印章和模板数据，交接至该部门ID下，未填写交接至公司根部门。
	ReceiveDeptId *string `json:"ReceiveDeptId,omitnil" name:"ReceiveDeptId"`
}

func (r *DeleteIntegrationDepartmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationDepartmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "DeptId")
	delete(f, "Agent")
	delete(f, "ReceiveDeptId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntegrationDepartmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationDepartmentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteIntegrationDepartmentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntegrationDepartmentResponseParams `json:"Response"`
}

func (r *DeleteIntegrationDepartmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationEmployeesRequestParams struct {
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 待移除员工的信息，userId和openId二选一，必填一个，如果需要指定交接人的话，ReceiveUserId或者ReceiveOpenId字段二选一
	Employees []*Staff `json:"Employees,omitnil" name:"Employees"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId需填充子企业Id
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DeleteIntegrationEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 待移除员工的信息，userId和openId二选一，必填一个，如果需要指定交接人的话，ReceiveUserId或者ReceiveOpenId字段二选一
	Employees []*Staff `json:"Employees,omitnil" name:"Employees"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId需填充子企业Id
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DeleteIntegrationEmployeesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationEmployeesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Employees")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntegrationEmployeesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationEmployeesResponseParams struct {
	// 员工删除数据
	DeleteEmployeeResult *DeleteStaffsResult `json:"DeleteEmployeeResult,omitnil" name:"DeleteEmployeeResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteIntegrationEmployeesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntegrationEmployeesResponseParams `json:"Response"`
}

func (r *DeleteIntegrationEmployeesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationEmployeesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationRoleUsersRequestParams struct {
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 角色id
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 用户信息,最多 200 个用户，并且 UserId 和 OpenId 二选一，其他字段不需要传
	Users []*UserInfo `json:"Users,omitnil" name:"Users"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DeleteIntegrationRoleUsersRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 角色id
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 用户信息,最多 200 个用户，并且 UserId 和 OpenId 二选一，其他字段不需要传
	Users []*UserInfo `json:"Users,omitnil" name:"Users"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DeleteIntegrationRoleUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationRoleUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "RoleId")
	delete(f, "Users")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntegrationRoleUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationRoleUsersResponseParams struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteIntegrationRoleUsersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntegrationRoleUsersResponseParams `json:"Response"`
}

func (r *DeleteIntegrationRoleUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationRoleUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSealPoliciesRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 印章授权编码数组。这个参数跟下面的SealId其中一个必填，另外一个可选填
	PolicyIds []*string `json:"PolicyIds,omitnil" name:"PolicyIds"`

	// 印章ID。这个参数跟上面的PolicyIds其中一个必填，另外一个可选填
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 待授权的员工ID
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DeleteSealPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 印章授权编码数组。这个参数跟下面的SealId其中一个必填，另外一个可选填
	PolicyIds []*string `json:"PolicyIds,omitnil" name:"PolicyIds"`

	// 印章ID。这个参数跟上面的PolicyIds其中一个必填，另外一个可选填
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 待授权的员工ID
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DeleteSealPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSealPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "PolicyIds")
	delete(f, "SealId")
	delete(f, "UserIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSealPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSealPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSealPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSealPoliciesResponseParams `json:"Response"`
}

func (r *DeleteSealPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSealPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStaffsResult struct {
	// 删除员工的成功数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessEmployeeData []*SuccessDeleteStaffData `json:"SuccessEmployeeData,omitnil" name:"SuccessEmployeeData"`

	// 删除员工的失败数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedEmployeeData []*FailedDeleteStaffData `json:"FailedEmployeeData,omitnil" name:"FailedEmployeeData"`
}

type Department struct {
	// 部门id
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 部门名称
	DepartmentName *string `json:"DepartmentName,omitnil" name:"DepartmentName"`
}

// Predefined struct for user
type DescribeExtendedServiceAuthInfosRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 要查询的扩展服务类型。
	// 默认为空，即查询当前支持的所有扩展服务信息。
	// 若需查询单个扩展服务的开通情况，请传递相应的值，如下所示：
	// <ul><li>OPEN_SERVER_SIGN：企业静默签署</li>
	// <li>OVERSEA_SIGN：企业与港澳台居民签署合同</li>
	// <li>MOBILE_CHECK_APPROVER：使用手机号验证签署方身份</li>
	// <li>PAGING_SEAL：骑缝章</li>
	// <li>BATCH_SIGN：批量签署</li></ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil" name:"ExtendServiceType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DescribeExtendedServiceAuthInfosRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 要查询的扩展服务类型。
	// 默认为空，即查询当前支持的所有扩展服务信息。
	// 若需查询单个扩展服务的开通情况，请传递相应的值，如下所示：
	// <ul><li>OPEN_SERVER_SIGN：企业静默签署</li>
	// <li>OVERSEA_SIGN：企业与港澳台居民签署合同</li>
	// <li>MOBILE_CHECK_APPROVER：使用手机号验证签署方身份</li>
	// <li>PAGING_SEAL：骑缝章</li>
	// <li>BATCH_SIGN：批量签署</li></ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil" name:"ExtendServiceType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DescribeExtendedServiceAuthInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtendedServiceAuthInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ExtendServiceType")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtendedServiceAuthInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtendedServiceAuthInfosResponseParams struct {
	// 服务开通和授权的信息列表，根据查询类型返回所有支持的扩展服务开通和授权状况，或者返回特定扩展服务的开通和授权状况。
	AuthInfoList []*ExtendAuthInfo `json:"AuthInfoList,omitnil" name:"AuthInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeExtendedServiceAuthInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtendedServiceAuthInfosResponseParams `json:"Response"`
}

func (r *DescribeExtendedServiceAuthInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtendedServiceAuthInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUrlsRequestParams struct {
	// 调用方用户信息，UserId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 文件对应的业务类型，目前支持：
	// - 流程 "FLOW"，如需下载合同文件请选择此项
	// - 模板 "TEMPLATE"
	// - 文档 "DOCUMENT"
	// - 印章  “SEAL”
	BusinessType *string `json:"BusinessType,omitnil" name:"BusinessType"`

	// 业务编号的数组，如流程编号、模板编号、文档编号、印章编号。如需下载合同文件请传入FlowId
	// 最大支持20个资源
	BusinessIds []*string `json:"BusinessIds,omitnil" name:"BusinessIds"`

	// 下载后的文件命名，只有FileType为zip的时候生效
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 文件类型，"JPG", "PDF","ZIP"等
	FileType *string `json:"FileType,omitnil" name:"FileType"`

	// 指定资源起始偏移量，默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 指定资源数量，查询全部资源则传入-1
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 下载url过期时间，单位秒。0: 按默认值5分钟，允许范围：1s~24x60x60s(1天)
	UrlTtl *int64 `json:"UrlTtl,omitnil" name:"UrlTtl"`

	// 暂不开放
	//
	// Deprecated: CcToken is deprecated.
	CcToken *string `json:"CcToken,omitnil" name:"CcToken"`

	// 暂不开放
	//
	// Deprecated: Scene is deprecated.
	Scene *string `json:"Scene,omitnil" name:"Scene"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DescribeFileUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，UserId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 文件对应的业务类型，目前支持：
	// - 流程 "FLOW"，如需下载合同文件请选择此项
	// - 模板 "TEMPLATE"
	// - 文档 "DOCUMENT"
	// - 印章  “SEAL”
	BusinessType *string `json:"BusinessType,omitnil" name:"BusinessType"`

	// 业务编号的数组，如流程编号、模板编号、文档编号、印章编号。如需下载合同文件请传入FlowId
	// 最大支持20个资源
	BusinessIds []*string `json:"BusinessIds,omitnil" name:"BusinessIds"`

	// 下载后的文件命名，只有FileType为zip的时候生效
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 文件类型，"JPG", "PDF","ZIP"等
	FileType *string `json:"FileType,omitnil" name:"FileType"`

	// 指定资源起始偏移量，默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 指定资源数量，查询全部资源则传入-1
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 下载url过期时间，单位秒。0: 按默认值5分钟，允许范围：1s~24x60x60s(1天)
	UrlTtl *int64 `json:"UrlTtl,omitnil" name:"UrlTtl"`

	// 暂不开放
	CcToken *string `json:"CcToken,omitnil" name:"CcToken"`

	// 暂不开放
	Scene *string `json:"Scene,omitnil" name:"Scene"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DescribeFileUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUrlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "BusinessType")
	delete(f, "BusinessIds")
	delete(f, "FileName")
	delete(f, "FileType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "UrlTtl")
	delete(f, "CcToken")
	delete(f, "Scene")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUrlsResponseParams struct {
	// 文件URL信息；
	// 链接不是永久链接，有效期5分钟后链接失效。
	FileUrls []*FileUrl `json:"FileUrls,omitnil" name:"FileUrls"`

	// URL数量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFileUrlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileUrlsResponseParams `json:"Response"`
}

func (r *DescribeFileUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowBriefsRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 需要查询的流程ID列表，限制最大100个
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DescribeFlowBriefsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 需要查询的流程ID列表，限制最大100个
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DescribeFlowBriefsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowBriefsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowBriefsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowBriefsResponseParams struct {
	// 流程列表
	FlowBriefs []*FlowBrief `json:"FlowBriefs,omitnil" name:"FlowBriefs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFlowBriefsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowBriefsResponseParams `json:"Response"`
}

func (r *DescribeFlowBriefsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowBriefsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowComponentsRequestParams struct {
	// 操作者信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 流程(合同)的编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DescribeFlowComponentsRequest struct {
	*tchttp.BaseRequest
	
	// 操作者信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 流程(合同)的编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DescribeFlowComponentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowComponentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowComponentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowComponentsResponseParams struct {
	// 流程关联的填写控件信息，按照参与方进行分类返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientComponentInfos []*RecipientComponentInfo `json:"RecipientComponentInfos,omitnil" name:"RecipientComponentInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFlowComponentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowComponentsResponseParams `json:"Response"`
}

func (r *DescribeFlowComponentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowEvidenceReportRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 签署报告编号
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DescribeFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 签署报告编号
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DescribeFlowEvidenceReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowEvidenceReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ReportId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowEvidenceReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowEvidenceReportResponseParams struct {
	// 出证报告PDF的下载 URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportUrl *string `json:"ReportUrl,omitnil" name:"ReportUrl"`

	// 签署报告出证任务的状态
	// <ul>
	// <li>EvidenceStatusExecuting : 出证任务在执行中</li>
	// <li>EvidenceStatusSuccess : 出证任务执行成功</li>
	// <li>EvidenceStatusFailed : 出证任务执行失败</li>
	// </ul>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFlowEvidenceReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowEvidenceReportResponseParams `json:"Response"`
}

func (r *DescribeFlowEvidenceReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowEvidenceReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowInfoRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 需要查询的流程ID列表，限制最大100个
	// 
	// 如果查询合同组的信息,不要传此参数
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同组ID, 如果传此参数会忽略FlowIds入参
	//  所以如传此参数不要传FlowIds参数
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`
}

type DescribeFlowInfoRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 需要查询的流程ID列表，限制最大100个
	// 
	// 如果查询合同组的信息,不要传此参数
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同组ID, 如果传此参数会忽略FlowIds入参
	//  所以如传此参数不要传FlowIds参数
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`
}

func (r *DescribeFlowInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "Agent")
	delete(f, "FlowGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowInfoResponseParams struct {
	// 签署流程信息
	FlowDetailInfos []*FlowDetailInfo `json:"FlowDetailInfos,omitnil" name:"FlowDetailInfos"`

	// 合同组ID
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 合同组名称
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFlowInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowInfoResponseParams `json:"Response"`
}

func (r *DescribeFlowInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowTemplatesRequestParams struct {
	// 调用方员工/经办人信息
	// UserId 必填，在企业控制台组织架构中可以查到员工的UserId
	// 注：请保证员工有相关的角色权限
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 代理相关应用信息
	// 如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询内容类型
	// 0-模板列表及详情（默认）
	// 1-仅模板列表
	ContentType *int64 `json:"ContentType,omitnil" name:"ContentType"`

	// 搜索条件，本字段用于指定模板Id进行查询。
	// Key：template-id
	// Values：需要查询的模板Id列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，此处指定第几页，如果不传默从第一页返回。页码从0开始，即首页为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 指定每页多少条数据，如果不传默认为20，单页最大200。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 用于查询指定应用号下单模板列表。
	// ApplicationId不为空，查询指定应用下的模板列表
	// ApplicationId为空，查询所有应用下的模板列表
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 默认为false，查询SaaS模板库列表；
	// 为true，查询第三方应用集成平台企业模板库管理列表
	//
	// Deprecated: IsChannel is deprecated.
	IsChannel *bool `json:"IsChannel,omitnil" name:"IsChannel"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 暂未开放
	//
	// Deprecated: GenerateSource is deprecated.
	GenerateSource *uint64 `json:"GenerateSource,omitnil" name:"GenerateSource"`
}

type DescribeFlowTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 调用方员工/经办人信息
	// UserId 必填，在企业控制台组织架构中可以查到员工的UserId
	// 注：请保证员工有相关的角色权限
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 代理相关应用信息
	// 如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询内容类型
	// 0-模板列表及详情（默认）
	// 1-仅模板列表
	ContentType *int64 `json:"ContentType,omitnil" name:"ContentType"`

	// 搜索条件，本字段用于指定模板Id进行查询。
	// Key：template-id
	// Values：需要查询的模板Id列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，此处指定第几页，如果不传默从第一页返回。页码从0开始，即首页为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 指定每页多少条数据，如果不传默认为20，单页最大200。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 用于查询指定应用号下单模板列表。
	// ApplicationId不为空，查询指定应用下的模板列表
	// ApplicationId为空，查询所有应用下的模板列表
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 默认为false，查询SaaS模板库列表；
	// 为true，查询第三方应用集成平台企业模板库管理列表
	IsChannel *bool `json:"IsChannel,omitnil" name:"IsChannel"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 暂未开放
	GenerateSource *uint64 `json:"GenerateSource,omitnil" name:"GenerateSource"`
}

func (r *DescribeFlowTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "ContentType")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ApplicationId")
	delete(f, "IsChannel")
	delete(f, "Organization")
	delete(f, "GenerateSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowTemplatesResponseParams struct {
	// 模板详情列表
	Templates []*TemplateInfo `json:"Templates,omitnil" name:"Templates"`

	// 查询到的总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFlowTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowTemplatesResponseParams `json:"Response"`
}

func (r *DescribeFlowTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationDepartmentsRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 查询类型 0-查询单个部门节点 1-单个部门节点及一级子节点部门列表
	QueryType *uint64 `json:"QueryType,omitnil" name:"QueryType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 部门ID,与DeptOpenId二选一,优先DeptId,都为空时获取根节点数据
	DeptId *string `json:"DeptId,omitnil" name:"DeptId"`

	// 客户系统部门ID,与DeptId二选一,优先DeptId,都为空时获取根节点数据
	DeptOpenId *string `json:"DeptOpenId,omitnil" name:"DeptOpenId"`
}

type DescribeIntegrationDepartmentsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 查询类型 0-查询单个部门节点 1-单个部门节点及一级子节点部门列表
	QueryType *uint64 `json:"QueryType,omitnil" name:"QueryType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 部门ID,与DeptOpenId二选一,优先DeptId,都为空时获取根节点数据
	DeptId *string `json:"DeptId,omitnil" name:"DeptId"`

	// 客户系统部门ID,与DeptId二选一,优先DeptId,都为空时获取根节点数据
	DeptOpenId *string `json:"DeptOpenId,omitnil" name:"DeptOpenId"`
}

func (r *DescribeIntegrationDepartmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationDepartmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "QueryType")
	delete(f, "Agent")
	delete(f, "DeptId")
	delete(f, "DeptOpenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationDepartmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationDepartmentsResponseParams struct {
	// 部门列表
	Departments []*IntegrationDepartment `json:"Departments,omitnil" name:"Departments"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIntegrationDepartmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationDepartmentsResponseParams `json:"Response"`
}

func (r *DescribeIntegrationDepartmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationDepartmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationEmployeesRequestParams struct {
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 指定每页多少条数据，单页最大20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询过滤实名用户，Key为Status，Values为["IsVerified"]
	// 根据第三方系统openId过滤查询员工时,Key为StaffOpenId,Values为["OpenId","OpenId",...]
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0，最大20000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeIntegrationEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 指定每页多少条数据，单页最大20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询过滤实名用户，Key为Status，Values为["IsVerified"]
	// 根据第三方系统openId过滤查询员工时,Key为StaffOpenId,Values为["OpenId","OpenId",...]
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0，最大20000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeIntegrationEmployeesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationEmployeesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Limit")
	delete(f, "Agent")
	delete(f, "Filters")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationEmployeesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationEmployeesResponseParams struct {
	// 员工数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Employees []*Staff `json:"Employees,omitnil" name:"Employees"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0，最大20000
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 指定每页多少条数据，单页最大20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 符合条件的员工数量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIntegrationEmployeesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationEmployeesResponseParams `json:"Response"`
}

func (r *DescribeIntegrationEmployeesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationEmployeesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationMainOrganizationUserRequestParams struct {
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type DescribeIntegrationMainOrganizationUserRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

func (r *DescribeIntegrationMainOrganizationUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationMainOrganizationUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationMainOrganizationUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationMainOrganizationUserResponseParams struct {
	// 主企业员工账号信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntegrationMainOrganizationUser *IntegrationMainOrganizationUser `json:"IntegrationMainOrganizationUser,omitnil" name:"IntegrationMainOrganizationUser"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIntegrationMainOrganizationUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationMainOrganizationUserResponseParams `json:"Response"`
}

func (r *DescribeIntegrationMainOrganizationUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationMainOrganizationUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationRolesRequestParams struct {
	// 操作人信息，UserId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 指定每页多少条数据，单页最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询的关键字段:
	// Key:"RoleType",Values:["1"]查询系统角色，Values:["2"]查询自定义角色
	// Key:"RoleStatus",Values:["1"]查询启用角色，Values:["2"]查询禁用角色
	// Key:"IsGroupRole"，Values:["0"]:查询非集团角色，Values:["1"]表示查询集团角色
	// Key:"IsReturnPermissionGroup"，Values:["0"]:表示接口不返回角色对应的权限树字段，Values:["1"]表示接口返回角色对应的权限树字段
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0，最大2000
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeIntegrationRolesRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息，UserId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 指定每页多少条数据，单页最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询的关键字段:
	// Key:"RoleType",Values:["1"]查询系统角色，Values:["2"]查询自定义角色
	// Key:"RoleStatus",Values:["1"]查询启用角色，Values:["2"]查询禁用角色
	// Key:"IsGroupRole"，Values:["0"]:查询非集团角色，Values:["1"]表示查询集团角色
	// Key:"IsReturnPermissionGroup"，Values:["0"]:表示接口不返回角色对应的权限树字段，Values:["1"]表示接口返回角色对应的权限树字段
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0，最大2000
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeIntegrationRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Limit")
	delete(f, "Agent")
	delete(f, "Filters")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationRolesResponseParams struct {
	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0，最大2000
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 指定每页多少条数据，单页最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 符合查询条件的总的角色数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 企业角色信息列表
	IntegrateRoles []*IntegrateRole `json:"IntegrateRoles,omitnil" name:"IntegrateRoles"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIntegrationRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationRolesResponseParams `json:"Response"`
}

func (r *DescribeIntegrationRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationGroupOrganizationsRequestParams struct {
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 指定每页多少条数据，单页最大1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询成员企业的企业名，模糊匹配
	Name *string `json:"Name,omitnil" name:"Name"`

	// 成员企业加入集团的当前状态:1-待授权;2-已授权待激活;3-拒绝授权;4-已解除;5-已加入
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 是否导出当前成员企业数据
	Export *bool `json:"Export,omitnil" name:"Export"`

	// 成员企业机构 ID，在PC控制台 集团管理可获取
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeOrganizationGroupOrganizationsRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息，userId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 指定每页多少条数据，单页最大1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询成员企业的企业名，模糊匹配
	Name *string `json:"Name,omitnil" name:"Name"`

	// 成员企业加入集团的当前状态:1-待授权;2-已授权待激活;3-拒绝授权;4-已解除;5-已加入
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 是否导出当前成员企业数据
	Export *bool `json:"Export,omitnil" name:"Export"`

	// 成员企业机构 ID，在PC控制台 集团管理可获取
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeOrganizationGroupOrganizationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationGroupOrganizationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Name")
	delete(f, "Status")
	delete(f, "Export")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationGroupOrganizationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationGroupOrganizationsResponseParams struct {
	// 查询到的符合条件的成员企业总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 已授权待激活的企业数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinedTotal *uint64 `json:"JoinedTotal,omitnil" name:"JoinedTotal"`

	// 已加入的企业数量(废弃,请使用ActivatedTotal)
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ActivedTotal is deprecated.
	ActivedTotal *uint64 `json:"ActivedTotal,omitnil" name:"ActivedTotal"`

	// 如果入参Export为 true 时使用，表示导出Excel的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExportUrl *string `json:"ExportUrl,omitnil" name:"ExportUrl"`

	// 成员企业信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*GroupOrganization `json:"List,omitnil" name:"List"`

	// 已加入的企业数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivatedTotal *uint64 `json:"ActivatedTotal,omitnil" name:"ActivatedTotal"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOrganizationGroupOrganizationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationGroupOrganizationsResponseParams `json:"Response"`
}

func (r *DescribeOrganizationGroupOrganizationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationGroupOrganizationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationSealsRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 返回最大数量，最大为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询信息类型，为0时不返回授权用户，为1时返回
	InfoType *int64 `json:"InfoType,omitnil" name:"InfoType"`

	// 印章id（没有输入返回所有）
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 印章类型列表（都是组织机构印章）。
	// 为空时查询所有类型的印章。
	// 目前支持以下类型：
	// OFFICIAL：企业公章；
	// CONTRACT：合同专用章；
	// ORGANIZATION_SEAL：企业印章(图片上传创建)；
	// LEGAL_PERSON_SEAL：法定代表人章
	SealTypes []*string `json:"SealTypes,omitnil" name:"SealTypes"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询的印章状态列表。
	// 取值为空，只查询启用状态的印章；
	// 取值ALL，查询所有状态的印章；
	// 取值CHECKING，查询待审核的印章；
	// 取值SUCCESS，查询启用状态的印章；
	// 取值FAIL，查询印章审核拒绝的印章；
	// 取值DISABLE，查询已停用的印章；
	// 取值STOPPED，查询已终止的印章；
	// 取值VOID，查询已作废的印章；
	// 取值INVALID，查询已失效的印章；
	SealStatuses []*string `json:"SealStatuses,omitnil" name:"SealStatuses"`
}

type DescribeOrganizationSealsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 返回最大数量，最大为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询信息类型，为0时不返回授权用户，为1时返回
	InfoType *int64 `json:"InfoType,omitnil" name:"InfoType"`

	// 印章id（没有输入返回所有）
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 印章类型列表（都是组织机构印章）。
	// 为空时查询所有类型的印章。
	// 目前支持以下类型：
	// OFFICIAL：企业公章；
	// CONTRACT：合同专用章；
	// ORGANIZATION_SEAL：企业印章(图片上传创建)；
	// LEGAL_PERSON_SEAL：法定代表人章
	SealTypes []*string `json:"SealTypes,omitnil" name:"SealTypes"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询的印章状态列表。
	// 取值为空，只查询启用状态的印章；
	// 取值ALL，查询所有状态的印章；
	// 取值CHECKING，查询待审核的印章；
	// 取值SUCCESS，查询启用状态的印章；
	// 取值FAIL，查询印章审核拒绝的印章；
	// 取值DISABLE，查询已停用的印章；
	// 取值STOPPED，查询已终止的印章；
	// 取值VOID，查询已作废的印章；
	// 取值INVALID，查询已失效的印章；
	SealStatuses []*string `json:"SealStatuses,omitnil" name:"SealStatuses"`
}

func (r *DescribeOrganizationSealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationSealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InfoType")
	delete(f, "SealId")
	delete(f, "SealTypes")
	delete(f, "Agent")
	delete(f, "SealStatuses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationSealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationSealsResponseParams struct {
	// 在设置了SealId时返回0或1，没有设置时返回公司的总印章数量，可能比返回的印章数组数量多
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 查询到的印章结果数组
	Seals []*OccupiedSeal `json:"Seals,omitnil" name:"Seals"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOrganizationSealsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationSealsResponseParams `json:"Response"`
}

func (r *DescribeOrganizationSealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationSealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThirdPartyAuthCodeRequestParams struct {
	// 电子签小程序跳转客户小程序时携带的授权查看码
	AuthCode *string `json:"AuthCode,omitnil" name:"AuthCode"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DescribeThirdPartyAuthCodeRequest struct {
	*tchttp.BaseRequest
	
	// 电子签小程序跳转客户小程序时携带的授权查看码
	AuthCode *string `json:"AuthCode,omitnil" name:"AuthCode"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DescribeThirdPartyAuthCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThirdPartyAuthCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuthCode")
	delete(f, "Operator")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeThirdPartyAuthCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThirdPartyAuthCodeResponseParams struct {
	// 用户是否实名，VERIFIED 为实名，UNVERIFIED 未实名
	VerifyStatus *string `json:"VerifyStatus,omitnil" name:"VerifyStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeThirdPartyAuthCodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeThirdPartyAuthCodeResponseParams `json:"Response"`
}

func (r *DescribeThirdPartyAuthCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThirdPartyAuthCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserAutoSignStatusRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** : 电子处方场景</li></ul>
	// 
	// 注: `现在仅支持电子处方场景`
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 要查询状态的用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DescribeUserAutoSignStatusRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** : 电子处方场景</li></ul>
	// 
	// 注: `现在仅支持电子处方场景`
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 要查询状态的用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DescribeUserAutoSignStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserAutoSignStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserAutoSignStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserAutoSignStatusResponseParams struct {
	// 查询用户是否已开通自动签
	IsOpen *bool `json:"IsOpen,omitnil" name:"IsOpen"`

	// 自动签许可生效时间。当且仅当已通过许可开通自动签时有值。
	// 
	// 值为unix时间戳,单位为秒。
	LicenseFrom *int64 `json:"LicenseFrom,omitnil" name:"LicenseFrom"`

	// 自动签许可到期时间。当且仅当已通过许可开通自动签时有值。
	// 
	// 值为unix时间戳,单位为秒。
	LicenseTo *int64 `json:"LicenseTo,omitnil" name:"LicenseTo"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。
	// 
	// <ul><li>**0**: 使用个人自动签账号许可进行开通，个人自动签账号许可有效期1年，注: `不可解绑释放更换他人`</li>
	// <li>**1**: 不使用个人自动签账号许可进行开通</li></ul>
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserAutoSignStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserAutoSignStatusResponseParams `json:"Response"`
}

func (r *DescribeUserAutoSignStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserAutoSignStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUserAutoSignRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** 电子处方</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 需要关闭自动签的个人的信息，如姓名，证件信息等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DisableUserAutoSignRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** 电子处方</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 需要关闭自动签的个人的信息，如姓名，证件信息等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *DisableUserAutoSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUserAutoSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableUserAutoSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUserAutoSignResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableUserAutoSignResponse struct {
	*tchttp.BaseResponse
	Response *DisableUserAutoSignResponseParams `json:"Response"`
}

func (r *DisableUserAutoSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUserAutoSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmbedUrlOption struct {
	// 合同详情预览，允许展示控件信息
	// <br/>true：允许在合同详情页展示控件
	// <br/>false：不允许在合同详情页展示控件
	// <br/>默认false，合同详情页不展示控件
	ShowFlowDetailComponent *bool `json:"ShowFlowDetailComponent,omitnil" name:"ShowFlowDetailComponent"`

	// 模板预览，允许展示模板控件信息
	// <br/>true：允许在模板预览页展示控件
	// <br/>false：不允许在模板预览页展示控件
	// <br/>默认false，模板预览页不展示控件
	ShowTemplateComponent *bool `json:"ShowTemplateComponent,omitnil" name:"ShowTemplateComponent"`
}

type ExtendAuthInfo struct {
	// 扩展服务的类型，可能是以下值：
	// <ul><li>OPEN_SERVER_SIGN：企业静默签署</li>
	// <li>OVERSEA_SIGN：企业与港澳台居民签署合同</li>
	// <li>MOBILE_CHECK_APPROVER：使用手机号验证签署方身份</li>
	// <li>PAGING_SEAL：骑缝章</li>
	// <li>BATCH_SIGN：批量签署</li></ul>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 扩展服务的名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 扩展服务的开通状态：
	// ENABLE：开通
	// DISABLE：未开通
	Status *string `json:"Status,omitnil" name:"Status"`

	// 操作扩展服务的操作人UserId，员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUserId *string `json:"OperatorUserId,omitnil" name:"OperatorUserId"`

	// 扩展服务的操作时间，格式为Unix标准时间戳（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateOn *int64 `json:"OperateOn,omitnil" name:"OperateOn"`

	// 该扩展服务若可以授权，此参数对应授权人员的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasAuthUserList []*HasAuthUser `json:"HasAuthUserList,omitnil" name:"HasAuthUserList"`
}

type FailedCreateRoleData struct {
	// 用户userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 角色id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleIds []*string `json:"RoleIds,omitnil" name:"RoleIds"`
}

type FailedCreateStaffData struct {
	// 员工名
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// 员工手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 传入的企微账号id
	WeworkOpenId *string `json:"WeworkOpenId,omitnil" name:"WeworkOpenId"`

	// 失败原因
	Reason *string `json:"Reason,omitnil" name:"Reason"`
}

type FailedDeleteStaffData struct {
	// 员工在电子签的userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 员工在第三方平台的openId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 失败原因
	Reason *string `json:"Reason,omitnil" name:"Reason"`
}

type FailedUpdateStaffData struct {
	// 用户传入的名称
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// 用户传入的手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 失败原因
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 用户Id
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 员工在第三方平台的openId
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`
}

type FileInfo struct {
	// 文件ID
	FileId *string `json:"FileId,omitnil" name:"FileId"`

	// 文件名
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 文件大小，单位为Byte
	FileSize *int64 `json:"FileSize,omitnil" name:"FileSize"`

	// 文件上传时间，格式为Unix标准时间戳（秒）
	CreatedOn *int64 `json:"CreatedOn,omitnil" name:"CreatedOn"`
}

type FileUrl struct {
	// 下载文件的URL，有效期为输入的UrlTtl，默认5分钟
	Url *string `json:"Url,omitnil" name:"Url"`

	// 下载文件的附加信息。如果是pdf文件，会返回pdf文件每页的有效高宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	Option *string `json:"Option,omitnil" name:"Option"`
}

type FillApproverInfo struct {
	// 对应模板中的参与方ID
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 签署人来源
	// WEWORKAPP: 企业微信
	// <br/>仅【企微或签】时指定WEWORKAPP
	ApproverSource *string `json:"ApproverSource,omitnil" name:"ApproverSource"`

	// 企业自定义账号ID
	// <br/>当ApproverSource为WEWORKAPP的企微或签场景下，必须指企业自有应用获取企微明文的userid
	CustomUserId *string `json:"CustomUserId,omitnil" name:"CustomUserId"`

	// 补充签署人姓名
	ApproverName *string `json:"ApproverName,omitnil" name:"ApproverName"`

	// 补充签署人手机号
	ApproverMobile *string `json:"ApproverMobile,omitnil" name:"ApproverMobile"`
}

type FilledComponent struct {
	// 控件Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentId *string `json:"ComponentId,omitnil" name:"ComponentId"`

	// 控件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// 控件填写状态；0-未填写；1-已填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentFillStatus *string `json:"ComponentFillStatus,omitnil" name:"ComponentFillStatus"`

	// 控件填写内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentValue *string `json:"ComponentValue,omitnil" name:"ComponentValue"`

	// 控件所属参与方Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentRecipientId *string `json:"ComponentRecipientId,omitnil" name:"ComponentRecipientId"`

	// 图片填充控件下载链接，如果是图片填充控件时，这里返回图片的下载链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type Filter struct {
	// 查询过滤条件的Key
	Key *string `json:"Key,omitnil" name:"Key"`

	// 查询过滤条件的Value列表
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type FlowApproverDetail struct {
	// 签署时的相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMessage *string `json:"ApproveMessage,omitnil" name:"ApproveMessage"`

	// 签署方姓名
	ApproveName *string `json:"ApproveName,omitnil" name:"ApproveName"`

	// 签署方的签署状态
	// 0：还没有发起
	// 1：流程中 没有开始处理
	// 2：待签署
	// 3：已签署
	// 4：已拒绝
	// 5：已过期
	// 6：已撤销
	// 7：还没有预发起
	// 8：待填写
	// 9：因为各种原因而终止
	// 10：填写完成
	// 15：已解除
	// 19：转他人处理
	ApproveStatus *int64 `json:"ApproveStatus,omitnil" name:"ApproveStatus"`

	// 模板配置中的参与方ID,与控件绑定
	ReceiptId *string `json:"ReceiptId,omitnil" name:"ReceiptId"`

	// 客户自定义的用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomUserId *string `json:"CustomUserId,omitnil" name:"CustomUserId"`

	// 签署人手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 签署顺序，如果是有序签署，签署顺序从小到大
	SignOrder *int64 `json:"SignOrder,omitnil" name:"SignOrder"`

	// 签署人签署时间，时间戳，单位秒
	ApproveTime *int64 `json:"ApproveTime,omitnil" name:"ApproveTime"`

	// 签署方类型，ORGANIZATION-企业员工，PERSON-个人，ENTERPRISESERVER-企业静默签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveType *string `json:"ApproveType,omitnil" name:"ApproveType"`

	// 签署方侧用户来源，如WEWORKAPP-企业微信等
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverSource *string `json:"ApproverSource,omitnil" name:"ApproverSource"`

	// 客户自定义签署方标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomApproverTag *string `json:"CustomApproverTag,omitnil" name:"CustomApproverTag"`

	// 签署方企业Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 签署方企业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 签署参与人在本流程中的编号ID（每个流程不同），可用此ID来定位签署参与人在本流程的签署节点，也可用于后续创建签署链接等操作。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"SignId,omitnil" name:"SignId"`
}

type FlowApproverUrlInfo struct {
	// 签署链接(短链形式呈现)。请注意保密，不要将其外泄给无关用户。
	// 注: `注意该链接有效期为30分钟`
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignUrl *string `json:"SignUrl,omitnil" name:"SignUrl"`

	// 签署参与人类型 
	// <ul><li> **1** :个人参与方</li></ul>
	// 
	// 注: `现在仅支持个人参与方`
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverType *int64 `json:"ApproverType,omitnil" name:"ApproverType"`

	// 签署人姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverName *string `json:"ApproverName,omitnil" name:"ApproverName"`

	// 签署人手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverMobile *string `json:"ApproverMobile,omitnil" name:"ApproverMobile"`

	// 签署链接(长链形式呈现)。请注意保密，不要将其外泄给无关用户。
	// 注: `注意该链接有效期为30分钟`
	// 注意：此字段可能返回 null，表示取不到有效值。
	LongUrl *string `json:"LongUrl,omitnil" name:"LongUrl"`
}

type FlowBrief struct {
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 合同流程的名称。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同流程的类别分类（如销售合同/入职合同等）。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 合同流程当前的签署状态, 会存在下列的状态值
	// <ul><li> **0** : 未开启流程(合同中不存在填写环节)</li>
	// <li> **1** : 待签署</li>
	// <li> **2** : 部分签署</li>
	// <li> **3** : 已拒签</li>
	// <li> **4** : 已签署</li>
	// <li> **5** : 已过期</li>
	// <li> **6** : 已撤销</li>
	// <li> **7** : 未开启流程(合同中存在填写环节)</li>
	// <li> **8** : 等待填写</li>
	// <li> **9** : 部分填写</li>
	// <li> **10** : 已拒填</li>
	// <li> **21** : 已解除</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowStatus *int64 `json:"FlowStatus,omitnil" name:"FlowStatus"`

	// 合同流程创建时间，格式为Unix标准时间戳（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedOn *int64 `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 当合同流程状态为已拒签（即 FlowStatus=3）或已撤销（即 FlowStatus=6）时，此字段 FlowMessage 为拒签或撤销原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMessage *string `json:"FlowMessage,omitnil" name:"FlowMessage"`

	//  合同流程发起方的员工编号, 即员工在腾讯电子签平台的唯一身份标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`
}

type FlowCreateApprover struct {
	// 在指定签署方时，可选择企业B端或个人C端等不同的参与者类型，可选类型如下:
	// 0：企业
	// 1：个人
	// 3：企业静默签署
	// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。静默签署仅进行盖章操作，不能自动签名。
	// 7: 个人自动签署，适用于个人自动签场景。
	// 注: 个人自动签场景为白名单功能，使用前请联系对接的客户经理沟通。
	ApproverType *int64 `json:"ApproverType,omitnil" name:"ApproverType"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 
	// 注: `当approverType=0(企业签署方) 或 approverType=3(企业静默签署)时，必须指定`
	// 
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 在未指定签署人电子签UserId情况下，为必填参数
	ApproverName *string `json:"ApproverName,omitnil" name:"ApproverName"`

	// 签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此合同签署方。
	// 
	// 在未指定签署人电子签UserId情况下，为必填参数
	ApproverMobile *string `json:"ApproverMobile,omitnil" name:"ApproverMobile"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li></ul>
	ApproverIdCardType *string `json:"ApproverIdCardType,omitnil" name:"ApproverIdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母（但“I”、“O”除外），后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitnil" name:"ApproverIdCardNumber"`

	// 签署方经办人在模板中的参与方ID
	// <br/>模板发起合同时，该参数为必填项
	// <br/>文件发起合同是，该参数无序传值
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 签署意愿确认渠道，默认为WEIXINAPP:人脸识别
	// 
	// 注: 将要废弃, 用ApproverSignTypes签署人签署合同时的认证方式代替, 新客户可请用ApproverSignTypes来设置
	VerifyChannel []*string `json:"VerifyChannel,omitnil" name:"VerifyChannel"`

	// 通知签署方经办人的方式,  有以下途径:
	// <ul><li>  **sms**  :  (默认)短信</li>
	// <li>   **none**   : 不通知</li></ul>
	// 
	// 注: `发起方也是签署方时不给此签署方发送短信`
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 合同强制需要阅读全文，无需传此参数
	IsFullText *bool `json:"IsFullText,omitnil" name:"IsFullText"`

	// 合同的强制预览时间：3~300s，未指定则按合同页数计算
	PreReadTime *uint64 `json:"PreReadTime,omitnil" name:"PreReadTime"`

	// 签署人userId，仅支持本企业的员工userid， 可在控制台组织管理处获得
	// 
	// 注: `若传此字段 则以userid的信息为主，会覆盖传递过来的签署人基本信息， 包括姓名，手机号，证件类型等信息`
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 字段已经废弃，当前只支持true，默认为true
	Required *bool `json:"Required,omitnil" name:"Required"`

	// 在企微场景下使用，需设置参数为**WEWORKAPP**，以表明合同来源于企微。
	ApproverSource *string `json:"ApproverSource,omitnil" name:"ApproverSource"`

	// 在企业微信场景下，表明该合同流程为或签，其最大长度为64位字符串。
	// 所有参与或签的人员均需具备该标识。
	// 注意，在合同中，不同的或签参与人必须保证其CustomApproverTag唯一。
	// 如果或签签署人为本方企业微信参与人，则需要指定ApproverSource参数为WEWORKAPP。
	CustomApproverTag *string `json:"CustomApproverTag,omitnil" name:"CustomApproverTag"`

	// 已经废弃, 快速注册相关信息
	RegisterInfo *RegisterInfo `json:"RegisterInfo,omitnil" name:"RegisterInfo"`

	// 签署人个性化能力值，如是否可以转发他人处理、是否可以拒签等功能开关。
	ApproverOption *ApproverOption `json:"ApproverOption,omitnil" name:"ApproverOption"`

	// 签署完前端跳转的url，暂未使用
	//
	// Deprecated: JumpUrl is deprecated.
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 签署ID
	// - 发起流程时系统自动补充
	// - 创建签署链接时，可以通过查询详情接口获得签署人的SignId，然后可传入此值为该签署人创建签署链接，无需再传姓名、手机号、证件号等其他信息
	SignId *string `json:"SignId,omitnil" name:"SignId"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li>**false**：（默认）不需要审批，直接签署。</li>
	// <li>**true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li>如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li>如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	ApproverNeedSignReview *bool `json:"ApproverNeedSignReview,omitnil" name:"ApproverNeedSignReview"`

	// 签署人签署控件， 此参数仅针对文件发起（CreateFlowByFiles）生效
	// 
	// 合同中的签署控件列表，列表中可支持下列多种签署控件,控件的详细定义参考开发者中心的Component结构体
	// <ul><li> 个人签名/印章</li>
	// <li> 企业印章</li>
	// <li> 骑缝章等签署控件</li></ul>
	// 
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	SignComponents []*Component `json:"SignComponents,omitnil" name:"SignComponents"`

	// 签署人填写控件 此参数仅针对文件发起（CreateFlowByFiles）生效
	// 
	// 合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>动态表格等填写控件</li></ul>
	// 
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 签署方控件类型为 SIGN_SIGNATURE时，可以指定签署方签名方式
	// 	HANDWRITE – 手写签名
	// 	OCR_ESIGN -- AI智能识别手写签名
	// 	ESIGN -- 个人印章类型
	// 	SYSTEM_ESIGN -- 系统签名（该类型可以在用户签署时根据用户姓名一键生成一个签名来进行签署）
	ComponentLimitType []*string `json:"ComponentLimitType,omitnil" name:"ComponentLimitType"`

	// 指定个人签署方查看合同的校验方式,可以传值如下:
	// <ul><li>  **1**   : （默认）人脸识别,人脸识别后才能合同内容</li>
	// <li>  **2**  : 手机号验证, 用户手机号和参与方手机号(ApproverMobile)相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// </li></ul>
	// 注: 
	// <ul><li>如果合同流程设置ApproverVerifyType查看合同的校验方式,    则忽略此签署人的查看合同的校验方式</li>
	// <li>此字段不可传多个校验方式</li></ul>
	// 
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	// 
	// .
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil" name:"ApproverVerifyTypes"`

	// 您可以指定签署方签署合同的认证校验方式，可传递以下值：
	// <ul><li>**1**：人脸认证，需进行人脸识别成功后才能签署合同；</li>
	// <li>**2**：签署密码，需输入与用户在腾讯电子签设置的密码一致才能校验成功进行合同签署；</li>
	// <li>**3**：运营商三要素，需到运营商处比对手机号实名信息（名字、手机号、证件号）校验一致才能成功进行合同签署。</li></ul>
	// 注：
	// <ul><li>默认情况下，认证校验方式为人脸认证和签署密码两种形式；</li>
	// <li>您可以传递多种值，表示可用多种认证校验方式。</li></ul>
	// 
	// 注:
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	ApproverSignTypes []*uint64 `json:"ApproverSignTypes,omitnil" name:"ApproverSignTypes"`
}

type FlowDetailInfo struct {
	// 合同(流程)的ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 合同(流程)的名字
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同(流程)的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 流程状态
	// - 0 还没有发起
	// - 1 待签署
	// - 2 部分签署
	// - 3 已拒签
	// - 4 已签署
	// - 5 已过期
	// - 6 已撤销
	// - 7 还没有预发起
	// - 8 等待填写
	// - 9 部分填写
	// - 10 拒填
	// - 21 已解除
	FlowStatus *int64 `json:"FlowStatus,omitnil" name:"FlowStatus"`

	// 合同(流程)的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMessage *string `json:"FlowMessage,omitnil" name:"FlowMessage"`

	// 流程的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同(流程)的创建时间戳，单位秒
	CreatedOn *int64 `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 合同(流程)的签署方数组
	FlowApproverInfos []*FlowApproverDetail `json:"FlowApproverInfos,omitnil" name:"FlowApproverInfos"`

	// 合同(流程)的关注方信息列表
	CcInfos []*FlowApproverDetail `json:"CcInfos,omitnil" name:"CcInfos"`

	// 合同发起人UserId
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil" name:"Creator"`
}

type FlowGroupInfo struct {
	// 合同（流程）的名称
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同（流程）的签署方信息
	Approvers []*ApproverInfo `json:"Approvers,omitnil" name:"Approvers"`

	// 发起合同（流程）的资源Id,此资源必须是PDF文件,来自UploadFiles,使用文件发起合同(流程)组时必传
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 发起合同（流程）的模板Id,用模板发起合同（流程）组时必填
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 合同（流程）的类型
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 合同（流程）的描述
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同（流程）的截止时间戳，单位秒。默认是一年
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 合同（流程）的回调地址
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 第三方平台传递过来的信息, 限制1024字符 格式必须是base64的
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 合同（流程）的签署是否是无序签, true - 无序。 false - 有序, 默认 
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 合同（流程）发起方的填写控件, 由发起方进行在发起时进行填充
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 本企业（发起方）是否需要签署审批，若需要审批则只允许查看不允许签署，需要您调用接口CreateFlowSignReview提交审批结果。
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 本企业（发起方）自动签署，需要您在发起合同时给印章控件指定自动签的印章。
	AutoSignScene *string `json:"AutoSignScene,omitnil" name:"AutoSignScene"`
}

type FlowGroupOptions struct {
	// 发起合同（流程）组的合同（流程）签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证
	// 参数说明：此参数仅在合同组文件发起有效，可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 发起合同（流程）组本方企业经办人通知方式
	// 签署通知类型：sms--短信，none--不通知
	SelfOrganizationApproverNotifyType *string `json:"SelfOrganizationApproverNotifyType,omitnil" name:"SelfOrganizationApproverNotifyType"`

	// 发起合同（流程）组他方经办人通知方式
	// 签署通知类型：sms--短信，none--不通知
	OtherApproverNotifyType *string `json:"OtherApproverNotifyType,omitnil" name:"OtherApproverNotifyType"`
}

type FormField struct {
	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// TEXT - 文本内容
	// MULTI_LINE_TEXT - 文本内容
	// CHECK_BOX - true/false
	// FILL_IMAGE、ATTACHMENT - 附件的FileId，需要通过UploadFiles接口上传获取
	// SELECTOR - 选项值
	// DYNAMIC_TABLE - 传入json格式的表格内容，具体见数据结构FlowInfo：https://cloud.tencent.com/document/api/1420/61525#FlowInfo
	ComponentValue *string `json:"ComponentValue,omitnil" name:"ComponentValue"`

	// 控件id，和ComponentName选择一项传入即可
	ComponentId *string `json:"ComponentId,omitnil" name:"ComponentId"`

	// 控件名字，最大长度不超过30字符，和ComponentId选择一项传入即可
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`
}

// Predefined struct for user
type GetTaskResultApiRequestParams struct {
	// 任务Id，通过接口CreateConvertTaskApi或CreateMergeFileTask得到的返回任务id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 操作人信息,UserId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`
}

type GetTaskResultApiRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id，通过接口CreateConvertTaskApi或CreateMergeFileTask得到的返回任务id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 操作人信息,UserId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`
}

func (r *GetTaskResultApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskResultApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "Organization")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskResultApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskResultApiResponseParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务状态，需要关注的状态
	// 0  :NeedTranform   - 任务已提交
	// 4  :Processing     - 文档转换中
	// 8  :TaskEnd        - 任务处理完成
	// -2 :DownloadFailed - 下载失败
	// -6 :ProcessFailed  - 转换失败
	// -13:ProcessTimeout - 转换文件超时
	TaskStatus *int64 `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// 状态描述，需要关注的状态
	// NeedTranform   - 任务已提交
	// Processing     - 文档转换中
	// TaskEnd        - 任务处理完成
	// DownloadFailed - 下载失败
	// ProcessFailed  - 转换失败
	// ProcessTimeout - 转换文件超时
	TaskMessage *string `json:"TaskMessage,omitnil" name:"TaskMessage"`

	// 资源Id，也是FileId，用于文件发起时使用
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTaskResultApiResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskResultApiResponseParams `json:"Response"`
}

func (r *GetTaskResultApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskResultApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupOrganization struct {
	// 成员企业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 成员企业别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// 成员企业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 更新时间，时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 成员企业加入集团的当前状态:1-待授权;2-已授权待激活;3-拒绝授权;4-已解除;5-已加入
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 是否为集团主企业
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMainOrganization *bool `json:"IsMainOrganization,omitnil" name:"IsMainOrganization"`

	// 企业社会信用代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 企业超管信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminInfo *Admin `json:"AdminInfo,omitnil" name:"AdminInfo"`

	// 企业许可证
	// 注意：此字段可能返回 null，表示取不到有效值。
	License *string `json:"License,omitnil" name:"License"`

	// 企业许可证过期时间，时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseExpireTime *uint64 `json:"LicenseExpireTime,omitnil" name:"LicenseExpireTime"`

	// 成员企业加入集团时间，时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinTime *uint64 `json:"JoinTime,omitnil" name:"JoinTime"`

	// 是否使用自建审批流引擎（即不是企微审批流引擎），true-是，false-否
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowEngineEnable *bool `json:"FlowEngineEnable,omitnil" name:"FlowEngineEnable"`
}

type HasAuthUser struct {
	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 当前员工的归属情况，可能值是：
	// MainOrg：在集团企业的场景下，返回此值代表是归属主企业
	// CurrentOrg：在普通企业场景下返回此值；或者在集团企业的场景下，返回此值代表归属子企业
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongTo *string `json:"BelongTo,omitnil" name:"BelongTo"`
}

type IntegrateRole struct {
	// 角色id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 角色名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// 角色状态，1-启用，2-禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleStatus *uint64 `json:"RoleStatus,omitnil" name:"RoleStatus"`

	// 是否是集团角色，true-是，false-否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsGroupRole *bool `json:"IsGroupRole,omitnil" name:"IsGroupRole"`

	// 管辖的子企业列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubOrgIdList []*string `json:"SubOrgIdList,omitnil" name:"SubOrgIdList"`

	// 权限树
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`
}

type IntegrationDepartment struct {
	// 部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeptId *string `json:"DeptId,omitnil" name:"DeptId"`

	// 部门名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeptName *string `json:"DeptName,omitnil" name:"DeptName"`

	// 父部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentDeptId *string `json:"ParentDeptId,omitnil" name:"ParentDeptId"`

	// 客户系统部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeptOpenId *string `json:"DeptOpenId,omitnil" name:"DeptOpenId"`

	// 序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *uint64 `json:"OrderNo,omitnil" name:"OrderNo"`
}

type IntegrationMainOrganizationUser struct {
	// 主企业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainOrganizationId *string `json:"MainOrganizationId,omitnil" name:"MainOrganizationId"`

	// 主企业员工UserId
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainUserId *string `json:"MainUserId,omitnil" name:"MainUserId"`

	// 主企业员工名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil" name:"UserName"`
}

// Predefined struct for user
type ModifyApplicationCallbackInfoRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 操作类型：
	// 1-新增
	// 2-删除
	OperateType *int64 `json:"OperateType,omitnil" name:"OperateType"`

	// 企业应用回调信息
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitnil" name:"CallbackInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type ModifyApplicationCallbackInfoRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 操作类型：
	// 1-新增
	// 2-删除
	OperateType *int64 `json:"OperateType,omitnil" name:"OperateType"`

	// 企业应用回调信息
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitnil" name:"CallbackInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *ModifyApplicationCallbackInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationCallbackInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "OperateType")
	delete(f, "CallbackInfo")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationCallbackInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationCallbackInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyApplicationCallbackInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationCallbackInfoResponseParams `json:"Response"`
}

func (r *ModifyApplicationCallbackInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationCallbackInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationDepartmentRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子签部门ID,通过DescribeIntegrationDepartments接口可以获取
	DeptId *string `json:"DeptId,omitnil" name:"DeptId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 电子签父部门ID，通过DescribeIntegrationDepartments接口可以获取
	ParentDeptId *string `json:"ParentDeptId,omitnil" name:"ParentDeptId"`

	// 部门名称，不超过50个字符
	DeptName *string `json:"DeptName,omitnil" name:"DeptName"`

	// 客户系统部门ID，不超过64个字符
	DeptOpenId *string `json:"DeptOpenId,omitnil" name:"DeptOpenId"`

	// 排序号,1~30000范围内
	OrderNo *uint64 `json:"OrderNo,omitnil" name:"OrderNo"`
}

type ModifyIntegrationDepartmentRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子签部门ID,通过DescribeIntegrationDepartments接口可以获取
	DeptId *string `json:"DeptId,omitnil" name:"DeptId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 电子签父部门ID，通过DescribeIntegrationDepartments接口可以获取
	ParentDeptId *string `json:"ParentDeptId,omitnil" name:"ParentDeptId"`

	// 部门名称，不超过50个字符
	DeptName *string `json:"DeptName,omitnil" name:"DeptName"`

	// 客户系统部门ID，不超过64个字符
	DeptOpenId *string `json:"DeptOpenId,omitnil" name:"DeptOpenId"`

	// 排序号,1~30000范围内
	OrderNo *uint64 `json:"OrderNo,omitnil" name:"OrderNo"`
}

func (r *ModifyIntegrationDepartmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationDepartmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "DeptId")
	delete(f, "Agent")
	delete(f, "ParentDeptId")
	delete(f, "DeptName")
	delete(f, "DeptOpenId")
	delete(f, "OrderNo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIntegrationDepartmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationDepartmentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyIntegrationDepartmentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIntegrationDepartmentResponseParams `json:"Response"`
}

func (r *ModifyIntegrationDepartmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationRoleRequestParams struct {
	// 角色Id，可通过接口 DescribeIntegrationRoles 查询获取
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 权限树
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`

	// 集团角色的话，需要传递集团子企业列表，如果是全选，则传1
	SubOrganizationIds []*string `json:"SubOrganizationIds,omitnil" name:"SubOrganizationIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type ModifyIntegrationRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色Id，可通过接口 DescribeIntegrationRoles 查询获取
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 权限树
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`

	// 集团角色的话，需要传递集团子企业列表，如果是全选，则传1
	SubOrganizationIds []*string `json:"SubOrganizationIds,omitnil" name:"SubOrganizationIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *ModifyIntegrationRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleId")
	delete(f, "Name")
	delete(f, "Operator")
	delete(f, "Description")
	delete(f, "PermissionGroups")
	delete(f, "SubOrganizationIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIntegrationRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationRoleResponseParams struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyIntegrationRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIntegrationRoleResponseParams `json:"Response"`
}

func (r *ModifyIntegrationRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OccupiedSeal struct {
	// 电子印章编号
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 电子印章名称
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 电子印章授权时间戳，单位秒
	CreateOn *int64 `json:"CreateOn,omitnil" name:"CreateOn"`

	// 电子印章授权人的UserId
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 电子印章策略Id
	SealPolicyId *string `json:"SealPolicyId,omitnil" name:"SealPolicyId"`

	// 印章状态，有以下六种：CHECKING（审核中）SUCCESS（已启用）FAIL（审核拒绝）CHECKING-SADM（待超管审核）DISABLE（已停用）STOPPED（已终止）
	SealStatus *string `json:"SealStatus,omitnil" name:"SealStatus"`

	// 审核失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitnil" name:"FailReason"`

	// 印章图片url，5分钟内有效
	Url *string `json:"Url,omitnil" name:"Url"`

	// 印章类型,OFFICIAL-企业公章, CONTRACT-合同专用章,ORGANIZATIONSEAL-企业印章(本地上传印章类型),LEGAL_PERSON_SEAL-法人印章
	SealType *string `json:"SealType,omitnil" name:"SealType"`

	// 用印申请是否为永久授权，true-是，false-否
	IsAllTime *bool `json:"IsAllTime,omitnil" name:"IsAllTime"`

	// 授权人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedUsers []*AuthorizedUser `json:"AuthorizedUsers,omitnil" name:"AuthorizedUsers"`
}

type OrganizationInfo struct {
	// 机构在平台的编号，内部字段，暂未开放
	//
	// Deprecated: OrganizationId is deprecated.
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 用户渠道，内部字段，暂未开放
	//
	// Deprecated: Channel is deprecated.
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 用户在渠道的机构编号，内部字段，暂未开放
	//
	// Deprecated: OrganizationOpenId is deprecated.
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 用户真实的IP，内部字段，暂未开放
	//
	// Deprecated: ClientIp is deprecated.
	ClientIp *string `json:"ClientIp,omitnil" name:"ClientIp"`

	// 机构的代理IP，内部字段，暂未开放
	//
	// Deprecated: ProxyIp is deprecated.
	ProxyIp *string `json:"ProxyIp,omitnil" name:"ProxyIp"`
}

type PdfVerifyResult struct {
	// 验签结果。0-签名域未签名；1-验签成功； 3-验签失败；4-未找到签名域：文件内没有签名域；5-签名值格式不正确。
	VerifyResult *int64 `json:"VerifyResult,omitnil" name:"VerifyResult"`

	// 签署平台，如果文件是在腾讯电子签平台签署，则返回腾讯电子签，如果文件不在腾讯电子签平台签署，则返回其他平台。
	SignPlatform *string `json:"SignPlatform,omitnil" name:"SignPlatform"`

	// 签署人名称
	SignerName *string `json:"SignerName,omitnil" name:"SignerName"`

	// 签署时间戳，单位秒
	SignTime *int64 `json:"SignTime,omitnil" name:"SignTime"`

	// 签名算法
	SignAlgorithm *string `json:"SignAlgorithm,omitnil" name:"SignAlgorithm"`

	// 签名证书序列号
	CertSn *string `json:"CertSn,omitnil" name:"CertSn"`

	// 证书起始时间戳，单位毫秒
	CertNotBefore *int64 `json:"CertNotBefore,omitnil" name:"CertNotBefore"`

	// 证书过期时间戳，单位毫秒
	CertNotAfter *int64 `json:"CertNotAfter,omitnil" name:"CertNotAfter"`

	// 签名域横坐标，单位pt
	ComponentPosX *float64 `json:"ComponentPosX,omitnil" name:"ComponentPosX"`

	// 签名域纵坐标，单位pt
	ComponentPosY *float64 `json:"ComponentPosY,omitnil" name:"ComponentPosY"`

	// 签名域宽度，单位pt
	ComponentWidth *float64 `json:"ComponentWidth,omitnil" name:"ComponentWidth"`

	// 签名域高度，单位pt
	ComponentHeight *float64 `json:"ComponentHeight,omitnil" name:"ComponentHeight"`

	// 签名域所在页码，1～N
	ComponentPage *int64 `json:"ComponentPage,omitnil" name:"ComponentPage"`
}

type Permission struct {
	// 权限名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 权限key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 权限类型 1前端，2后端
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 是否隐藏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hide *int64 `json:"Hide,omitnil" name:"Hide"`

	// 数据权限标签 1:表示根节点，2:表示叶子结点
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataLabel *int64 `json:"DataLabel,omitnil" name:"DataLabel"`

	// 数据权限独有，1:关联其他模块鉴权，2:表示关联自己模块鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *int64 `json:"DataType,omitnil" name:"DataType"`

	// 数据权限独有，表示数据范围，1：全公司，2:部门及下级部门，3:自己
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRange *int64 `json:"DataRange,omitnil" name:"DataRange"`

	// 关联权限, 表示这个功能权限要受哪个数据权限管控
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataTo *string `json:"DataTo,omitnil" name:"DataTo"`

	// 父级权限key
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentKey *string `json:"ParentKey,omitnil" name:"ParentKey"`

	// 是否选中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsChecked *bool `json:"IsChecked,omitnil" name:"IsChecked"`

	// 子权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*Permission `json:"Children,omitnil" name:"Children"`
}

type PermissionGroup struct {
	// 权限组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 权限组key
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupKey *string `json:"GroupKey,omitnil" name:"GroupKey"`

	// 是否隐藏分组，0否1是
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hide *int64 `json:"Hide,omitnil" name:"Hide"`

	// 权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*Permission `json:"Permissions,omitnil" name:"Permissions"`
}

type Recipient struct {
	// 签署参与者ID，唯一标识
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 参与者类型。
	// 默认为空。
	// ENTERPRISE-企业；
	// INDIVIDUAL-个人；
	// PROMOTER-发起方
	RecipientType *string `json:"RecipientType,omitnil" name:"RecipientType"`

	// 描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// 是否需要验证，
	// 默认为false-不需要验证
	RequireValidation *bool `json:"RequireValidation,omitnil" name:"RequireValidation"`

	// 是否需要签署，
	// 默认为true-需要签署
	RequireSign *bool `json:"RequireSign,omitnil" name:"RequireSign"`

	// 此参与方添加的顺序，从0～N
	RoutingOrder *int64 `json:"RoutingOrder,omitnil" name:"RoutingOrder"`

	// 是否需要发送，
	// 默认为true-需要发送
	RequireDelivery *bool `json:"RequireDelivery,omitnil" name:"RequireDelivery"`

	// 邮箱地址
	Email *string `json:"Email,omitnil" name:"Email"`

	// 电话号码
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 关联的用户ID，电子签系统的用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 发送方式，默认为EMAIL。
	// EMAIL-邮件；
	// MOBILE-手机短信；
	// WECHAT-微信通知
	DeliveryMethod *string `json:"DeliveryMethod,omitnil" name:"DeliveryMethod"`

	// 参与方的一些附属信息，json格式
	RecipientExtra *string `json:"RecipientExtra,omitnil" name:"RecipientExtra"`
}

type RecipientComponentInfo struct {
	// 参与方Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 参与方填写状态
	// 0-未填写
	// 1-已填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientFillStatus *string `json:"RecipientFillStatus,omitnil" name:"RecipientFillStatus"`

	// 是否为发起方
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPromoter *bool `json:"IsPromoter,omitnil" name:"IsPromoter"`

	// 填写控件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*FilledComponent `json:"Components,omitnil" name:"Components"`
}

type RegisterInfo struct {
	// 法人姓名
	LegalName *string `json:"LegalName,omitnil" name:"LegalName"`

	// 社会统一信用代码
	//
	// Deprecated: Uscc is deprecated.
	Uscc *string `json:"Uscc,omitnil" name:"Uscc"`

	// 社会统一信用代码
	UnifiedSocialCreditCode *string `json:"UnifiedSocialCreditCode,omitnil" name:"UnifiedSocialCreditCode"`
}

type ReleasedApprover struct {
	// 签署人姓名，最大长度50个字。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 签署人手机号。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 要更换的原合同参与人RecipientId编号。(可通过接口<a href="https://qian.tencent.com/developers/companyApis/queryFlows/DescribeFlowInfo/">DescribeFlowInfo</a>查询签署人的RecipientId编号)<br/>
	RelievedApproverReceiptId *string `json:"RelievedApproverReceiptId,omitnil" name:"RelievedApproverReceiptId"`

	// 指定签署人类型，目前仅支持
	// <ul><li> **ORGANIZATION**：企业（默认值）</li>
	// <li> **ENTERPRISESERVER**：企业静默签</li></ul>
	ApproverType *string `json:"ApproverType,omitnil" name:"ApproverType"`

	// 签署控件类型，支持自定义企业签署方的签署控件类型
	// <ul><li> **SIGN_SEAL**：默认为印章控件类型（默认值）</li>
	// <li> **SIGN_SIGNATURE**：手写签名控件类型</li></ul>
	ApproverSignComponentType *string `json:"ApproverSignComponentType,omitnil" name:"ApproverSignComponentType"`

	// 参与方在合同中的角色是按照创建合同的时候来排序的; 解除协议默认会将第一个参与人叫甲方, 第二个叫乙方,第三个叫丙方，以此类推。如果您需要改动参与人的角色名字, 可以设置此签署方自定义控件别名字段，最大20个字。
	ApproverSignRole *string `json:"ApproverSignRole,omitnil" name:"ApproverSignRole"`
}

type RelieveInfo struct {
	// 解除理由，最大支持200个字
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 解除后仍然有效的条款，保留条款，最大支持200个字
	RemainInForceItem *string `json:"RemainInForceItem,omitnil" name:"RemainInForceItem"`

	// 原合同事项处理-费用结算，最大支持200个字
	OriginalExpenseSettlement *string `json:"OriginalExpenseSettlement,omitnil" name:"OriginalExpenseSettlement"`

	// 原合同事项处理-其他事项，最大支持200个字
	OriginalOtherSettlement *string `json:"OriginalOtherSettlement,omitnil" name:"OriginalOtherSettlement"`

	// 其他约定，最大支持200个字
	OtherDeals *string `json:"OtherDeals,omitnil" name:"OtherDeals"`
}

type RemindFlowRecords struct {
	// 合同流程是否可以催办：
	// true - 可以，false - 不可以。
	// 若无法催办，将返回RemindMessage以解释原因。
	CanRemind *bool `json:"CanRemind,omitnil" name:"CanRemind"`

	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 在合同流程无法催办的情况下，系统将返回RemindMessage以阐述原因。
	RemindMessage *string `json:"RemindMessage,omitnil" name:"RemindMessage"`
}

type ReviewerInfo struct {
	// 姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`
}

type SealInfo struct {
	// 印章ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 印章类型。LEGAL_PERSON_SEAL: 法定代表人章；
	// ORGANIZATIONSEAL：企业印章；
	// OFFICIAL：企业公章；
	// CONTRACT：合同专用章
	// 注意：此字段可能返回 null，表示取不到有效值。
	SealType *string `json:"SealType,omitnil" name:"SealType"`

	// 印章名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SealName *string `json:"SealName,omitnil" name:"SealName"`
}

type SignQrCode struct {
	// 二维码ID，为32位字符串。
	QrCodeId *string `json:"QrCodeId,omitnil" name:"QrCodeId"`

	// 二维码URL，可通过转换二维码的工具或代码组件将此URL转化为二维码，以便用户扫描进行流程签署。
	QrCodeUrl *string `json:"QrCodeUrl,omitnil" name:"QrCodeUrl"`

	// 二维码的有截止时间，格式为Unix标准时间戳（秒）。
	// 一旦超过二维码的有效期限，该二维码将自动失效。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`
}

type SignUrl struct {
	// 跳转至电子签名小程序签署的链接地址。
	// 适用于客户端APP及小程序直接唤起电子签名小程序。
	AppSignUrl *string `json:"AppSignUrl,omitnil" name:"AppSignUrl"`

	// 签署链接有效时间，格式类似"2022-08-05 15:55:01"
	EffectiveTime *string `json:"EffectiveTime,omitnil" name:"EffectiveTime"`

	// 跳转至电子签名小程序签署的链接地址，格式类似于https://essurl.cn/xxx。
	// 打开此链接将会展示H5中间页面，随后唤起电子签名小程序以进行合同签署。
	HttpSignUrl *string `json:"HttpSignUrl,omitnil" name:"HttpSignUrl"`
}

type Staff struct {
	// 用户在电子签平台的id
	// 注：创建和更新场景无需填写
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 显示的用户名/昵称
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// 用户手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 用户邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil" name:"Email"`

	// 用户在第三方平台id，如需在此接口提醒员工实名，该参数不传
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 员工角色
	// 注：创建和更新场景无需填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*StaffRole `json:"Roles,omitnil" name:"Roles"`

	// 员工部门
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil" name:"Department"`

	// 员工是否实名
	// 注：创建和更新场景无需填写
	Verified *bool `json:"Verified,omitnil" name:"Verified"`

	// 员工创建时间戳，单位秒
	// 注：创建和更新场景无需填写
	CreatedOn *int64 `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 员工实名时间戳，单位秒
	// 注：创建和更新场景无需填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifiedOn *int64 `json:"VerifiedOn,omitnil" name:"VerifiedOn"`

	// 员工是否离职：0-未离职，1-离职
	// 注：创建和更新场景无需填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuiteJob *int64 `json:"QuiteJob,omitnil" name:"QuiteJob"`

	// 员工离职交接人用户id
	// 注：创建和更新场景无需填写
	ReceiveUserId *string `json:"ReceiveUserId,omitnil" name:"ReceiveUserId"`

	// 员工离职交接人用户OpenId
	// 注：创建和更新场景无需填写
	ReceiveOpenId *string `json:"ReceiveOpenId,omitnil" name:"ReceiveOpenId"`

	// 企业微信用户账号ID
	// 注：仅企微类型的企业创建员工接口支持该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeworkOpenId *string `json:"WeworkOpenId,omitnil" name:"WeworkOpenId"`
}

type StaffRole struct {
	// 角色id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`
}

// Predefined struct for user
type StartFlowRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 此处需要传入[创建签署流程接口](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow)得到的FlowId。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 客户端Token，保持接口幂等性,最大长度64个字符
	//
	// Deprecated: ClientToken is deprecated.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 若在创建签署流程时指定了关注人CcInfos，此参数可设定向关注人发送短信通知的类型：
	// <ul><li> **0** :合同发起时通知通知对方来查看合同（默认）</li>
	// <li> **1** : 签署完成后通知对方来查看合同</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil" name:"CcNotifyType"`
}

type StartFlowRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 此处需要传入[创建签署流程接口](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow)得到的FlowId。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 若在创建签署流程时指定了关注人CcInfos，此参数可设定向关注人发送短信通知的类型：
	// <ul><li> **0** :合同发起时通知通知对方来查看合同（默认）</li>
	// <li> **1** : 签署完成后通知对方来查看合同</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil" name:"CcNotifyType"`
}

func (r *StartFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "ClientToken")
	delete(f, "Agent")
	delete(f, "CcNotifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartFlowResponseParams struct {
	// 发起成功后返回的状态，根据合同流程的不同，返回不同状态：
	// <ul><li> **START** : 发起成功, 合同进入签署环节</li>
	// <li> **REVIEW** : 提交审核成功, 合同需要发起审核, 发起方企业通过接口审核通过后合同才进入签署环境  `白名单功能，使用前请联系对接的客户经理沟通。`</li>
	// <li> **EXECUTING** : 已提交发起任务且PDF合同正在合成中, 等PDF合同合成成功后进入签署环节</li></ul>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartFlowResponse struct {
	*tchttp.BaseResponse
	Response *StartFlowResponseParams `json:"Response"`
}

func (r *StartFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SuccessCreateStaffData struct {
	// 员工名
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// 员工手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 员工在电子签平台的id
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 提示，当创建已存在未实名用户时，该字段有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitnil" name:"Note"`

	// 传入的企微账号id
	WeworkOpenId *string `json:"WeworkOpenId,omitnil" name:"WeworkOpenId"`
}

type SuccessDeleteStaffData struct {
	// 员工名
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// 员工手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 员工在电子签平台的id
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type SuccessUpdateStaffData struct {
	// 传入的用户名称
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// 传入的手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 用户Id
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type TemplateInfo struct {
	// 模板ID，模板的唯一标识
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 模板名
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 模板描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 模板关联的资源ID列表
	DocumentResourceIds []*string `json:"DocumentResourceIds,omitnil" name:"DocumentResourceIds"`

	// 生成模板的文件基础信息
	FileInfos []*FileInfo `json:"FileInfos,omitnil" name:"FileInfos"`

	// 附件关联的资源ID
	AttachmentResourceIds []*string `json:"AttachmentResourceIds,omitnil" name:"AttachmentResourceIds"`

	// 签署顺序
	// 无序 -1
	// 有序为序列数字 0,1,2
	SignOrder []*int64 `json:"SignOrder,omitnil" name:"SignOrder"`

	// 模板中的签署参与方列表
	Recipients []*Recipient `json:"Recipients,omitnil" name:"Recipients"`

	// 模板的填充控件列表
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 模板中的签署控件列表
	SignComponents []*Component `json:"SignComponents,omitnil" name:"SignComponents"`

	// 模板状态
	// -1:不可用
	// 0:草稿态
	// 1:正式态，可以正常使用
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 模板的创建者信息，电子签系统用户ID
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 模板创建的时间戳，格式为Unix标准时间戳（秒）
	CreatedOn *int64 `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 发起方参与信息Promoter
	Promoter *Recipient `json:"Promoter,omitnil" name:"Promoter"`

	// 模板类型：
	// 1  静默签,
	// 3  普通模板
	TemplateType *int64 `json:"TemplateType,omitnil" name:"TemplateType"`

	// 模板可用状态：
	// 1 启用（默认）
	// 2 停用
	Available *int64 `json:"Available,omitnil" name:"Available"`

	// 创建模板的企业ID，电子签的机构ID
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 模板预览链接，有效时间5分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitnil" name:"PreviewUrl"`

	// 模板版本。默认为空时，全数字字符，初始版本为yyyyMMdd001。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateVersion *string `json:"TemplateVersion,omitnil" name:"TemplateVersion"`

	// 模板是否已发布：
	// true-已发布
	// false-未发布
	// 注意：此字段可能返回 null，表示取不到有效值。
	Published *bool `json:"Published,omitnil" name:"Published"`

	// 模板内部指定的印章列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateSeals []*SealInfo `json:"TemplateSeals,omitnil" name:"TemplateSeals"`

	// 模板内部指定的印章列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Seals is deprecated.
	Seals []*SealInfo `json:"Seals,omitnil" name:"Seals"`
}

// Predefined struct for user
type UnbindEmployeeUserIdWithClientOpenIdRequestParams struct {
	// 用户信息，OpenId与UserId二选一必填一个，OpenId是第三方客户ID，userId是用户实名后的电子签生成的ID,当传入客户系统openId，传入的openId需与电子签员工userId绑定，且参数Channel必填，Channel值为INTEGRATE；当传入参数UserId，Channel无需指定(参数用法参考示例)
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子签系统员工UserId
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 客户系统OpenId
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type UnbindEmployeeUserIdWithClientOpenIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息，OpenId与UserId二选一必填一个，OpenId是第三方客户ID，userId是用户实名后的电子签生成的ID,当传入客户系统openId，传入的openId需与电子签员工userId绑定，且参数Channel必填，Channel值为INTEGRATE；当传入参数UserId，Channel无需指定(参数用法参考示例)
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子签系统员工UserId
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 客户系统OpenId
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *UnbindEmployeeUserIdWithClientOpenIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindEmployeeUserIdWithClientOpenIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "UserId")
	delete(f, "OpenId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindEmployeeUserIdWithClientOpenIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindEmployeeUserIdWithClientOpenIdResponseParams struct {
	// 解绑是否成功，1表示成功，0表示失败
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnbindEmployeeUserIdWithClientOpenIdResponse struct {
	*tchttp.BaseResponse
	Response *UnbindEmployeeUserIdWithClientOpenIdResponseParams `json:"Response"`
}

func (r *UnbindEmployeeUserIdWithClientOpenIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindEmployeeUserIdWithClientOpenIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIntegrationEmployeesRequestParams struct {
	// 当前用户信息，UserId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 员工信息，不超过100个。
	// 根据UserId或OpenId更新员工，必填一个，优先UserId。
	// 可更新Mobile、DisplayName、Email和Department.DepartmentId字段，其他字段暂不支持
	Employees []*Staff `json:"Employees,omitnil" name:"Employees"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId需填充子企业Id
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type UpdateIntegrationEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 当前用户信息，UserId必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 员工信息，不超过100个。
	// 根据UserId或OpenId更新员工，必填一个，优先UserId。
	// 可更新Mobile、DisplayName、Email和Department.DepartmentId字段，其他字段暂不支持
	Employees []*Staff `json:"Employees,omitnil" name:"Employees"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId需填充子企业Id
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

func (r *UpdateIntegrationEmployeesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIntegrationEmployeesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Employees")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateIntegrationEmployeesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIntegrationEmployeesResponseParams struct {
	// 更新成功的用户列表
	SuccessEmployeeData []*SuccessUpdateStaffData `json:"SuccessEmployeeData,omitnil" name:"SuccessEmployeeData"`

	// 更新失败的用户列表
	FailedEmployeeData []*FailedUpdateStaffData `json:"FailedEmployeeData,omitnil" name:"FailedEmployeeData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateIntegrationEmployeesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateIntegrationEmployeesResponseParams `json:"Response"`
}

func (r *UpdateIntegrationEmployeesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIntegrationEmployeesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadFile struct {
	// Base64编码后的文件内容
	FileBody *string `json:"FileBody,omitnil" name:"FileBody"`

	// 文件名，最大长度不超过200字符
	FileName *string `json:"FileName,omitnil" name:"FileName"`
}

// Predefined struct for user
type UploadFilesRequestParams struct {
	// 文件对应业务类型
	// 1. TEMPLATE - 模板； 文件类型：.pdf/.doc/.docx/.html
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html
	// 3. SEAL - 印章； 文件类型：.jpg/.jpeg/.png
	BusinessType *string `json:"BusinessType,omitnil" name:"BusinessType"`

	// 调用方信息，其中OperatorId为必填字段，即用户的UserId
	Caller *Caller `json:"Caller,omitnil" name:"Caller"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitnil" name:"FileInfos"`

	// 文件类型， 默认通过文件内容解析得到文件类型，客户可以显示的说明上传文件的类型。
	// 如：PDF 表示上传的文件 xxx.pdf的文件类型是 PDF
	FileType *string `json:"FileType,omitnil" name:"FileType"`

	// 此参数只对 PDF 文件有效。是否将pdf灰色矩阵置白
	// true--是，处理置白
	// 默认为false--否，不处理
	CoverRect *bool `json:"CoverRect,omitnil" name:"CoverRect"`

	// 用户自定义ID数组，与上传文件一一对应
	CustomIds []*string `json:"CustomIds,omitnil" name:"CustomIds"`

	// 不再使用，上传文件链接数组，最多支持20个URL
	//
	// Deprecated: FileUrls is deprecated.
	FileUrls *string `json:"FileUrls,omitnil" name:"FileUrls"`
}

type UploadFilesRequest struct {
	*tchttp.BaseRequest
	
	// 文件对应业务类型
	// 1. TEMPLATE - 模板； 文件类型：.pdf/.doc/.docx/.html
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html
	// 3. SEAL - 印章； 文件类型：.jpg/.jpeg/.png
	BusinessType *string `json:"BusinessType,omitnil" name:"BusinessType"`

	// 调用方信息，其中OperatorId为必填字段，即用户的UserId
	Caller *Caller `json:"Caller,omitnil" name:"Caller"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitnil" name:"FileInfos"`

	// 文件类型， 默认通过文件内容解析得到文件类型，客户可以显示的说明上传文件的类型。
	// 如：PDF 表示上传的文件 xxx.pdf的文件类型是 PDF
	FileType *string `json:"FileType,omitnil" name:"FileType"`

	// 此参数只对 PDF 文件有效。是否将pdf灰色矩阵置白
	// true--是，处理置白
	// 默认为false--否，不处理
	CoverRect *bool `json:"CoverRect,omitnil" name:"CoverRect"`

	// 用户自定义ID数组，与上传文件一一对应
	CustomIds []*string `json:"CustomIds,omitnil" name:"CustomIds"`

	// 不再使用，上传文件链接数组，最多支持20个URL
	FileUrls *string `json:"FileUrls,omitnil" name:"FileUrls"`
}

func (r *UploadFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessType")
	delete(f, "Caller")
	delete(f, "FileInfos")
	delete(f, "FileType")
	delete(f, "CoverRect")
	delete(f, "CustomIds")
	delete(f, "FileUrls")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFilesResponseParams struct {
	// 文件id数组
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 上传成功文件数量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UploadFilesResponse struct {
	*tchttp.BaseResponse
	Response *UploadFilesResponseParams `json:"Response"`
}

func (r *UploadFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {
	// 用户在平台的编号
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户的来源渠道，一般不用传，特定场景根据接口说明传值
	//
	// Deprecated: Channel is deprecated.
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 用户在渠道的编号，一般不用传，特定场景根据接口说明传值
	//
	// Deprecated: OpenId is deprecated.
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 用户真实IP，内部字段，暂未开放
	//
	// Deprecated: ClientIp is deprecated.
	ClientIp *string `json:"ClientIp,omitnil" name:"ClientIp"`

	// 用户代理IP，内部字段，暂未开放
	//
	// Deprecated: ProxyIp is deprecated.
	ProxyIp *string `json:"ProxyIp,omitnil" name:"ProxyIp"`
}

type UserThreeFactor struct {
	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母（但“I”、“O”除外），后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`
}

// Predefined struct for user
type VerifyPdfRequestParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type VerifyPdfRequest struct {
	*tchttp.BaseRequest
	
	// 流程ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

func (r *VerifyPdfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPdfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyPdfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPdfResponseParams struct {
	// 验签结果，1-文件未被篡改，全部签名在腾讯电子签完成； 2-文件未被篡改，部分签名在腾讯电子签完成；3-文件被篡改；4-异常：文件内没有签名域；5-异常：文件签名格式错误
	VerifyResult *int64 `json:"VerifyResult,omitnil" name:"VerifyResult"`

	// 验签结果详情，每个签名域对应的验签结果。状态值：1-验签成功，在电子签签署；2-验签成功，在其他平台签署；3-验签失败；4-pdf文件没有签名域；5-文件签名格式错误
	PdfVerifyResults []*PdfVerifyResult `json:"PdfVerifyResults,omitnil" name:"PdfVerifyResults"`

	// 验签序列号
	VerifySerialNo *string `json:"VerifySerialNo,omitnil" name:"VerifySerialNo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VerifyPdfResponse struct {
	*tchttp.BaseResponse
	Response *VerifyPdfResponseParams `json:"Response"`
}

func (r *VerifyPdfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPdfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebThemeConfig struct {
	// 是否页面底部显示电子签logo
	// <br/>true：允许在页面底部隐藏电子签logo
	// <br/>false：不允许允许在页面底部隐藏电子签logo
	// <br/>默认false，不隐藏logo
	DisplaySignBrandLogo *bool `json:"DisplaySignBrandLogo,omitnil" name:"DisplaySignBrandLogo"`

	// 主题颜色
	// <br/>支持十六进制颜色值以及RGB格式颜色值，例如：#D54941，rgb(213, 73, 65)
	WebEmbedThemeColor *string `json:"WebEmbedThemeColor,omitnil" name:"WebEmbedThemeColor"`
}