package defs

import "github.com/iancoleman/orderedmap"

// Hoxy 22:46:24 INFO >>## [Mission/endTurn]#156794553100029

// Mission/endTurn
type SMissionEndTurn struct {
	ChangeBelong1          *orderedmap.OrderedMap `json:"change_belong1,omitempty"`
	ChangeBelong           *orderedmap.OrderedMap `json:"change_belong,omitempty"`
	BuildingChangeBelong   *[]interface{}         `json:"building_change_belong,omitempty"`
	BuildingChangeBelong1  *[]interface{}         `json:"building_change_belong1,omitempty"`
	BuildingDefenderChange *[]interface{}         `json:"building_defender_change,omitempty"`
	// TODO: define lose event
	MissionLoseResult *[]interface{} `json:"mission_lose_result,omitempty"`
	// Need a custom marshaller for this, when the json value is [] the struct should be nil
	MissionWinResult *struct {
		Rank   string `json:"rank"`
		Medal4 int    `json:"medal4"`
		Open   struct {
			Num3 int `json:"3"`
		} `json:"open"`
		UserExp   string `json:"user_exp"`
		RewardGun struct {
			GunWithUserID string `json:"gun_with_user_id"`
			GunID         string `json:"gun_id"`
		} `json:"reward_gun"`
		MissionInfo struct {
			Turn                    string `json:"turn"`
			EnemydieNum             string `json:"enemydie_num"`
			EnemydieNumKillbyfriend string `json:"enemydie_num_killbyfriend"`
			GundieNum               string `json:"gundie_num"`
		} `json:"mission_info"`
	} `json:"mission_win_result,omitempty"`
	AllyBattle            *[]interface{} `json:"ally_battle,omitempty"`
	Type5Score            string         `json:"type5_score"`
	FairySkillReturn      []interface{}  `json:"fairy_skill_return"`
	FairySkillPerform     []interface{}  `json:"fairy_skill_perform"`
	FairySkillOnSpot      []interface{}  `json:"fairy_skill_on_spot"`
	FairySkillOnTeam      []interface{}  `json:"fairy_skill_on_team"`
	FairySkillOnEnemy     []interface{}  `json:"fairy_skill_on_enemy"`
	FairySkillOnSquad     []interface{}  `json:"fairy_skill_on_squad"`
	AllyInstanceTransform []interface{}  `json:"ally_instance_transform"`
	AllyInstanceBetray    []interface{}  `json:"ally_instance_betray"`
	MissionControl        struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
	} `json:"mission_control"`
}
