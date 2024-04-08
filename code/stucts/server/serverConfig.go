package server

type ServerConfig struct {
	Admins                    Admins                `json:"admins"`
	Bans                      Bans                  `json:"bans"`
	RepeatedChatMessages      []RepeatedChatMessage `json:"repeatedChatMessages"`
	ScheduledChatMessages     []interface{}         `json:"scheduledChatMessages"`
	ServerMessage             []string              `json:"serverMessage"`
	ChatMessagesUtcTime       bool                  `json:"chatMessagesUtcTime"`
	RepeatedChatMessagesCycle bool                  `json:"repeatedChatMessagesCycle"`
	StatsFileUpdateInterval   int                   `json:"statsFileUpdateIntervalSeconds"`
	BanReloadInterval         int                   `json:"banReloadIntervalMinutes"`
	StatsFileName             string                `json:"statsFileName"`
	StatsSaveConnectedPlayers bool                  `json:"statsSaveConnectedPlayers"`
	EventsApiToken            string                `json:"eventsApiToken"`
	EventsApiAddress          string                `json:"eventsApiAddress"`
	EventsApiRatelimit        int                   `json:"eventsApiRatelimitSeconds"`
	ServerMessageHeaderImage  string                `json:"serverMessageHeaderImage"`
	ServerMessageDiscordLink  string                `json:"serverMessageDiscordLink"`
	ServerMessageOpen         bool                  `json:"serverMessageOpen"`
	EventsApiEventsEnabled    []string              `json:"eventsApiEventsEnabled"`
}

type Admins map[string]string

type Bans map[string]string

type RepeatedChatMessage struct {
	Message         string `json:"message"`
	IntervalMinutes int    `json:"intervalMinutes"`
}
