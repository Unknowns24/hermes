package mails

import (
	hermes "github.com/unknowns24/hermes/pkg/mails"
)

type InviteCode struct {
}

func (w *InviteCode) Name() string {
	return "invite_code"
}

func (w *InviteCode) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please copy your invite code:",
					InviteCode:   "123456",
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
}
