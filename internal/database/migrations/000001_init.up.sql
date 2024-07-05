create table produtos (
	pro_codigo serial primary key,
	pro_descricao varchar not null,
	pro_preco numeric(10, 2) not null
)