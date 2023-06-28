-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Apr 15, 2023 at 04:13 PM
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
(1, 'Abandonment'),
(2, 'Mistrust and Abuse'),
(3, 'Vulnerability'),
(4, 'Dependence'),
(5, 'Emotional deprivation'),
(6, 'Social exclusion'),
(7, 'Defectiveness'),
(8, 'Failure'),
(9, 'Subjugation'),
(10, 'Unrelenting standards'),
(11, 'Entitlement');

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
(1, '[As a child] I find myself clinging to people i\'m close to because i\'m afraid they\'ll leave me.', 1),
(2, '[Now] I find myself clinging to people i\'m close to because i\'m afraid they\'ll leave me.', 1),
(3, '[As a child] I worry a lot that the people i love will find someone else they prefer and leave me.', 1),
(4, '[Now] I worry a lot that the people i love will find someone else they prefer and leave me.', 1),
(5, '[As a child] I am usually on the lookout for people\'s ulterior motives; I don\'t trust people easily.', 2),
(6, '[Now] I am usually on the lookout for people\'s ulterior motives; I don\'t trust people easily.', 2),
(7, '[As a child] I feel i cannot let my guard down around other people or they will hurt me', 2),
(8, '[Now] I feel i cannot let my guard down around other people or they will hurt me', 2),
(9, '[As a child] I worry more than the average person about danger that I will get sick or that some harm will come to me.', 3),
(10, '[Now] I worry more than the average person about danger that I will get sick or that some harm will come to me.', 3),
(11, '[As a child] I worry that I (or my family) will lose money and become destitute or dependent on others.', 3),
(12, '[Now] I worry that I (or my family) will lose money and become destitute or dependent on others.', 3),
(13, '[As a child] I do not feel I can cope well by myself, so I feel I need other people to help me get by.', 4),
(14, '[Now] I do not feel I can cope well by myself, so I feel I need other people to help me get by.', 4),
(15, '[As a child] My parents and I tend to overinvolved in each other\'s lives and problems.', 4),
(16, '[Now] My parents and I tend to overinvolved in each other\'s lives and problems.', 4),
(17, '[As a child] I have not had someone to nurture me, share him/herself with me, or care deeply about what happens to me.', 5),
(18, '[Now] I have not had someone to nurture me, share him/herself with me, or care deeply about what happens to me.', 5),
(19, '[As a child] People have not been there to meet my emotional needs for understanding, empathy, guidance, advice, and support.', 5),
(20, '[Now] People have not been there to meet my emotional needs for understanding, empathy, guidance, advice, and support.', 5),
(21, '[As a child] I feel like I do not belong. I am different. I do not really fit in.', 6),
(22, '[Now] I feel like I do not belong. I am different. I do not really fit in.', 6),
(23, '[As a child] I\'m dull and boring. I don\'t know what to say socially.', 6),
(24, '[Now] I\'m dull and boring. I don\'t know what to say socially', 6),
(25, '[As a child] No one I desire who knew the real me -with all my defects exposed- could love me.', 7),
(26, '[Now] No one I desire who knew the real me -with all my defects exposed- could love me.', 7),
(27, '[As a child] I am ashamed of myself, I am unworthy of the love, attention, and respect of others.', 7),
(28, '[Now] I am ashamed of myself, I am unworthy of the love, attention, and respect of others.', 7),
(29, '[As a child] I am not as intelligent or capable as most people when it comes to work (or school)', 8),
(30, '[Now] I am not as intelligent or capable as most people when it comes to work (or school)', 8),
(31, '[As a child] I often feel inadequate because I do not measure up to others in terms of talent, intelligence, and success.', 8),
(32, '[Now] I often feel inadequate because I do not measure up to others in terms of talent, intelligence, and success.', 8),
(33, '[As a child] I feel that I have no choice but to give in to other people\'s wishes; otherwise they will retaliate or reject me in some way.', 9),
(34, '[Now] I feel that I have no choice but to give in to other people\'s wishes; otherwise they will retaliate or reject me in some way.', 9),
(35, '[As a child] People see me as doing too much for others and not enough for myself.', 9),
(36, '[Now] People see me as doing too much for others and not enough for myself.', 9),
(37, '[As a child] I try to do my best, I can\'t settle for good enough. I like to be number one at what I do.', 10),
(38, '[Now] I try to do my best, I can\'t settle for good enough. I like to be number one at what I do.', 10),
(39, '[As a child] I have so much to accomplish that there is almost no time to relax and really enjoy myself.', 10),
(40, '[Now] I have so much to accomplish that there is almost no time to relax and really enjoy myself.', 10),
(41, '[As a child] I feel that I shouldn\'t have to follow the normal rules and conventions other people do.', 11),
(42, '[Now] I feel that I shouldn\'t have to follow the normal rules and conventions other people do.', 11),
(43, '[As a child] I can\'t seem to discipline myself to complete routine, boring tasks or to control my emotions.', 11),
(44, '[Now] I can\'t seem to discipline myself to complete routine, boring tasks or to control my emotions.', 11);

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
