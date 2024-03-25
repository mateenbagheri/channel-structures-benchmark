package heavystruct

import "math/rand"

type HeavyStruct struct {
	Field1  string
	Field2  string
	Field3  string
	Field4  string
	Field5  string
	Field6  string
	Field7  string
	Field8  string
	Field9  string
	Field10 string
	Field11 string
	Field12 string
	Field13 string
	Field14 string
	Field15 string
	Field16 string
	Field17 string
	Field18 string
	Field19 string
	Field20 string
}

func createHeavyStructWithRandomValues() HeavyStruct {
	return HeavyStruct{
		Field1:  createRandomString(),
		Field2:  createRandomString(),
		Field3:  createRandomString(),
		Field4:  createRandomString(),
		Field5:  createRandomString(),
		Field6:  createRandomString(),
		Field7:  createRandomString(),
		Field8:  createRandomString(),
		Field9:  createRandomString(),
		Field10: createRandomString(),
		Field11: createRandomString(),
		Field12: createRandomString(),
		Field13: createRandomString(),
		Field14: createRandomString(),
		Field15: createRandomString(),
		Field16: createRandomString(),
		Field17: createRandomString(),
		Field18: createRandomString(),
		Field19: createRandomString(),
		Field20: createRandomString(),
	}
}

func createRandomString() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	result := make([]byte, 10)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
