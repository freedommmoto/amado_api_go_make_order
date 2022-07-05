-- name: MakeNewProduct :one
INSERT INTO product ("name", stock, price)
VALUES ($1, $2 , $3) RETURNING *;

-- name: UpdateProduct :one
UPDATE product SET name=$2 , stock=$3 , price=$4 , active=$5 , modified=now()
               WHERE id_product = $1 RETURNING *;

-- name: UpdateProductStock :one
UPDATE product SET stock=$2 , modified=now() WHERE id_product = $1 RETURNING *;

-- name: DeleteProduct :one
UPDATE product SET deleted=now() WHERE id_product = $1 RETURNING *;

-- name: ListProduct :many
SELECT *
FROM product
ORDER BY id_product LIMIT $1
OFFSET $2;

-- name: getProduct :one
SELECT *
FROM product
WHERE id_product = $1;

-- name: getProductLowStock :one
SELECT *
FROM product
WHERE stock < $1;

