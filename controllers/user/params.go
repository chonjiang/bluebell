package user

// validator内置验证参数与语法：https://blog.csdn.net/zhaozuoyou/article/details/127812519

// RegisterParam 注册接口参数
type RegisterParam struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=120"`
	Gender     int8   `json:"gender"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// LoginParam 注册接口参数
type LoginParam struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
