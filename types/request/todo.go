package api_request

type CreateTodoPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
	UserID      int    `json:"user_id"`
}

type UpdateTodoPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

type GetTodoQueryParams struct {
	Page        int  `json:"page" validate:"gt=0" default:"1"`
	Size        int  `json:"size" validate:"gt=0" default:"10"`
	ID          int  `json:"id"`
	UserID      int  `json:"user_id"`
	IsCompleted bool `json:"is_completed"`
}

type GetTodoByIDQueryParams struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}
