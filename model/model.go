package model

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  uint8
}

func UserCreate() {
	//migrator := GLOBAL_DB.Migrator()

	//if !migrator.HasTable(&User{}) {
	//migrator.CreateTable(&User{})
	//}
	// 添加数据 只创建name字段 不创建age字段
	//if err := GLOBAL_DB.Select("Name").Create(&User{Name: "coder4j", Age: 18}).Error; err != nil {
	//	fmt.Println("添加数据失败")
	//}
	//// 除了Name字段 其他字段都创建
	//if err := GLOBAL_DB.Omit("Name").Create(&User{Name: "coder4j", Age: 18}).Error; err != nil {
	//	fmt.Println("添加数据失败")
	//}
	// 使用切片添加多条数据
	GLOBAL_DB.Create(&[]User{
		{Name: "coder1j", Age: 11},
		{Name: "coder2j", Age: 12},
		{Name: "coder3j", Age: 13},
		{Name: "coder4j", Age: 14},
	})

}

func UserQuery() {
	// works because model is specified using `db.Model()`
	var res map[string]interface{}
	GLOBAL_DB.Model(&User{}).First(&res)
	fmt.Println(res)
	var user User
	var userMap map[string]interface{} = make(map[string]interface{})

	// works because destination struct is passed in
	GLOBAL_DB.First(&user)
	fmt.Println(user)

	GLOBAL_DB.Model(&User{}).Take(&res)
	fmt.Println(res)

	GLOBAL_DB.Model(&User{}).Last(&res)
	fmt.Println(res)
	//获取第一条记录（主键升序）
	//where 也可支持map或结构体查询
	query := GLOBAL_DB.Where("name = ? and age = ? ", "coder4j", 21).First(&user)

	fmt.Println(errors.Is(query.Error, gorm.ErrRecordNotFound))
	fmt.Println(userMap)
	fmt.Println(user)

}

type Info struct {
	gorm.Model
	Money int
	DogId uint
}
type Dog struct {
	gorm.Model
	Name     string
	GirlGods []GirlGod `gorm:"many2many:dog_girl_god;"`
	Info     Info
}
type GirlGod struct {
	gorm.Model
	Name string
	Dogs []Dog `gorm:"many2many:dog_girl_god;"`
}

func Orm() {
	//d := Dog{
	//	Name: "我是舔狗",
	//}
	//g := God{
	//	Name: "啊啊啊",
	//	Dog:  d,
	//}
	//GLOBAL_DB.AutoMigrate(&God{}, &Dog{})
	//GLOBAL_DB.Create(&g)
	// 查询
	// 预加载查询
	var gril GirlGod
	GLOBAL_DB.Preload("Dog").First(&gril, 2)
	marshal, err := sonic.Marshal(gril)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", marshal)

}

func OrmOne2More() {
	//GLOBAL_DB.AutoMigrate(&Info{})
	//d1 := Dog{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name: "我是舔狗1号",
	//}
	//d2 := Dog{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//	Name: "我是舔狗2号",
	//}
	//g := God{
	//	Name: "我是女神",
	//	Dog:  []Dog{d1, d2},
	//}
	//GLOBAL_DB.Create(&g)
	// 查询
	var gril GirlGod
	// 查询条件只适用于当前预加载的那一层model
	//GLOBAL_DB.Preload("Dog.Info").First(&gril, 1)

	// 提前连表join
	GLOBAL_DB.Preload("Dog", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("money > ?", 200)
	}).First(&gril, 1)

	fmt.Println(gril)
}

func OrmM2M() {
	//GLOBAL_DB.AutoMigrate(&Dog{}, &GirlGod{}, &Info{})
	//i1 := Info{
	//	Money: 100,
	//}
	//i2 := Info{
	//	Money: 200000,
	//}
	//g1 := GirlGod{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name: "我是女神1号",
	//}
	//g2 := GirlGod{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//	Name: "我是女神2号",
	//}
	//d1 := Dog{
	//	Name: "我是舔狗1号",
	//	GirlGods: []GirlGod{
	//		g2,
	//	},
	//	Info: i1,
	//}
	//d2 := Dog{
	//	Name: "我是舔狗2号",
	//	GirlGods: []GirlGod{
	//		g1, g2,
	//	},
	//	Info: i2,
	//}
	//GLOBAL_DB.Create(&[]Dog{d1, d2})
	d := Dog{
		Model: gorm.Model{
			ID: 4,
		},
	}

	//GLOBAL_DB.Preload("GirlGods").Find(&d)
	// 查询舔狗4号的女神
	var grils []GirlGod
	//GLOBAL_DB.Model(&d).Preload("Dogs.Info").Association("GirlGods").Find(&grils)
	// 只查询钱数大于100的舔狗
	GLOBAL_DB.Model(&d).Preload("Dogs", queryMoney).Association("GirlGods").Find(&grils)

	fmt.Println(grils)
}

func queryMoney(db *gorm.DB) *gorm.DB {
	return db.Joins("Info").Where("money > ?", 100)
}

type TMG struct {
	ID   uint
	Name string
}

// 事务
func OrmTransaction() {
	createTableifNotExists(TMG{}, GLOBAL_DB)
	flag := false
	GLOBAL_DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&TMG{
			Name: "coder4j",
		})
		tx.Create(&TMG{
			Name: "aaa",
		})

		tx.Create(&TMG{
			Name: "eee",
		})
		if flag {
			return nil
		} else {
			return errors.New("测试回滚")
		}
	})
}

func createTableifNotExists(obj any, db *gorm.DB) {
	migrator := db.Migrator()
	if !migrator.HasTable(&obj) {
		err := migrator.CreateTable(&obj)
		if err != nil {
			return
		}
	}
}
