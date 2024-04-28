package mails

import (
	hermes "github.com/unknowns24/hermes/pkg/mails"
)

type Maintenance struct {
}

func (w *Maintenance) Name() string {
	return "maintenance"
}

func (w *Maintenance) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			FreeMarkdown: `
> _Hermes_ service will shutdown the **1st August 2017** for maintenance operations. 

Services will be unavailable based on the following schedule:

| Services | Downtime |
| :------:| :-----------: |
| Service A | 2AM to 3AM |
| Service B | 4AM to 5AM |
| Service C | 5AM to 6AM |

Feel free to contact us for any question regarding this matter at [support@hermes-example.com](mailto:support@hermes-example.com) or in our [Gitter](https://gitter.im/)

`,
		},
	}
}
