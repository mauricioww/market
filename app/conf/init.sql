CREATE TABLE IF NOT EXISTS PRODUCTS(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name char(60),
    supplier_id INT,
    category_id INT,
    units_in_stock INT,
    unit_price FLOAT,
    discontinued BOOLEAN
);