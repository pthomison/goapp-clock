package clock

import (
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	defaultTimeFormat = "3:04:05.00pm (MST)"
	defaultDateFormat = "Jan 2, 2006"
)

type Clock struct {
	app.Compo

	ticker *time.Ticker

	// rendered fields
	time         string
	date         string
	frequencyErr error

	timeFormat string
	dateFormat string

	frequency time.Duration
}

func NewClock(timeFormat string, dateFormat string, updateFreq time.Duration) *Clock {
	c := &Clock{
		timeFormat: timeFormat,
		dateFormat: dateFormat,
		frequency:  updateFreq,

		frequencyErr: nil,
	}

	c.ticker = time.NewTicker(updateFreq)

	return c
}

func (c *Clock) UpdateFrequency(ctx app.Context, freq time.Duration) {
	c.ticker.Stop()
	c.ticker = time.NewTicker(freq)

	go func() {
		for range c.ticker.C {
			c.RefreshClock(ctx)
		}
	}()
}

func (c *Clock) RefreshClock(ctx app.Context) {
	ctx.Dispatch(func(_ app.Context) {
		now := time.Now()
		c.time = now.Format(c.timeFormat)
		c.date = now.Format(c.dateFormat)
	})
}

func (c *Clock) OnMount(ctx app.Context) {
	c.RefreshClock(ctx)
	c.UpdateFrequency(ctx, c.frequency)
}

func (c *Clock) Render() app.UI {

	frequencyColor := "has-text-grey-lighter"
	if c.frequencyErr != nil {
		frequencyColor = "has-text-danger"
	}

	timeView := app.P().
		Class("has-text-grey-lighter").
		Class("subtitle is-spaced").
		Text(c.time)

	dateView := app.P().
		Class("has-text-grey-lighter").
		Class("title is-1").
		Text(c.date)

	optionsView := app.Div().Body(
		app.Form().Body(
			app.Label().
				Class("label").
				Class("has-text-grey-lighter").
				Text("Time Format"),
			app.Input().
				Class("input").
				OnInput(c.ValueTo(&c.timeFormat)).Value(c.timeFormat),

			app.Label().
				Class("label").
				Class("has-text-grey-lighter").
				Text("Date Format"),
			app.Input().
				Class("input").
				OnInput(c.ValueTo(&c.dateFormat)).Value(c.dateFormat),

			app.Label().
				Class("label").
				Class(frequencyColor).
				Text("Refresh Frequency"),
			app.Input().
				Class("input").
				OnInput(func(ctx app.Context, e app.Event) {
					newFreq := ctx.JSSrc().Get("value").String()
					h, err := time.ParseDuration(newFreq)
					c.frequencyErr = err
					if err == nil {
						c.UpdateFrequency(ctx, h)
					}

				}).Value(c.frequency),
		),
	)

	return app.Div().
		Class("level").
		Body(
			app.Div().Class("level-item has-text-centered").Body(
				app.Div().Body(
					timeView,
					dateView,
				),
			),
			app.Div().Class("level-item has-text-centered").Body(
				optionsView,
			),
		)
}
