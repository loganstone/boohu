// +build !tcell

package main

import (
	"errors"

	termbox "github.com/nsf/termbox-go"
)

type termui struct {
}

func WindowsPalette() {
	ColorBgLOS = color(termbox.ColorWhite)
	ColorBgDark = color(termbox.ColorBlack)
	ColorBg = color(termbox.ColorBlack)
	ColorBgCloud = color(termbox.ColorWhite)
	ColorFgLOS = color(termbox.ColorBlack)
	ColorFgDark = color(termbox.ColorWhite)
	ColorFg = color(termbox.ColorWhite)
	ColorFgPlayer = color(termbox.ColorBlue)
	ColorFgMonster = color(termbox.ColorRed)
	ColorFgSleepingMonster = color(termbox.ColorCyan)
	ColorFgWanderingMonster = color(termbox.ColorMagenta)
	ColorFgConfusedMonster = color(termbox.ColorGreen)
	ColorFgCollectable = color(termbox.ColorYellow)
	ColorFgStairs = color(termbox.ColorMagenta)
	ColorFgGold = color(termbox.ColorYellow)
	ColorFgHPok = color(termbox.ColorGreen)
	ColorFgHPwounded = color(termbox.ColorYellow)
	ColorFgHPcritical = color(termbox.ColorRed)
	ColorFgMPok = color(termbox.ColorBlue)
	ColorFgMPpartial = color(termbox.ColorMagenta)
	ColorFgMPcritical = color(termbox.ColorRed)
	ColorFgStatusGood = color(termbox.ColorBlue)
	ColorFgStatusBad = color(termbox.ColorRed)
	ColorFgStatusOther = color(termbox.ColorYellow)
	ColorFgTargetMode = color(termbox.ColorCyan)
	ColorFgTemporalWall = color(termbox.ColorCyan)
}

func (ui *termui) Init() error {
	return termbox.Init()
}

func (ui *termui) Close() {
	termbox.Close()
}

func (ui *termui) PostInit() {
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
}

func (ui *termui) Clear() {
	termbox.Clear(termbox.Attribute(ColorFg), termbox.Attribute(ColorBg))
}

func (ui *termui) Flush() {
	termbox.Flush()
}

func (ui *termui) Interrupt() {
	termbox.Interrupt()
}

func (ui *termui) HideCursor() {
	termbox.HideCursor()
}

func (ui *termui) SetCursor(pos position) {
	termbox.SetCursor(pos.X, pos.Y)
}

func (ui *termui) SetCell(x, y int, r rune, fg, bg color) {
	termbox.SetCell(x, y, r, termbox.Attribute(fg), termbox.Attribute(bg))
}

func (ui *termui) Reverse(c color) color {
	return color(termbox.Attribute(c) | termbox.AttrReverse)
}

func (ui *termui) WaitForContinue(g *game) {
loop:
	for {
		switch tev := termbox.PollEvent(); tev.Type {
		case termbox.EventKey:
			if tev.Ch == 0 {
				switch tev.Key {
				case termbox.KeyEsc, termbox.KeySpace:
					break loop
				}
			}
			if tev.Ch == ' ' {
				break loop
			}
		}
	}
}

func (ui *termui) PromptConfirmation(g *game) bool {
	for {
		switch tev := termbox.PollEvent(); tev.Type {
		case termbox.EventKey:
			if tev.Ch == 'Y' || tev.Ch == 'y' {
				return true
			}
		}
		return false
	}
}

func (ui *termui) PressAnyKey() error {
	for {
		switch tev := termbox.PollEvent(); tev.Type {
		case termbox.EventKey:
			return nil
		case termbox.EventInterrupt:
			return errors.New("interrupted")
		case termbox.EventMouse:
			if tev.Ch == 0 && tev.Key == termbox.MouseLeft {
				return nil
			}
		}
	}
}

