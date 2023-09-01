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

package v20181217

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ApplyParam struct {
	// 审批流中表单唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 表单value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil" name:"Value"`

	// 表单参数描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type ApproveOpinion struct {
	// 方式 1:输入文字反馈  2:预设选项
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 审批意见
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*string `json:"Content,omitnil" name:"Content"`
}

type ApproveUser struct {
	// 用户uin
	Uin *uint64 `json:"Uin,omitnil" name:"Uin"`

	// 用户类型 (1:用户  2:用户组)
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 用户描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitnil" name:"Nick"`

	// 动态获取Scf
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scf *Scf `json:"Scf,omitnil" name:"Scf"`
}

// Predefined struct for user
type GetBpaasApproveDetailRequestParams struct {
	// 审批id
	ApproveId *uint64 `json:"ApproveId,omitnil" name:"ApproveId"`
}

type GetBpaasApproveDetailRequest struct {
	*tchttp.BaseRequest
	
	// 审批id
	ApproveId *uint64 `json:"ApproveId,omitnil" name:"ApproveId"`
}

func (r *GetBpaasApproveDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBpaasApproveDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApproveId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBpaasApproveDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBpaasApproveDetailResponseParams struct {
	// 申请人uin
	ApplyUin *uint64 `json:"ApplyUin,omitnil" name:"ApplyUin"`

	// 申请人主账号
	ApplyOwnUin *uint64 `json:"ApplyOwnUin,omitnil" name:"ApplyOwnUin"`

	// 申请人昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyUinNick *string `json:"ApplyUinNick,omitnil" name:"ApplyUinNick"`

	// 审批流id
	BpaasId *uint64 `json:"BpaasId,omitnil" name:"BpaasId"`

	// 审批流名称
	BpaasName *string `json:"BpaasName,omitnil" name:"BpaasName"`

	// 申请参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationParams []*ApplyParam `json:"ApplicationParams,omitnil" name:"ApplicationParams"`

	// 申请原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 申请单状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nodes []*StatusNode `json:"Nodes,omitnil" name:"Nodes"`

	// 正在审批的节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApprovingNodeId *string `json:"ApprovingNodeId,omitnil" name:"ApprovingNodeId"`

	// 更新时间，时间格式：2021-12-12 10:12:10	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetBpaasApproveDetailResponse struct {
	*tchttp.BaseResponse
	Response *GetBpaasApproveDetailResponseParams `json:"Response"`
}

func (r *GetBpaasApproveDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBpaasApproveDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OutApproveBpaasApplicationRequestParams struct {
	// 状态  1:通过  2:拒绝
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 审批单id
	ApproveId *uint64 `json:"ApproveId,omitnil" name:"ApproveId"`

	// 审批意见
	Msg *string `json:"Msg,omitnil" name:"Msg"`
}

type OutApproveBpaasApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 状态  1:通过  2:拒绝
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 审批单id
	ApproveId *uint64 `json:"ApproveId,omitnil" name:"ApproveId"`

	// 审批意见
	Msg *string `json:"Msg,omitnil" name:"Msg"`
}

func (r *OutApproveBpaasApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OutApproveBpaasApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "ApproveId")
	delete(f, "Msg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OutApproveBpaasApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OutApproveBpaasApplicationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type OutApproveBpaasApplicationResponse struct {
	*tchttp.BaseResponse
	Response *OutApproveBpaasApplicationResponseParams `json:"Response"`
}

func (r *OutApproveBpaasApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OutApproveBpaasApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Scf struct {
	// Scf函数地域id
	ScfRegion *string `json:"ScfRegion,omitnil" name:"ScfRegion"`

	// Scf函数地域
	ScfRegionName *string `json:"ScfRegionName,omitnil" name:"ScfRegionName"`

	// Scf函数名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScfName *string `json:"ScfName,omitnil" name:"ScfName"`

	// Scf函数入参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*ScfParam `json:"Params,omitnil" name:"Params"`
}

type ScfParam struct {
	// 参数Key
	Key *string `json:"Key,omitnil" name:"Key"`

	// 参数类型 1用户输入 2预设参数 3表单参数
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 参数值
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 参数描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type StatusNode struct {
	// 节点id
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`

	// 节点名称
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// 节点类型 1:审批节点 2:执行节点 3:条件节点
	NodeType *uint64 `json:"NodeType,omitnil" name:"NodeType"`

	// 下一个节点
	NextNode *string `json:"NextNode,omitnil" name:"NextNode"`

	// 审批意见模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Opinion *ApproveOpinion `json:"Opinion,omitnil" name:"Opinion"`

	// scf函数名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScfName *string `json:"ScfName,omitnil" name:"ScfName"`

	// 状态（0：待审批，1：审批通过，2：拒绝，3：scf执行失败，4：scf执行成功）18: 外部审批中
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubStatus *uint64 `json:"SubStatus,omitnil" name:"SubStatus"`

	// 审批节点审批人
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApprovedUin []*uint64 `json:"ApprovedUin,omitnil" name:"ApprovedUin"`

	// 审批时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 审批意见信息 审批节点:审批人意见  执行节点:scf函数执行日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 有权限审批该节点的uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users *ApproveUser `json:"Users,omitnil" name:"Users"`

	// 是否有权限审批该节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsApprove *bool `json:"IsApprove,omitnil" name:"IsApprove"`

	// 审批id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveId *string `json:"ApproveId,omitnil" name:"ApproveId"`

	// 审批方式 0或签 1会签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMethod *uint64 `json:"ApproveMethod,omitnil" name:"ApproveMethod"`

	// 审批节点审批类型，1人工审批 2自动通过 3自动决绝 4外部审批scf
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveType *uint64 `json:"ApproveType,omitnil" name:"ApproveType"`

	// 外部审批类型 scf:0或null ; CKafka:1
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallMethod *uint64 `json:"CallMethod,omitnil" name:"CallMethod"`

	// CKafka - 接入资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataHubId *string `json:"DataHubId,omitnil" name:"DataHubId"`

	// CKafka - 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// CKafka - 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	CKafkaRegion *string `json:"CKafkaRegion,omitnil" name:"CKafkaRegion"`

	// 外部审批Url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalUrl *string `json:"ExternalUrl,omitnil" name:"ExternalUrl"`

	// 并行节点 3-4
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParallelNodes *string `json:"ParallelNodes,omitnil" name:"ParallelNodes"`

	// scf拒绝时返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RejectedCloudFunctionMsg *string `json:"RejectedCloudFunctionMsg,omitnil" name:"RejectedCloudFunctionMsg"`

	// 上一个节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrevNode *string `json:"PrevNode,omitnil" name:"PrevNode"`
}