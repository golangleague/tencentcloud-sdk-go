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

package v20181201

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeMaterialListRequestParams struct {
	// 活动Id
	ActivityId *int64 `json:"ActivityId,omitnil" name:"ActivityId"`

	// 素材Id
	MaterialId *string `json:"MaterialId,omitnil" name:"MaterialId"`

	// 每次拉取条数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeMaterialListRequest struct {
	*tchttp.BaseRequest
	
	// 活动Id
	ActivityId *int64 `json:"ActivityId,omitnil" name:"ActivityId"`

	// 素材Id
	MaterialId *string `json:"MaterialId,omitnil" name:"MaterialId"`

	// 每次拉取条数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeMaterialListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityId")
	delete(f, "MaterialId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaterialListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaterialListResponseParams struct {
	// 素材列表数据
	MaterialInfos []*PublicMaterialInfos `json:"MaterialInfos,omitnil" name:"MaterialInfos"`

	// 素材条数
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMaterialListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaterialListResponseParams `json:"Response"`
}

func (r *DescribeMaterialListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceFusionLiteRequestParams struct {
	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitnil" name:"ModelId"`

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil" name:"MergeInfos"`

	// 返回图像方式（url 或 base64) ，二选一。默认url, url有效期为30天。
	RspImgType *string `json:"RspImgType,omitnil" name:"RspImgType"`

	// 请注意，鉴政服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitnil" name:"CelebrityIdentify"`

	// 算法引擎参数:  1）选脸版 - youturecreat; 2）优享版 - youtu1vN； 3）畅享版 - ptu； 4）随机 - ALL;  默认为活动选择的算法
	Engine *string `json:"Engine,omitnil" name:"Engine"`
}

type FaceFusionLiteRequest struct {
	*tchttp.BaseRequest
	
	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitnil" name:"ModelId"`

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil" name:"MergeInfos"`

	// 返回图像方式（url 或 base64) ，二选一。默认url, url有效期为30天。
	RspImgType *string `json:"RspImgType,omitnil" name:"RspImgType"`

	// 请注意，鉴政服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitnil" name:"CelebrityIdentify"`

	// 算法引擎参数:  1）选脸版 - youturecreat; 2）优享版 - youtu1vN； 3）畅享版 - ptu； 4）随机 - ALL;  默认为活动选择的算法
	Engine *string `json:"Engine,omitnil" name:"Engine"`
}

func (r *FaceFusionLiteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceFusionLiteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ModelId")
	delete(f, "MergeInfos")
	delete(f, "RspImgType")
	delete(f, "CelebrityIdentify")
	delete(f, "Engine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FaceFusionLiteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceFusionLiteResponseParams struct {
	// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 鉴政结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReviewResultSet []*FuseFaceReviewResult `json:"ReviewResultSet,omitnil" name:"ReviewResultSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type FaceFusionLiteResponse struct {
	*tchttp.BaseResponse
	Response *FaceFusionLiteResponseParams `json:"Response"`
}

func (r *FaceFusionLiteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceFusionLiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceFusionRequestParams struct {
	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitnil" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitnil" name:"RspImgType"`

	// 图片 base64 数据。请确保人脸为正脸，无旋转。若某些手机拍摄后人脸被旋转，请使用图片的 EXIF 信息对图片进行旋转处理；请勿在 base64 数据中包含头部，如“data:image/jpeg;base64,”。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 历史遗留字段，无需填写。因为融合只需提取人脸特征，不需要鉴黄。
	PornDetect *int64 `json:"PornDetect,omitnil" name:"PornDetect"`

	// 0表示不需要不适宜内容识别，1表示需要不适宜内容识别。默认值为0。
	// 请注意，不适宜内容识别服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitnil" name:"CelebrityIdentify"`

	// 图片Url地址
	Url *string `json:"Url,omitnil" name:"Url"`

	// 融合参数
	FuseParam *FuseParam `json:"FuseParam,omitnil" name:"FuseParam"`
}

type FaceFusionRequest struct {
	*tchttp.BaseRequest
	
	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitnil" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitnil" name:"RspImgType"`

	// 图片 base64 数据。请确保人脸为正脸，无旋转。若某些手机拍摄后人脸被旋转，请使用图片的 EXIF 信息对图片进行旋转处理；请勿在 base64 数据中包含头部，如“data:image/jpeg;base64,”。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 历史遗留字段，无需填写。因为融合只需提取人脸特征，不需要鉴黄。
	PornDetect *int64 `json:"PornDetect,omitnil" name:"PornDetect"`

	// 0表示不需要不适宜内容识别，1表示需要不适宜内容识别。默认值为0。
	// 请注意，不适宜内容识别服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitnil" name:"CelebrityIdentify"`

	// 图片Url地址
	Url *string `json:"Url,omitnil" name:"Url"`

	// 融合参数
	FuseParam *FuseParam `json:"FuseParam,omitnil" name:"FuseParam"`
}

func (r *FaceFusionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceFusionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ModelId")
	delete(f, "RspImgType")
	delete(f, "Image")
	delete(f, "PornDetect")
	delete(f, "CelebrityIdentify")
	delete(f, "Url")
	delete(f, "FuseParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FaceFusionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceFusionResponseParams struct {
	// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 不适宜内容识别结果
	ReviewResultSet []*FuseFaceReviewResult `json:"ReviewResultSet,omitnil" name:"ReviewResultSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type FaceFusionResponse struct {
	*tchttp.BaseResponse
	Response *FaceFusionResponseParams `json:"Response"`
}

func (r *FaceFusionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceFusionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceInfo struct {
	// 人脸框的横坐标
	X *int64 `json:"X,omitnil" name:"X"`

	// 人脸框的纵坐标
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// 人脸框的宽度
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 人脸框的高度
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

type FaceRect struct {
	// 人脸框左上角横坐标。
	X *int64 `json:"X,omitnil" name:"X"`

	// 人脸框左上角纵坐标。
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 人脸框高度。
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

// Predefined struct for user
type FuseFaceRequestParams struct {
	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitnil" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitnil" name:"RspImgType"`

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil" name:"MergeInfos"`

	// 脸型融合比例，数值越高，融合后的脸型越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中脸型参数数值。（换脸版算法暂不支持此参数调整）
	FuseProfileDegree *int64 `json:"FuseProfileDegree,omitnil" name:"FuseProfileDegree"`

	// 五官融合比例，数值越高，融合后的五官越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中五官参数数值。（换脸版算法暂不支持此参数调整）
	FuseFaceDegree *int64 `json:"FuseFaceDegree,omitnil" name:"FuseFaceDegree"`

	// 0表示不需要不适宜内容识别，1表示需要不适宜内容识别。默认值为0。
	// 请注意，不适宜内容识别服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitnil" name:"CelebrityIdentify"`

	// 融合参数
	FuseParam *FuseParam `json:"FuseParam,omitnil" name:"FuseParam"`
}

type FuseFaceRequest struct {
	*tchttp.BaseRequest
	
	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitnil" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitnil" name:"RspImgType"`

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil" name:"MergeInfos"`

	// 脸型融合比例，数值越高，融合后的脸型越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中脸型参数数值。（换脸版算法暂不支持此参数调整）
	FuseProfileDegree *int64 `json:"FuseProfileDegree,omitnil" name:"FuseProfileDegree"`

	// 五官融合比例，数值越高，融合后的五官越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中五官参数数值。（换脸版算法暂不支持此参数调整）
	FuseFaceDegree *int64 `json:"FuseFaceDegree,omitnil" name:"FuseFaceDegree"`

	// 0表示不需要不适宜内容识别，1表示需要不适宜内容识别。默认值为0。
	// 请注意，不适宜内容识别服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitnil" name:"CelebrityIdentify"`

	// 融合参数
	FuseParam *FuseParam `json:"FuseParam,omitnil" name:"FuseParam"`
}

func (r *FuseFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ModelId")
	delete(f, "RspImgType")
	delete(f, "MergeInfos")
	delete(f, "FuseProfileDegree")
	delete(f, "FuseFaceDegree")
	delete(f, "CelebrityIdentify")
	delete(f, "FuseParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FuseFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FuseFaceResponseParams struct {
	// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。
	FusedImage *string `json:"FusedImage,omitnil" name:"FusedImage"`

	// 不适宜内容识别结果。该数组的顺序和请求中mergeinfo的顺序一致，一一对应
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReviewResultSet []*FuseFaceReviewResult `json:"ReviewResultSet,omitnil" name:"ReviewResultSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type FuseFaceResponse struct {
	*tchttp.BaseResponse
	Response *FuseFaceResponseParams `json:"Response"`
}

func (r *FuseFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FuseFaceReviewDetail struct {
	// 保留字段
	Field *string `json:"Field,omitnil" name:"Field"`

	// 人员名称
	Label *string `json:"Label,omitnil" name:"Label"`

	// 对应识别label的置信度，分数越高意味违法违规可能性越大。 
	// 0到70，Suggestion建议为PASS； 
	// 70到80，Suggestion建议为REVIEW； 
	// 80到100，Suggestion建议为BLOCK。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 识别场景的审核结论：  
	// PASS：正常 
	// REVIEW：疑似  
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`
}

type FuseFaceReviewResult struct {
	// 保留字段
	Category *string `json:"Category,omitnil" name:"Category"`

	// 状态码， 0为处理成功，其他值为处理失败
	Code *string `json:"Code,omitnil" name:"Code"`

	// 对应状态码信息描述
	CodeDescription *string `json:"CodeDescription,omitnil" name:"CodeDescription"`

	// 保留字段
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 保留字段
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 审核详细内容
	DetailSet []*FuseFaceReviewDetail `json:"DetailSet,omitnil" name:"DetailSet"`
}

type FuseParam struct {
	// 图片编码参数
	ImageCodecParam *ImageCodecParam `json:"ImageCodecParam,omitnil" name:"ImageCodecParam"`
}

type ImageCodecParam struct {
	// 元数据
	MetaData []*MetaData `json:"MetaData,omitnil" name:"MetaData"`
}

type MaterialFaceList struct {
	// 人脸序号
	FaceId *string `json:"FaceId,omitnil" name:"FaceId"`

	// 人脸框信息
	FaceInfo *FaceInfo `json:"FaceInfo,omitnil" name:"FaceInfo"`
}

type MergeInfo struct {
	// 输入图片base64
	Image *string `json:"Image,omitnil" name:"Image"`

	// 输入图片url
	Url *string `json:"Url,omitnil" name:"Url"`

	// 上传的图片人脸位置信息（人脸框）
	InputImageFaceRect *FaceRect `json:"InputImageFaceRect,omitnil" name:"InputImageFaceRect"`

	// 控制台上传的素材人脸ID，不填默认取最大人脸
	TemplateFaceID *string `json:"TemplateFaceID,omitnil" name:"TemplateFaceID"`
}

type MetaData struct {
	// MetaData的Key
	MetaKey *string `json:"MetaKey,omitnil" name:"MetaKey"`

	// MetaData的Value
	MetaValue *string `json:"MetaValue,omitnil" name:"MetaValue"`
}

type PublicMaterialInfos struct {
	// 素材Id
	MaterialId *string `json:"MaterialId,omitnil" name:"MaterialId"`

	// 素材状态
	MaterialStatus *int64 `json:"MaterialStatus,omitnil" name:"MaterialStatus"`

	// 脸型参数P图
	BlendParamPtu *int64 `json:"BlendParamPtu,omitnil" name:"BlendParamPtu"`

	// 五官参数P图
	PositionParamPtu *int64 `json:"PositionParamPtu,omitnil" name:"PositionParamPtu"`

	// 脸型参数优图
	BlendParamYoutu *int64 `json:"BlendParamYoutu,omitnil" name:"BlendParamYoutu"`

	// 五官参数优图
	PositionParamYoutu *int64 `json:"PositionParamYoutu,omitnil" name:"PositionParamYoutu"`

	// 素材COS地址
	Url *string `json:"Url,omitnil" name:"Url"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 人脸信息
	MaterialFaceList []*MaterialFaceList `json:"MaterialFaceList,omitnil" name:"MaterialFaceList"`
}