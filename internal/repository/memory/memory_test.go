package memory_test

import (
	"context"
	"testing"

	"github.com/capcom6/labeled-storage/internal/repository/data"
	"github.com/capcom6/labeled-storage/internal/repository/memory"
)

const (
	keyOne = "one"
	keyTwo = "two"
)

var mockItemSingleLabel = data.NewItem("single", map[string]string{"bar": "baz"})
var mockItemMultiLabel = data.NewItem("multi", map[string]string{"bar": "baz", "qux": "quux"})

func TestMemory(t *testing.T) {
	repo := memory.New()
	ctx := context.Background()

	t.Run("Add", func(t *testing.T) {
		_, err := repo.Replace(ctx, keyOne, *mockItemSingleLabel)
		if err != nil {
			t.Fatal(err)
		}

		_, err = repo.Replace(ctx, keyTwo, *mockItemMultiLabel)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Get", func(t *testing.T) {
		item, err := repo.Get(ctx, keyOne)
		if err != nil {
			t.Fatal(err)
		}
		if item.Value() != mockItemSingleLabel.Value() {
			t.Fatal("unexpected item")
		}
	})

	t.Run("List", func(t *testing.T) {
		items, err := repo.List(ctx, map[string]string{"bar": "baz"})
		if err != nil {
			t.Fatal(err)
		}
		if len(items) != 2 {
			t.Fatal("unexpected items")
		}
	})

	t.Run("List Complex", func(t *testing.T) {
		items, err := repo.List(ctx, map[string]string{"bar": "baz", "qux": "quux"})
		if err != nil {
			t.Fatal(err)
		}
		if len(items) != 1 || items[0].Value() != mockItemMultiLabel.Value() {
			t.Fatal("unexpected item")
		}
	})

	t.Run("Replace", func(t *testing.T) {
		_, err := repo.Replace(ctx, keyOne, *mockItemMultiLabel)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Get Replaced", func(t *testing.T) {
		item, err := repo.Get(ctx, keyOne)
		if err != nil {
			t.Fatal(err)
		}
		if item.Value() != mockItemMultiLabel.Value() {
			t.Fatal("unexpected item")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cnt, err := repo.DeleteMany(ctx, map[string]string{"bar": "baz"})
		if err != nil {
			t.Fatal(err)
		}
		if cnt != 2 {
			t.Fatal("unexpected count")
		}
	})

	t.Run("Get Deleted", func(t *testing.T) {
		_, err := repo.Get(ctx, keyOne)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}
