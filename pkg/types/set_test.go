package types

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewSet(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		want     Set[int]
	}{
		{
			name:     "empty set",
			elements: []int{},
			want:     make(Set[int]),
		},
		{
			name:     "single element",
			elements: []int{1},
			want:     Set[int]{1: struct{}{}},
		},
		{
			name:     "multiple elements",
			elements: []int{1, 2, 3},
			want:     Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name:     "duplicate elements",
			elements: []int{1, 2, 1, 3, 2},
			want:     Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSet(tt.elements...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	tests := []struct {
		name    string
		initial Set[string]
		add     string
		want    Set[string]
	}{
		{
			name:    "add to empty set",
			initial: make(Set[string]),
			add:     "apple",
			want:    Set[string]{"apple": struct{}{}},
		},
		{
			name:    "add to non-empty set",
			initial: Set[string]{"banana": struct{}{}, "orange": struct{}{}},
			add:     "apple",
			want:    Set[string]{"apple": struct{}{}, "banana": struct{}{}, "orange": struct{}{}},
		},
		{
			name:    "add existing element",
			initial: Set[string]{"apple": struct{}{}, "banana": struct{}{}},
			add:     "apple",
			want:    Set[string]{"apple": struct{}{}, "banana": struct{}{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.Add(tt.add)
			if !reflect.DeepEqual(tt.initial, tt.want) {
				t.Errorf("After Add() = %v, want %v", tt.initial, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	tests := []struct {
		name    string
		initial Set[int]
		remove  int
		want    Set[int]
	}{
		{
			name:    "remove from empty set",
			initial: make(Set[int]),
			remove:  1,
			want:    make(Set[int]),
		},
		{
			name:    "remove existing element",
			initial: Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			remove:  2,
			want:    Set[int]{1: struct{}{}, 3: struct{}{}},
		},
		{
			name:    "remove non-existing element",
			initial: Set[int]{1: struct{}{}, 3: struct{}{}},
			remove:  2,
			want:    Set[int]{1: struct{}{}, 3: struct{}{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.Remove(tt.remove)
			if !reflect.DeepEqual(tt.initial, tt.want) {
				t.Errorf("After Remove() = %v, want %v", tt.initial, tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	set := NewSet(1, 2, 3)

	tests := []struct {
		name    string
		element int
		want    bool
	}{
		{
			name:    "contains existing element",
			element: 2,
			want:    true,
		},
		{
			name:    "does not contain non-existing element",
			element: 4,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := set.Contains(tt.element); got != tt.want {
				t.Errorf("Set.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		set  Set[string]
		want bool
	}{
		{
			name: "empty set",
			set:  make(Set[string]),
			want: true,
		},
		{
			name: "non-empty set",
			set:  Set[string]{"a": struct{}{}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.IsEmpty(); got != tt.want {
				t.Errorf("Set.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Size(t *testing.T) {
	tests := []struct {
		name string
		set  Set[int]
		want int
	}{
		{
			name: "empty set",
			set:  make(Set[int]),
			want: 0,
		},
		{
			name: "single element",
			set:  Set[int]{1: struct{}{}},
			want: 1,
		},
		{
			name: "multiple elements",
			set:  Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Size(); got != tt.want {
				t.Errorf("Set.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Elements(t *testing.T) {
	tests := []struct {
		name string
		set  Set[string]
		want []string
	}{
		{
			name: "empty set",
			set:  make(Set[string]),
			want: []string{},
		},
		{
			name: "single element",
			set:  Set[string]{"a": struct{}{}},
			want: []string{"a"},
		},
		{
			name: "multiple elements",
			set:  Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			want: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Elements()
			// Sort to ensure deterministic comparison
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Elements() = %v, want %v", got, tt.want)
			}
		})
	}
}