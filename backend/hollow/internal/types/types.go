package tpyes

type User struct {
	ID          int64  // 用户ID
	Status      int64  // 用户状态 0=正常 1=管理员 2=封禁
	Username    string // 用户名
	Email       string // 邮箱
	Phone       int64  // 手机号
	Created_at  int64  // 创建时间
	Deleted_at  int64  // 删除时间
	Updated_at  int64  // 更新时间
	Password    string // 密码
	Nickname    string // 昵称
	Mfa_enabled bool   // 是否启用MFA
	Mfa_secret  string // MFA秘钥
}

type Leaf struct {
	ID         int64  // 消息ID
	Owner      int64  // 所述对象
	Created_at int64  // 创建时间
	Status     int64  // 消息状态 0=匿名 1=实名 2=封禁
	Message    string // 消息
	Liked      int64  // 点赞数
}

type Comment struct {
	ID         int64  // 评论ID
	Owner      int64  // 用户
	Root       int64  // 归属帖子id
	Father     int64  // 父楼id 0=顶级 其他数字则为上一层id
	Created_at int64  // 创建时间
	Status     int64  // 状态 0=匿名 1=实名 2=封禁（删除）
	Message    string // 评论内容
	Liked      int64  // 点赞数
}

type Like struct {
	Timestamp int64 // 评论时间(充当点赞ID)
	User      int64 // 用户
	Comment   int64 // 评论ID
}

type ShortMsg struct {
	Code      string // 返回OK代表请求成功。 其他：https://help.aliyun.com/document_detail/101346.html
	Message   string // 状态码的描述。
	BizId     string // 发送回执ID。
	RequestId string // 请求ID。
}

type Report struct {
	Id         int64  // 举报id
	Type       int64  // 举报类型 0=帖子 1=评论 2=用户
	Status     int64  // 当前状态 0=未审核 1=审核通过
	Reporter   int64  // 举报人ID
	Report_id  int64  // 举报对象ID
	Reason     string // 举报描述
	Message    string // 原始文本
	Reply      string // 管理员回复
	Created_at int64  // 创建时间
	Updated_at int64  // 最后更新
}
