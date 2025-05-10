package types

import "strings"

type ItemRequest struct {
	Id     int
	ApiKey string
}

type ItemResponse struct {
	StatusCode    int     `json:"-"`
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	ItemTypeId    int     `json:"itemTypeId"`
	ItemSubTypeId int     `json:"itemSubTypeId"`
	Attack        *int    `json:"attack"`  // For weapons
	Defense       *int    `json:"defense"` // For armors/headgears
	Weight        int     `json:"weight"`
	RequiredLevel *int    `json:"requiredLevel"`
	Slots         *int    `json:"slots"`
	Job           *int    `json:"job"`
	Location      *string `json:"location"` // e.g., "Upper"
}

type ItemYAML struct {
	Id            int    `yaml:"Id"`
	AegisName     string `yaml:"AegisName"`
	Name          string `yaml:"Name"`
	Type          int    `yaml:"Type"`
	Buy           int    `yaml:"Buy"`
	Sell          int    `yaml:"Sell"`
	Weight        int    `yaml:"Weight"`
	DEF           int    `yaml:"DEF,omitempty"`
	ATK           int    `yaml:"ATK,omitempty"`
	Slots         int    `yaml:"Slots"`
	Job           int    `yaml:"Job"`
	Upper         int    `yaml:"Upper"` // Usually 63 = all
	Loc           int    `yaml:"Loc"`   // 256 = Upper, 512 = Mid, 1 = Lower, 769 = All
	EquipLevelMin int    `yaml:"EquipLevelMin"`
	Refineable    bool   `yaml:"Refineable"`
	View          int    `yaml:"View"`
}

func ConvertToItemYAML(resp ItemResponse) ItemYAML {
	locMap := map[string]int{
		"Upper":              256,
		"Middle":             512,
		"Lower":              1,
		"Upper,Middle,Lower": 769,
	}

	loc := 0
	if resp.Location != nil {
		loc = locMap[*resp.Location]
	}

	slots := 0
	if resp.Slots != nil {
		slots = *resp.Slots
	}

	def := 0
	if resp.Defense != nil {
		def = *resp.Defense
	}

	atk := 0
	if resp.Attack != nil {
		atk = *resp.Attack
	}

	lvl := 0
	if resp.RequiredLevel != nil {
		lvl = *resp.RequiredLevel
	}

	job := 0
	if resp.Job != nil {
		job = *resp.Job
	}

	return ItemYAML{
		Id:            resp.Id,
		AegisName:     strings.ReplaceAll(strings.Title(strings.ReplaceAll(resp.Name, " ", "_")), "-", ""),
		Name:          resp.Name,
		Type:          resp.ItemTypeId,
		Buy:           20,
		Sell:          10,
		Weight:        resp.Weight,
		DEF:           def,
		ATK:           atk,
		Slots:         slots,
		Job:           job,
		Upper:         63,
		Loc:           loc,
		EquipLevelMin: lvl,
		Refineable:    true,
		View:          0,
	}
}
