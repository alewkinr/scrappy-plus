package main

import (
	"bytes"
	"context"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/alewkinr/scrappy-plus/internal"
	"github.com/alewkinr/scrappy-plus/pkg"
	"github.com/gocolly/colly"
)

const (
	ExitCodeOK  = iota // ExitCodeOK – код выхода при успешном завершении
	ExitCodeErr        // ExitCodeErr — код выхода при ошибке
)

// isButtonActive — проверяем что кнопка "Зарегистрироваться" активна
func isButtonActive(e *colly.HTMLElement) bool {
	applyBtnClasses := e.ChildAttr(".dacha-event-card__button-wrapper > .dacha-event-card__button", "class")
	return !strings.Contains(applyBtnClasses, "ui-button_disabled")
}

// parse — парсим URL и наполняем шаблон одержимым
func parse(settings *internal.Settings) (*internal.EventList, error) {
	c := colly.NewCollector()

	events := make([]internal.Event, 0)
	c.OnHTML(".dacha-event-card", func(e *colly.HTMLElement) {
		event := internal.Event{}
		event.Title = e.ChildText(".dacha-event-card__title > span")
		event.IsAvailable = isButtonActive(e)
		event.BookURL = settings.ParseURL
		events = append(events, event)
	})

	if visitURLErr := c.Visit(settings.ParseURL); visitURLErr != nil {
		return nil, visitURLErr
	}

	return &internal.EventList{Events: events}, nil
}

// run — запуск программы
func run(ctx context.Context) error {
	// defaults
	log := pkg.NewLogger()
	settings := internal.MustInitConfig()

	// senders
	sender := pkg.MustInitSender(settings.TelegramChannelID, settings.TelegramToken)

	// templates
	tmpl, parseTemplateErr := template.ParseFiles("./pkg/templates/message.tmpl")
	if parseTemplateErr != nil {
		log.Errorf("Ошибка парсинга шаблона сообщения %s", parseTemplateErr)
		return parseTemplateErr
	}

	eventList, parseEventsErr := parse(settings)
	if parseEventsErr != nil {
		log.Errorf("Ошибка при парсинге URL %s: %s", settings.ParseURL, parseEventsErr)
		return parseEventsErr
	}

	buf := new(bytes.Buffer)
	if executeTmplErr := tmpl.Execute(buf, eventList); executeTmplErr != nil {
		log.Errorf("Ошибка наполнения шаблона данными: %s", executeTmplErr)
		return executeTmplErr
	}

	if sendMsgErr := sender.Send(ctx, buf.String()); sendMsgErr != nil {
		log.Errorf("Ошибка при отправке сообщения: %s", sendMsgErr)
		return sendMsgErr
	}

	return nil
}

// main — главная функция
func main() {
	ctx, cFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cFunc()

	if runErr := run(ctx); runErr != nil {
		// nolint:gocritic
		os.Exit(ExitCodeErr)
	}

	os.Exit(ExitCodeOK)
}
