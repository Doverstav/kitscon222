package view

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/doverstav/kitscon222/charm_demo/commands"
)

func AddPresentationUpdate(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyTab {
			// Some focus logic here
			currentFocus := m.PresentationInputFocus
			if currentFocus == 4 {
				return m, nil
			}
			m = blurIndex(m, currentFocus)
			m = focusIndex(m, currentFocus+1)
			m.PresentationInputFocus = currentFocus + 1

			return m, nil
		} else if msg.Type == tea.KeyShiftTab {
			currentFocus := m.PresentationInputFocus
			if currentFocus == 0 {
				return m, nil
			}
			m = blurIndex(m, currentFocus)
			m = focusIndex(m, currentFocus-1)
			m.PresentationInputFocus = currentFocus - 1

			return m, nil
		} else if msg.Type == tea.KeyEscape || msg.Type == tea.KeyEsc {
			m.CurrentView = PRESENTATION_LIST

			return m, nil
		} else if msg.String() == "ctrl+j" { // Ctrl + Enter reads as ctrl+j for some reason
			rating, _ := strconv.Atoi(m.PresentationRatingInput.Value())
			// Save new presentation
			return m, commands.SavePresentation(
				m.DB,
				m.SelectedKitscon.Id,
				m.PresentationTitleInput.Value(),
				m.PresentationPresenterInput.Value(),
				m.PresentationDescriptionInput.Value(),
				rating,
				m.PresentationReviewInput.Value(),
			)
		}

	}

	var titleCmd tea.Cmd
	m.PresentationTitleInput, titleCmd = m.PresentationTitleInput.Update(msg)
	var presenterCmd tea.Cmd
	m.PresentationPresenterInput, presenterCmd = m.PresentationPresenterInput.Update(msg)
	var descriptionCmd tea.Cmd
	m.PresentationDescriptionInput, descriptionCmd = m.PresentationDescriptionInput.Update(msg)
	var ratingCmd tea.Cmd
	m.PresentationRatingInput, ratingCmd = m.PresentationRatingInput.Update(msg)
	var reviewCmd tea.Cmd
	m.PresentationReviewInput, reviewCmd = m.PresentationReviewInput.Update(msg)

	return m, tea.Batch(titleCmd, presenterCmd, descriptionCmd, ratingCmd, reviewCmd)
}

func AddPresentationView(m Model) string {
	return docStyle.Render(
		inputStyle.Render(m.PresentationTitleInput.View()) + "\n" +
			inputStyle.Render(m.PresentationPresenterInput.View()) + "\n" +
			inputStyle.Render(m.PresentationDescriptionInput.View()) + "\n" +
			inputStyle.Render(m.PresentationRatingInput.View()) + "\n" +
			inputStyle.Render(m.PresentationReviewInput.View()),
	)
}

// ---------- HELPERS -------------

func blurIndex(m Model, index int) Model {
	if index == 0 {
		m.PresentationTitleInput.Blur()
		return m
	} else if index == 1 {
		m.PresentationPresenterInput.Blur()
		return m
	} else if index == 2 {
		m.PresentationDescriptionInput.Blur()
		return m
	} else if index == 3 {
		m.PresentationRatingInput.Blur()
		return m
	} else if index == 4 {
		m.PresentationReviewInput.Blur()
		return m
	} else {
		return m
	}
}

func focusIndex(m Model, index int) Model {
	if index == 0 {
		m.PresentationTitleInput.Focus()
		return m
	} else if index == 1 {
		m.PresentationPresenterInput.Focus()
		return m
	} else if index == 2 {
		m.PresentationDescriptionInput.Focus()
		return m
	} else if index == 3 {
		m.PresentationRatingInput.Focus()
		return m
	} else if index == 4 {
		m.PresentationReviewInput.Focus()
		return m
	} else {
		return m
	}
}
