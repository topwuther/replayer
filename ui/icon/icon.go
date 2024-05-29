package icon

import (
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var MenuIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.NavigationMenu)
	return icon
}()

var SaveIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ContentSave)
	return icon
}()

var RefreshIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.NavigationRefresh)
	return icon
}()

var RestaurantMenuIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.MapsRestaurantMenu)
	return icon
}()

var AccountBalanceIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionAccountBalance)
	return icon
}()

var AccountBoxIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionAccountBox)
	return icon
}()

var CartIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionAddShoppingCart)
	return icon
}()

var HomeIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionHome)
	return icon
}()

var SettingsIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionSettings)
	return icon
}()

var OtherIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionHelp)
	return icon
}()

var HeartIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionFavorite)
	return icon
}()

var PlusIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ContentAdd)
	return icon
}()

var EditIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ContentCreate)
	return icon
}()

var VisibilityIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionVisibility)
	return icon
}()

var PlayIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.AVPlayArrow)
	return icon
}()

var PauseIcon *widget.Icon = func() *widget.Icon {
	icon,_ := widget.NewIcon(icons.AVPause)
	return icon
}()

var StopIcon *widget.Icon = func() *widget.Icon {
	icon,_ := widget.NewIcon(icons.AVStop)
	return icon
}()

var NextIcon *widget.Icon = func() *widget.Icon {
	icon,_ := widget.NewIcon(icons.AVSkipNext)
	return icon
}()

var PreviousIcon *widget.Icon = func() *widget.Icon {
	icon,_ := widget.NewIcon(icons.AVSkipPrevious)
	return icon
}()

