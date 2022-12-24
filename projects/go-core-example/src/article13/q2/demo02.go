package main

import "fmt"

type Cat struct {
	name           string // 名字
	scientificName string // 学名
	category       string // 动物学基本分类
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}

func main() {
	cat := New("a001", "ShangHai", "cat")
	// 该类型的指针类型的方法集合囊括了前者的所有方法，包括所有值方法和指针方法。
	// 严格来讲，我们在这样的基本类型的值上只能调用它的值方法，
	// 但是，GO语言会适时地为我们进行自动转换，使我们的值也可以调用它的指针方法。
	cat.SetName("name02") // (&cat).SetName("name02")
	fmt.Printf("cat : %s\n", cat)

	// 值方法的接收者是该方法所属的那个类型值的一个副本，对该副本的修改不会体现到原值上
	cat.SetNameOfCopy("copy name03")
	fmt.Printf("cat : %s \n", cat)

	// 一个指针类型实现了某某接口类型，但它的基本类型却不一定能够作为该接口的实现类型
	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}
	_, ok := interface{}(cat).(Pet)
	fmt.Printf("cat 实现了Pet接口：%v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*cat 实现了Pet接口：%v\n", ok)
}
