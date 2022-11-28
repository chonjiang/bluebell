package demo

// validator内置验证参数与语法：https://blog.csdn.net/zhaozuoyou/article/details/127812519

// SignUpParam 注册接口参数
type SignUpParam struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	BirthDate  string `json:"birth" binding:"required,datetime=2006-01-02,checkDate"`
}
