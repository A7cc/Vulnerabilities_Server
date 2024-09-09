package define

// 接收登录参数的结构体
type LoginPassWordRequest struct {
	// 登录用户名
	UserName string `json:"username"`
	// 登录密码
	Password string `json:"password"`
	// 登录验证码
	Code string `json:"code"`
}

// 登录后返回数据
type LoginPasswordResponse struct {
	Uid           uint   `json:"uid"`
	Authorization string `json:"Authorization"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Phone         string `json:"phone"`
	Sex           string `json:"sex"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	RoleLevel     uint   `json:"rolelevel"`
	Introduce     string `json:"introduce"`
	Created_at    string `json:"created_at"`
}

// 用于登录成功后返回token结构体类型
type LoginPasswordReply struct {
	Authorization string `json:"Authorization"`
}

// 获取管理员列表参数结构体
type GetUserListRequest struct {
	*QueryRequest
}

// 关键字和分页信息结构体
type QueryRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	Keyword string `json:"keyword" form:"keyword"`
	Status  int    `json:"status" form:"status"`
}

// 返回管理员信息结构体
type GetUserListReply struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Avatar     string `json:"avatar"`
	Phone      string `json:"phone"`
	Sex        string `json:"sex"`
	Email      string `json:"email"`
	Status     bool   `json:"status"`
	Role_id    uint   `json:"role_id"`
	Role       string `json:"role"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

// 接收添加管理员表单数据结构体
type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Status   bool   `json:"status"`
	Role_id  uint   `json:"role_id"`
	Email    string `json:"email"`
	Remarks  string `json:"remarks"`
}

// 获取管理员信息结构体
type GetUserDetailReply struct {
	ID uint `json:"id"`
	AddUserRequest
}

// 接收更新管理员信息参数结构体
type UpdateUserRequest struct {
	ID uint `json:"id"`
	AddUserRequest
}

//获取角色列表查询参数结构体
type GetRoleListRequest struct {
	*QueryRequest
}

// 返回角色信息结构体
type GetRoleListReply struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Level      string `json:"level"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

// 接收添加角色表单数据结构体
type AddRoleRequest struct {
	Name    string `json:"name"`
	Level   uint   `json:"level"`
	Remarks string `json:"remarks"`
}

// 获取角色信息结构体
type GetRoleDetailReply struct {
	ID uint `json:"id"`
	AddRoleRequest
}

// 接收更新角色信息参数结构体
type UpdateRoleRequest struct {
	ID uint `json:"id"`
	AddRoleRequest
}

// 修改密码
type UpdateInfoType struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Sex      string `json:"sex"`
	Avatar   string `json:"avatar"`
}

// ==========食物信息===========
// 获取食物列表参数结构体
type GetFoodListRequest struct {
	*QueryRequest
}

// 返回食物信息结构体
type GetFoodListReply struct {
	ID            uint    `json:"id"`
	Foodname      string  `json:"foodname"`
	User_id       uint    `json:"userid"`
	User          string  `json:"user"`
	Foodicon      string  `json:"foodicon"`
	Foodprocedure string  `json:"foodprocedure"`
	Video         string  `json:"video"`
	Remarks       string  `json:"remarks"`
	Created_at    string  `json:"created_at"`
	Updated_at    string  `json:"updated_at"`
	Price         float64 `json:"price"`
}

// 接收添加食物表单数据结构体
type AddFoodRequest struct {
	Foodname      string  `json:"foodname"`
	User_id       uint    `json:"user_id"`
	FoodIcon      string  `json:"foodicon"`
	FoodProcedure string  `json:"foodprocedure"`
	Video         string  `json:"video"`
	Remarks       string  `json:"remarks"`
	Price         float64 `json:"price"`
}

// 获取食物信息结构体
type GetFoodDetailReply struct {
	ID uint `json:"id"`
	AddFoodRequest
}

// 接收更新食物信息参数结构体
type UpdateFoodRequest struct {
	ID uint `json:"id"`
	AddFoodRequest
}

// ==========订单信息===========
// 获取订单列表参数结构体
type GetOrderListRequest struct {
	*QueryRequest
}

// 返回订单信息结构体
type GetOrderListReply struct {
	ID         uint   `json:"id"`
	User       string `json:"user"`
	Food_id    string `json:"foodid"`
	Food       string `json:"food"`
	Num        string `json:"num"`
	Remarks    string `json:"remarks"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

// 接收添加订单表单数据结构体
type AddOrderRequest struct {
	User    string `json:"user"`
	Food    string `json:"food"`
	Num     uint8  `json:"num"`
	Remarks string `json:"remarks"`
}

// 获取订单信息结构体
type GetOrderDetailReply struct {
	ID uint `json:"id"`
	AddOrderRequest
}
