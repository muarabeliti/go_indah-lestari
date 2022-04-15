CREATE TABLE IF NOT EXISTS pivot(
    id INT NOT NULL AUTO_INCREMENT,
    id_product INT NOT NULL,
    id_log_product INT NOT NULL,
    qty INT NOT NULL,
    price DECIMAL(9,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY ( id ),
    FOREIGN KEY (id_product) REFERENCES product(id),
    FOREIGN KEY (id_log_product) REFERENCES log_product(id)
);