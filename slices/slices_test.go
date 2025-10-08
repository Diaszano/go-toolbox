package slices_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Diaszano/go-toolbox/slices"
	"github.com/stretchr/testify/require"
)

func TestForEach(t *testing.T) {
	input := []int{1, 2, 3}
	var sum int
	slices.ForEach(input, func(v int) {
		sum += v
	})
	require.Equal(t, 6, sum)
}

func TestForEachWithIndex(t *testing.T) {
	input := []string{"a", "b", "c"}
	var result string
	slices.ForEachWithIndex(input, func(i int, v string) {
		result += fmt.Sprintf("%s%d", v, i)
	})
	require.Equal(t, "a0b1c2", result)
}

func TestTryForEach(t *testing.T) {
	input := []int{1, 2, 3, 4}
	err := slices.TryForEach(input, func(v int) error {
		if v == 3 {
			return errors.New("fail at 3")
		}
		return nil
	})
	require.Error(t, err)
	require.Equal(t, "fail at 3", err.Error())
}

func TestTryForEachWithIndex(t *testing.T) {
	input := []int{5, 6, 7}
	err := slices.TryForEachWithIndex(input, func(i, v int) error {
		if i == 1 {
			return errors.New("fail at index 1")
		}
		return nil
	})
	require.Error(t, err)
	require.Equal(t, "fail at index 1", err.Error())
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	output := slices.Map(input, func(v int) int { return v * 2 })
	require.Equal(t, []int{2, 4, 6}, output)
}

func TestMapWithIndex(t *testing.T) {
	input := []int{10, 20, 30}
	output := slices.MapWithIndex(input, func(i, v int) int { return v + i })
	require.Equal(t, []int{10, 21, 32}, output)
}

func TestTryMap(t *testing.T) {
	input := []int{1, 2, 3, 4}
	output, err := slices.TryMap(input, func(v int) (int, error) {
		if v == 3 {
			return 0, errors.New("fail at 3")
		}
		return v * 3, nil
	})
	require.Error(t, err)
	require.Equal(t, "fail at 3", err.Error())
	require.Nil(t, output)
}

func TestTryMapWithIndex(t *testing.T) {
	input := []int{1, 2, 3}
	output, err := slices.TryMapWithIndex(input, func(i, v int) (int, error) {
		if i == 2 {
			return 0, errors.New("fail at index 2")
		}
		return v + i, nil
	})
	require.Error(t, err)
	require.Equal(t, "fail at index 2", err.Error())
	require.Nil(t, output)
}
