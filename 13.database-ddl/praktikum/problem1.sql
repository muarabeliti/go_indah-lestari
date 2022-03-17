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
  `sattus_user` varchar(255),
  `gender` bool,
  `created_at` text,
  `created_up` text
);
