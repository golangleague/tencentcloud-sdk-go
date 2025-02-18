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

package v20201214

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CbsInfo struct {
	// cbs存储大小，单位TB
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// cbs存储类型，默认为SSD
	Type *string `json:"Type,omitnil" name:"Type"`
}

type CosCapacity struct {
	// 已购cos的总容量大小，单位GB
	TotalCapacity *float64 `json:"TotalCapacity,omitnil" name:"TotalCapacity"`

	// 剩余可用cos的容量大小，单位GB
	TotalFreeCapacity *float64 `json:"TotalFreeCapacity,omitnil" name:"TotalFreeCapacity"`

	// 已用cos的容量大小，单位GB
	TotalUsedCapacity *float64 `json:"TotalUsedCapacity,omitnil" name:"TotalUsedCapacity"`
}

type CosInfo struct {
	// COS存储大小，单位TB
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// COS存储类型，默认为cos
	Type *string `json:"Type,omitnil" name:"Type"`
}

// Predefined struct for user
type CreateDedicatedClusterOrderRequestParams struct {
	// 专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// order关联的专用集群类型数组
	DedicatedClusterTypes []*DedicatedClusterTypeInfo `json:"DedicatedClusterTypes,omitnil" name:"DedicatedClusterTypes"`

	// order关联的cos存储信息
	CosInfo *CosInfo `json:"CosInfo,omitnil" name:"CosInfo"`

	// order关联的cbs存储信息
	CbsInfo *CbsInfo `json:"CbsInfo,omitnil" name:"CbsInfo"`

	// 购买来源，默认为cloudApi
	PurchaseSource *string `json:"PurchaseSource,omitnil" name:"PurchaseSource"`

	// 当调用API接口提交订单时，需要提交DedicatedClusterOrderId
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil" name:"DedicatedClusterOrderId"`
}

type CreateDedicatedClusterOrderRequest struct {
	*tchttp.BaseRequest
	
	// 专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// order关联的专用集群类型数组
	DedicatedClusterTypes []*DedicatedClusterTypeInfo `json:"DedicatedClusterTypes,omitnil" name:"DedicatedClusterTypes"`

	// order关联的cos存储信息
	CosInfo *CosInfo `json:"CosInfo,omitnil" name:"CosInfo"`

	// order关联的cbs存储信息
	CbsInfo *CbsInfo `json:"CbsInfo,omitnil" name:"CbsInfo"`

	// 购买来源，默认为cloudApi
	PurchaseSource *string `json:"PurchaseSource,omitnil" name:"PurchaseSource"`

	// 当调用API接口提交订单时，需要提交DedicatedClusterOrderId
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil" name:"DedicatedClusterOrderId"`
}

func (r *CreateDedicatedClusterOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "DedicatedClusterTypes")
	delete(f, "CosInfo")
	delete(f, "CbsInfo")
	delete(f, "PurchaseSource")
	delete(f, "DedicatedClusterOrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDedicatedClusterOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDedicatedClusterOrderResponseParams struct {
	// 专用集群订单id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil" name:"DedicatedClusterOrderId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDedicatedClusterOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateDedicatedClusterOrderResponseParams `json:"Response"`
}