func (ui *termui) HandlePlayerTurn(g *game, ev event) bool {
getKey:
	for {
		ui.DrawDungeonView(g, false)
		var err error
		switch tev := termbox.PollEvent(); tev.Type {
		case termbox.EventKey:
			if tev.Ch == 0 {
				switch tev.Key {
				case termbox.KeyArrowUp:
					tev.Ch = 'k'
				case termbox.KeyArrowRight:
					tev.Ch = 'l'
				case termbox.KeyArrowDown:
					tev.Ch = 'j'
				case termbox.KeyArrowLeft:
					tev.Ch = 'h'
				case termbox.KeyCtrlW:
					ui.EnterWizard(g)
					continue getKey
				case termbox.KeyCtrlQ:
					if ui.Quit(g) {
						return true
					}
					continue getKey
				case termbox.KeyCtrlP:
					tev.Ch = 'm'
				}
			}
			var again, quit bool
			err, again, quit = ui.HandleCharacter(g, ev, tev.Ch)
			if again {
				continue getKey
			} else if quit {
				return true
			}
			if err != nil {
				g.Print(err.Error())
				continue getKey
			}
			return false
		case termbox.EventMouse:
			action := false
			if tev.Ch == 0 {
				switch tev.Key {
				case termbox.MouseLeft:
					pos := position{X: tev.MouseX, Y: tev.MouseY}
					err, action = ui.GoToPos(g, ev, pos)
				case termbox.MouseRight:
					pos := position{X: tev.MouseX, Y: tev.MouseY}
					var again bool
					again, action = ui.ExaminePos(g, ev, pos)
					if again {
						continue getKey
					}
				}
			}
			if err != nil {
				g.Print(err.Error())
				continue getKey
			}
			if action {
				return false
			}
		}
	}
}

func (ui *termui) Scroll(n int) (m int, quit bool) {
	switch tev := termbox.PollEvent(); tev.Type {
	case termbox.EventKey:
		if tev.Ch == 0 {
			switch tev.Key {
			case termbox.KeyEsc, termbox.KeySpace:
				quit = true
				return n, quit
			}
		}
		switch tev.Ch {
		case 'u':
			n -= 12
		case 'd':
			n += 12
		case 'j':
			n++
		case 'k':
			n--
		case ' ':
			quit = true
		}
	}
	return n, quit
}

func (ui *termui) TargetModeEvent(g *game, targ Targetter, pos position, data *examineData) bool {
	switch tev := termbox.PollEvent(); tev.Type {
	case termbox.EventKey:
		if tev.Ch == 0 {
			switch tev.Key {
			case termbox.KeyArrowUp:
				tev.Ch = 'k'
			case termbox.KeyArrowRight:
				tev.Ch = 'l'
			case termbox.KeyArrowDown:
				tev.Ch = 'j'
			case termbox.KeyArrowLeft:
				tev.Ch = 'h'
			case termbox.KeyEsc:
				return true
			case termbox.KeyEnter:
				tev.Ch = '.'
			}
		}
		if ui.CursorCharAction(g, targ, tev.Ch, pos, data) {
			return true
		}
	case termbox.EventMouse:
		if tev.Ch == 0 {
			switch tev.Key {
			case termbox.MouseLeft:
				if ui.CursorMouseLeft(g, targ, pos) {
					return true
				}
			case termbox.MouseRight:
				data.npos = position{X: tev.MouseX, Y: tev.MouseY}
			}
		}
	}
	return false
}

func (ui *termui) Select(g *game, ev event, l int) (index int, alternate bool, err error) {
	for {
		switch tev := termbox.PollEvent(); tev.Type {
		case termbox.EventKey:
			if tev.Ch == 0 {
				switch tev.Key {
				case termbox.KeyEsc, termbox.KeySpace:
					return -1, false, errors.New("Ok, then.")
				}
			}
			if 97 <= tev.Ch && int(tev.Ch) < 97+l {
				return int(tev.Ch - 97), false, nil
			}
			if tev.Ch == '?' {
				return -1, true, nil
			}
			if tev.Ch == ' ' {
				return -1, false, errors.New("Ok, then.")
			}
		}
	}
}