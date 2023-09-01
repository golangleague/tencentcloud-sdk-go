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

package v20190311

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateBotRequestParams struct {
	// 机器人名称
	BotName *string `json:"BotName,omitnil" name:"BotName"`

	// 机器人中文名称
	BotCnName *string `json:"BotCnName,omitnil" name:"BotCnName"`
}

type CreateBotRequest struct {
	*tchttp.BaseRequest
	
	// 机器人名称
	BotName *string `json:"BotName,omitnil" name:"BotName"`

	// 机器人中文名称
	BotCnName *string `json:"BotCnName,omitnil" name:"BotCnName"`
}

func (r *CreateBotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotName")
	delete(f, "BotCnName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBotResponseParams struct {
	// 任务ID
	TaskRequestId *string `json:"TaskRequestId,omitnil" name:"TaskRequestId"`

	// 任务信息
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBotResponse struct {
	*tchttp.BaseResponse
	Response *CreateBotResponseParams `json:"Response"`
}

func (r *CreateBotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Group struct {
	// 消息类型参考互联网MIME类型标准，当前仅支持"text/plain"。	
	ContentType *string `json:"ContentType,omitnil" name:"ContentType"`

	// 返回内容以链接形式提供。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 普通文本。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil" name:"Content"`
}

// Predefined struct for user
type ResetRequestParams struct {
	// 机器人标识
	BotId *string `json:"BotId,omitnil" name:"BotId"`

	// 子账户id，每个终端对应一个
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 机器人版本号。BotVersion/BotEnv二选一：二者均填，仅BotVersion有效；二者均不填，默认BotEnv=dev
	BotVersion *string `json:"BotVersion,omitnil" name:"BotVersion"`

	// 机器人环境{dev:测试;release:线上}。BotVersion/BotEnv二选一：二者均填，仅BotVersion有效；二者均不填，默认BotEnv=dev
	BotEnv *string `json:"BotEnv,omitnil" name:"BotEnv"`
}

type ResetRequest struct {
	*tchttp.BaseRequest
	
	// 机器人标识
	BotId *string `json:"BotId,omitnil" name:"BotId"`

	// 子账户id，每个终端对应一个
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 机器人版本号。BotVersion/BotEnv二选一：二者均填，仅BotVersion有效；二者均不填，默认BotEnv=dev
	BotVersion *string `json:"BotVersion,omitnil" name:"BotVersion"`

	// 机器人环境{dev:测试;release:线上}。BotVersion/BotEnv二选一：二者均填，仅BotVersion有效；二者均不填，默认BotEnv=dev
	BotEnv *string `json:"BotEnv,omitnil" name:"BotEnv"`
}

func (r *ResetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotId")
	delete(f, "UserId")
	delete(f, "BotVersion")
	delete(f, "BotEnv")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetResponseParams struct {
	// 当前会话状态。取值:"start"/"continue"/"complete"
	// 注意：此字段可能返回 null，表示取不到有效值。
	DialogStatus *string `json:"DialogStatus,omitnil" name:"DialogStatus"`

	// 匹配到的机器人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotName *string `json:"BotName,omitnil" name:"BotName"`

	// 匹配到的意图名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentName *string `json:"IntentName,omitnil" name:"IntentName"`

	// 机器人回答
	ResponseText *string `json:"ResponseText,omitnil" name:"ResponseText"`

	// 语义解析的槽位结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotInfoList []*SlotInfo `json:"SlotInfoList,omitnil" name:"SlotInfoList"`

	// 透传字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionAttributes *string `json:"SessionAttributes,omitnil" name:"SessionAttributes"`

	// 用户说法。该说法是用户原生说法或ASR识别结果，未经过语义优化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil" name:"Question"`

	// tts合成pcm音频存储链接。仅当请求参数NeedTts=true时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaveUrl *string `json:"WaveUrl,omitnil" name:"WaveUrl"`

	// tts合成的pcm音频。二进制数组经过base64编码(暂时不返回)
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaveData *string `json:"WaveData,omitnil" name:"WaveData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetResponse struct {
	*tchttp.BaseResponse
	Response *ResetResponseParams `json:"Response"`
}

func (r *ResetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResponseMessage struct {
	// 消息组列表。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupList []*Group `json:"GroupList,omitnil" name:"GroupList"`
}

type SlotInfo struct {
	// 槽位名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotName *string `json:"SlotName,omitnil" name:"SlotName"`

	// 槽位值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotValue *string `json:"SlotValue,omitnil" name:"SlotValue"`
}

// Predefined struct for user
type TextProcessRequestParams struct {
	// 机器人标识，用于定义抽象机器人。
	BotId *string `json:"BotId,omitnil" name:"BotId"`

	// 终端标识，每个终端(或线程)对应一个，区分并发多用户。
	TerminalId *string `json:"TerminalId,omitnil" name:"TerminalId"`

	// 请求的文本。
	InputText *string `json:"InputText,omitnil" name:"InputText"`

	// 机器人版本，取值"dev"或"release"，{调试版本：dev；线上版本：release}。
	BotEnv *string `json:"BotEnv,omitnil" name:"BotEnv"`

	// 透传字段，透传给用户自定义的WebService服务。
	SessionAttributes *string `json:"SessionAttributes,omitnil" name:"SessionAttributes"`
}

type TextProcessRequest struct {
	*tchttp.BaseRequest
	
	// 机器人标识，用于定义抽象机器人。
	BotId *string `json:"BotId,omitnil" name:"BotId"`

	// 终端标识，每个终端(或线程)对应一个，区分并发多用户。
	TerminalId *string `json:"TerminalId,omitnil" name:"TerminalId"`

	// 请求的文本。
	InputText *string `json:"InputText,omitnil" name:"InputText"`

	// 机器人版本，取值"dev"或"release"，{调试版本：dev；线上版本：release}。
	BotEnv *string `json:"BotEnv,omitnil" name:"BotEnv"`

	// 透传字段，透传给用户自定义的WebService服务。
	SessionAttributes *string `json:"SessionAttributes,omitnil" name:"SessionAttributes"`
}

func (r *TextProcessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextProcessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotId")
	delete(f, "TerminalId")
	delete(f, "InputText")
	delete(f, "BotEnv")
	delete(f, "SessionAttributes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextProcessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextProcessResponseParams struct {
	// 当前会话状态{会话开始: START; 会话中: COUTINUE; 会话结束: COMPLETE}。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DialogStatus *string `json:"DialogStatus,omitnil" name:"DialogStatus"`

	// 匹配到的机器人名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotName *string `json:"BotName,omitnil" name:"BotName"`

	// 匹配到的意图名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentName *string `json:"IntentName,omitnil" name:"IntentName"`

	// 槽位信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotInfoList []*SlotInfo `json:"SlotInfoList,omitnil" name:"SlotInfoList"`

	// 原始的用户说法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputText *string `json:"InputText,omitnil" name:"InputText"`

	// 透传字段，由用户自定义的WebService服务返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionAttributes *string `json:"SessionAttributes,omitnil" name:"SessionAttributes"`

	// 机器人对话的应答文本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseText *string `json:"ResponseText,omitnil" name:"ResponseText"`

	// 结果类型 {中间逻辑出错:0; 任务型机器人:1; 问答型机器人:2; 闲聊型机器人:3; 未匹配上，返回预设兜底话术:5; 未匹配上，返回相似问题列表:6}。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultType *string `json:"ResultType,omitnil" name:"ResultType"`

	// 机器人应答。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseMessage *ResponseMessage `json:"ResponseMessage,omitnil" name:"ResponseMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TextProcessResponse struct {
	*tchttp.BaseResponse
	Response *TextProcessResponseParams `json:"Response"`
}

func (r *TextProcessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextResetRequestParams struct {
	// 机器人标识，用于定义抽象机器人。
	BotId *string `json:"BotId,omitnil" name:"BotId"`

	// 终端标识，每个终端(或线程)对应一个，区分并发多用户。
	TerminalId *string `json:"TerminalId,omitnil" name:"TerminalId"`

	// 机器人版本，取值"dev"或"release"，{调试版本：dev；线上版本：release}。
	BotEnv *string `json:"BotEnv,omitnil" name:"BotEnv"`
}

type TextResetRequest struct {
	*tchttp.BaseRequest
	
	// 机器人标识，用于定义抽象机器人。
	BotId *string `json:"BotId,omitnil" name:"BotId"`

	// 终端标识，每个终端(或线程)对应一个，区分并发多用户。
	TerminalId *string `json:"TerminalId,omitnil" name:"TerminalId"`

	// 机器人版本，取值"dev"或"release"，{调试版本：dev；线上版本：release}。
	BotEnv *string `json:"BotEnv,omitnil" name:"BotEnv"`
}

func (r *TextResetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextResetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotId")
	delete(f, "TerminalId")
	delete(f, "BotEnv")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextResetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextResetResponseParams struct {
	// 当前会话状态，取值："START"/"COUTINUE"/"COMPLETE"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DialogStatus *string `json:"DialogStatus,omitnil" name:"DialogStatus"`

	// 匹配到的机器人名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotName *string `json:"BotName,omitnil" name:"BotName"`

	// 匹配到的意图名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentName *string `json:"IntentName,omitnil" name:"IntentName"`

	// 槽位信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotInfoList []*SlotInfo `json:"SlotInfoList,omitnil" name:"SlotInfoList"`

	// 原始的用户说法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputText *string `json:"InputText,omitnil" name:"InputText"`

	// 透传字段，由用户自定义的WebService服务返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionAttributes *string `json:"SessionAttributes,omitnil" name:"SessionAttributes"`

	// 机器人对话的应答文本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseText *string `json:"ResponseText,omitnil" name:"ResponseText"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TextResetResponse struct {
	*tchttp.BaseResponse
	Response *TextResetResponseParams `json:"Response"`
}

func (r *TextResetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextResetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}