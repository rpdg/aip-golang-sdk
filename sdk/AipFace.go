package sdk

import (
	"aip-face-sdk/sdk/lib"
)

const (
	/**
	 * 人脸检测 detect api url
	 * @var string
	 */
	DETECT_URL = "https://aip.baidubce.com/rest/2.0/face/v3/detect"

	/**
	 * 人脸搜索 search api url
	 * @var string
	 */
	SEARCH_URL = "https://aip.baidubce.com/rest/2.0/face/v3/search"

	/**
	 * 人脸搜索 M:N 识别 multi_search api url
	 * @var string
	 */
	MULTI_SEARCH_URL = "https://aip.baidubce.com/rest/2.0/face/v3/multi-search"

	/**
	 * 人脸注册 user_add api url
	 * @var string
	 */
	USER_ADD_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add"

	/**
	 * 人脸更新 user_update api url
	 * @var string
	 */
	USER_UPDATE_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/update"

	/**
	 * 人脸删除 face_delete api url
	 * @var string
	 */
	FACE_DELETE_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/face/delete"

	/**
	 * 用户信息查询 user_get api url
	 * @var string
	 */
	USER_GET_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/get"

	/**
	 * 获取用户人脸列表 face_getlist api url
	 * @var string
	 */
	FACE_GET_LIST_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/face/getlist"

	/**
	 * 获取用户列表 group_getusers api url
	 * @var string
	 */
	GROUP_GET_USER_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/getusers"

	/**
	 * 复制用户 user_copy api url
	 * @var string
	 */
	USER_COPY_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/copy"

	/**
	 * 删除用户 user_delete api url
	 * @var string
	 */
	USER_DELETE_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/delete"

	/**
	 * 创建用户组 group_add api url
	 * @var string
	 */
	GROUP_ADD_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/add"

	/**
	 * 删除用户组 group_delete api url
	 * @var string
	 */
	GROUP_DELETE_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/delete"

	/**
	 * 组列表查询 group_getlist api url
	 * @var string
	 */
	GROUP_GET_LIST_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/getlist"

	/**
	 * 身份验证 person_verify api url
	 * @var string
	 */
	PERSON_VERIFY_URL = "https://aip.baidubce.com/rest/2.0/face/v3/person/verify"

	/**
	 * 语音校验码接口 video_sessioncode api url
	 * @var string
	 */
	VIDEO_SESSIONCODE_URL = "https://aip.baidubce.com/rest/2.0/face/v1/faceliveness/sessioncode"
)

const (
	/**
	 * 在线活体检测 faceverify api url
	 * @var string
	 */
	FACE_VERIFY_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceverify"

	/**
	 * 人脸比对 match api url
	 * @var string
	 */
	MATCH_URL = "https://aip.baidubce.com/rest/2.0/face/v3/match"
)

type result map[string]interface{}

type AipFace struct {
	lib.AipBase
}

func NewAipFace() *AipFace {
	return &AipFace{}
}

/**
 * 人脸检测接口
 *
 * @param string $image - 图片信息(**总数据大小应小于10M**)，图片上传方式根据image_type来判断
 * @param string $imageType - 图片类型     <br> **BASE64**:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M； <br>**URL**:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)； <br>**FACE_TOKEN**: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   face_field 包括**age,beauty,expression,face_shape,gender,glasses,landmark,landmark72，landmark150，race,quality,eye_status,emotion,face_type信息**  <br> 逗号分隔. 默认只返回face_token、人脸框、概率和旋转角度
 *   max_face_num 最多处理人脸的数目，默认值为1，仅检测图片中面积最大的那个人脸；**最大值10**，检测图片中面积最大的几张人脸。
 *   face_type 人脸的类型 **LIVE**表示生活照：通常为手机、相机拍摄的人像图片、或从网络获取的人像图片等**IDCARD**表示身份证芯片照：二代身份证内置芯片中的人像照片 **WATERMARK**表示带水印证件照：一般为带水印的小图，如公安网小图 **CERT**表示证件照片：如拍摄的身份证、工卡、护照、学生证等证件图片 默认**LIVE**
 *   liveness_control 活体检测控制  **NONE**: 不进行控制 **LOW**:较低的活体要求(高通过率 低攻击拒绝率) **NORMAL**: 一般的活体要求(平衡的攻击拒绝率, 通过率) **HIGH**: 较高的活体要求(高攻击拒绝率 低通过率) **默认NONE**
 * @return array
 */
