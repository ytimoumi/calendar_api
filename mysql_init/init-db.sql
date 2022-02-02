CREATE DATABASE IF NOT EXISTS `giskard_db` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE `giskard_db`;
-- CREATE USER 'sf_user' IDENTIFIED BY 'UzTQ4356NffpN3i9';
CREATE DATABASE IF NOT EXISTS `giskard_db`;
-- GRANT ALL PRIVILEGES ON giskard_db.* TO `sf_user`;
USE giskard_db;
--
-- Hôte : mysql_giskard

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `giskard_db`
--

-- --------------------------------------------------------

--
-- Structure de la table `availabilities`
--

CREATE TABLE `availabilities` (
                                  `id` int(11) NOT NULL,
                                  `start` datetime NOT NULL,
                                  `end` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Structure de la table `reservations`
--

CREATE TABLE `reservations` (
                                `id` int(11) NOT NULL,
                                `start` datetime NOT NULL,
                                `end` datetime NOT NULL,
                                `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `availabilities`
--
ALTER TABLE `availabilities`
    ADD PRIMARY KEY (`id`);

--
-- Index pour la table `reservations`
--
ALTER TABLE `reservations`
    ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `availabilities`
--
ALTER TABLE `availabilities`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `reservations`
--
ALTER TABLE `reservations`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

