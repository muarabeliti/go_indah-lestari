-- phpMyAdmin SQL Dump
-- version 4.5.1
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: 19 Mar 2022 pada 17.56
-- Versi Server: 10.1.19-MariaDB
-- PHP Version: 5.6.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `alta_online_shop`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `operators`
--

CREATE TABLE `operators` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `status` tinyint(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `operators`
--

INSERT INTO `operators` (`id`, `name`, `created_at`, `updated_at`, `status`) VALUES
(9, 'alta', '2022-03-20 07:25:00', '0000-00-00 00:00:00', 1),
(9, 'alta', '2022-03-20 07:25:00', '0000-00-00 00:00:00', 1),
(3, 'people', '2022-03-13 02:00:00', '0000-00-00 00:00:00', 1),
(3, 'people', '2022-03-13 02:00:00', '0000-00-00 00:00:00', 1),
(3, 'people', '2022-03-19 16:44:42', '0000-00-00 00:00:00', 1),
(3, 'people', '2022-03-19 16:44:47', '0000-00-00 00:00:00', 1),
(1, 'who', '2022-03-19 16:45:29', '0000-00-00 00:00:00', 1),
(1, 'who', '2022-03-19 16:45:36', '0000-00-00 00:00:00', 1),
(4, 'halim', '2022-03-15 03:30:00', '0000-00-00 00:00:00', 0),
(4, 'halim', '2022-03-15 03:30:00', '0000-00-00 00:00:00', 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `payment_method`
--

CREATE TABLE `payment_method` (
  `PK id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `product`
--

CREATE TABLE `product` (
  `PK id` int(11) NOT NULL,
  `product_type_id` int(11) NOT NULL,
  `operator_id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `product`
--

INSERT INTO `product` (`PK id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `update_at`) VALUES
(1, 1, 1, '12443', 'halim', 4, '2022-03-17 02:15:00', '0000-00-00 00:00:00'),
(1, 1, 1, '12443', 'halim', 4, '2022-03-17 02:15:00', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `product_description`
--

CREATE TABLE `product_description` (
  `PK id` int(11) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `product_type`
--

CREATE TABLE `product_type` (
  `id` int(11) NOT NULL,
  `name` varchar(60) NOT NULL,
  `status` tinyint(4) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `product_type`
--

INSERT INTO `product_type` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
(0, 'alta', 0, '2022-03-19 16:13:59', '0000-00-00 00:00:00'),
(1, 'buku', 0, '2022-03-19 16:18:56', '0000-00-00 00:00:00'),
(2, 'paperbag', 0, '2022-03-19 16:18:56', '0000-00-00 00:00:00'),
(2, 'pulpen', 0, '2022-03-20 14:35:00', '0000-00-00 00:00:00'),
(2, 'pulpen', 0, '2022-03-20 14:35:00', '0000-00-00 00:00:00'),
(3, 'tas', 3, '2022-03-14 23:00:00', '0000-00-00 00:00:00'),
(3, 'tas', 3, '2022-03-14 23:00:00', '0000-00-00 00:00:00'),
(3, 'tas', 3, '2022-03-20 03:27:00', '0000-00-00 00:00:00'),
(3, 'tas', 3, '2022-03-20 03:27:00', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `pk id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `payment_method_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL,
  `total_qty` int(1) NOT NULL,
  `total_price` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaction_details`
--

CREATE TABLE `transaction_details` (
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL,
  `qty` int(11) NOT NULL,
  `price` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `PK id` int(11) NOT NULL,
  `status` smallint(6) NOT NULL,
  `dob` date NOT NULL,
  `gender` char(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `user`
--

INSERT INTO `user` (`PK id`, `status`, `dob`, `gender`, `created_at`, `update_at`) VALUES
(1, 1, '2022-03-07', 'm', '2022-03-22 01:31:00', '0000-00-00 00:00:00'),
(1, 1, '2022-03-07', 'm', '2022-03-22 01:31:00', '0000-00-00 00:00:00'),
(2, 4, '2022-03-23', 'p', '2022-03-14 17:00:00', '0000-00-00 00:00:00'),
(2, 4, '2022-03-23', 'p', '2022-03-14 17:00:00', '0000-00-00 00:00:00'),
(3, 5, '2022-03-23', 'p', '2022-03-17 03:00:19', '0000-00-00 00:00:00'),
(3, 5, '2022-03-23', 'p', '2022-03-17 03:00:19', '0000-00-00 00:00:00');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
