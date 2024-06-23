# ************************************************************
# Sequel Ace SQL dump
# Version 20067
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: localhost (MySQL 8.0.36)
# Database: dating_app
# Generation Time: 2024-06-23 13:48:22 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `phone` varchar(16) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `url_photo` varchar(255) DEFAULT NULL,
  `date_birth` datetime DEFAULT NULL,
  `gender` enum('male','female','') DEFAULT NULL,
  `about_me` longtext,
  `instragram_url` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `subscription` enum('free','premium') DEFAULT 'free',
  `verify` enum('yes','no') DEFAULT 'no',
  `created_at` bigint DEFAULT NULL,
  `updated_at` bigint DEFAULT NULL,
  `status_like` longtext,
  PRIMARY KEY (`phone`),
  UNIQUE KEY `idx_user_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`phone`, `email`, `name`, `password`, `url_photo`, `date_birth`, `gender`, `about_me`, `instragram_url`, `city`, `country`, `subscription`, `verify`, `created_at`, `updated_at`, `status_like`)
VALUES
	('216-253-6879','Jean_Mosciski@hotmail.com','Margaret Sauer','$2a$14$csbm1pk5NdlvjJqi1ZbdT.fYdnBqFqKJvDs.yP4ZRSPrZTlNAtAQ6','https://example.com/photo.jpg','1989-06-01 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','North Conradton','Lebanon','free','no',1719114330,1719114330,NULL),
	('250-487-0294','Ben.Wilderman@hotmail.com','Beatrice Ankunding','$2a$14$OWHSMsdzRtALEor7yL4i2.9zSs8AjbSDxGjlmK0rBh0YGxxKb.QBi','https://example.com/photo.jpg','2000-04-28 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Lawrence','Saint Vincent and the Grenadines','free','no',1719114251,1719114251,NULL),
	('252-656-8153','Karlee_Casper@yahoo.com','Jesus Greenholt','$2a$14$25fi0HHir98pgVmq/UOM/.JsdOZFbVjlCtd0ms63rQENb/qd3zKaq','https://example.com/photo.jpg','1988-08-19 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Mission Viejo','Bhutan','free','no',1719114269,1719114269,NULL),
	('254-937-6042','Lia9@hotmail.com','Joanne Hegmann','$2a$14$FwGdKeG5A4gyjZVpnw/mpe4oMkr2oULHUcF6kazrikvMtCvfHZgt2','https://example.com/photo.jpg','1987-12-17 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Weberport','Fiji','free','no',1719114369,1719114369,NULL),
	('255-236-4436','Violet_Wuckert38@gmail.com','Stanley Witting','$2a$14$BC4xl9fozvoUElETe92/bue8KmOEgmDD6nG0EK4yGHCPy6jvofgY6','https://example.com/photo.jpg','1992-01-31 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','New Jocelyn','New Zealand','free','no',1719114343,1719114343,NULL),
	('285-420-7737','Ewald_OKon@gmail.com','Sergio Zboncak','$2a$14$RUoyPuAjCDQDTBZkgjJoz.C6mczZOZhzlnvO9UGbgLQHPhEFZjU5W','https://example.com/photo.jpg','1982-10-03 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Port Kallieberg','Slovakia (Slovak Republic)','free','no',1719114337,1719114337,NULL),
	('392-398-5548','Colleen8@gmail.com','Darren Baumbach','$2a$14$9lkXXhf.2lkGdcxfJoTEaOOBdq9zuMTGuD6.phgLvkl0xoGwZcFd6','https://example.com/photo.jpg','1991-12-23 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','West Jamarcusshire','Portugal','free','no',1719114292,1719114292,NULL),
	('419-560-8134','Seth96@gmail.com','Nelson Becker','$2a$14$U4inZGubgOSc0oCAZVZZKuW9bbjXdRAPYWQ/lakjvY/tCl.wHHxcu','https://example.com/photo.jpg','1987-10-06 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Lake Taureanland','Madagascar','free','no',1719114302,1719114302,NULL),
	('420-868-3296','Lucio.Hegmann@gmail.com','Teresa McDermott','$2a$14$TqSkxRauf8qIHzMC981k2u/yumVWmgfOAqBIge0oEbtDnjn4nURVO','https://example.com/photo.jpg','1997-10-29 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Arecibo','Qatar','free','no',1719114308,1719114308,NULL),
	('440-749-7824','Wilburn_Stark@yahoo.com','Emmett Carroll','$2a$14$zo1ISWOPnorqYqA0.2lp4e83Q9kDgmON1cIBuEMCfIYRQtQhg7mXe','https://example.com/photo.jpg','2004-06-14 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Port Kristoffer','Norway','free','no',1719114295,1719114295,NULL),
	('449-452-7144','Dessie52@yahoo.com','Peggy King','$2a$14$LBwxu9CxyUsyqmwRFM/qa.28RM0GrAiujP03dd70P3w0Wx7W89GJO','https://example.com/photo.jpg','1990-03-05 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','South Santiago','Iceland','free','no',1719114340,1719114340,NULL),
	('452-521-8482','Marie13@yahoo.com','Willie Hilll','$2a$14$UX3wHo.X8Ug8DdRXshW.tuTZQ4cej9e/uSIcxtXKxr4Mx66mm1SJK','https://example.com/photo.jpg','1994-08-19 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','New Germaineville','Cocos (Keeling) Islands','free','no',1719114324,1719114324,NULL),
	('525-251-0238','Jevon_Schulist@yahoo.com','Dean Treutel Jr.','$2a$14$y5KvOQ3Trsepa9h7rkT0I.WpAtpSEYr3BeA9NTIZB/iwHY.bgDlAO','https://example.com/photo.jpg','1998-12-09 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','West Keenan','Uganda','free','no',1719114260,1719114260,NULL),
	('527-490-5267','Duncan76@gmail.com','Bonnie King','$2a$14$hnp48mtpW30ole4leo5PoOhxThlWMCLdmgJ8oWyHCqUtnvFU0Jmoy','https://example.com/photo.jpg','2003-10-06 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','New Kelly','Anguilla','free','no',1719114300,1719114300,NULL),
	('541-368-1532','Bennett55@gmail.com','Gladys Boehm','$2a$14$rBjDkwchKo4urs7Nb614JuOkfwlagvp8ZgGnJ6yctxYY1Gn3zBqSO','https://example.com/photo.jpg','1998-02-10 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Newton','Bahamas','free','no',1719114326,1719114326,NULL),
	('570-221-8526','Alfonzo36@hotmail.com','Ralph Nitzsche','$2a$14$gYXVXtoLO3Fw4Z8GHMryk.Oh6wbZgjeTSO.eRgoObyjPZei24XIrK','https://example.com/photo.jpg','2000-07-18 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Chesleymouth','Saint Lucia','free','no',1719114286,1719114286,NULL),
	('575-543-6150','Kobe.Pfannerstill@yahoo.com','Arlene Hayes','$2a$14$UTvSusTxDXY5Mxm2Ex/eWeL2qyqBbWT3WiJhJ1wVYng3Emel.cpPu','https://example.com/photo.jpg','1995-03-14 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','East Domenica','Dominica','free','no',1719114311,1719114311,NULL),
	('605-231-4755','Meda97@gmail.com','Dewey Schmitt','$2a$14$ryLTxV.u9JTzdK1I6Njgg.0KLW5r7ylX/B1ok1UHJec8vwtKU.AZK','https://example.com/photo.jpg','1983-10-08 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Champlintown','Republic of Korea','free','no',1719114290,1719114290,NULL),
	('620-603-9996','Zelma76@hotmail.com','Debra Brekke','$2a$14$5rOxdnuhgdgka45jAPtkT.HU8Kz83bzHQ59S3JbvZnqTk6dGtTaDO','https://example.com/photo.jpg','1986-04-21 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Towson','Kenya','free','no',1719114313,1719114313,NULL),
	('629-287-5482','Velva_MacGyver@gmail.com','Ramon Dicki','$2a$14$LXMxVoI.6XXvuvrcGTxT9On2QxJucFEVvRFvQxxhaTt0sMuywUGEm','https://example.com/photo.jpg','1983-01-08 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Mitchellmouth','Iceland','free','no',1719114278,1719114278,NULL),
	('672-973-1546','Eloise41@gmail.com','Shawna Hayes','$2a$14$4rI5Dr53Gv9tIe18HCV4zebHO5hAUCCDWGuLHhFzggjK/BnjWyb0G','https://example.com/photo.jpg','1983-09-29 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','West Filibertoside','Paraguay','free','no',1719114305,1719114305,NULL),
	('676-420-3809','Lizeth41@gmail.com','Rodney Ruecker','$2a$14$d6Hi9edyi0WHNj3jL45pluBDsh8oIq9riAbBMNUMrJ54aPy5VCvmK','https://example.com/photo.jpg','1980-07-01 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','South Keltonshire','Pitcairn Islands','free','no',1719114274,1719114274,NULL),
	('676-823-0101','Danny_Erdman86@hotmail.com','Elmer Keefe','$2a$14$7e42hu9Ct5NSJ3ITSCj8D.lE5nQYpsPVsE7P298nb/C/i0pLc1M2y','https://example.com/photo.jpg','1983-09-29 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','South Bend','Malaysia','free','no',1719114332,1719114332,NULL),
	('715-953-4984','Elise_Corkery@gmail.com','Nicole Heathcote I','$2a$14$RSV0pdc3GgLwuFIzFptXAOVUEJOJdogH33.vh7GvBYao3k4C6qSp.','https://example.com/photo.jpg','1991-09-05 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Lake Myrtle','Trinidad and Tobago','free','no',1719114321,1719114321,NULL),
	('780-202-7121','Novella79@yahoo.com','Alan Daniel','$2a$14$2H2p3ryqhTlcvWSiRAwHOuks8FVKqVsLzfq6D.SA/CLJflpBV3FgW','https://example.com/photo.jpg','1991-08-04 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Wilmington','Maldives','free','no',1719114372,1719114372,NULL),
	('801-282-0959','Nathan91@gmail.com','Salvador Treutel','$2a$14$SnGLcZAPaA7MngqgC0eaBO9Pt/MB9pNHyE09I5iWf8hUq/uTrJHH6','https://example.com/photo.jpg','1982-06-16 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Hilpertburgh','Tonga','free','no',1719114258,1719114258,NULL),
	('828-449-5010','Woodrow.Carroll40@yahoo.com','Patti Walker','$2a$14$hb.KJ/mHB8KwqTQI.vutLefBYqtLvYXLUNWAzpYofOMA7brXfvxjG','https://example.com/photo.jpg','1988-09-21 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Fosterfurt','Colombia','free','no',1719114271,1719114271,NULL),
	('852-733-8542','Chelsea.Dare78@gmail.com','Alfred Smitham','$2a$14$y/KWr1WPz6.9YI11VeDJDeGp2tDRnzdM5yeUMl8AZDOThQv7oPiBW','https://example.com/photo.jpg','1987-04-18 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Vadastad','Greece','free','no',1719114284,1719114284,NULL),
	('885-721-7683','Tracy54@yahoo.com','Corey Bauch I','$2a$14$gUbmKtc244l.qI1H5AUDcuHcLAIChNN5UWAAu50s.XljEqKXKyDO2','https://example.com/photo.jpg','2003-01-09 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','North Nikkostad','Senegal','free','no',1719114335,1719114335,NULL),
	('905-696-5271','Mariane87@hotmail.com','Marcos Rodriguez','$2a$14$T6rI1/WLrl1HhLUpXZyN/.aScYhSxBqlVYdPdGun2VmMbUUPLkGB6','https://example.com/photo.jpg','1981-09-17 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','North Roel','Ethiopia','free','no',1719114281,1719114281,NULL),
	('906-551-5213','Katlyn_Blick39@gmail.com','Gretchen Reilly DVM','$2a$14$VzJj3NQTqF55IVHdZTuPgOBwVjNEev/ivVROkAHD934TovkRVxym6','https://example.com/photo.jpg','1992-05-05 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Mohammedbury','United States of America','free','no',1719114319,1719114319,NULL),
	('910-971-5250','Kale79@yahoo.com','Antonio Jast','$2a$14$y0FPY6JY.QJt3tTE4E8GCehxyrn8oij3QmCGyR/R5eI2hfJJihG1C','https://example.com/photo.jpg','1997-01-20 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Metachester','Kuwait','free','no',1719114254,1719114254,NULL),
	('912-494-0835','Kiera45@yahoo.com','Peter Reinger','$2a$14$Vb5jp3DbTfv9DoNBH7Z16uBIh2HhYlgRaHD7jYSVI7pMAF.cG03c2','https://example.com/photo.jpg','1987-04-28 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','West Jerel','Vietnam','free','no',1719114276,1719114276,NULL),
	('955-761-8827','Zakary_Leuschke@yahoo.com','Edna Tillman MD','$2a$14$XKvRUGLaspREZ02XTSzAq.QaBSNWBPUXm/C9fZT9Q88yPAPm1lQuy','https://example.com/photo.jpg','1999-03-27 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Kokomo','South Africa','free','no',1719114297,1719114297,NULL),
	('957-693-8346','Perry6@yahoo.com','Nancy Barrows','$2a$14$Kq.K8hye8QPWvBWaKgIUY.FL/fiC4ueowCtRhtwGxV6XKx1xE3XHG','https://example.com/photo.jpg','2003-01-16 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Aylabury','Lebanon','free','no',1719114315,1719114315,NULL),
	('994-739-5650','Hershel_Lockman@yahoo.com','Irene Blick','$2a$14$BFqIbo2xgcDUmxEWycYKT.p7SSKBEJAQ.ca4KZWFFRadvYUxwYICy','https://example.com/photo.jpg','1987-12-02 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','New Estevanstad','Malta','free','no',1719114256,1719114256,NULL),
	('999-505-6260','Zoie55@gmail.com','Robin Bauch','$2a$14$XWjY2ZpRBws0Om8krnNZy.CqtBN193aQIVeEBXeiMoaBB2Oj9CFJO','https://example.com/photo.jpg','1986-04-09 00:00:00','female','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Port Yasmeenland','El Salvador','free','no',1719114263,1719114263,NULL),
	('999-844-3312','Landen.Swaniawski@gmail.com','Tami Kemmer MD','$2a$14$/7lBOIKCxE7na23ZrF/igesPJRQiJqUZfC4YMpSjoNGfXYKC9CtSe','https://example.com/photo.jpg','2000-02-07 00:00:00','male','Deleniti aut reiciendis id vitae corrupti. Perferendis blanditiis consequatur.','','Gerlachhaven','Faroe Islands','free','no',1719114375,1719114375,NULL);

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user_liked
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_liked`;

CREATE TABLE `user_liked` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `phone` varchar(16) DEFAULT NULL,
  `phone_liked` varchar(16) DEFAULT NULL,
  `created_at` bigint DEFAULT NULL,
  `updated_at` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_user_liked` (`phone`),
  KEY `fk_user_user_liked_phone` (`phone_liked`),
  CONSTRAINT `fk_user_user_liked` FOREIGN KEY (`phone`) REFERENCES `user` (`phone`),
  CONSTRAINT `fk_user_user_liked_phone` FOREIGN KEY (`phone_liked`) REFERENCES `user` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
