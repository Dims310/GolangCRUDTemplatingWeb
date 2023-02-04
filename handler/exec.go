package handler

import (
	"rest-api-sim/connection"
	"rest-api-sim/helper"
	"time"
)

func QueryExec(name string) string {
	resultNames, resultAbs, _ := QueryChecker(name)

	if !resultNames {
		time := time.Now()
		exec_script := "INSERT INTO user(name, presence) VALUES (?, ?);"
		_, err := connection.Connection().Exec(exec_script, name, time.Format("2006-01-02 15:04:05"))
		helper.Output(err)
		return "Presence success, have good day! :)"
	} else if resultNames && time.Now().Hour() >= 16 && !resultAbs {
		time := time.Now()
		exec_script := "UPDATE user SET absence=(?), already_absence=(?) WHERE name=(?);"
		_, err := connection.Connection().Exec(exec_script, time.Format("2006-01-02 15:04:05"), true, name)
		helper.Output(err)
		return "Absence success, have good day! :)"
	} else if resultNames && resultAbs {
		return "Records already absence. Thank you :)"
	} else {
		return "Records already presence. You can do absences past or equal to 4 pm."
	}
}

func QueryExecDel(name string) string {
	resultNames, _, resultId := QueryChecker(name)

	if resultNames {
		exec_script := "DELETE FROM user WHERE id=(?)"
		_, err := connection.Connection().Exec(exec_script, resultId)
		helper.Output(err)
		return "Records " + name + " successfully deleted."
	} else {
		return "No records found with name " + name
	}

}
