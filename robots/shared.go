package robots

import "log"

// Robots contains a map of robots
var Robots = make(map[string]func() Robot)

// RegisterRobot registers a command and init function for a robot
func RegisterRobot(command string, robotInitFunction func() Robot) {
	if _, ok := Robots[command]; ok {
		log.Printf("There are two robots mapped to %s!", command)
	} else {
		Robots[command] = robotInitFunction
	}
}
