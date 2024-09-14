// Package window defines standard windows for application GUIs.
// This package has been copied from fyne.io/fyne/v2/dialog
// So we have to update and maintain this forever. :(
package window // import "fyne.io/fyne/v2/dialog"

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	padWidth  = 32
	padHeight = 16
)

// Dialog is the common API for any dialog window with a single dismiss button
type Dialog interface {
	Show()
	Hide()
	SetDismissText(label string)
	SetOnClosed(closed func())
	Refresh()
	Resize(size fyne.Size)
	SetFixedSize(fix bool)

	// Since: 2.1
	MinSize() fyne.Size
}

// Declare conformity to Dialog interface
var _ Dialog = (*dialog)(nil)

type dialog struct {
	callback    func(bool)
	title       string
	icon        fyne.Resource
	desiredSize fyne.Size

	win     fyne.Window
	content fyne.CanvasObject
	dismiss *widget.Button
	parent  fyne.App

	// allows derived dialogs to inject logic that runs before Show()
	beforeShowHook func()
}

func (d *dialog) Hide() {
	d.hideWithResponse(false)
}

// MinSize returns the size that this window should not shrink below
func (d *dialog) MinSize() fyne.Size {
	return d.win.Canvas().Size()
}

func (d *dialog) Show() {
	if d.beforeShowHook != nil {
		d.beforeShowHook()
	}
	if !d.desiredSize.IsZero() {
		d.win.Resize(d.desiredSize)
	}
	d.win.Show()
}

func (d *dialog) Refresh() {
	d.win.Canvas().Refresh(d.win.Content())
}

// Resize dialog, call this function after dialog show
func (d *dialog) Resize(size fyne.Size) {
	d.desiredSize = size
	d.win.Resize(size)
}

// SetFixedSize sets the window to fixed size and centers it on the screen
func (d *dialog) SetFixedSize(fix bool) {
	d.win.SetFixedSize(true)
	d.win.CenterOnScreen()
}

// SetDismissText allows custom text to be set in the dismiss button
// This is a no-op for dialogs without dismiss buttons.
func (d *dialog) SetDismissText(label string) {
	if d.dismiss == nil {
		return
	}

	d.dismiss.SetText(label)
	d.Refresh()
}

// SetOnClosed allows to set a callback function that is called when
// the dialog is closed
func (d *dialog) SetOnClosed(closed func()) {
	// if there is already a callback set, remember it and call both
	originalCallback := d.callback

	d.callback = func(response bool) {
		closed()
		if originalCallback != nil {
			originalCallback(response)
		}
	}
}

func (d *dialog) hideWithResponse(resp bool) {
	d.win.Hide()
	if d.callback != nil {
		d.callback(resp)
	}
}

func (d *dialog) create(buttons fyne.CanvasObject) {
	var image fyne.CanvasObject
	if d.icon != nil {
		image = &canvas.Image{Resource: d.icon}
		fmt.Println("Icon found: ", d.icon.Name())
	} else {
		image = &layout.Spacer{}
		fmt.Println("No icon found: Use spacer", image.MinSize())
	}

	fmt.Printf("Content size (create): %fx%f\n", d.content.Size().Width, d.content.Size().Height)
	content := container.New(&dialogLayout{d: d},
		image,
		newThemedBackground(),
		d.content,
		buttons,
	)

	d.win = d.parent.NewWindow(d.title)
	d.win.SetContent(content)
}

func (d *dialog) setButtons(buttons fyne.CanvasObject) {
	d.win.Content().(*fyne.Container).Objects[3] = buttons
	d.Refresh()
}

// ===============================================================
// ThemedBackground
// ===============================================================

type themedBackground struct {
	widget.BaseWidget
}

func newThemedBackground() *themedBackground {
	t := &themedBackground{}
	t.ExtendBaseWidget(t)
	return t
}

func (t *themedBackground) CreateRenderer() fyne.WidgetRenderer {
	t.ExtendBaseWidget(t)
	rect := canvas.NewRectangle(theme.Color(theme.ColorNameOverlayBackground))
	return &themedBackgroundRenderer{rect, []fyne.CanvasObject{rect}}
}

type themedBackgroundRenderer struct {
	rect    *canvas.Rectangle
	objects []fyne.CanvasObject
}

func (renderer *themedBackgroundRenderer) Destroy() {
}

func (renderer *themedBackgroundRenderer) Layout(size fyne.Size) {
	renderer.rect.Resize(size)
}

func (renderer *themedBackgroundRenderer) MinSize() fyne.Size {
	return renderer.rect.MinSize()
}

func (renderer *themedBackgroundRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

func (renderer *themedBackgroundRenderer) Refresh() {
	renderer.rect.FillColor = theme.Color(theme.ColorNameBackground)
}

// ===============================================================
// DialogLayout
// ===============================================================

type dialogLayout struct {
	d *dialog
}

func (l *dialogLayout) Layout(obj []fyne.CanvasObject, size fyne.Size) {
	btnMin := obj[3].MinSize()

	// icon
	iconHeight := btnMin.Height * 2
	obj[0].Resize(fyne.NewSize(iconHeight, iconHeight))
	obj[0].Move(fyne.NewPos(theme.Padding(), theme.Padding()))

	// background
	obj[1].Move(fyne.NewPos(0, 0))
	obj[1].Resize(size)

	// content
	fmt.Printf("Content size (Layout): %fx%f\n", obj[2].Size().Width, obj[2].Size().Height)
	contentStart := float32(padHeight) / 2
	contentEnd := obj[3].Position().Y - theme.Padding()
	obj[2].Move(fyne.NewPos(padWidth/2, contentStart))
	obj[2].Resize(fyne.NewSize(size.Width-padWidth, contentEnd-contentStart))

	// buttons
	obj[3].Resize(btnMin)
	obj[3].Move(fyne.NewPos(size.Width/2-(btnMin.Width/2), size.Height-padHeight-btnMin.Height))
}

func (l *dialogLayout) MinSize(obj []fyne.CanvasObject) fyne.Size {
	contentMin := obj[2].MinSize()
	btnMin := obj[3].MinSize()

	width := fyne.Max(contentMin.Width, btnMin.Width) + padWidth
	height := contentMin.Height + btnMin.Height + theme.Padding() + padHeight*2

	return fyne.NewSize(width, height)
}
