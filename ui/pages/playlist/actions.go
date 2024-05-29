package playlist

import (
	"math/rand"

	"github.com/topwuther/replayer/ui/pages/settings"
)

func autoAction() {
	player.Stop()
	repeat := settings.SettingConfig.Repeat
	random := settings.SettingConfig.Random
	single := settings.SettingConfig.Single
	switch {
	case random:
		playIndex = rand.Intn(len(musics))
	case repeat:
		playIndex++
		if playIndex == len(musics) {
			playIndex = 0
		}
	}
	if single {
		player.Play(&musics[playIndex])
	}
}

func nextAction() {
	player.Stop()
	random := settings.SettingConfig.Random
	switch {
	case random:
		playIndex = rand.Intn(len(musics))
	default:
		playIndex++
		if playIndex == len(musics) {
			playIndex = 0
		}
	}
	player.Play(&musics[playIndex])
}

func prevAction() {
	player.Stop()
	random := settings.SettingConfig.Random
	switch {
	case random:
		playIndex = rand.Intn(len(musics))
	default:
		playIndex--
		if playIndex == -1 {
			playIndex = len(musics) - 1
		}
	}
	player.Play(&musics[playIndex])
}
