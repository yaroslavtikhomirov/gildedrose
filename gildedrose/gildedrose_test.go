package gildedrose_test

import (
	"encoding/json"
	"testing"

	"github.com/yaroslavtikhomirov/gildedrose/gildedrose"
)

func Test_ShouldDeclineByOne(t *testing.T) {
	items := []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	gildedrose.UpdateQuality(items)

	expected := []*gildedrose.Item{
		{"+5 Dexterity Vest", 9, 19},
		{"Aged Brie", 1, 1},
		{"Elixir of the Mongoose", 4, 6},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		{"Conjured Mana Cake", 2, 4},
	}

	got, _ := json.MarshalIndent(items, "", "  ")
	want, _ := json.MarshalIndent(expected, "", "  ")

	if string(got) != string(want) {
		t.Errorf("want %s\n but got %s\n", string(want), string(got))
	}
}