func (face *AipFace) Detect(image string, imageType string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["image"] = image
	data["image_type"] = imageType

	for k, v := range options {
		data[k] = v
	}

	return face.Request(DETECT_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 人脸搜索接口
 *
 * @param string $image - 图片信息(**总数据大小应小于10M**)，图片上传方式根据image_type来判断
 * @param string $imageType - 图片类型     <br> **BASE64**:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M； <br>**URL**:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)； <br>**FACE_TOKEN**: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
 * @param string $groupIdList - 从指定的group中进行查找 用逗号分隔，**上限20个**
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   max_face_num 最多处理人脸的数目<br>**默认值为1(仅检测图片中面积最大的那个人脸)** **最大值10**
 *   match_threshold 匹配阈值（设置阈值后，score低于此阈值的用户信息将不会返回） 最大100 最小0 默认80 <br>**此阈值设置得越高，检索速度将会越快，推荐使用默认阈值`80`**
 *   quality_control 图片质量控制  **NONE**: 不进行控制 **LOW**:较低的质量要求 **NORMAL**: 一般的质量要求 **HIGH**: 较高的质量要求 **默认 NONE**
 *   liveness_control 活体检测控制  **NONE**: 不进行控制 **LOW**:较低的活体要求(高通过率 低攻击拒绝率) **NORMAL**: 一般的活体要求(平衡的攻击拒绝率, 通过率) **HIGH**: 较高的活体要求(高攻击拒绝率 低通过率) **默认NONE**
 *   user_id 当需要对特定用户进行比对时，指定user_id进行比对。即人脸认证功能。
 *   max_user_num 查找后返回的用户数量。返回相似度最高的几个用户，默认为1，最多返回50个。
 * @return array
 */
func (face *AipFace) Search(image string, imageType string, groupIdList string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["image"] = image
	data["image_type"] = imageType
	data["group_id_list"] = groupIdList

	for k, v := range options {
		data[k] = v
	}

	return face.Request(SEARCH_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 人脸搜索 M:N 识别接口
 *
 * @param string $image - 图片信息(**总数据大小应小于10M**)，图片上传方式根据image_type来判断
 * @param string $imageType - 图片类型     <br> **BASE64**:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M； <br>**URL**:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)； <br>**FACE_TOKEN**: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
 * @param string $groupIdList - 从指定的group中进行查找 用逗号分隔，**上限20个**
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   max_face_num 最多处理人脸的数目<br>**默认值为1(仅检测图片中面积最大的那个人脸)** **最大值10**
 *   match_threshold 匹配阈值（设置阈值后，score低于此阈值的用户信息将不会返回） 最大100 最小0 默认80 <br>**此阈值设置得越高，检索速度将会越快，推荐使用默认阈值`80`**
 *   quality_control 图片质量控制  **NONE**: 不进行控制 **LOW**:较低的质量要求 **NORMAL**: 一般的质量要求 **HIGH**: 较高的质量要求 **默认 NONE**
 *   liveness_control 活体检测控制  **NONE**: 不进行控制 **LOW**:较低的活体要求(高通过率 低攻击拒绝率) **NORMAL**: 一般的活体要求(平衡的攻击拒绝率, 通过率) **HIGH**: 较高的活体要求(高攻击拒绝率 低通过率) **默认NONE**
 *   max_user_num 查找后返回的用户数量。返回相似度最高的几个用户，默认为1，最多返回50个。
 * @return array
 */
func (face *AipFace) MultiSearch(image string, imageType string, groupIdList string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["image"] = image
	data["image_type"] = imageType
	data["group_id_list"] = groupIdList

	for k, v := range options {
		data[k] = v
	}

	return face.Request(MULTI_SEARCH_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 人脸注册接口
 *
 * @param string $image - 图片信息(总数据大小应小于10M)，图片上传方式根据image_type来判断。注：组内每个uid下的人脸图片数目上限为20张
 * @param string $imageType - 图片类型     <br> **BASE64**:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M； <br>**URL**:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)； <br>**FACE_TOKEN**: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
 * @param string $groupId - 用户组id（由数字、字母、下划线组成），长度限制128B
 * @param string $userId - 用户id（由数字、字母、下划线组成），长度限制128B
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   user_info 用户资料，长度限制256B
 *   quality_control 图片质量控制  **NONE**: 不进行控制 **LOW**:较低的质量要求 **NORMAL**: 一般的质量要求 **HIGH**: 较高的质量要求 **默认 NONE**
 *   liveness_control 活体检测控制  **NONE**: 不进行控制 **LOW**:较低的活体要求(高通过率 低攻击拒绝率) **NORMAL**: 一般的活体要求(平衡的攻击拒绝率, 通过率) **HIGH**: 较高的活体要求(高攻击拒绝率 低通过率) **默认NONE**
 *   action_type 操作方式  APPEND: 当user_id在库中已经存在时，对此user_id重复注册时，新注册的图片默认会追加到该user_id下,REPLACE : 当对此user_id重复注册时,则会用新图替换库中该user_id下所有图片,默认使用APPEND
 * @return array
 */
func (face *AipFace) AddUser(image string, imageType string, groupId string, userId string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["image"] = image
	data["image_type"] = imageType
	data["group_id"] = groupId
	data["user_id"] = userId

	for k, v := range options {
		data[k] = v
	}

	return face.Request(MULTI_SEARCH_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 人脸更新接口
 *
 * @param string $image - 图片信息(**总数据大小应小于10M**)，图片上传方式根据image_type来判断
 * @param string $imageType - 图片类型     <br> **BASE64**:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M； <br>**URL**:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)； <br>**FACE_TOKEN**: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
 * @param string $groupId - 更新指定groupid下uid对应的信息
 * @param string $userId - 用户id（由数字、字母、下划线组成），长度限制128B
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   user_info 用户资料，长度限制256B
 *   quality_control 图片质量控制  **NONE**: 不进行控制 **LOW**:较低的质量要求 **NORMAL**: 一般的质量要求 **HIGH**: 较高的质量要求 **默认 NONE**
 *   liveness_control 活体检测控制  **NONE**: 不进行控制 **LOW**:较低的活体要求(高通过率 低攻击拒绝率) **NORMAL**: 一般的活体要求(平衡的攻击拒绝率, 通过率) **HIGH**: 较高的活体要求(高攻击拒绝率 低通过率) **默认NONE**
 *   action_type 操作方式  APPEND: 当user_id在库中已经存在时，对此user_id重复注册时，新注册的图片默认会追加到该user_id下,REPLACE : 当对此user_id重复注册时,则会用新图替换库中该user_id下所有图片,默认使用APPEND
 * @return array
 */
func (face *AipFace) UpdateUser(image string, imageType string, groupId string, userId string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["image"] = image
	data["image_type"] = imageType
	data["group_id"] = groupId
	data["user_id"] = userId

	for k, v := range options {
		data[k] = v
	}

	return face.Request(USER_UPDATE_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 人脸删除接口
 *
 * @param string $userId - 用户id（由数字、字母、下划线组成），长度限制128B
 * @param string $groupId - 用户组id（由数字、字母、下划线组成），长度限制128B
 * @param string $faceToken - 需要删除的人脸图片token，（由数字、字母、下划线组成）长度限制64B
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 * @return array
 */
func (face *AipFace) FaceDelete(userId string, groupId string, faceToken string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["user_id"] = userId
	data["group_id"] = groupId
	data["face_token"] = faceToken

	for k, v := range options {
		data[k] = v
	}

	return face.Request(FACE_DELETE_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 用户信息查询接口
 *
 * @param string $userId - 用户id（由数字、字母、下划线组成），长度限制128B
 * @param string $groupId - 用户组id（由数字、字母、下划线组成），长度限制128B
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 * @return array
 */
func (face *AipFace) GetUser(userId string, groupId string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["user_id"] = userId
	data["group_id"] = groupId

	for k, v := range options {
		data[k] = v
	}

	return face.Request(USER_GET_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 获取用户人脸列表接口
 *
 * @param string $userId - 用户id（由数字、字母、下划线组成），长度限制128B
 * @param string $groupId - 用户组id（由数字、字母、下划线组成），长度限制128B
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 * @return array
 */
func (face *AipFace) FaceGetList(userId string, groupId string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["user_id"] = userId
	data["group_id"] = groupId

	for k, v := range options {
		data[k] = v
	}

	return face.Request(FACE_GET_LIST_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 获取用户列表接口
 *
 * @param string $groupId - 用户组id（由数字、字母、下划线组成），长度限制128B
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   start 默认值0，起始序号
 *   length 返回数量，默认值100，最大值1000
 * @return array
 */
func (face *AipFace) GetGroupUsers(groupId string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["group_id"] = groupId

	for k, v := range options {
		data[k] = v
	}

	return face.Request(GROUP_GET_USER_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 复制用户接口
 *
 * @param string $userId - 用户id（由数字、字母、下划线组成），长度限制128B
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   src_group_id 从指定组里复制信息
 *   dst_group_id 需要添加用户的组id
 * @return array
 */
func (face *AipFace) UserCopy(userId string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["user_id"] = userId

	for k, v := range options {
		data[k] = v
	}

	return face.Request(USER_COPY_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 删除用户接口
 *
 * @param string $groupId - 用户组id（由数字、字母、下划线组成），长度限制128B
 * @param string $userId - 用户id（由数字、字母、下划线组成），长度限制128B
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 * @return array
 */
func (face *AipFace) DeleteUser(groupId string, userId string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["group_id"] = groupId
	data["user_id"] = userId

	for k, v := range options {
		data[k] = v
	}

	return face.Request(USER_DELETE_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 创建用户组接口
 *
 * @param string $groupId - 用户组id（由数字、字母、下划线组成），长度限制128B
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 * @return array
 */
func (face *AipFace) GroupAdd(groupId string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["group_id"] = groupId

	for k, v := range options {
		data[k] = v
	}

	return face.Request(GROUP_ADD_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 删除用户组接口
 *
 * @param string $groupId - 用户组id（由数字、字母、下划线组成），长度限制128B
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 * @return array
 */
func (face *AipFace) GroupDelete(groupId string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["group_id"] = groupId

	for k, v := range options {
		data[k] = v
	}

	return face.Request(GROUP_DELETE_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 组列表查询接口
 *
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   start 默认值0，起始序号
 *   length 返回数量，默认值100，最大值1000
 * @return array
 */
func (face *AipFace) GetGroupList(options map[string]interface{}) map[string]interface{} {
	data := make(result)

	for k, v := range options {
		data[k] = v
	}

	return face.Request(GROUP_GET_LIST_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 身份验证接口
 *
 * @param string $image - 图片信息(**总数据大小应小于10M**)，图片上传方式根据image_type来判断
 * @param string $imageType - 图片类型     <br> **BASE64**:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M； <br>**URL**:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)； <br>**FACE_TOKEN**: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
 * @param string $idCardNumber - 身份证号（真实身份证号号码）
 * @param string $name - utf8，姓名（真实姓名，和身份证号匹配）
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   quality_control 图片质量控制  **NONE**: 不进行控制 **LOW**:较低的质量要求 **NORMAL**: 一般的质量要求 **HIGH**: 较高的质量要求 **默认 NONE**
 *   liveness_control 活体检测控制  **NONE**: 不进行控制 **LOW**:较低的活体要求(高通过率 低攻击拒绝率) **NORMAL**: 一般的活体要求(平衡的攻击拒绝率, 通过率) **HIGH**: 较高的活体要求(高攻击拒绝率 低通过率) **默认NONE**
 * @return array
 */
func (face *AipFace) PersonVerify(image string, imageType string, idCardNumber string, name string, options map[string]interface{}) map[string]interface{} {
	data := make(result)

	data["image"] = image
	data["image_type"] = imageType
	data["id_card_number"] = idCardNumber
	data["name"] = name

	for k, v := range options {
		data[k] = v
	}

	return face.Request(PERSON_VERIFY_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 语音校验码接口接口
 *
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   appid 百度云创建应用时的唯一标识ID
 * @return array
 */
func (face *AipFace) VideoSessionCode(options map[string]interface{}) map[string]interface{} {
	data := make(result)

	for k, v := range options {
		data[k] = v
	}

	return face.Request(VIDEO_SESSIONCODE_URL, data, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 在线活体检测接口
 *
 * @param array $images
 * @return array
 */
func (face *AipFace) FaceVerify(images []result) map[string]interface{} {
	return face.Request(FACE_VERIFY_URL, images, map[string]string{
		"Content-Type": "application/json",
	})
}

/**
 * 人脸比对接口
 *
 * @param array $images
 * @return array
 */
func (face *AipFace) Match(images []result) map[string]interface{} {
	return face.Request(MATCH_URL, images, map[string]string{
		"Content-Type": "application/json",
	})
}
