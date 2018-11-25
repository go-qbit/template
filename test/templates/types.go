package templates

type User struct {
	Name     string
	Lastname string
	Age      uint8
	IsMan    bool
}

type TestType struct {
	StructField struct {
		F1 map[string]int
	}
}

type StringMap map[string]string
