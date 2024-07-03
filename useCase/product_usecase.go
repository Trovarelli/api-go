package usecase

import (
	"api-curriculos/model"
	"api-curriculos/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{repository: repo}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) GetProductById(id int) (*model.Product, error) {
	return pu.repository.GetProductById(id)
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	id, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	//Como as informações não seráo alteradas pelo banco de dados (como por exemplo uma trigger criado_em) essa abordagem foi escolhida devido sua melhor perfomance e simplicidade
	return model.Product{
		Id:    id,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
