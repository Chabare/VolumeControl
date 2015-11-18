package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func main() {
	// regex, _ := regexp.Compile("Front Left: Playback [0-9]+ \\[[0-9]+%\\] \\[on\\]")
	regex, _ := regexp.Compile("[0-9]+")
	out, _ := exec.Command("amixer").Output()

	currentVolume, _ := strconv.Atoi(regex.FindAllString(string(out), 4)[3])

	var perc float64 = 5

	if len(os.Args) > 2 {
		var err error

		perc, err = strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			log.Printf("Error: The value isn't a number: %s\n", os.Args[2])
		}
	}

	if os.Args[1] == "-l" || os.Args[1] == "--lower" {
		lowerVolume(currentVolume, perc)
	} else {
		raiseVolume(currentVolume, perc)
	}
}

func lowerVolume(vol int, perc float64) {
	if int(perc) == 100 {
		changeVolume(0)
		return
	}

	fac := int(65535 * perc / 100)
	if vol-fac >= 0 {
		vol = vol - fac
	}
	changeVolume(vol)
}

func raiseVolume(vol int, perc float64) {
	if perc == 100 {
		changeVolume(65535)
		return
	}

	fac := int(65535 * perc / 100)
	if vol+fac-1 <= 65535 {
		vol = vol + fac
	}

	changeVolume(vol)
}

func changeVolume(newVolume int) {
	c := exec.Command("amixer", "set", "Master", strconv.FormatInt(int64(newVolume), 10))
	c.Run()
}
