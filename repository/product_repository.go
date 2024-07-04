package repository

import (
	"api-curriculos/model"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type ProductRepository struct {
	connection *sql.DB
}

type ProductDB struct {
	Id           int
	Product_name string
	Price        float64
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query, args, err := sq.Select("pro_codigo", "pro_descricao", "pro_preco").From("produtos").ToSql()
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	rows, err := pr.connection.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.Id,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(p model.Product) (int, error) {
	var id int
	query, args, err := sq.
		Insert("produtos").
		Columns("pro_descricao", "pro_preco").
		Values(p.Name, p.Price).
		Suffix("returning pro_codigo").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = pr.connection.QueryRow(query, args...).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, args, err := sq.
		Select("pro_codigo", "pro_descricao", "pro_preco").
		From("produtos").
		Where(sq.Eq{"pro_codigo": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product
	err = pr.connection.QueryRow(query, args...).Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Retorna nil se não houver linhas encontradas
		}
		fmt.Println(err)
		return nil, err
	}

	return &product, nil
}
