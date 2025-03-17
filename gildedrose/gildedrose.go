package gildedrose

/**
All items have a SellIn value which denotes the number of days we have to sell the item
All items have a Quality value which denotes how valuable the item is
At the end of each day our system lowers both values for every item

Once the sell by date has passed, Quality degrades twice as fast
The Quality of an item is never negative
“Aged Brie” actually increases in Quality the older it gets
The Quality of an item is never more than 50
“Sulfuras”, being a legendary item, never has to be sold or decreases in Quality
“Backstage passes”, like aged brie, increases in Quality as its SellIn value approaches;
    Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
    Quality drops to 0 after the concert
**/

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		item.SellIn, item.Quality = calculate(item.Name, item.SellIn, item.Quality)
	}
}

func calculate(name string, sellIn, quality int) (int, int) {
	sellInChange := -1
	qualityChange := -1

	switch name {
	case "Sulfuras, Hand of Ragnaros":
		return sellIn, quality
	case "Aged Brie":
		qualityChange = 1
	case "Backstage passes to a TAFKAL80ETC concert":
		qualityChange = 1
		if sellIn <= 10 {
			qualityChange = 3
		}
		if sellIn <= 5 {
			qualityChange = 2
		}
	case "Conjured Mana Cake":
		qualityChange = -2
	}

	return sellIn + sellInChange, incQuality(quality, qualityChange)
}

const (
	qualityMax = 50
	qualityMin = 0
)

func incQuality(q, inc int) int {
	sum := q + inc
	if sum > qualityMax {
		return qualityMax
	}
	if sum < qualityMin {
		return qualityMin
	}
	return sum
}
