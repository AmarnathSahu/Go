package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveItemsFromList(t *testing.T) {
	type args struct {
		originalList []string
		removedList  []string
	}

	tests := []struct {
		name string
		arg  args
		want []string
	}{
		{
			name: "should return empty list",
			arg:  args{originalList: []string{}, removedList: []string{"sales"}},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveItemsFromList(tt.arg.originalList, tt.arg.removedList)
			assert.Equal(t, tt.want, got)

		})
	}
}
