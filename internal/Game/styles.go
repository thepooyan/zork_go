package Game

import "github.com/charmbracelet/lipgloss"

type Style int

const (
  pointer Style = iota
  story
)

func GetStyle(s Style) lipgloss.Style {
  smap := map[Style] lipgloss.Style {
    pointer: lipgloss.NewStyle().Foreground(lipgloss.Color("#eb246c")).Bold(true),
    story: lipgloss.NewStyle().Italic(true).Bold(true).Foreground(lipgloss.Color("#aAFAFA")).Align(lipgloss.Right),
  }
  return smap[s]
}

type emoji int

const (
	laugh emoji = iota
	smirk
	wink
	thumbsUp
	sad
	angry
	heart
	clap
	surprised
	thinking
	crying
	sleeping
	confused
	grinning
	sweat
	hug
	shushing
	nerd
	cool
	rollingEyes
	heartEyes
	kiss
	party
	zany
	moneyFace
	tired
	scream
	facepalm
	shrug
	monocle
	explodingHead
)

func Emoji(e emoji) string {
	emap := map[emoji]string{
		laugh:          "\U0001F606",
		smirk:          "\U0001F60F",
		wink:           "\U0001F609",
		thumbsUp:       "\U0001F44D",
		sad:            "\U0001F622",
		angry:          "\U0001F620",
		heart:          "\U0001F496",
		clap:           "\U0001F44F",
		surprised:      "\U0001F62E",
		thinking:       "\U0001F914",
		crying:         "\U0001F62D",
		sleeping:       "\U0001F634",
		confused:       "\U0001F615",
		grinning:       "\U0001F600",
		sweat:          "\U0001F613",
		hug:            "\U0001F917",
		shushing:       "\U0001F92B",
		nerd:           "\U0001F913",
		cool:           "\U0001F60E",
		rollingEyes:    "\U0001F644",
		heartEyes:      "\U0001F60D",
		kiss:           "\U0001F618",
		party:          "\U0001F973",
		zany:           "\U0001F92A",
		moneyFace:      "\U0001F911",
		tired:          "\U0001F62B",
		scream:         "\U0001F631",
		facepalm:       "\U0001F926",
		shrug:          "\U0001F937",
		monocle:        "\U0001F9D0",
		explodingHead:  "\U0001F92F",
	}

	return emap[e]
}
