package requests

type SignupReq struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	ChessComUsername string `json:"chess_com_username"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateChessComUsernameReq struct {
	ChessComUsername string `json:"chess_com_username"`
}

type ChangePasswordReq struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type Logout struct {
}
