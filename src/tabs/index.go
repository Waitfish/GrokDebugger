package tabs

import (
	"encoding/json"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/vjeantet/grok"
)

func RunIndexTab(app fyne.App, myWin fyne.Window) *container.TabItem {
	g, _ := grok.New()

	// new Vbox

	// define binding variable

	bindStr := binding.NewString()
	bindNameOnly := binding.NewBool()
	bindNameOnly.Set(true)
	grokExpTextWidget := widget.NewMultiLineEntry()
	sampleDataTextWidget := widget.NewMultiLineEntry()
	outputTextWidget := widget.NewLabelWithData(bindStr)
	nameCaptureOnleyBoolWidget := widget.NewCheckWithData("Named Captures Only", bindNameOnly)
	nameCaptureOnleyBoolWidget.OnChanged = func(b bool) {
		g, _ = grok.NewWithConfig(&grok.Config{NamedCapturesOnly: b})
	}

	form := widget.NewForm(
		&widget.FormItem{Text: "Grok:", Widget: grokExpTextWidget},
		&widget.FormItem{Text: "Sample Data:", Widget: sampleDataTextWidget},
		&widget.FormItem{Text: "Output:", Widget: outputTextWidget},
		&widget.FormItem{Text: "", Widget: nameCaptureOnleyBoolWidget},
	)

	grokExpTextWidget.OnChanged = func(s string) {
		sampleDataList := strings.Split(sampleDataTextWidget.Text, "\n")
		var resJson []map[string]string
		for _, line := range sampleDataList {
			value, err := g.Parse(s, line)
			if err != nil {
				log.Print(err)
				resJson = append(resJson, map[string]string{"": ""})
			} else {
				resJson = append(resJson, value)
			}
		}
		b, _ := json.MarshalIndent(resJson, "", "    ")
		bindStr.Set(string(b))

	}

	sampleDataTextWidget.OnChanged = func(s string) {
		sampleDataList := strings.Split(s, "\n")
		var resJson []map[string]string
		for _, line := range sampleDataList {
			value, err := g.Parse(grokExpTextWidget.Text, line)
			if err != nil {
				log.Print(err)
				resJson = append(resJson, map[string]string{"": ""})
			} else {
				resJson = append(resJson, value)
			}
		}
		b, _ := json.MarshalIndent(resJson, "", "    ")
		bindStr.Set(string(b))
	}

	co := container.NewGridWithColumns(1, form)
	indexTabItem := container.NewTabItem("Run", co)
	return indexTabItem
}
