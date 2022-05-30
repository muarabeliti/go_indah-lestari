CREATE TABLE `product` (
  `product_type` integer PRIMARY KEY,
  `product_description` varchar(50),
  `operator` integer,
  `payment_method` bool
);

CREATE TABLE `data_pembelian` (
  `nama` varchar(50),
  `alamat` text,
  `tanggal_lahir` integer,
  `status_user` varchar(255),
  `gender` bool,
  `created_at` text,
  `created_up` text
);

ALTER TABLE `product` ADD FOREIGN KEY (`operator`) REFERENCES `product` (`product_type`);

ALTER TABLE `data_pembelian` ADD FOREIGN KEY (`created_at`) REFERENCES `product` (`product_description`);
