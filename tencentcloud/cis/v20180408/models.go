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

package v20180408

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Container struct {
	// 容器启动命令
	Command *string `json:"Command,omitnil" name:"Command"`

	// 容器启动参数
	Args []*string `json:"Args,omitnil" name:"Args"`

	// 容器环境变量
	EnvironmentVars []*EnvironmentVar `json:"EnvironmentVars,omitnil" name:"EnvironmentVars"`

	// 镜像
	Image *string `json:"Image,omitnil" name:"Image"`

	// 容器名，由小写字母、数字和 - 组成，由小写字母开头，小写字母或数字结尾，且长度不超过 63个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// CPU，单位：核
	Cpu *float64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存，单位：Gi
	Memory *float64 `json:"Memory,omitnil" name:"Memory"`

	// 重启次数
	RestartCount *uint64 `json:"RestartCount,omitnil" name:"RestartCount"`

	// 当前状态
	CurrentState *ContainerState `json:"CurrentState,omitnil" name:"CurrentState"`

	// 上一次状态
	PreviousState *ContainerState `json:"PreviousState,omitnil" name:"PreviousState"`

	// 容器工作目录
	WorkingDir *string `json:"WorkingDir,omitnil" name:"WorkingDir"`

	// 容器ID
	ContainerId *string `json:"ContainerId,omitnil" name:"ContainerId"`
}

type ContainerInstance struct {
	// 容器实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 容器实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 容器实例所属VpcId
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 容器实例所属SubnetId
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 容器实例状态
	State *string `json:"State,omitnil" name:"State"`

	// 容器列表
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`

	// 重启策略
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 启动时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Vpc名称
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// VpcCidr
	VpcCidr *string `json:"VpcCidr,omitnil" name:"VpcCidr"`

	// SubnetName
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 子网Cidr
	SubnetCidr *string `json:"SubnetCidr,omitnil" name:"SubnetCidr"`

	// 内网IP
	LanIp *string `json:"LanIp,omitnil" name:"LanIp"`
}

type ContainerLog struct {
	// 容器名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 日志
	Log *string `json:"Log,omitnil" name:"Log"`

	// 日志记录时间
	Time *string `json:"Time,omitnil" name:"Time"`
}

type ContainerState struct {
	// 容器运行开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 容器状态
	State *string `json:"State,omitnil" name:"State"`

	// 状态详情
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 容器运行结束时间
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`

	// 容器运行退出码
	ExitCode *int64 `json:"ExitCode,omitnil" name:"ExitCode"`
}

// Predefined struct for user
type CreateContainerInstanceRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// vpcId
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// subnetId
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 容器实例名称，由小写字母、数字和 - 组成，由小写字母开头，小写字母或数字结尾，且长度不超过 40个字符
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 重启策略（Always,OnFailure,Never）
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 容器列表
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`
}

type CreateContainerInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// vpcId
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// subnetId
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 容器实例名称，由小写字母、数字和 - 组成，由小写字母开头，小写字母或数字结尾，且长度不超过 40个字符
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 重启策略（Always,OnFailure,Never）
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 容器列表
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`
}

