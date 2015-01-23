
CREATE TABLE `assignments` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `lecture_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `due_date` timestamp NOT NULL,
  `commit_hash` varchar(255) NOT NULL,
  `tex` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `comments` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ref_type` varchar(255) NOT NULL,
  `ref_id` int(11) unsigned NOT NULL,
  `ref_version` int(11) unsigned NOT NULL,
  `ref_line` int(11) unsigned NOT NULL,
  `parent_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `timestamp` timestamp NOT NULL,
  `text` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `lectures` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `university_id` int(11) unsigned NOT NULL,
  `name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `university_id` (`university_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `solutions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `assignment_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `commit_hash` varchar(255) NOT NULL,
  `tex` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `universities` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `versions` (
  `commit_hash` varchar(255) NOT NULL DEFAULT '',
  `ref_type` varchar(255) NOT NULL,
  `ref_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `version` int(11) unsigned NOT NULL,
  PRIMARY KEY (`commit_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
