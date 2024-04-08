package sports

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/maps"
	"syreclabs.com/go/faker"
)

const (
	TopicNewFootballMatch = "football-match-new"
)

type FootballMatch struct {
	ID          uuid.UUID     `json:"id"`
	HomeTeam    *FootballTeam `json:"home_team"`
	AwayTeam    *FootballTeam `json:"away_team"`
	Stadium     string        `json:"stadium"`
	Round       int           `json:"round"`
	Competition string        `json:"competition"`
	Country     string        `json:"country"`
	KickOff     time.Time     `json:"kick_off"`
}

type FootballTeam struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Stadium string    `json:"stadium"`
}

func NewFootballMatch() *FootballMatch {
	competition := getRandomElement(footballLeagues)

	teams := teamsByLeague[competition]

	homeTeam := getRandomElement(teams)
	awayTeam := getRandomElement(teams)

	for awayTeam.ID == homeTeam.ID {
		awayTeam = getRandomElement(teams)
	}

	return &FootballMatch{
		ID:          uuid.New(),
		HomeTeam:    homeTeam,
		AwayTeam:    awayTeam,
		Stadium:     homeTeam.Stadium,
		Round:       getRandomRoundNumber(len(teams)),
		Competition: competition,
		Country:     countryByLeague[competition],
		KickOff:     faker.Time().Forward(7 * 24 * time.Hour),
	}
}

func getRandomElement[T any](arr []T) T {
	// Create a new rand
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random index within the range of the array length
	randomIndex := r.Intn(len(arr))

	// Return the element at the randomly chosen index
	return arr[randomIndex]
}

func getRandomRoundNumber(numberOfTeams int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return 1 + r.Intn(numberOfTeams*2)
}

