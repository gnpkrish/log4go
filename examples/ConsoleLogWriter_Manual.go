/*
* @Author: Narayanaperumal Gurusamy
* @Date:   2015-10-04 22:58:24
* @Last Modified by:   Narayanaperumal Gurusamy
* @Last Modified time: 2015-10-04 23:01:25
 */
package main

import (
	"time"
)

import l4g "github.com/gnpkrish/log4go"

func main() {
	log := l4g.NewLogger()
	log.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
}
