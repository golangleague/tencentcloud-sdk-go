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

package v20210519

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CheckChainRequestParams struct {
	// 群组ID
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// did服务机构名称
	AgencyName *string `json:"AgencyName,omitnil" name:"AgencyName"`
}

type CheckChainRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// did服务机构名称
	AgencyName *string `json:"AgencyName,omitnil" name:"AgencyName"`
}

func (r *CheckChainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckChainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ClusterId")
	delete(f, "AgencyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckChainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckChainResponseParams struct {
	// 1为盟主，0为非盟主
	RoleType *int64 `json:"RoleType,omitnil" name:"RoleType"`

	// 链ID
	ChainId *string `json:"ChainId,omitnil" name:"ChainId"`

	// 应用名称
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckChainResponse struct {
	*tchttp.BaseResponse
	Response *CheckChainResponseParams `json:"Response"`
}

func (r *CheckChainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckChainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCredentialRequestParams struct {
	// 参数集合，详见示例
	FunctionArg *FunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`

	// 参数集合，详见示例
	TransactionArg *TransactionArg `json:"TransactionArg,omitnil" name:"TransactionArg"`

	// 版本
	VersionCredential *string `json:"VersionCredential,omitnil" name:"VersionCredential"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitnil" name:"UnSigned"`
}

type CreateCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 参数集合，详见示例
	FunctionArg *FunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`

	// 参数集合，详见示例
	TransactionArg *TransactionArg `json:"TransactionArg,omitnil" name:"TransactionArg"`

	// 版本
	VersionCredential *string `json:"VersionCredential,omitnil" name:"VersionCredential"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitnil" name:"UnSigned"`
}

func (r *CreateCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	delete(f, "TransactionArg")
	delete(f, "VersionCredential")
	delete(f, "UnSigned")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCredentialResponseParams struct {
	// Credential的具体信息
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateCredentialResponseParams `json:"Response"`
}

func (r *CreateCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSelectiveCredentialRequestParams struct {
	// 参数集合
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`

	// 批露策略id
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

type CreateSelectiveCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 参数集合
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`

	// 批露策略id
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

func (r *CreateSelectiveCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSelectiveCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSelectiveCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSelectiveCredentialResponseParams struct {
	// 凭证字符串
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSelectiveCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateSelectiveCredentialResponseParams `json:"Response"`
}

func (r *CreateSelectiveCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSelectiveCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPrivateKeyRequestParams struct {
	// 网络ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// 私钥
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`
}

type CreateTDidByPrivateKeyRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// 私钥
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`
}

func (r *CreateTDidByPrivateKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPrivateKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "PrivateKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByPrivateKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPrivateKeyResponseParams struct {
	// did的具体信息
	Did *string `json:"Did,omitnil" name:"Did"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidByPrivateKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByPrivateKeyResponseParams `json:"Response"`
}

func (r *CreateTDidByPrivateKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPrivateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPublicKeyRequestParams struct {
	// 网络ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// 身份公钥
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// 加密公钥
	EncryptPubKey *string `json:"EncryptPubKey,omitnil" name:"EncryptPubKey"`
}

type CreateTDidByPublicKeyRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// 身份公钥
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// 加密公钥
	EncryptPubKey *string `json:"EncryptPubKey,omitnil" name:"EncryptPubKey"`
}

func (r *CreateTDidByPublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "PublicKey")
	delete(f, "EncryptPubKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByPublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPublicKeyResponseParams struct {
	// did具体信息
	Did *string `json:"Did,omitnil" name:"Did"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidByPublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByPublicKeyResponseParams `json:"Response"`
}

func (r *CreateTDidByPublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidRequestParams struct {
	// 群组ID
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 部署机构为1，否则为0
	Relegation *uint64 `json:"Relegation,omitnil" name:"Relegation"`
}

type CreateTDidRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 部署机构为1，否则为0
	Relegation *uint64 `json:"Relegation,omitnil" name:"Relegation"`
}

func (r *CreateTDidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ClusterId")
	delete(f, "Relegation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidResponseParams struct {
	// TDID
	Did *string `json:"Did,omitnil" name:"Did"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidResponseParams `json:"Response"`
}

func (r *CreateTDidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialStatus struct {
	// 凭证唯一id
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`

	// 凭证状态（0：吊销；1：有效）
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 凭证颁发者Did
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 凭证摘要
	// 注意：此字段可能返回 null，表示取不到有效值。
	Digest *string `json:"Digest,omitnil" name:"Digest"`

	// 凭证签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Signature *string `json:"Signature,omitnil" name:"Signature"`

	// 更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeStamp *uint64 `json:"TimeStamp,omitnil" name:"TimeStamp"`
}

type FunctionArg struct {
	// CPT ID
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`

	// 签发者 did
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 过期时间
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// 声明
	ClaimJson *string `json:"ClaimJson,omitnil" name:"ClaimJson"`
}

// Predefined struct for user
type GetAuthorityIssuerRequestParams struct {
	// tdid
	Did *string `json:"Did,omitnil" name:"Did"`
}

type GetAuthorityIssuerRequest struct {
	*tchttp.BaseRequest
	
	// tdid
	Did *string `json:"Did,omitnil" name:"Did"`
}

func (r *GetAuthorityIssuerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthorityIssuerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAuthorityIssuerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAuthorityIssuerResponseParams struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 区块链网络id
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 区块链群组id
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// 权威机构did
	Did *string `json:"Did,omitnil" name:"Did"`

	// 机构备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 注册时间
	RegisterTime *string `json:"RegisterTime,omitnil" name:"RegisterTime"`

	// 认证时间
	RecognizeTime *string `json:"RecognizeTime,omitnil" name:"RecognizeTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetAuthorityIssuerResponse struct {
	*tchttp.BaseResponse
	Response *GetAuthorityIssuerResponseParams `json:"Response"`
}

func (r *GetAuthorityIssuerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthorityIssuerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCptInfoRequestParams struct {
	// Cpt索引
	CptIndex *uint64 `json:"CptIndex,omitnil" name:"CptIndex"`
}

type GetCptInfoRequest struct {
	*tchttp.BaseRequest
	
	// Cpt索引
	CptIndex *uint64 `json:"CptIndex,omitnil" name:"CptIndex"`
}

func (r *GetCptInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCptInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CptIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCptInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCptInfoResponseParams struct {
	// CptJsonData的具体信息
	CptJsonData *string `json:"CptJsonData,omitnil" name:"CptJsonData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCptInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetCptInfoResponseParams `json:"Response"`
}

func (r *GetCptInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCptInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStatusRequestParams struct {
	// 凭证id
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`
}

type GetCredentialStatusRequest struct {
	*tchttp.BaseRequest
	
	// 凭证id
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`
}

func (r *GetCredentialStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStatusResponseParams struct {
	// 凭证状态信息
	CredentialStatus *CredentialStatus `json:"CredentialStatus,omitnil" name:"CredentialStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCredentialStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialStatusResponseParams `json:"Response"`
}

func (r *GetCredentialStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidDocumentRequestParams struct {
	// tdid
	Did *string `json:"Did,omitnil" name:"Did"`
}

type GetDidDocumentRequest struct {
	*tchttp.BaseRequest
	
	// tdid
	Did *string `json:"Did,omitnil" name:"Did"`
}

func (r *GetDidDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidDocumentResponseParams struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// DID文档
	Document *string `json:"Document,omitnil" name:"Document"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDidDocumentResponse struct {
	*tchttp.BaseResponse
	Response *GetDidDocumentResponseParams `json:"Response"`
}

func (r *GetDidDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Proof struct {
	// 创建时间
	Created *int64 `json:"Created,omitnil" name:"Created"`

	// 创建着did
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// salt值
	SaltJson *string `json:"SaltJson,omitnil" name:"SaltJson"`

	// 签名
	SignatureValue *string `json:"SignatureValue,omitnil" name:"SignatureValue"`

	// type类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

// Predefined struct for user
type RegisterCptRequestParams struct {
	// 群组ID
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// CptJson的具体信息
	CptJson *string `json:"CptJson,omitnil" name:"CptJson"`

	// cptId 不填默认自增
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`
}

type RegisterCptRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// CptJson的具体信息
	CptJson *string `json:"CptJson,omitnil" name:"CptJson"`

	// cptId 不填默认自增
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`
}

func (r *RegisterCptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterCptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ClusterId")
	delete(f, "CptJson")
	delete(f, "CptId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterCptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterCptResponseParams struct {
	// 凭证模板索引
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 凭证模板id
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RegisterCptResponse struct {
	*tchttp.BaseResponse
	Response *RegisterCptResponseParams `json:"Response"`
}

func (r *RegisterCptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterCptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCredentialStatusRequestParams struct {
	// 凭证状态
	CredentialStatus *CredentialStatus `json:"CredentialStatus,omitnil" name:"CredentialStatus"`
}

type SetCredentialStatusRequest struct {
	*tchttp.BaseRequest
	
	// 凭证状态
	CredentialStatus *CredentialStatus `json:"CredentialStatus,omitnil" name:"CredentialStatus"`
}

func (r *SetCredentialStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCredentialStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCredentialStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCredentialStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetCredentialStatusResponse struct {
	*tchttp.BaseResponse
	Response *SetCredentialStatusResponseParams `json:"Response"`
}

func (r *SetCredentialStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCredentialStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransactionArg struct {
	// 凭证did
	InvokerTDid *string `json:"InvokerTDid,omitnil" name:"InvokerTDid"`
}

// Predefined struct for user
type VerifyCredentialRequestParams struct {
	// 参数集合
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`
}

type VerifyCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 参数集合
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`
}

func (r *VerifyCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyCredentialResponseParams struct {
	// 是否验证成功
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 验证返回码
	VerifyCode *uint64 `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 验证消息
	VerifyMessage *string `json:"VerifyMessage,omitnil" name:"VerifyMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VerifyCredentialResponse struct {
	*tchttp.BaseResponse
	Response *VerifyCredentialResponseParams `json:"Response"`
}

func (r *VerifyCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyFunctionArg struct {
	// CPT ID
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`

	// issuer did
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 过期时间
	ExpirationDate *int64 `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// 声明
	ClaimJson *string `json:"ClaimJson,omitnil" name:"ClaimJson"`

	// 颁发时间
	IssuanceDate *int64 `json:"IssuanceDate,omitnil" name:"IssuanceDate"`

	// context值
	Context *string `json:"Context,omitnil" name:"Context"`

	// id值
	Id *string `json:"Id,omitnil" name:"Id"`

	// 签名值
	Proof *Proof `json:"Proof,omitnil" name:"Proof"`

	// type值
	Type []*string `json:"Type,omitnil" name:"Type"`
}