package orders

type Service interface {
	FindAll() ([]Order, error)
	FindByID(ID FindOrderInput) (Order, error)
	Save(input SaveOrderInput) (Order, error)
	Update(ID FindOrderInput, input UpdateOrderInput) (Order, error)
	Delete(ID FindOrderInput) (Order, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Order, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *service) FindByID(ID FindOrderInput) (Order, error) {
	order, err := s.repository.FindByID(ID.ID)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (s *service) Save(input SaveOrderInput) (Order, error) {
	var order Order
	order.CustomerName = input.CustomerName
	for _, item := range input.Items {
		order.Items = append(order.Items, Item{
			Code:        item.Code,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	newOrder, err := s.repository.Save(order)
	if err != nil {
		return newOrder, err
	}
	return newOrder, nil
}

func (s *service) Update(ID FindOrderInput, input UpdateOrderInput) (Order, error) {
	order, err := s.repository.FindByID(ID.ID)
	if err != nil {
		return order, err
	}

	order.CustomerName = input.CustomerName

	var orderItems []Item
	for _, item := range input.Items {
		orderItems = append(orderItems, Item{
			ID:          item.ID,
			Code:        item.Code,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	order.Items = orderItems

	newOrder, err := s.repository.Update(order)
	if err != nil {
		return newOrder, err
	}
	return newOrder, nil
}

func (s *service) Delete(ID FindOrderInput) (Order, error) {
	order, err := s.repository.FindByID(ID.ID)
	if err != nil {
		return order, err
	}

	newOrder, err := s.repository.Delete(order)
	if err != nil {
		return newOrder, err
	}
	return newOrder, nil
}
