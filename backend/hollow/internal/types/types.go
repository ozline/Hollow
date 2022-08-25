package tpyes

type User struct {
	ID          int64  //用户ID
	Status      int64  //用户状态
	Username    string //用户名
	Email       string //邮箱
	Phone       int64  //手机号
	Created_at  int64  //创建时间
	Deleted_at  int64  //删除时间
	Updated_at  int64  //更新时间
	Password    string //密码
	Nickname    string //昵称
	Mfa_enabled bool   //是否启用MFA
	Mfa_secret  string //MFA秘钥
}

type Leaf struct {
	ID         int64  //消息ID
	Owner      int64  //所述对象
	Created_at int64  //创建时间
	Status     int64  //消息状态 0=匿名 1=实名 2=封禁
	Message    string //消息
	Liked      int64  //点赞数
}

type Comment struct {
	ID         int64  //评论ID
	Owner      int64  //用户
	Root       int64  //归属帖子id
	Father     int64  //父楼id 0=顶级 其他数字则为上一层id
	Created_at int64  //创建时间
	Status     int64  //状态 0=匿名 1=实名 2=封禁（删除）
	Message    string //评论内容
	Liked      int64  //点赞数
}

type Like struct {
	Timestamp int64 //评论时间(充当点赞ID)
	User      int64 //用户
	Comment   int64 //评论ID
}
