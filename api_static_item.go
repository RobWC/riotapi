package riotapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ItemDataTags tags that provide modifiers of what data is returned on an item request
type ItemDataTag string

const (
	ItemDataNil                  ItemDataTag = ""
	ItemDataAll                  ItemDataTag = "all"
	ItemDataColloq               ItemDataTag = "colloq"
	ItemDataConsumeOnFull        ItemDataTag = "consumeOnFull"
	ItemDataConsumed             ItemDataTag = "consumed"
	ItemDataDepth                ItemDataTag = "depth"
	ItemDataEffect               ItemDataTag = "effect"
	ItemDataGold                 ItemDataTag = "gold"
	ItemDataGroups               ItemDataTag = "groups"
	ItemDataHideFromAll          ItemDataTag = "hideFromAll"
	ItemDataImage                ItemDataTag = "image"
	ItemDataInStore              ItemDataTag = "inStore"
	ItemDataInfo                 ItemDataTag = "info"
	ItemDataMaps                 ItemDataTag = "maps"
	ItemDataRequiredChampion     ItemDataTag = "requiredChampion"
	ItemDataSanitizedDescription ItemDataTag = "sanitizedDescription"
	ItemDataSpecialRecipe        ItemDataTag = "specialRecipe"
	ItemDataStacks               ItemDataTag = "stacks"
	ItemDataStats                ItemDataTag = "stats"
	ItemDataTags                 ItemDataTag = "tags"
	ItemDataTree                 ItemDataTag = "tree"
)

// https://developer.riotgames.com/api/methods#!/1055/3621

// StaticItem get static champion data
// If version is not defined then it uses the current version
func (c *APIClient) StaticItem(version string, itemData ...ItemDataTag) (id *ItemData, err error) {
	var req *http.Request

	// setup api query options
	q := url.Values{}
	if version != "" {
		q.Add("version", version)
	}

	if len(itemData) > 0 {
		for i := range itemData {
			q.Add("itemListData", string(itemData[i]))
		}
	}

	req, err = c.genStaticRequest("GET", "v1.2", "item", q)
	if err != nil {
		return id, err
	}
	data, err := c.do(req, true)
	if err != nil {
		return id, err
	}

	err = json.Unmarshal(data, &id)
	if err != nil {
		return nil, err
	}

	return id, err

}

// https://developer.riotgames.com/api/methods#!/1055/3627

// StaticItemByID get an item by ID
// If version is not defined then it uses the current version
func (c *APIClient) StaticItemByID(id int, version string, itemData ...ItemDataTag) (it *Item, err error) {
	var req *http.Request

	// setup api query options
	q := url.Values{}
	if version != "" {
		q.Add("version", version)
	}

	if len(itemData) > 0 {
		for i := range itemData {
			q.Add("itemData", string(itemData[i]))
		}
	}

	req, err = c.genStaticRequest("GET", "v1.2", strings.Join([]string{"item", strconv.Itoa(id)}, "/"), q)
	if err != nil {
		return it, err
	}
	data, err := c.do(req, true)
	if err != nil {
		return it, err
	}

	err = json.Unmarshal(data, &it)
	if err != nil {
		return nil, err
	}

	return it, err
}