func (r *CreateDedicatedClusterOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDedicatedClusterRequestParams struct {
	// 专用集群所属的SiteId
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`

	// 专用集群的名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 专用集群所属的可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 专用集群的描述
	Description *string `json:"Description,omitnil" name:"Description"`
}

type CreateDedicatedClusterRequest struct {
	*tchttp.BaseRequest
	
	// 专用集群所属的SiteId
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`

	// 专用集群的名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 专用集群所属的可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 专用集群的描述
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *CreateDedicatedClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteId")
	delete(f, "Name")
	delete(f, "Zone")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDedicatedClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDedicatedClusterResponseParams struct {
	// 创建的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDedicatedClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateDedicatedClusterResponseParams `json:"Response"`
}

func (r *CreateDedicatedClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSiteRequestParams struct {
	// 站点名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 站点所在国家
	Country *string `json:"Country,omitnil" name:"Country"`

	// 站点所在省份
	Province *string `json:"Province,omitnil" name:"Province"`

	// 站点所在城市
	City *string `json:"City,omitnil" name:"City"`

	// 站点所在地区的详细地址信息
	AddressLine *string `json:"AddressLine,omitnil" name:"AddressLine"`

	// 站点描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 注意事项
	Note *string `json:"Note,omitnil" name:"Note"`

	// 您将使用光纤类型将CDC设备连接到网络。有单模和多模两种选项。
	FiberType *string `json:"FiberType,omitnil" name:"FiberType"`

	// 您将CDC连接到网络时采用的光学标准。此字段取决于上行链路速度、光纤类型和到上游设备的距离。
	OpticalStandard *string `json:"OpticalStandard,omitnil" name:"OpticalStandard"`

	// 电源连接器类型
	PowerConnectors *string `json:"PowerConnectors,omitnil" name:"PowerConnectors"`

	// 从机架上方还是下方供电。
	PowerFeedDrop *string `json:"PowerFeedDrop,omitnil" name:"PowerFeedDrop"`

	// 最大承重(KG)
	MaxWeight *int64 `json:"MaxWeight,omitnil" name:"MaxWeight"`

	// 功耗(KW)
	PowerDrawKva *int64 `json:"PowerDrawKva,omitnil" name:"PowerDrawKva"`

	// 网络到腾讯云Region区域的上行链路速度
	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitnil" name:"UplinkSpeedGbps"`

	// 将CDC连接到网络时，每台CDC网络设备(每个机架 2 台设备)使用的上行链路数量。
	UplinkCount *int64 `json:"UplinkCount,omitnil" name:"UplinkCount"`

	// 是否满足下面环境条件：
	// 1、场地没有材料要求或验收标准会影响 CDC 设备配送和安装。
	// 2、确定的机架位置包含:
	// 温度范围为 41 到 104°F (5 到 40°C)。
	// 湿度范围为 10°F (-12°C)和 8% RH (相对湿度)到 70°F(21°C)和 80% RH。
	// 机架位置的气流方向为从前向后，且应具有足够的 CFM (每分钟立方英尺)。CFM 必须是 CDC 配置的 kVA 功耗值的 145.8 倍。
	ConditionRequirement *bool `json:"ConditionRequirement,omitnil" name:"ConditionRequirement"`

	// 是否满足下面的尺寸条件：
	// 您的装货站台可以容纳一个机架箱(高 x 宽 x 深 = 94" x 54" x 48")。
	// 您可以提供从机架(高 x 宽 x 深 = 80" x 24" x 48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。
	// 在最终的 CDC安置位置，前部间隙可以为 48" 或更大，后部间隙可以为 24" 或更大。
	DimensionRequirement *bool `json:"DimensionRequirement,omitnil" name:"DimensionRequirement"`

	// 是否提供冗余的上游设备(交换机或路由器)，以便两台  网络设备都能连接到网络设备。
	RedundantNetworking *bool `json:"RedundantNetworking,omitnil" name:"RedundantNetworking"`

	// 站点所在地区的邮编
	PostalCode *int64 `json:"PostalCode,omitnil" name:"PostalCode"`

	// 站点所在地区的详细地址信息（补充）
	OptionalAddressLine *string `json:"OptionalAddressLine,omitnil" name:"OptionalAddressLine"`

	// 是否需要腾讯云团队协助完成机架支撑工作
	NeedHelp *bool `json:"NeedHelp,omitnil" name:"NeedHelp"`

	// 是否电源冗余
	RedundantPower *bool `json:"RedundantPower,omitnil" name:"RedundantPower"`

	// 上游断路器是否具备
	BreakerRequirement *bool `json:"BreakerRequirement,omitnil" name:"BreakerRequirement"`
}

type CreateSiteRequest struct {
	*tchttp.BaseRequest
	
	// 站点名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 站点所在国家
	Country *string `json:"Country,omitnil" name:"Country"`

	// 站点所在省份
	Province *string `json:"Province,omitnil" name:"Province"`

	// 站点所在城市
	City *string `json:"City,omitnil" name:"City"`

	// 站点所在地区的详细地址信息
	AddressLine *string `json:"AddressLine,omitnil" name:"AddressLine"`

	// 站点描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 注意事项
	Note *string `json:"Note,omitnil" name:"Note"`

	// 您将使用光纤类型将CDC设备连接到网络。有单模和多模两种选项。
	FiberType *string `json:"FiberType,omitnil" name:"FiberType"`

	// 您将CDC连接到网络时采用的光学标准。此字段取决于上行链路速度、光纤类型和到上游设备的距离。
	OpticalStandard *string `json:"OpticalStandard,omitnil" name:"OpticalStandard"`

	// 电源连接器类型
	PowerConnectors *string `json:"PowerConnectors,omitnil" name:"PowerConnectors"`

	// 从机架上方还是下方供电。
	PowerFeedDrop *string `json:"PowerFeedDrop,omitnil" name:"PowerFeedDrop"`

	// 最大承重(KG)
	MaxWeight *int64 `json:"MaxWeight,omitnil" name:"MaxWeight"`

	// 功耗(KW)
	PowerDrawKva *int64 `json:"PowerDrawKva,omitnil" name:"PowerDrawKva"`

	// 网络到腾讯云Region区域的上行链路速度
	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitnil" name:"UplinkSpeedGbps"`

	// 将CDC连接到网络时，每台CDC网络设备(每个机架 2 台设备)使用的上行链路数量。
	UplinkCount *int64 `json:"UplinkCount,omitnil" name:"UplinkCount"`

	// 是否满足下面环境条件：
	// 1、场地没有材料要求或验收标准会影响 CDC 设备配送和安装。
	// 2、确定的机架位置包含:
	// 温度范围为 41 到 104°F (5 到 40°C)。
	// 湿度范围为 10°F (-12°C)和 8% RH (相对湿度)到 70°F(21°C)和 80% RH。
	// 机架位置的气流方向为从前向后，且应具有足够的 CFM (每分钟立方英尺)。CFM 必须是 CDC 配置的 kVA 功耗值的 145.8 倍。
	ConditionRequirement *bool `json:"ConditionRequirement,omitnil" name:"ConditionRequirement"`

	// 是否满足下面的尺寸条件：
	// 您的装货站台可以容纳一个机架箱(高 x 宽 x 深 = 94" x 54" x 48")。
	// 您可以提供从机架(高 x 宽 x 深 = 80" x 24" x 48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。
	// 在最终的 CDC安置位置，前部间隙可以为 48" 或更大，后部间隙可以为 24" 或更大。
	DimensionRequirement *bool `json:"DimensionRequirement,omitnil" name:"DimensionRequirement"`

	// 是否提供冗余的上游设备(交换机或路由器)，以便两台  网络设备都能连接到网络设备。
	RedundantNetworking *bool `json:"RedundantNetworking,omitnil" name:"RedundantNetworking"`

	// 站点所在地区的邮编
	PostalCode *int64 `json:"PostalCode,omitnil" name:"PostalCode"`

	// 站点所在地区的详细地址信息（补充）
	OptionalAddressLine *string `json:"OptionalAddressLine,omitnil" name:"OptionalAddressLine"`

	// 是否需要腾讯云团队协助完成机架支撑工作
	NeedHelp *bool `json:"NeedHelp,omitnil" name:"NeedHelp"`

	// 是否电源冗余
	RedundantPower *bool `json:"RedundantPower,omitnil" name:"RedundantPower"`

	// 上游断路器是否具备
	BreakerRequirement *bool `json:"BreakerRequirement,omitnil" name:"BreakerRequirement"`
}

func (r *CreateSiteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSiteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Country")
	delete(f, "Province")
	delete(f, "City")
	delete(f, "AddressLine")
	delete(f, "Description")
	delete(f, "Note")
	delete(f, "FiberType")
	delete(f, "OpticalStandard")
	delete(f, "PowerConnectors")
	delete(f, "PowerFeedDrop")
	delete(f, "MaxWeight")
	delete(f, "PowerDrawKva")
	delete(f, "UplinkSpeedGbps")
	delete(f, "UplinkCount")
	delete(f, "ConditionRequirement")
	delete(f, "DimensionRequirement")
	delete(f, "RedundantNetworking")
	delete(f, "PostalCode")
	delete(f, "OptionalAddressLine")
	delete(f, "NeedHelp")
	delete(f, "RedundantPower")
	delete(f, "BreakerRequirement")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSiteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSiteResponseParams struct {
	// 创建Site生成的id
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSiteResponse struct {
	*tchttp.BaseResponse
	Response *CreateSiteResponseParams `json:"Response"`
}

func (r *CreateSiteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DedicatedCluster struct {
	// 专用集群id。如"cluster-xxxxx"。
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// 专用集群所属可用区名称。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 专用集群的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 专用集群的名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 专用集群的生命周期。如"PENDING"。
	LifecycleStatus *string `json:"LifecycleStatus,omitnil" name:"LifecycleStatus"`

	// 专用集群的创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 专用集群所属的站点id。
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`
}

type DedicatedClusterInstanceType struct {
	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 规格名称
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 网卡类型，例如：25代表25G网卡
	NetworkCard *int64 `json:"NetworkCard,omitnil" name:"NetworkCard"`

	// 实例的CPU核数，单位：核。
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// 实例内存容量，单位：`GB`。
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// 实例机型系列。
	InstanceFamily *string `json:"InstanceFamily,omitnil" name:"InstanceFamily"`

	// 机型名称。
	TypeName *string `json:"TypeName,omitnil" name:"TypeName"`

	// 本地存储块数量。
	StorageBlockAmount *int64 `json:"StorageBlockAmount,omitnil" name:"StorageBlockAmount"`

	// 内网带宽，单位Gbps。
	InstanceBandwidth *float64 `json:"InstanceBandwidth,omitnil" name:"InstanceBandwidth"`

	// 网络收发包能力，单位万PPS。
	InstancePps *int64 `json:"InstancePps,omitnil" name:"InstancePps"`

	// 处理器型号。
	CpuType *string `json:"CpuType,omitnil" name:"CpuType"`

	// 实例的GPU数量。
	Gpu *int64 `json:"Gpu,omitnil" name:"Gpu"`

	// 实例的FPGA数量。
	Fpga *int64 `json:"Fpga,omitnil" name:"Fpga"`

	// 机型描述
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 实例是否售卖。取值范围： <br><li>SELL：表示实例可购买<br><li>SOLD_OUT：表示实例已售罄。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type DedicatedClusterOrder struct {
	// 专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// 专用集群类型id（移到下一层级，已经废弃，后续将删除）
	DedicatedClusterTypeId *string `json:"DedicatedClusterTypeId,omitnil" name:"DedicatedClusterTypeId"`

	// 支持的存储类型列表（移到下一层级，已经废弃，后续将删除）
	SupportedStorageType []*string `json:"SupportedStorageType,omitnil" name:"SupportedStorageType"`

	// 支持的上连交换机的链路传输速率(GiB)（移到下一层级，已经废弃，后续将删除）
	SupportedUplinkSpeed []*int64 `json:"SupportedUplinkSpeed,omitnil" name:"SupportedUplinkSpeed"`

	// 支持的实例族列表（移到下一层级，已经废弃，后续将删除）
	SupportedInstanceFamily []*string `json:"SupportedInstanceFamily,omitnil" name:"SupportedInstanceFamily"`

	// 地板承重要求(KG)
	Weight *int64 `json:"Weight,omitnil" name:"Weight"`

	// 功率要求(KW)
	PowerDraw *float64 `json:"PowerDraw,omitnil" name:"PowerDraw"`

	// 订单状态
	OrderStatus *string `json:"OrderStatus,omitnil" name:"OrderStatus"`

	// 订单创建的时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 大订单ID
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil" name:"DedicatedClusterOrderId"`

	// 订单类型，创建CREATE或扩容EXTEND
	Action *string `json:"Action,omitnil" name:"Action"`

	// 子订单详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DedicatedClusterOrderItems []*DedicatedClusterOrderItem `json:"DedicatedClusterOrderItems,omitnil" name:"DedicatedClusterOrderItems"`

	// cpu值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// mem值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mem *int64 `json:"Mem,omitnil" name:"Mem"`

	// gpu值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *int64 `json:"Gpu,omitnil" name:"Gpu"`

	// 0代表未支付，1代表已支付
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayStatus *int64 `json:"PayStatus,omitnil" name:"PayStatus"`

	// 支付方式，一次性、按月、按年
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayType *string `json:"PayType,omitnil" name:"PayType"`

	// 购买时长的单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 购买时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 订单类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`
}

type DedicatedClusterOrderItem struct {
	// 专用集群类型id
	DedicatedClusterTypeId *string `json:"DedicatedClusterTypeId,omitnil" name:"DedicatedClusterTypeId"`

	// 支持的存储类型列表
	SupportedStorageType []*string `json:"SupportedStorageType,omitnil" name:"SupportedStorageType"`

	// 支持的上连交换机的链路传输速率(GiB)
	SupportedUplinkSpeed []*int64 `json:"SupportedUplinkSpeed,omitnil" name:"SupportedUplinkSpeed"`

	// 支持的实例族列表
	SupportedInstanceFamily []*string `json:"SupportedInstanceFamily,omitnil" name:"SupportedInstanceFamily"`

	// 地板承重要求(KG)
	Weight *int64 `json:"Weight,omitnil" name:"Weight"`

	// 功率要求(KW)
	PowerDraw *float64 `json:"PowerDraw,omitnil" name:"PowerDraw"`

	// 订单状态
	SubOrderStatus *string `json:"SubOrderStatus,omitnil" name:"SubOrderStatus"`

	// 订单创建的时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 子订单ID
	SubOrderId *string `json:"SubOrderId,omitnil" name:"SubOrderId"`

	// 关联的集群规格数量
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 规格简单描述
	Name *string `json:"Name,omitnil" name:"Name"`

	// 规格详细描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// CPU数
	TotalCpu *int64 `json:"TotalCpu,omitnil" name:"TotalCpu"`

	// 内存数
	TotalMem *int64 `json:"TotalMem,omitnil" name:"TotalMem"`

	// GPU数
	TotalGpu *int64 `json:"TotalGpu,omitnil" name:"TotalGpu"`

	// 规格英文名
	TypeName *string `json:"TypeName,omitnil" name:"TypeName"`

	// 规格展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeFormat *string `json:"ComputeFormat,omitnil" name:"ComputeFormat"`

	// 规格类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeFamily *string `json:"TypeFamily,omitnil" name:"TypeFamily"`

	// 0未支付，1已支付
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubOrderPayStatus *int64 `json:"SubOrderPayStatus,omitnil" name:"SubOrderPayStatus"`
}

type DedicatedClusterType struct {
	// 配置id
	DedicatedClusterTypeId *string `json:"DedicatedClusterTypeId,omitnil" name:"DedicatedClusterTypeId"`

	// 配置描述，对应描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 配置名称，对应计算资源类型
	Name *string `json:"Name,omitnil" name:"Name"`

	// 创建配置的时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 支持的存储类型列表
	SupportedStorageType []*string `json:"SupportedStorageType,omitnil" name:"SupportedStorageType"`

	// 支持的上连交换机的链路传输速率
	SupportedUplinkGiB []*int64 `json:"SupportedUplinkGiB,omitnil" name:"SupportedUplinkGiB"`

	// 支持的实例族列表
	SupportedInstanceFamily []*string `json:"SupportedInstanceFamily,omitnil" name:"SupportedInstanceFamily"`

	// 地板承重要求(KG)
	Weight *int64 `json:"Weight,omitnil" name:"Weight"`

	// 功率要求(KW)
	PowerDrawKva *float64 `json:"PowerDrawKva,omitnil" name:"PowerDrawKva"`

	// 显示计算资源规格详情，存储等资源不显示；对应规格
	ComputeFormatDesc *string `json:"ComputeFormatDesc,omitnil" name:"ComputeFormatDesc"`
}

type DedicatedClusterTypeInfo struct {
	// 集群类型Id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 集群类型个数
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

// Predefined struct for user
type DeleteDedicatedClustersRequestParams struct {
	// 要删除的专用集群id
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil" name:"DedicatedClusterIds"`
}

type DeleteDedicatedClustersRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的专用集群id
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil" name:"DedicatedClusterIds"`
}

func (r *DeleteDedicatedClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDedicatedClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDedicatedClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDedicatedClustersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDedicatedClustersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDedicatedClustersResponseParams `json:"Response"`
}

func (r *DeleteDedicatedClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDedicatedClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSitesRequestParams struct {
	// 要删除的站点id列表
	SiteIds []*string `json:"SiteIds,omitnil" name:"SiteIds"`
}

type DeleteSitesRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的站点id列表
	SiteIds []*string `json:"SiteIds,omitnil" name:"SiteIds"`
}

func (r *DeleteSitesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSitesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSitesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSitesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSitesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSitesResponseParams `json:"Response"`
}

func (r *DeleteSitesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterCosCapacityRequestParams struct {
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`
}

type DescribeDedicatedClusterCosCapacityRequest struct {
	*tchttp.BaseRequest
	
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterCosCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterCosCapacityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterCosCapacityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterCosCapacityResponseParams struct {
	// 本集群内cos容量信息，单位：‘GB’
	CosCapacity *CosCapacity `json:"CosCapacity,omitnil" name:"CosCapacity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDedicatedClusterCosCapacityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterCosCapacityResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterCosCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterCosCapacityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterHostStatisticsRequestParams struct {
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// 宿主机id
	HostId *string `json:"HostId,omitnil" name:"HostId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间范围精度，1分钟/5分钟
	Period *string `json:"Period,omitnil" name:"Period"`
}

type DescribeDedicatedClusterHostStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// 宿主机id
	HostId *string `json:"HostId,omitnil" name:"HostId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间范围精度，1分钟/5分钟
	Period *string `json:"Period,omitnil" name:"Period"`
}

func (r *DescribeDedicatedClusterHostStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterHostStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "HostId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterHostStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterHostStatisticsResponseParams struct {
	// 该集群内宿主机的统计信息列表
	HostStatisticSet []*HostStatistic `json:"HostStatisticSet,omitnil" name:"HostStatisticSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDedicatedClusterHostStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterHostStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterHostStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterHostStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterHostsRequestParams struct {
	// 集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeDedicatedClusterHostsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeDedicatedClusterHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterHostsResponseParams struct {
	// 宿主机信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostInfoSet []*HostInfo `json:"HostInfoSet,omitnil" name:"HostInfoSet"`

	// 宿主机总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDedicatedClusterHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterHostsResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterInstanceTypesRequestParams struct {
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`
}

type DescribeDedicatedClusterInstanceTypesRequest struct {
	*tchttp.BaseRequest
	
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterInstanceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterInstanceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterInstanceTypesResponseParams struct {
	// 支持的实例规格列表
	DedicatedClusterInstanceTypeSet []*DedicatedClusterInstanceType `json:"DedicatedClusterInstanceTypeSet,omitnil" name:"DedicatedClusterInstanceTypeSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDedicatedClusterInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterInstanceTypesResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterOrdersRequestParams struct {
	// 按照专用集群id过滤
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil" name:"DedicatedClusterIds"`

	// 按照专用集群订单id过滤
	DedicatedClusterOrderIds *string `json:"DedicatedClusterOrderIds,omitnil" name:"DedicatedClusterOrderIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 订单状态为过滤条件：PENDING INCONSTRUCTION DELIVERING DELIVERED EXPIRED CANCELLED  OFFLINE
	Status *string `json:"Status,omitnil" name:"Status"`

	// 订单类型为过滤条件：CREATE  EXTEND
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`
}

type DescribeDedicatedClusterOrdersRequest struct {
	*tchttp.BaseRequest
	
	// 按照专用集群id过滤
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil" name:"DedicatedClusterIds"`

	// 按照专用集群订单id过滤
	DedicatedClusterOrderIds *string `json:"DedicatedClusterOrderIds,omitnil" name:"DedicatedClusterOrderIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 订单状态为过滤条件：PENDING INCONSTRUCTION DELIVERING DELIVERED EXPIRED CANCELLED  OFFLINE
	Status *string `json:"Status,omitnil" name:"Status"`

	// 订单类型为过滤条件：CREATE  EXTEND
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`
}

func (r *DescribeDedicatedClusterOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterIds")
	delete(f, "DedicatedClusterOrderIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "ActionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterOrdersResponseParams struct {
	// 专用集群订单列表
	DedicatedClusterOrderSet []*DedicatedClusterOrder `json:"DedicatedClusterOrderSet,omitnil" name:"DedicatedClusterOrderSet"`

	// 符合条件的专用集群订单总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDedicatedClusterOrdersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterOrdersResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterOverviewRequestParams struct {
	// 集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`
}

type DescribeDedicatedClusterOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterOverviewResponseParams struct {
	// 云服务器数量
	CvmCount *uint64 `json:"CvmCount,omitnil" name:"CvmCount"`

	// 宿主机数量
	HostCount *uint64 `json:"HostCount,omitnil" name:"HostCount"`

	// vpn通道状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpnConnectionState *string `json:"VpnConnectionState,omitnil" name:"VpnConnectionState"`

	// vpn网关监控数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpngwBandwidthData *VpngwBandwidthData `json:"VpngwBandwidthData,omitnil" name:"VpngwBandwidthData"`

	// 本地网关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalNetInfo *LocalNetInfo `json:"LocalNetInfo,omitnil" name:"LocalNetInfo"`

	// vpn网关通道监控数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpnConnectionBandwidthData []*VpngwBandwidthData `json:"VpnConnectionBandwidthData,omitnil" name:"VpnConnectionBandwidthData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDedicatedClusterOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterOverviewResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterTypesRequestParams struct {
	// 模糊匹配专用集群配置名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 待查询的专用集群配置id列表
	DedicatedClusterTypeIds []*string `json:"DedicatedClusterTypeIds,omitnil" name:"DedicatedClusterTypeIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 是否只查询计算规格类型
	IsCompute *bool `json:"IsCompute,omitnil" name:"IsCompute"`
}

type DescribeDedicatedClusterTypesRequest struct {
	*tchttp.BaseRequest
	
	// 模糊匹配专用集群配置名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 待查询的专用集群配置id列表
	DedicatedClusterTypeIds []*string `json:"DedicatedClusterTypeIds,omitnil" name:"DedicatedClusterTypeIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 是否只查询计算规格类型
	IsCompute *bool `json:"IsCompute,omitnil" name:"IsCompute"`
}

func (r *DescribeDedicatedClusterTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DedicatedClusterTypeIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IsCompute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterTypesResponseParams struct {
	// 专用集群配置列表
	DedicatedClusterTypeSet []*DedicatedClusterType `json:"DedicatedClusterTypeSet,omitnil" name:"DedicatedClusterTypeSet"`

	// 符合条件的个数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDedicatedClusterTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterTypesResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClustersRequestParams struct {
	// 按照一个或者多个实例ID查询。实例ID形如：`cluster-xxxxxxxx`
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil" name:"DedicatedClusterIds"`

	// 按照可用区名称过滤
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil" name:"SiteIds"`

	// 按照专用集群生命周期过滤
	LifecycleStatuses []*string `json:"LifecycleStatuses,omitnil" name:"LifecycleStatuses"`

	// 模糊匹配专用集群名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeDedicatedClustersRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个实例ID查询。实例ID形如：`cluster-xxxxxxxx`
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil" name:"DedicatedClusterIds"`

	// 按照可用区名称过滤
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil" name:"SiteIds"`

	// 按照专用集群生命周期过滤
	LifecycleStatuses []*string `json:"LifecycleStatuses,omitnil" name:"LifecycleStatuses"`

	// 模糊匹配专用集群名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeDedicatedClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterIds")
	delete(f, "Zones")
	delete(f, "SiteIds")
	delete(f, "LifecycleStatuses")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClustersResponseParams struct {
	// 符合查询条件的专用集群列表
	DedicatedClusterSet []*DedicatedCluster `json:"DedicatedClusterSet,omitnil" name:"DedicatedClusterSet"`

	// 符合条件的专用集群数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDedicatedClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClustersResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedSupportedZonesRequestParams struct {
	// 传入region列表
	Regions []*int64 `json:"Regions,omitnil" name:"Regions"`
}

type DescribeDedicatedSupportedZonesRequest struct {
	*tchttp.BaseRequest
	
	// 传入region列表
	Regions []*int64 `json:"Regions,omitnil" name:"Regions"`
}

func (r *DescribeDedicatedSupportedZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedSupportedZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Regions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedSupportedZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedSupportedZonesResponseParams struct {
	// 支持的可用区列表
	ZoneSet []*RegionZoneInfo `json:"ZoneSet,omitnil" name:"ZoneSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDedicatedSupportedZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedSupportedZonesResponseParams `json:"Response"`
}

func (r *DescribeDedicatedSupportedZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedSupportedZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesDetailRequestParams struct {
	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil" name:"SiteIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 按照站定名称模糊匹配
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DescribeSitesDetailRequest struct {
	*tchttp.BaseRequest
	
	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil" name:"SiteIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 按照站定名称模糊匹配
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *DescribeSitesDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSitesDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesDetailResponseParams struct {
	// 站点详情
	SiteDetailSet []*SiteDetail `json:"SiteDetailSet,omitnil" name:"SiteDetailSet"`

	// 符合条件的站点总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSitesDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSitesDetailResponseParams `json:"Response"`
}

func (r *DescribeSitesDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesRequestParams struct {
	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil" name:"SiteIds"`

	// 模糊匹配站点名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSitesRequest struct {
	*tchttp.BaseRequest
	
	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil" name:"SiteIds"`

	// 模糊匹配站点名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeSitesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteIds")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSitesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesResponseParams struct {
	// 符合查询条件的站点列表
	SiteSet []*Site `json:"SiteSet,omitnil" name:"SiteSet"`

	// 符合条件的站点数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSitesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSitesResponseParams `json:"Response"`
}

func (r *DescribeSitesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailData struct {
	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamps []*float64 `json:"Timestamps,omitnil" name:"Timestamps"`

	// 对应的具体值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*float64 `json:"Values,omitnil" name:"Values"`
}

type HostInfo struct {
	// 宿主机IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 云服务类型
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 宿主机运行状态
	HostStatus *string `json:"HostStatus,omitnil" name:"HostStatus"`

	// 宿主机类型
	HostType *string `json:"HostType,omitnil" name:"HostType"`

	// cpu可用数
	CpuAvailable *uint64 `json:"CpuAvailable,omitnil" name:"CpuAvailable"`

	// cpu总数
	CpuTotal *uint64 `json:"CpuTotal,omitnil" name:"CpuTotal"`

	// 内存可用数
	MemAvailable *uint64 `json:"MemAvailable,omitnil" name:"MemAvailable"`

	// 内存总数
	MemTotal *uint64 `json:"MemTotal,omitnil" name:"MemTotal"`

	// 运行时间
	RunTime *string `json:"RunTime,omitnil" name:"RunTime"`

	// 到期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 宿主机id
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostId *string `json:"HostId,omitnil" name:"HostId"`
}

type HostStatistic struct {
	// 宿主机规格
	HostType *string `json:"HostType,omitnil" name:"HostType"`

	// 宿主机机型系列
	HostFamily *string `json:"HostFamily,omitnil" name:"HostFamily"`

	// 宿主机的CPU核数，单位：核
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 宿主机内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 该规格宿主机的数量
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 平均cpu负载百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuAverage *float64 `json:"CpuAverage,omitnil" name:"CpuAverage"`

	// 平均内存使用率百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemAverage *float64 `json:"MemAverage,omitnil" name:"MemAverage"`

	// 平均网络流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetAverage *float64 `json:"NetAverage,omitnil" name:"NetAverage"`

	// cpu详细监控数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuDetailData *DetailData `json:"CpuDetailData,omitnil" name:"CpuDetailData"`

	// 内存详细数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemDetailData *DetailData `json:"MemDetailData,omitnil" name:"MemDetailData"`

	// 网络速率详细数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetRateDetailData *DetailData `json:"NetRateDetailData,omitnil" name:"NetRateDetailData"`

	// 网速包详细数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetPacketDetailData *DetailData `json:"NetPacketDetailData,omitnil" name:"NetPacketDetailData"`
}

type InBandwidth struct {
	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamps []*float64 `json:"Timestamps,omitnil" name:"Timestamps"`

	// 时间对应的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*float64 `json:"Values,omitnil" name:"Values"`
}

type LocalNetInfo struct {
	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 网络id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 路由信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BGPRoute *string `json:"BGPRoute,omitnil" name:"BGPRoute"`

	// 本地IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalIp *string `json:"LocalIp,omitnil" name:"LocalIp"`
}

// Predefined struct for user
type ModifyDedicatedClusterInfoRequestParams struct {
	// 本地专用集群ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// 集群的新名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 集群的新可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 集群的新描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 集群所在站点
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`
}

type ModifyDedicatedClusterInfoRequest struct {
	*tchttp.BaseRequest
	
	// 本地专用集群ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`

	// 集群的新名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 集群的新可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 集群的新描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 集群所在站点
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`
}

func (r *ModifyDedicatedClusterInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDedicatedClusterInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "Name")
	delete(f, "Zone")
	delete(f, "Description")
	delete(f, "SiteId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDedicatedClusterInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDedicatedClusterInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDedicatedClusterInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDedicatedClusterInfoResponseParams `json:"Response"`
}

func (r *ModifyDedicatedClusterInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDedicatedClusterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrderStatusRequestParams struct {
	// 要更新成的状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 大订单ID
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil" name:"DedicatedClusterOrderId"`

	// 小订单ID
	SubOrderIds []*string `json:"SubOrderIds,omitnil" name:"SubOrderIds"`
}

type ModifyOrderStatusRequest struct {
	*tchttp.BaseRequest
	
	// 要更新成的状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 大订单ID
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil" name:"DedicatedClusterOrderId"`

	// 小订单ID
	SubOrderIds []*string `json:"SubOrderIds,omitnil" name:"SubOrderIds"`
}

func (r *ModifyOrderStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrderStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "DedicatedClusterOrderId")
	delete(f, "SubOrderIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOrderStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrderStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyOrderStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOrderStatusResponseParams `json:"Response"`
}

func (r *ModifyOrderStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrderStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySiteDeviceInfoRequestParams struct {
	// 机房ID
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`

	// 您将使用光纤类型将CDC设备连接到网络。有单模和多模两种选项。
	FiberType *string `json:"FiberType,omitnil" name:"FiberType"`

	// 您将CDC连接到网络时采用的光学标准。此字段取决于上行链路速度、光纤类型和到上游设备的距离。
	OpticalStandard *string `json:"OpticalStandard,omitnil" name:"OpticalStandard"`

	// 电源连接器类型
	PowerConnectors *string `json:"PowerConnectors,omitnil" name:"PowerConnectors"`

	// 从机架上方还是下方供电。
	PowerFeedDrop *string `json:"PowerFeedDrop,omitnil" name:"PowerFeedDrop"`

	// 最大承重(KG)
	MaxWeight *int64 `json:"MaxWeight,omitnil" name:"MaxWeight"`

	// 功耗(KW)
	PowerDrawKva *int64 `json:"PowerDrawKva,omitnil" name:"PowerDrawKva"`

	// 网络到腾讯云Region区域的上行链路速度
	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitnil" name:"UplinkSpeedGbps"`

	// 将CDC连接到网络时，每台CDC网络设备(每个机架 2 台设备)使用的上行链路数量。
	UplinkCount *int64 `json:"UplinkCount,omitnil" name:"UplinkCount"`

	// 是否满足下面环境条件：
	// 1、场地没有材料要求或验收标准会影响 CDC 设备配送和安装。
	// 2、确定的机架位置包含:
	// 温度范围为 41 到 104°F (5 到 40°C)。
	// 湿度范围为 10°F (-12°C)和 8% RH (相对湿度)到 70°F(21°C)和 80% RH。
	// 机架位置的气流方向为从前向后，且应具有足够的 CFM (每分钟立方英尺)。CFM 必须是 CDC 配置的 kVA 功耗值的 145.8 倍。
	ConditionRequirement *bool `json:"ConditionRequirement,omitnil" name:"ConditionRequirement"`

	// 是否满足下面的尺寸条件：
	// 您的装货站台可以容纳一个机架箱(高 x 宽 x 深 = 94" x 54" x 48")。
	// 您可以提供从机架(高 x 宽 x 深 = 80" x 24" x 48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。
	// 在最终的 CDC安置位置，前部间隙可以为 48" 或更大，后部间隙可以为 24" 或更大。
	DimensionRequirement *bool `json:"DimensionRequirement,omitnil" name:"DimensionRequirement"`

	// 是否提供冗余的上游设备(交换机或路由器)，以便两台  网络设备都能连接到网络设备。
	RedundantNetworking *bool `json:"RedundantNetworking,omitnil" name:"RedundantNetworking"`

	// 是否需要腾讯云团队协助完成机架支撑工作
	NeedHelp *bool `json:"NeedHelp,omitnil" name:"NeedHelp"`

	// 是否电源冗余
	RedundantPower *bool `json:"RedundantPower,omitnil" name:"RedundantPower"`

	// 上游断路器是否具备
	BreakerRequirement *bool `json:"BreakerRequirement,omitnil" name:"BreakerRequirement"`
}

type ModifySiteDeviceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 机房ID
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`

	// 您将使用光纤类型将CDC设备连接到网络。有单模和多模两种选项。
	FiberType *string `json:"FiberType,omitnil" name:"FiberType"`

	// 您将CDC连接到网络时采用的光学标准。此字段取决于上行链路速度、光纤类型和到上游设备的距离。
	OpticalStandard *string `json:"OpticalStandard,omitnil" name:"OpticalStandard"`

	// 电源连接器类型
	PowerConnectors *string `json:"PowerConnectors,omitnil" name:"PowerConnectors"`

	// 从机架上方还是下方供电。
	PowerFeedDrop *string `json:"PowerFeedDrop,omitnil" name:"PowerFeedDrop"`

	// 最大承重(KG)
	MaxWeight *int64 `json:"MaxWeight,omitnil" name:"MaxWeight"`

	// 功耗(KW)
	PowerDrawKva *int64 `json:"PowerDrawKva,omitnil" name:"PowerDrawKva"`

	// 网络到腾讯云Region区域的上行链路速度
	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitnil" name:"UplinkSpeedGbps"`

	// 将CDC连接到网络时，每台CDC网络设备(每个机架 2 台设备)使用的上行链路数量。
	UplinkCount *int64 `json:"UplinkCount,omitnil" name:"UplinkCount"`

	// 是否满足下面环境条件：
	// 1、场地没有材料要求或验收标准会影响 CDC 设备配送和安装。
	// 2、确定的机架位置包含:
	// 温度范围为 41 到 104°F (5 到 40°C)。
	// 湿度范围为 10°F (-12°C)和 8% RH (相对湿度)到 70°F(21°C)和 80% RH。
	// 机架位置的气流方向为从前向后，且应具有足够的 CFM (每分钟立方英尺)。CFM 必须是 CDC 配置的 kVA 功耗值的 145.8 倍。
	ConditionRequirement *bool `json:"ConditionRequirement,omitnil" name:"ConditionRequirement"`

	// 是否满足下面的尺寸条件：
	// 您的装货站台可以容纳一个机架箱(高 x 宽 x 深 = 94" x 54" x 48")。
	// 您可以提供从机架(高 x 宽 x 深 = 80" x 24" x 48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。
	// 在最终的 CDC安置位置，前部间隙可以为 48" 或更大，后部间隙可以为 24" 或更大。
	DimensionRequirement *bool `json:"DimensionRequirement,omitnil" name:"DimensionRequirement"`

	// 是否提供冗余的上游设备(交换机或路由器)，以便两台  网络设备都能连接到网络设备。
	RedundantNetworking *bool `json:"RedundantNetworking,omitnil" name:"RedundantNetworking"`

	// 是否需要腾讯云团队协助完成机架支撑工作
	NeedHelp *bool `json:"NeedHelp,omitnil" name:"NeedHelp"`

	// 是否电源冗余
	RedundantPower *bool `json:"RedundantPower,omitnil" name:"RedundantPower"`

	// 上游断路器是否具备
	BreakerRequirement *bool `json:"BreakerRequirement,omitnil" name:"BreakerRequirement"`
}

func (r *ModifySiteDeviceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySiteDeviceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteId")
	delete(f, "FiberType")
	delete(f, "OpticalStandard")
	delete(f, "PowerConnectors")
	delete(f, "PowerFeedDrop")
	delete(f, "MaxWeight")
	delete(f, "PowerDrawKva")
	delete(f, "UplinkSpeedGbps")
	delete(f, "UplinkCount")
	delete(f, "ConditionRequirement")
	delete(f, "DimensionRequirement")
	delete(f, "RedundantNetworking")
	delete(f, "NeedHelp")
	delete(f, "RedundantPower")
	delete(f, "BreakerRequirement")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySiteDeviceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySiteDeviceInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySiteDeviceInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifySiteDeviceInfoResponseParams `json:"Response"`
}

func (r *ModifySiteDeviceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySiteDeviceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySiteInfoRequestParams struct {
	// 机房ID
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`

	// 站点名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 站点描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 注意事项
	Note *string `json:"Note,omitnil" name:"Note"`

	// 站点所在国家
	Country *string `json:"Country,omitnil" name:"Country"`

	// 站点所在省份
	Province *string `json:"Province,omitnil" name:"Province"`

	// 站点所在城市
	City *string `json:"City,omitnil" name:"City"`

	// 站点所在地区的邮编
	PostalCode *string `json:"PostalCode,omitnil" name:"PostalCode"`

	// 站点所在地区的详细地址信息
	AddressLine *string `json:"AddressLine,omitnil" name:"AddressLine"`
}

type ModifySiteInfoRequest struct {
	*tchttp.BaseRequest
	
	// 机房ID
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`

	// 站点名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 站点描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 注意事项
	Note *string `json:"Note,omitnil" name:"Note"`

	// 站点所在国家
	Country *string `json:"Country,omitnil" name:"Country"`

	// 站点所在省份
	Province *string `json:"Province,omitnil" name:"Province"`

	// 站点所在城市
	City *string `json:"City,omitnil" name:"City"`

	// 站点所在地区的邮编
	PostalCode *string `json:"PostalCode,omitnil" name:"PostalCode"`

	// 站点所在地区的详细地址信息
	AddressLine *string `json:"AddressLine,omitnil" name:"AddressLine"`
}

func (r *ModifySiteInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySiteInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Note")
	delete(f, "Country")
	delete(f, "Province")
	delete(f, "City")
	delete(f, "PostalCode")
	delete(f, "AddressLine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySiteInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySiteInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySiteInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifySiteInfoResponseParams `json:"Response"`
}

func (r *ModifySiteInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySiteInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutBandwidth struct {
	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamps []*float64 `json:"Timestamps,omitnil" name:"Timestamps"`

	// 对应时间的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*float64 `json:"Values,omitnil" name:"Values"`
}

type RegionZoneInfo struct {
	// Region id
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// ZoneInfo数组
	Zones []*ZoneInfo `json:"Zones,omitnil" name:"Zones"`
}

type Site struct {
	// 站点名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 站点id
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`

	// 站点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 站点创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type SiteDetail struct {
	// 站点id
	SiteId *string `json:"SiteId,omitnil" name:"SiteId"`

	// 站点名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 站点描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 站点创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 光纤类型
	FiberType *string `json:"FiberType,omitnil" name:"FiberType"`

	// 网络到腾讯云Region区域的上行链路速度
	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitnil" name:"UplinkSpeedGbps"`

	// 将CDC连接到网络时，每台CDC网络设备(每个机架 2 台设备)使用的上行链路数量。
	UplinkCount *int64 `json:"UplinkCount,omitnil" name:"UplinkCount"`

	// 将CDC连接到网络时采用的光学标准
	OpticalStandard *string `json:"OpticalStandard,omitnil" name:"OpticalStandard"`

	// 是否提供冗余的上游设备(交换机或路由器)，以便两台  网络设备都能连接到网络设备。
	RedundantNetworking *bool `json:"RedundantNetworking,omitnil" name:"RedundantNetworking"`

	// 电源连接器类型
	PowerConnectors *string `json:"PowerConnectors,omitnil" name:"PowerConnectors"`

	// 从机架上方还是下方供电。
	PowerFeedDrop *string `json:"PowerFeedDrop,omitnil" name:"PowerFeedDrop"`

	// 功耗(KW)
	PowerDrawKva *float64 `json:"PowerDrawKva,omitnil" name:"PowerDrawKva"`

	// 是否满足下面环境条件：
	// 1、场地没有材料要求或验收标准会影响 CDC 设备配送和安装。
	// 2、确定的机架位置包含:
	// 温度范围为 41 到 104°F (5 到 40°C)。
	// 湿度范围为 10°F (-12°C)和 8% RH (相对湿度)到 70°F(21°C)和 80% RH。
	// 机架位置的气流方向为从前向后，且应具有足够的 CFM (每分钟立方英尺)。CFM 必须是 CDC 配置的 kVA 功耗值的 145.8 倍。
	ConditionRequirement *bool `json:"ConditionRequirement,omitnil" name:"ConditionRequirement"`

	// 是否满足下面的尺寸条件：
	// 您的装货站台可以容纳一个机架箱(高 x 宽 x 深 = 94" x 54" x 48")。
	// 您可以提供从机架(高 x 宽 x 深 = 80" x 24" x 48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。
	// 在最终的 CDC安置位置，前部间隙可以为 48" 或更大，后部间隙可以为 24" 或更大。
	DimensionRequirement *bool `json:"DimensionRequirement,omitnil" name:"DimensionRequirement"`

	// 最大承重(KG)
	MaxWeight *int64 `json:"MaxWeight,omitnil" name:"MaxWeight"`

	// 站点地址
	AddressLine *string `json:"AddressLine,omitnil" name:"AddressLine"`

	// 站点所在地区的详细地址信息（补充）
	OptionalAddressLine *string `json:"OptionalAddressLine,omitnil" name:"OptionalAddressLine"`

	// 是否需要腾讯云团队协助完成机架支撑工作
	NeedHelp *bool `json:"NeedHelp,omitnil" name:"NeedHelp"`

	// 上游断路器是否具备
	BreakerRequirement *bool `json:"BreakerRequirement,omitnil" name:"BreakerRequirement"`

	// 是否电源冗余
	RedundantPower *bool `json:"RedundantPower,omitnil" name:"RedundantPower"`

	// 站点所在国家
	Country *string `json:"Country,omitnil" name:"Country"`

	// 站点所在省份
	Province *string `json:"Province,omitnil" name:"Province"`

	// 站点所在城市
	City *string `json:"City,omitnil" name:"City"`

	// 站点所在地区的邮编
	PostalCode *int64 `json:"PostalCode,omitnil" name:"PostalCode"`
}

type VpngwBandwidthData struct {
	// 出带宽流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutBandwidth *OutBandwidth `json:"OutBandwidth,omitnil" name:"OutBandwidth"`

	// 入带宽流量
	InBandwidth *InBandwidth `json:"InBandwidth,omitnil" name:"InBandwidth"`
}

type ZoneInfo struct {
	// 可用区名称
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 可用区描述
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 可用区ID
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 可用区状态，包含AVAILABLE和UNAVAILABLE。AVAILABLE代表可用，UNAVAILABLE代表不可用。
	ZoneState *string `json:"ZoneState,omitnil" name:"ZoneState"`
}