package lib

import (
	"strconv"
	"time"

	"github.com/fatih/color"
)

type durationTimer struct {
	duration  string
	startTime time.Time
	endTime   time.Time
}

func NewTimer(duration string) *durationTimer {
	return &durationTimer{
		duration: duration,
	}
}

func (t *durationTimer) start() {
	t.startTime = time.Now()
	duration, err := time.ParseDuration(t.duration + "m")
	if err != nil {
		panic(err)
	}
	t.endTime = t.startTime.Add(duration)
}

func Process(workDuration string, breakDuration string, breaksCounts string) error {
	ColorTimerStartWork := color.New(color.Bold).Add(color.Underline).Add(color.FgRed)
	ColorTimerStartBreak := color.New(color.Bold).Add(color.Underline).Add(color.FgGreen)
	breaksCountsInt, _ := strconv.Atoi(breaksCounts)

	count := 0
	for {
		workTime(workDuration, *ColorTimerStartWork)
		if count == breaksCountsInt {
			break
		}
		breakTime(breakDuration, *ColorTimerStartBreak)
		count++
	}
	return nil //don't forget to make error handling
}

func workTime(workDuration string, workColor color.Color) {
	workColor.Print("Work")
	timer := NewTimer(workDuration)
	timer.start()
	for {
		if time.Now().After(timer.endTime) {
			break
		} else {
			workColor.Printf("Time left: %.2f min\n", timer.endTime.Sub(time.Now().Round(time.Second)).Minutes())
			time.Sleep(5 * time.Second)
		}
	}
	workColor.Add(color.FgBlue).Printf("Work Done,time to chill \n")
	PlaySound("sounds/workDoneSoundmp3.mp3")
}

func breakTime(breakDuration string, breakColor color.Color) {
	breakColor.Print("Break")
	timer := NewTimer(breakDuration)
	timer.start()
	for {
		if time.Now().After(timer.endTime) {
			break
		} else {
			breakColor.Printf("Time left: %.2f  min\n", timer.endTime.Sub(time.Now().Round(time.Second)).Minutes())
			time.Sleep(5 * time.Second)
		}
	}
	breakColor.Add(color.FgBlue).Printf("Break Done,time to work \n")
	PlaySound("sounds/breakDoneSoundmp3.mp3")
}
