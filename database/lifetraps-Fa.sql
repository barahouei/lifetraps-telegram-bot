-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: May 15, 2023 at 08:31 PM
-- Server version: 10.11.2-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `lifetraps`
--

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `cid` int(11) NOT NULL,
  `category` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`cid`, `category`) VALUES
(1, 'رهاشدگی - Abandonment'),
(2, 'بی‌اعتمادی و سوءاستفاده - Mistrust and Abuse'),
(3, 'آسیب‌پذیری - Vulnerability'),
(4, 'وابستگی - Dependence'),
(5, 'محرومیت عاطفی - Emotional Deprivation'),
(6, 'انزوای اجتماعی - Social Exclusion'),
(7, 'نقص - Defectiveness'),
(8, 'شکست - Failure'),
(9, 'اطاعت - Subjugation'),
(10, 'معیارهای سخت‌گیرانه - Unrelenting Standards'),
(11, 'استحقاق - Entitlement');

-- --------------------------------------------------------

--
-- Table structure for table `questions`
--

CREATE TABLE `questions` (
  `qid` int(11) NOT NULL,
  `question` text NOT NULL,
  `cid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `questions`
--

INSERT INTO `questions` (`qid`, `question`, `cid`) VALUES
(1, '«وقتی بچه بودم» به نزدیکانم می‌چسبیدم چون می‌ترسیدم ترکم کنن.', 1),
(2, '«الان» به نزدیکانم می‌چسبم چون می‌ترسم ترکم کنن.', 1),
(3, '«وقتی بچه بودم» خیلی نگران بودم کسانی که دوستشون دارم یه نفر دیگه رو پیدا کنن که به من ترجیحش بدن و ترکم کنن.', 1),
(4, '«الان» خیلی نگرانم کسانی که دوستشون دارم یه نفر دیگه رو پیدا کنن که به من ترجیحش بدن و ترکم کنن.', 1),
(5, '«وقتی بچه بودم» معمولا دنبال انگیزه‌های پنهان آدم‌ها می‌گشتم؛ به راحتی به کسی اعتماد نمی‌کردم.', 2),
(6, '«الان» معمولا دنبال انگیزه‌های پنهان آدم‌ها می‌گردم؛ به راحتی به کسی اعتماد نمی‌کنم.', 2),
(7, '«وقتی بچه بودم» حس می‌کردم نمی‌تونم گاردم رو اطراف دیگران پایین بیارم وگرنه بهم صدمه می‌زنن.', 2),
(8, '«الان» حس می‌کنم نمی‌تونم گاردم رو اطراف دیگران پایین بیارم وگرنه بهم صدمه می‌زنن.', 2),
(9, '«وقتی بچه بودم» بیشتر از آدم‌های معمولی نگران خطر بودم، می‌ترسیدم مریض بشم یا آسیب ببینم.', 3),
(10, '«الان» بیشتر از آدم‌های معمولی نگران خطرم، می‌ترسم مریض بشم یا آسیب ببینم.', 3),
(11, '«وقتی بچه بودم» نگران بودم خودم یا خانواده‌م پول از دست بدیم و بی‌بضاعت و وابسته به دیگران بشیم.', 3),
(12, '«الان» نگرانم خودم یا خانواده‌م پول از دست بدیم و بی‌بضاعت و وابسته به دیگران بشیم.', 3),
(13, '«وقتی بچه بودم» حس می‌کردم نمی‌تونم تنهایی از پس خودم بر بیام و حس می‌کردم به کمک دیگران احتیاج دارم تا بتونم کارها رو انجام بدم.', 4),
(14, '«الان» حس می‌کنم نمی‌تونم تنهایی از پس خودم بر بیام و حس می‌کنم به کمک دیگران احتیاج دارم تا بتونم کارها رو انجام بدم.', 4),
(15, '«وقتی بچه بودم» من و والدینم بیش از حد تو زندگی و مشکلات همدیگه درگیر می‌شدیم.', 4),
(16, '«الان» من و والدینم بیش از حد تو زندگی و مشکلات همدیگه درگیر می‌شیم.', 4),
(17, '«وقتی بچه بودم» کسی رو نداشتم ازم مراقبت کنه، خودش رو با من به اشتراک بذاره یا عمیقا براش مهم باشه چه اتفاقی برای من می‌افته.', 5),
(18, '«الان» کسی رو ندارم ازم مراقبت کنه، خودش رو با من به اشتراک بذاره یا عمیقا براش مهم باشه چه اتفاقی برای من می‌افته', 5),
(19, '«وقتی بچه بودم» کسی نبود تا نیازهای عاطفیم برای درک، همدلی، راهنمایی، مشاوره و پشتیبانی رو برآورده کنه.', 5),
(20, '«الان» کسی نیست که نیازهای عاطفیم برای درک، همدلی، راهنمایی، مشاوره و پشتیبانی رو برآورده کنه.', 5),
(21, '«وقتی بچه بودم» حس می‌کردم به جایی تعلق ندارم. فرق می‌کردم. نمی‌تونستم جا بی‌افتم.', 6),
(22, '«الان» حس می‌کنم به جایی تعلق ندارم. فرق می‌کنم. نمی‌تونم جا بی‌افتم.', 6),
(23, '«وقتی بچه بودم» خیلی گرفته و خسته‌کننده بودم. نمی‌دونستم برای ارتباط اجتماعی چی بگم.', 6),
(24, '«الان» خیلی گرفته و خسته‌کننده‌م. نمی‌دونم برای ارتباط اجتماعی چی بگم.', 6),
(25, '«وقتی بچه بودم» اون کسی که من آرزوش رو داشتم اگه من واقعی رو با تموم عیب‌هام می‌دید نمی‌تونست عاشقم بشه.', 7),
(26, '«الان» اون کسی که من آرزوش رو دارم اگه من واقعی رو با تموم عیب‌هام ببینه نمی‌تونه عاشقم بشه.', 7),
(27, '«وقتی بچه بودم» از خودم خجالت می‌کشیدم، ارزش عشق، توجه و احترام دیگران رو نداشتم.', 7),
(28, '«الان» از خودم خجالت می‌کشم، ارزش عشق، توجه و احترام دیگران رو ندارم.', 7),
(29, '«وقتی بچه بودم» نوبت کار یا مدرسه که می‌رسید به‌اندازه بقیه باهوش یا توانا نبودم.', 8),
(30, '«الان» نوبت کار یا دانشگاه که می‌شه به اندازه بقیه باهوش یا توانا نیستم.', 8),
(31, '«وقتی بچه بودم» اغلب حس می‌کردم کافی نیستم چون به‌اندازه بقیه بااستعداد، باهوش و موفق نبودم.', 8),
(32, '«الان» اغلب حس می‌کنم کافی نیستم چون به‌اندازه بقیه بااستعداد، باهوش و موفق نیستم.', 8),
(33, '«وقتی بچه بودم» حس می‌کردم چاره‌ای ندارم جز این که به خواسته‌های دیگران تن بدم وگرنه تلافیش رو سرم در می‌آوردن یا یه طوری طردم می‌کردن.', 9),
(34, '«الان» حس می‌کنم چاره‌ای ندارم جز این که به خواسته‌های دیگران تن بدم وگرنه تلافیش رو سرم در میارن یا یه طوری طردم می‌کنن.', 9),
(35, '«وقتی بچه بودم» آدم‌ها من رو به‌عنوان کسی می‌دیدن که برای دیگران خیلی کارها انجام می‌دادم، اما به‌اندازه کافی برای خودم کاری نمی‌کردم.', 9),
(36, '«الان» آدم‌ها من رو به‌عنوان کسی می‌بینن که برای دیگران خیلی کارها انجام می‌دم، اما به‌اندازه کافی برای خودم کاری نمی‌کنم.', 9),
(37, '«وقتی بچه بودم» سعی می‌کردم تمام تلاشم رو کنم، نمی‌تونستم به به‌اندازه کافی خوب بودن قانع باشم. دوست داشتم توی کاری که انجام می‌دم بهترین باشم.', 10),
(38, '«الان» سعی می‌کنم تمام تلاشم رو کنم، نمی‌تونم به به‌اندازه کافی خوب بودن قانع باشم. دوست دارم توی کاری که انجام می‌دم بهترین باشم.', 10),
(39, '«وقتی بچه بودم» اونقدر کار برای انجام دادن داشتم که تقریبا هیچ وقتی برای استراحت و تفریح نداشتم.', 10),
(40, '«الان» اینقدر کار برای انجام دادن دارم که تقریبا هیچ وقتی برای استراحت و تفریح ندارم.', 10),
(41, '«وقتی بچه بودم» حس می‌کردم نباید مجبور باشم از قوانین و قواعد معمول که دیگران پیروی می‌کنن پیروی کنم.', 11),
(42, '«الان» حس می‌کنم نباید مجبور باشم از قوانین و قواعد معمول که دیگران پیروی می‌کنن پیروی کنم.', 11),
(43, '«وقتی بچه بودم» نمی‌تونستم خودم رو به تمام کردن کارهای تکراری، خسته کننده یا کنترل احساساتم متعهد کنم.', 11),
(44, '«الان» به‌نظر میاد نمی‌تونم خودم رو به تمام کردن کارهای تکراری، خسته کننده یا کنترل احساساتم متعهد کنم.', 11);

-- --------------------------------------------------------

--
-- Table structure for table `scores`
--

CREATE TABLE `scores` (
  `sid` int(11) NOT NULL,
  `telegram_id` bigint(64) NOT NULL,
  `score` int(11) NOT NULL,
  `qid` int(11) NOT NULL,
  `cid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `uid` int(11) NOT NULL,
  `telegram_id` bigint(64) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `tested` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`cid`);

--
-- Indexes for table `questions`
--
ALTER TABLE `questions`
  ADD PRIMARY KEY (`qid`);

--
-- Indexes for table `scores`
--
ALTER TABLE `scores`
  ADD PRIMARY KEY (`sid`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`uid`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `cid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `questions`
--
ALTER TABLE `questions`
  MODIFY `qid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=45;

--
-- AUTO_INCREMENT for table `scores`
--
ALTER TABLE `scores`
  MODIFY `sid` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `uid` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
