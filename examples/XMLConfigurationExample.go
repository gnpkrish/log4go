/*
* @Author: Narayanaperumal Gurusamy
* @Date:   2015-10-04 22:58:45
* @Last Modified by:   Narayanaperumal Gurusamy
* @Last Modified time: 2015-10-04 23:01:04
 */
package main

import l4g "github.com/gnpkrish/log4go"

func main() {
	// Load the configuration (isn't this easy?)
	l4g.LoadConfiguration("example.xml")

	// And now we're ready!
	l4g.Finest("This will only go to those of you really cool UDP kids!  If you change enabled=true.")
	l4g.Debug("Oh no!  %d + %d = %d!", 2, 2, 2+2)
	l4g.Info("About that time, eh chaps?")
}