var (
	teamsByLeague = map[string][]*FootballTeam{
		"Premier League":   premierLeagueTeams,
		"La Liga":          laLigaTeams,
		"Serie A":          serieATeams,
		"Bundesliga":       bundesligaTeams,
		"Ligue 1":          ligue1Teams,
		"EPL Championship": eplChampionshipTeams,
		"La Liga 2":        laLiga2Teams,
		"Serie B":          serieBTeams,
		"Bundesliga 2":     bundesliga2Teams,
		"Ligue 2":          ligue2Teams,
	}

	countryByLeague = map[string]string{
		"Premier League":   "England",
		"La Liga":          "Spain",
		"Serie A":          "Italy",
		"Bundesliga":       "Germany",
		"Ligue 1":          "France",
		"EPL Championship": "England",
		"La Liga 2":        "Spain",
		"Serie B":          "Italy",
		"Bundesliga 2":     "Germany",
		"Ligue 2":          "France",
	}

	footballLeagues = maps.Keys(teamsByLeague)

	premierLeagueTeams = []*FootballTeam{
		{uuid.New(), "Arsenal F.C.", "Emirates Stadium"},
		{uuid.New(), "Aston Villa F.C.", "Villa Park"},
		{uuid.New(), "Brentford F.C.", "Brentford Community Stadium"},
		{uuid.New(), "Brighton & Hove Albion F.C.", "American Express Community Stadium"},
		{uuid.New(), "Burnley F.C.", "Turf Moor"},
		{uuid.New(), "Chelsea F.C.", "Stamford Bridge"},
		{uuid.New(), "Crystal Palace F.C.", "Selhurst Park Stadium"},
		{uuid.New(), "Everton F.C.", "Goodison Park"},
		{uuid.New(), "Fulham F.C.", "Craven Cottage"},
		{uuid.New(), "Leeds United F.C.", "Elland Road"},
		{uuid.New(), "Leicester City F.C.", "King Power Stadium"},
		{uuid.New(), "Liverpool F.C.", "Anfield"},
		{uuid.New(), "Manchester City F.C.", "Etihad Stadium"},
		{uuid.New(), "Manchester United F.C.", "Old Trafford"},
		{uuid.New(), "Newcastle United F.C.", "St James' Park"},
		{uuid.New(), "Nottingham Forest F.C.", "City Ground"},
		{uuid.New(), "Sheffield United F.C.", "Bramall Lane"},
		{uuid.New(), "Tottenham Hotspur F.C.", "Tottenham Hotspur Stadium"},
		{uuid.New(), "West Ham United F.C.", "London Stadium"},
		{uuid.New(), "Wolverhampton Wanderers F.C.", "Molineux Stadium"},
	}

	laLigaTeams = []*FootballTeam{
		{uuid.New(), "Athletic Bilbao", "San Mamés"},
		{uuid.New(), "Atlético Madrid", "Wanda Metropolitano"},
		{uuid.New(), "Barcelona", "Camp Nou"},
		{uuid.New(), "Celta Vigo", "Abanca-Balaídos"},
		{uuid.New(), "Elche CF", "Martínez Valero"},
		{uuid.New(), "Espanyol", "RCDE Stadium"},
		{uuid.New(), "Getafe CF", "Coliseum Alfonso Pérez"},
		{uuid.New(), "Granada CF", "Nuevo Los Cármenes"},
		{uuid.New(), "Levante UD", "Ciutat de València"},
		{uuid.New(), "Mallorca", "Son Moix"},
		{uuid.New(), "Osasuna", "El Sadar"},
		{uuid.New(), "Real Betis", "Benito Villamarín"},
		{uuid.New(), "Real Madrid", "Santiago Bernabéu"},
		{uuid.New(), "Real Sociedad", "Reale Arena"},
		{uuid.New(), "Sevilla FC", "Ramón Sánchez Pizjuán"},
		{uuid.New(), "Valencia CF", "Mestalla"},
		{uuid.New(), "Villarreal CF", "Estadio de la Cerámica"},
	}

	serieATeams = []*FootballTeam{
		{uuid.New(), "AC Milan", "San Siro"},
		{uuid.New(), "Atalanta BC", "Gewiss Stadium"},
		{uuid.New(), "Bologna FC 1909", "Renato Dall'Ara Stadium"},
		{uuid.New(), "Cagliari Calcio", "Sardegna Arena"},
		{uuid.New(), "Empoli F.C.", "Carlo Castellani Stadium"},
		{uuid.New(), "FC Internazionale Milano", "San Siro"},
		{uuid.New(), "ACF Fiorentina", "Artemio Franchi Stadium"},
		{uuid.New(), "Frosinone Calcio", "Stadio Benito Stirpe"},
		{uuid.New(), "Genoa CFC", "Luigi Ferraris Stadium"},
		{uuid.New(), "Hellas Verona FC", "Marcantonio Bentegodi Stadium"},
		{uuid.New(), "Juventus FC", "Allianz Stadium"},
		{uuid.New(), "S.S. Lazio", "Stadio Olimpico"},
		{uuid.New(), "US Lecce", "Stadio Ettore Giardiniero - Via del Mare"},
		{uuid.New(), "AC Monza", "Brianteo Stadium"},
		{uuid.New(), "SSC Napoli", "Stadio Diego Armando Maradona"},
		{uuid.New(), "Salernitana 1919", "Arechi Stadium"},
		{uuid.New(), "Sassuolo Calcio", "Mapei Stadium - Città del Tricolore"},
		{uuid.New(), "Torino FC", "Olympic Grande Torino Stadium"},
		{uuid.New(), "Udinese Calcio", "Stadio Friuli"},
	}

	bundesligaTeams = []*FootballTeam{
		{uuid.New(), "FC Bayern Munich", "Allianz Arena"},
		{uuid.New(), "Borussia Dortmund", "Signal Iduna Park"},
		{uuid.New(), "RB Leipzig", "Red Bull Arena"},
		{uuid.New(), "Borussia Mönchengladbach", "Borussia-Park"},
		{uuid.New(), "VfL Wolfsburg", "Volkswagen Arena"},
		{uuid.New(), "Eintracht Frankfurt", "Deutsche Bank Park"},
		{uuid.New(), "Bayer 04 Leverkusen", "BayArena"},
		{uuid.New(), "FC Union Berlin", "Stadion An der Alten Försterei"},
		{uuid.New(), "SC Freiburg", "Schwarzwald-Stadion"},
		{uuid.New(), "TSG 1899 Hoffenheim", "PreZero Arena"},
		{uuid.New(), "FC Köln", "RheinEnergieStadion"},
		{uuid.New(), "Hertha BSC", "Olympiastadion"},
		{uuid.New(), "FSV Mainz 05", "Opel Arena"},
		{uuid.New(), "Arminia Bielefeld", "SchücoArena"},
		{uuid.New(), "FC Augsburg", "WWK Arena"},
		{uuid.New(), "SV Werder Bremen", "Weserstadion"},
		{uuid.New(), "FC Schalke 04", "VELTINS-Arena"},
	}

	ligue1Teams = []*FootballTeam{
		{uuid.New(), "AC Ajaccio", "Stade François Coty"},
		{uuid.New(), "Amiens SC", "Stade Crédit Agricole de la Licorne"},
		{uuid.New(), "AS Nancy", "Stade Marcel Picot"},
		{uuid.New(), "Clermont Foot", "Stade Gabriel Montpied"},
		{uuid.New(), "Dijon FCO", "Stade Gaston Gérard"},
		{uuid.New(), "EA Guingamp", "Stade de Roudourou"},
		{uuid.New(), "En Avant Troyes", "Stade de l'Aube"},
		{uuid.New(), "FC Bastia-Borgo", "Stade Armand Cesari"},
		{uuid.New(), "FC Chambly", "Stade Pierre Brisson"},
		{uuid.New(), "Grenoble Foot 38", "Stade des Alpes"},
		{uuid.New(), "Le Havre AC", "Stade Océane"},
		{uuid.New(), "Nîmes Olympique", "Stade des Costières"},
		{uuid.New(), "Paris FC", "Stade Charléty"},
		{uuid.New(), "Paris Saint-Germain", "Parc des Princes"},
		{uuid.New(), "Pau FC", "Stade du Hameau"},
		{uuid.New(), "Quevilly-Rouen Métropole", "Stade Robert Diochon"},
		{uuid.New(), "Rodez AF", "Stade Paul Lignon"},
		{uuid.New(), "SM Caen", "Stade Michel d'Ornano"},
		{uuid.New(), "Toulouse FC", "Stadium de Toulouse"},
		{uuid.New(), "USL Dunkerque", "Stade Marcel-Tribut"},
		{uuid.New(), "Valenciennes FC", "Stade du Hainaut"},
	}

	laLiga2Teams = []*FootballTeam{
		{uuid.New(), "CD Leganés", "Estadio Municipal de Butarque"},
		{uuid.New(), "SD Eibar", "Ipurua"},
		{uuid.New(), "RCD Espanyol de Barcelona (Espanyol)", "RCDE Stadium"},
		{uuid.New(), "Racing de Santander", "El Sardinero"},
		{uuid.New(), "Elche CF", "Estadio Manuel Martínez Valero"},
		{uuid.New(), "Real Valladolid CF (Real Valladolid)", "Estadio Nuevo José Zorrilla"},
		{uuid.New(), "Real Oviedo", "Estadio Carlos Tartiere"},
		{uuid.New(), "Racing Club de Ferrol (Racing Ferrol)", "Estadio Municipal de A Malata"},
		{uuid.New(), "Burgos CF", "El Plantío"},
		{uuid.New(), "Sporting de Gijón", "El Molinón"},
		{uuid.New(), "Levante UD", "Estadi Ciutat de València"},
		{uuid.New(), "CD Tenerife", "Estadio Heliodoro Rodríguez López"},
		{uuid.New(), "CD Eldense", "Nuevo Pepico Amat"},
		{uuid.New(), "SD Huesca", "El Alcoraz"},
		{uuid.New(), "Real Zaragoza", "Estadio La Romareda"},
		{uuid.New(), "FC Cartagena", "Estadio Cartagonova"},
		{uuid.New(), "CD Mirandés", "Estadio Municipal de Anduva"},
		{uuid.New(), "AD Alcorcón", "Santo Domingo Municipal Stadium"},
		{uuid.New(), "Albacete Balompié", "Estadio Carlos Belmonte"},
		{uuid.New(), "FC Andorra", "Estadi Nacional"},
		{uuid.New(), "SD Amorebieta", "Instalaciones de Lezama"},
		{uuid.New(), "Villarreal CF B", "Ciudad Deportiva de Villarreal"},
	}

	eplChampionshipTeams = []*FootballTeam{
		{uuid.New(), "Birmingham City", "St Andrew's Stadium"},
		{uuid.New(), "Blackburn Rovers", "Ewood Park"},
		{uuid.New(), "Blackpool", "Bloomfield Road"},
		{uuid.New(), "Bristol City", "Ashton Gate Stadium"},
		{uuid.New(), "Burnley", "Turf Moor"},
		{uuid.New(), "Cardiff City", "Cardiff City Stadium"},
		{uuid.New(), "Coventry City", "Coventry Building Society Arena"},
		{uuid.New(), "Huddersfield Town", "John Smith's Stadium"},
		{uuid.New(), "Hull City", "MKM Stadium"},
		{uuid.New(), "Luton Town", "Kenilworth Road"},
		{uuid.New(), "Middlesbrough", "Riverside Stadium"},
		{uuid.New(), "Millwall", "The Den"},
		{uuid.New(), "Norwich City", "Carrow Road"},
		{uuid.New(), "Preston North End", "Deepdale"},
		{uuid.New(), "Queens Park Rangers", "Loftus Road"},
		{uuid.New(), "Rotherham United", "New York Stadium"},
		{uuid.New(), "Sheffield United", "Bramall Lane"},
		{uuid.New(), "Stoke City", "bet365 Stadium"},
		{uuid.New(), "Sunderland", "Stadium of Light"},
		{uuid.New(), "Swansea City", "Swansea.com Stadium"},
		{uuid.New(), "Watford", "Vicarage Road"},
		{uuid.New(), "West Bromwich Albion", "The Hawthorns"},
	}

	serieBTeams = []*FootballTeam{
		{uuid.New(), "AS Cittadella", "Stadio Pier Cesare Tombolato"},
		{uuid.New(), "Benevento", "Stadio Ciro Vigorito"},
		{uuid.New(), "Brescia", "Stadio Mario Rigamonti"},
		{uuid.New(), "Cesena", "Stadio Dino Manuzzi"},
		{uuid.New(), "Città di Fasano", "Stadio Comunale (Fasano)"},
		{uuid.New(), "Cremonese", "Stadio Giovanni Zini"},
		{uuid.New(), "Crotone", "Stadio Ezio Scida"},
		{uuid.New(), "Feralpisalò", "Stadio Lino Turina"},
		{uuid.New(), "L.R. Vicenza Virtus", "Stadio Romeo Menti"},
		{uuid.New(), "Monza", "Stadio Brianteo"},
		{uuid.New(), "Novara", "Stadio Silvio Piola"},
		{uuid.New(), "Perugia", "Stadio Renato Curi"},
		{uuid.New(), "Pordenone", "Stadio Guido Teghil"},
		{uuid.New(), "Reggiana", "Stadio Città del Tricolore"},
		{uuid.New(), "Salernitana", "Stadio Arechi"},
		{uuid.New(), "SPAL", "Stadio Paolo Mazza"},
		{uuid.New(), "Ternana", "Stadio Libero Liberati"},
		{uuid.New(), "Venezia", "Stadio Pierluigi Penzo"},
		{uuid.New(), "Vicenza", "Stadio Romeo Menti"},
		{uuid.New(), "Virtus Entella", "Stadio Comunale di Chiavari"},
	}

	bundesliga2Teams = []*FootballTeam{
		{uuid.New(), "FC Heidenheim 1846", "Voith-Arena"},
		{uuid.New(), "FC Kaiserslautern", "Fritz-Walter-Stadion"},
		{uuid.New(), "FC Nürnberg", "Max-Morlock-Stadion"},
		{uuid.New(), "Eintracht Braunschweig", "Eintracht-Stadion"},
		{uuid.New(), "FC Erzgebirge Aue", "Erzgebirgsstadion"},
		{uuid.New(), "FC Ingolstadt 04", "Audi Sportpark"},
		{uuid.New(), "FC St. Pauli", "Millerntor-Stadion"},
		{uuid.New(), "FSV Zwickau", "Stadion Zwickau"},
		{uuid.New(), "Hansa Rostock", "DKB-Arena"},
		{uuid.New(), "Karlsruher SC", "Wildparkstadion"},
		{uuid.New(), "SC Paderborn 07", "Benteler-Arena"},
		{uuid.New(), "SV Darmstadt 98", "Merck-Stadion am Böllenfalltor"},
		{uuid.New(), "SV Sandhausen", "BWT-Stadion am Hardtwald"},
		{uuid.New(), "SV Wehen Wiesbaden", "BRITA-Arena"},
		{uuid.New(), "Türkgücü München", "Städtisches Stadion an der Grünwalder Straße"},
		{uuid.New(), "VfL Osnabrück", "Bremer Brücke"},
		{uuid.New(), "VfL Bochum 1848", "Vonovia Ruhrstadion"},
	}

	ligue2Teams = []*FootballTeam{
		{uuid.New(), "AJ Auxerre", "Stade de l'Abbé-Deschamps"},
		{uuid.New(), "Amiens SC", "Stade Crédit Agricole de la Licorne"},
		{uuid.New(), "AS Nancy", "Stade Marcel Picot"},
		{uuid.New(), "Clermont Foot", "Stade Gabriel Montpied"},
		{uuid.New(), "Dijon FCO", "Stade Gaston Gérard"},
		{uuid.New(), "EA Guingamp", "Stade de Roudourou"},
		{uuid.New(), "En Avant Troyes", "Stade de l'Aube"},
		{uuid.New(), "Grenoble Foot 38", "Stade des Alpes"},
		{uuid.New(), "Havre AC", "Stade Océane"},
		{uuid.New(), "Nîmes Olympique", "Stade des Costières"},
		{uuid.New(), "Paris FC", "Stade Charléty"},
		{uuid.New(), "Pau FC", "Stade du Hameau"},
		{uuid.New(), "Quevilly-Rouen Métropole", "Stade Robert Diochon"},
		{uuid.New(), "Rodez AF", "Stade Paul Lignon"},
		{uuid.New(), "SM Caen", "Stade Michel d'Ornano"},
		{uuid.New(), "Toulouse FC", "Stadium de Toulouse"},
		{uuid.New(), "USL Dunkerque", "Stade Marcel-Tribut"},
		{uuid.New(), "Valenciennes FC", "Stade du Hainaut"},
	}
)
