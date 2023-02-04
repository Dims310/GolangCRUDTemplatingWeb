package handler

import (
	"rest-api-sim/connection"
	"rest-api-sim/model"
)

func QueryAllData() model.User {
	var sUser model.User
	rows, _ := connection.Connection().Query("SELECT id, name, presence, absence, already_absence FROM user")
	defer rows.Close()

	for rows.Next() {
		var id int
		var names, presence, absence string
		var alreadyAbs bool
		rows.Scan(&id, &names, &presence, &absence, &alreadyAbs)

		sUser.ID = append(sUser.ID, id)
		sUser.Name = append(sUser.Name, names)
		sUser.Presence = append(sUser.Presence, presence)
		sUser.Absence = append(sUser.Absence, absence)
		sUser.AlreadyAbs = append(sUser.AlreadyAbs, alreadyAbs)
	}

	return sUser
}

func QueryChecker(name string) (rNames bool, rAbs bool, rId int) {
	resultId := QueryAllData().ID
	resultName := QueryAllData().Name
	resultAbs := QueryAllData().AlreadyAbs

	rNames = false

	for i := 0; i < len(resultName); i++ {
		if resultName[i] == name {
			rNames = true
			rAbs = resultAbs[i]
			rId = resultId[i]
			break
		}
	}
	return
}
