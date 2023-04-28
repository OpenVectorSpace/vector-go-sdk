package sdk_wrapper

import "net/http"

/*
  These functions work only on an OSKR bot. You can call them on a production bot, but they will return an error.
  They rely on some components running only on dev builds (like the Anim/Engine web servers)
*/

// Triggers a Fake back button press, that will in turn trigger a wake word event (unless Alexa is enabled)

func TriggerWakeWord() error {
	return SetAnimProcessVariable("FakeButtonPressType", "1")
}

// Sets a variable exposed by the vic-anim process using the OSKR web server

func SetAnimProcessVariable(key string, value string) error {
	ipAddr := Robot.GetIPAddress()
	targetUrl := "http://" + ipAddr + ":8889/consolevarset?key=" + key + "&value=" + value

	resp, err := http.Get(targetUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}
