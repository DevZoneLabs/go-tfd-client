package models

type BaseStat struct {
	ID    string  `json:"stat_id"`
	Value float64 `json:"stat_value"`
}

type VoidBattleResponse struct {
	ID   string `json:"void_battle_id"`
	Name string `json:"void_battle_name"`
}

type TitleResponse struct {
	ID   string `json:"title_id"`
	Name string `json:"title_name"`
}
