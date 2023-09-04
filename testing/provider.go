package testing

import "fmt"

type SimpleRepository struct {
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

func (repository *SimpleRepository) Create() {
	fmt.Println("create")
}

type SimpleService struct {
	SimpleRepository *SimpleRepository
}

func NewSimpleService(simpleRepository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: simpleRepository,
	}
}

func (service *SimpleService) Save() {
	service.SimpleRepository.Create()
	fmt.Println("save")
}

type SimpleController struct {
	SimpleService *SimpleService
}

func NewSimpleController(simpleService *SimpleService) *SimpleController {
	return &SimpleController{
		SimpleService: simpleService,
	}
}

func (controller *SimpleController) PostUser() {
	controller.SimpleService.Save()
	fmt.Println("post")
}

type SimpleRoute struct {
	SimpleController *SimpleController
}

func NewSimpleRoute(simpleController *SimpleController) *SimpleRoute {
	return &SimpleRoute{
		SimpleController: simpleController,
	}
}

func (route *SimpleRoute) RouteUser() {
	// route get, post, update, delete
	route.SimpleController.PostUser()
	fmt.Println("some route")
}
