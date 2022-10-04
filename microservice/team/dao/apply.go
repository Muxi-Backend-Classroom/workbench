package dao

import (
	team "workbench-team/proto"
	m "workbench/model"
)

func GetGroup(id uint32) (GroupsModel, error) {
	var group = GroupsModel{}
	query := m.DB.Self.Model(&ApplysModel{}).Where("id=?", id)
	err := query.First(&group).Error
	return group, err
}

func GetTeam(id uint32) (TeamsModel, error) {
	var team = TeamsModel{}
	query := m.DB.Self.Model(&GroupsModel{}).Where("id=?", id)
	err := query.First(&team).Error
	return team, err
}

func UpdateTeamAndGroup(id, group, team uint32) error {
	query := m.DB.Self.Table("users").Where("id=?", id)
	err := query.Updates(map[string]interface{}{
		"group_id": group,
		"team_id":  team,
	}).Error
	return err
}

func GetAllApplys() ([]*team.Apply, error) {
	var rets = make([]*team.Apply, 0)
	query := m.DB.Self.Table("applys").Where("1=1")
	err := query.Find(&rets).Error
	return rets, err
}

func DeleteApply(apply ApplysModel) error {
	return m.DB.Self.Delete(&apply).Error
}

func GetApply(id uint32) (ApplysModel, error) {
	var apply = ApplysModel{}
	query := m.DB.Self.Where("id=?", id)
	err := query.First(&apply).Error
	return apply, err
}
