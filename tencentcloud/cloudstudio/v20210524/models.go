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

package v20210524

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AgentSpaceDTO struct {
	// 工作空间名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 镜像id
	ImageId *int64 `json:"ImageId,omitnil" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitnil" name:"ImageName"`

	// 云服务器登录名称
	RemoteUser *string `json:"RemoteUser,omitnil" name:"RemoteUser"`

	// 云服务器登录地址
	RemoteHost *string `json:"RemoteHost,omitnil" name:"RemoteHost"`

	// 云服务器登录端口
	RemotePort *string `json:"RemotePort,omitnil" name:"RemotePort"`

	// 工作空间类型
	WorkspaceType *string `json:"WorkspaceType,omitnil" name:"WorkspaceType"`

	// 工作空间版本
	WorkspaceVersion *int64 `json:"WorkspaceVersion,omitnil" name:"WorkspaceVersion"`

	// 工作空间资源结构
	WorkspaceResourceDTO *WorkspaceResourceDTO `json:"WorkspaceResourceDTO,omitnil" name:"WorkspaceResourceDTO"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`
}

// Predefined struct for user
type CreateCustomizeTemplatesRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 无
	UserDefinedTemplateParams *UserDefinedTemplateParams `json:"UserDefinedTemplateParams,omitnil" name:"UserDefinedTemplateParams"`
}

type CreateCustomizeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 无
	UserDefinedTemplateParams *UserDefinedTemplateParams `json:"UserDefinedTemplateParams,omitnil" name:"UserDefinedTemplateParams"`
}

func (r *CreateCustomizeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "UserDefinedTemplateParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomizeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizeTemplatesResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceTemplateInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCustomizeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomizeTemplatesResponseParams `json:"Response"`
}

func (r *CreateCustomizeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByAgentRequestParams struct {
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 无
	AgentSpaceDTO *AgentSpaceDTO `json:"AgentSpaceDTO,omitnil" name:"AgentSpaceDTO"`
}

type CreateWorkspaceByAgentRequest struct {
	*tchttp.BaseRequest
	
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 无
	AgentSpaceDTO *AgentSpaceDTO `json:"AgentSpaceDTO,omitnil" name:"AgentSpaceDTO"`
}

func (r *CreateWorkspaceByAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "AgentSpaceDTO")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceByAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByAgentResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceInfoDTO `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWorkspaceByAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceByAgentResponseParams `json:"Response"`
}

func (r *CreateWorkspaceByAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByTemplateRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitnil" name:"TemplateId"`

	// 工作空间名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

type CreateWorkspaceByTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitnil" name:"TemplateId"`

	// 工作空间名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *CreateWorkspaceByTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "TemplateId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceByTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByTemplateResponseParams struct {
	// 创建工作空间返回的信息
	Data *WorkspaceInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWorkspaceByTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceByTemplateResponseParams `json:"Response"`
}

func (r *CreateWorkspaceByTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByVersionControlRequestParams struct {
	// 工作空间结构
	WorkspaceDTO *WorkspaceDTO `json:"WorkspaceDTO,omitnil" name:"WorkspaceDTO"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`
}

type CreateWorkspaceByVersionControlRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间结构
	WorkspaceDTO *WorkspaceDTO `json:"WorkspaceDTO,omitnil" name:"WorkspaceDTO"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`
}

func (r *CreateWorkspaceByVersionControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByVersionControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceDTO")
	delete(f, "CloudStudioSessionTeam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceByVersionControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByVersionControlResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceInfoDTO `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWorkspaceByVersionControlResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceByVersionControlResponseParams `json:"Response"`
}

func (r *CreateWorkspaceByVersionControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByVersionControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceTemporaryTokenRequestParams struct {
	// 创建工作空间凭证 DTO
	WorkspaceTokenDTO *WorkspaceTokenDTO `json:"WorkspaceTokenDTO,omitnil" name:"WorkspaceTokenDTO"`
}

type CreateWorkspaceTemporaryTokenRequest struct {
	*tchttp.BaseRequest
	
	// 创建工作空间凭证 DTO
	WorkspaceTokenDTO *WorkspaceTokenDTO `json:"WorkspaceTokenDTO,omitnil" name:"WorkspaceTokenDTO"`
}

func (r *CreateWorkspaceTemporaryTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceTemporaryTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceTokenDTO")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceTemporaryTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceTemporaryTokenResponseParams struct {
	// 工作空间临时访问 token 信息
	Data *WorkspaceTokenInfoV0 `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWorkspaceTemporaryTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceTemporaryTokenResponseParams `json:"Response"`
}

func (r *CreateWorkspaceTemporaryTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceTemporaryTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomizeTemplatesPresetsInfo struct {
	// 模板tag列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 模板图标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icons []*string `json:"Icons,omitnil" name:"Icons"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Templates *UserDefinedTemplateParams `json:"Templates,omitnil" name:"Templates"`
}

// Predefined struct for user
type DeleteCustomizeTemplatesByIdRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

type DeleteCustomizeTemplatesByIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteCustomizeTemplatesByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomizeTemplatesByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomizeTemplatesByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomizeTemplatesByIdResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCustomizeTemplatesByIdResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomizeTemplatesByIdResponseParams `json:"Response"`
}

func (r *DeleteCustomizeTemplatesByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomizeTemplatesByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesByIdRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

type DescribeCustomizeTemplatesByIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeCustomizeTemplatesByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomizeTemplatesByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesByIdResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceTemplateInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomizeTemplatesByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomizeTemplatesByIdResponseParams `json:"Response"`
}

func (r *DescribeCustomizeTemplatesByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesPresetsRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`
}

type DescribeCustomizeTemplatesPresetsRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`
}

func (r *DescribeCustomizeTemplatesPresetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesPresetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "SpaceKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomizeTemplatesPresetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesPresetsResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CustomizeTemplatesPresetsInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomizeTemplatesPresetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomizeTemplatesPresetsResponseParams `json:"Response"`
}

func (r *DescribeCustomizeTemplatesPresetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesPresetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`
}

type DescribeCustomizeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`
}

func (r *DescribeCustomizeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomizeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*WorkspaceTemplateInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomizeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomizeTemplatesResponseParams `json:"Response"`
}

func (r *DescribeCustomizeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceEnvListRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`
}

type DescribeWorkspaceEnvListRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`
}

func (r *DescribeWorkspaceEnvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceEnvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceEnvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceEnvListResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ImageUserDTO `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkspaceEnvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceEnvListResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceEnvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceEnvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceIsReadyRequestParams struct {
	// 工作空间 spaceKey
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`
}

type DescribeWorkspaceIsReadyRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间 spaceKey
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`
}

func (r *DescribeWorkspaceIsReadyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceIsReadyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceIsReadyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceIsReadyResponseParams struct {
	// 工作空间是否就绪
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkspaceIsReadyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceIsReadyResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceIsReadyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceIsReadyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceNameExistRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 工作空间名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil" name:"WorkspaceId"`
}

type DescribeWorkspaceNameExistRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 工作空间名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil" name:"WorkspaceId"`
}

func (r *DescribeWorkspaceNameExistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceNameExistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Name")
	delete(f, "WorkspaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceNameExistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceNameExistResponseParams struct {
	// 工作空间信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceInfoDTO `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkspaceNameExistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceNameExistResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceNameExistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceNameExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceStatusListRequestParams struct {
	// xxx
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`
}

type DescribeWorkspaceStatusListRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`
}

func (r *DescribeWorkspaceStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceStatusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceStatusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceStatusListResponseParams struct {
	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*WorkspaceStatusInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkspaceStatusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceStatusListResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceStatusRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`
}

type DescribeWorkspaceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`
}

func (r *DescribeWorkspaceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "SpaceKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceStatusResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceStatusInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkspaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceStatusResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageUserDTO struct {
	// 镜像模板ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 镜像模板名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// Tag时间
	Tag *string `json:"Tag,omitnil" name:"Tag"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 中文描述
	DescriptionCN *string `json:"DescriptionCN,omitnil" name:"DescriptionCN"`

	// 图标地址
	IconUrl *string `json:"IconUrl,omitnil" name:"IconUrl"`

	// 创建人
	Author *string `json:"Author,omitnil" name:"Author"`

	// 访问状态
	Visible *string `json:"Visible,omitnil" name:"Visible"`

	// 版本
	WorkspaceVersion *int64 `json:"WorkspaceVersion,omitnil" name:"WorkspaceVersion"`

	// 分类
	Sort *int64 `json:"Sort,omitnil" name:"Sort"`
}

// Predefined struct for user
type ModifyCustomizeTemplateVersionControlRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitnil" name:"TemplateId"`

	// 仓库地址
	Url *string `json:"Url,omitnil" name:"Url"`

	// 代码仓库分支/标签
	Ref *string `json:"Ref,omitnil" name:"Ref"`

	// 代码仓库 ref 类型
	RefType *string `json:"RefType,omitnil" name:"RefType"`
}

type ModifyCustomizeTemplateVersionControlRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitnil" name:"TemplateId"`

	// 仓库地址
	Url *string `json:"Url,omitnil" name:"Url"`

	// 代码仓库分支/标签
	Ref *string `json:"Ref,omitnil" name:"Ref"`

	// 代码仓库 ref 类型
	RefType *string `json:"RefType,omitnil" name:"RefType"`
}

func (r *ModifyCustomizeTemplateVersionControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplateVersionControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "TemplateId")
	delete(f, "Url")
	delete(f, "Ref")
	delete(f, "RefType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizeTemplateVersionControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizeTemplateVersionControlResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceTemplateInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCustomizeTemplateVersionControlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomizeTemplateVersionControlResponseParams `json:"Response"`
}

func (r *ModifyCustomizeTemplateVersionControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplateVersionControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizeTemplatesFullByIdRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 自定义模板参数
	UserDefinedTemplateParams *UserDefinedTemplateParams `json:"UserDefinedTemplateParams,omitnil" name:"UserDefinedTemplateParams"`
}

type ModifyCustomizeTemplatesFullByIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 自定义模板参数
	UserDefinedTemplateParams *UserDefinedTemplateParams `json:"UserDefinedTemplateParams,omitnil" name:"UserDefinedTemplateParams"`
}

func (r *ModifyCustomizeTemplatesFullByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplatesFullByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Id")
	delete(f, "UserDefinedTemplateParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizeTemplatesFullByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizeTemplatesFullByIdResponseParams struct {
	// 自定义模板返回值
	Data *WorkspaceTemplateInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCustomizeTemplatesFullByIdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomizeTemplatesFullByIdResponseParams `json:"Response"`
}

func (r *ModifyCustomizeTemplatesFullByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplatesFullByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizeTemplatesPartByIdRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 自定义模板Patched参数
	UserDefinedTemplatePatchedParams *UserDefinedTemplatePatchedParams `json:"UserDefinedTemplatePatchedParams,omitnil" name:"UserDefinedTemplatePatchedParams"`
}

type ModifyCustomizeTemplatesPartByIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 自定义模板Patched参数
	UserDefinedTemplatePatchedParams *UserDefinedTemplatePatchedParams `json:"UserDefinedTemplatePatchedParams,omitnil" name:"UserDefinedTemplatePatchedParams"`
}

func (r *ModifyCustomizeTemplatesPartByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplatesPartByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Id")
	delete(f, "UserDefinedTemplatePatchedParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizeTemplatesPartByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizeTemplatesPartByIdResponseParams struct {
	// 自定义模板返回值
	Data *WorkspaceTemplateInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCustomizeTemplatesPartByIdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomizeTemplatesPartByIdResponseParams `json:"Response"`
}

func (r *ModifyCustomizeTemplatesPartByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplatesPartByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkspaceAttributesRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 工作空间ID
	WorkspaceId *int64 `json:"WorkspaceId,omitnil" name:"WorkspaceId"`

	// 工作空间名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 工作空间描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// xxx
	PriceId *int64 `json:"PriceId,omitnil" name:"PriceId"`
}

type ModifyWorkspaceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 工作空间ID
	WorkspaceId *int64 `json:"WorkspaceId,omitnil" name:"WorkspaceId"`

	// 工作空间名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 工作空间描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// xxx
	PriceId *int64 `json:"PriceId,omitnil" name:"PriceId"`
}

func (r *ModifyWorkspaceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkspaceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "WorkspaceId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "PriceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkspaceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkspaceAttributesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWorkspaceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkspaceAttributesResponseParams `json:"Response"`
}

func (r *ModifyWorkspaceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkspaceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverWorkspaceRequestParams struct {
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 无
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`
}

type RecoverWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 无
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`
}

func (r *RecoverWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "SpaceKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverWorkspaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecoverWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *RecoverWorkspaceResponseParams `json:"Response"`
}

func (r *RecoverWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveWorkspaceRequestParams struct {
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 无
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`

	// 是否强制，true或者false
	Force *bool `json:"Force,omitnil" name:"Force"`
}

type RemoveWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 无
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`

	// 是否强制，true或者false
	Force *bool `json:"Force,omitnil" name:"Force"`
}

func (r *RemoveWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "SpaceKey")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveWorkspaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RemoveWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *RemoveWorkspaceResponseParams `json:"Response"`
}

func (r *RemoveWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkspaceRequestParams struct {
	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`
}

type RunWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`
}

func (r *RunWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceKey")
	delete(f, "CloudStudioSessionTeam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkspaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RunWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *RunWorkspaceResponseParams `json:"Response"`
}

func (r *RunWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWorkspaceRequestParams struct {
	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 是否强制终止，true或者false
	Force *string `json:"Force,omitnil" name:"Force"`
}

type StopWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitnil" name:"CloudStudioSessionTeam"`

	// 是否强制终止，true或者false
	Force *string `json:"Force,omitnil" name:"Force"`
}

func (r *StopWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceKey")
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWorkspaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *StopWorkspaceResponseParams `json:"Response"`
}

func (r *StopWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserDefinedTemplateParams struct {
	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板图标地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icon *string `json:"Icon,omitnil" name:"Icon"`

	// 模板标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 模板来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil" name:"Source"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 模板仓库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlType *string `json:"VersionControlType,omitnil" name:"VersionControlType"`

	// 模板地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlUrl *string `json:"VersionControlUrl,omitnil" name:"VersionControlUrl"`
}

type UserDefinedTemplatePatchedParams struct {
	// 模板来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil" name:"Source"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板图标地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icon *string `json:"Icon,omitnil" name:"Icon"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 模板标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

type UserInfoRsp struct {
	// 用户ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 用户验证信息
	AuthenticationUserInfo *UserSubInfo `json:"AuthenticationUserInfo,omitnil" name:"AuthenticationUserInfo"`

	// 头像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitnil" name:"Avatar"`

	// 介绍
	// 注意：此字段可能返回 null，表示取不到有效值。
	Features *string `json:"Features,omitnil" name:"Features"`

	// 状况
	PreviewStatus *int64 `json:"PreviewStatus,omitnil" name:"PreviewStatus"`
}

type UserSubInfo struct {
	// 团队名称
	Team *string `json:"Team,omitnil" name:"Team"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 昵称
	NickName *string `json:"NickName,omitnil" name:"NickName"`

	// 是否为管理员
	IsAdmin *bool `json:"IsAdmin,omitnil" name:"IsAdmin"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsTrial *bool `json:"IsTrial,omitnil" name:"IsTrial"`
}

type WorkspaceDTO struct {
	// 工作空间名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 代码来源类型
	VersionControlType *string `json:"VersionControlType,omitnil" name:"VersionControlType"`

	// 镜像id
	ImageId *int64 `json:"ImageId,omitnil" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitnil" name:"ImageName"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 工作空间版本
	WorkspaceVersion *int64 `json:"WorkspaceVersion,omitnil" name:"WorkspaceVersion"`

	// 工作空间资源结构
	WorkspaceResourceDTO *WorkspaceResourceDTO `json:"WorkspaceResourceDTO,omitnil" name:"WorkspaceResourceDTO"`

	// 代码仓库地址
	VersionControlUrl *string `json:"VersionControlUrl,omitnil" name:"VersionControlUrl"`

	// 代码Ref是分支还是标签
	VersionControlRef *string `json:"VersionControlRef,omitnil" name:"VersionControlRef"`

	// 代码Ref地址
	VersionControlRefType *string `json:"VersionControlRefType,omitnil" name:"VersionControlRefType"`

	// 快照Uid
	SnapshotUid *string `json:"SnapshotUid,omitnil" name:"SnapshotUid"`

	// 模板id
	TemplateId *int64 `json:"TemplateId,omitnil" name:"TemplateId"`

	// 价格id
	PriceId *int64 `json:"PriceId,omitnil" name:"PriceId"`

	// 初始化状态
	InitializeStatus *int64 `json:"InitializeStatus,omitnil" name:"InitializeStatus"`

	// 描述
	VersionControlDesc *string `json:"VersionControlDesc,omitnil" name:"VersionControlDesc"`
}

type WorkspaceInfo struct {
	// 工作空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *int64 `json:"WorkspaceId,omitnil" name:"WorkspaceId"`

	// 工作空间标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`

	// 工作空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type WorkspaceInfoDTO struct {
	// 工作空间创建时间
	CreateDate *string `json:"CreateDate,omitnil" name:"CreateDate"`

	// 空间key
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil" name:"WorkspaceId"`
}

type WorkspaceResourceDTO struct {
	// CPU核心数
	CpuCoreNumber *uint64 `json:"CpuCoreNumber,omitnil" name:"CpuCoreNumber"`

	// 一般内存
	NormalMemory *uint64 `json:"NormalMemory,omitnil" name:"NormalMemory"`

	// 系统存储
	SystemStorage *uint64 `json:"SystemStorage,omitnil" name:"SystemStorage"`

	// 用户存储
	UserStorage *uint64 `json:"UserStorage,omitnil" name:"UserStorage"`

	// GPU数
	GpuNumber *uint64 `json:"GpuNumber,omitnil" name:"GpuNumber"`

	// GPU内存
	GpuMemory *uint64 `json:"GpuMemory,omitnil" name:"GpuMemory"`
}

type WorkspaceShareInfo struct {
	// 共享或不共享状态
	Status *bool `json:"Status,omitnil" name:"Status"`

	// 是否与我共享
	// 注意：此字段可能返回 null，表示取不到有效值。
	WithMe *bool `json:"WithMe,omitnil" name:"WithMe"`

	// 开始共享的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// 停止共享的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 停止共享的时间
	Users []*UserInfoRsp `json:"Users,omitnil" name:"Users"`
}

type WorkspaceStatusInfo struct {
	// 空间ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 空间名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 所属人
	Owner *UserInfoRsp `json:"Owner,omitnil" name:"Owner"`

	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`

	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 最后操作时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOpsDate *string `json:"LastOpsDate,omitnil" name:"LastOpsDate"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 共享状态
	Share *WorkspaceShareInfo `json:"Share,omitnil" name:"Share"`

	// 空间类型
	WorkspaceType *string `json:"WorkspaceType,omitnil" name:"WorkspaceType"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 空间版本
	WorkspaceVersion *int64 `json:"WorkspaceVersion,omitnil" name:"WorkspaceVersion"`

	// 图标地址
	ImageIcon *string `json:"ImageIcon,omitnil" name:"ImageIcon"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateDate *string `json:"CreateDate,omitnil" name:"CreateDate"`

	// 版本控制地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlUrl *string `json:"VersionControlUrl,omitnil" name:"VersionControlUrl"`

	// 版本控制描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlDesc *string `json:"VersionControlDesc,omitnil" name:"VersionControlDesc"`

	// 版本控制引用
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlRef *string `json:"VersionControlRef,omitnil" name:"VersionControlRef"`

	// 版本控制引用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlRefType *string `json:"VersionControlRefType,omitnil" name:"VersionControlRefType"`

	// 版本控制类型
	VersionControlType *string `json:"VersionControlType,omitnil" name:"VersionControlType"`

	// 模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *int64 `json:"TemplateId,omitnil" name:"TemplateId"`

	// 快照ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotUid *string `json:"SnapshotUid,omitnil" name:"SnapshotUid"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecDesc *string `json:"SpecDesc,omitnil" name:"SpecDesc"`

	// CPU数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`
}

type WorkspaceTemplateInfo struct {
	// 模板ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 模板分类
	Category *string `json:"Category,omitnil" name:"Category"`

	// 模板名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 中文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptionEN *string `json:"DescriptionEN,omitnil" name:"DescriptionEN"`

	// 模板标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags *string `json:"Tags,omitnil" name:"Tags"`

	// 模板图标地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icon *string `json:"Icon,omitnil" name:"Icon"`

	// 默认仓库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlType *string `json:"VersionControlType,omitnil" name:"VersionControlType"`

	// 默认仓库地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlUrl *string `json:"VersionControlUrl,omitnil" name:"VersionControlUrl"`

	// 默认仓库描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlDesc *string `json:"VersionControlDesc,omitnil" name:"VersionControlDesc"`

	// 默认仓库所属人
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlOwner *string `json:"VersionControlOwner,omitnil" name:"VersionControlOwner"`

	// 默认仓库引用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlRef *string `json:"VersionControlRef,omitnil" name:"VersionControlRef"`

	// 默认仓库引用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlRefType *string `json:"VersionControlRefType,omitnil" name:"VersionControlRefType"`

	// 用户自定义仓库地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVersionControlUrl *string `json:"UserVersionControlUrl,omitnil" name:"UserVersionControlUrl"`

	// 用户自定义仓库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVersionControlType *string `json:"UserVersionControlType,omitnil" name:"UserVersionControlType"`

	// 用户自定义仓库引用
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVersionControlRef *string `json:"UserVersionControlRef,omitnil" name:"UserVersionControlRef"`

	// 用户自定义仓库引用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVersionControlRefType *string `json:"UserVersionControlRefType,omitnil" name:"UserVersionControlRefType"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevFile *string `json:"DevFile,omitnil" name:"DevFile"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	PluginFile *string `json:"PluginFile,omitnil" name:"PluginFile"`

	// 是否标记
	Marked *bool `json:"Marked,omitnil" name:"Marked"`

	// 标记状态
	MarkAt *int64 `json:"MarkAt,omitnil" name:"MarkAt"`

	// 创建时间
	CreateDate *string `json:"CreateDate,omitnil" name:"CreateDate"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModified *string `json:"LastModified,omitnil" name:"LastModified"`

	// 编号
	Sort *int64 `json:"Sort,omitnil" name:"Sort"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotUid *string `json:"SnapshotUid,omitnil" name:"SnapshotUid"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *int64 `json:"UserId,omitnil" name:"UserId"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Author *string `json:"Author,omitnil" name:"Author"`

	// 是否属于当前用户
	Me *bool `json:"Me,omitnil" name:"Me"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorAvatar *string `json:"AuthorAvatar,omitnil" name:"AuthorAvatar"`
}

type WorkspaceTokenDTO struct {
	// 工作空间 SpaceKey
	SpaceKey *string `json:"SpaceKey,omitnil" name:"SpaceKey"`

	// token过期时间，单位是秒，默认 3600
	TokenExpiredLimitSec *uint64 `json:"TokenExpiredLimitSec,omitnil" name:"TokenExpiredLimitSec"`
}

type WorkspaceTokenInfoV0 struct {
	// 访问工作空间临时凭证
	Token *string `json:"Token,omitnil" name:"Token"`

	// token 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`
}