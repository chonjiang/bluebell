package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"time"
)

// 自定义字段校验方法

//type SignUpParam struct {
//	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
//	Name       string `json:"name" binding:"required"`
//	Email      string `json:"email" binding:"required,email"`
//	Password   string `json:"password" binding:"required"`
//	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
//	需要使用自定义校验方法checkDate校验日期字段
//	BirthDate      string `json:"date" binding:"required,datetime=2006-01-02,checkDate"`
//}

// 自定义方法验证传入的日期参数不得大于今天的日期
func checkDate(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}

	if date.After(time.Now()) {
		return false
	}

	return true
}

// 定义好了字段及其自定义校验方法后，就需要将它们联系起来并注册到校验器实例中
func registerDateChecker() error {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("checkDate", checkDate); err != nil {
			return err
		}
	}
	return nil
}

func init() {
	err := registerDateChecker()
	if err != nil {
		panic(err.Error())
	}
}
