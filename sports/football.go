package sports

import (
	"fmt"
	"github.com/google/uuid"
	"syreclabs.com/go/faker"
	"time"
)

type Match struct {
	ID          uuid.UUID `json:"id"`
	HomeTeam    string    `json:"home_team"`
	AwayTeam    string    `json:"away_team"`
	Stadium     string    `json:"stadium"`
	Round       string    `json:"round"`
	Competition string    `json:"competition"`
	KickOff     time.Time `json:"kick_off"`
	Note        string    `json:"note"`
}

func NewMatch() *Match {
	return &Match{
		ID:       uuid.New(),
		HomeTeam: faker.Team().Name(),
		AwayTeam: faker.Team().Name(),
		Stadium:  fmt.Sprintf("%s Arena", faker.Address().City()),
		KickOff:  faker.Time().Forward(7 * 24 * time.Hour),
		Note:     faker.Lorem().Paragraph(5),
	}
}
