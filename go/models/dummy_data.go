package models

import (
	"time"
)

type DummyData struct {
	Name        string
	DummyString string
	DummyInt    int
	DummyFloat  float64
	DummyBool   bool

	DummyEnumString ElementType

	DummyEnumInt DummnyTypeInt

	DummyTime                time.Time
	DummyDuration            time.Duration
	DummyPointerToGongStruct *AnotherDummyData
}

type DummnyTypeInt int

// values for EnumType
const (
	ONE DummnyTypeInt = 1
	TWO DummnyTypeInt = 2
)
