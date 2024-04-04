package auth

type AuthRepository interface {
	ValidateUser(string) (bool, error)
}

// бизнес логика авторизации будет дергать метод авторизации
type AuthService struct {
	AuthRepository
}

func NewAuthService(repo AuthRepository) *AuthService {
	service := &AuthService{
		AuthRepository: repo,
	}

	return service
}

func (ase *AuthService) AuthUser(token string) (bool, error) {
	ok, err := ase.AuthRepository.ValidateUser(token)
	if err != nil {
		return ok, err
	}
	//в куки контроллер сохранит результат
	//а тута просто псевдовалидация с заделом на будущее
	return ok, nil
}
