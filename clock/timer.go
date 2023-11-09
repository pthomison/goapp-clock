package clock

import (
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sirupsen/logrus"
)

type Clock struct {
	app.Compo

	time string
	date string
}

func (c *Clock) RefreshClock(ctx app.Context) {
	ctx.Dispatch(func(_ app.Context) {
		now := time.Now()
		c.time = now.Format("3:04:05.000pm (MST)")
		c.date = now.Format("Jan 2, 2006")
	})
}

func (c *Clock) OnMount(ctx app.Context) {
	c.RefreshClock(ctx)

	timerChan := time.NewTicker(1 * time.Millisecond)

	go func() {
		for _ = range timerChan.C {
			c.RefreshClock(ctx)
		}
	}()

	logrus.Info("complete")
}

func (c *Clock) Render() app.UI {
	timeView := app.H1().
		Class("has-text-grey-lighter").
		Class("title").
		Text(c.time)

	dateView := app.H1().
		Class("has-text-grey-lighter").
		Class("title").
		Text(c.date)

	return app.Div().Body(
		timeView,
		dateView,
	)
}
