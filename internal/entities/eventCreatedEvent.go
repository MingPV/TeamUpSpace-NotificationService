package entities

type EventCreatedEvent struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	EventName       string `json:"event_name"`
	StartAt         string `json:"start_at"`
	EndAt           string `json:"end_at"`
	MainImageUrl    string `json:"main_image_url"`
	RegisterStartDt string `json:"register_start_dt"`
	RegisterCloseDt string `json:"register_close_dt"`
}
