package db

import (
	"math"

	"github.com/mbolis/yello/model"
)

var allLists = model.Lists{
	Lists: []model.List{
		{Id: 1, Title: "list n.1"},
		{Id: 2, Title: "list n.2"},
		{Id: 3, Title: "list n.3"},
	},
}

func GetAllLists() model.Lists {
	return allLists
}

func CreateNewList(l *model.List) error {
	var maxId uint
	for _, list := range allLists.Lists {
		maxId = uint(math.Max(float64(maxId), float64(list.Id)))
	}

	l.Id = maxId + 1
	allLists.Lists = append(allLists.Lists, *l)
	return nil
}
