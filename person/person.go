package person

type Person struct {
	name string "未导出的字段，在外部访问时将不可见，会抛出 unexported错误"
	age  int    "同上"
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) (newName string) {
	p.name = name

	return
}
