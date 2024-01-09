package profile

import pc "basic-coding-kulina/modules/usecase/user/profile"

type ProfileHandler struct {
	profileUsecase pc.ProfileUsecase
}

func New(profileUsecase pc.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		profileUsecase,
	}
}
