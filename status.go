package main

type status int

const (
	StatusBerserk status = iota
	StatusSlow
	StatusExhausted
	StatusSwift
	StatusLignification
	StatusConfusion
	StatusNausea
	StatusDisabledShield
	StatusCorrosion
	StatusFlames // fake status
	StatusHidden
	StatusUnhidden
	StatusDig
	StatusSwap
	StatusShadows
)

func (st status) Good() bool {
	switch st {
	case StatusBerserk, StatusSwift, StatusDig, StatusSwap, StatusShadows, StatusHidden:
		return true
	default:
		return false
	}
}

func (st status) Bad() bool {
	switch st {
	case StatusSlow, StatusConfusion, StatusNausea, StatusDisabledShield, StatusCorrosion, StatusUnhidden:
		return true
	default:
		return false
	}
}

func (st status) String() string {
	switch st {
	case StatusBerserk:
		return "Berserk"
	case StatusSlow:
		return "Slow"
	case StatusExhausted:
		return "Exhausted"
	case StatusSwift:
		return "Swift"
	case StatusLignification:
		return "Lignified"
	case StatusConfusion:
		return "Confused"
	case StatusNausea:
		return "Nausea"
	case StatusDisabledShield:
		return "-Shield"
	case StatusCorrosion:
		return "Corroded"
	case StatusFlames:
		return "Flames"
	case StatusHidden:
		return "Hidden"
	case StatusUnhidden:
		return "Unhidden"
	case StatusDig:
		return "Dig"
	case StatusSwap:
		return "Swap"
	case StatusShadows:
		return "Shadows"
	default:
		// should not happen
		return "unknown"
	}
}

func (st status) Short() string {
	switch st {
	case StatusBerserk:
		return "Be"
	case StatusSlow:
		return "Sl"
	case StatusExhausted:
		return "Ex"
	case StatusSwift:
		return "Sw"
	case StatusLignification:
		return "Li"
	case StatusConfusion:
		return "Co"
	case StatusNausea:
		return "Na"
	case StatusDisabledShield:
		return "-S"
	case StatusCorrosion:
		return "Co"
	case StatusFlames:
		return "Fl"
	case StatusDig:
		return "Di"
	case StatusSwap:
		return "Sw"
	case StatusShadows:
		return "Sh"
	default:
		// should not happen
		return "?"
	}
}
