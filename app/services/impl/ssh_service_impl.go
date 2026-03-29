package impl

import (
	"time"

	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/app/repositories"
	"github.com/create-go-app/fiber-go-template/pkg/httpctx"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SshService struct {
	repos repositories.Repositories
}

func NewSshService(repos repositories.Repositories) *SshService {
	return &SshService{repos: repos}
}

func (s *SshService) CreateSSHConnection(c *fiber.Ctx) error {
	ssh := &models.SSH{}

	if err := c.BodyParser(ssh); err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusBadRequest,
			true,
			err.Error(),
		)
	}

	validate := utils.NewValidator()

	ssh.ID = uuid.New()
	ssh.CreatedAt = time.Now()
	ssh.UserID = httpctx.UserID(c)

	if err := validate.Struct(ssh); err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusBadRequest,
			true,
			err.Error(),
		)
	}
	if err := s.repos.Ssh.CreateSSHConnection(ssh); err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusInternalServerError,
			true,
			err.Error(),
		)
	}

	return utils.NewResponse(
		c,
		fiber.StatusCreated,
		false,
		"Connection created successfully",
		fiber.Map{
			"ssh": ssh,
		},
	)
}

func (s *SshService) GetAllSSHConnectionsByUserID(c *fiber.Ctx) error {

	userId := httpctx.UserID(c)
	foundedSSHConnections, err := s.repos.Ssh.GetAllSSHConnectionsByUserId(userId)

	if err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusBadRequest,
			true,
			err.Error(),
		)
	}

	return utils.NewResponse(
		c,
		fiber.StatusOK,
		false,
		"List of connections",
		fiber.Map{
			"data": foundedSSHConnections,
		},
	)
}
