package utils_test

import (
	"testing"
	"time"

	"github.com/DoWithLogic/go-rbac/pkg/utils"
	"github.com/stretchr/testify/require"
)

// Test for NewPointer function
func TestNewPointer(t *testing.T) {
	t.Run("CreatePointerFromValue", func(t *testing.T) {
		value := 42
		ptr := utils.NewPointer(value)
		require.NotNil(t, ptr)
		require.Equal(t, value, *ptr) // Assert that the value pointed to is correct
	})

	t.Run("CreatePointerFromString", func(t *testing.T) {
		strValue := "Hello, World!"
		ptr := utils.NewPointer(strValue)
		require.NotNil(t, ptr)
		require.Equal(t, strValue, *ptr) // Assert that the string value is correct
	})

	t.Run("PointerToStruct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		person := Person{Name: "Alice", Age: 30}
		ptr := utils.NewPointer(person)
		require.NotNil(t, ptr)
		require.Equal(t, person, *ptr) // Assert that the struct value is correct
	})
}

// Test for GetPointerValue function
func TestGetPointerValue(t *testing.T) {
	t.Run("GetValueFromPointer", func(t *testing.T) {
		value := 99
		ptr := utils.NewPointer(value)
		retrievedValue := utils.GetPointerValue(ptr)
		require.Equal(t, value, retrievedValue) // Assert the retrieved value is correct
	})

	t.Run("GetValueFromNilPointer", func(t *testing.T) {
		var ptr *int // Declare a nil pointer
		retrievedValue := utils.GetPointerValue(ptr)
		require.Equal(t, 0, retrievedValue) // Assert that the default zero value for int is returned
	})

	t.Run("GetValueFromPointerToString", func(t *testing.T) {
		strValue := "Go is awesome!"
		ptr := utils.NewPointer(strValue)
		retrievedValue := utils.GetPointerValue(ptr)
		require.Equal(t, strValue, retrievedValue) // Assert the retrieved string value is correct
	})
}

func TestGetValueOrDefault(t *testing.T) {
	t.Run("PointerIsNil", func(t *testing.T) {
		var nilPointer *int
		defaultValue := 42
		result := utils.GetValueOrDefault(nilPointer, defaultValue)
		require.Equal(t, defaultValue, result, "Expected default value when pointer is nil")
	})

	t.Run("PointerHasValue", func(t *testing.T) {
		value := 10
		result := utils.GetValueOrDefault(&value, 42)
		require.Equal(t, value, result, "Expected the value pointed to by the pointer")
	})

	t.Run("TimePointerIsNil", func(t *testing.T) {
		var nilTimePointer *time.Time
		defaultTime := time.Now()
		result := utils.GetValueOrDefault(nilTimePointer, defaultTime)
		require.WithinDuration(t, defaultTime, result, time.Second, "Expected default time when pointer is nil")
	})

	t.Run("TimePointerHasValue", func(t *testing.T) {
		someTime := time.Date(2024, 10, 4, 10, 0, 0, 0, time.UTC)
		result := utils.GetValueOrDefault(&someTime, time.Now())
		require.Equal(t, someTime, result, "Expected the time value pointed to by the pointer")
	})
}
