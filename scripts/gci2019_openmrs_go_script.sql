CREATE TABLE hospital
(
    id                 SERIAL PRIMARY KEY,
    name               VARCHAR(100),
    max_patient_amount INT
);

CREATE TABLE location
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100),
    hospital_id INT,
    FOREIGN KEY (hospital_id) REFERENCES hospital (id)
);

CREATE TABLE patient
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(100),
    illness        VARCHAR(200),
    birth_date     DATE,
    location_id    INT NOT NULL,
    FOREIGN KEY (location_id) REFERENCES location (id)
);

INSERT INTO hospital (name, max_patient_amount)
VALUES ('Local Hospital', 2000),
       ('24 hours hospital', 1000);

INSERT INTO location (name, hospital_id)
VALUES ('ICU ward', 1),
       ('Lab', 2),
       ('Pharmacy', 1),
       ('daycare ward', 2);

INSERT INTO patient (name, illness, birth_date, location_id)
VALUES ('krish', 'Fever', '2002-03-11', 2),
       ('krishna', 'Cold', '2002-12-17', 3),
       ('Gokul', 'AIDS', '1978-06-04', 3),
       ('Girish', 'Malaria', '1991-09-29', 1),
       ('Dinesh', 'Asthama', '1998-05-26', 4),
       ('Rohit', 'brain fever', '2002-11-21', 2),
       ('Sachin', 'Dengue', '1993-06-08', 4),
       ('Dhoni', 'Common Cold', '2005-05-05', 1),
       ('Coby', 'jaundice', '2006-06-05', 2);