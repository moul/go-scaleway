// Copyright (C) 2015  Nicolas Lamirault <nicolas.lamirault@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package logging

import (
	//"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

// InitLogging define log level
func InitLogging(level log.Level) {
	//log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stderr)
	log.SetLevel(level)
}

// Debug print message using the Debug level color
func Debug(args ...interface{}) {
	log.Debug("[OnlineLabs] ", args)
}

// Info print message using the INFO level color
func Info(args ...interface{}) {
	log.Info("[OnlineLabs] ", args)
}

// Warn print message using the WARN level color
func Warn(args ...interface{}) {
	log.Warn("[OnlineLabs] ", args)
}

// Error print message using the ERROR level color
func Error(args ...interface{}) {
	log.Error("[OnlineLabs] ", args)
}