func (r *CreateContainerInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContainerInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceName")
	delete(f, "RestartPolicy")
	delete(f, "Containers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateContainerInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContainerInstanceResponseParams struct {
	// 容器实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateContainerInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateContainerInstanceResponseParams `json:"Response"`
}

func (r *CreateContainerInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContainerInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContainerInstanceRequestParams struct {
	// 容器实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type DeleteContainerInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 容器实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

func (r *DeleteContainerInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteContainerInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteContainerInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContainerInstanceResponseParams struct {
	// 操作信息
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteContainerInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteContainerInstanceResponseParams `json:"Response"`
}

func (r *DeleteContainerInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteContainerInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerInstanceEventsRequestParams struct {
	// 容器实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type DescribeContainerInstanceEventsRequest struct {
	*tchttp.BaseRequest
	
	// 容器实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

func (r *DescribeContainerInstanceEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerInstanceEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContainerInstanceEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerInstanceEventsResponseParams struct {
	// 容器实例事件列表
	EventList []*Event `json:"EventList,omitnil" name:"EventList"`

	// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeContainerInstanceEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContainerInstanceEventsResponseParams `json:"Response"`
}

func (r *DescribeContainerInstanceEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerInstanceEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerInstanceRequestParams struct {
	// 容器实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type DescribeContainerInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 容器实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

func (r *DescribeContainerInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContainerInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerInstanceResponseParams struct {
	// 容器实例详细信息
	ContainerInstance *ContainerInstance `json:"ContainerInstance,omitnil" name:"ContainerInstance"`

	// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeContainerInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContainerInstanceResponseParams `json:"Response"`
}

func (r *DescribeContainerInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerInstancesRequestParams struct {
	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件。
	// - Zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。
	// - VpcId - String - 是否必填：否 -（过滤条件）按照VpcId过滤。
	// - InstanceName - String - 是否必填：否 -（过滤条件）按照容器实例名称做模糊查询。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeContainerInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件。
	// - Zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。
	// - VpcId - String - 是否必填：否 -（过滤条件）按照VpcId过滤。
	// - InstanceName - String - 是否必填：否 -（过滤条件）按照容器实例名称做模糊查询。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeContainerInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContainerInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerInstancesResponseParams struct {
	// 容器实例列表
	ContainerInstanceList []*ContainerInstance `json:"ContainerInstanceList,omitnil" name:"ContainerInstanceList"`

	// 容器实例总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeContainerInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContainerInstancesResponseParams `json:"Response"`
}

func (r *DescribeContainerInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerLogRequestParams struct {
	// 容器实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitnil" name:"ContainerName"`

	// 日志显示尾部行数
	Tail *uint64 `json:"Tail,omitnil" name:"Tail"`

	// 日志起始时间
	SinceTime *string `json:"SinceTime,omitnil" name:"SinceTime"`
}

type DescribeContainerLogRequest struct {
	*tchttp.BaseRequest
	
	// 容器实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitnil" name:"ContainerName"`

	// 日志显示尾部行数
	Tail *uint64 `json:"Tail,omitnil" name:"Tail"`

	// 日志起始时间
	SinceTime *string `json:"SinceTime,omitnil" name:"SinceTime"`
}

func (r *DescribeContainerLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "ContainerName")
	delete(f, "Tail")
	delete(f, "SinceTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContainerLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerLogResponseParams struct {
	// 容器日志数组
	ContainerLogList []*ContainerLog `json:"ContainerLogList,omitnil" name:"ContainerLogList"`

	// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeContainerLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContainerLogResponseParams `json:"Response"`
}

func (r *DescribeContainerLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvironmentVar struct {
	// 环境变量名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 环境变量值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Event struct {
	// 事件首次出现时间
	FirstSeen *string `json:"FirstSeen,omitnil" name:"FirstSeen"`

	// 事件上次出现时间
	LastSeen *string `json:"LastSeen,omitnil" name:"LastSeen"`

	// 事件等级
	Level *string `json:"Level,omitnil" name:"Level"`

	// 事件出现次数
	Count *string `json:"Count,omitnil" name:"Count"`

	// 事件出现原因
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 事件消息
	Message *string `json:"Message,omitnil" name:"Message"`
}

type Filter struct {
	// 过滤字段，可选值 - Zone，VpcId，InstanceName
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤值列表
	ValueList []*string `json:"ValueList,omitnil" name:"ValueList"`
}

// Predefined struct for user
type InquiryPriceCreateCisRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// CPU，单位：核
	Cpu *float64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存，单位：Gi
	Memory *float64 `json:"Memory,omitnil" name:"Memory"`
}

type InquiryPriceCreateCisRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// CPU，单位：核
	Cpu *float64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存，单位：Gi
	Memory *float64 `json:"Memory,omitnil" name:"Memory"`
}

func (r *InquiryPriceCreateCisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateCisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Cpu")
	delete(f, "Memory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateCisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateCisResponseParams struct {
	// 价格
	Price *Price `json:"Price,omitnil" name:"Price"`

	// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceCreateCisResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateCisResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateCisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateCisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Price struct {
	// 原价，单位：元
	DiscountPrice *float64 `json:"DiscountPrice,omitnil" name:"DiscountPrice"`

	// 折扣价，单位：元
	OriginalPrice *float64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`
}