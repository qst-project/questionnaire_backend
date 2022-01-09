CREATE TABLE Users
(
    google_id SERIAL not null unique,
);

CREATE TABLE Questionnaire
(
    id SERIAL not null unique,
    title varchar(255) not null,
    creator_id int references Users (google_id) on delete cascade not null,
);

CREATE TABLE Question
(
    id int SERIAL not null,
--     unique???
    order int not null,
    questionnaire int references Questionnaire (id) on delete cascade not null,
--     как будем хранить типы? мб 0,1,2 = radio, checkbox, text. Или в стринге
    type int not null,
    text varchar (255) not null,
);

CREATE TABLE RadioPossibleAnswer
(
    id SERIAL not null unique,
    question int references  Question (id) on delete cascade not null,
    text varchar (255) not null ,
);

CREATE TABLE RadioAnswer
(
    id SERIAL not null unique,
    radioPossibleAnswer int references RadioPossibleAnswer (id) in delete cascade not null,
    user_id int,
);

CREATE TABLE CheckboxPossibleAnswer
(
    id SERIAL not null unique,
    question int references  Question (id) on delete cascade not null,
    text varchar (255) not null,
);

CREATE TABLE CheckboxAnswer
(
    id SERIAL not null unique,
    checkboxPossibleAnswer int references CheckboxPossibleAnswer(id),
    user_id int,
);

CREATE TABLE TextPossibleAnswer
(
    id SERIAL not null unique,
    question int references  Question (id) on delete cascade not null,
    text varchar (255) not null,
    placeholder varchar (255),
);

CREATE TABLE TextAnswer
(
  textPossibleAnswer int references TextPossibleAnswer(id) on delete cascade not null,
  user_id int,
  answer varchar (255) not null,
);