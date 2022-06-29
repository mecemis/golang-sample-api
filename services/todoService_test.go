package services

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-sample-api/mocks/repositories"
	"golang-sample-api/models"
	"testing"
)

var mockRepo *repositories.MockTodoRepository
var service TodoService

var FakeData = []models.Todo{
	{primitive.NewObjectID(), "title1", "content1"},
	{primitive.NewObjectID(), "title2", "content2"},
	{primitive.NewObjectID(), "title3", "content3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repositories.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepo)
	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultTodoService_TodoGetAll(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(FakeData, nil)
	result, err := service.TodoGetAll()

	assert.NotEmpty(t, result)

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}